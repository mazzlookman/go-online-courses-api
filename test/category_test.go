package test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateCategorySuccess(t *testing.T) {
	CreateAuthorTest()
	token := GetAuthorToken()
	defer DeleteAuthorTest()
	defer DeleteCategory()

	request := strings.NewReader(`{"name":"backend"}`)
	req := httptest.NewRequest("POST", "/api/v1/categories", request)
	req.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	Router.ServeHTTP(w, req)
	response := w.Result()

	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "backend", mapResponse["data"].(map[string]interface{})["name"])
}

func TestCreateCategoryErrorUnauthorized(t *testing.T) {
	CreateAuthorTest()
	token := GetAuthorToken()
	defer DeleteAuthorTest()
	defer DeleteCategory()

	request := strings.NewReader(`{"name":"backend"}`)
	req := httptest.NewRequest("POST", "/api/v1/categories", request)
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

func TestCreateCategoryErrorValidation(t *testing.T) {
	CreateAuthorTest()
	token := GetAuthorToken()
	defer DeleteAuthorTest()
	defer DeleteCategory()

	request := strings.NewReader(`{"name":""}`)
	req := httptest.NewRequest("POST", "/api/v1/categories", request)
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
