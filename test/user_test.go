package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-pzn-restful-api/test/util"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"os"
	"strings"
	"testing"
)

func TestUserRegister(t *testing.T) {
	reqBody := strings.NewReader(`{"name": "test","email": "test@test.com","password": "123"}`)
	req := httptest.NewRequest("POST", "/api/v1/users", reqBody)
	w := httptest.NewRecorder()

	util.R.ServeHTTP(w, req)
	response := w.Result()

	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "test", mapResponse["data"].(map[string]interface{})["name"])

	util.DeleteUserTest()
}

func TestUserLogin(t *testing.T) {
	util.CreateUserTest()
	reqBody := strings.NewReader(`{"email": "test@test.com","password": "123"}`)

	req := httptest.NewRequest("POST", "/api/v1/users/login", reqBody)
	w := httptest.NewRecorder()
	util.R.ServeHTTP(w, req)

	response := w.Result()
	bytes, _ := io.ReadAll(response.Body)

	var mapResponse map[string]interface{}
	json.Unmarshal(bytes, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "test", mapResponse["data"].(map[string]interface{})["name"])
	assert.NotNil(t, mapResponse["data"].(map[string]any)["token"])

	util.DeleteUserTest()
}

func TestGetUser(t *testing.T) {
	util.CreateUserTest()
	token := util.GetTokenAfterLogin()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/users", nil)
	req.Header.Add("Authorization", "Bearer "+token)
	util.R.ServeHTTP(w, req)

	response := w.Result()
	bytes, _ := io.ReadAll(response.Body)

	var mapResponse map[string]interface{}
	json.Unmarshal(bytes, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
	assert.NotNil(t, mapResponse["data"].(map[string]any)["id"])
	assert.Equal(t, "test", mapResponse["data"].(map[string]any)["name"])

	util.DeleteUserTest()
}

func TestUploadAvatar(t *testing.T) {
	util.CreateUserTest()
	token := util.GetTokenAfterLogin()

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

	req := httptest.NewRequest("POST", "/api/v1/users/avatars", body)
	req.Header.Add("Content-Type", multipartWriter.FormDataContentType())
	req.Header.Add("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	util.R.ServeHTTP(rec, req)

	response := rec.Result()
	bytes, _ := io.ReadAll(response.Body)

	var mapResponse map[string]interface{}
	json.Unmarshal(bytes, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)

	util.DeleteUserTest()
}

func TestLogout(t *testing.T) {
	util.CreateUserTest()
	token := util.GetTokenAfterLogin()

	req := httptest.NewRequest("POST", "/api/v1/users/logout", nil)
	req.Header.Add("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	util.R.ServeHTTP(rec, req)

	response := rec.Result()
	bytes, _ := io.ReadAll(response.Body)

	var mapResponse map[string]interface{}
	json.Unmarshal(bytes, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "", mapResponse["data"].(map[string]any)["token"])
	util.DeleteUserTest()
}

func TestCreateNewUser(t *testing.T) {
	util.CreateUserTest()
}
