package test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-pzn-restful-api/app"
	"go-pzn-restful-api/repository"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserRegister(t *testing.T) {
	r := app.NewRouter()
	reqBody := strings.NewReader(`{
									  "name": "Aqib",
									  "email": "aqib@test.com",
									  "password": "123"
									}`)
	req := httptest.NewRequest("POST", "/api/v1/users", reqBody)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)
	response := rec.Result()

	body, _ := io.ReadAll(response.Body)
	var mapResponse map[string]interface{}
	json.Unmarshal(body, &mapResponse)

	fmt.Println(mapResponse)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "Aqib", mapResponse["data"].(map[string]interface{})["name"])
}

func TestFindByEmailRepo(t *testing.T) {
	userRepository := repository.NewUserRepository(app.DBConnection())
	user, err := userRepository.FindByEmail("ucuptest.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
