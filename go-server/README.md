## use examle

```code

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-server 

docker build -t registry.cn-hangzhou.aliyuncs.com/amrom/dubbo-go-server:0.0.16 . --no-cache 

docker push registry.cn-hangzhou.aliyuncs.com/amrom/dubbo-go-server:0.0.16

```
