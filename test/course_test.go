package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/repository"
	"go-pzn-restful-api/service"
	"go-pzn-restful-api/test/util"
	"testing"
)

var (
	courseRepository = repository.NewCourseRepository(util.Db)
	courseService    = service.NewCourseService(courseRepository)
)

func TestGetCourseBySlug(t *testing.T) {
	course := courseService.FindBySlug("docker")
	//if err != nil {
	//	t.Error(err)
	//}

	marshal, _ := json.Marshal(course)
	t.Log(string(marshal))
}
func TestGetByAuthor(t *testing.T) {
	courseResponses := courseService.FindByAuthorID(11)

	marshal, _ := json.Marshal(courseResponses)
	t.Log(string(marshal))
}

func TestUserEnrolled(t *testing.T) {
	usersEnrolled := courseRepository.UsersEnrolled(domain.UserCourse{
		CourseID: 9,
		UserID:   14,
	})

	assert.Equal(t, 2, usersEnrolled)
}