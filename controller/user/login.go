package user

import (
	"github.com/gin-gonic/gin"
	"github.com/llh/uniapp-qiandao/pkg/app"
	"github.com/llh/uniapp-qiandao/service"
	"github.com/llh/uniapp-qiandao/viewmodel"
)

func Login(ctx *gin.Context) {
	var loginRequest viewmodel.LoginRequest
	if err := ctx.Bind(&loginRequest); err != nil {
		app.SendResponse(ctx, app.ErrBind, nil)
		return
	}
	login, err := service.Login(loginRequest)
	if err != nil {
		app.SendResponse(ctx, err, nil)
		return
	}
	app.SendResponse(ctx, nil, login)
}
