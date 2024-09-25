package api

import (
	"github.com/gin-gonic/gin"
	"go-backend-portfolio/controllers" // Controllers'ni import qilish
)

func SetupRouter(userController *controllers.UserController) *gin.Engine {
	router := gin.Default()

	// Statik fayllar uchun yo'nalish (masalan, favicon)
	router.Static("/public", "./public")

	// Root yo'nalish uchun javob
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Go Backend Portfolio!",
		})
	})

	// Boshqa API yo'nalishlari
	router.POST("/users", userController.CreateUser)

	return router
}
