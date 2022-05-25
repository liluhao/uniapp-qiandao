package user

import (
	"github.com/gin-gonic/gin"
	"github.com/llh/uniapp-qiandao/pkg/app"
	"github.com/llh/uniapp-qiandao/service"
	"github.com/llh/uniapp-qiandao/viewmodel"
)

func ForgetPassword(ctx *gin.Context) {
	var forgetPasswordRequest viewmodel.ForgetPasswordRequest
	if err := ctx.Bind(&forgetPasswordRequest); err != nil {
		app.SendResponse(ctx, app.ErrBind, nil)
		return
	}

	if err := service.ForgetPassword(forgetPasswordRequest); err != nil {
		app.SendResponse(ctx, err, nil)
		return
	}
	app.SendResponse(ctx, nil, nil)
}
