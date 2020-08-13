package controllers

import (
	"fmt"
	"log"
	"teastore/api/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	// postgres database driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Server is for grouping the global variables
type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

// InitDB is the core function which initializes connection to the db
func (server *Server) InitDB(DbName, DbUser, DbPass, DbType, DbHost, DbPort string) {
	var err error

	// This is the format used to connect to Postgres via gorm
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPass)

	// Passing the above string to establish connection
	server.DB, err = gorm.Open(DbType, DBURL)

	if err != nil {
		fmt.Println("Could not connect to the Postgres Database")
		log.Fatal("Error: ", err)
	} else {
		fmt.Println("Connection to Postgres Database established.")
	}

	// Models to be placed in automigrate() params
	server.DB.Debug().AutoMigrate(&models.User{})

}

// InitServer starts the backend server and configures html rendering
func (server *Server) InitServer(Port string) {
	gin.SetMode(gin.ReleaseMode)
	server.Router = gin.New()

	// GOTO routes.go/initRoutes()
	server.initRoutes()

	// Load HTML and Static files
	server.Router.LoadHTMLGlob("templates/*.html")
	server.Router.Static("/css", "templates/css")

	// Running the server
	fmt.Printf("Listening to port %s", Port)
	server.Router.Run(Port)
}
