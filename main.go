package main

// use the gin package for the rest api
import (
	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()

	// Setting up the engine and behind the scenes it creates HTTP server
	// It has some basic functionality

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080

}
