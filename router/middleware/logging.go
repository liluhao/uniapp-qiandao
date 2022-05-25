package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/willf/pad"
	"io/ioutil"
	"qiandao/pkg/app"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logging() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 记录访问接口起始时间
		start := time.Now().UTC()

		// 记录访问的接口的path
		path := context.Request.URL.Path

		// 跳过健康检查请求
		if path == "/sd/health" || path == "/sd/ram" || path == "/sd/cpu" || path == "/sd/disk" {
			return
		}

		// 读取body内容
		var bodyBytes []byte
		if context.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(context.Request.Body)
		}

		// 由于已经读取过 Request 的 Body 数据了，后续再读会读不到(读取过后会被置空)
		// 所以这里需要自己重新再构建一个 ReadCloser 赋值给原先的 Body
		// 用 NopCloser 简单的包装一下
		context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		method := context.Request.Method
		// ClientIP 实现了一种尽力而为的算法来返回真实的客户端 IP。
		// 它在后台调用了 c.RemoteIP()，以检查远程 IP 是否是受信任的代理。
		// 如果是，它将尝试解析 Engine.RemoteIPHeaders 中定义的标头（默认为 [X-Forwarded-For, X-Real-Ip]）。
		// 如果标头在语法上无效或远程 IP 不对应于受信任的代理，则返回远程 IP（来自 Request.RemoteAddr)
		ip := context.ClientIP()

		blw := &bodyLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: context.Writer,
		}
		context.Writer = blw

		// 调用该请求的剩余处理程序
		context.Next()

		// 定义结束时间.
		end := time.Now().UTC()
		// 计算出调用该接口耗时
		latency := end.Sub(start)

		code, message := -1, ""
		// 获取 code 和 message
		var response app.Response

		// Unmarshal函数解析json编码的数据并将结果存入v指向的值
		// 要将json数据解码写入一个结构体，函数会匹配输入对象的键和Marshal使用的键
		if err := json.Unmarshal(blw.body.Bytes(), &response); err != nil {
			log.Errorf(err, "response body 不能解码到 model。响应结构体正文: `%s`", blw.body.Bytes())
			code = app.InternalServerError.Code
			message = err.Error()
		} else {
			code = response.Code
			message = response.Message
		}
		log.Infof("%-13s | %-12s | %s %s | {code: %d, message: %s}", latency, ip, pad.Right(method, 5, ""), path, code, message)
	}
}
