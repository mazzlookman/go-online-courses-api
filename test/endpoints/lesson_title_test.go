package endpoints

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-pzn-restful-api/test/helper"
	"io"
	"log"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func deleteAllLessonTitles() {
	helper.DeleteLessonTitleTest()
	helper.DeleteCategoryCoursesTest()
	helper.DeleteCourseTest()
	helper.DeleteAuthorTest()
	helper.DeleteCategoryTest()
	log.Println("All data in database has been deleted")
}

func TestCreateLessonTitleSuccess(t *testing.T) {
	courseTest := helper.GetCourseTest()
	authorToken := helper.GetAuthorToken()

	defer deleteAllLessonTitles()

	request := strings.NewReader(`{"title": "title1","in_order": 1}`)

	req := httptest.NewRequest("POST", "/api/v1/authors/courses/"+strconv.Itoa(courseTest.ID)+"/lesson-titles", request)
	req.Header.Add("Authorization", "Bearer "+authorToken)
	w := httptest.NewRecorder()

	helper.Router.ServeHTTP(w, req)

	response := w.Result()
	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
	assert.NotZero(t, mapResponse["data"].(map[string]interface{})["id"])
}

func TestCreateLessonTitleErrorValidation(t *testing.T) {
	courseTest := helper.GetCourseTest()
	authorToken := helper.GetAuthorToken()

	defer deleteAllLessonTitles()

	request := strings.NewReader(`{"title": "","in_order": 1}`)

	req := httptest.NewRequest("POST", "/api/v1/authors/courses/"+strconv.Itoa(courseTest.ID)+"/lesson-titles", request)
	req.Header.Add("Authorization", "Bearer "+authorToken)
	w := httptest.NewRecorder()

	helper.Router.ServeHTTP(w, req)

	response := w.Result()
	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 400, response.StatusCode)
}

func TestCreateLessonTitleErrorCourseNotFound(t *testing.T) {
	courseTest := helper.GetCourseTest()
	authorToken := helper.GetAuthorToken()

	defer deleteAllLessonTitles()

	request := strings.NewReader(`{"title": "title1","in_order": 1}`)

	req := httptest.NewRequest("POST", "/api/v1/authors/courses/"+strconv.Itoa(courseTest.ID+1)+"/lesson-titles", request)
	req.Header.Add("Authorization", "Bearer "+authorToken)
	w := httptest.NewRecorder()

	helper.Router.ServeHTTP(w, req)

	response := w.Result()
	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 404, response.StatusCode)
}

func TestCreateLessonTitleErrorUnauthorized(t *testing.T) {
	request := strings.NewReader(`{"title": "title1","in_order": 1}`)

	req := httptest.NewRequest("POST", "/api/v1/authors/courses/"+strconv.Itoa(1)+"/lesson-titles", request)
	//req.Header.Add("Authorization", "Bearer "+authorToken)
	w := httptest.NewRecorder()

	helper.Router.ServeHTTP(w, req)

	response := w.Result()
	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 401, response.StatusCode)
}

func TestUpdateLessonTitleSuccess(t *testing.T) {
	courseTest := helper.GetCourseTest()
	authorToken := helper.GetAuthorToken()
	lessonTitleTest := helper.CreateLessonTitleTest(courseTest.ID, courseTest.AuthorID)

	defer deleteAllLessonTitles()

	request := strings.NewReader(`{"title": "title1","in_order": 2}`)

	req := httptest.NewRequest("PATCH", "/api/v1/authors/courses/"+strconv.Itoa(courseTest.ID)+"/lesson-titles/"+strconv.Itoa(lessonTitleTest.ID), request)
	req.Header.Add("Authorization", "Bearer "+authorToken)
	w := httptest.NewRecorder()

	helper.Router.ServeHTTP(w, req)

	response := w.Result()
	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, lessonTitleTest.ID, int(mapResponse["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, 2, int(mapResponse["data"].(map[string]interface{})["in_order"].(float64)))
}

func TestUpdateLessonTitleErrorValidation(t *testing.T) {
	courseTest := helper.GetCourseTest()
	authorToken := helper.GetAuthorToken()
	lessonTitleTest := helper.CreateLessonTitleTest(courseTest.ID, courseTest.AuthorID)

	defer deleteAllLessonTitles()

	request := strings.NewReader(`{"title": "","in_order": 2}`)

	req := httptest.NewRequest("PATCH", "/api/v1/authors/courses/"+strconv.Itoa(courseTest.ID)+"/lesson-titles/"+strconv.Itoa(lessonTitleTest.ID), request)
	req.Header.Add("Authorization", "Bearer "+authorToken)
	w := httptest.NewRecorder()

	helper.Router.ServeHTTP(w, req)

	response := w.Result()
	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 400, response.StatusCode)
}

func TestUpdateLessonTitleErrorNotFound(t *testing.T) {
	courseTest := helper.GetCourseTest()
	authorToken := helper.GetAuthorToken()
	lessonTitleTest := helper.CreateLessonTitleTest(courseTest.ID, courseTest.AuthorID)

	defer deleteAllLessonTitles()

	request := strings.NewReader(`{"title": "title1","in_order": 2}`)

	req := httptest.NewRequest("PATCH", "/api/v1/authors/courses/"+strconv.Itoa(courseTest.ID)+"/lesson-titles/"+strconv.Itoa(lessonTitleTest.ID+1), request)
	req.Header.Add("Authorization", "Bearer "+authorToken)
	w := httptest.NewRecorder()

	helper.Router.ServeHTTP(w, req)

	response := w.Result()
	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 404, response.StatusCode)
}

func TestUpdateLessonTitleErrorUnauthorized(t *testing.T) {
	request := strings.NewReader(`{"title": "title1","in_order": 2}`)

	req := httptest.NewRequest("PATCH", "/api/v1/authors/courses/"+strconv.Itoa(1)+"/lesson-titles/"+strconv.Itoa(1), request)
	//req.Header.Add("Authorization", "Bearer "+authorToken)
	w := httptest.NewRecorder()

	helper.Router.ServeHTTP(w, req)

	response := w.Result()
	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 401, response.StatusCode)
}

func TestGetByCourseIDSuccess(t *testing.T) {
	course := helper.GetCourseTest()
	helper.CreateLessonTitleTest(course.ID, course.AuthorID)

	defer deleteAllLessonTitles()

	req := httptest.NewRequest("GET", "/api/v1/enrolled/courses/"+strconv.Itoa(course.ID)+"/lesson-titles", nil)
	w := httptest.NewRecorder()

	helper.Router.ServeHTTP(w, req)

	response := w.Result()
	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 1, len(mapResponse["data"].([]any)))
}

func TestGetByCourseIDErrorCourseNotFound(t *testing.T) {
	course := helper.GetCourseTest()
	helper.CreateLessonTitleTest(course.ID, course.AuthorID)

	defer deleteAllLessonTitles()

	req := httptest.NewRequest("GET", "/api/v1/enrolled/courses/"+strconv.Itoa(course.ID+1)+"/lesson-titles", nil)
	w := httptest.NewRecorder()

	helper.Router.ServeHTTP(w, req)

	response := w.Result()
	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 404, response.StatusCode)
}
