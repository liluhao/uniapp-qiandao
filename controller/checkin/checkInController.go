package checkin

import (
	"github.com/gin-gonic/gin"
	"github.com/llh/uniapp-qiandao/pkg/app"
	"github.com/llh/uniapp-qiandao/service"
	"github.com/llh/uniapp-qiandao/viewmodel"
)

//创建签到
func CreateCheckin(ctx *gin.Context) {
	viewCheckin := new(viewmodel.CreateCheckin)
	err := ctx.ShouldBindJSON(viewCheckin)
	if err != nil {
		app.SendResponse(ctx, app.ErrBind, nil)
		return
	}
	if viewCheckin.CreatorID == "" || viewCheckin.CheckinCode == "" || viewCheckin.Duration <= 0 ||
		viewCheckin.LessonID == "" || viewCheckin.Longitude == "" || viewCheckin.Latitude == "" {
		app.SendResponse(ctx, app.ErrParamNull, nil)
		return
	}
	createCheckin, err := service.CreateCheckin(viewCheckin)
	if err != nil {
		app.SendResponse(ctx, err, nil)
		return
	}
	app.SendResponse(ctx, app.OK, createCheckin.CheckinID)
}

//学生签到
func StuCheckIn(ctx *gin.Context) {
	checkIn := new(viewmodel.Checkin)
	err := ctx.ShouldBindJSON(checkIn)
	if err != nil {
		app.SendResponse(ctx, app.ErrBind, nil)
		return
	}
	if checkIn.CheckinID == "" || checkIn.UserID == "" || checkIn.CheckinCode == "" {
		app.SendResponse(ctx, app.ErrParamNull, nil)
		return
	}
	checkinResponse, err := service.StuCheckin(checkIn)
	if err != nil {
		app.SendResponse(ctx, err, checkinResponse)
		return
	}
	app.SendResponse(ctx, app.OK, checkinResponse)
}

// 获取签到详情
func GetCheckinDetails(ctx *gin.Context) {
	checkinID := ctx.Query("checkin_id")
	if checkinID == "" {
		app.SendResponse(ctx, app.ErrParamNull, nil)
		return
	}
	checkinDetailsResponse, err := service.GetCheckinDetails(checkinID)
	if err != nil {
		app.SendResponse(ctx, err, nil)
		return
	}
	app.SendResponse(ctx, app.OK, checkinDetailsResponse)
}

//获取已创建签到列表
func GetCreatedCheckinLst(ctx *gin.Context) {
	userID := ctx.Query("user_id")
	if userID == "" {
		app.SendResponse(ctx, app.ErrParamNull, nil)
		return
	}
	response, err := service.GetCreatedCheckInList(userID)
	if err != nil {
		app.SendResponse(ctx, err, nil)
		return
	}
	app.SendResponse(ctx, app.OK, response)
}

//获取签到记录列表
func GetCheckinRecLst(ctx *gin.Context) {
	userID := ctx.Query("user_id")
	if userID == "" {
		app.SendResponse(ctx, app.ErrParamNull, nil)
		return
	}
	response, err := service.GetCheckinRecList(userID)
	if err != nil {
		app.SendResponse(ctx, err, nil)
		return
	}
	app.SendResponse(ctx, app.OK, response)
}
