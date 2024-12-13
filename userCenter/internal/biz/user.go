package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

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
	ProducerStart(context.Context) error
	ConsumerStart(context.Context) error
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
	uc.mq.ProducerStart(ctx)
	uc.mq.SndMsg(ctx, "test", "Hello RocketMQ Go Client!")
}

func (uc *UserUsecase) ActiveConsumer(ctx context.Context) {
	//uc.mq.DealMsg(ctx, "test")
	uc.mq.DealMsgWithTimeElapse(ctx, "test")
	uc.mq.DealMsgOrderly(ctx, "orderlyTest")
	uc.mq.ConsumerStart(ctx)
}

func (uc *UserUsecase) PublishMsg(ctx context.Context) {
	uc.mq.SndMsg(ctx, "test", "Hello RocketMQ Go Client!")
}
