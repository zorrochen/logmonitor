package handler

import (
	. "logmonitor/base/http_server"
)

func Init() {
	Register("GET", "/ping", pingHandler)
	Register("POST", "/logmonitor/error", LogMonitorHandler)
}
