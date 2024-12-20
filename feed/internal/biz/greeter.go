package biz

import (
	"context"

	w1 "moonChat/feedInterface/api/helloworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	EwrUserNotFound = errors.NotFound(w1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Greeter is a Greeter model.
type Greeter struct {
	Hello string
}

// GreeterRepo is a Greater repo.
type GreeterRepo interface {
	Save(context.Context, *Greeter) (*Greeter, error)
	Update(context.Context, *Greeter) (*Greeter, error)
	FindByID(context.Context, int64) (*Greeter, error)
	ListByHello(context.Context, string) ([]*Greeter, error)
	ListAll(context.Context) ([]*Greeter, error)
}

type GreeterMQ interface {
	SndMsgSync(context.Context, string, string, string) (string, error)
	SndMsgSync_ano(context.Context, string, string, string) (string, error)
	SndMsgAsync(ctx context.Context, topic string, content string, tag string) error
	SndMsgOneWay(ctx context.Context, topic string, content string, tag string) error
	SndMsgBatch(ctx context.Context, topic string, content string, tag string) error
	SndMsgOrderly(ctx context.Context, topic string, content string, tag string) error
	SndMsgDelay(ctx context.Context, topic string, content string, tag string) error
	SndMsgTrans(ctx context.Context, topic string, content string, tag string) error
	SndMsgDelayAnyTime(ctx context.Context, topic string, content string, tag string, delayInterval int64) error
	ClientsStart(ctx context.Context) error
	TopicsCreate(ctx context.Context) error
}

// GreeterUsecase is a Greeter usecase.
type GreeterUsecase struct {
	repo GreeterRepo
	mq   GreeterMQ
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase(repo GreeterRepo, mq GreeterMQ, logger log.Logger) *GreeterUsecase {
	return &GreeterUsecase{repo: repo, mq: mq, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *GreeterUsecase) CreateGreeter(ctx context.Context, g *Greeter) (*Greeter, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}

func (uc *GreeterUsecase) ActiveProducer(ctx context.Context) {
	uc.mq.TopicsCreate(ctx)
	uc.mq.ClientsStart(ctx)
}

func (uc *GreeterUsecase) PublishMsg(ctx context.Context, tag string) {
	/*
		uc.mq.SndMsgDelay(ctx, "broadCastingTest", "Hello RocketMQ Go Client snd msg delay!", tag)
		uc.mq.SndMsgDelayAnyTime(ctx, "broadCastingTest", "Hello RocketMQ Go Client snd msg delay any time!", tag, 15)
	*/
	//uc.mq.SndMsgOneWay(ctx, "test", "Hello RocketMQ Go Client snd msg one way!", tag)
	//uc.mq.SndMsgAsync(ctx, "test", "Hello RocketMQ Go Client snd msg asynccccccccccccccccccc!", tag)
	uc.mq.SndMsgSync(ctx, "test", "Hello RocketMQ Go Client snd msg sync!", tag)
	//uc.mq.SndMsgSync_ano(ctx, "test", "Hello RocketMQ Go Client snd msg sync anooooooo!", tag)
	//uc.mq.SndMsgTrans(ctx, "test", "Hello RocketMQ Go Client snd msg trans!", tag)
	//uc.mq.SndMsgBatch(ctx, "test", "Hello RocketMQ Go Client snd msg batch!", tag)
	//uc.mq.SndMsgOrderly(ctx, "orderlyTest", "Hello RocketMQ Go Client snd msg orderly!", tag)
}
