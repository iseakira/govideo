package models

import "time"

type Views struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	UserID    string    `json:"userId,omitempty"`
	VideoID   int       `json:"videoId,omitempty"`
	Video     Video     `gorm:"constraint:CASCADE,OnDelete"`
	CreatedAt time.Time `json:"createdAt,omitempty" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" gorm:"autoUpdateTime"`
}

func CreateViews(userID string,videoID int) (*Views,error) {
	views := &Views{
		VideoID: videoID,
		UserID:userID,
	}
	if err := DbConnection.Create(views).Error; err != nil {
		return nil,err
	}
	return views,nil
}

func TotalViews(videoID int) (total int64,err error) {
	if err = DbConnection.Model(&Views{}).Where("video_id=?",videoID).Count(&total).Error; err != nil {
		return total,err
	}
	return total,err
}

