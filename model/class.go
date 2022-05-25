package model

import "time"

type Class struct {
	ClassId       string `json:"class_id"       gorm:"primary_key;column:class_id;not null"`
	ClassName     string `json:"class_name"     gorm:"column:class_name"`
	ClassCapacity int    `json:"class_capacity" gorm:"column:class_capacity"`
	CreateId      string `json:"create_id"      gorm:"column:create_id"`

	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"-"` //json:"-": 字段被本包忽略
	DeletedAt *time.Time `gorm:"column:deleted_at" sql:"index" json:"-"`
}
