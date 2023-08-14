package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestCreateCourseSuccess(t *testing.T) {
	CreateAuthorTest()
	CreateCategoryTest()
	token := GetAuthorToken()

	defer DeleteCategoryTest()
	defer DeleteAuthorTest()
	defer DeleteCourseTest()
	defer DeleteCategoryCoursesTest()

	request := strings.NewReader(`{
							  "title": "Golang",
							  "slug": "golang",
							  "description": "Desc",
							  "perks": "p1,p2,p3",
							  "price": 99000,
							  "category": "Backend"
							}`)

	req := httptest.NewRequest("POST", "/api/v1/courses", request)
	req.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	Router.ServeHTTP(w, req)
	response := w.Result()

	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
	assert.NotZero(t, "author_id", mapResponse["data"].(map[string]interface{})["author_id"])
}

func TestCreateCourseErrorValidation(t *testing.T) {
	CreateAuthorTest()
	CreateCategoryTest()
	token := GetAuthorToken()

	defer DeleteCategoryTest()
	defer DeleteAuthorTest()
	defer DeleteCourseTest()
	defer DeleteCategoryCoursesTest()

	request := strings.NewReader(`{
							  "title": "",
							  "slug": "",
							  "description": "",
							  "perks": "p1,p2,p3",
							  "price": 99000,
							  "category": "Backend"
							}`)

	req := httptest.NewRequest("POST", "/api/v1/courses", request)
	req.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	Router.ServeHTTP(w, req)
	response := w.Result()

	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 400, response.StatusCode)
}

func TestCreateCourseErrorUnauthorized(t *testing.T) {
	CreateAuthorTest()
	token := GetAuthorToken()

	defer DeleteCategoryTest()
	defer DeleteAuthorTest()
	defer DeleteCourseTest()
	defer DeleteCategoryCoursesTest()

	request := strings.NewReader(`{
							  "title": "",
							  "slug": "",
							  "description": "",
							  "perks": "p1,p2,p3",
							  "price": 99000,
							  "category": "Backend"
							}`)

	req := httptest.NewRequest("POST", "/api/v1/courses", request)
	req.Header.Add("Authorization", "Bear "+token)
	w := httptest.NewRecorder()

	Router.ServeHTTP(w, req)
	response := w.Result()

	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 401, response.StatusCode)
}

func TestUploadBannerSuccess(t *testing.T) {
	CreateCategoryTest()
	authorTest := CreateAuthorTest()
	courseTest := CreateCourseTest(authorTest.ID)
	token := GetAuthorToken()

	defer DeleteCategoryTest()
	defer DeleteAuthorTest()
	defer DeleteCourseTest()
	defer DeleteCategoryCoursesTest()

	request := new(bytes.Buffer)
	writer := multipart.NewWriter(request)
	formFile, _ := writer.CreateFormFile("banner", "golang.jpg")

	pathFile := `C:\Users\moham\Downloads\all.jpg`
	file, _ := os.Open(pathFile)
	io.Copy(formFile, file)
	writer.Close()

	req := httptest.NewRequest("PUT", "/api/v1/courses/"+strconv.Itoa(courseTest.ID)+"/banners", request)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	Router.ServeHTTP(w, req)

	response := w.Result()
	bytes, _ := io.ReadAll(response.Body)

	var mapResponse map[string]interface{}
	json.Unmarshal(bytes, &mapResponse)

	log.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
	assert.NotEmpty(t, mapResponse["data"].(map[string]any)["is_uploaded"])
}

func TestUploadBannerErrorUnauthorized(t *testing.T) {
	request := new(bytes.Buffer)
	writer := multipart.NewWriter(request)
	formFile, _ := writer.CreateFormFile("banner", "golang.jpg")

	pathFile := `C:\Users\moham\Downloads\all.jpg`
	file, _ := os.Open(pathFile)
	io.Copy(formFile, file)
	writer.Close()

	req := httptest.NewRequest("PUT", "/api/v1/courses/23/banners", request)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	//req.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	Router.ServeHTTP(w, req)

	response := w.Result()
	bytes, _ := io.ReadAll(response.Body)

	var mapResponse map[string]interface{}
	json.Unmarshal(bytes, &mapResponse)

	log.Println(mapResponse)
	assert.Equal(t, 401, response.StatusCode)
}

