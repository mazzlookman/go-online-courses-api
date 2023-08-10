package test

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go-pzn-restful-api/test/util"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAuthorLogin(t *testing.T) {
	authorTest := util.CreateAuthorTest()
	reqBody := strings.NewReader(`{"email": "test@test.com","password": "123"}`)

	w := httptest.NewRecorder()
	context, eng := gin.CreateTestContext(w)
	req := httptest.NewRequest("POST", "/api/v1/authors/login", reqBody)
	eng.ServeHTTP(w, req)

	response := w.Result()
	bytes, _ := io.ReadAll(response.Body)

	var mapResponse map[string]interface{}
	json.Unmarshal(bytes, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, authorTest.ID, context.MustGet("author_id"))
	assert.Equal(t, "test", mapResponse["data"].(map[string]interface{})["name"])
	assert.NotNil(t, mapResponse["data"].(map[string]any)["token"])

	util.DeleteAuthorTest()
}

func TestDeleteAuthor(t *testing.T) {
	util.DeleteAuthorTest()
}

func TestGetByID(t *testing.T) {
	author := util.GetAuthorByID(11)
	marshal, _ := json.Marshal(author)
	t.Log(string(marshal))
}
