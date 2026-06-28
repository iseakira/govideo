package models

import (
	"bytes"
	"io"
	"ms-api/config"
	"os"
	"path/filepath"
	"strconv"
)

type Thumbnail struct {
	File io.ReadCloser
}

func NewThumbnail(file *os.File) *Thumbnail {
	return &Thumbnail{
		File:file,
	}
}

func (t *Thumbnail) Bytes() []byte {
	buffer := new(bytes.Buffer)
	io.Copy(buffer,t.File)
	return buffer.Bytes()
}

func (t *Thumbnail) Close() error {
	return t.File.Close()
}

func GetThumbnail(videoID int) (*Thumbnail,error) {
	basePath := filepath.Join(config.Config.AssetsDirPath,strconv.Itoa(videoID))
	thumbnailFilePath := filepath.Join(basePath, "thumbnail", config.Config.AssetsThumbnailFileName)
	file,err := os.Open(thumbnailFilePath)
	if err != nil {
		return nil,err
	}
	thumbnail := NewThumbnail(file)
	return thumbnail,nil

}