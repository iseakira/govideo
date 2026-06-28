package controllers

import "github.com/gin-gonic/gin"

type HTTPError struct {
	Code int `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

func APIError(ctx *gin.Context,statusCode int, err error) {
	httpEer := HTTPError{
		Code: statusCode,
		Message: err.Error(),
	}
	ctx.JSON(statusCode,httpEer)
}