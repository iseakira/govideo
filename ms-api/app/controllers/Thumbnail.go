package controllers

import (
	"net/http"
	"strconv"

	"ms-api/app/models"

	"github.com/gin-gonic/gin"
)

type Thumbnail struct {
}

func NewThumbnail() *Thumbnail {
	return &Thumbnail{}
}

func (c *Thumbnail) GetThumbnail(ctx *gin.Context) {
	videoIdStr := ctx.Param("id")
	videoId, err := strconv.Atoi(videoIdStr)
	if err != nil {
		APIError(ctx, http.StatusBadRequest, err)
		return
	}
	thumbnailImage, err := models.GetThumbnail(videoId)
	if err != nil {
		APIError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.Data(http.StatusOK, "image/jpeg", thumbnailImage.Bytes())
}
