package controller

import (
	"github.com/gin-gonic/gin"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/service"
	"strconv"
)

type CourseControllerImpl struct {
	service.CourseService
}

func (c *CourseControllerImpl) UserEnrolled(ctx *gin.Context) {
	userID := ctx.MustGet("current_user").(web.UserResponse).ID
	courseID, err := strconv.Atoi(ctx.Param("courseID"))
	helper.PanicIfError(err)

	userEnrolled := c.CourseService.UserEnrolled(userID, courseID)

	ctx.JSON(200,
		helper.APIResponse(200, "Success to enrolled",
			gin.H{"users_enrolled": userEnrolled}),
	)
}

func (c *CourseControllerImpl) GetAll(ctx *gin.Context) {
	courseResponses := c.CourseService.FindAll()
	ctx.JSON(200,
		helper.APIResponse(200, "List of course", courseResponses),
	)
}

func (c *CourseControllerImpl) GetByAuthorID(ctx *gin.Context) {
	param := ctx.Param("authorID")
	authorID, _ := strconv.Atoi(param)
	courseResponse := c.CourseService.FindByAuthorID(authorID)
	ctx.JSON(200,
		helper.APIResponse(200, "List of course", courseResponse),
	)
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
