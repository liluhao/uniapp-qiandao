package model

import (
	"time"
)

// Checkin 签到结构体
type Checkin struct {
	CheckinID   string    `json:"checkin_id" gorm:"primaryKey"`
	CreatorID   string    `json:"creator"`
	LessonID    string    `json:"lesson_id"`
	BeginTime   time.Time `json:"begin_time"`   //开始时间
	EndTime     time.Time `json:"end_time"`     //结束时间
	CheckinCode string    `json:"checkin_code"` //签到码
	Longitude   string    `json:"longitude"`    //经度
	Latitude    string    `json:"latitude"`     //纬度

	CreatedAt time.Time  `gorm:"created_at"`
	UpdatedAt time.Time  `gorm:"updated_at"`
	DeletedAt *time.Time `gorm:"deleted_at"`
}
