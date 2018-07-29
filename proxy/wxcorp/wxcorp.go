package wxcorp

import (
	"encoding/json"
	"errors"
	"fmt"
	"logmonitor/proxy"
	"net/http"
)

//================= WxGetToken =================
type WxGetTokenReq struct {
	Corpid     string `json:"corpid"`
	Corpsecret string `json:"corpsecret"`
}

type WxGetTokenResp struct {
	AccessToken string `json:"access_token"`
	Errorcode   int64  `json:"errorcode"`
	Errormsg    string `json:"errormsg"`
	ExpiresIn   int64  `json:"expires_in"`
}

//获取AccessToken
func WxGetToken(req WxGetTokenReq) (*WxGetTokenResp, error) {
	rst := &WxGetTokenResp{}

	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", req.Corpid, req.Corpsecret)
	httpcode, body := proxy.Get(url)
	if httpcode != http.StatusOK {
		return nil, errors.New("request failed.")
	}

	err := json.Unmarshal(body, rst)
	if err != nil {
		return nil, err
	}

	return rst, nil
}

//================= WxCorpMsgSend =================
type WxCorpMsgSendReq struct {
	Agentid int64  `json:"agentid"`
	Msgtype string `json:"msgtype"`
	Safe    int64  `json:"safe"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
	Toparty string `json:"toparty"`
	Totag   string `json:"totag"`
	Touser  string `json:"touser"`
}

type WxCorpMsgSendResp struct {
	Errcode      int64  `json:"errcode"`
	Errmsg       string `json:"errmsg"`
	Invalidparty string `json:"invalidparty"`
	Invalidtag   string `json:"invalidtag"`
	Invaliduser  string `json:"invaliduser"`
}

//企业号消息发送
func WxCorpMsgSend(req WxCorpMsgSendReq, token string) (*WxCorpMsgSendResp, error) {
	rst := &WxCorpMsgSendResp{}

	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", token)
	httpcode, body := proxy.PostJson(url, req)
	if httpcode != http.StatusOK {
		return nil, errors.New("request failed.")
	}

	err := json.Unmarshal(body, rst)
	if err != nil {
		return nil, err
	}

	return rst, nil
}
