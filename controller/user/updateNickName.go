package user

import (
	"github.com/gin-gonic/gin"
	"github.com/llh/uniapp-qiandao/pkg/app"
	"github.com/llh/uniapp-qiandao/service"
	"github.com/llh/uniapp-qiandao/viewmodel"
)

// UpdateNickName 修改用户昵称 controller
func UpdateNickName(ctx *gin.Context) {
	var updateNickName viewmodel.UpdateNickNameRequest
	if err := ctx.Bind(&updateNickName); err != nil {
		app.SendResponse(ctx, app.ErrBind, nil)
		return
	}
	if err := service.UpdateNickName(updateNickName); err != nil {
		app.SendResponse(ctx, err, nil)
		return
	}
	app.SendResponse(ctx, nil, nil)
}
