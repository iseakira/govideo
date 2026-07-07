package main

import (
	"io"
	"log"
	"ms-api/app/controllers"
	"ms-api/auth"
	"ms-api/config"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	timeout "github.com/vearne/gin-timeout"
)

func LoggingSettings(logFile string) {
	logfile, err := os.OpenFile(logFile,os.O_RDWR|os.O_CREATE|os.O_APPEND,0666)

	if err != nil {
		log.Fatalf("file=logFile err=%s",err.Error())
	}


	multiLogFile := io.MultiWriter(os.Stdout,logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.SetOutput(multiLogFile)
}

func StartServer() {
	engine := gin.Default()

	engine.Use(
		timeout.Timeout(
			timeout.WithTimeout(config.Config.APITimeout),
			timeout.WithErrorHttpCode(http.StatusRequestTimeout),
			timeout.WithCallBack(func(r *http.Request) {

				log.Println("Request Timeout:",r.URL.String())
			})),
	)

	rateController := controllers.NewRate()
	thumbnailController := controllers.NewThumbnail()
	videoController := controllers.NewVideo()
	viewsController := controllers.NewViews()
	internalRequestMiddleWare := auth.NewInternalRequestMiddleWare(config.Config.InternalServiceSecret)

	v1 := engine.Group("/api/v1")

	{
		v1.GET("/health", controllers.Health)

		videos := v1.Group("/videos")
		{
			videos.GET(":id/rate/average",rateController.Average)
			videos.GET(":id/thumbnail",thumbnailController.GetThumbnail)
			videos.GET(":id/views/total",viewsController.Total)

			if config.Config.AuthEnable {
				videos.GET("",videoController.List)
				videos.GET(":id", videoController.Get)
				videos.GET(":id/rate",internalRequestMiddleWare.CheckInternalAuth(),rateController.Get)
				videos.POST("upload",internalRequestMiddleWare.CheckInternalAuth(),videoController.Upload)
				videos.POST(":id/views",internalRequestMiddleWare.CheckInternalAuth(),viewsController.Add)
				videos.PATCH(":id/rate",internalRequestMiddleWare.CheckInternalAuth(),rateController.Update)
			}else {
				videos.GET("",videoController.List)
				videos.GET(":id",videoController.Get)
				videos.GET(":id/rate",rateController.Get)
				videos.POST("upload",videoController.Upload)
				videos.POST(":id/views",viewsController.Add)
				videos.PATCH(":id/rate",rateController.Update)
			}
		}

	}

	engine.Run(":8080")
}

func main() {
	LoggingSettings(config.Config.LogFile)
	StartServer()
}