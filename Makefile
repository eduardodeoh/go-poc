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
	


