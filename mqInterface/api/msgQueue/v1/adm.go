package v1

import (
	"context"
	"fmt"

	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewMqAdm)
var BrokerAddr = "127.0.0.1:9876"

type MqAdm struct {
	Adm admin.Admin
}

func NewMqAdm() *MqAdm {
	adm, err := admin.NewAdmin(
		admin.WithResolver(primitive.NewPassthroughResolver([]string{BrokerAddr})),
		//admin.WithCredentials()
	)

	if err != nil {
		fmt.Println("NewMqAdm error: ", err)
		panic(err)
	}

	/*
		tmp := new(MqAdm)
		tmp.Adm = adm
	*/

	return &MqAdm{
		Adm: adm,
	}

}

func (a *MqAdm) CreateTopic(newTopic, brokerAddr string) error {
	err := a.Adm.CreateTopic(context.Background(),
		admin.WithTopicCreate(newTopic),
		admin.WithBrokerAddrCreate(brokerAddr),
	)

	if err != nil {
		fmt.Println("CreateTopic error: ", err)
	}

	return err
}

func (a *MqAdm) DeleteTopic(topic, brokerAddr string) error {
	err := a.Adm.DeleteTopic(context.Background(),
		admin.WithTopicDelete(topic),
		admin.WithBrokerAddrDelete(brokerAddr),
	)

	if err != nil {
		fmt.Println("DeleteTopic error: ", err)
	}

	return err
}
