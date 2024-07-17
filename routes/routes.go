package routes

import (
	"example.com/app/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents) //fetches all events
	server.GET("/users", getUsers)
	server.GET("/events/:id", getEvent)  //fetch a single event

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("events/:id/register", registerForEvent)
	authenticated.DELETE("events/:id/register", cancelRegistration)
	 
	server.POST("/signup", signUp) //create new user
	server.POST("/login", logIn)

}