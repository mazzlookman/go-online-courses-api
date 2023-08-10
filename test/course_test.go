package test

import (
	"encoding/json"
	"go-pzn-restful-api/repository"
	"go-pzn-restful-api/test/util"
	"testing"
)

var (
	courseRepository = repository.NewCourseRepository(util.Db)
)

func TestGetCourse(t *testing.T) {
	course, err := courseRepository.FindBySlug("docker")
	if err != nil {
		t.Error(err)
	}

	marshal, _ := json.Marshal(course)
	t.Log(string(marshal))
}
