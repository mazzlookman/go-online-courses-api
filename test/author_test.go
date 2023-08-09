package test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/test/util"
	"io"
	"net/http/httptest"
	"testing"
)

func TestCreateAuthor(t *testing.T) {
	authorInputRequest := web.AuthorRegisterInput{
		Name:    "test",
		Profile: "Test Intro",
	}
	m, _ := json.Marshal(authorInputRequest)
	body := bytes.NewReader(m)

	req := httptest.NewRequest("POST", "/api/v1/authors", body)
	req.Header.Add("role", "admin")
	w := httptest.NewRecorder()

	util.R.ServeHTTP(w, req)

	response := w.Result()
	readAll, _ := io.ReadAll(response.Body)

	mapResponse := map[string]any{}
	json.Unmarshal(readAll, &mapResponse)

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "test", mapResponse["data"].(map[string]any)["name"])
}
