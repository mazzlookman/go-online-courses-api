package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go-pzn-restful-api/auth"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/service"
	"net/http"
	"strings"
)

func JwtAuthMiddleware(jwtAuth auth.JwtAuth, userService service.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		getHeader := ctx.GetHeader("Authorization")
		if !strings.Contains(getHeader, "Bearer") {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				helper.APIResponse(http.StatusUnauthorized, "Unauthorized", "Who you are?"))
			return
		}
		valueHeader := strings.Split(getHeader, " ")
		token := valueHeader[1]

		validateJwtToken, err := jwtAuth.ValidateJwtToken(token)
		if !validateJwtToken.Valid || err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				helper.APIResponse(http.StatusUnauthorized, "Unauthorized", "Who you are?"))
			return
		}

		claims := validateJwtToken.Claims.(jwt.MapClaims)
		userID := int(claims["user_id"].(float64))

		findByID := userService.FindByID(userID)
		ctx.Set("current_user", findByID)
	}
}
