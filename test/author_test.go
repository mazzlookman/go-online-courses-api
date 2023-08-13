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
	"strings"
	"testing"
)

func TestAuthorRegisterSuccess(t *testing.T) {
	defer DeleteAuthorTest()
	body := new(bytes.Buffer)
	multipartWriter := multipart.NewWriter(body)
	multipartWriter.WriteField("name", "author")
	multipartWriter.WriteField("email", "author@author.com")
	multipartWriter.WriteField("password", "123")
	multipartWriter.WriteField("profile", "author")
	multipartWriter.CreateFormFile("avatar", "author.jpg")
	multipartWriter.Close()

	req := httptest.NewRequest("POST", "/api/v1/authors", body)
	req.Header.Add("Content-Type", multipartWriter.FormDataContentType())
	w := httptest.NewRecorder()
	Router.ServeHTTP(w, req)

	response := w.Result()
	bytes, _ := io.ReadAll(response.Body)

	var mapResponse map[string]interface{}
	json.Unmarshal(bytes, &mapResponse)

	log.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "author@author.com", mapResponse["data"].(map[string]any)["email"])
}

func TestAuthorLoginSuccess(t *testing.T) {
	CreateAuthorTest()
	defer DeleteAuthorTest()
	reqBody := strings.NewReader(`{"email": "author@author.com","password": "123"}`)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/authors/login", reqBody)
	Router.ServeHTTP(w, req)

	response := w.Result()
	bytes, _ := io.ReadAll(response.Body)

	var mapResponse map[string]interface{}
	json.Unmarshal(bytes, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "author", mapResponse["data"].(map[string]interface{})["name"])
	assert.NotNil(t, mapResponse["data"].(map[string]any)["token"])
}

func TestAuthorLoginErrorValidation(t *testing.T) {
	CreateAuthorTest()
	defer DeleteAuthorTest()
	reqBody := strings.NewReader(`{"email": "authorauthor.com","password": ""}`)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/authors/login", reqBody)
	Router.ServeHTTP(w, req)

	response := w.Result()
	bytes, _ := io.ReadAll(response.Body)

	var mapResponse map[string]interface{}
	json.Unmarshal(bytes, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 400, response.StatusCode)
}

func TestAuthorLoginErrorUsernameOrPasswordIsWrong(t *testing.T) {
	CreateAuthorTest()
	defer DeleteAuthorTest()
	reqBody := strings.NewReader(`{"email": "a@author.com","password": "123"}`)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/authors/login", reqBody)
	Router.ServeHTTP(w, req)

	response := w.Result()
	bytes, _ := io.ReadAll(response.Body)

	var mapResponse map[string]interface{}
	json.Unmarshal(bytes, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 404, response.StatusCode)
}

func TestAuthorLogoutSuccess(t *testing.T) {
	CreateAuthorTest()
	token := GetAuthorToken()
	defer DeleteAuthorTest()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/authors/logout", nil)
	req.Header.Add("Authorization", "Bearer "+token)
	Router.ServeHTTP(w, req)

	response := w.Result()
	bytes, _ := io.ReadAll(response.Body)

	var mapResponse map[string]interface{}
	json.Unmarshal(bytes, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
}

func TestAuthorLogoutErrorUnauthorized(t *testing.T) {
	CreateAuthorTest()
	token := GetAuthorToken()
	defer DeleteAuthorTest()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/authors/logout", nil)
	req.Header.Add("Authorization", "Bear "+token)
	Router.ServeHTTP(w, req)

	response := w.Result()
	bytes, _ := io.ReadAll(response.Body)

	var mapResponse map[string]interface{}
	json.Unmarshal(bytes, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 401, response.StatusCode)
}
