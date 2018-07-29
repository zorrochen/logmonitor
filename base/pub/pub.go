package pub

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

// 拷贝http请求的body
func CopyHttpRequestBody(r *http.Request) string {
	buf, _ := ioutil.ReadAll(r.Body)
	body := ioutil.NopCloser(bytes.NewBuffer(buf))
	r.Body = body
	return string(buf)
}

// 拷贝http请求的body([]byte)
func CopyHttpRequestBodyBytes(r *http.Request) []byte {
	buf, _ := ioutil.ReadAll(r.Body)
	body := ioutil.NopCloser(bytes.NewBuffer(buf))
	r.Body = body
	return buf
}

// 正则校验
func RegexpCheck(sourceStr string, ruleStr string) bool {
	reg := regexp.MustCompile(ruleStr)
	car_match := reg.FindAllString(sourceStr, -1)
	if car_match == nil {
		return false
	}
	return true
}

// MD5
func MD5(s string) (res string) {
	h := md5.New()
	h.Write([]byte(s))
	res = hex.EncodeToString(h.Sum(nil))
	res = strings.ToUpper(res)
	return
}
