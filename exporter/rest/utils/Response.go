package utils

type SimpleResponse struct {
	Code int `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Msg string `json:"msg,omitempty"`
}

func Succeed(data interface{}) *SimpleResponse {
	return &SimpleResponse{Code: 0, Data: data}
}

func Failed(msg string) *SimpleResponse {
	return &SimpleResponse{Code: 1, Msg: msg}
}