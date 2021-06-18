package dao

import (
	"context"
	"fmt"
	"github.com/1211ciel/shadow-im/calendar/conf"
	"go.etcd.io/etcd/clientv3"
	"log"
	"testing"
)

func TestEtcdCli(t *testing.T) {
	members, err := getEtcd().Discover(context.Background(), "job.rpc")
	if err != nil {
		t.Fatal(err)
	}
	for _, member := range members {
		log.Println(string(member.Key), string(member.Value))
	}

}
func TestEtcdWatcher(t *testing.T) {
	getEtcd().Watch(
		context.Background(),
		"ciel",
		func(event *clientv3.Event) {
			fmt.Println("put", string(event.Kv.Key), string(event.Kv.Value))
		},
		func(event *clientv3.Event) {
			fmt.Println("del ", string(event.Kv.Key), string(event.Kv.Value))
		},
	)
}
func TestEtcd_PutAlive(t *testing.T) {
	err := getEtcd().Register(context.Background(), "ciel/123", "ok", 5)
	if err != nil {
		t.Fatal(err)
	}
	select {}
}
func TestEtcd_Put(t *testing.T) {
	_, err := getEtcd().Put(context.Background(), "ciel/123", "233")
	if err != nil {
		t.Fatal(err)
	}
}

func getEtcd() *Etcd {
	cli, err := NewEtcd(&conf.EtcdCli{
		DialTimeout: 5,
		Endpoints:   []string{"127.0.0.1:2379"},
	})
	if err != nil {
		panic(err)
	}
	return cli
}
