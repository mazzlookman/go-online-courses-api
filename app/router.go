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
)

func NewRouter() *gin.Engine {
	DBMigrate(db)
	router := gin.Default()
	router.Use(gin.CustomRecovery(middleware.ErrorHandler))

	v1 := router.Group("/api/v1")

	// User endpoints
	v1.POST("/users", userController.Register)
	v1.POST("/users/login", userController.Login)
	v1.PUT("/users/avatars", middleware.UserJwtAuthMiddleware(jwtAuth, userService), userController.UploadAvatar)
	v1.GET("/users", middleware.UserJwtAuthMiddleware(jwtAuth, userService), userController.GetByID)
	v1.POST("/users/logout", middleware.UserJwtAuthMiddleware(jwtAuth, userService), userController.Logout)

	// Author endpoints
	v1.POST("/authors", authorController.Register)
	v1.POST("/authors/login", authorController.Login)
	v1.POST("/authors/logout", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), authorController.Logout)

	// Category endpoints
	v1.POST("/categories", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), categoryController.Create)

	// Course endpoints
	v1.POST("/courses", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), courseController.Create)
	v1.PUT("/courses/:courseID/banners", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), courseController.UploadBanner)
	v1.GET("/courses/authors/:authorID", courseController.GetByAuthorID)
	v1.GET("/courses/:slug", courseController.GetBySlug)
	v1.GET("/courses", courseController.GetAll)
	v1.GET("/courses/categories/:categoryName", courseController.GetByCategory)
	v1.GET("/courses/enrolled", middleware.UserJwtAuthMiddleware(jwtAuth, userService), courseController.GetByUserID)
	v1.POST("/courses/:courseID/enrolled", middleware.UserJwtAuthMiddleware(jwtAuth, userService), courseController.UserEnrolled)

	// Lesson title endpoints
	v1.POST("/authors/courses/:courseID/lesson-titles", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), lessonTitleController.Create)
	v1.PATCH("/authors/courses/:courseID/lesson-titles/:ltID", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), lessonTitleController.Update)
	//add payment middleware (later)
	v1.GET("/courses/enrolled/:courseID/lesson-titles", lessonTitleController.GetByCourseID)

	// Lesson content endpoints
	v1.POST("authors/courses/:courseID/lesson-titles/:ltID/lesson-contents", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), lessonContentController.Create)
	v1.PATCH("authors/courses/:courseID/lesson-titles/:ltID/lesson-contents/:lcID", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), lessonContentController.Update)
	//add payment middleware (later)
	v1.GET("/courses/enrolled/lesson-titles/:ltID/lesson-contents", lessonContentController.GetByLessonTitleID)

	return router
}
