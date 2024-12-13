package v1

import (
	"fmt"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

type Suber int

const (
	Producer Suber = iota
	PushConsumer
	ProducerOrderly
	ConsumerOrderly
)

type MQ struct {
	Producer        rocketmq.Producer
	PushConsumer    rocketmq.PushConsumer
	ProducerOrderly rocketmq.Producer
	ConsumerOrderly rocketmq.PushConsumer
}

type MQ_Config map[Suber]*MQ_Suber_Config

type MQ_Suber_Config struct {
	GroupName   string
	NameSvrAddr string
}

func NewMQ(c *MQ_Config) *MQ {
	pd := newProducer(c)
	pc := newConsumer(c)
	po := newOrderlyProducer(c)
	co := newOrderlyConsumer(c)
	return &MQ{
		Producer:        pd,
		PushConsumer:    pc,
		ProducerOrderly: po,
		ConsumerOrderly: co,
	}
}

func newProducer(cf *MQ_Config) rocketmq.Producer {
	if (*cf)[Producer] == nil {
		return nil
	}
	p, err := rocketmq.NewProducer(
		//producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
		producer.WithNameServer([]string{(*cf)[Producer].NameSvrAddr}),
		producer.WithRetry(2),
		producer.WithGroupName((*cf)[Producer].GroupName),
	)
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		return nil
	}

	return p
}

func newConsumer(cf *MQ_Config) rocketmq.PushConsumer {
	if (*cf)[PushConsumer] == nil {
		return nil
	}
	c, err := rocketmq.NewPushConsumer(
		//consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
		consumer.WithNameServer([]string{(*cf)[PushConsumer].NameSvrAddr}),
		consumer.WithGroupName((*cf)[PushConsumer].GroupName),
	)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return c
}

func newOrderlyProducer(cf *MQ_Config) rocketmq.Producer {
	if (*cf)[ProducerOrderly] == nil {
		return nil
	}
	c, err := rocketmq.NewProducer(
		producer.WithNameServer([]string{(*cf)[ProducerOrderly].NameSvrAddr}),
		producer.WithGroupName((*cf)[ProducerOrderly].GroupName),
		producer.WithRetry(2),
		producer.WithQueueSelector(producer.NewHashQueueSelector()),
	)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return c
}

func newOrderlyConsumer(cf *MQ_Config) rocketmq.PushConsumer {
	if (*cf)[ConsumerOrderly] == nil {
		return nil
	}
	c, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{(*cf)[ConsumerOrderly].NameSvrAddr}),
		consumer.WithGroupName((*cf)[ConsumerOrderly].GroupName),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithConsumeFromWhere(consumer.ConsumeFromFirstOffset),
		consumer.WithConsumerOrder(true),
	)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return c
}
