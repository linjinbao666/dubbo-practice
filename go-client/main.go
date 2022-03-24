package main

import (
	"context"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

import (
	"go-client/api"
)

var greeterProvider = &api.UserProviderClientImpl{}

func init() {
	// validate consumer greeterProvider ptr
	config.SetConsumerService(greeterProvider)
}

func main() {
	// init rootConfig with config api
	rc := config.NewRootConfigBuilder().
		SetConsumer(config.NewConsumerConfigBuilder().
			AddReference("UserProviderClientImpl", config.NewReferenceConfigBuilder().
				SetProtocol("tri").
				SetRegistryIDs("registryKey").
				Build()).
			Build()).
		AddRegistry("registryKey",
			config.NewRegistryConfigBuilder().
				SetAddress("127.0.0.1:2181").
				SetProtocol("kubernetes").
				Build()).
		Build()

	// start dubbo-go framework with configuration
	if err := config.Load(config.WithRootConfig(rc)); err != nil {
		panic(err)
	}

	// run rpc invocation
	testSayHello()
}

func testSayHello() {
	ctx := context.Background()

	req := api.HelloRequest{
		Name: "laurence",
	}
	user, err := greeterProvider.SayHello(ctx, &req)
	if err != nil {
		panic(err)
	}

	logger.Infof("Receive user = %+v\n", user)
}
