package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-pzn-restful-api/test/helper"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestUserRegisterSuccess(t *testing.T) {
	defer helper.DeleteUserTest()
	reqBody := strings.NewReader(`{"name": "user","email": "user@user.com","password": "123"}`)
	req := httptest.NewRequest("POST", "/api/v1/users", reqBody)
	w := httptest.NewRecorder()

	helper.Router.ServeHTTP(w, req)

	response := w.Result()
	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)
	fmt.Println(mapResponse)

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "user", mapResponse["data"].(map[string]interface{})["name"])
}

func TestUserRegisterErrorValidation(t *testing.T) {
	defer helper.DeleteUserTest()
	reqBody := strings.NewReader(`{"name": "test","email": "test@test.com","password": ""}`)
	req := httptest.NewRequest("POST", "/api/v1/users", reqBody)
	w := httptest.NewRecorder()

	helper.Router.ServeHTTP(w, req)

	response := w.Result()
	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)
	fmt.Println(mapResponse)

	assert.Equal(t, 400, response.StatusCode)
}

func TestUserLoginSuccess(t *testing.T) {
	helper.CreateUserTest()
	defer helper.DeleteUserTest()

	reqBody := strings.NewReader(`{"email": "user@user.com","password": "123"}`)

	req := httptest.NewRequest("POST", "/api/v1/users/login", reqBody)
	w := httptest.NewRecorder()
	helper.Router.ServeHTTP(w, req)

	response := w.Result()
	bytes, _ := io.ReadAll(response.Body)

	var mapResponse map[string]interface{}
	json.Unmarshal(bytes, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "user", mapResponse["data"].(map[string]interface{})["name"])
	assert.NotNil(t, mapResponse["data"].(map[string]any)["token"])
}

func TestUserLoginErrorUsernameOrPasswordIsWrong(t *testing.T) {
	helper.CreateUserTest()
	defer helper.DeleteUserTest()

	reqBody := strings.NewReader(`{"email": "user@user.com","password": "12"}`)

	req := httptest.NewRequest("POST", "/api/v1/users/login", reqBody)
	w := httptest.NewRecorder()
	helper.Router.ServeHTTP(w, req)

	response := w.Result()
	bytes, _ := io.ReadAll(response.Body)

	var mapResponse map[string]interface{}
	json.Unmarshal(bytes, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 404, response.StatusCode)
}

func TestUserLoginErrorValidation(t *testing.T) {
	helper.CreateUserTest()
	defer helper.DeleteUserTest()

	reqBody := strings.NewReader(`{"email": "useruser.com","password": ""}`)

	req := httptest.NewRequest("POST", "/api/v1/users/login", reqBody)
	w := httptest.NewRecorder()

	helper.Router.ServeHTTP(w, req)

	response := w.Result()
	bytes, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(bytes, &mapResponse)
	fmt.Println(mapResponse)

	assert.Equal(t, 400, response.StatusCode)
}

func TestGetUserByIdSuccess(t *testing.T) {
	helper.CreateUserTest()
	defer helper.DeleteUserTest()

	token := helper.GetTokenAfterLogin()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/users", nil)
	req.Header.Add("Authorization", "Bearer "+token)
	helper.Router.ServeHTTP(w, req)

	response := w.Result()
	bytes, _ := io.ReadAll(response.Body)

	var mapResponse map[string]interface{}
	json.Unmarshal(bytes, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
	assert.NotNil(t, mapResponse["data"].(map[string]any)["id"])
	assert.Equal(t, "user", mapResponse["data"].(map[string]any)["name"])
}

func TestGetUserByIdErrorUnauthorized(t *testing.T) {
	helper.CreateUserTest()
	defer helper.DeleteUserTest()
	// random token
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMSJ9.vE_WTAksI9VxxurKeIDb-OETLYJzmecOuqs0FZJJ6kE"

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/users", nil)
	req.Header.Add("Authorization", "Bearer "+token)
	helper.Router.ServeHTTP(w, req)

	response := w.Result()
	bytes, _ := io.ReadAll(response.Body)

	var mapResponse map[string]interface{}
	json.Unmarshal(bytes, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 401, response.StatusCode)
}

func TestUploadAvatarSuccess(t *testing.T) {
	helper.CreateUserTest()
	defer helper.DeleteUserTest()
	token := helper.GetTokenAfterLogin()

	body := new(bytes.Buffer)
	multipartWriter := multipart.NewWriter(body)
	createFormFile, _ := multipartWriter.CreateFormFile("avatar", "user.jpg")

	file, err := os.Open(`C:\Users\moham\Downloads\fc_page-0001.jpg`)
	assert.Nil(t, err)
	io.Copy(createFormFile, file)

	multipartWriter.Close()

	req := httptest.NewRequest("PUT", "/api/v1/users/avatars", body)
	req.Header.Add("Content-Type", multipartWriter.FormDataContentType())
	req.Header.Add("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	helper.Router.ServeHTTP(rec, req)

	response := rec.Result()
	bytes, _ := io.ReadAll(response.Body)

	var mapResponse map[string]interface{}
	json.Unmarshal(bytes, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
}

func TestUploadAvatarErrorUnAuthorized(t *testing.T) {
	helper.CreateUserTest()
	defer helper.DeleteUserTest()
	token := helper.GetTokenAfterLogin()

	body := new(bytes.Buffer)
	multipartWriter := multipart.NewWriter(body)
	createFormFile, _ := multipartWriter.CreateFormFile("avatar", "user.jpg")

	file, err := os.Open(`C:\Users\moham\Downloads\fc_page-0001.jpg`)
	assert.Nil(t, err)
	io.Copy(createFormFile, file)

	multipartWriter.Close()

	req := httptest.NewRequest("PUT", "/api/v1/users/avatars", body)
	req.Header.Add("Content-Type", multipartWriter.FormDataContentType())
	req.Header.Add("Authorization", "Bear "+token)
	rec := httptest.NewRecorder()
	helper.Router.ServeHTTP(rec, req)

	response := rec.Result()
	bytes, _ := io.ReadAll(response.Body)

	var mapResponse map[string]interface{}
	json.Unmarshal(bytes, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 401, response.StatusCode)
}

func TestLogoutSuccess(t *testing.T) {
	helper.CreateUserTest()
	defer helper.DeleteUserTest()
	token := helper.GetTokenAfterLogin()

	req := httptest.NewRequest("POST", "/api/v1/users/logout", nil)
	req.Header.Add("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	helper.Router.ServeHTTP(rec, req)

	response := rec.Result()
	bytes, _ := io.ReadAll(response.Body)

	var mapResponse map[string]interface{}
	json.Unmarshal(bytes, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "", mapResponse["data"].(map[string]any)["token"])
}

func TestLogoutErrorUnauthorized(t *testing.T) {
	helper.CreateUserTest()
	defer helper.DeleteUserTest()
	token := helper.GetTokenAfterLogin()

	req := httptest.NewRequest("POST", "/api/v1/users/logout", nil)
	req.Header.Add("Authorization", "Bear "+token)
	rec := httptest.NewRecorder()
	helper.Router.ServeHTTP(rec, req)

	response := rec.Result()
	bytes, _ := io.ReadAll(response.Body)

	var mapResponse map[string]interface{}
	json.Unmarshal(bytes, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 401, response.StatusCode)
}
