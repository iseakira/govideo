package controllers

import (
	"log"
	"net/http"
	"strconv"

	"ms-api/app/models"
	"ms-api/auth"

	"github.com/gin-gonic/gin"
)


type Rate struct {

}

func NewRate() *Rate {
	return &Rate{}
}

func (c *Rate) Get(ctx *gin.Context) {
	userID, err := auth.GetRequestUserID(ctx,true)
	if err != nil {
		log.Println(err)
		APIError(ctx, http.StatusUnauthorized, err)
		return
	}

	videoIDStr := ctx.Param("id")
	videoID,err := strconv.Atoi(videoIDStr)

	if err != nil {
		log.Println(err)
		APIError(ctx, http.StatusBadRequest, err)
		return
	}

	rate,err := models.GetRate(userID,videoID)

		if err != nil {
		log.Println(err)
		APIError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, rate)
}


func (c *Rate) Update(ctx *gin.Context) {
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
	rate := &models.Rate{
		UserID:  userID,
		VideoID: videoID,
	}

	if err := ctx.BindJSON(&rate); err != nil {
		APIError(ctx, http.StatusInternalServerError, err)
		return
	}

	rate.UserID = userID
	rate.VideoID = videoID

		_, err = models.CreateOrUpdateRate(rate.UserID, rate.VideoID, rate.Value)
	if err != nil {
		log.Println(err)
		APIError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

type ResponseRateAverage struct {
	Value float32 `json:"value"`
}


func (c *Rate) Average(ctx *gin.Context) {
	videoIDStr := ctx.Param("id")
	videoID, err := strconv.Atoi(videoIDStr)
	if err != nil {
		log.Println(err)
		APIError(ctx, http.StatusBadRequest, err)
		return
	}
	value, err := models.RateAverage(videoID)
	if err != nil {
		log.Println(err)
	}
	resp := ResponseRateAverage{
		Value: value,
	}
	ctx.JSON(http.StatusOK, resp)
}