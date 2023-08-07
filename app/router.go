package app

import (
	"github.com/gin-gonic/gin"
	"go-pzn-restful-api/controller"
	"go-pzn-restful-api/middleware"
	"go-pzn-restful-api/repository"
	"go-pzn-restful-api/service"
)

func UserControllerInit() controller.UserController {
	db := DBConnection()
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	return userController
}

func NewRouter() *gin.Engine {
	userController := UserControllerInit()

	router := gin.Default()
	router.Use(gin.CustomRecovery(middleware.ErrorHandler))

	v1 := router.Group("/api/v1")

	v1.POST("/users", userController.Register)
	v1.GET("/users/:userID", userController.GetByID)
	v1.POST("/users/login", userController.Login)

	return router
}
