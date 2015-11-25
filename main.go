package main

import (
	"github.com/jiangshengwu/aliyun-sdk-for-go/log"
	"github.com/jiangshengwu/aliyun-sdk-for-go/ecs"
)

func main() {
	// Initialize client
	cli := ecs.NewClient(
		"n4qTpXAgxmjtqTd5",
		"gjXgWdw3fIPNMu2PbiOkzHmKBSjqvx",
		"", //ResourceOwnerAccount", // optional, set to empty string if it's no need for you
	)

	// Perform request
	result, err := cli.Region.DescribeRegions(map[string]string{
		//result, err := cli.Listener.DescribeLoadBalancerHTTPListenerAttribute(map[string]string{
		//result, err := cli.BackendServer.DescribeHealthStatus(map[string]string{
		"RegionId":            "cn-beijing",
		"SecurityGroupId":     "sg-253yqueea",
		"LoadBalancerId":      "15137ad7161-cn-beijing-btc-a01",
		"LoadBalancerStatus":  "active",
		"LoadBalancerName":    "slbtest",
		"ListenerPort":        "1234",
		"AccessControlStatus": "open_white_list",
		"BackendServerPort":   "4321",
		"Bandwidth":           "1000",
		"StickySession":       "off",
		"HealthCheck":         "off",
		"BackendServers":      "['i-25qiriruf']",
	})

	if err != nil {
		log.Error(err)
		return
	}
	log.Info(result)
}
