package main

import (
	"go-backend-portfolio/api"
	"go-backend-portfolio/config"
	"go-backend-portfolio/controllers"
	"go-backend-portfolio/repository"
	"go-backend-portfolio/services"
)

func main() {
	db := config.ConnectDatabase()

	// User repository, service va controllerni yaratish
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// Routerni sozlash
	router := api.SetupRouter(userController)

	// Serverni ishlatish
	router.Run(":8080")
}
