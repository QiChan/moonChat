package handler

import (
	"context"
	"net/http"

	jsondata "moonChat/gateway/jsonData"
	pb "moonChat/userCenterInterface/api/userInfo/v1"

	"github.com/gin-gonic/gin"
	"github.com/ruiboma/warlock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var UserCenterConnPool *warlock.Pool

func InitUserCenterGrpcConnPool() *warlock.Pool {
	cfg := warlock.NewConfig()
	srvInstances := GetSrvInstances(context.Background(), "userCenter")
	tmp := make([]string, 0)
	for _, ins := range srvInstances {
		tmp = append(tmp, GetGrpcAddrs(ins.Endpoints)...)
	}
	cfg.ServerAdds = &tmp
	cfg.OverflowCap = false

	pool, err := warlock.NewWarlock(cfg, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return pool

}

func GetUserHandler(c *gin.Context) {
	conn, close, err := UserCenterConnPool.Acquire()
	if err != nil {
		panic(err)
	}
	defer close()

	if err = c.ShouldBindJSON(&fromData); err != nil {
		c.JSON(500, gin.H{"parse json error: ": err.Error()})
	}

	req := &pb.GetUserRequest{Tag: jsondata.GetUserReq.Tag}
	client := pb.NewUserClient(conn)
	resp, err := client.GetUser(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "grpc执行GetUser()失败: " + err.Error()})
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "posted",
		"tag":     resp.Tag,
		"message": resp.Msg,
		"nick":    fromData.Nick,
	})
}
