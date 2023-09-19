package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-pzn-restful-api/test/helper"
	"io"
	"log"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

func deleteAllLessonContents() {
	helper.DeleteLessonContentTest()
	helper.DeleteLessonTitleTest()
	helper.DeleteCategoryCoursesTest()
	helper.DeleteCourseTest()
	helper.DeleteAuthorTest()
	helper.DeleteCategoryTest()
	log.Println("All data in database has been deleted")
}

func TestCreateLessonContentSuccess(t *testing.T) {
	defer deleteAllLessonContents()
	courseTest := helper.GetCourseTest()
	authorToken := helper.GetAuthorToken()
	lessonTitleTest := helper.CreateLessonTitleTest(courseTest.Id, courseTest.AuthorId)

	reqBody := new(bytes.Buffer)
	multipartWriter := multipart.NewWriter(reqBody)

	multipartWriter.WriteField("in_order", "1")
	formFile, _ := multipartWriter.CreateFormFile("content", "content.mov")

	pathFile := `D:\CodingTraining\Coding_Tutor\GO\Dasar\Belajar_Go-Lang_-_1_Pengenalan_Go-Lang.mov`
	file, _ := os.Open(pathFile)
	io.Copy(formFile, file)

	multipartWriter.Close()

	req := httptest.NewRequest(
		"POST",
		"/api/v1/authors/courses/"+strconv.Itoa(courseTest.Id)+"/lesson-titles/"+strconv.Itoa(lessonTitleTest.Id)+"/lesson-contents",
		reqBody,
	)
	req.Header.Add("Content-Type", multipartWriter.FormDataContentType())
	req.Header.Add("Authorization", "Bearer "+authorToken)
	w := httptest.NewRecorder()

	helper.Router.ServeHTTP(w, req)

	response := w.Result()
	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
}

func TestCreateLessonContentErrorNotACourseAuthor(t *testing.T) {
	defer deleteAllLessonContents()
	courseTest := helper.GetCourseTest()
	notACourseAuthorToken, _ := helper.JwtAuth.GenerateJwtToken("author", 1)
	lessonTitleTest := helper.CreateLessonTitleTest(courseTest.Id, courseTest.AuthorId)

	reqBody := new(bytes.Buffer)
	multipartWriter := multipart.NewWriter(reqBody)

	multipartWriter.WriteField("in_order", "1")
	formFile, _ := multipartWriter.CreateFormFile("content", "content.mov")

	pathFile := `D:\CodingTraining\Coding_Tutor\GO\Dasar\Belajar_Go-Lang_-_1_Pengenalan_Go-Lang.mov`
	file, _ := os.Open(pathFile)
	io.Copy(formFile, file)

	multipartWriter.Close()

	req := httptest.NewRequest(
		"POST",
		"/api/v1/authors/courses/"+strconv.Itoa(courseTest.Id)+"/lesson-titles/"+strconv.Itoa(lessonTitleTest.Id)+"/lesson-contents",
		reqBody,
	)
	req.Header.Add("Content-Type", multipartWriter.FormDataContentType())
	req.Header.Add("Authorization", "Bearer "+notACourseAuthorToken)
	w := httptest.NewRecorder()

	helper.Router.ServeHTTP(w, req)

	response := w.Result()
	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 401, response.StatusCode)
	assert.Equal(t, "You're not an author of this courses", mapResponse["data"])
}

func TestCreateLessonContentErrorCourseNotFound(t *testing.T) {
	defer deleteAllLessonContents()
	courseTest := helper.GetCourseTest()
	authorToken := helper.GetAuthorToken()
	lessonTitleTest := helper.CreateLessonTitleTest(courseTest.Id, courseTest.AuthorId)

	reqBody := new(bytes.Buffer)
	multipartWriter := multipart.NewWriter(reqBody)

	multipartWriter.WriteField("in_order", "1")
	formFile, _ := multipartWriter.CreateFormFile("content", "content.mov")

	pathFile := `D:\CodingTraining\Coding_Tutor\GO\Dasar\Belajar_Go-Lang_-_1_Pengenalan_Go-Lang.mov`
	file, _ := os.Open(pathFile)
	io.Copy(formFile, file)

	multipartWriter.Close()

	req := httptest.NewRequest(
		"POST",
		"/api/v1/authors/courses/"+strconv.Itoa(courseTest.Id+1)+"/lesson-titles/"+strconv.Itoa(lessonTitleTest.Id)+"/lesson-contents",
		reqBody,
	)
	req.Header.Add("Content-Type", multipartWriter.FormDataContentType())
	req.Header.Add("Authorization", "Bearer "+authorToken)
	w := httptest.NewRecorder()

	helper.Router.ServeHTTP(w, req)

	response := w.Result()
	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 404, response.StatusCode)
	assert.Equal(t, "Course not found", mapResponse["data"])
}

