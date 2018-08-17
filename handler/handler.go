package handler

import (
	"github.com/gin-gonic/gin"
	"logmonitor/base/log"
	"net/http"
)

func pingHandler(c *gin.Context) {
	req := pingReq{}
	err := c.Bind(&req)
	if err != nil {
		log.LOG.E("%v,  %v\n", req, err)
		ErrResp(c, http.StatusBadRequest)
		return
	}

	resp, err := ping(req)
	if err != nil {
		log.LOG.E("%v, %v\n", req, err)
		ErrResp(c, http.StatusInternalServerError)
		return
	}

	SuccResp(c, resp)
}

func LogMonitorHandler(c *gin.Context) {
	req := LogMonitorReq{}
	err := c.Bind(&req)
	if err != nil {
		log.LOG.E("%v, %v\n", req, err)
		ErrResp(c, http.StatusBadRequest)
		return
	}

	resp, err := LogMonitor(req)
	if err != nil {
		log.LOG.E("%v, %v\n", req, err)
		ErrResp(c, http.StatusInternalServerError)
		return
	}

	SuccResp(c, resp)
}
