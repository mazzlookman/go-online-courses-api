package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/service"
	"strconv"
)

type CourseControllerImpl struct {
	service.CourseService
}

func (c *CourseControllerImpl) GetByCategory(ctx *gin.Context) {
	courseResponses := c.CourseService.FindByCategory(ctx.Param("categoryName"))

	ctx.JSON(200,
		helper.APIResponse(200, "List of courses", courseResponses),
	)
}

func (c *CourseControllerImpl) GetByUserId(ctx *gin.Context) {
	userId := ctx.MustGet("current_user").(web.UserResponse).Id
	courseResponses := c.CourseService.FindByUserId(userId)

	ctx.JSON(200,
		helper.APIResponse(200, "List of courses", courseResponses),
	)
}

func (c *CourseControllerImpl) UploadBanner(ctx *gin.Context) {
	courseId, _ := strconv.Atoi(ctx.Param("courseId"))
	fileHeader, _ := ctx.FormFile("banner")

	pathFile := fmt.Sprintf("assets/images/banners/%d-%s", courseId, fileHeader.Filename)
	uploadBanner := c.CourseService.UploadBanner(courseId, pathFile)

	ctx.SaveUploadedFile(fileHeader, pathFile)

	ctx.JSON(200,
		helper.APIResponse(200, "Banner is successfully uploaded",
			gin.H{"is_uploaded": uploadBanner}),
	)
}

func (c *CourseControllerImpl) UserEnrolled(ctx *gin.Context) {
	user := ctx.MustGet("current_user").(web.UserResponse)
	courseId, err := strconv.Atoi(ctx.Param("courseId"))
	helper.PanicIfError(err)

	c.CourseService.UserEnrolled(user.Id, courseId)

	ctx.JSON(200,
		helper.APIResponse(200, "Success to enrolled",
			gin.H{"enrolled_by": user.Name}),
	)
}

func (c *CourseControllerImpl) GetAll(ctx *gin.Context) {
	courseResponses := c.CourseService.FindAll()
	ctx.JSON(200,
		helper.APIResponse(200, "List of courses", courseResponses),
	)
}

func (c *CourseControllerImpl) GetByAuthorId(ctx *gin.Context) {
	param := ctx.Param("authorId")
	authorId, _ := strconv.Atoi(param)
	courseResponse := c.CourseService.FindByAuthorId(authorId)
	ctx.JSON(200,
		helper.APIResponse(200, "List of courses", courseResponse),
	)
}

func (c *CourseControllerImpl) GetBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")
	courseResponse := c.CourseService.FindBySlug(slug)
	ctx.JSON(200,
		helper.APIResponse(200, "Course detail", courseResponse),
	)
}

func (c *CourseControllerImpl) Create(ctx *gin.Context) {
	request := web.CourseCreateInput{}
	err := ctx.ShouldBindJSON(&request)
	helper.PanicIfError(err)

	authorId := ctx.MustGet("current_author").(web.AuthorResponse).Id
	request.AuthorId = authorId

	courseResponse := c.CourseService.Create(request)
	ctx.JSON(200,
		helper.APIResponse(200, "Course has been created", courseResponse),
	)
}

func NewCourseController(courseService service.CourseService) CourseController {
	return &CourseControllerImpl{CourseService: courseService}
}
