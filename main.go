package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/llh/uniapp-qiandao/config"
	"github.com/llh/uniapp-qiandao/router"
	"github.com/llh/uniapp-qiandao/store"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var (
	//类似于标准库中plag.String方法，在此-c参数设设置为0,返回*string类型，设为全局变量
	cfg = pflag.StringP("config", "c", "", "api config file path.")
)

func main() {
	pflag.Parse()
	// 初始化配置文件并监控配置文件变化进行热加载程序
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// 数据库初始化并建立连接
	store.DB.Init()
	defer store.DB.Close()

	gin.SetMode(viper.GetString("run_mode"))

	g := gin.New()
	router.Load(
		g,
		// 加入多个中间件...
	)

	// ping 服务器以确保路由正常工作(健康检查)
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("路由没有响应，或者启动时间过长.", err)
		}
		log.Info("路由启动成功.")
		log.Infof("开始监听 http 地址上的传入请求: %s", viper.GetString("addr"))
	}()
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// 通过向 /health 发送 GET 请求来 Ping 服务器。
		resp, err := http.Get(viper.GetString("url") + "/api/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		// 休眠一秒钟以继续下一次 ping。
		log.Info("等待路由，1秒后重试.")
		time.Sleep(time.Second)
	}
	return errors.New("无法连接到路由")
}
