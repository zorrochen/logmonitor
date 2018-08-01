package task

import (
	"fmt"
	"logmonitor/base/log"
	"logmonitor/config"
	"logmonitor/proxy/wxcorp"
	"logmonitor/task/data"
	"time"
)

// 监控数据分析
func TaskMonitorStat() {
	for {
		MonitorStat()
		//频率: 1分钟/次
		<-time.After(time.Minute)
	}
}

// 监控触发
func TaskMonitorTrigger() {
	for {
		MonitorTrigger()
		//频率: 1分钟/次
		<-time.After(time.Minute)
	}
}

//================= MonitorStat =================
func MonitorStat() {
	for _, v := range data.GlobalSrvList() {
		data.SrvStat(v).StatSet()
	}
}

//================= MonitorTrigger =================
func MonitorTrigger() {
	for _, v := range data.GlobalSrvList() {
		totolErrCnt := 0
		logStatList, err := data.SrvStat(v).GetStatMap()
		if err != nil {
			continue
		}

		//统计指定时间内错误总数
		totolErrCnt = 0
		for k, v := range logStatList {
			if k > config.Cfg.MonitorTime {
				break
			}
			totolErrCnt += v
		}

		//触发
		if totolErrCnt >= config.Cfg.MonitorThreshold {
			//CD
			if time.Now().Local().Unix()-data.SrvStat(v).LastTriggerTimeStamp() < int64(config.Cfg.MonitorCD*60) {
				continue
			}

			//发送告警
			log.LOG.I("[%s] trigger!!!!! %d", v, totolErrCnt)
			content := fmt.Sprintf("[服务异常] %s，近%d分钟，错误条数%d，请保持关注!", v, config.Cfg.MonitorTime, totolErrCnt)
			msgsend(content)
			data.SrvStat(v).SetLastTriggerTimeStamp(time.Now().Local().Unix())
		}
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
