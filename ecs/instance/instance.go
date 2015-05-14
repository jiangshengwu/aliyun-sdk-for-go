package instance

import (
	"fmt"

	"github.com/jiangshengwu/aliyun-sdk-for-go/util"
)

type IInstance interface {
	StartInstance(params map[string]string) (string, error)
	StopInstance(params map[string]string) (string, error)
	RebootInstance(params map[string]string) (string, error)
	DescribeInstanceStatus(params map[string]string) (string, error)
	DescribeInstanceAttribute(params map[string]string) (string, error)
	DescribeInstances(params map[string]string) (string, error)
}

type InstanceOperator struct {
}

func (op *InstanceOperator) StartInstance(params map[string]string) (string, error) {
	return requestAPI(params)
}

func (op *InstanceOperator) StopInstance(params map[string]string) (string, error) {
	return requestAPI(params)
}

func (op *InstanceOperator) RebootInstance(params map[string]string) (string, error) {
	return requestAPI(params)
}

func (op *InstanceOperator) DescribeInstanceStatus(params map[string]string) (string, error) {
	return requestAPI(params)
}

func (op *InstanceOperator) DescribeInstanceAttribute(params map[string]string) (string, error) {
	return requestAPI(params)
}

func (op *InstanceOperator) DescribeInstances(params map[string]string) (string, error) {
	return requestAPI(params)
}

func requestAPI(params map[string]string) (string, error) {
	query := util.GetQueryFromMap(params)
	req := &util.AliyunRequest{
		util.ECS_HOST + query[1:],
	}
	fmt.Println(req.Url)
	result, err := req.DoGetRequest()
	return result, err
}
