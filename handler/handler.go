package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"logmonitor/base/log"
)

func pingHandler(c *gin.Context) {
	req := pingReq{}
	err := c.Bind(&req)
	if err != nil {
		log.LOG.E("%v, %v\n", req, err)
		return
	}

	resp, err := ping(req)
	if err != nil {
		log.LOG.E("%v, %v\n", req, err)
		return
	}

	respJson, _ := json.Marshal(resp)
	c.Set("respJson", respJson)
	c.JSON(200, resp)
}

func LogMonitorHandler(c *gin.Context) {
	req := LogMonitorReq{}
	err := c.Bind(&req)
	if err != nil {
		log.LOG.E("%v, %v\n", req, err)
		return
	}

	resp, err := LogMonitor(req)
	if err != nil {
		log.LOG.E("%v, %v\n", req, err)
		return
	}

	respJson, _ := json.Marshal(resp)
	c.Set("respJson", respJson)
	c.JSON(200, resp)
}