func TestUpdateLessonContentSuccessContentOnly(t *testing.T) {
	defer deleteAllLessonContents()
	courseTest := helper.GetCourseTest()
	authorToken := helper.GetAuthorToken()
	lessonTitleTest := helper.CreateLessonTitleTest(courseTest.Id, courseTest.AuthorId)
	lessonContentTest := helper.CreateLessonContentTest(courseTest.AuthorId, courseTest.Id, lessonTitleTest.Id)

	courseIDStr := strconv.Itoa(courseTest.Id)
	ltIDStr := strconv.Itoa(lessonTitleTest.Id)
	lcIDStr := strconv.Itoa(lessonContentTest.Id)

	reqBody := new(bytes.Buffer)
	multipartWriter := multipart.NewWriter(reqBody)

	//multipartWriter.WriteField("in_order", "1")
	formFile, _ := multipartWriter.CreateFormFile("content", "content_updated.mov")

	pathFile := `D:\CodingTraining\Coding_Tutor\GO\Dasar\Belajar_Go-Lang_-_1_Pengenalan_Go-Lang.mov`
	file, _ := os.Open(pathFile)
	io.Copy(formFile, file)

	multipartWriter.Close()

	req := httptest.NewRequest(
		"PATCH",
		"/api/v1/authors/courses/"+courseIDStr+"/lesson-titles/"+ltIDStr+"/lesson-contents/"+lcIDStr,
		reqBody,
	)
	req.Header.Add("Content-Type", multipartWriter.FormDataContentType())
	req.Header.Add("Authorization", "Bearer "+authorToken)
	w := httptest.NewRecorder()

	helper.Router.ServeHTTP(w, req)

	response := w.Result()
	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 1, int(mapResponse["data"].(map[string]any)["in_order"].(float64)))
	assert.Equal(t, "assets/contents/content_updated.mov", mapResponse["data"].(map[string]any)["content"])
}

func TestUpdateLessonContentSuccessInOrderOnly(t *testing.T) {
	defer deleteAllLessonContents()
	courseTest := helper.GetCourseTest()
	authorToken := helper.GetAuthorToken()
	lessonTitleTest := helper.CreateLessonTitleTest(courseTest.Id, courseTest.AuthorId)
	lessonContentTest := helper.CreateLessonContentTest(courseTest.AuthorId, courseTest.Id, lessonTitleTest.Id)

	courseIDStr := strconv.Itoa(courseTest.Id)
	ltIDStr := strconv.Itoa(lessonTitleTest.Id)
	lcIDStr := strconv.Itoa(lessonContentTest.Id)

	reqBody := new(bytes.Buffer)
	multipartWriter := multipart.NewWriter(reqBody)
	multipartWriter.WriteField("in_order", "2")
	multipartWriter.Close()

	req := httptest.NewRequest(
		"PATCH",
		"/api/v1/authors/courses/"+courseIDStr+"/lesson-titles/"+ltIDStr+"/lesson-contents/"+lcIDStr,
		reqBody,
	)
	req.Header.Add("Content-Type", multipartWriter.FormDataContentType())
	req.Header.Add("Authorization", "Bearer "+authorToken)
	w := httptest.NewRecorder()

	helper.Router.ServeHTTP(w, req)

	response := w.Result()
	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 2, int(mapResponse["data"].(map[string]any)["in_order"].(float64)))
	assert.Equal(t, "assets/contents/content.mov", mapResponse["data"].(map[string]any)["content"])
}

