package dao

import (
	"context"
	"github.com/1211ciel/fall/calendar/conf"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"log"
	"time"
)

type Etcd struct {
	cli *clientv3.Client
}

func NewEtcd(c *conf.EtcdCli) (*Etcd, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   c.Endpoints,
		DialTimeout: time.Second * time.Duration(c.DialTimeout),
	})
	if err != nil {
		return nil, err
	}
	return &Etcd{cli: cli}, nil
}
func (e *Etcd) Close() {
	err := e.cli.Close()
	if err != nil {
		log.Println(err)
	}
}

// WatchAction 监听的动作
type WatchAction func(*clientv3.Event)

// Watch 服务监听
func (e *Etcd) Watch(ctx context.Context, keyPrefix string, putAction, delAction WatchAction) {
	log.Println("亲爱的ciel 我已经开始监听啦~")
	watch := e.cli.Watch(ctx, keyPrefix, clientv3.WithPrefix())
	for response := range watch {
		for _, event := range response.Events {
			switch event.Type {
			case mvccpb.PUT:
				putAction(event)
			case mvccpb.DELETE:
				delAction(event)
			}
		}
	}
}

func (e *Etcd) Put(ctx context.Context, key, val string) (*clientv3.PutResponse, error) {
	return e.cli.Put(ctx, key, val)
}

// Discover  服务发现
func (e *Etcd) Discover(ctx context.Context, keyPrefix string) ([]*mvccpb.KeyValue, error) {
	get, err := e.cli.Get(ctx, keyPrefix, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	return get.Kvs, nil
}

// Register 服务注册
func (e *Etcd) Register(ctx context.Context, k, v string, ttl int64) error {
	grant, err := e.cli.Grant(ctx, ttl)
	if err != nil {
		return err
	}
	_, err = e.cli.Put(ctx, k, v, clientv3.WithLease(grant.ID))
	if err != nil {
		return err
	}
	keepAlive, err := e.cli.KeepAlive(ctx, grant.ID)
	if err != nil {
		return err
	}
	log.Println(keepAlive)
	//for {
	//	ka := <-keepAlive
	//	log.Println(ka.TTL)
	//}
	return nil
}
