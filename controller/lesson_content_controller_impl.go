package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/service"
	"strconv"
)

type LessonContentControllerImpl struct {
	service.LessonContentService
}

func (c *LessonContentControllerImpl) GetByLessonTitleID(ctx *gin.Context) {
	ltID, err := strconv.Atoi(ctx.Param("ltID"))
	helper.PanicIfError(err)

	lessonContentsResponse := c.LessonContentService.FindByLessonTitleID(ltID)

	ctx.JSON(200,
		helper.APIResponse(200, "List of lesson contents", lessonContentsResponse),
	)
}

func (c *LessonContentControllerImpl) Update(ctx *gin.Context) {
	input := web.LessonContentCreateInput{}
	err := ctx.ShouldBind(&input)
	authorID := ctx.MustGet("current_author").(web.AuthorResponse).ID
	courseID, err := strconv.Atoi(ctx.Param("courseID"))
	ltID, err := strconv.Atoi(ctx.Param("ltID"))
	lcID, err := strconv.Atoi(ctx.Param("lcID"))
	helper.PanicIfError(err)
	input.AuthorID = authorID
	input.CourseID = courseID
	input.LessonTitleID = ltID

	fileHeader, err := ctx.FormFile("content")
	if err != nil {
		input.Content = ""
		lessonContentResponse := c.LessonContentService.Update(lcID, input)
		ctx.JSON(200,
			helper.APIResponse(200, "Lesson content successfully updated", lessonContentResponse),
		)
		return
	}
	pathContent := fmt.Sprintf("assets/contents/%s", fileHeader.Filename)
	input.Content = pathContent
	err = ctx.SaveUploadedFile(fileHeader, pathContent)
	helper.PanicIfError(err)

	lessonContentResponse := c.LessonContentService.Update(lcID, input)

	ctx.JSON(200,
		helper.APIResponse(200, "Lesson content successfully updated", lessonContentResponse),
	)

}

func (c *LessonContentControllerImpl) Create(ctx *gin.Context) {
	input := web.LessonContentCreateInput{}
	err := ctx.ShouldBind(&input)
	authorID := ctx.MustGet("current_author").(web.AuthorResponse).ID
	courseID, err := strconv.Atoi(ctx.Param("courseID"))
	ltID, err := strconv.Atoi(ctx.Param("ltID"))
	helper.PanicIfError(err)
	input.AuthorID = authorID
	input.CourseID = courseID
	input.LessonTitleID = ltID

	fileHeader, err := ctx.FormFile("content")
	pathContent := fmt.Sprintf("assets/contents/%s", fileHeader.Filename)
	input.Content = pathContent
	err = ctx.SaveUploadedFile(fileHeader, pathContent)
	helper.PanicIfError(err)

	lessonContentResponse := c.LessonContentService.Create(input)

	ctx.JSON(200,
		helper.APIResponse(200, "Lesson content successfully created", lessonContentResponse),
	)
}

func NewLessonContentController(lessonContentService service.LessonContentService) LessonContentController {
	return &LessonContentControllerImpl{LessonContentService: lessonContentService}
}
