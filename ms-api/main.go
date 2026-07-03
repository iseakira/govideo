package main

import (
	"io"
	"log"
	"ms-api/app/controllers"
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
	engine.Run(":8080")

	v1 := engine.Group("/api/v1")

	v1.GET("/health", controllers.Health)


}

func main() {
	LoggingSettings(config.Config.LogFile)
	StartServer()
}