package controllers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	// postgres database driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Server is for grouping the global db variables
type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

// Init here, is the core function which initializes connection to the db
func (server *Server) Init(DbName, DbUser, DbPass, DbType, DbHost, DbPort string) {
	var err error

	// This is the format used to connect to Postgres via gorm
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPass)

	// Passing the above string to establish connection
	server.DB, err = gorm.Open("postgres", DBURL)

	if err != nil {
		fmt.Println("Could not connect to the Postgres Database")
		log.Fatal("Error: ", err)
	} else {
		fmt.Println("Connection to Postgres Database established.")
	}

	// Models to be placed in automigrate() params
	server.DB.Debug().AutoMigrate()
}

// Run starts the backend server and configures html rendering
func (server *Server) Run(Port string) {
	server.Router = gin.New()

	// Load HTML and Static files
	server.Router.LoadHTMLGlob("templates/*.html")
	server.Router.Static("/css", "templates/css")

	server.initRoutes()

	// Running the server
	fmt.Printf("Listening to port %s", Port)

	server.Router.Run(Port)
}
