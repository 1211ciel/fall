package dao

import (
	"github.com/1211ciel/fall/calendar/conf"
	"github.com/nsqio/go-nsq"
	"time"
)

func NewNsqPub(c *conf.NsqServer) *nsq.Producer {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(c.Addr, config)
	if err != nil {
		panic(err)
	}
	return producer
}

func NewNSQSub(c *conf.NqsCli, topic, channel string, handle nsq.Handler) *nsq.Consumer {
	config := nsq.NewConfig()
	config.LookupdPollInterval = 15 * time.Second
	con, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		panic(err)
	}
	con.AddHandler(handle)
	err = con.ConnectToNSQDs(c.Addresses)
	if err != nil {
		panic(err)
	}
	return con
}
