package controller

import (
	"github.com/gin-gonic/gin"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/service"
)

type CourseControllerImpl struct {
	service.CourseService
}

func (c *CourseControllerImpl) GetBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")
	courseResponse := c.CourseService.FindBySlug(slug)
	ctx.JSON(200,
		helper.APIResponse(200, "Detail of course", courseResponse),
	)
}

func (c *CourseControllerImpl) Create(ctx *gin.Context) {
	request := web.CourseInputRequest{}
	err := ctx.ShouldBindJSON(&request)
	helper.PanicIfError(err)

	authorID := ctx.MustGet("current_author").(web.AuthorResponse).ID
	request.AuthorID = authorID

	courseResponse := c.CourseService.Create(request)
	ctx.JSON(200,
		helper.APIResponse(200, "Course has been created", courseResponse),
	)
}

func NewCourseController(courseService service.CourseService) CourseController {
	return &CourseControllerImpl{CourseService: courseService}
}
