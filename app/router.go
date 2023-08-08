package app

import (
	"github.com/gin-gonic/gin"
	"go-pzn-restful-api/auth"
	"go-pzn-restful-api/controller"
	"go-pzn-restful-api/middleware"
	"go-pzn-restful-api/repository"
	"go-pzn-restful-api/service"
)

var (
	jwtAuth        = auth.NewJwtAuth()
	db             = DBConnection()
	userRepository = repository.NewUserRepository(db)
	userService    = service.NewUserService(userRepository, jwtAuth)
)

func userControllerInit() controller.UserController {
	return controller.NewUserController(userService)
}

func authorControllerInit() controller.AuthorController {
	authorRepository := repository.NewAuthorRepository(db)
	authorService := service.NewAuthorService(authorRepository)

	return controller.NewAuthorController(authorService)
}

func NewRouter() *gin.Engine {
	userController := userControllerInit()
	authorController := authorControllerInit()

	router := gin.Default()
	router.Use(gin.CustomRecovery(middleware.ErrorHandler))

	v1 := router.Group("/api/v1")

	// User endpoints
	v1.POST("/users", userController.Register)
	v1.POST("/users/login", userController.Login)
	v1.POST("/users/logout", middleware.JwtAuthMiddleware(jwtAuth, userService), userController.Logout)
	v1.GET("/users", middleware.JwtAuthMiddleware(jwtAuth, userService), userController.GetByID)
	v1.POST("/users/avatars", middleware.JwtAuthMiddleware(jwtAuth, userService), userController.UploadAvatar)

	// Author endpoints
	v1.POST("/authors", authorController.Create)

	return router
}
