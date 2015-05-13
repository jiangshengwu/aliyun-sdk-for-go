package main

import (
	"fmt"

	"github.com/jiangshengwu/aliyun-sdk-for-go/client"
)

func main() {
	// Initialize client
	cli := &client.EcsClient{}
	cli.AccessKeyId = "myKeyId"
	cli.AccessKeySecret = "myKeySecret"
	cli.Format = "JSON"
	cli.Init() // Ready to use

	// Perform request
	result, _ := cli.Execute("DescribeInstanceAttribute", map[string]string{
		"InstanceId": "i-253op6931",
	})
	fmt.Println(result)
}
