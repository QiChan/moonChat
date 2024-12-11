package mq

import (
	"context"
	"fmt"

	"moonChat/feed/internal/biz"
	v1 "moonChat/mqInterface/api/msgQueue/v1"

	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewGreeterMQ, v1.NewMQ)

type greeterMQ struct {
	mq  *v1.MQ
	log *log.Helper
}

// NewFeedMQ .
func NewGreeterMQ(mq *v1.MQ, logger log.Logger) biz.GreeterMQ {
	return &greeterMQ{
		mq:  mq,
		log: log.NewHelper(logger),
	}
}

func (r *greeterMQ) SndMsg(ctx context.Context, topic string, content string, tag string) (string, error) {
	msg := primitive.NewMessage(topic, []byte(content)).WithTag(tag)
	// 可以用json把想要传递的对象序列化成字节切片，刚好就是[]byte类型
	res, err := r.mq.Producer.SendSync(context.Background(), msg)
	if err != nil {
		fmt.Printf("send message error: %s\n", err)
		return err.Error(), err
	} else {
		fmt.Printf("send message success: result=%s\n", res.String())
		return res.String(), nil
	}
}

func (r *greeterMQ) DealMsg(ctx context.Context, topic string) error {
	selector := consumer.MessageSelector{
		Type:       consumer.TAG,
		Expression: "TagA || TagC",
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

	return nil
}

func (r *greeterMQ) ProducerStart(ctx context.Context) error {
	err := r.mq.Producer.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		return err
	}
	return nil
}

func (r *greeterMQ) ConsumerStart(ctx context.Context) error {
	err := r.mq.PushConsumer.Start()
	if err != nil {
		fmt.Printf("start consumer error: %s", err.Error())
		return err
	}
	return nil
}
