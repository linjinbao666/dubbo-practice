FROM golang

#ENV AppLogConfFile /go/src/dubbo-practice/conf/log.yml
#ENV DUBBO_GO_CONFIG_PATH  /go/src/dubbo-practice/conf/dubbogo.yaml

WORKDIR /go/src/dubbo-practice/go-server

#ADD dubbogo.yaml /go/src/dubbo-practice/conf/dubbogo.yaml
ADD go-server /usr/local/bin/go-server

CMD /usr/local/bin/go-server
