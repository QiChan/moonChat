package handler

import (
	"context"
	"net/http"
	"strings"

	wb "moonChat/gateway/websocket"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/gorilla/websocket"
	"github.com/hashicorp/consul/api"
)

var (
	ConsulAddr = "127.0.0.1:8500"

	WsPool *wb.ConnectionPool

	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // 允许所有来源的WebSocket连接，生产环境中应该更严格
		},
	}
)

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

func UpgradeHttpToWs(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	WsPool.HandleConnection(conn)
}
