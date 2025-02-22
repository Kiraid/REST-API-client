package main

import (
	"example.com/app/db"
	"example.com/app/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	db.CreateTable()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}


