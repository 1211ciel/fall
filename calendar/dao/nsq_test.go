package dao

import (
	"github.com/1211ciel/shadow-im/calendar/conf"
	"github.com/nsqio/go-nsq"
	"log"
	"testing"
)

func TestNewNsqPub(t *testing.T) {
	nsqPub := NewNsqPub(&conf.NsqServer{
		Addr: "localhost:4150",
	})
	err := nsqPub.Publish("ciel", []byte("hello"))
	if err != nil {
		t.Fatal(err.Error())
	}
}

type handle struct{}

func (h handle) HandleMessage(message *nsq.Message) error {
	message.Finish()
	log.Println(string(message.Body))
	return nil
}
