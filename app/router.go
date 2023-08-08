package app

import (
	"github.com/gin-gonic/gin"
	"go-pzn-restful-api/auth"
	"go-pzn-restful-api/controller"
	"go-pzn-restful-api/middleware"
	"go-pzn-restful-api/repository"
	"go-pzn-restful-api/service"
)

func NewRouter() *gin.Engine {
	// User
	db := DBConnection()
	jwtAuth := auth.NewJwtAuth()
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, jwtAuth)
	userController := controller.NewUserController(userService)

	router := gin.Default()
	router.Use(gin.CustomRecovery(middleware.ErrorHandler))

	v1 := router.Group("/api/v1")

	v1.POST("/users", userController.Register)
	v1.POST("/users/login", userController.Login)
	v1.POST("/users/logout", middleware.JwtAuthMiddleware(jwtAuth, userService), userController.Logout)
	v1.GET("/users", middleware.JwtAuthMiddleware(jwtAuth, userService), userController.GetByID)
	v1.POST("/users/avatars", middleware.JwtAuthMiddleware(jwtAuth, userService), userController.UploadAvatar)

	return router
}
