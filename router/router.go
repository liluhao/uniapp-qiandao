package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qiandao/controller/checkin"

	"qiandao/controller/class"
	"qiandao/controller/lesson"
	"qiandao/controller/sd"
	"qiandao/controller/user"
	"qiandao/router/middleware"
)

func Load(engine *gin.Engine, handlerFunc ...gin.HandlerFunc) *gin.Engine {
	engine.Use(gin.Recovery())
	engine.Use(handlerFunc...)

	// NoRoute()是默认情况下都返回404代码
	engine.NoRoute(func(context *gin.Context) {
		// 将给定的字符串写入响应正文
		context.String(http.StatusNotFound, "API路由错误")
	})

	authAPI := engine.Group("/api/auth")
	{
		authAPI.POST("/register", user.Register)
		authAPI.POST("/login", user.Login)
		authAPI.PUT("/update-forget-password", user.ForgetPassword)
	}

	userAPI := engine.Group("/api/user", middleware.Auth())
	{
		userAPI.PUT("/update-user", user.UpdateUserInfo)
		userAPI.PUT("/update-email", user.UpdateEmail)
		userAPI.PUT("/update-nick-name", user.UpdateNickName)
		userAPI.PUT("/update-password", user.UpdatePassword)
	}

	classAPI := engine.Group("/api/class")
	{
		classAPI.POST("", class.Create)
		classAPI.GET("", class.GetAllClass)
	}

	// 课程
	lessonApi := engine.Group("/api/lesson")
	{
		// 创建课程
		lessonApi.POST("", lesson.CreateLesson)
		// 获取创建的课程列表
		lessonApi.GET("/user", lesson.GetCreateLessonList)
		//获取加入的课程列表
		lessonApi.GET("/join", lesson.GetJoinLessonList)
		// 编辑课程
		lessonApi.PUT("/editor", lesson.EditorLesson)
		// 移除课程
		lessonApi.DELETE("/del", lesson.RemoveLesson)
	}

	// 检查http健康的路由组
	svcd := engine.Group("/api/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	// 签到
	checkInApi := engine.Group("/api/checkin")
	{
		// 创建签到
		checkInApi.POST("createCheckin", checkin.CreateCheckin)
		// 学生签到
		checkInApi.POST("", checkin.StuCheckIn)
		// 获取签到详情
		checkInApi.GET("getCheckinDetails", checkin.GetCheckinDetails)
		// 获取已创建签到列表
		checkInApi.GET("getCreatedCheckinList", checkin.GetCreatedCheckinLst)
		// 获取签到记录列表
		checkInApi.GET("getCheckinRecList", checkin.GetCheckinRecLst)
	}

	return engine
}
