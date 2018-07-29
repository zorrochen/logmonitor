package http_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"logmonitor/base/log"
	"logmonitor/base/pub"
	"time"
)

type regInfo struct {
	Method  string
	Uri     string
	Handler []gin.HandlerFunc
}

var _RegInfo []regInfo = make([]regInfo, 0, 50)

func Register(method, uri string, handler ...gin.HandlerFunc) {
	info := regInfo{method, uri, handler}
	_RegInfo = append(_RegInfo, info)
}

func Run(ip string, port string) {
	r := gin.New()

	// 中间件
	r.Use(ReqSimplePrint)

	// 将注册的handle加入路由
	for _, info := range _RegInfo {
		switch info.Method {
		case "GET":
			r.GET(info.Uri, info.Handler...)
		case "POST":
			r.POST(info.Uri, info.Handler...)
		}
	}

	// 未知路由，添加自定义处理
	r.NoRoute(NotFoundPrint)

	// 启动http服务
	r.Run(fmt.Sprintf("%s:%s", ip, port))
}

// 请求入口，添加统一打印
// TODO:添加用户id,版本号
func ReqSimplePrint(c *gin.Context) {
	// 打印请求
	log.LOG.Info("[ReqSimplePrint] %s, %s", c.Request.URL, pub.CopyHttpRequestBody(c.Request))

	// 计算耗时
	t := time.Now()
	c.Next()
	ts := time.Since(t)
	timecost := fmt.Sprintf("%d", ts/time.Millisecond)

	// 打印结果
	log.LOG.Info("[ReqSimplePrint-resp] %s, %s(ms)", c.Request.URL, timecost)
}

// 404记录
func NotFoundPrint(c *gin.Context) {
	log.LOG.Warn("[NotFoundPrint] %s, %s", c.Request.URL, pub.CopyHttpRequestBody(c.Request))
}
