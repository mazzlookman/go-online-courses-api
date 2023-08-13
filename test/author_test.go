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

func TestAuthorLogin(t *testing.T) {
	CreateAuthorTest()
	reqBody := strings.NewReader(`{"email": "test@test.com","password": "123"}`)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/authors/login", reqBody)
	Router.ServeHTTP(w, req)

	response := w.Result()
	bytes, _ := io.ReadAll(response.Body)

	var mapResponse map[string]interface{}
	json.Unmarshal(bytes, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
	//assert.Equal(t, authorTest.ID, context.MustGet("author_id"))
	assert.Equal(t, "test", mapResponse["data"].(map[string]interface{})["name"])
	assert.NotNil(t, mapResponse["data"].(map[string]any)["token"])

	DeleteAuthorTest()
}

func TestDeleteAuthor(t *testing.T) {
	DeleteAuthorTest()
}

func TestGetByID(t *testing.T) {
	author := GetAuthorByID(11)
	marshal, _ := json.Marshal(author)
	t.Log(string(marshal))
}
