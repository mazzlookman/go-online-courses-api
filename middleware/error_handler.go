package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-pzn-restful-api/helper"
	"net/http"
)

func ErrorHandler(ctx *gin.Context, err any) {
	if badRequestError(ctx, err) {
		return
	}

	if validationErrors(ctx, err) {
		return
	}

	if notFoundError(ctx, err) {
		return
	}

	if unauthorizedError(ctx, err) {
		return
	}

	internalServerError(ctx, err)
}

func badRequestError(ctx *gin.Context, err any) bool {
	unauthorized, ok := err.(helper.BadRequestError)
	if ok {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			helper.APIResponse(http.StatusBadRequest, "Bad Request", helper.NewResponseErrorKey(unauthorized.Error)),
		)
		return true
	} else {
		return false
	}
}

func unauthorizedError(ctx *gin.Context, err any) bool {
	unauthorized, ok := err.(helper.UnauthorizedError)
	if ok {
		ctx.Writer.WriteHeader(http.StatusUnauthorized)
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			helper.APIResponse(http.StatusUnauthorized, "Unauthorized", helper.NewResponseErrorKey(unauthorized.Error)),
		)
		return true
	} else {
		return false
	}
}

func notFoundError(ctx *gin.Context, err any) bool {
	notFound, ok := err.(helper.NotFoundError)
	if ok {
		ctx.Writer.WriteHeader(http.StatusNotFound)
		ctx.AbortWithStatusJSON(
			http.StatusNotFound,
			helper.APIResponse(http.StatusNotFound, "Not Found", helper.NewResponseErrorKey(notFound.Error)),
		)
		return true
	} else {
		return false
	}
}

func validationErrors(ctx *gin.Context, err any) bool {
	errors, ok := err.(validator.ValidationErrors)
	if ok {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			helper.APIResponse(http.StatusBadRequest, "Bad Request", helper.NewResponseErrorKey(errors.Error())),
		)
		return true
	} else {
		return false
	}
}

func internalServerError(ctx *gin.Context, err any) {
	ctx.Writer.WriteHeader(http.StatusInternalServerError)
	ctx.AbortWithStatusJSON(
		http.StatusInternalServerError,
		helper.APIResponse(http.StatusInternalServerError, "Internal Server Error", helper.NewResponseErrorKey(fmt.Sprintf("%v", err))),
	)
}
