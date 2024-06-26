SHELL=/bin/bash

#===========================#
#== Environment Variables ==#
#===========================#

# Project's environment variables
ENV := $(PWD)/.env
include $(ENV)

# Export all variables to sub-make
export

#================================#
#== Makefile Default Variables ==#
#================================#
docker_compose_file ?= docker-compose.yml
migration_name = migration_name
seed_name = seed_name
db_url = postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}
db_migrations_dir = ${PWD}/internal/infra/postgresql/migrations
db_seeds_dir = ${PWD}/internal/infra/postgresql/seeds


#==============#
#== DATABASE ==#
#==============#

db-create: ## Create a database
db-create:
	@echo "Creating database $(DB_NAME)..."
	docker-compose -f ${docker_compose_file} exec db createdb -h localhost -U postgres $(DB_NAME)

db-drop: ## Drop a database
db-drop:
	@echo "Dropping database $(DB_NAME)..."
	docker-compose -f ${docker_compose_file} exec db dropdb -h localhost -U postgres --if-exists $(DB_NAME)

db-shell: ## Access database console
db-shell:
	@echo "Connecting to database console..."
	docker compose -f ${docker_compose_file} exec db psql -U postgres -d $(DB_NAME)

db-reset: ## Reset the database
db-reset:
	@$(MAKE) name=$(DB_NAME) db-drop
	@$(MAKE) name=$(DB_NAME) db-create
	@$(MAKE) db-migrate

#========================#
#== DATABASE MIGRATION ==#
#========================#
name ?= ${migration_name}

db-migrate: ## Run migrations
db-migrate:
	migrate -database ${db_url} -path ${db_migrations_dir} up

db-rollback: ## Rollback migrations
db-rollback:
	migrate -database ${db_url} -path ${db_migrations_dir} down 1

db-migration: ## Create database migration file. e.g `make db-migration name=put_your_migration_name_here`
db-migration:
	migrate create -ext sql -dir ${db_migrations_dir} $(name)
	

#========================#
#== DATABASE SEED DATA ==#
#========================#
db-seed-file: ## Creates a new seed file with using the provided name. e.g `make db-seed-file name=put_your_seed_file_name_here`
db-seed-file:
	@echo "Creating seed data file"
	@echo "--Seed files MUST be idempotent by design" > $(db_seeds_dir)/$$(date +%Y%m%d%H%M%S)_$(seed_name).sql


sql_files = $(wildcard $(db_seeds_dir)/*.sql)

db-seed: ## Seeds database
db-seed:
	@$(foreach file,$(sql_files), echo -e "\n Seeding $(file)"; psql -f $(file) $(db_url);)


#==================#
#== CODE QUALITY ==#
#==================#
code-quality: ## Run Go linters to check code quality
code-quality:
	@echo "Checking code quality..."
	golangci-lint run ./...


#==================#
#== CODE TESTING ==#
#==================#
test-code: ## Run Go tests
test-code:
	@echo "Testing code..."
	go test -v ./...

test-cover: ## Calculate test coverage
test-cover:
	@echo "Generating code coverage..."
	go test -coverprofile=coverage.out -v ./...

test-cover-report-html: ## Show generated test coverage
test-cover-report-html:
	go tool cover -html=coverage.out

test-cover-report-cli: ## Show generated test coverage
test-cover-report-cli:
	go tool cover -func=coverage.out