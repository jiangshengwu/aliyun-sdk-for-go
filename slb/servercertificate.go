package slb

import "github.com/jiangshengwu/aliyun-sdk-for-go/util"

type ServerCertificateService interface {
	UploadServerCertificate(params map[string]string) (UploadServerCertificateResponse, error)
	DeleteServerCertificate(params map[string]string) (DeleteServerCertificateResponse, error)
	DescribeServerCertificates(params map[string]string) (DescribeServerCertificatesResponse, error)
	SetServerCertificateName(params map[string]string) (SetServerCertificateNameResponse, error)
}

type ServerCertificateOperator struct {
	Common *CommonParam
}

// Response struct for UploadServerCertificate
type UploadServerCertificateResponse struct {
	util.ErrorResponse
	ServerCertificateId   string `json:"ServerCertificateId"`
	ServerCertificateName string `json:"ServerCertificateName"`
	Fingerprint           string `json:"Fingerprint"`
}

// Response struct for DeleteServerCertificate
type DeleteServerCertificateResponse struct {
	util.ErrorResponse
}

// Response struct for DescribeServerCertificates
type DescribeServerCertificatesResponse struct {
	util.ErrorResponse
	ServerCertificates ServerCertificateList `json:"ServerCertificates"`
}

type ServerCertificateList struct {
	ServerCertificate []ServerCertificateType `json:"ServerCertificate"`
}

type ServerCertificateType struct {
	ServerCertificateId   string `json:"ServerCertificateId"`
	ServerCertificateName string `json:"ServerCertificateName"`
	RegionId              string `json:"RegionId"`
	Fingerprint           string `json:"Fingerprint"`
}

// Response struct for SetServerCertificateName
type SetServerCertificateNameResponse struct {
	util.ErrorResponse
}

func (op *ServerCertificateOperator) UploadServerCertificate(params map[string]string) (UploadServerCertificateResponse, error) {
	var resp UploadServerCertificateResponse
	err := op.Common.Request(GetFuncName(1), params, &resp)
	return resp, err
}

func (op *ServerCertificateOperator) DeleteServerCertificate(params map[string]string) (DeleteServerCertificateResponse, error) {
	var resp DeleteServerCertificateResponse
	err := op.Common.Request(GetFuncName(1), params, &resp)
	return resp, err
}

func (op *ServerCertificateOperator) DescribeServerCertificates(params map[string]string) (DescribeServerCertificatesResponse, error) {
	var resp DescribeServerCertificatesResponse
	err := op.Common.Request(GetFuncName(1), params, &resp)
	return resp, err
}

func (op *ServerCertificateOperator) SetServerCertificateName(params map[string]string) (SetServerCertificateNameResponse, error) {
	var resp SetServerCertificateNameResponse
	err := op.Common.Request(GetFuncName(1), params, &resp)
	return resp, err
}
