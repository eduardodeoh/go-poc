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