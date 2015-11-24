package slb

import "github.com/jiangshengwu/aliyun-sdk-for-go/util"

type BackendServerService interface {
	AddBackendServers(params map[string]string) (AddBackendServersResponse, error)
	RemoveBackendServers(params map[string]string) (RemoveBackendServersResponse, error)
	SetBackendServers(params map[string]string) (SetBackendServersResponse, error)
	DescribeHealthStatus(params map[string]string) (DescribeHealthStatusResponse, error)
}

type BackendServerOperator struct {
	Common *CommonParam
}

type BackendServerList struct {
	BackendServer []BackendServerType `json:"BackendServer"`
}

type BackendServerType struct {
	ServerId           string `json:"ServerId"`
	Weight             int    `json:"Weight"`
	ServerHealthStatus string `json:"ServerHealthStatus"`
}

// Response struct for AddBackendServers
type AddBackendServersResponse struct {
	util.ErrorResponse
	LoadBalancerId string            `json:"LoadBalancerId"`
	BackendServers BackendServerList `json:"BackendServers"`
}

// Response struct for RemoveBackendServers
type RemoveBackendServersResponse struct {
	util.ErrorResponse
	LoadBalancerId string            `json:"LoadBalancerId"`
	BackendServers BackendServerList `json:"BackendServers"`
}

// Response struct for SetBackendServers
type SetBackendServersResponse struct {
	util.ErrorResponse
	LoadBalancerId string            `json:"LoadBalancerId"`
	BackendServers BackendServerList `json:"BackendServers"`
}

// Response struct for DescribeHealthStatus
type DescribeHealthStatusResponse struct {
	util.ErrorResponse
	BackendServers BackendServerList `json:"BackendServers"`
}

func (op *BackendServerOperator) AddBackendServers(params map[string]string) (AddBackendServersResponse, error) {
	var resp AddBackendServersResponse
	err := op.Common.Request(GetFuncName(1), params, &resp)
	return resp, err
}

func (op *BackendServerOperator) RemoveBackendServers(params map[string]string) (RemoveBackendServersResponse, error) {
	var resp RemoveBackendServersResponse
	err := op.Common.Request(GetFuncName(1), params, &resp)
	return resp, err
}

func (op *BackendServerOperator) SetBackendServers(params map[string]string) (SetBackendServersResponse, error) {
	var resp SetBackendServersResponse
	err := op.Common.Request(GetFuncName(1), params, &resp)
	return resp, err
}

func (op *BackendServerOperator) DescribeHealthStatus(params map[string]string) (DescribeHealthStatusResponse, error) {
	var resp DescribeHealthStatusResponse
	err := op.Common.Request(GetFuncName(1), params, &resp)
	return resp, err
}
