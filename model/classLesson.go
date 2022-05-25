package model

import (
	"time"
)

type ClassLesson struct {
	ClassLessonID  string    `json:"class_lesson_id"`
	ClassID        string    `json:"class_id"`   // 班级id
	LessonID       string    `json:"lesson_id"` //  课程id
	CreatedAt    time.Time   `gorm:"created_at"`
	UpdatedAt    time.Time   `gorm:"updated_at"`
	DeletedAt   *time.Time   `gorm:"deleted_at" sql:"index"`


}
