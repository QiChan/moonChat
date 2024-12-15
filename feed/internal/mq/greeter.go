package mq

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"moonChat/feed/internal/biz"
	v1 "moonChat/mqInterface/api/msgQueue/v1"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewGreeterMQ)

type MQ struct {
	Producer_1      rocketmq.Producer
	Producer_2      rocketmq.Producer
	Producer_3      rocketmq.Producer
	Producer_4      rocketmq.Producer
	Producer_5      rocketmq.Producer
	Producer_6      rocketmq.Producer
	ProducerOrderly rocketmq.Producer
	TransProducer_1 rocketmq.TransactionProducer
}

type greeterMQ struct {
	mq  *MQ
	log *log.Helper
}

func NewMQ() *MQ {
	tmp := &MQ{
		Producer_1:      v1.NewProducer("127.0.0.1:9876", "feedProducer_test"),
		Producer_2:      v1.NewProducer("127.0.0.1:9876", "feedProducer_test"),
		Producer_3:      v1.NewProducer("127.0.0.1:9876", "feedProducer_test"),
		Producer_4:      v1.NewProducer("127.0.0.1:9876", "feedProducer_test"),
		Producer_5:      v1.NewProducer("127.0.0.1:9876", "feedProducer_test"),
		Producer_6:      v1.NewProducer("127.0.0.1:9876", "feedProducer_test"),
		ProducerOrderly: v1.NewOrderlyProducer("127.0.0.1:9876", "feedOrderlyProducer_test"),
		TransProducer_1: v1.NewTransProducer("127.0.0.1:9876", NewDemoListener()),
	}
	return tmp
}

// NewFeedMQ .
func NewGreeterMQ(logger log.Logger) biz.GreeterMQ {
	tmp := &greeterMQ{
		mq:  NewMQ(),
		log: log.NewHelper(logger),
	}

	return tmp
}

func (r *greeterMQ) SndMsgSync(ctx context.Context, topic string, content string, tag string) (string, error) {
	msg := primitive.NewMessage(topic, []byte(content)).WithTag(tag)
	// 可以用json把想要传递的对象序列化成字节切片，刚好就是[]byte类型
	res, err := r.mq.Producer_1.SendSync(context.Background(), msg)
	if err != nil {
		fmt.Printf("send message error: %s\n", err)
		return err.Error(), err
	} else {
		fmt.Printf("send message success: result=%s\n", res.String())
		return res.String(), nil
	}
}

func (r *greeterMQ) SndMsgAsync(ctx context.Context, topic string, content string, tag string) error {
	err := r.mq.Producer_2.SendAsync(ctx, func(ctx context.Context, result *primitive.SendResult, e error) {
		if e != nil {
			fmt.Println("snd async message error: ", e.Error())
		} else {
			fmt.Println("snd async message success: ", result.String())
		}
	}, primitive.NewMessage(topic, []byte(content)).WithTag(tag))

	if err != nil {
		fmt.Println("snd async message error: ", err.Error())
		return err
	}
	return nil
}

func (r *greeterMQ) SndMsgOneWay(ctx context.Context, topic string, content string, tag string) error {
	msg := primitive.NewMessage(topic, []byte(content)).WithTag(tag)
	err := r.mq.Producer_3.SendOneWay(ctx, msg)
	if err != nil {
		fmt.Println("send oneway message error: ", err)
		return err
	} else {
		fmt.Println("send oneway message success")
		return nil
	}
}

func (r *greeterMQ) SndMsgBatch(ctx context.Context, topic string, content string, tag string) error {
	msgs := []*primitive.Message{}
	for i := 0; i < 10; i++ {
		msgs = append(msgs, primitive.NewMessage(topic, []byte(content+strconv.Itoa(i))).WithTag(tag))
	}
	res, err := r.mq.Producer_4.SendSync(ctx, msgs...)
	if err != nil {
		fmt.Printf("send batch message error: %s\n", err)
	} else {
		fmt.Printf("send batch message success: result=%s\n", res.String())
	}

	return nil
}

func (r *greeterMQ) SndMsgOrderly(ctx context.Context, topic string, content string, tag string) error {
	// 能保证同一orderId下的消息是顺序的
	for i := 0; i < 3; i++ {
		orderId := strconv.Itoa(i)
		for j := 1; j < 5; j++ {
			msg := &primitive.Message{
				Topic: topic,
				Body:  []byte("订单: " + orderId + content + " 步骤: " + strconv.Itoa(j)),
			}
			msg.WithShardingKey(orderId).WithTag(tag) // *关键 用于分片
			_, err := r.mq.ProducerOrderly.SendSync(context.Background(), msg)
			if err != nil {
				fmt.Printf("send message orderly err: %s", err)
				continue
			}
		}
	}
	return nil
}

func (r *greeterMQ) SndMsgDelay(ctx context.Context, topic string, content string, tag string) error {
	for i := 0; i < 10; i++ {
		msg := primitive.NewMessage(topic, []byte(content))
		msg.WithDelayTimeLevel(3).WithTag(tag)
		res, err := r.mq.Producer_5.SendSync(context.Background(), msg)

		if err != nil {
			fmt.Printf("send delay message error: %s\n", err)
		} else {
			fmt.Printf("send delay message success: result=%s\n", res.String())
		}
	}
	return nil
}

func (r *greeterMQ) SndMsgDelayAnyTime(ctx context.Context, topic string, content string, tag string, delayInterval int64) error {
	for i := 0; i < 10; i++ {
		msg := primitive.NewMessage(topic, []byte(content+strconv.Itoa(i)+" timestamp:"+strconv.FormatInt(time.Now().Unix(), 10)))
		msg.WithDelayTimestamp(time.Now().Add(time.Duration(delayInterval) * time.Second)).WithTag(tag)
		res, err := r.mq.Producer_6.SendSync(context.Background(), msg)

		if err != nil {
			fmt.Printf("send delay any time message error: %s\n", err)
		} else {
			fmt.Printf("send delay any time message success: result=%s\n", res.String())
		}
	}
	return nil
}

func (r *greeterMQ) SndMsgTrans(ctx context.Context, topic string, content string, tag string) error {
	for i := 0; i < 10; i++ {
		res, err := r.mq.TransProducer_1.SendMessageInTransaction(context.Background(),
			primitive.NewMessage(topic, []byte(content+strconv.Itoa(i))).WithTag(tag))

		if err != nil {
			fmt.Printf("send trans message error: %s\n", err)
			return err
		} else {
			fmt.Printf("send trans message success: result=%s\n", res.String())
		}
	}
	return nil
}

func (r *greeterMQ) ClientsStart(ctx context.Context) error {
	var err error
	err = r.mq.Producer_1.Start()
	err = r.mq.Producer_2.Start()
	err = r.mq.Producer_3.Start()
	err = r.mq.Producer_4.Start()
	err = r.mq.Producer_5.Start()
	err = r.mq.Producer_6.Start()
	err = r.mq.ProducerOrderly.Start()
	err = r.mq.TransProducer_1.Start()

	if err != nil {
		fmt.Println("start producer error: ", err.Error())
		return err
	}

	return nil
}
