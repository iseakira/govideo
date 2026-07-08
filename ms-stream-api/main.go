package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"ms-stream-api/app/controllers"
	"ms-stream-api/config"

	"github.com/gin-gonic/gin"
	timeout "github.com/vearne/gin-timeout"
)

func LoggingSettings(logFile string) {
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("file=logFile err=%s", err.Error())
	}
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
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
				log.Println("Request Timeout : ", r.URL.String())
			})),
	)
	streamController := controllers.NewStream()

	v1 := engine.Group("/api/v1")
	{
		v1.GET("/health", controllers.Health)

		stream := v1.Group("/stream")
		{
			stream.GET("/ping", func(context *gin.Context) {
				context.JSON(200, gin.H{
					"message": "pong",
				})
			})
			stream.GET(":id/playlist", streamController.GetM3u8File)
			stream.GET(":id/:segName", streamController.GetHlsFile)
		}
	}
	engine.Run(":8081")

}

func main() {
	LoggingSettings(config.Config.LogFile)
	StartServer()
}