Golang PoC 

# Database - PostgreSQL
### Install Go package migrate for migrations - CLI Mode
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
    asdf reshim

### Database model for this project
[ERD Model](https://eli.thegreenplace.net/images/2021/mooc-dbschema.png)

### Create migrations
    make db-migration name=your_migration_name

### Run migrations
    make db-migrate

### Create seed files
    make db-seed-file seed_name=your_seed_file_name_here
    
### Seed data
    make db-seed

# Go Linters
### [golangci-lint](https://golangci-lint.run/)
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.58.1
    golangci-lint run ./...

### [staticcheck](https://staticcheck.dev/)
    go install honnef.co/go/tools/cmd/staticcheck@latest
    staticcheck ./...

### [go-critic](https://go-critic.com/)
    go install -v github.com/go-critic/go-critic/cmd/gocritic@latest
    gocritic check ./...

### Govet
    go vet ./...

### Gofmt
    go fmt ./...