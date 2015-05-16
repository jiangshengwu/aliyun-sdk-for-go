package util

import "fmt"

type SdkError struct {
	Resp ErrorResponse
}

func (err *SdkError) Error() string {
	return fmt.Sprintf("Request Error:\nHostId: %s\nCode: %s\nMessage: %s", err.Resp.HostId, err.Resp.Code, err.Resp.Message)
}
