package ecs

import (
	"testing"
	"fmt"
	"time"
)

func TestCreateVpc(t *testing.T) {
	cli := NewClient("YouAppKeyID", "YourAppKeySecret", "")

	// test create vpc
	createResp, err := cli.Vpc.CreateVpc(map[string]string{
		"RegionId":"cn-beijing",
		"CidrBlock":"10.0.0.0/8",
		"VpcName":"testvpc",
		"Description": "testvpc",
	})
	if err != nil {
		t.Errorf("runtime exception:", err)
		return
	}
	fmt.Println(createResp)

	defer func() {
		// test delete vpc
		deleteResp, err := cli.Vpc.DeleteVpc(map[string]string {
			"VpcId": createResp.VpcId,
		})
		if err != nil {
			t.Errorf("runtime exception:", err)
			return
		}
		fmt.Println(deleteResp)
	}()

	// test vrouter
	descrVRouterRsp, err := cli.VRouterService.DescribeVRouters(map[string]string {
		"RegionId": "cn-beijing",
	})
	if err != nil {
		t.Errorf("runtime exception:", err)
		return
	}
	fmt.Println(descrVRouterRsp)

	descrRouteTableRsp, err := cli.RouteService.DescribeRouteTables(map[string]string {
		"VRouterId": descrVRouterRsp.VRouters.VRouter[0].VRouterId,
	})
	if err != nil {
		t.Errorf("runtime exception:", err)
		return
	}
	fmt.Println(descrRouteTableRsp)

	// test vswitch
	descrVSwitchRsp, err := cli.VSwitchService.DescribeVSwitches(map[string]string {
		"VpcId": createResp.VpcId,
	})
	if err != nil {
		t.Errorf("runtime exception:", err)
		return
	}
	fmt.Println(descrVSwitchRsp)

	descResp, err := cli.Vpc.DescribeVpcs(map[string]string {
		"RegionId":"cn-beijing",
	});
	if err != nil {
		t.Errorf("runtime exception:", err)
		return
	}
	fmt.Println(descResp)

	time.Sleep(time.Second * 5)
}