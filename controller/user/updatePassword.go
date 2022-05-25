package user

import (
	"github.com/gin-gonic/gin"
	"github.com/llh/uniapp-qiandao/pkg/app"
	"github.com/llh/uniapp-qiandao/service"
	"github.com/llh/uniapp-qiandao/viewmodel"
)

func UpdatePassword(ctx *gin.Context) {
	var updatePasswordRequest viewmodel.UpdatePasswordRequest
	if err := ctx.Bind(&updatePasswordRequest); err != nil {
		app.SendResponse(ctx, app.ErrBind, nil)
		return
	}
	if err := service.UpdatePassword(updatePasswordRequest); err != nil {
		app.SendResponse(ctx, err, nil)
		return
	}
	app.SendResponse(ctx, nil, nil)
}
