package class

import (
	"github.com/gin-gonic/gin"
	"github.com/llh/uniapp-qiandao/pkg/app"
	"github.com/llh/uniapp-qiandao/pkg/util"
	"github.com/llh/uniapp-qiandao/service"
)

// GetAllClass 根据limit与offset获取班级列表
func GetAllClass(ctx *gin.Context) {
	var pageCondition util.PageRequest
	if err := ctx.Bind(&pageCondition); err != nil {
		app.SendResponse(ctx, app.ErrBind, nil)
		return
	}
	//获取班级列表
	class, err := service.GetAllClass(pageCondition)
	if err != nil {
		app.SendResponse(ctx, err, nil)
		return
	}
	app.SendResponse(ctx, nil, class)
}
