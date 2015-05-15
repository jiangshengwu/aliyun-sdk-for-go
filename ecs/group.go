package ecs

import (
	"encoding/json"

	"github.com/jiangshengwu/aliyun-sdk-for-go/log"
	"github.com/jiangshengwu/aliyun-sdk-for-go/util"
)

type GroupService interface {
	CreateSecurityGroup(params map[string]string) (CreateSecurityGroupResponse, error)
	AuthorizeSecurityGroup(params map[string]string) (AuthorizeSecurityGroupResponse, error)
	DescribeSecurityGroupAttribute(params map[string]string) (DescribeSecurityGroupAttributeResponse, error)
	DescribeSecurityGroups(params map[string]string) (DescribeSecurityGroupsResponse, error)
}

type GroupOperator struct {
	Common *CommonParam
}

// Response struct for CreateSecurityGroup
type CreateSecurityGroupResponse struct {
	util.ErrorResponse
}

// Response struct for AuthorizeSecurityGroup
type AuthorizeSecurityGroupResponse struct {
	util.ErrorResponse
}

// Response struct for DescribeSecurityGroupAttribute
type DescribeSecurityGroupAttributeResponse struct {
	util.ErrorResponse
	RegionId        string      `json:"RegionId"`
	SecurityGroupId string      `json:"SecurityGroupId"`
	Description     string      `json:"Description"`
	AllPermissions  Permissions `json:"Permissions"`
	VpcId           string      `json:"VpcId"`
}

type Permissions struct {
	AllPermission []PermissionType `json:"Permission"`
}

// See http://docs.aliyun.com/?spm=5176.775974174.2.4.BYfRJ2#/ecs/open-api/datatype&permissiontype
type PermissionType struct {
	IpProtocol              string `json:"IpProtocol"`
	PortRange               string `json:"PortRange"`
	SourceCidrIp            string `json:"SourceCidrIp"`
	SourceGroupId           string `json:"SourceGroupId"`
	SourceGroupOwnerAccount string `json:"SourceGroupOwnerAccount"`
	DestCidrIp              string `json:"DestCidrIp"`
	DestGroupId             string `json:"DestGroupId"`
	DestGroupOwnerAccount   string `json:"DestGroupOwnerAccount"`
	Policy                  string `json:"Policy"`
	NicType                 string `json:"NicType"`
	Priority                string `json:"Priority"`
}

// Response struct for DescribeSecurityGroupsResponse
type DescribeSecurityGroupsResponse struct {
	util.ErrorResponse
}

func (op *GroupOperator) CreateSecurityGroup(params map[string]string) (CreateSecurityGroupResponse, error) {
	var resp CreateSecurityGroupResponse
	action := GetFuncName(1)
	p := op.Common.ResolveAllParams(action, params)
	result, err := RequestAPI(p)
	if err != nil {
		return CreateSecurityGroupResponse{}, err
	}
	log.Infoln(result)
	json.Unmarshal([]byte(result), &resp)
	return resp, nil
}

func (op *GroupOperator) AuthorizeSecurityGroup(params map[string]string) (AuthorizeSecurityGroupResponse, error) {
	var resp AuthorizeSecurityGroupResponse
	action := GetFuncName(1)
	p := op.Common.ResolveAllParams(action, params)
	result, err := RequestAPI(p)
	if err != nil {
		return AuthorizeSecurityGroupResponse{}, err
	}
	log.Infoln(result)
	json.Unmarshal([]byte(result), &resp)
	return resp, nil
}

func (op *GroupOperator) DescribeSecurityGroupAttribute(params map[string]string) (DescribeSecurityGroupAttributeResponse, error) {
	var resp DescribeSecurityGroupAttributeResponse
	action := GetFuncName(1)
	p := op.Common.ResolveAllParams(action, params)
	result, err := RequestAPI(p)
	if err != nil {
		return DescribeSecurityGroupAttributeResponse{}, err
	}
	log.Infoln(result)
	json.Unmarshal([]byte(result), &resp)
	return resp, nil
}

func (op *GroupOperator) DescribeSecurityGroups(params map[string]string) (DescribeSecurityGroupsResponse, error) {
	var resp DescribeSecurityGroupsResponse
	action := GetFuncName(1)
	p := op.Common.ResolveAllParams(action, params)
	result, err := RequestAPI(p)
	if err != nil {
		return DescribeSecurityGroupsResponse{}, err
	}
	log.Infoln(result)
	json.Unmarshal([]byte(result), &resp)
	return resp, nil
}
