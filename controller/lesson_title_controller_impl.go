package controller

import (
	"github.com/gin-gonic/gin"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/service"
	"strconv"
)

type LessonTitleControllerImpl struct {
	service.LessonTitleService
}

func (c *LessonTitleControllerImpl) Update(ctx *gin.Context) {
	input := web.LessonTitleCreateInput{}
	err := ctx.ShouldBindJSON(&input)
	helper.PanicIfError(err)
	ltID, _ := strconv.Atoi(ctx.Param("ltID"))

	lessonTitleResponse := c.LessonTitleService.Update(ltID, input)
	ctx.JSON(200,
		helper.APIResponse(200, "Lesson title is successfully updated", lessonTitleResponse))
}

func (c *LessonTitleControllerImpl) GetByCourseID(ctx *gin.Context) {
	courseID, _ := strconv.Atoi(ctx.Param("courseID"))
	lessonTitlesResponse := c.LessonTitleService.FindByCourseID(courseID)

	ctx.JSON(200,
		helper.APIResponse(200, "List of lesson titles", lessonTitlesResponse))
}

func (c *LessonTitleControllerImpl) Create(ctx *gin.Context) {
	input := web.LessonTitleCreateInput{}
	err := ctx.ShouldBindJSON(&input)
	helper.PanicIfError(err)

	courseID, _ := strconv.Atoi(ctx.Param("courseID"))
	input.CourseID = courseID

	authorID := ctx.MustGet("current_author").(web.AuthorResponse).ID
	input.AuthorID = authorID

	lessonTitleResponse := c.LessonTitleService.Create(input)
	ctx.JSON(200,
		helper.APIResponse(200, "Lesson title is successfully created", lessonTitleResponse))
}

func NewLessonTitleController(lessonTitleService service.LessonTitleService) LessonTitleController {
	return &LessonTitleControllerImpl{LessonTitleService: lessonTitleService}
}
