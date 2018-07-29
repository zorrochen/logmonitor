package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"logmonitor/base/log"
	"logmonitor/base/pub"
	"logmonitor/config"
	"logmonitor/proxy/wxcorp"
	"net/http"
	"time"
)

type LogInfo_t struct {
	Content  string `json:"content" binding:"required"`
	FileLine string `json:"file_line" binding:"required"`
	SrvName  string `json:"srv_name" binding:"required"`
}

func LogMonitorHandler(c *gin.Context) {
	loginfo := LogInfo_t{}
	err := c.Bind(&loginfo)
	if err != nil {
		log.LOG.E("%v\n", err)
		return
	}

	srv := getSrvStat(loginfo.SrvName)
	srv.currentLogCount++

	c.JSON(http.StatusOK, "OK")
}

var globalStat map[string]*SrvLogStat

type SrvLogStat struct {
	srvName              string
	currentLogCount      int
	lastTriggerTimeStamp int64
	logStatCache         *pub.LRUCache
}

func getSrvStat(srvname string) *SrvLogStat {
	_, ok := globalStat[srvname]
	if !ok {
		newLogStat := &SrvLogStat{
			srvName:         srvname,
			currentLogCount: 0,
			logStatCache:    pub.NewLRUCache(60),
		}
		globalStat[srvname] = newLogStat
	}
	return globalStat[srvname]
}

func MonitorTask() {
	globalStat = map[string]*SrvLogStat{}
	for {
		for _, v := range globalStat {
			log.LOG.I("[MonitorTask] [%s] begin...", v.srvName)
			currentTimeStr := time.Now().Local().Format("15:04")
			v.logStatCache.Set(currentTimeStr, v.currentLogCount)
			log.LOG.I("[MonitorTask] [%s] set(%s, %d)", v.srvName, currentTimeStr, v.currentLogCount)
			v.currentLogCount = 0
			log.LOG.I("[MonitorTask] [%s] complete...", v.srvName)
		}
		<-time.After(time.Minute)
	}
}

func MonitorTriggerTask() {
	for {
		for _, v := range globalStat {
			totolErrCnt := 0
			//log.LOG.I("[MonitorTriggerTask] [%s] begin...", v.srvName)
			logStatList, ok, err := v.logStatCache.GetAllWithoutUpdate()
			if err != nil || !ok {
				goto LOOP_END
			}

			//统计指定时间内错误总数
			totolErrCnt = 0
			for k, v := range logStatList {
				if k > config.Cfg.MonitorTime {
					break
				}
				totolErrCnt += v.(int)
			}

			//触发
			if totolErrCnt >= config.Cfg.MonitorThreshold {
				//CD
				if time.Now().Local().Unix()-v.lastTriggerTimeStamp < int64(config.Cfg.MonitorCD*60) {
					goto LOOP_END
				}

				//发送告警
				log.LOG.I("[%s] trigger!!!!! %d", v.srvName, totolErrCnt)
				content := fmt.Sprintf("[服务异常] %s，近%d分钟，错误条数%d，请保持关注!", v.srvName, config.Cfg.MonitorTime, totolErrCnt)
				msgsend(content)
				v.lastTriggerTimeStamp = time.Now().Local().Unix()
			}
		LOOP_END:
			continue
			//log.LOG.I("[MonitorTriggerTask] [%s] complete...", v.srvName)
		}
		//探测频率
		<-time.After(time.Second)
	}
}

func msgsend(content string) {
	req1 := wxcorp.WxGetTokenReq{}
	req1.Corpid = config.Cfg.WxCorpID
	req1.Corpsecret = config.Cfg.WxCorpSecret
	resp1, _ := wxcorp.WxGetToken(req1)

	req2 := wxcorp.WxCorpMsgSendReq{}
	req2.Agentid = int64(config.Cfg.WxCorpAgent)
	req2.Msgtype = "text"
	req2.Safe = 0
	req2.Text.Content = content
	req2.Toparty = ""
	req2.Totag = ""
	req2.Touser = "@all"
	wxcorp.WxCorpMsgSend(req2, resp1.AccessToken)
}
