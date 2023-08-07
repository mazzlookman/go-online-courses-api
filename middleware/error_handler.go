package middleware

import (
	"github.com/gin-gonic/gin"
	"go-pzn-restful-api/helper"
	"net/http"
)

func ErrorHandler(ctx *gin.Context, err any) {
	internalServerError(ctx, err)
}

func internalServerError(ctx *gin.Context, err any) {
	ctx.Next()
	ctx.Writer.WriteHeader(http.StatusInternalServerError)
	ctx.AbortWithStatusJSON(
		http.StatusInternalServerError,
		helper.APIResponse(http.StatusInternalServerError, "Internal Server Error", err),
	)
}
