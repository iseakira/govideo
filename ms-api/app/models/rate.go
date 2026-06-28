package models

import "time"

type Rate struct {
	ID        int       `json:"id" gorm:"unique;autoIncrement;not null"`
	UserID    string    `json:"userId,omitempty" gorm:"primaryKey"`
	VideoID   int       `json:"videoId,omitempty" gorm:"primaryKey"`
	Video     Video     `json:"-" gorm:"constraint:CASCADE,OnDelete"`
	Value     float32   `json:"value,omitempty" binding:"numeric"`
	CreatedAt time.Time `json:"createdAt,omitempty" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" gorm:"autoUpdateTime"`
}

func NewRate(UserID string, videoID int, value float32) *Rate {
	return &Rate{
		UserID :UserID,
		VideoID: videoID,
		Value: value,
	}
}

func GetRate(UserID string, videoID int) (rate *Rate,err error) {
	if err = DbConnection.Where("user_id=? AND video_id=?",UserID,videoID).Find(&rate).Error; err != nil {
		return nil,err
	}
	return rate,nil
}

func CreateOrUpdateRate(userID string,videoID int,value float32) (*Rate,error) {
	r := NewRate(userID  ,videoID,value)
	result := DbConnection.Where("user_id=? AND video_id=?",userID,videoID).Updates(r)

	if result.Error != nil {
		return r,result.Error
	}
	if result.RowsAffected == 0 {
		if err := DbConnection.Omit("id").Create(r).Error; err != nil {
			return r,err
		}
	}
		return r,nil
}


func RateAverage(videoID int) (avg float32,err error) {
	if err = DbConnection.Model(&Rate{}).
	Where("video_id=?",videoID).
	Select("COALECSCE(avg(value),0)").
	Scan(&avg).Error; err != nil {
		return avg,err
	}
	return avg,nil
}