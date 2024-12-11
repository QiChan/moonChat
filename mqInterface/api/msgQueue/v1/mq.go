package v1

import (
	"fmt"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

type MQ struct {
	Producer     rocketmq.Producer
	PushConsumer rocketmq.PushConsumer
}

func NewMQ() *MQ {
	pd := newProducer()
	pc := newConsumer()
	return &MQ{
		Producer:     pd,
		PushConsumer: pc,
	}
}

func newProducer() rocketmq.Producer {
	p, err := rocketmq.NewProducer(
		//producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
		producer.WithNameServer([]string{"127.0.0.1:9876"}),
		producer.WithRetry(2),
		producer.WithGroupName("testGroup"),
	)
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		return nil
	}

	return p
}

func newConsumer() rocketmq.PushConsumer {
	c, err := rocketmq.NewPushConsumer(
		//consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
		consumer.WithNameServer([]string{"127.0.0.1:9876"}),
		consumer.WithGroupName("testGroup"),
	)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return c
}
