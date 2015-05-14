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
cli := &client.EcsClient{}
cli.AccessKeyId = "Your Access Key Id"
cli.AccessKeySecret = "Your Access Key Secrect"
cli.Format = "JSON" //optional
cli.Init() // Ready to use

// Perform request
result, err := cli.Execute("StopInstance", map[string]string{
    "InstanceId": "i-253op6931",
})
if (err != nil) {
    log.Fatal(err)
    return
}
fmt.Println(result)
```