package test

import (
	"go-pzn-restful-api/app"
	"go-pzn-restful-api/auth"
	"go-pzn-restful-api/controller"
	"go-pzn-restful-api/repository"
	"go-pzn-restful-api/service"
)

var (
	Router = app.NewRouter()
	Db     = app.DBConnection()
)

// User
var (
	UserRepository = repository.NewUserRepository(Db)
	JwtAuth        = auth.NewJwtAuth()
	UserService    = service.NewUserService(UserRepository, JwtAuth)
	UserController = controller.NewUserController(UserService)
)

// Author
var (
	AuthorRepository = repository.NewAuthorRepository(Db)
	AuthorService    = service.NewAuthorService(AuthorRepository, JwtAuth)
	AuthorController = controller.NewAuthorController(AuthorService)
)

var (
	CategoryRepository = repository.NewCategoryRepository(Db)
	CategoryService    = service.NewCategoryService(CategoryRepository)
	CategoryController = controller.NewCategoryController(CategoryService)
)

// Course
var (
	CourseRepository = repository.NewCourseRepository(Db)
	CourseService    = service.NewCourseService(CourseRepository)
	CourseController = controller.NewCourseController(CourseService)
)
