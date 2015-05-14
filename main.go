package main

import (
	"fmt"

	"log"

	"github.com/jiangshengwu/aliyun-sdk-for-go/client"
)

func main() {
	// Initialize client
	cli := &client.EcsClient{}
	cli.AccessKeyId = "testId"
	cli.AccessKeySecret = "testKey"
	cli.Format = "JSON" //optional
	cli.Init()          // Ready to use

	// Perform request
	result, err := cli.Execute("StartInstance", map[string]string{
		"InstanceId": "i-253op6931",
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(result)
}
