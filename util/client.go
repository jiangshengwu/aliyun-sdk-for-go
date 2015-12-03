package util

import (
	"time"

	"encoding/json"

	"github.com/jiangshengwu/aliyun-sdk-for-go/log"
)

type Client struct {
	Common *CommonParam
}

// struct for common parameters
type CommonParam struct {
	AccessKeyId          string
	AccessKeySecret      string
	ResourceOwnerAccount string
	Attr                 map[string]string
}

type ClientInterface interface {
	GetClientName() string
	GetVersion() string
	GetSignatureMethod() string
	GetSignatureVersion() string
}

func (c *CommonParam) RequestAPI(params map[string]string) (string, error) {
	query := GetQueryFromMap(params)
	req := &AliyunRequest{}
	req.Url = c.Attr["Host"] + query
	log.Debug(req.Url)
	result, err := req.DoGetRequest()
	return result, err
}

// Generate all parameters include Signature
func (c *CommonParam) ResolveAllParams(action string, params map[string]string) map[string]string {
	if params == nil {
		params = make(map[string]string, len(c.Attr))
	}
	// Process parameters
	for key, value := range c.Attr {
		params[key] = value
	}
	params["Action"] = action
	if c.ResourceOwnerAccount != "" {
		params["ResourceOwnerAccount"] = c.ResourceOwnerAccount
	}
	params["TimeStamp"] = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	params["SignatureNonce"] = GetGuid()
	sign := MapToSign(params, c.AccessKeySecret, c.Attr["HttpMethod"])
	params["Signature"] = sign
	return params
}

func (c *CommonParam) Request(action string, params map[string]string, response interface{}) error {
	p := c.ResolveAllParams(action, params)
	result, err := c.RequestAPI(p)
	if err != nil {
		return err
	}
	log.Debug(result)
	err = json.Unmarshal([]byte(result), response)
	if err != nil {
		log.Error(err)
	}
	return nil
}
