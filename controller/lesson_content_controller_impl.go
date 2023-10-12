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

func (c *LessonContentControllerImpl) GetById(ctx *gin.Context) {
	isUserHas := ctx.MustGet("isUserHas").(bool)

	if isUserHas == false {
		ctx.AbortWithStatusJSON(200,
			helper.APIResponse(
				200, "List of lesson contents",
				gin.H{"is_user_has": isUserHas, "message": "You must unlock this course first"},
			),
		)
		return
	}

	lcId, _ := strconv.Atoi(ctx.Param("lcId"))
	findById := c.LessonContentService.FindById(lcId)

	ctx.JSON(200, helper.APIResponse(200, "Detail of lesson content", findById))
}

func (c *LessonContentControllerImpl) GetByLessonTitleId(ctx *gin.Context) {
	ltId, err := strconv.Atoi(ctx.Param("ltId"))
	helper.PanicIfError(err)

	lessonContentsResponse := c.LessonContentService.FindByLessonTitleId(ltId)

	ctx.JSON(200,
		helper.APIResponse(200, "List of lesson contents", lessonContentsResponse),
	)
}

func (c *LessonContentControllerImpl) Update(ctx *gin.Context) {
	input := web.LessonContentUpdateInput{}
	err := ctx.ShouldBindJSON(&input)
	authorId := ctx.MustGet("current_author").(web.AuthorResponse).Id
	courseId, err := strconv.Atoi(ctx.Param("courseId"))
	lcId, err := strconv.Atoi(ctx.Param("lcId"))
	helper.PanicIfError(err)

	input.AuthorId = authorId
	input.CourseId = courseId

	fileHeader, err := ctx.FormFile("content")
	if err != nil {
		input.Content = ""
		lessonContentResponse := c.LessonContentService.Update(lcId, input)
		ctx.JSON(200,
			helper.APIResponse(200, "Lesson content successfully updated", lessonContentResponse),
		)
		return
	}
	pathContent := fmt.Sprintf("assets/contents/%s", fileHeader.Filename)
	input.Content = pathContent
	err = ctx.SaveUploadedFile(fileHeader, pathContent)
	helper.PanicIfError(err)

	lessonContentResponse := c.LessonContentService.Update(lcId, input)

	ctx.JSON(200,
		helper.APIResponse(200, "Lesson content successfully updated", lessonContentResponse),
	)

}

func (c *LessonContentControllerImpl) Create(ctx *gin.Context) {
	input := web.LessonContentCreateInput{}
	err := ctx.ShouldBindJSON(&input)

	authorId := ctx.MustGet("current_author").(web.AuthorResponse).Id
	courseId, err := strconv.Atoi(ctx.Param("courseId"))
	ltId, err := strconv.Atoi(ctx.Param("ltId"))
	helper.PanicIfError(err)

	input.AuthorId = authorId
	input.CourseId = courseId
	input.LessonTitleId = ltId

	fileHeader, err := ctx.FormFile("content")
	pathContent := fmt.Sprintf("assets/contents/%d-%s", courseId, fileHeader.Filename)
	input.Content = pathContent

	lessonContentResponse := c.LessonContentService.Create(input)

	err = ctx.SaveUploadedFile(fileHeader, pathContent)
	helper.PanicIfError(err)

	ctx.JSON(200,
		helper.APIResponse(200, "Lesson content successfully created", lessonContentResponse),
	)
}

func NewLessonContentController(lessonContentService service.LessonContentService) LessonContentController {
	return &LessonContentControllerImpl{LessonContentService: lessonContentService}
}
