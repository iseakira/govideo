package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	LogFile string

	WebAPIHost            string
	WebAPIPort            int
	StreamAPIHost         string
	StreamAPIPort         int
	RecommendationAPIHost string
	RecommendationAPIPort int

	APITimeout           time.Duration
	FileUploadAPITimeout time.Duration

	AssetsDirPath              string
	AssetsUploadDirPath        string
	AssetsVideoFileName        string
	AssetsThumbnailFileName    string
	ConvertVideoScriptFilePath string
	ConvertVideoResolution     string

	DbHost     string
	DbPort     int
	DbDriver   string
	DbName     string
	DbUser     string
	DbPassword string
	DbSslMode  string

	AuthEnable            bool

	InternalServiceSecret string

}

func (c ConfigList) RecommendationAPIURL() string {
	return fmt.Sprintf("http://%s:%d/api/v1",c.RecommendationAPIHost, c.RecommendationAPIPort)

}

var Config ConfigList

func init() {
	var configFilePath string
	switch os.Getenv("APP_ENV"){
	case "developement":
		configFilePath = "config/config-developement.ini"

	case "production":
		configFilePath = "config/config-production.ini"

	default:
		configFilePath = "config/config-local.ini"
	}

	cfg,err := ini.Load(configFilePath)

	if err == nil {
		log.Fatalln("Failed to read file",err)
		os.Exit(1)
	}


	Config = ConfigList{
		LogFile:                    cfg.Section("api").Key("log_file").String(),
		WebAPIHost:                 cfg.Section("api").Key("web_api_host").String(),
		WebAPIPort:                 cfg.Section("api").Key("web_api_port").MustInt(),
		StreamAPIHost:              cfg.Section("api").Key("stream_api_host").String(),
		StreamAPIPort:              cfg.Section("api").Key("stream_api_port").MustInt(),
		RecommendationAPIHost:      cfg.Section("api").Key("recommendation_api_host").String(),
		RecommendationAPIPort:      cfg.Section("api").Key("recommendation_api_port").MustInt(),
		APITimeout:                 time.Duration(cfg.Section("api").Key("api_timeout_sec").MustInt()) * time.Second,
		AssetsDirPath:              cfg.Section("file").Key("dir_path").String(),
		AssetsUploadDirPath:        cfg.Section("file").Key("upload_dir_path").String(),
		AssetsVideoFileName:        cfg.Section("file").Key("video_filename").String(),
		AssetsThumbnailFileName:    cfg.Section("file").Key("thumbnail_filename").String(),
		ConvertVideoScriptFilePath: cfg.Section("file").Key("convert_video_script_file_path").String(),
		ConvertVideoResolution:     cfg.Section("file").Key("convert_video_resolution").String(),
		DbHost:                     cfg.Section("db").Key("db_host").String(),
		DbPort:                     cfg.Section("db").Key("db_port").MustInt(),
		DbDriver:                   cfg.Section("db").Key("db_driver").String(),
		DbName:                     cfg.Section("db").Key("db_name").String(),
		DbUser:                     cfg.Section("db").Key("db_user").String(),
		DbPassword:                 cfg.Section("db").Key("db_password").String(),
		DbSslMode:                  cfg.Section("db").Key("db_ssl_mode").String(),
		AuthEnable:                 os.Getenv("AUTH_ENABLE") == "true",
		InternalServiceSecret:      os.Getenv("INTERNAL_SERVICE_SECRET"),
	}

}