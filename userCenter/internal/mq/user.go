package mq

import (
	"context"
	"fmt"

	v1 "moonChat/mqInterface/api/msgQueue/v1"
	"moonChat/userCenter/internal/biz"

	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewUserMQ, v1.NewMQ)

type userMQ struct {
	mq  *v1.MQ
	log *log.Helper
}

// NewGreeterRepo .
func NewUserMQ(mq *v1.MQ, logger log.Logger) biz.UserMQ {
	return &userMQ{
		mq:  mq,
		log: log.NewHelper(logger),
	}
}

func (r *userMQ) SndMsg(ctx context.Context, topic string, content string) (string, error) {
	msg := primitive.NewMessage(topic, []byte(content))
	res, err := r.mq.Producer.SendSync(context.Background(), msg)
	if err != nil {
		fmt.Printf("send message error: %s\n", err)
		return err.Error(), err
	} else {
		fmt.Printf("send message success: result=%s\n", res.String())
		return res.String(), nil
	}
}

func (r *userMQ) DealMsg(ctx context.Context, topic string) error {
	selector := consumer.MessageSelector{
		Type:       consumer.TAG,
		Expression: "",
	}
	err := r.mq.PushConsumer.Subscribe(topic, selector, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, msg := range msgs {
			fmt.Printf("subscribe callback: %v \n", string(msg.Body))
		}
		return consumer.ConsumeSuccess, nil
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	err = r.mq.PushConsumer.Start()
	if err != nil {
		fmt.Printf("start consumer error: %s", err.Error())
		return nil
	}

	return nil
}

func (r *userMQ) ProducerStart(ctx context.Context) error {
	err := r.mq.Producer.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		return err
	}
	return nil
}

func (r *userMQ) ConsumerStart(ctx context.Context) error {
	err := r.mq.PushConsumer.Start()
	if err != nil {
		fmt.Printf("start consumer error: %s", err.Error())
		return err
	}
	return nil
}
