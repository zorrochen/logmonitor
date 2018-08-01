package handler

//================= ping =================
type pingReq struct{}

type pingResp struct {
	ErrorCode int64  `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

//ping测试
func ping(req pingReq) (*pingResp, error) {
	rst := &pingResp{}
	return rst, nil
}
