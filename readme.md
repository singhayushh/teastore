

# About the project development

## Instructions to run

1. Install Postgres 13, Go 1.13+, Redis-Server
2. Create postgres user, password, database for the project
3. Open .env file and update all required details
4. Make sure postgres and redis services are running
5. Hit go run main.go in the terminal

## Other Details

<details> <summary>Expand</summary>

> Some parts forked from @mindinvertory under MIT License

## Short definitions of packages and modules used

-   gorm : It is the ORM library in Go which provides user friendly functions to interact with database. It supports features like ORM, Associations, Hooks, Preloading, Transaction, Auto Migration, Logger etc.
-   gin : Gin is a web framework for Go language. Here gin is used for increase performance and productivity.
-   godotenv : Basically used for load env variables from .env file.
-   fresh : It is a live reloader for go-lang, handy for developers when making frequent changes to the code.

## Current Working Directory

The root directory of this repository is named `teastore` having the following path:

`$GOPATH/src/github.com/<username>/teaStore/`

## Running the project

If you have [Fresh](https://github.com/gravityblast/fresh) installed, simple type `fresh` in console, else `go run main.go`.

To get rid of the debugger warning logs, set mode to `release` in the `.env` file

## Components used generally

### 1. Api helpers

Basically contains the helper functions used in returning api responses, HTTP status codes, default messages etc.

### 2. Controllers

Contains handler functions for particular route to be called when an api is called.

### 3. Helpers

Contains helper functions used in all apis

### 4. Middlewares

Middleware to be used for the project

### 5. Models

Database tables to be used as models struct

### 6. Resources

Resources contains all structures other than models which can be used as responses

### 7. Routers

Resources define the routes for your project

### 8. Seeder

It is optional, but if you want to insert lots of dummy records in your database, then you can use seeder.

### 9. Services

All the core apis for your projects should be within services.

### 10. Storage

It is generally for storage purpose.

### 11. Templates

Contains the HTML templates used in your project

### 12. .env

Contains environment variables.

</details>