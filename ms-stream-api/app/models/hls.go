package models

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"ms-stream-api/config"
)

type HlsFile struct {
	File io.ReadCloser
}

func NewHlsFile(file *os.File) HlsFile {
	return HlsFile{
		File: file,
	}
}

func (t HlsFile) Bytes() []byte {
	buffer := new(bytes.Buffer)
	io.Copy(buffer,t.File)
	return buffer.Bytes()
}

func GetHlsFilePath(videoId int, segName string) (string,error) {

	if strings.Contains(segName, "..") || strings.Contains(segName, "/") {
		return "", errors.New("invalid segment name")
	}
	basePath := fmt.Sprintf("%s/%d", config.Config.AssetsDirPath, videoId)
	absBase, err := filepath.Abs(basePath)
	if err != nil {
		return "", err
	}
	joined := filepath.Join(absBase, "hls", segName)
	absJoined, err := filepath.Abs(joined)
	if err != nil {
		return "", err
	}
	if !strings.HasPrefix(absJoined, absBase+string(filepath.Separator)) {
		return "", errors.New("invalid segment name")
	}
	return absJoined, nil


}


func GetHlsFile(videoId int, segName string) (*HlsFile, error) {
	m3u8FilePath, err := GetHlsFilePath(videoId, segName)
	if err != nil {
		return nil, err
	}
	file, err := os.Open(m3u8FilePath)
	if err != nil {
		return nil, err
	}
	hlsFile := NewHlsFile(file)
	return &hlsFile, nil
}