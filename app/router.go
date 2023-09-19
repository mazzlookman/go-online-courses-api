package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-pzn-restful-api/auth"
	"go-pzn-restful-api/controller"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/middleware"
	"go-pzn-restful-api/repository"
	"go-pzn-restful-api/service"
)

var (
	jwtAuth = auth.NewJwtAuth()
	db      = DBConnection()

	// user
	userRepository = repository.NewUserRepository(db)
	userService    = service.NewUserService(userRepository, jwtAuth)
	userController = controller.NewUserController(userService)

	// author
	authorRepository = repository.NewAuthorRepository(db)
	authorService    = service.NewAuthorService(authorRepository, jwtAuth)
	authorController = controller.NewAuthorController(authorService)

	// course
	courseRepository = repository.NewCourseRepository(db)
	courseService    = service.NewCourseService(courseRepository)
	courseController = controller.NewCourseController(courseService)

	// category
	categoryRepository = repository.NewCategoryRepository(db)
	categoryService    = service.NewCategoryService(categoryRepository)
	categoryController = controller.NewCategoryController(categoryService)

	// lesson_title
	lessonTitleRepository = repository.NewLessonTitleRepository(db)
	lessonTitleService    = service.NewLessonTitleService(lessonTitleRepository, courseService)
	lessonTitleController = controller.NewLessonTitleController(lessonTitleService)

	// lesson_content
	lessonContentRepository = repository.NewLessonContentRepository(db)
	lessonContentService    = service.NewLessonContentService(lessonContentRepository, courseService)
	lessonContentController = controller.NewLessonContentController(lessonContentService)

	transactionController = controller.NewTransactionController(
		service.NewTransactionService(repository.NewTransactionRepository(db), courseService),
	)
)

func NewRouter() *gin.Engine {
	helper.EnvInit()
	DBMigrate(db)

	router := gin.Default()
	router.Use(cors.Default())
	router.Use(gin.CustomRecovery(middleware.ErrorHandler))

	v1 := router.Group("/api/v1")

	// User endpoints
	v1.POST("/users", userController.Register)
	v1.POST("/users/login", userController.Login)
	v1.PUT("/users/avatars", middleware.UserJwtAuthMiddleware(jwtAuth, userService), userController.UploadAvatar)
	v1.GET("/users", middleware.UserJwtAuthMiddleware(jwtAuth, userService), userController.GetById)
	v1.POST("/users/logout", middleware.UserJwtAuthMiddleware(jwtAuth, userService), userController.Logout)

	// Author endpoints
	v1.POST("/authors", authorController.Register)
	v1.GET("/authors", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), authorController.GetById)
	v1.POST("/authors/login", authorController.Login)
	v1.PUT("/authors/avatars", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), authorController.UploadAvatar)
	v1.POST("/authors/logout", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), authorController.Logout)

	// Category endpoints
	v1.POST("/categories", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), categoryController.Create)

	// Course endpoints
	v1.POST("/courses", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), courseController.Create)
	v1.PUT("/courses/:courseId/banners", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), courseController.UploadBanner)
	v1.GET("/courses/authors/:authorID", courseController.GetByAuthorId)
	v1.GET("/courses/:slug", courseController.GetBySlug)
	v1.GET("/courses", courseController.GetAll)
	v1.GET("/courses/categories/:categoryName", courseController.GetByCategory)
	v1.GET("/courses/enrolled", middleware.UserJwtAuthMiddleware(jwtAuth, userService), courseController.GetByUserId)
	v1.POST("/courses/:courseId/enrolled", middleware.UserJwtAuthMiddleware(jwtAuth, userService), courseController.UserEnrolled)

	// Lesson title endpoints
	v1.POST("/authors/courses/:courseId/lesson-titles", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), lessonTitleController.Create)
	v1.PATCH("/authors/courses/:courseId/lesson-titles/:ltId", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), lessonTitleController.Update)
	v1.GET("/courses/enrolled/:courseId/lesson-titles", lessonTitleController.GetByCourseId)

	// Lesson content endpoints
	v1.POST("authors/courses/:courseId/lesson-titles/:ltId/lesson-contents", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), lessonContentController.Create)
	v1.PATCH("authors/courses/:courseId/lesson-contents/:lcId", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), lessonContentController.Update)
	v1.GET("/c/:courseId/lesson-titles/:ltId/lesson-contents", lessonContentController.GetByLessonTitleId)
	v1.GET("/c/:courseId/lesson-titles/lesson-contents/:lcId",
		middleware.UserJwtAuthMiddleware(jwtAuth, userService),
		middleware.MidtransPaymentMiddleware(courseService),
		lessonContentController.GetById,
	)

	// transaction endpoints
	v1.POST("/courses/:courseId/transactions", middleware.UserJwtAuthMiddleware(jwtAuth, userService), transactionController.EarnPaidCourse)
	v1.POST("/transactions/notifications", transactionController.MidtransNotification)

	return router
}
