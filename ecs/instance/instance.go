package instance

import (
	"fmt"

	"github.com/jiangshengwu/aliyun-sdk-for-go/util"
)

type IInstance interface {
	DescribeInstanceStatus(params map[string]string) (string, error)
	DescribeInstanceAttribute(params map[string]string) (string, error)
	DescribeInstances(params map[string]string) (string, error)
}

type InstanceOperator struct {
}

func (op *InstanceOperator) DescribeInstanceStatus(params map[string]string) (string, error) {
	query := util.GetQueryFromMap(params)
	req := &util.AliyunRequest{
		util.ECS_HOST + query[1:],
	}
	fmt.Println(req.Url)
	result, err := req.DoGetRequest()
	return result, err
}

func (op *InstanceOperator) DescribeInstanceAttribute(params map[string]string) (string, error) {
	query := util.GetQueryFromMap(params)
	req := &util.AliyunRequest{
		util.ECS_HOST + query[1:],
	}
	fmt.Println(req.Url)
	result, err := req.DoGetRequest()
	return result, err
}

func (op *InstanceOperator) DescribeInstances(params map[string]string) (string, error) {
	query := util.GetQueryFromMap(params)
	req := &util.AliyunRequest{
		util.ECS_HOST + query[1:],
	}
	fmt.Println(req.Url)
	result, err := req.DoGetRequest()
	return result, err
}
