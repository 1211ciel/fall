package dao

import (
	"fmt"
	"github.com/1211ciel/fall/calendar/conf"
	"github.com/nsqio/go-nsq"
	"log"
	"testing"
)

func TestNewNsqPub(t *testing.T) {
	nsqPub := NewNsqPub(&conf.NsqServer{
		Addr: "localhost:4150",
	})
	err := nsqPub.Publish("ciel", []byte("hello2"))
	if err != nil {
		t.Fatal(err.Error())
	}
}
func TestNewNSQSub(t *testing.T) {
	go func() {
		NewNSQSub(&conf.NsqCli{
			Addresses: []string{"localhost:4150"},
		}, "ciel", "first", &handle{})
	}()
	select {}
}

type handle struct{}

// 消息处理
func (h handle) HandleMessage(message *nsq.Message) error {
	message.Finish()
	log.Println(string(message.Body))
	return nil
}

// 2021-7-2 1
func TestNewNSQSub2021_7_2(t *testing.T) {
	pub := NewNsqPub(&conf.NsqServer{
		Addr: "localhost:4150",
	})
	if err := pub.Publish("ciel", []byte("hello")); err != nil {
		t.Fatal(err.Error())
	}
}
func TestNSQSub2021_7_2(t *testing.T) {
	go func() {
		NewNSQSub(&conf.NsqCli{Addresses: []string{"localhost:4150"}}, "ciel", "fist", &nsqHandle2021_7_2{})
	}()
	select {}
}

type nsqHandle2021_7_2 struct {
}

func (nsqHandle2021_7_2) HandleMessage(msg *nsq.Message) error {
	msg.Finish()
	log.Println(string(msg.Body))
	return nil
}

// 2021-7-2 2
func TestNSQPub21722(t *testing.T) {
	pub := NewNsqPub(&conf.NsqServer{Addr: "localhost:4150"})
	if err := pub.Publish("ciel", []byte("hello3")); err != nil {
		t.Fatal(err.Error())
	}
}

type handleNSQSub21722 struct {
}

func (h handleNSQSub21722) HandleMessage(message *nsq.Message) error {
	message.Finish()
	log.Println(string(message.Body))
	return nil
}

func TestNSQSub21722(t *testing.T) {
	go func() {
		NewNSQSub(&conf.NsqCli{Addresses: []string{"localhost:4150"}}, "ciel", "first", handleNSQSub21722{})
	}()
	select {}
}

func TestNSQPub3(t *testing.T) {
	pub := NewNsqPub(&conf.NsqServer{Addr: "localhost:4150"})
	err := pub.Publish("ciel", []byte("hello 7 3"))
	if err != nil {
		t.Fatal(err)
	}
}

type nsqSubHandler struct {
}

func (n nsqSubHandler) HandleMessage(m *nsq.Message) error {
	m.Finish()
	fmt.Println(string(m.Body))
	return nil
}

func TestNSQSub3(t *testing.T) {
	go func() {
		NewNSQSub(&conf.NsqCli{Addresses: []string{"localhost:4150"}}, "ciel", "1", nsqSubHandler{})
	}()
	select {}
}
