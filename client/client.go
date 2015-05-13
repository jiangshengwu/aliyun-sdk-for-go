// Aliyun API client package
package client

type IClient interface {
	Init()
	GetClientName() string
}
