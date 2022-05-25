package user

import (
	"github.com/gin-gonic/gin"
	"github.com/llh/uniapp-qiandao/pkg/app"
	"github.com/llh/uniapp-qiandao/service"
	"github.com/llh/uniapp-qiandao/viewmodel"
)

// Register 用户注册 controller
func Register(ctx *gin.Context) {
	var registerRequest viewmodel.RegisterRequest
	if err := ctx.Bind(&registerRequest); err != nil { //返回错误
		app.SendResponse(ctx, app.ErrBind, nil)
		return
	}
	if err := service.CreateUser(&registerRequest); err != nil {
		app.SendResponse(ctx, err, nil) //返回错误
		return
	}
	app.SendResponse(ctx, nil, nil) //返回成果
}
