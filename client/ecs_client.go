package client

import (
	"reflect"
	"time"

	"github.com/jiangshengwu/aliyun-sdk-for-go/ecs/instance"
	"github.com/jiangshengwu/aliyun-sdk-for-go/util"
)

type EcsClient struct {
	AccessKeyId          string
	AccessKeySecret      string
	Format               string
	ResourceOwnerAccount string
	attr                 map[string]string
	instance.InstanceOperator
}

func (client *EcsClient) Init() {
	client.attr = map[string]string{
		"Version":          "2014-05-26",
		"AccessKeyId":      client.AccessKeyId,
		"SignatureMethod":  "HMAC-SHA1",
		"SignatureVersion": "1",
	}
}

func (client *EcsClient) GetVersion() string {
	return client.attr["Version"]
}

func (client *EcsClient) GetSignatureMethod() string {
	return client.attr["SignatureMethod"]
}

func (client *EcsClient) GetSignatureVersion() string {
	return client.attr["SignatureVersion"]
}

func (client *EcsClient) Execute(action string, params map[string]string) (string, error) {
	if params == nil {
		params = make(map[string]string, len(client.attr))
	}
	// Process parameters
	for key, value := range client.attr {
		params[key] = value
	}
	params["Action"] = action
	if client.Format != "" {
		params["Format"] = client.Format
	}
	if client.ResourceOwnerAccount != "" {
		params["ResourceOwnerAccount"] = client.ResourceOwnerAccount
	}
	params["TimeStamp"] = time.Now().UTC().Format("2006-01-02T15:04:05")
	params["SignatureNonce"] = util.GetGuid()
	sign := util.MapToSign(params, client.AccessKeySecret, util.ECS_HTTP_METHOD)
	params["Signature"] = sign

	// Call request func
	m := reflect.ValueOf(client).MethodByName(action)
	in := make([]reflect.Value, 1)
	in[0] = reflect.ValueOf(params)
	result := m.Call(in)
	var err error = nil
	e := result[1].Interface()
	if e != nil {
		err = result[1].Interface().(error)
	}
	return result[0].String(), err
}
