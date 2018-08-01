package handler

import (
	"logmonitor/task/data"
)

type LogMonitorReq struct {
	Content  string `json:"content" binding:"required"`
	FileLine string `json:"file_line" binding:"required"`
	SrvName  string `json:"srv_name" binding:"required"`
}

type LogMonitorResp struct {
	ErrorCode int64  `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

func LogMonitor(req LogMonitorReq) (*LogMonitorResp, error) {
	rst := &LogMonitorResp{}

	data.SrvStat(req.SrvName).Set()

	return rst, nil
}
