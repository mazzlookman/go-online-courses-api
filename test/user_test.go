package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"os"
	"strings"
	"testing"
)

func TestUserRegisterSuccess(t *testing.T) {
	defer DeleteUserTest()
	reqBody := strings.NewReader(`{"name": "user","email": "user@user.com","password": "123"}`)
	req := httptest.NewRequest("POST", "/api/v1/users", reqBody)
	w := httptest.NewRecorder()

	Router.ServeHTTP(w, req)
	response := w.Result()

	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "user", mapResponse["data"].(map[string]interface{})["name"])
}

func TestUserRegisterErrorValidation(t *testing.T) {
	reqBody := strings.NewReader(`{"name": "test","email": "test@test.com","password": ""}`)
	req := httptest.NewRequest("POST", "/api/v1/users", reqBody)
	w := httptest.NewRecorder()

	Router.ServeHTTP(w, req)
	response := w.Result()

	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 400, response.StatusCode)
	//assert.Equal(t, "test", mapResponse["data"].(map[string]interface{})["name"])

	DeleteUserTest()
}

func TestUserLoginSuccess(t *testing.T) {
	CreateUserTest()
	reqBody := strings.NewReader(`{"email": "user@user.com","password": "123"}`)

	req := httptest.NewRequest("POST", "/api/v1/users/login", reqBody)
	w := httptest.NewRecorder()
	Router.ServeHTTP(w, req)

	response := w.Result()
	bytes, _ := io.ReadAll(response.Body)

	var mapResponse map[string]interface{}
	json.Unmarshal(bytes, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "user", mapResponse["data"].(map[string]interface{})["name"])
	assert.NotNil(t, mapResponse["data"].(map[string]any)["token"])

	DeleteUserTest()
}

func TestUserLoginErrorUsernameOrPasswordIsWrong(t *testing.T) {
	CreateUserTest()
	reqBody := strings.NewReader(`{"email": "user@user.com","password": "12"}`)

	req := httptest.NewRequest("POST", "/api/v1/users/login", reqBody)
	w := httptest.NewRecorder()
	Router.ServeHTTP(w, req)

	response := w.Result()
	bytes, _ := io.ReadAll(response.Body)

	var mapResponse map[string]interface{}
	json.Unmarshal(bytes, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 404, response.StatusCode)

	DeleteUserTest()
}

func TestUserLoginErrorValidation(t *testing.T) {
	CreateUserTest()
	reqBody := strings.NewReader(`{"email": "useruser.com","password": ""}`)

	req := httptest.NewRequest("POST", "/api/v1/users/login", reqBody)
	w := httptest.NewRecorder()
	Router.ServeHTTP(w, req)

	response := w.Result()
	bytes, _ := io.ReadAll(response.Body)

	var mapResponse map[string]interface{}
	json.Unmarshal(bytes, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 400, response.StatusCode)

	DeleteUserTest()
}

func TestGetUserByIdSuccess(t *testing.T) {
	CreateUserTest()
	token := GetTokenAfterLogin()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/users", nil)
	req.Header.Add("Authorization", "Bearer "+token)
	Router.ServeHTTP(w, req)

	response := w.Result()
	bytes, _ := io.ReadAll(response.Body)

	var mapResponse map[string]interface{}
	json.Unmarshal(bytes, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
	assert.NotNil(t, mapResponse["data"].(map[string]any)["id"])
	assert.Equal(t, "user", mapResponse["data"].(map[string]any)["name"])

	DeleteUserTest()
}

func TestGetUserByIdErrorUnauthorized(t *testing.T) {
	CreateUserTest()
	// random token
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMSJ9.vE_WTAksI9VxxurKeIDb-OETLYJzmecOuqs0FZJJ6kE"

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/users", nil)
	req.Header.Add("Authorization", "Bearer "+token)
	Router.ServeHTTP(w, req)

	response := w.Result()
	bytes, _ := io.ReadAll(response.Body)

	var mapResponse map[string]interface{}
	json.Unmarshal(bytes, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 401, response.StatusCode)

	DeleteUserTest()
}

func TestUploadAvatarSuccess(t *testing.T) {
	CreateUserTest()
	defer DeleteUserTest()
	token := GetTokenAfterLogin()

	body := new(bytes.Buffer)
	multipartWriter := multipart.NewWriter(body)

	// Register multipart header
	fileHeader := make(textproto.MIMEHeader)
	fileHeader.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "avatar", "fc_page-0001.jpg"))
	writer, err := multipartWriter.CreatePart(fileHeader)
	assert.Nil(t, err)

	// Copy file to file multipart writer
	file, err := os.Open(`C:\Users\moham\Downloads\fc_page-0001.jpg`)
	assert.Nil(t, err)
	io.Copy(writer, file)

	// close the writer before making the request
	multipartWriter.Close()

	req := httptest.NewRequest("PUT", "/api/v1/users/avatars", body)
	req.Header.Add("Content-Type", multipartWriter.FormDataContentType())
	req.Header.Add("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	Router.ServeHTTP(rec, req)

	response := rec.Result()
	bytes, _ := io.ReadAll(response.Body)

	var mapResponse map[string]interface{}
	json.Unmarshal(bytes, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
}

func TestUploadAvatarErrorUnAuthorized(t *testing.T) {
	CreateUserTest()
	defer DeleteUserTest()
	token := GetTokenAfterLogin()

	body := new(bytes.Buffer)
	multipartWriter := multipart.NewWriter(body)

	// Register multipart header
	fileHeader := make(textproto.MIMEHeader)
	fileHeader.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "avatar", "fc_page-0001.jpg"))
	writer, err := multipartWriter.CreatePart(fileHeader)
	assert.Nil(t, err)

	// Copy file to file multipart writer
	file, err := os.Open(`C:\Users\moham\Downloads\fc_page-0001.jpg`)
	assert.Nil(t, err)
	io.Copy(writer, file)

	// close the writer before making the request
	multipartWriter.Close()

	req := httptest.NewRequest("PUT", "/api/v1/users/avatars", body)
	req.Header.Add("Content-Type", multipartWriter.FormDataContentType())
	req.Header.Add("Authorization", "Bear "+token)
	rec := httptest.NewRecorder()
	Router.ServeHTTP(rec, req)

	response := rec.Result()
	bytes, _ := io.ReadAll(response.Body)

	var mapResponse map[string]interface{}
	json.Unmarshal(bytes, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 401, response.StatusCode)
}

func TestLogoutSuccess(t *testing.T) {
	CreateUserTest()
	defer DeleteUserTest()
	token := GetTokenAfterLogin()

	req := httptest.NewRequest("POST", "/api/v1/users/logout", nil)
	req.Header.Add("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	Router.ServeHTTP(rec, req)

	response := rec.Result()
	bytes, _ := io.ReadAll(response.Body)

	var mapResponse map[string]interface{}
	json.Unmarshal(bytes, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "", mapResponse["data"].(map[string]any)["token"])
}

func TestLogoutErrorUnauthorized(t *testing.T) {
	CreateUserTest()
	defer DeleteUserTest()
	token := GetTokenAfterLogin()

	req := httptest.NewRequest("POST", "/api/v1/users/logout", nil)
	req.Header.Add("Authorization", "Bear "+token)
	rec := httptest.NewRecorder()
	Router.ServeHTTP(rec, req)

	response := rec.Result()
	bytes, _ := io.ReadAll(response.Body)

	var mapResponse map[string]interface{}
	json.Unmarshal(bytes, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 401, response.StatusCode)
}
