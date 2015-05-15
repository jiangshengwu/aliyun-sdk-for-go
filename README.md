# Aliyun SDK for Go

The encapsulation of `Aliyun` HTTP API

### Status

![Travis CI](https://travis-ci.org/jiangshengwu/aliyun-sdk-for-go.svg)

Alpha

Just provide some ECS API support at present

### Install:

```bash
go get -u github.com/jiangshengwu/aliyun-sdk-for-go
```

### ECS Client Usage:

```go
// Initialize client
cli := ecs.NewClient(
    "Your Access Key Id",
    "Your Access Key Secret",
    "ResourceOwnerAccount", // optional, set to empty string if it's no need for you
)

// Perform request
result, err := cli.Group.DescribeSecurityGroupAttribute(map[string]string{
    "RegionId":        "cn-beijing",
    "SecurityGroupId": "sg-25rh80j7f",
})
if err != nil {
    log.Errorln(err)
    return
}
log.Infoln(result)
```