func TestUploadBannerErrorCourseNotFound(t *testing.T) {
	CreateAuthorTest()
	token := GetAuthorToken()

	defer DeleteAuthorTest()

	request := new(bytes.Buffer)
	writer := multipart.NewWriter(request)
	formFile, _ := writer.CreateFormFile("banner", "golang.jpg")

	pathFile := `C:\Users\moham\Downloads\all.jpg`
	file, _ := os.Open(pathFile)
	io.Copy(formFile, file)
	writer.Close()

	req := httptest.NewRequest("PUT", "/api/v1/courses/300/banners", request)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	Router.ServeHTTP(w, req)

	response := w.Result()
	bytes, _ := io.ReadAll(response.Body)

	var mapResponse map[string]interface{}
	json.Unmarshal(bytes, &mapResponse)
	log.Println(mapResponse)

	assert.Equal(t, 404, response.StatusCode)
}

func TestGetByAuthorIdSuccess(t *testing.T) {
	CreateCategoryTest()
	author := CreateAuthorTest()
	CreateCourseTest(author.ID)

	defer DeleteCategoryTest()
	defer DeleteAuthorTest()
	defer DeleteCourseTest()
	defer DeleteCategoryCoursesTest()

	req := httptest.NewRequest("GET", "/api/v1/courses/authors/"+strconv.Itoa(author.ID), nil)
	w := httptest.NewRecorder()

	Router.ServeHTTP(w, req)
	response := w.Result()

	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 1, len(mapResponse["data"].([]any)))
}

func TestGetByAuthorIdErrorCoursesNotFound(t *testing.T) {
	CreateCategoryTest()
	author := CreateAuthorTest()
	CreateCourseTest(author.ID)

	defer DeleteCategoryTest()
	defer DeleteAuthorTest()
	defer DeleteCourseTest()
	defer DeleteCategoryCoursesTest()

	req := httptest.NewRequest("GET", "/api/v1/courses/authors/"+strconv.Itoa(author.ID+1), nil)
	w := httptest.NewRecorder()

	Router.ServeHTTP(w, req)
	response := w.Result()

	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 404, response.StatusCode)
}

func TestGetBySlugSuccess(t *testing.T) {
	CreateCategoryTest()
	author := CreateAuthorTest()
	courseTest := CreateCourseTest(author.ID)

	defer DeleteCategoryTest()
	defer DeleteAuthorTest()
	defer DeleteCourseTest()
	defer DeleteCategoryCoursesTest()

	req := httptest.NewRequest("GET", "/api/v1/courses/"+courseTest.Slug, nil)
	w := httptest.NewRecorder()

	Router.ServeHTTP(w, req)
	response := w.Result()

	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "golang", mapResponse["data"].(map[string]any)["slug"])
}

func TestGetBySlugErrorCourseNotFound(t *testing.T) {
	CreateCategoryTest()
	author := CreateAuthorTest()
	CreateCourseTest(author.ID)

	defer DeleteCategoryTest()
	defer DeleteAuthorTest()
	defer DeleteCourseTest()
	defer DeleteCategoryCoursesTest()

	req := httptest.NewRequest("GET", "/api/v1/courses/notfound", nil)
	w := httptest.NewRecorder()

	Router.ServeHTTP(w, req)
	response := w.Result()

	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 404, response.StatusCode)
}

func TestGetAllSuccess(t *testing.T) {
	CreateCategoryTest()
	author := CreateAuthorTest()
	CreateCourseTest(author.ID)

	defer DeleteCategoryTest()
	defer DeleteAuthorTest()
	defer DeleteCourseTest()
	defer DeleteCategoryCoursesTest()

	req := httptest.NewRequest("GET", "/api/v1/courses", nil)
	w := httptest.NewRecorder()

	Router.ServeHTTP(w, req)
	response := w.Result()

	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 1, len(mapResponse["data"].([]any)))
}

func TestGetAllErrorCoursesNotFound(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/v1/courses", nil)
	w := httptest.NewRecorder()

	Router.ServeHTTP(w, req)
	response := w.Result()

	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 404, response.StatusCode)
}

func TestGetByCategorySuccess(t *testing.T) {
	categoryTest := CreateCategoryTest()
	author := CreateAuthorTest()
	CreateCourseTest(author.ID)

	defer DeleteCategoryTest()
	defer DeleteAuthorTest()
	defer DeleteCourseTest()
	defer DeleteCategoryCoursesTest()

	req := httptest.NewRequest("GET", "/api/v1/courses/categories/"+categoryTest.Name, nil)
	w := httptest.NewRecorder()

	Router.ServeHTTP(w, req)
	response := w.Result()

	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 1, len(mapResponse["data"].([]any)))
}

