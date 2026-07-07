package models

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"ms-stream-api/config"
)

type M3u8File struct {
	File io.ReadCloser
}

func NewM3u8File(file *os.File) M3u8File {
	return M3u8File{
		File:file,
	}
}

func (t M3u8File) Bytes() []byte {
	buffer := new(bytes.Buffer)
	io.Copy(buffer,t.File)
	return buffer.Bytes()
}

func GetM3u8FilePath(videoId int) string {
	basePath := fmt.Sprintf("%s%d",config.Config.AssetsDirPath,videoId)
		return fmt.Sprintf("%s/hls/%s", basePath, config.Config.AssetsM3u8FileName)

}

func GetM3u8File(videoId int) (*M3u8File,error) {
		m3u8FilePath := GetM3u8FilePath(videoId)
		file, err := os.Open(m3u8FilePath)

		if err != nil {
			return nil,err
		}
		m3u8File := NewM3u8File(file)
		return &m3u8File,nil

}