package mq

import (
	"context"
	"fmt"
	"time"

	v1 "moonChat/mqInterface/api/msgQueue/v1"
	"moonChat/userCenter/internal/biz"

	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewUserMQ)

type userMQ struct {
	mq  *v1.MQ
	log *log.Helper
}

// NewGreeterRepo .
func NewUserMQ(logger log.Logger) biz.UserMQ {
	return &userMQ{
		mq: v1.NewMQ(&v1.MQ_Config{
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
		}),
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

	return nil
}

func (r *userMQ) DealMsgWithTimeElapse(ctx context.Context, topic string) error {
	selector := consumer.MessageSelector{
		Type:       consumer.TAG,
		Expression: "",
	}
	err := r.mq.PushConsumer.Subscribe(topic, selector, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, msg := range msgs {
			t := time.Now().UnixNano()/int64(time.Millisecond) - msg.BornTimestamp
			fmt.Printf("Receive message[msgId=%s] %d ms later\n", msg.MsgId, t)
		}

		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func (r *userMQ) DealMsgOrderly(ctx context.Context, topic string) error {
	selector := consumer.MessageSelector{
		Type:       consumer.TAG,
		Expression: "",
	}
	err := r.mq.ConsumerOrderly.Subscribe(topic, selector, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		/*
			orderlyCtx, _ := primitive.GetOrderlyCtx(ctx)
			fmt.Printf("orderly context: %v\n", orderlyCtx)
		*/
		fmt.Printf("subscribe orderly callback: %v \n", msgs)
		return consumer.ConsumeSuccess, nil
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}

func (r *userMQ) ProducerStart(ctx context.Context) error {
	err := r.mq.Producer.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		return err
	}

	err = r.mq.ProducerOrderly.Start()
	if err != nil {
		fmt.Printf("start producerOrderly error: %s", err.Error())
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

	err = r.mq.ConsumerOrderly.Start()
	if err != nil {
		fmt.Printf("start consumerOrderly error: %s", err.Error())
		return err
	}

	return nil
}
