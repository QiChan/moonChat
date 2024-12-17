package mq

import (
	"context"
	"fmt"
	"time"

	v1 "moonChat/mqInterface/api/msgQueue/v1"
	"moonChat/userCenter/internal/biz"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewUserMQ, NewMQ)

type MQ struct {
	PushConsumer_1      rocketmq.PushConsumer
	PushConsumer_2      rocketmq.PushConsumer
	ConsumerOrderly_1   rocketmq.PushConsumer
	ConsumerOrderly_2   rocketmq.PushConsumer
	BroadCastConsumer_1 rocketmq.PushConsumer
	BroadCastConsumer_2 rocketmq.PushConsumer
}

var Mq_Suber_arr = []rocketmq.PushConsumer{}

func NewMQ() *MQ {
	tmp := &MQ{
		PushConsumer_1:      v1.NewConsumer("127.0.0.1:9876", "userCenterConsumer_test", "pushConsumer1"),
		PushConsumer_2:      v1.NewConsumer("127.0.0.1:9876", "userCenterConsumer_test", "pushConsumer2"),
		ConsumerOrderly_1:   v1.NewOrderlyConsumer("127.0.0.1:9876", "userCenterOrderlyConsumer_orderlyTest", "orderlyConsumer1"),
		ConsumerOrderly_2:   v1.NewOrderlyConsumer("127.0.0.1:9876", "userCenterOrderlyConsumer_orderlyTest", "orderlyConsumer2"),
		BroadCastConsumer_1: v1.NewBroadCastConsumer("127.0.0.1:9876", "userCenterBroadcastConsumer_broadcastTest", "broadcastConsumer1"),
		BroadCastConsumer_2: v1.NewBroadCastConsumer("127.0.0.1:9876", "userCenterBroadcastConsumer_broadcastTest", "broadcastConsumer2"),
	}

	Mq_Suber_arr = append(Mq_Suber_arr, tmp.PushConsumer_1, tmp.PushConsumer_2, tmp.ConsumerOrderly_1, tmp.ConsumerOrderly_2, tmp.BroadCastConsumer_1, tmp.BroadCastConsumer_2)

	return tmp
}

type userMQ struct {
	mq  *MQ
	log *log.Helper
}

// NewGreeterRepo .
func NewUserMQ(mq *MQ, logger log.Logger) biz.UserMQ {
	return &userMQ{
		mq:  mq,
		log: log.NewHelper(logger),
	}
}

func (r *userMQ) DealMsg(ctx context.Context, topic string) error {
	selector := consumer.MessageSelector{
		Type:       consumer.TAG,
		Expression: "",
	}
	err := r.mq.PushConsumer_1.Subscribe(topic, selector, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
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
	err := r.mq.PushConsumer_1.Subscribe(topic, selector, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, msg := range msgs {
			t := time.Now().UnixNano()/int64(time.Millisecond) - msg.BornTimestamp
			fmt.Printf("subscribe 1 Receive message[msg=%s] %d ms later\n", msg, t)
		}

		return consumer.ConsumeSuccess, nil
	})

	err = r.mq.PushConsumer_2.Subscribe(topic, selector, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, msg := range msgs {
			t := time.Now().UnixNano()/int64(time.Millisecond) - msg.BornTimestamp
			fmt.Printf("subscribe 2 Receive message[msg=%s] %d ms later\n", msg, t)
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
	err := r.mq.ConsumerOrderly_1.Subscribe(topic, selector, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		fmt.Printf("subscribe orderly 1 callback: %v \n", msgs)
		return consumer.ConsumeSuccess, nil
	})

	err = r.mq.ConsumerOrderly_2.Subscribe(topic, selector, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		fmt.Printf("subscribe orderly 2 callback: %v \n", msgs)
		return consumer.ConsumeSuccess, nil
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}

func (r *userMQ) DealMsgBroadCasting(ctx context.Context, topic string) error {
	err := r.mq.BroadCastConsumer_1.Subscribe(topic, consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		fmt.Printf("subscribe broadcasting 1 callback: %v \n", msgs)
		return consumer.ConsumeSuccess, nil
	})

	err = r.mq.BroadCastConsumer_2.Subscribe(topic, consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		fmt.Printf("subscribe broadcasting 2 callback: %v \n", msgs)
		return consumer.ConsumeSuccess, nil
	})

	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}

func (r *userMQ) ClientsStart(ctx context.Context) error {
	for idx, cli := range Mq_Suber_arr {
		if err := cli.Start(); err != nil {
			fmt.Println("userCenter consumer start err, idx: ", idx, "err: ", err)
			//return err
		}
	}

	return nil
}
