package biz

import (
	"context"
	v1 "moonChat/mqInterface/api/msgQueue/v1"

	"github.com/go-kratos/kratos/v2/log"
)

// 改成从配置里读取
var MqConfig = &v1.MQ_Config{
	v1.Producer: nil,
	v1.PushConsumer: &v1.MQ_Suber_Config{
		GroupName:   "userCenterConsumer_test",
		NameSvrAddr: "127.0.0.1:9876",
	},
	v1.ProducerOrderly: nil,
	v1.ConsumerOrderly: &v1.MQ_Suber_Config{
		GroupName:   "userCenterOrderlyConsumer_orderlyTest",
		NameSvrAddr: "127.0.0.1:9876",
	},
	v1.BroadCastConsumer: &v1.MQ_Suber_Config{
		GroupName:   "userCenterBroadCastConsumer_broadCastingTest",
		NameSvrAddr: "127.0.0.1:9876",
	},
}

type UserInfo struct {
	ID       int64  `json:"id"`
	UserName string `json:"userName"`
}

type UserRepo interface {
	FindByID(context.Context, int64) (*UserInfo, error)
}

type UserMQ interface {
	SndMsg(context.Context, string, string) (string, error)
	DealMsg(context.Context, string) error
	DealMsgWithTimeElapse(ctx context.Context, topic string) error
	DealMsgOrderly(ctx context.Context, topic string) error
	DealMsgBroadCasting(ctx context.Context, topic string) error
	ClientsStart(ctx context.Context, config *v1.MQ_Config) error
}

type UserUsecase struct {
	repo UserRepo
	mq   UserMQ
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, mq UserMQ, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, mq: mq, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) CreateUserInfo() *UserInfo {
	return &UserInfo{ID: 1, UserName: "test"}
}

func (uc *UserUsecase) ActiveProducer(ctx context.Context) {
	uc.mq.ClientsStart(ctx, MqConfig)
	uc.mq.SndMsg(ctx, "test", "Hello RocketMQ Go Client!")
}

func (uc *UserUsecase) ActiveConsumer(ctx context.Context) {
	//uc.mq.DealMsg(ctx, "test")
	uc.mq.DealMsgWithTimeElapse(ctx, "test")
	uc.mq.DealMsgOrderly(ctx, "orderlyTest")
	uc.mq.DealMsgBroadCasting(ctx, "broadCastingTest")
	uc.mq.ClientsStart(ctx, MqConfig)
}

func (uc *UserUsecase) PublishMsg(ctx context.Context) {
	uc.mq.SndMsg(ctx, "test", "Hello RocketMQ Go Client!")
}
