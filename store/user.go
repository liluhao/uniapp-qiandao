package store

import (
	"github.com/lexkong/log"
	"qiandao/model"
	"qiandao/pkg/app"
	"qiandao/viewmodel"
)

// IsExistUser 判断用户表中某个值是不是存在了; field:数据库中要判断的字段 , param : 与field对应的值
func IsExistUser(field, param string, a *model.User) bool {
	isPhone := DB.Self.Where(field+" = ?", param).First(&a)
	if isPhone.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// CreateUserMapper 创建用户
func CreateUserMapper(user *model.User, classId string) error {
	//用户事务
	tx := DB.Self.Begin()
	// 创建用户
	db := tx.Create(&user)
	if db.Error != nil {
		tx.Rollback()
		log.Errorf(db.Error, "创建班级失败")
		return app.InternalServerError
	}
	log.Infof("创建班级：成功创建:%v条记录", db.RowsAffected)
	// 维护用户班级中间表
	intermediateTable := tx.Create(&model.Connection{
		ClassId: classId,
		UserId:  user.UserId,
	})
	if intermediateTable.Error != nil {
		//失败则回滚事务
		tx.Rollback()
		log.Errorf(intermediateTable.Error, "维护中间表失败")
		return app.InternalServerError
	}
	log.Infof("维护班级用户中间表：成功创建:%v条记录", intermediateTable.RowsAffected)
	tx.Commit()
	return nil
}

// UpdateUserMapper 修改用户信息
func UpdateUserMapper(updateUser viewmodel.UpdateUserInfoRequest) error {
	result := DB.Self.Model(model.User{}).Where("user_id = ?", updateUser.UserId).Updates(model.User{
		Email:    updateUser.Email,
		RealName: updateUser.RealName,
		Hobby:    updateUser.Hobby,
		Address:  updateUser.Address,
		Sex:      updateUser.Sex,
		Age:      updateUser.Age,
	})
	if result.Error != nil {
		log.Errorf(result.Error, "修改用户信息失败")
		return app.InternalServerError
	}
	log.Infof("修改用户信息：成功修改 %v 条记录", result.RowsAffected)
	return nil
}

// UpdateEmailMapper 修改邮箱
func UpdateEmailMapper(updateEmail viewmodel.UpdateEmailRequest) error {
	result := DB.Self.Model(&model.User{}).Where("user_id = ?", updateEmail.UserId).Update("email", updateEmail.Email)
	if result.Error != nil {
		log.Errorf(result.Error, "修改邮箱失败")
		return app.InternalServerError
	}
	log.Infof("修改邮箱：成功修改 %v 条记录", result.RowsAffected)
	return nil
}

// UpdateNickNameMapper 修改昵称
func UpdateNickNameMapper(updateNickName viewmodel.UpdateNickNameRequest) error {
	result1 := DB.Self.Model(&model.User{}).Where("user_id = ?", updateNickName.UserId).Update("nick_name", updateNickName.NickName)
	if result1.Error != nil {
		log.Errorf(result1.Error, "修改昵称失败")
		return app.InternalServerError
	}
	log.Infof("修改昵称：成功修改 %v 条记录", result1.RowsAffected)
	return nil
}

// GetPasswordById 根据用户ID查找对应用户的密码
func GetPasswordById(userID string) (string, error) {
	user := &model.User{}
	result := DB.Self.Select("password").Where("user_id = ?", userID).Find(&user)
	if result.Error != nil {
		log.Errorf(result.Error, "查找用户密码失败")
		return "", app.InternalServerError
	}
	return user.Password, nil
}

// UpdatePasswordByFieldMapper 根据规定的字段修改密码 ; field:数据库中要判断的字段 , param : 与field对应的值
func UpdatePasswordByFieldMapper(filed, condition, password string) error {
	//更新密码
	result := DB.Self.Model(&model.User{}).Where(filed+" = ?", condition).Update("password", password)
	if result.Error != nil {
		log.Errorf(result.Error, "修改密码失败")
		return app.InternalServerError
	}
	return nil
}

// GetEmailByPhone 根据手机号查找对应用户的邮箱
func GetEmailByPhone(phone string) (string, error) {
	user := &model.User{}
	//Select只查询email字段
	result := DB.Self.Select("email").Where("phone = ?", phone).Find(&user)
	if result.Error != nil {
		log.Errorf(result.Error, "查找用户邮箱失败")
		return "", app.InternalServerError
	}
	return user.Email, nil
}

// GetUserInfoByPhone 根据手机号获取用户信息
func GetUserInfoByPhone(phone string) (*viewmodel.UserInfo, error) {

	// userInfo := DB.Self.Where("phone = ?", phone).First(&info)

	result := &viewmodel.UserInfo{}

	//查询model.User结构体，再内嵌到viewmodel.UserInfo结构体里
	//利用左外连接查询
	scan := DB.Self.Model(&model.User{}).
		Select("user.user_id, "+
			"user.phone, "+
			"user.password, "+
			"user.role, "+
			"user.email, "+
			"user.real_name, "+
			"user.nick_name, "+
			"user.hobby, "+
			"user.address, "+
			"user.sex, "+
			"user.age, "+
			"user.class_id, "+
			"class.class_name").
		Joins("left join class on user.class_id = class.class_id").
		Where("user.phone = ?", phone).Scan(&result)

	//*DB.Scan函数的返回值还只有*DB,不会返回errro，所以通过RowsAffected是否为0来判断是否成功
	if scan.RowsAffected == 0 {
		log.Errorf(scan.Error, "找不到账号为: %v 的信息", phone)
		return &viewmodel.UserInfo{}, app.ErrAccountDoesNotExist
	}

	log.Infof("查找到账号为: %v 的信息", phone)
	return result, nil
}
