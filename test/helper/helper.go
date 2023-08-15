package helper

import (
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
	"io"
	"log"
	"os"
)

// User
func CreateUserTest() web.UserResponse {
	input := web.UserRegisterInput{
		Name:     "user",
		Email:    "user@user.com",
		Password: "123",
	}
	log.Println("User registered")
	return UserService.Register(input)
}

func DeleteUserTest() {
	db, _ := Db.DB()
	result, _ := db.Exec("delete from users")
	rowsAffected, _ := result.RowsAffected()
	log.Println("User has been deleted, rows affected: ", rowsAffected)
}

func GetTokenAfterLogin() string {
	login := UserService.Login(web.UserLoginInput{
		Email:    "user@user.com",
		Password: "123",
	})

	return login.Token
}

func DeleteUserCoursesTest() {
	db, _ := Db.DB()
	result, _ := db.Exec("delete from user_courses")
	rowsAffected, _ := result.RowsAffected()
	log.Println("User courses has been deleted, rows affected: ", rowsAffected)
}

// Author
func CreateAuthorTest() web.AuthorResponse {
	input := web.AuthorRegisterInput{
		Name:     "author",
		Email:    "author@author.com",
		Password: "123",
		Profile:  "Profile",
		Avatar:   "assets/images/avatars/author.jpg",
	}

	log.Println("Author has been created")
	return AuthorService.Register(input)
}

func DeleteAuthorTest() {
	//err := AuthorRepository.Delete("author@author.com")
	db, _ := Db.DB()
	result, _ := db.Exec("delete from authors")
	rowsAffected, _ := result.RowsAffected()
	log.Println("Authors has been deleted, rows affected: ", rowsAffected)
}

func GetAuthorToken() string {
	login := AuthorService.Login(web.AuthorLoginInput{
		Email:    "author@author.com",
		Password: "123",
	})

	return login.Token
}

// Category
func CreateCategoryTest() web.CategoryResponse {
	return CategoryService.Create(web.CategoryCreateInput{Name: "backend"})
}

func DeleteCategoryTest() {
	//tx := Db.Delete(&domain.Category{}, "name=?", "backend")
	db, _ := Db.DB()
	result, _ := db.Exec("delete from categories")
	rowsAffected, _ := result.RowsAffected()
	log.Println("Category has been deleted, rows affected: ", rowsAffected)
}

// Course
func CreateCourseTest(authorID int) web.CourseResponse {
	return CourseService.Create(web.CourseCreateInput{
		AuthorID:    authorID,
		Title:       "Golang",
		Slug:        "golang",
		Description: "Desc",
		Perks:       "p1,p2,p3",
		Price:       99000,
		Category:    "backend",
	})
}

func CreateUserCoursesTest(userID int, courseID int) domain.UserCourse {
	userCourse, err := CourseRepository.UsersEnrolled(domain.UserCourse{
		CourseID: courseID,
		UserID:   userID,
	})

	if err != nil {
		log.Fatalln(err)
	}

	return userCourse
}

func DeleteCourseTest() {
	db, _ := Db.DB()
	result, _ := db.Exec("delete from courses")
	rowsAffected, _ := result.RowsAffected()
	log.Println("Course has been deleted, rows affected: ", rowsAffected)
}

func DeleteCategoryCoursesTest() {
	db, _ := Db.DB()
	result, _ := db.Exec("delete from category_courses")
	rowsAffected, _ := result.RowsAffected()
	log.Println("Category_courses has been deleted, rows affected: ", rowsAffected)
	//tx := Db.Delete(&domain.CategoryCourse{})
}

func GetCourseTest() web.CourseResponse {
	CreateCategoryTest()
	authorTest := CreateAuthorTest()
	return CreateCourseTest(authorTest.ID)
}

// Lesson Title
func DeleteLessonTitleTest() {
	db, _ := Db.DB()
	result, _ := db.Exec("delete from lesson_titles")
	rowsAffected, _ := result.RowsAffected()
	log.Println("Lesson titles has been deleted, rows affected: ", rowsAffected)
}

func CreateLessonTitleTest(courseID int, authorID int) web.LessonTitleResponse {
	return LessonTitleService.Create(web.LessonTitleCreateInput{
		CourseID: courseID,
		Title:    "title1",
		InOrder:  1,
		AuthorID: authorID,
	})
}

// Lesson Content
func CreateLessonContentTest(authorID int, courseID int, ltID int) web.LessonContentResponse {
	pathFile := `D:\CodingTraining\Coding_Tutor\GO\Dasar\Belajar_Go-Lang_-_1_Pengenalan_Go-Lang.mov`
	create, _ := os.Create("assets/contents/content.mov")
	open, _ := os.Open(pathFile)
	io.Copy(create, open)
	create.Close()
	open.Close()

	return LessonContentService.Create(web.LessonContentCreateInput{
		AuthorID:      authorID,
		CourseID:      courseID,
		LessonTitleID: ltID,
		Content:       "assets/contents/content.mov",
		InOrder:       1,
	})
}

func DeleteLessonContentTest() {
	db, _ := Db.DB()
	result, _ := db.Exec("delete from lesson_contents")
	rowsAffected, _ := result.RowsAffected()
	log.Println("Lesson content has been deleted, rows affected: ", rowsAffected)
}
