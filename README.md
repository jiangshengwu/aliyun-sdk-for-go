# Aliyun SDK for Go

The encapsulation of `Aliyun` HTTP API

### Build Status

![Travis CI](https://travis-ci.org/jiangshengwu/aliyun-sdk-for-go.svg)

Version 0.1.0(beta)

## Progress

[Here](http://docs.aliyun.com/?spm=5176.775974174.2.4.BYfRJ2#/ecs/open-api/apisummary) to read more about ECS API.

Implemented ECS API

* Instance
* Disk
* Snapshot
* Image
* Network
* Security Group
* Region
* Monitor
* Other
* Vpc
* VRouter
* VSwitch
* Route

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
result, err := cli.SecurityGroup.DescribeSecurityGroupAttribute(map[string]string{
    "RegionId":        "cn-beijing",
    "SecurityGroupId": "sg-25rh80j7f",
})

if err != nil {
    log.Error(err)
    return
}
log.Info(result)
```