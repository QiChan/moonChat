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
	TransProducer
	BroadCastConsumer
)

type MQ struct {
	Producer          rocketmq.Producer
	PushConsumer      rocketmq.PushConsumer
	ProducerOrderly   rocketmq.Producer
	ConsumerOrderly   rocketmq.PushConsumer
	TransProducer     rocketmq.TransactionProducer
	BroadCastConsumer rocketmq.PushConsumer
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
	bc := newBroadCastConsumer(c)
	return &MQ{
		Producer:          pd,
		PushConsumer:      pc,
		ProducerOrderly:   po,
		ConsumerOrderly:   co,
		BroadCastConsumer: bc,
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
		consumer.WithConsumerModel(consumer.BroadCasting),
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

func newBroadCastConsumer(cf *MQ_Config) rocketmq.PushConsumer {
	if (*cf)[BroadCastConsumer] == nil {
		return nil
	}
	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName((*cf)[BroadCastConsumer].GroupName),
		consumer.WithNameServer([]string{(*cf)[BroadCastConsumer].NameSvrAddr}),
		consumer.WithConsumeFromWhere(consumer.ConsumeFromFirstOffset),
		consumer.WithConsumerModel(consumer.BroadCasting),
	)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return c
}

// producer和consumer都是broker的cli
func (r *MQ) StartOneCli(suber Suber) error {
	var err error
	switch suber {
	case Producer:
		err = r.Producer.Start()
	case PushConsumer:
		err = r.PushConsumer.Start()
	case ProducerOrderly:
		err = r.ProducerOrderly.Start()
	case ConsumerOrderly:
		err = r.ConsumerOrderly.Start()
	case TransProducer:
		err = r.TransProducer.Start()
	case BroadCastConsumer:
		err = r.BroadCastConsumer.Start()
	}
	if err != nil {
		fmt.Printf("start %d error: %s", suber, err.Error())
		return err
	}

	return nil
}

func (r *MQ) StartAllCli(config *MQ_Config) error {
	for k, v := range *config {
		if v != nil {
			if err := r.StartOneCli(k); err != nil {
				panic(err)
			}

		}
	}

	return nil
}
