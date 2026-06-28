package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"ms-api/config"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Video struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	UserID    string    `json:"userID,omitempty"`
	Title     string    `json:"title,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty" gorm:"autoCreateTime"`
	UpdatedAt        time.Time      `json:"updatedAt,omitempty" gorm:"autoUpdateTime"`
	TotalViews       int64          `json:"totalViews" gorm:"column:total_views;->"`
	AverageRate      float64        `json:"averageRate" gorm:"column:average_rate;->"`
	File             multipart.File `json:"-" gorm:"-"`
	UploadedFilePath string         `json:"-" gorm:"-"`
	Ctx              *gin.Context   `json:"-" gorm:"-"`
}

// GetVideo は指定されたIDのビデオレコードをデータベースから取得する。
// videoID: 取得したいビデオのID
// 戻り値: ビデオオブジェクトとエラー
func GetVideo(videoID int) (*Video,error) {
	video := Video{}

	if err := DbConnection.First(&video,videoID).Error; err != nil {
		return nil,err
	}
	return &video,nil
}

// CreateVideo は指定されたユーザーIDとタイトルで新しいビデオレコードをデータベースに作成する。
// userID: ビデオを所有するユーザーのID
// title: ビデオのタイトル
// 戻り値: 作成されたビデオオブジェクトとエラー
func CreateVideo(userID,title string) (*Video,error) {
	v := &Video{
		UserID: userID,
		Title: title,
	}
	if err := DbConnection.Create(&v).Error; err != nil {
		return nil,err
	}
	return v,nil
}

// Update は現在のビデオオブジェクトの内容をデータベースに保存して更新する。
// 戻り値: 更新に失敗した場合のエラー
func (v *Video) Update() error {
	if err := DbConnection.Save(v).Error; err != nil {
		return err
	}
	return nil
}

// Delete は現在のビデオレコードをデータベースから削除する。
// 戻り値: 削除に失敗した場合のエラー
func (v *Video) Delete() error {
	if err := DbConnection.Delete(v).Error; err != nil {
		return err
	}
	return nil
}

// makeProcessingDirectory はビデオIDをもとにビデオ処理用のディレクトリを作成する。
// ビデオIDをディレクトリ名とした処理用フォルダを設定パスに作成する。
// 戻り値: 作成したディレクトリのパスとエラー
func (v *Video) makeProcessingDirectory() (string,error) {
	path := filepath.Join(config.Config.AssetsUploadDirPath,strconv.Itoa(v.ID))
	if err := os.MkdirAll(path,0755); err != nil {
		return "",err
	}
	return path,nil
}

// saveUploadedVideo はアップロードされたビデオファイルを処理用ディレクトリに保存する。
// 処理ディレクトリを作成してから、アップロードファイルをコピーして保存する。
// また、保存したファイルの絶対パスをv.UploadedFilePathに設定する。
// 戻り値: ファイル保存に失敗した場合のエラー
func (v *Video) saveUploadedVideo() error {
	processingDirectory,err := v.makeProcessingDirectory()
	if err != nil {
		return err
	}
	uploadedVideoFilePath := filepath.Join(processingDirectory, config.Config.AssetsThumbnailFileName)

	newFile,err := os.Create(uploadedVideoFilePath)
	if err != nil {
		return err
	}

	defer newFile.Close()
	if _,err := io.Copy(newFile,v.File); err != nil {
		return err
	}
	absUploadedVideoFilePath, err := filepath.Abs(uploadedVideoFilePath)

	if err != nil {
		return err
	}
	v.UploadedFilePath = absUploadedVideoFilePath
	return nil
}

// ConvertVideo はアップロードされたビデオを保存してから、外部スクリプトを使用してビデオを変換する。
// saveUploadedVideoでファイルを保存した後、設定されたスクリプトを実行して指定解像度に変換する。
// 変換完了後、アップロード元のファイルは削除される。
// 戻り値: ビデオ変換に失敗した場合のエラー
func (v *Video) ConvertVideo() error {
	if err := v.saveUploadedVideo(); err != nil {
		return err
	}

	dstDirPath,_ := filepath.Abs(filepath.Join(config.Config.AssetsDirPath,strconv.Itoa(v.ID)))

	cmd := exec.CommandContext(v.Ctx,"/bin/sh",config.Config.ConvertVideoScriptFilePath,
	v.UploadedFilePath,dstDirPath,config.Config.ConvertVideoResolution)

	log.Println(cmd)
	if err := cmd.Run(); err != nil {
		return err
	}
	os.RemoveAll(v.UploadedFilePath)

	return nil

}

