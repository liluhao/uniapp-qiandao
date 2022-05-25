package model

import "time"

type Lesson struct {
	LessonID 		string  `json:"lesson_id" gorm:"primaryKey"`
	LessonName 	    string  `json:"lesson_name"`
	LessonCreator 	string  `json:"lesson_creator"`

	CreatedAt   time.Time    `gorm:"created_at"`
	UpdatedAt   time.Time    `gorm:"updated_at"`
	DeletedAt   *time.Time   `gorm:"deleted_at" sql:"index"`
}
