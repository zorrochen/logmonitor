package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

var retMap map[int]string

func init() {
	retMap = map[int]string{
		http.StatusBadRequest:          "参数错误",
		http.StatusInternalServerError: "系统错误",
	}
}

func SuccResp(c *gin.Context, data interface{}) {
	//注入respJson,支持中间件统一响应日志打印
	respJson, _ := json.Marshal(data)
	c.Set("respJson", string(respJson))

	//返回前端
	c.JSON(http.StatusOK, data)
}

type errResp struct {
	ErrMsg string `json:"errMsg"`
}

func ErrResp(c *gin.Context, errcode int) {
	resp := errResp{
		ErrMsg: retMap[errcode],
	}

	//注入respJson,支持中间件统一响应日志打印
	respJson, _ := json.Marshal(resp)
	c.Set("respJson", string(respJson))

	//返回前端
	c.JSON(errcode, resp)
}
