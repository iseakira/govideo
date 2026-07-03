package controllers

import (
	"log"
	"net/http"
	"strconv"

	"ms-api/app/models"
	"ms-api/auth"

	"github.com/gin-gonic/gin"
)

type Views struct {
}

func NewViews() *Views {
	return &Views{}
}

type ResponseViewsCount struct {
	Value int64 `json:"value"`
}

func (c *Views) Total(ctx *gin.Context) {
	videoIDStr := ctx.Param("id")
	videoID, err := strconv.Atoi(videoIDStr)
	if err != nil {
		log.Println(err)
		APIError(ctx, http.StatusBadRequest, err)
		return
	}
	count, err := models.TotalViews(videoID)
	if err != nil {
		log.Println(err)
		APIError(ctx, http.StatusNotFound, err)
		return
	}
	resp := ResponseViewsCount{
		Value: count,
	}
	ctx.JSON(http.StatusOK, resp)

}


func (c *Views) Add(ctx *gin.Context) {
	userID, err := auth.GetRequestUserID(ctx, true)
	if err != nil {
		log.Println(err)
		APIError(ctx, http.StatusUnauthorized, err)
		return
	}
	videoIDStr := ctx.Param("id")
	videoID, err := strconv.Atoi(videoIDStr)
	if err != nil {
		log.Println(err)
		APIError(ctx, http.StatusBadRequest, err)
		return
	}
	views, err := models.CreateViews(userID, videoID)
	if err != nil {
		log.Println(err)
		APIError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, views)



}

