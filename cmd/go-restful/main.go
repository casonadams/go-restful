package main

import (
	"fmt"
	"github.com/casonadams/go-restful/internal/pkg/controller"
	"github.com/casonadams/go-restful/internal/pkg/database"
	"github.com/casonadams/go-restful/internal/pkg/middleware"
	"github.com/casonadams/go-restful/internal/pkg/model"
	"github.com/gin-gonic/gin"
)

func main() {
	loadDatabase()
	serveApplication()
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Entry{})
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/entry", controller.AddEntry)
	protectedRoutes.GET("/entry", controller.GetAllEntries)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}
