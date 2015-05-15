package ecs

type InstanceService interface {
	StartInstance(params map[string]string) (string, error)
	StopInstance(params map[string]string) (string, error)
	RebootInstance(params map[string]string) (string, error)
	DescribeInstanceStatus(params map[string]string) (string, error)
	DescribeInstanceAttribute(params map[string]string) (string, error)
	DescribeInstances(params map[string]string) (string, error)
}

type InstanceOperator struct {
	Common *CommonParam
}

func (op *InstanceOperator) StartInstance(params map[string]string) (string, error) {
	return RequestAPI(params)
}

func (op *InstanceOperator) StopInstance(params map[string]string) (string, error) {
	return RequestAPI(params)
}

func (op *InstanceOperator) RebootInstance(params map[string]string) (string, error) {
	return RequestAPI(params)
}

func (op *InstanceOperator) DescribeInstanceStatus(params map[string]string) (string, error) {
	return RequestAPI(params)
}

func (op *InstanceOperator) DescribeInstanceAttribute(params map[string]string) (string, error) {
	return RequestAPI(params)
}

func (op *InstanceOperator) DescribeInstances(params map[string]string) (string, error) {
	return RequestAPI(params)
}
