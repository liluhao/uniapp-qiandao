package model

import (
	"github.com/go-playground/validator/v10"
	"time"
)

// User 用户表
type User struct {
	UserId   string `json:"user_id"   gorm:"primary_key;column:user_id;not null"`
	Phone    string `json:"phone"     gorm:"column:phone;not null"`
	Password string `json:"password"  gorm:"column:password;not null"`
	Email    string `json:"email"     gorm:"column:email"`
	Role     int    `json:"role"      gorm:"column:role"`
	ClassId  string `json:"class_id"  gorm:"column:class_id;"`
	RealName string `json:"real_name" gorm:"column:real_name"` //真实名字
	NickName string `json:"nick_name" gorm:"column:nick_name"` //外号
	Hobby    string `json:"hobby"     gorm:"column:hobby"`
	Address  string `json:"address"   gorm:"column:address"`
	Sex      int    `json:"sex"       gorm:"column:sex"`
	Age      int    `json:"age"       gorm:"column:age"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// Validate 对请求参数的校验
func (u *User) Validate() error {
	// 创建一个验证器，这个验证器可以指定选项、添加自定义约束，然后在调用他的Struct()方法来验证各种结构对象的字段是否符合定义的约束
	v := validator.New()
	return v.Struct(u)
}
