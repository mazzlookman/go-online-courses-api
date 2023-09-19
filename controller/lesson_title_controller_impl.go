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
	ltId, _ := strconv.Atoi(ctx.Param("ltId"))

	lessonTitleResponse := c.LessonTitleService.Update(ltId, input)
	ctx.JSON(200,
		helper.APIResponse(200, "Lesson title is successfully updated", lessonTitleResponse))
}

func (c *LessonTitleControllerImpl) GetByCourseId(ctx *gin.Context) {
	courseId, _ := strconv.Atoi(ctx.Param("courseId"))
	lessonTitlesResponse := c.LessonTitleService.FindByCourseId(courseId)

	ctx.JSON(200,
		helper.APIResponse(200, "List of lesson titles", lessonTitlesResponse))
}

func (c *LessonTitleControllerImpl) Create(ctx *gin.Context) {
	input := web.LessonTitleCreateInput{}
	err := ctx.ShouldBindJSON(&input)
	helper.PanicIfError(err)

	courseId, _ := strconv.Atoi(ctx.Param("courseId"))
	input.CourseId = courseId

	authorId := ctx.MustGet("current_author").(web.AuthorResponse).Id
	input.AuthorId = authorId

	lessonTitleResponse := c.LessonTitleService.Create(input)
	ctx.JSON(200,
		helper.APIResponse(200, "Lesson title is successfully created", lessonTitleResponse))
}

func NewLessonTitleController(lessonTitleService service.LessonTitleService) LessonTitleController {
	return &LessonTitleControllerImpl{LessonTitleService: lessonTitleService}
}
