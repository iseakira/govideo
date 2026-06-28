package controllers

import (
	"errors"
	"fmt"
	"log"
	"ms-api/app/models"
	"ms-api/auth"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Video struct {
}

func NewVideo() *Video {
	return &Video{}
}
//1. 状態（依存）を持つ／持たせたいとき
// → DB接続、S3クライアントなどをフィールドに入れる

//2. 関連する操作を1つの型にグループ化したいとき
  // → v.Upload, v.List, v.Get をまとめてルーターに登録
func (c *Video) Upload(ctx *gin.Context) {
	userID, err := auth.GetRequestUserID(ctx,true)
	if err != nil {
		log.Println(err)
		APIError(ctx,http.StatusUnauthorized,err)
		return
	}

	file,_, err := ctx.Request.FormFile("file")
	if err != nil {
		log.Println(err)
		APIError(ctx,http.StatusInternalServerError,err)
		return
	}

	title := ctx.PostForm("title")
	video,err := models.VideoUpload(ctx,userID,title,file)

	if err != nil {
		log.Println(err)
		APIError(ctx,http.StatusInternalServerError,err)
		return
	}
	ctx.JSON(http.StatusCreated,video)
}

type VideoSortType string

const (
	VideoSortTypePopular VideoSortType = "popular"
	VideoSortTypeRecommended VideoSortType="recommended"
)


func (v VideoSortType) Valid() error {
	switch v {
	case VideoSortTypePopular,VideoSortTypeRecommended:
		return nil
	default:
	err := fmt.Sprintf("invalid type %s",v)
	return errors.New(err)
	}
}


func (c *Video) List(ctx *gin.Context) {
	sortTypeStr := ctx.Query("sortType")
	limitStr := ctx.Query("limit")
	var limit int

	if limitStr != "" {
		value,err := strconv.Atoi(limitStr)
		if err != nil {
			log.Println(err)
			APIError(ctx,http.StatusBadRequest,err)
			return
		} else {
			limit=value
		}
	}

	sortType := models.VideoSortType(sortTypeStr)
	if err := sortType.Valid(); err != nil {
		log.Println(err)
		APIError(ctx,http.StatusBadRequest,err)
		return
	}

	userID := ctx.Query("user_id")
	if sortType == models.VideoSortTypeRecommended && userID == "" {
		userID = auth.AnonymousUserID
	}
	videos,err := models.VideoList(sortType,limit,userID)

	if err != nil {
		log.Println(err)
		APIError(ctx,http.StatusInternalServerError,err)
		return
	}
	ctx.JSON(http.StatusOK,videos)

	select {
	case <-ctx.Done():
		log.Println(ctx.Err())
		return
	default:
		return
	}
}

func (c *Video) Get(ctx *gin.Context) {
	videoIdStr := ctx.Param("id")
	videoId,err := strconv.Atoi(videoIdStr)
	if err != nil {
		log.Println(err)
		APIError(ctx,http.StatusBadRequest,err)
		return
	}
	video,err := models.GetVideo(videoId)

	if err != nil {
		log.Println(err)
		APIError(ctx,http.StatusNotFound,err)
		return
	}
	ctx.JSON(http.StatusOK,video)
}






