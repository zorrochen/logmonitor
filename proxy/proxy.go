package proxy

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"logmonitor/base/log"
	"logmonitor/base/pub"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"time"
)

//结构体转GET请求的querystr
func Struct2Querystr(s interface{}) string {
	if reflect.TypeOf(s).Kind() != reflect.Struct {
		return ""
	}

	uv := url.Values{}
	for i := 0; i < reflect.TypeOf(s).NumField(); i++ {
		uv.Add(reflect.TypeOf(s).Field(i).Tag.Get("json"), reflect.ValueOf(s).Field(i).String())
	}

	return uv.Encode()
}

func Get(url string) (int, []byte) {
	return reqProxy(url, "GET", nil, nil, 0)
}

func PostJson(url string, req interface{}) (int, []byte) {
	reqjson, _ := json.Marshal(req)

	header := map[string]string{
		"Content-Type": "application/json; charset=utf8",
	}

	return reqProxy(url, "POST", header, reqjson, 0)
}

// ================= 初始化 =================
// 长连接的httpclient
var httpClient *http.Client

// 长连接配置
func init() {
	httpClient = http.DefaultClient
	httpClient.Transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
		DisableCompression:  true,
	}
}

// ================= req =================
func reqProxy(url string, method string, header map[string]string, data []byte, timeout int64) (int, []byte) {
	// 请求初始化
	request, err := http.NewRequest(method, url, bytes.NewReader(data))
	if err != nil {
		return -1, nil
	}

	// 添加header
	for k, v := range header {
		request.Header.Add(k, v)
	}

	// 默认超时时间：3s
	if timeout == 0 {
		timeout = 3
	}

	// 添加超时
	ctx, _ := context.WithTimeout(context.TODO(), time.Duration(timeout)*time.Second)
	request = request.WithContext(ctx)

	// 添加请求打印
	log.LOG.Debug("[proxy-req] %s, %v, %s", request.URL.String(), request.Header, pub.CopyHttpRequestBody(request))

	// 发送请求
	resp, err := httpClient.Do(request)
	if err != nil {
		return -1, nil
	}
	defer resp.Body.Close()

	// 读取响应body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, nil
	}

	// 添加返回打印
	log.LOG.Debug("[proxy-return] %s, %s", request.URL.String(), string(body))

	// return
	return resp.StatusCode, body
}
