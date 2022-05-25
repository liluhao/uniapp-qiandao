package store

import (
	"github.com/jinzhu/gorm"
	"github.com/llh/uniapp-qiandao/model"
)

type Tx struct {
	tx *gorm.DB
}

func GetTx() *Tx {
	return &Tx{tx: DB.Self}
}

func (tx *Tx) Begin() {
	tx.tx = tx.tx.Begin()
}

func (tx *Tx) RollBack() {
	tx.tx.Rollback()
}

func (tx *Tx) Commit() (err error) {
	err = tx.tx.Commit().Error
	return
}

// CreateCheckin  创建一个签到
func (tx *Tx) CreateCheckin(checkin *model.Checkin) (err error) {
	err = tx.tx.Create(checkin).Error
	return
}

// GetCheckinById 根据签到CheckinID从checkin获取一个签到
func (tx *Tx) GetCheckinById(checkinID string) (checkin model.Checkin, err error) {
	err = tx.tx.Where("checkin_id = ?", checkinID).First(&checkin).Error
	return
}

// GetCheckinByCreator  根据用户id获取一个签到
func (tx *Tx) GetCheckinByCreator(creatorID string) (checkinList []model.Checkin, err error) {
	err = tx.tx.Where("creator_id = ?", creatorID).Find(&checkinList).Error
	return
}

//func (tx *Tx) GetACheckin(field, fieldValue string) (checkinLst []model.Checkin, err error) {
//	err = tx.tx.Where(fmt.Sprintf("%v = ?", field), fieldValue).Find(&checkinLst).Error
//	return
//}

// GetLessonByID 根据课程id获取课程
func (tx *Tx) GetLessonByID(lessonID string) (lesson model.Lesson, err error) {
	err = tx.tx.Where("lesson_id = ?", lessonID).First(&lesson).Error
	return
}

// GetClassLstByLessonID 根据课程id获取需要签到的班级列表
func (tx *Tx) GetClassLstByLessonID(lessonID string) (classList []model.Class, err error) {
	err = tx.tx.Raw("select * from class where class_id IN (select class_id from class_lesson where lesson_id = ? and class_lesson.deleted_at is null) and class.deleted_at is null", lessonID).Scan(&classList).Error
	return
}

// GetStuLstByClassID 根据班级id获取需要签到的学生列表
func (tx *Tx) GetStuLstByClassID(classID string) (stuList []model.User, err error) {
	err = tx.tx.Raw("select * from user where class_id = ? and user.deleted_at is null", classID).Scan(&stuList).Error
	return
}

// AddCheckinRec 添加学生签到记录
func (tx *Tx) AddCheckinRec(stuCheckin *model.CheckinRec) (err error) {
	err = tx.tx.Create(stuCheckin).Error
	return
}

// UpdateCheckinRecStateByID  根据签到记录id更新学生签到状态
func (tx *Tx) UpdateCheckinRecStateByID(checkinRecID string, checkinRecState int) (err error) {
	err = tx.tx.Model(&model.CheckinRec{}).Where("checkin_rec_id = ?", checkinRecID).Update("state", checkinRecState).Error
	return
}

// GetCheckinRecByID  根据checkinRecID(CheckinID+viewStuCheckin.UserID)从checkinRec获取一个签到
func (tx *Tx) GetCheckinRecByID(checkinRecID string) (checkinRec model.CheckinRec, err error) {
	err = tx.tx.Where("checkin_rec_id = ?", checkinRecID).First(&checkinRec).Error
	return
}

// GetCheckinRecLstByCheckinID 根据签到id获取签到记录列表
func (tx *Tx) GetCheckinRecLstByCheckinID(checkinID string) (checkedInList []model.CheckinRec, err error) {
	err = tx.tx.Where("checkin_id = ?", checkinID).Find(&checkedInList).Error
	return
}

// GetCheckinRecByUserID  根据用户id获取某个用户需要签到的列表
func (tx *Tx) GetCheckinRecByUserID(userID string) (checkedInList []model.CheckinRec, err error) {
	err = tx.tx.Where("user_id = ?", userID).Order("end_time desc").Find(&checkedInList).Error
	return
}

func (tx *Tx) GetClassByUserID(userID string) (class model.Class, err error) {
	err = tx.tx.Raw("select * from class where class_id = (select class_id from connection where user_id = ? and connection.deleted_at is null) and class.deleted_at is null", userID).Scan(&class).Error
	return
}
