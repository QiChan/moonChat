package handler

import (
	"context"
	"strings"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/hashicorp/consul/api"
)

var ConsulAddr = "127.0.0.1:8500"

func GetGrpcAddrs(inputs []string) []string {
	var grpcAddrs []string
	for _, str := range inputs {
		if strings.Contains(str, "grpc") {
			// 去除协议前缀部分，获取地址和端口部分
			addr := strings.TrimPrefix(str, "grpc://")
			grpcAddrs = append(grpcAddrs, addr)
		}
	}
	return grpcAddrs
}

func GetSrvInstances(ctx context.Context, srvName string) []*registry.ServiceInstance {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = ConsulAddr
	consulClient, err := api.NewClient(consulConfig)
	//获取服务发现管理器
	discovery := consul.New(consulClient)
	if err != nil {
		panic(err)
	}

	instances, _ := discovery.GetService(ctx, srvName)

	return instances
}
