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
		helper.APIResponse(200, "List of course", courseResponses),
	)
}

func (c *CourseControllerImpl) GetByUserID(ctx *gin.Context) {
	userID := ctx.MustGet("current_user").(web.UserResponse).ID
	courseResponses := c.CourseService.FindByUserID(userID)

	ctx.JSON(200,
		helper.APIResponse(200, "List of course", courseResponses),
	)
}

func (c *CourseControllerImpl) UploadBanner(ctx *gin.Context) {
	courseID, _ := strconv.Atoi(ctx.Param("courseID"))
	fileHeader, _ := ctx.FormFile("banner")

	pathFile := fmt.Sprintf("assets/images/avatars/%d-%s", courseID, fileHeader.Filename)
	uploadBanner := c.CourseService.UploadBanner(courseID, pathFile)

	ctx.SaveUploadedFile(fileHeader, pathFile)

	ctx.JSON(200,
		helper.APIResponse(200, "Banner is successfully uploaded",
			gin.H{"is_uploaded": uploadBanner}),
	)
}

func (c *CourseControllerImpl) UserEnrolled(ctx *gin.Context) {
	user := ctx.MustGet("current_user").(web.UserResponse)
	courseID, err := strconv.Atoi(ctx.Param("courseID"))
	helper.PanicIfError(err)

	c.CourseService.UserEnrolled(user.ID, courseID)

	ctx.JSON(200,
		helper.APIResponse(200, "Success to enrolled",
			gin.H{"enrolled by": user.Name}),
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
