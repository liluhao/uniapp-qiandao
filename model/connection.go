package model

import "time"

type Connection struct {
	Id      int    `json:"id"       gorm:"primary_key;column:id;not null"`
	ClassId string `json:"class_id" gorm:"column:class_id"`
	UserId  string `json:"user_id"  gorm:"column:user_id"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
