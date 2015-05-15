package util

type Response interface {
}

type BaseResponse struct {
	RequestId string `json:"RequestId"`
}

type ErrorResponse struct {
	BaseResponse
	HostId  string `json:"HostId"`
	Code    string `json:"Code"`
	Message string `json:"Message"`
}
