# use example

```code

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-client 

docker build -t registry.cn-hangzhou.aliyuncs.com/amrom/dubbo-go-client:0.0.16 . --no-cache

docker push registry.cn-hangzhou.aliyuncs.com/amrom/dubbo-go-client:0.0.16

```