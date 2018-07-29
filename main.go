// 程序入口模块

package main

import (
	"flag"
	"fmt"
	"logmonitor/base/http_server"
	"logmonitor/base/log"
	"logmonitor/config"
	"logmonitor/handler"
)

var (
	ConfPath    = flag.String("confpath", "./config/config.yml", "Config File Path")
	LogCfgpath  = flag.String("logCfgpath", "./config/seelog.xml", "seelog.xml path")
	showVersion = flag.Bool("v", false, "print version string")
)

func main() {
	// 启动参数解析
	flag.Parse()

	// 版本查询
	if *showVersion {
		fmt.Println(version())
		return
	}
	// 配置初始化
	config.Init(*ConfPath)

	// 日志初始化
	log.LogInit("logmonitor", *LogCfgpath, false, "")

	// 业务模块初始化
	handler.Init()

	// 服务启动
	log.LOG.I("server start!!!")
	http_server.Run(config.Cfg.ServerIP, config.Cfg.ServerPort)
	log.LOG.I("server stoped!!!")
}

//====================version(from git) print====================
// 以下相关版本信息由go build编译时注入
// go build -ldflags "-X main.serverName="" -X main.gitCommit=`git rev-parse --short=7 HEAD` -X main.buildTime=`date +%Y-%m-%d_%H:%M:%S`"
var serverName string
var gitCommit string
var buildTime string

func version() string {
	return fmt.Sprintf("%s %s@%s", serverName, gitCommit, buildTime)
}
