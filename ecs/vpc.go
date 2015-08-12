package ecs

import (
	"github.com/jiangshengwu/aliyun-sdk-for-go/util"
)

type VpcService interface {
	/**
	 * RegionId: cn-beijing
	 * CidrBlock: 192.168.0.0/16, 172.16.0.0/12, 10.0.0.0/8
	 * VpcName:
	 * Description:
	 * ClientToken:
	 */
	CreateVpc(param map[string]string) (CreateVpcResponse, error)
	DeleteVpc(params map[string]string) (DeleteVpcResponse, error)
	DescribeVpcs(params map[string]string) (DescribeVpcsResponse, error)
	ModifyVpcAttribute(params map[string]string) (ModifyVpcAttributeResponse, error)
}

type VpcOperator struct {
	Common *CommonParam
}

// Response struct for AllocateEipAddress
type CreateVpcResponse struct {
	util.ErrorResponse
	VpcId        string `json:"VpcId"`
	VRouterId    string `json:"VRouterId"`
	RouteTableId string `json:"RouteTableId"`
}

type DeleteVpcResponse struct {
	util.ErrorResponse
}

type VpcSetTypes struct {
	Vpc []VpcSetType `json:"Vpc"`
}

type VpcSetType struct {
	VpcId        string `json:"VpcId"`
	RegionId     string `json:"RegionId"`
	Status       string `json:"Status"`
	VpcName      string `json:"VpcName"`
	VSwitchIds   string `json:"VSwitchIds"`
	CidrBlock    string `json:"CidrBlock"`
	VRouterId    string `json:"VRouterId"`
	Description  string `json:"Description"`
	CreationTime string `json:"CreationTime"`
}

type DescribeVpcsResponse struct {
	util.ErrorResponse
	util.PageResponse
	Vpcs VpcSetTypes `json:Vpcs`
}

type ModifyVpcAttributeResponse struct {
	util.ErrorResponse
}

func (op *VpcOperator) CreateVpc(params map[string]string) (CreateVpcResponse, error) {
	var resp CreateVpcResponse
	err := op.Common.Request(GetFuncName(1), params, &resp);
	return resp, err
}

func (op *VpcOperator) DeleteVpc(params map[string]string) (DeleteVpcResponse, error) {
	var resp DeleteVpcResponse
	err := op.Common.Request(GetFuncName(1), params, &resp);
	return resp, err
}

func (op *VpcOperator) DescribeVpcs(params map[string]string) (DescribeVpcsResponse, error) {
	var resp DescribeVpcsResponse
	err := op.Common.Request(GetFuncName(1), params, &resp);
	return resp, err
}

func (op *VpcOperator) ModifyVpcAttribute(params map[string]string) (ModifyVpcAttributeResponse, error) {
	var resp ModifyVpcAttributeResponse
	err := op.Common.Request(GetFuncName(1), params, &resp);
	return resp, err
}