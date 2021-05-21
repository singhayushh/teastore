package api

import (
	"log"
	"os"
	"teastore/api/controllers"

	"github.com/joho/godotenv"
)

var server = controllers.Server{}

// Connects to db and starts the server at defined port, requires .env to be present in root directory.
func Run() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn't access environmental variables")
	}

	// Call InitDB from ./controllers/config.go
	server.InitDB(os.Getenv("db_name"), os.Getenv("db_user"), os.Getenv("db_pass"), os.Getenv("db_type"), os.Getenv("db_host"), os.Getenv("db_port"))

	// Call InitServer from ./controllers/config.go
	server.InitServer(os.Getenv("port"))
}
