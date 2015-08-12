package ecs

import "github.com/jiangshengwu/aliyun-sdk-for-go/util"

type RouteService interface {
	/**
	 * RouteTableId(required): RouteTableId
	 * DestinationCidrBlock(required): RouteEntry 的目的网段
	 * NextHopType: 下一跳的类型，可选值为 Instance | Tunnel，默认值为 Instance
	 * NextHopId: 路由条目的下一跳
	 * ClientToken(optional):
	 */
	CreateRouteEntry(param map[string]string) (CreateRouteEntryResponse, error)

	/**
	 * RouteTableId(required):
	 * DestinationCidrBlock(required):
	 * NextHopId(required):
	 */
	DeleteRouteEntry(params map[string]string) (DeleteRouteEntryResponse, error)

	/**
	 * VRouterId(required)
	 * RouteTableId
	 * PageNumber
	 * PageSize
	 */
	DescribeRouteTables(params map[string]string) (DescribeRouteTablesResponse, error)
}

type RouteOperator struct {
	Common *CommonParam
}

type CreateRouteEntryResponse struct {
	util.ErrorResponse
}

type DeleteRouteEntryResponse struct {
	util.ErrorResponse
}

type DescribeRouteTablesResponse struct {
	util.ErrorResponse
	util.PageResponse
	RouteTables RouteTableSetTypes `json:"RouteTables"`
}

type RouteTableSetTypes struct {
	RouteTableSetType []RouteTableSetType `json:"RouteTable"`
}

type RouteTableSetType struct {
	VRouterId      string             `json:"VRouterId"`
	RouteTableId   string             `json:"RouteTableId"`
	RouteEntrys    RouteEntrySetTypes `json:"RouteEntrys"`
	RouteTableType string             `json:"RouteTableType"`
	CreationTime   string             `json:"CreationTime"`
}

type RouteEntrySetTypes struct {
	RouteEntry []RouteEntrySetType `json:"RouteEntry"`
}

type RouteEntrySetType struct {
	RouteTableId         string `json:"RouteTableId"`
	DestinationCidrBlock string `json:"DestinationCidrBlock"`
	Type                 string `json:"Type"`
	NextHopId            string `json:"NextHopId"`
	Status               string `json:"Status"`
}

func (op *RouteOperator) CreateRouteEntry(params map[string]string) (CreateRouteEntryResponse, error) {
	var resp CreateRouteEntryResponse
	err := op.Common.Request(GetFuncName(1), params, &resp)
	return resp, err
}

func (op *RouteOperator) DeleteRouteEntry(params map[string]string) (DeleteRouteEntryResponse, error) {
	var resp DeleteRouteEntryResponse
	err := op.Common.Request(GetFuncName(1), params, &resp)
	return resp, err
}

func (op *RouteOperator) DescribeRouteTables(params map[string]string) (DescribeRouteTablesResponse, error) {
	var resp DescribeRouteTablesResponse
	err := op.Common.Request(GetFuncName(1), params, &resp)
	return resp, err
}
