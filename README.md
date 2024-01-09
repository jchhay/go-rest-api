# go-rest-api-gin

## Summary

REST API implemented using the Gin Web Framework for performance and productivity. The Data Layer provides multiple data store options to showcase extensibility. GORM is also implemented to provide ORM to the database.
Configuration is managed by viper for reading config.yml file and/or environment variables (takes precedence over config file) as a means to accommodate 12 Factor App methodology. The project structure itself utilizes a clean architecture approach to separate areas of concerns. 

## Getting started

- [ ] Ensure to complete all the subsequent Setup instructions below if applicable
- [ ] Modify the configuration as needed:

```
# Database and Server configuration
config/config.yml (default database is memory, can be changed to sqlite)

# Database and Server configuration (overrides config.yml)
.env
```
- [ ] Execute the following on root:

```
go mod tidy
go run .
```
- [ ] Visit Swagger URL to test the API: http://localhost:3000/swagger/index.html


## Setup GoDotEnv - environment variable tool

- [ ] Using the following libraries:

```
go get github.com/joho/godotenv
```

- [ ] Create file:
```
.env
```

## Setup viper - configuration solution

- [ ] Using the following libraries:

```
go get github.com/spf13/viper
```

## Setup SQLITE - environment variable tool

- [ ] Using the following libraries:

```
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite

```

- [ ] For Windows, read the gcc compiler instructions to resolve CGO_ENABLED error:
```
https://pkg.go.dev/github.com/mattn/go-sqlite3#section-readme
```

## Setup Swagger - RESTful API doc generation

- [ ] Using the following libraries:

```
# Get swag
go install github.com/swaggo/swag/cmd/swag@latest
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files

# Generate docs
# If main is in root:
swag init 
# If main is in a sub directory:
swag init --dir <dir> --parseDependency --parseDependencyLevel 3 --parseInternal --output docs
```
- [ ] Test Swagger UI:

[Swagger UI URL](http://localhost:3000/swagger/index.html)


