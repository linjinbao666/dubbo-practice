FROM golang

#ENV AppLogConfFile /go/src/dubbo-practice/conf/log.yml
#ENV DUBBO_GO_CONFIG_PATH  /go/src/dubbo-practice/conf/dubbogo.yaml

WORKDIR /go/src/dubbo-practice/go-client

#ADD dubbogo.yaml /go/src/dubbo-practice/conf/dubbogo.yaml
ADD go-client /usr/local/bin/go-client

CMD /usr/local/bin/go-client