// VideoUpload はアップロードされたビデオファイルとメタデータで新しいビデオレコードをデータベースに作成する。
// トランザクション処理を使用して、ctx.Done()でリクエストキャンセルが検出された場合はロールバックされる。
// ctx: Ginのリクエストコンテキスト（キャンセル検出用）
// userID: ビデオを所有するユーザーのID
// title: ビデオのタイトル
// file: アップロードされたビデオファイル
// 戻り値: 作成されたビデオオブジェクトとエラー
func VideoUpload(ctx *gin.Context, userID,title string, file multipart.File) (*Video, error) {
	tx := DbConnection.Begin()

	v := Video{
		UserID: userID,
		Title: title,
		File: file,
		Ctx: ctx,
	}

	if err := tx.Create(&v).Error; err != nil {
		return nil,err
	}

	select {
	case <-ctx.Done():
		tx.Rollback()
		return nil,ctx.Err()

	default:
		tx.Commit()
		return &v,nil
	}

}

type VideoSortType string

const (
	VideoSortTypePopular   VideoSortType = "popular"
	VideoSortTypeRecommended VideoSortType= "recommended"
)

func (v VideoSortType) Valid() error {
	switch v {
	case VideoSortTypePopular, VideoSortTypeRecommended:
		return nil
	default:
	return errors.New("invalid type")
	}

}

func VideoList(sortType VideoSortType, limit int, userID string) ([]Video,error) {
	switch sortType{
	case VideoSortTypePopular:
		videos := []Video{}
		//.ModelはDbConnectionの作業用インスタンスをはやす
		query := DbConnection.Model(&Video{}).
		Select("videos.*,COALESCE(v.view_count,0) AS total_views, COALESCE(r.avg_rate,o) AS average_rate").
		Joins("LEFT JOIN (SELECT video_id, count(id) AS view_count FROM views GROUP BY video_id) AS v ON videos.id = v.video_id").
		Joins("LEFT JOIN (SELECT video_id, avg(value) AS avg_rate FOM rates GROUP BY video_id) AS r ON videos.id = r.video_id").
		Order("COALESCE(v.view_count,0) DESC")

		if limit != 0{
			query.Limit(limit)
		}
		if err := query.Find(&videos).Error; err != nil {
			return nil,err
		}
		return videos,nil

	case VideoSortTypeRecommended:
	url := fmt.Sprintf("%s/videos/recommende", config.Config.RecommendationAPIURL())
	request,err := http.NewRequest("GET",url,nil)
	if err != nil {
		log.Println(err)
		return nil,err
	}
	//.Queryはからのurls.value型定義
	params := request.URL.Query()
	if userID != "" {
		params.Add("user_id",userID)
	}
	if limit != 0{
		params.Add("limit",strconv.Itoa(limit))
	}
	request.URL.RawQuery = params.Encode()
	log.Printf("url=%s params=%s",url,params)
	response,err := http.DefaultClient.Do(request)

	if err != nil {
		return nil,err
	}
	defer response.Body.Close()

	byteArray,_ := io.ReadAll(response.Body)
	videos := []Video{}
	//&videos(Goの構造体)にjsonのbyte列を書き込み、バイト列ではなくなってGoの構造体へ
	err = json.Unmarshal(byteArray,&videos)
	if err != nil {
		return nil,err
	}
	if len(videos) > 0 {
		ids := make([]int,len(videos))
		for i,v := range videos {
			//[]intに動画のidを入れていく
			ids[i] = v.ID
		}
		type stat struct {
			VideoID int `gorm:"column:video_id"`
			TotalViews int64 `gorm:"column:total_views"`
			AverageRate float64 `gorm:"column:average_rate"`
		}
		stats := []stat{}
		DbConnection.Raw(
			"SELECT v.video_id, COALESCE(v.view_count, 0) AS total_views, COALESCE(r.avg_rate, 0) AS average_rate "+
					"FROM (SELECT video_id, count(id) AS view_count FROM views WHERE video_id IN ? GROUP BY video_id) AS v "+
					"FULL OUTER JOIN (SELECT video_id, avg(value) AS avg_rate FROM rates WHERE video_id IN ? GROUP BY video_id) AS r ON v.video_id = r.video_id",
				ids, ids,
		).Scan(&stats)
		statMap := make(map[int]stat,len(stats))

		for _, s := range stats {
			statMap[s.VideoID] = s
		}
		for i,v := range videos {
			if s,ok := statMap[v.ID]; ok {
				videos[i].TotalViews = s.TotalViews
				videos[i].AverageRate = s.AverageRate
			}
		}
	}
	return videos,nil
default:
	return []Video{},nil
	}

}