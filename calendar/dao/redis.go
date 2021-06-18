package dao

import (
	"context"
	"github.com/1211ciel/fall/calendar/conf"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

func NewRedis(c *conf.Redis) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     c.Idle,
		MaxActive:   c.Active,
		IdleTimeout: time.Second * time.Duration(c.IdleTimeout),
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(c.Network, c.Addr,
				redis.DialConnectTimeout(time.Millisecond*time.Duration(c.DialTimeout)),
				redis.DialReadTimeout(time.Millisecond*time.Duration(c.ReadTimeout)),
				redis.DialWriteTimeout(time.Millisecond*time.Duration(c.WriteTimeout)),
				redis.DialPassword(c.Auth),
			)
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
}

// ListenSubChannels 订阅redis监听
func ListenSubChannels(ctx context.Context, redisServerAddr string, onMessage func(channel string, data []byte) error, channels ...string) error {
	const healthCheckPeriod = time.Minute
	step := 1 // init cli
	c, err := redis.Dial("tcp", redisServerAddr,
		redis.DialReadTimeout(healthCheckPeriod+10*time.Second), // Read timeout on server should be greater than ping period.
		redis.DialWriteTimeout(10*time.Second))
	if err != nil {
		return err
	}
	defer c.Close()

	step = 2 // create pushConn
	psc := redis.PubSubConn{Conn: c}

	step = 3 // subscribe
	if err := psc.Subscribe(redis.Args{}.AddFlat(channels)...); err != nil {
		return err
	}

	step = 4 // create a  listener chan
	done := make(chan error, 1)

	step = 5 // Start a goroutine to receive notifications from the server.
	go func() {
		for {
			switch n := psc.Receive().(type) {
			case error:
				done <- n
				return
			case redis.Message:
				if err := onMessage(n.Channel, n.Data); err != nil {
					done <- err
					return
				}
			case redis.Subscription:
				log.Println()
			}
		}
	}()

	step = 6 // 健康检查
	ticker := time.NewTicker(healthCheckPeriod)
	defer ticker.Stop()
loop:
	for err == nil {
		select {
		case <-ticker.C:
			// Send ping to test health of connection and server. If
			// corresponding pong is not received, then receive on the connection will timeout and the receive goroutine will exit.
			if err = psc.Ping(""); err != nil {
				break loop
			}
		case <-ctx.Done():
			break loop
		case err := <-done: // Return error from the receive goroutine.
			return err
		}
	}

	// Signal the receiving goroutine to exit by unsubscribing from all channels.
	psc.Unsubscribe()

	// Wait for goroutine to complete.
	step++
	return <-done
}