func TestUpdateLessonContentErrorNotACourseAuthor(t *testing.T) {
	defer deleteAllLessonContents()
	courseTest := helper.GetCourseTest()
	authorToken, _ := helper.JwtAuth.GenerateJwtToken("author", 1)
	lessonTitleTest := helper.CreateLessonTitleTest(courseTest.Id, courseTest.AuthorId)
	lessonContentTest := helper.CreateLessonContentTest(courseTest.AuthorId, courseTest.Id, lessonTitleTest.Id)

	courseIDStr := strconv.Itoa(courseTest.Id)
	ltIDStr := strconv.Itoa(lessonTitleTest.Id)
	lcIDStr := strconv.Itoa(lessonContentTest.Id)

	reqBody := new(bytes.Buffer)
	multipartWriter := multipart.NewWriter(reqBody)

	multipartWriter.WriteField("in_order", "1")
	//formFile, _ := multipartWriter.CreateFormFile("content", "content_updated.mov")
	//
	//pathFile := `D:\CodingTraining\Coding_Tutor\GO\Dasar\Belajar_Go-Lang_-_1_Pengenalan_Go-Lang.mov`
	//file, _ := os.Open(pathFile)
	//io.Copy(formFile, file)

	multipartWriter.Close()

	req := httptest.NewRequest(
		"PATCH",
		"/api/v1/authors/courses/"+courseIDStr+"/lesson-titles/"+ltIDStr+"/lesson-contents/"+lcIDStr,
		reqBody,
	)
	req.Header.Add("Content-Type", multipartWriter.FormDataContentType())
	req.Header.Add("Authorization", "Bearer "+authorToken)
	w := httptest.NewRecorder()

	helper.Router.ServeHTTP(w, req)

	response := w.Result()
	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 401, response.StatusCode)
	assert.Equal(t, "You're not an author of this courses", mapResponse["data"])
}

func TestUpdateLessonContentErrorNotFound(t *testing.T) {
	defer deleteAllLessonContents()
	courseTest := helper.GetCourseTest()
	authorToken := helper.GetAuthorToken()
	lessonTitleTest := helper.CreateLessonTitleTest(courseTest.Id, courseTest.AuthorId)
	lessonContentTest := helper.CreateLessonContentTest(courseTest.AuthorId, courseTest.Id, lessonTitleTest.Id)

	courseIDStr := strconv.Itoa(courseTest.Id)
	ltIDStr := strconv.Itoa(lessonTitleTest.Id)
	lcIDStr := strconv.Itoa(lessonContentTest.Id + 1)

	reqBody := new(bytes.Buffer)
	multipartWriter := multipart.NewWriter(reqBody)
	multipartWriter.WriteField("in_order", "1")
	multipartWriter.Close()

	req := httptest.NewRequest(
		"PATCH",
		"/api/v1/authors/courses/"+courseIDStr+"/lesson-titles/"+ltIDStr+"/lesson-contents/"+lcIDStr,
		reqBody,
	)
	req.Header.Add("Content-Type", multipartWriter.FormDataContentType())
	req.Header.Add("Authorization", "Bearer "+authorToken)
	w := httptest.NewRecorder()

	helper.Router.ServeHTTP(w, req)

	response := w.Result()
	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 404, response.StatusCode)
	assert.Equal(t, "Lesson content not found", mapResponse["data"])
}

func TestGetByLessonTitleIDSuccess(t *testing.T) {
	defer deleteAllLessonContents()
	courseTest := helper.GetCourseTest()
	lessonTitleTest := helper.CreateLessonTitleTest(courseTest.Id, courseTest.AuthorId)
	helper.CreateLessonContentTest(courseTest.AuthorId, courseTest.Id, lessonTitleTest.Id)
	ltIDStr := strconv.Itoa(lessonTitleTest.Id)

	req := httptest.NewRequest(
		"GET",
		"/api/v1/enrolled/courses/lesson-titles/"+ltIDStr+"/lesson-contents",
		nil,
	)
	w := httptest.NewRecorder()

	helper.Router.ServeHTTP(w, req)

	response := w.Result()
	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 1, len(mapResponse["data"].([]any)))
	//assert.Equal(t, "Lesson content not found", mapResponse["data"])
}

func TestGetByLessonTitleErrorNotFound(t *testing.T) {
	defer deleteAllLessonContents()
	courseTest := helper.GetCourseTest()
	lessonTitleTest := helper.CreateLessonTitleTest(courseTest.Id, courseTest.AuthorId)
	helper.CreateLessonContentTest(courseTest.AuthorId, courseTest.Id, lessonTitleTest.Id)
	ltIDStr := strconv.Itoa(lessonTitleTest.Id + 1)

	req := httptest.NewRequest(
		"GET",
		"/api/v1/enrolled/courses/lesson-titles/"+ltIDStr+"/lesson-contents",
		nil,
	)
	w := httptest.NewRecorder()

	helper.Router.ServeHTTP(w, req)

	response := w.Result()
	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 404, response.StatusCode)
	assert.Equal(t, "Lesson contents not found", mapResponse["data"])
}
