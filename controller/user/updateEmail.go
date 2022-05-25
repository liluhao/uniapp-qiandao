package user

import (
	"github.com/gin-gonic/gin"
	"github.com/llh/uniapp-qiandao/pkg/app"
	"github.com/llh/uniapp-qiandao/service"
	"github.com/llh/uniapp-qiandao/viewmodel"
)

// UpdateEmail 修改用户邮箱 controller
func UpdateEmail(ctx *gin.Context) {
	var updateEmail viewmodel.UpdateEmailRequest
	if err := ctx.Bind(&updateEmail); err != nil {
		app.SendResponse(ctx, app.ErrBind, nil)
		return
	}
	if err := service.UpdateEmail(updateEmail); err != nil {
		app.SendResponse(ctx, err, nil)
		return
	}
	app.SendResponse(ctx, nil, nil)
}
