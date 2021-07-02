package dao

import (
	"github.com/1211ciel/fall/calendar/conf"
	"github.com/nsqio/go-nsq"
	"time"
)

/*
nsq windows 启动
1. 进入bin目录 cd nsq-1.2.0/bin
2. 启动nsqlookup ./nsqlookupd &
3. 启动nsqd ./nsqd --lookupd-tcp-address=127.0.0.1:4160 &
4. 运行 nsqadmin 管理 ./nsqadmin --lookupd-http-address=0.0.0.0:4161 --http-address=127.0.0.1:8761&
*/

// NewNsqPub  创建一个发布者
func NewNsqPub(c *conf.NsqServer) *nsq.Producer {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(c.Addr, config)
	if err != nil {
		panic(err)
	}
	return producer
}

// NewNSQSub 创建一个订阅者
func NewNSQSub(c *conf.NsqCli, topic, channel string, handle nsq.Handler) *nsq.Consumer {
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
