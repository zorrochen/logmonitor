package handler

import (
	"github.com/gin-gonic/gin"
	. "logmonitor/base/http_server"
	"net/http"
)

func Init() {
	go MonitorTask()
	go MonitorTriggerTask()

	Register("GET", "/ping", ping)
	Register("POST", "/logmonitor/error", LogMonitorHandler)
}

// ping
func ping(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}
