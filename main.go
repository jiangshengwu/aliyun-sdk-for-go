package main

import (
	"github.com/jiangshengwu/aliyun-sdk-for-go/ecs"
	"github.com/jiangshengwu/aliyun-sdk-for-go/log"
)

func main() {
	// Initialize client
	cli := ecs.NewClient(
		"Your Access Key Id",
		"Your Access Key Secret",
		"ResourceOwnerAccount", // optional, set to empty string if it's no need for you
	)

	// Perform request
	result, err := cli.SecurityGroup.DescribeSecurityGroupAttribute(map[string]string{
		"RegionId":        "cn-beijing",
		"SecurityGroupId": "sg-25rh80j7f",
	})

	if err != nil {
		log.Error(err)
		return
	}
	log.Info(result)
}
