package routes

import (
	"example.com/rest-api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	// This is our endpoint

	server.GET("/events", getEvents)    // GET, POST, PUT, PATCH, DELETE
	server.GET("/events/:id", getEvent) // events/1 or events/2 or events/3

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/events", createEvent)

	authenticated.PUT("/events/:id", updateEvent)

	authenticated.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signup)

	server.POST("/login", login)

}
