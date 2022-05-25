package viewmodel

// 用户模块的 request 和 response
// 用户模块的 request 和 response
// 用户模块的 request 和 response

// RegisterRequest 用户注册接口 接收参数
type RegisterRequest struct {
	Phone    string `json:"phone" form:"phone" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	Role     int    `json:"role" form:"role"`
	ClassId  string `json:"class_id" form:"class_id"`
}

// LoginRequest 登录请求结构体
type LoginRequest struct {
	Phone    string `json:"phone" form:"phone" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

//会model下的user表几乎一模一样
type UserInfo struct {
	UserId    string `json:"user_id"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Role      int    `json:"role"`
	ClassId   string `json:"class_id"`
	ClassName string `json:"class_name"`
	RealName  string `json:"real_name"`
	NickName  string `json:"nick_name"`
	Hobby     string `json:"hobby"`
	Address   string `json:"address"`
	Sex       int    `json:"sex"`
	Age       int    `json:"age"`
}

// LoginResponse 登录响应结构体
type LoginResponse struct {
	Token string    `json:"token"`
	User  *UserInfo `json:"user"`
}

// UpdateUserInfoRequest 修改用户信息请求结构体
type UpdateUserInfoRequest struct {
	UserId   string `json:"user_id" form:"user_id"`
	Email    string `json:"email" form:"email"`
	RealName string `json:"real_name" form:"real_name"`
	Hobby    string `json:"hobby" form:"hobby"`
	Address  string `json:"address" form:"address"`
	Sex      int    `json:"sex" form:"sex"`
	Age      int    `json:"age" form:"age"`
}

// UpdateEmailRequest 修改邮箱请求结构体
type UpdateEmailRequest struct {
	UserId string `json:"user_id" form:"user_id" binding:"required"`
	Email  string `json:"email" form:"email" binding:"required"`
}

// UpdateNickNameRequest 修改昵称请求结构体
type UpdateNickNameRequest struct {
	UserId   string `json:"user_id" form:"user_id" binding:"required"`
	NickName string `json:"nick_name" form:"nick_name" binding:"required"`
}

// UpdatePasswordRequest 修改密码请求结构体
type UpdatePasswordRequest struct {
	UserId             string `json:"user_id" form:"user_id" binding:"required"`
	OldPassword        string `json:"old_password" form:"old_password" binding:"required"`
	NewPassword        string `json:"new_password" form:"new_password" binding:"required"`
	NewConfirmPassword string `json:"new_confirm_password" form:"new_confirm_password" binding:"required"`
}

// ForgetPasswordRequest 忘记密码请求结构体
type ForgetPasswordRequest struct {
	Phone string `json:"phone" form:"phone" binding:"required"`
	Email string `json:"email" form:"email" binding:"required"`
}
