package main

import (
	"fmt"

	"log"

	"github.com/jiangshengwu/aliyun-sdk-for-go/client"
)

func main() {
	// Initialize client
	cli := &client.EcsClient{}
	cli.AccessKeyId = "BEfUjzyexQKIiGV3"
	cli.AccessKeySecret = "9M5fW99d9spQsgUCg34Od9lR7XMlXl"
	cli.Format = "JSON" //optional
	cli.Init()          // Ready to use

	// Perform request
	result, err := cli.Execute("StopInstance", map[string]string{
		"InstanceId": "i-253op6931",
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(result)
}
