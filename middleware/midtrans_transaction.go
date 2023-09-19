package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/service"
	"regexp"
	"strconv"
)

func MidtransPaymentMiddleware(courseService service.CourseService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.MustGet("current_user").(web.UserResponse).Id
		strCourseID := ctx.Param("courseId")
		courseID, _ := strconv.Atoi(strCourseID)
		cID := courseService.FindById(courseID).Id

		isUserHas := false

		courseIDByUserID := courseService.FindAllCourseIdByUserId(userID)
		for _, allCID := range courseIDByUserID {
			mustCompile := regexp.MustCompile(allCID)
			matchString := mustCompile.MatchString(fmt.Sprintf("%d", cID))
			if matchString {
				isUserHas = true
			}
		}

		ctx.Set("isUserHas", isUserHas)
	}
}
