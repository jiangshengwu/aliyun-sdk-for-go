package ecs

import "github.com/jiangshengwu/aliyun-sdk-for-go/util"

type NetworkService interface {
	AllocatePublicIpAddress(params map[string]string) (AllocatePublicIpAddressResponse, error)
	ModifyInstanceNetworkSpec(params map[string]string) (ModifyInstanceNetworkSpecResponse, error)
	AllocateEipAddress(params map[string]string) (AllocateEipAddressResponse, error)
	AssociateEipAddress(params map[string]string) (AssociateEipAddressResponse, error)
	DescribeEipAddresses(params map[string]string) (DescribeEipAddressesResponse, error)
	ModifyEipAddressAttribute(params map[string]string) (ModifyEipAddressAttributeResponse, error)
	UnassociateEipAddress(params map[string]string) (UnassociateEipAddressResponse, error)
	ReleaseEipAddress(params map[string]string) (ReleaseEipAddressResponse, error)
}

type NetworkOperator struct {
	Common *util.CommonParam
}

// Response struct for AllocatePublicIpAddress
type AllocatePublicIpAddressResponse struct {
	util.ErrorResponse
	IpAddress string `json:"IpAddress"`
}

// Response struct for ModifyInstanceNetworkSpec
type ModifyInstanceNetworkSpecResponse struct {
	util.ErrorResponse
}

// Response struct for AllocateEipAddress
type AllocateEipAddressResponse struct {
	util.ErrorResponse
	EipAddress   string `json:"EipAddress"`
	AllocationId string `json:"AllocationId"`
}

// Response struct for AssociateEipAddress
type AssociateEipAddressResponse struct {
	util.ErrorResponse
}

// Response struct for DescribeEipAddresses
type DescribeEipAddressesResponse struct {
	util.ErrorResponse
	util.PageResponse
	AllEipAddresses EipAddresses `json:"EipAddresses"`
}

type EipAddresses struct {
	AllEipAddress []EipAddressSetType `json:"EipAddress"`
}

// See http://docs.aliyun.com/?spm=5176.775974174.2.4.BYfRJ2#/ecs/open-api/datatype&eipaddresssettype
type EipAddressSetType struct {
	RegionId           string         `json:"RegionId"`
	IpAddress          string         `json:"IpAddress"`
	AllocationId       string         `json:"AllocationId"`
	Status             string         `json:"Status"`
	InstanceId         string         `json:"InstanceId"`
	Bandwidth          string         `json:"Bandwidth"`
	InternetChargeType string         `json:"InternetChargeType"`
	AllOperationLocks  OperationLocks `json:"OperationLocks"`
	AllocationTime     string         `json:"AllocationTime"`
}

// Response struct for ModifyEipAddressAttribute
type ModifyEipAddressAttributeResponse struct {
	util.ErrorResponse
}

// Response struct for UnassociateEipAddress
type UnassociateEipAddressResponse struct {
	util.ErrorResponse
}

// Response struct for ReleaseEipAddress
type ReleaseEipAddressResponse struct {
	util.ErrorResponse
}

func (op *NetworkOperator) AllocatePublicIpAddress(params map[string]string) (AllocatePublicIpAddressResponse, error) {
	var resp AllocatePublicIpAddressResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *NetworkOperator) ModifyInstanceNetworkSpec(params map[string]string) (ModifyInstanceNetworkSpecResponse, error) {
	var resp ModifyInstanceNetworkSpecResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *NetworkOperator) AllocateEipAddress(params map[string]string) (AllocateEipAddressResponse, error) {
	var resp AllocateEipAddressResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *NetworkOperator) AssociateEipAddress(params map[string]string) (AssociateEipAddressResponse, error) {
	var resp AssociateEipAddressResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *NetworkOperator) DescribeEipAddresses(params map[string]string) (DescribeEipAddressesResponse, error) {
	var resp DescribeEipAddressesResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *NetworkOperator) ModifyEipAddressAttribute(params map[string]string) (ModifyEipAddressAttributeResponse, error) {
	var resp ModifyEipAddressAttributeResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *NetworkOperator) UnassociateEipAddress(params map[string]string) (UnassociateEipAddressResponse, error) {
	var resp UnassociateEipAddressResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *NetworkOperator) ReleaseEipAddress(params map[string]string) (ReleaseEipAddressResponse, error) {
	var resp ReleaseEipAddressResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}
