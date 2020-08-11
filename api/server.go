package api

import (
	"log"
	"os"
	"teastore/api/controllers"

	"github.com/joho/godotenv"
)

var server = controllers.Server{}

// Run call db init and starts the server at defined port
func Run() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Couldn't access environmental variables")
	}

	server.Init(os.Getenv("db_name"), os.Getenv("db_user"), os.Getenv("db_pass"), os.Getenv("db_type"), os.Getenv("db_host"), os.Getenv("db_port"))

	server.Run(os.Getenv("port"))
}