func TestGetByCategoryErrorCoursesNotFound(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/v1/courses/categories/notfound", nil)
	w := httptest.NewRecorder()

	Router.ServeHTTP(w, req)
	response := w.Result()

	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 404, response.StatusCode)
}

func TestUserEnrolledSuccess(t *testing.T) {
	CreateUserTest()
	tokenUser := GetTokenAfterLogin()

	CreateCategoryTest()
	author := CreateAuthorTest()
	courseTest := CreateCourseTest(author.ID)

	defer DeleteCategoryTest()
	defer DeleteAuthorTest()
	defer DeleteCourseTest()
	defer DeleteCategoryCoursesTest()
	defer DeleteUserTest()
	defer DeleteUserCoursesTest()

	req := httptest.NewRequest("POST", "/api/v1/courses/"+strconv.Itoa(courseTest.ID)+"/enrolled", nil)
	req.Header.Add("Authorization", "Bearer "+tokenUser)
	w := httptest.NewRecorder()

	Router.ServeHTTP(w, req)
	response := w.Result()

	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
}

func TestUserEnrolledErrorUnauthorized(t *testing.T) {
	req := httptest.NewRequest("POST", "/api/v1/courses/200/enrolled", nil)
	//req.Header.Add("Authorization", "Bearer "+tokenUser)
	w := httptest.NewRecorder()

	Router.ServeHTTP(w, req)
	response := w.Result()

	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 401, response.StatusCode)
}

func TestUserEnrolledErrorCourseNotFound(t *testing.T) {
	CreateUserTest()
	tokenUser := GetTokenAfterLogin()

	CreateCategoryTest()
	author := CreateAuthorTest()
	courseTest := CreateCourseTest(author.ID)

	defer DeleteCategoryTest()
	defer DeleteAuthorTest()
	defer DeleteCourseTest()
	defer DeleteCategoryCoursesTest()
	defer DeleteUserTest()
	defer DeleteUserCoursesTest()

	req := httptest.NewRequest("POST", "/api/v1/courses/"+strconv.Itoa(courseTest.ID+1)+"/enrolled", nil)
	req.Header.Add("Authorization", "Bearer "+tokenUser)
	w := httptest.NewRecorder()

	Router.ServeHTTP(w, req)
	response := w.Result()

	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 404, response.StatusCode)
}

func TestGetByUserIDSuccess(t *testing.T) {
	userTest := CreateUserTest()
	tokenUser := GetTokenAfterLogin()

	CreateCategoryTest()
	authorTest := CreateAuthorTest()
	courseTest := CreateCourseTest(authorTest.ID)

	CreateUserCoursesTest(userTest.ID, courseTest.ID)

	defer DeleteCategoryTest()
	defer DeleteAuthorTest()
	defer DeleteCourseTest()
	defer DeleteCategoryCoursesTest()
	defer DeleteUserTest()
	defer DeleteUserCoursesTest()

	req := httptest.NewRequest("GET", "/api/v1/courses/enrolled", nil)
	req.Header.Add("Authorization", "Bearer "+tokenUser)
	w := httptest.NewRecorder()

	Router.ServeHTTP(w, req)
	response := w.Result()

	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 1, len(mapResponse["data"].([]any)))
}

func TestGetByUserIDErrorCoursesNotFound(t *testing.T) {
	CreateUserTest()
	tokenUser := GetTokenAfterLogin()

	CreateCategoryTest()
	authorTest := CreateAuthorTest()
	CreateCourseTest(authorTest.ID)

	//CreateUserCoursesTest(userTest.ID, courseTest.ID)

	defer DeleteCategoryTest()
	defer DeleteAuthorTest()
	defer DeleteCourseTest()
	defer DeleteCategoryCoursesTest()
	defer DeleteUserTest()

	req := httptest.NewRequest("GET", "/api/v1/courses/enrolled", nil)
	req.Header.Add("Authorization", "Bearer "+tokenUser)
	w := httptest.NewRecorder()

	Router.ServeHTTP(w, req)
	response := w.Result()

	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 404, response.StatusCode)
}

func TestGetByUserIDErrorUnauthorized(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/v1/courses/enrolled", nil)
	//req.Header.Add("Authorization", "Bearer "+tokenUser)
	w := httptest.NewRecorder()

	Router.ServeHTTP(w, req)
	response := w.Result()

	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 401, response.StatusCode)
}
