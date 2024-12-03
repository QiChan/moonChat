package service

import (
	"context"
	"log"

	v1 "moonChat/feedInterface/api/helloworld/v1"
	pb "moonChat/userCenterInterface/api/userInfo/v1"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/hashicorp/consul/api"
)

type UserService struct {
	pb.UnimplementedUserServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	return &pb.CreateUserReply{}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	return &pb.UpdateUserReply{}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	return &pb.DeleteUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = "127.0.0.1:8500"
	consulClient, err := api.NewClient(consulConfig)
	//获取服务发现管理器
	dis := consul.New(consulClient)
	if err != nil {
		log.Fatal(err)
	}
	//连接目标grpc服务器
	endPoint := "discovery:///feed"
	conn, err := grpc.DialInsecure(
		ctx,
		grpc.WithEndpoint(endPoint),
		grpc.WithDiscovery(dis),
	)

	if err != nil {
		return &pb.GetUserReply{Id: req.Id, Msg: "grpc连接feed服失败"}, nil
	}

	client := v1.NewGreeterClient(conn)
	resp, err := client.SayHello(ctx, &v1.HelloRequest{Name: "kratos"})
	if err != nil {
		return &pb.GetUserReply{Id: req.Id, Msg: "grpc执行SayHello()失败"}, nil
	}

	return &pb.GetUserReply{Id: req.Id, Msg: "grpc call from userCenter to feed success, resp: " + resp.Message}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}
