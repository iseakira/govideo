package controllers

import (
	"log"
	"net/http"
	"strconv"

	"ms-stream-api/app/models"

	"github.com/gin-gonic/gin"
)


type HTTPError struct {
	Code int `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`

}

func NewError(ctx *gin.Context,status int, err error) {
	httpError := HTTPError{
		Code: status,
		Message: err.Error(),
	}
	ctx.JSON(status,httpError)
}

type Stream struct {

}

func NewStream() *Stream {
	return &Stream{}
}

func (c *Stream) GetM3u8File(ctx *gin.Context) {
	videoIdStr := ctx.Param("id")
	videoId,err := strconv.Atoi(videoIdStr)

	if err != nil {
		log.Println(err)
		NewError(ctx,http.StatusNotFound,err)
		return
	}
	m3u8File, err := models.GetM3u8File(videoId)
	if err != nil {
		log.Println(err)
		NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.Data(http.StatusOK, "application/x-mpegURL", m3u8File.Bytes())

}

func (c *Stream) GetHlsFile(ctx *gin.Context) {
	videoIdStr := ctx.Param("id")
	segName := ctx.Param("segName")
	videoId,err := strconv.Atoi(videoIdStr)
	if err != nil {
		log.Println(err)
		NewError(ctx,http.StatusNotFound,err)
		return
	}
	hlsFile,err := models.GetHlsFile(videoId,segName)
	if err != nil {
		log.Println(err)
		NewError(ctx, http.StatusNotFound,err)
		return
	}
	ctx.Data(http.StatusOK, "video/Mp2T",hlsFile.Bytes())
}

