package main

import (
	"context"
	"log"

	pb "moonChat/userCenterInterface/api/userInfo/v1"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/hashicorp/consul/api"
)

var fromData struct {
	Tag  string `json:"Tag"`
	Nick string `json:"Nick"`
}

func main() {
	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		if err := c.BindJSON(&fromData); err != nil {
			c.JSON(500, gin.H{"error: ": err.Error()})
			panic(err)
		}

		grpcReq := &pb.GetUserRequest{Tag: fromData.Tag}
		resp, _ := DialGrpc(context.Background(), grpcReq)

		c.JSON(200, gin.H{
			"status":  "posted",
			"tag":     resp.Tag,
			"message": resp.Msg,
			"nick":    fromData.Nick,
		})
	})
	router.Run(":8091")
}

func DialGrpc(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = "127.0.0.1:8500"
	consulClient, err := api.NewClient(consulConfig)
	//获取服务发现管理器
	dis := consul.New(consulClient)
	if err != nil {
		log.Fatal(err)
	}
	//连接目标grpc服务器
	endPoint := "discovery:///userCenter"
	conn, err := grpc.DialInsecure(
		ctx,
		grpc.WithEndpoint(endPoint),
		grpc.WithDiscovery(dis),
	)

	if err != nil {
		return &pb.GetUserReply{Tag: req.Tag, Msg: "grpc连接userCenter服失败"}, nil
	}

	client := pb.NewUserClient(conn)
	resp, err := client.GetUser(ctx, &pb.GetUserRequest{Tag: req.Tag})
	if err != nil {
		return &pb.GetUserReply{Tag: req.Tag, Msg: "grpc执行GetUser()失败: " + err.Error()}, nil
	}

	return &pb.GetUserReply{Tag: req.Tag, Msg: "grpc call from gateway to userCenter success, resp: " + resp.Msg}, nil
}
