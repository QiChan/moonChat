package v1

import (
	"fmt"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

// reference
type MQ struct {
	Producer          rocketmq.Producer
	PushConsumer      rocketmq.PushConsumer
	ProducerOrderly   rocketmq.Producer
	ConsumerOrderly   rocketmq.PushConsumer
	TransProducer     rocketmq.TransactionProducer
	BroadCastConsumer rocketmq.PushConsumer
}

type MQ_Suber_Config struct {
	GroupName   string
	NameSvrAddr string
}

/*
func NewMQ(c *MQ_Config) *MQ {
	pd := newProducer(c)
	pc := newConsumer(c)
	po := newOrderlyProducer(c)
	co := newOrderlyConsumer(c)
	bc := newBroadCastConsumer(c)
	return &MQ{
		Producer:          pd,
		PushConsumer:      pc,
		ProducerOrderly:   po,
		ConsumerOrderly:   co,
		BroadCastConsumer: bc,
	}
}
*/

func NewProducer(nameSvrAddr string, groupName string) rocketmq.Producer {
	p, err := rocketmq.NewProducer(
		producer.WithNameServer([]string{nameSvrAddr}),
		producer.WithRetry(2),
		producer.WithGroupName(groupName),
	)
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		return nil
	}

	return p
}

func NewConsumer(nameSvrAddr string, groupName string) rocketmq.PushConsumer {
	c, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{nameSvrAddr}),
		consumer.WithGroupName(groupName),
		consumer.WithConsumerModel(consumer.BroadCasting),
	)
	if err != nil {
		fmt.Println("create consumer err", err.Error())
		return nil
	}

	return c
}

func NewOrderlyProducer(nameSvrAddr string, groupName string) rocketmq.Producer {
	c, err := rocketmq.NewProducer(
		producer.WithNameServer([]string{nameSvrAddr}),
		producer.WithGroupName(groupName),
		producer.WithRetry(2),
		producer.WithQueueSelector(producer.NewHashQueueSelector()),
	)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return c
}

func NewOrderlyConsumer(nameSvrAddr string, groupName string) rocketmq.PushConsumer {
	c, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{nameSvrAddr}),
		consumer.WithGroupName(groupName),
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

func NewBroadCastConsumer(nameSvrAddr string, groupName string) rocketmq.PushConsumer {
	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName(groupName),
		consumer.WithNameServer([]string{nameSvrAddr}),
		consumer.WithConsumeFromWhere(consumer.ConsumeFromFirstOffset),
		consumer.WithConsumerModel(consumer.BroadCasting),
	)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return c
}

func NewTransProducer(nameSvrAddr string, listener primitive.TransactionListener) rocketmq.TransactionProducer {
	tp, err := rocketmq.NewTransactionProducer(
		listener,
		producer.WithNameServer([]string{nameSvrAddr}),
		producer.WithRetry(2),
	)
	if err != nil {
		fmt.Println("create trans producer err: ", err.Error())
		return nil
	}

	return tp
}
