package client

import (
	"testing"

	"github.com/jiangshengwu/aliyun-sdk-for-go/ecs/instance"
)

func Test_Init1(t *testing.T) {
	cli := &EcsClient{
		"testKeyId1",
		"testKeySecrect1",
		"JSON",
		"True",
		nil,
		instance.InstanceOperator{},
	}
	cli.Init()

	if cli.Format != "JSON" {
		t.Error("Format value is incorrect.")
	}
	if cli.ResourceOwnerAccount != "True" {
		t.Error("ResourceOwnerAccount value is incorrect.")
	}
	if cli.GetSignatureMethod() != "HMAC-SHA1" {
		t.Error("SignatureMethod value is incorrect.")
	}
}

func Test_Init2(t *testing.T) {
	cli := &EcsClient{}
	cli.AccessKeyId = "testKeyId2"
	cli.AccessKeySecret = "testKeySecrect2"
	cli.Format = "XML"
	cli.Init()

	if cli.Format != "XML" {
		t.Error("Format value is incorrect.")
	}
	if cli.ResourceOwnerAccount != "" {
		t.Error("ResourceOwnerAccount value is incorrect.")
	}
	if cli.GetVersion() != "2014-05-26" {
		t.Error("Version value is incorrect.")
	}
}
