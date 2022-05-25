package viewmodel

// CreateCheckin 创建签到 请求结构体
type CreateCheckin struct {
	LessonID    string `json:"lesson_id" form:"lesson_id"`
	Duration    int    `json:"duration" form:"duration"`
	CreatorID   string `json:"creator_id" form:"creator_id"`
	CheckinCode string `json:"checkin_code" form:"checkin_code"`
	Longitude   string `json:"longitude" form:"longitude"`
	Latitude    string `json:"latitude" form:"latitude"`
}

// Checkin 学生签到 请求结构体
type Checkin struct {
	CheckinID   string `json:"checkin_id" form:"checkin_id"`
	UserID      string `json:"user_id" form:"user_id"` //用户id
	CheckinCode string `json:"checkin_code" form:"checkin_code"`
	Longitude   string `json:"longitude" form:"longitude"`
	Latitude    string `json:"latitude" form:"latitude"`
}

type List struct {
	ClassName string `json:"class_name"`
	UserName  string `json:"user_name"`
	State     string `json:"state"`
}

// CheckinDetailsResponse 签到详情 响应结构体
type CheckinDetailsResponse struct {
	TotalList     []List `json:"total_list"`
	CheckedInList []List `json:"checkin_list"`
	NotCheckList  []List `json:"not_check_list"`
}

// ListResponse 已创建签到列表/签到记录列表 响应结构体
type ListResponse struct {
	CheckinID    string `json:"checkin_id"`
	LessonName   string `json:"lesson_name"`
	BeginTime    string `json:"begin_time"`
	State        int    `json:"state"`
	CheckinState int    `json:"checkin_state"`
}
