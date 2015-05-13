package util

type SdkError struct {
	Message string
}

func (err *SdkError) Error() string {
	return err.Message
}
