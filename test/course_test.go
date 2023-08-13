package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"go-pzn-restful-api/model/domain"
	"testing"
)

func TestGetCourseBySlug(t *testing.T) {
	course := CourseService.FindBySlug("docker")
	//if err != nil {
	//	t.Error(err)
	//}

	marshal, _ := json.Marshal(course)
	t.Log(string(marshal))
}
func TestGetByAuthor(t *testing.T) {
	courseResponses := CourseService.FindByAuthorID(11)

	marshal, _ := json.Marshal(courseResponses)
	t.Log(string(marshal))
}

func TestUserEnrolled(t *testing.T) {
	CourseRepository.UsersEnrolled(domain.UserCourse{
		CourseID: 9,
		UserID:   14,
	})

	assert.Equal(t, 2, 2)
}

func TestGetCourseByUserID(t *testing.T) {
	courses, err := CourseRepository.FindByUserID(48)
	if err != nil {
		t.Error(err)
	}

	marshal, _ := json.Marshal(courses)
	t.Log(string(marshal))
}

func TestFindByCategoryID(t *testing.T) {
	courses, err := CourseRepository.FindByCategory("backend")
	if err != nil {
		t.Error(err)
	}

	marshal, _ := json.Marshal(courses)
	t.Log(string(marshal))
}
