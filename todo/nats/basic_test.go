package nats

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"testing"
	"time"
)

func TestNATS(t *testing.T) {
	// natsBasicTest(t)
	natsStreamBasicTest(t)
}

func natsBasicTest(t *testing.T) {
	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		t.Fatal(err)
	}

	// Simple Async Subscriber
	_, err = nc.Subscribe("ORDERS.*", func(m *nats.Msg) {
		t.Logf("Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		t.Fatal(err)
	}

	// Simple Publisher
	err = nc.Publish("ORDERS.a", []byte("Hello World"))
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(1 * time.Second)
}

func natsStreamBasicTest(t *testing.T) {
	var (
		nc  *nats.Conn
		js  nats.JetStreamContext
		err error
	)

	nc, err = nats.Connect(nats.DefaultURL)
	if err != nil {
		t.Fatal(fmt.Errorf("连接：%w", err))
	}

	js, err = nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		t.Fatal(fmt.Errorf("创建 context：%w", err))
	}

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "ORDERS",
		Subjects: []string{"ORDERS.*"},
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = js.Publish("ORDERS.scratch", []byte("hello"))
	if err != nil {
		t.Fatal(fmt.Errorf("发布消息：%w", err))
	}

	_, err = js.Subscribe("ORDERS.*", func(m *nats.Msg) {
		t.Logf("Received a JetStream message: %s\n", string(m.Data))
	})
	if err != nil {
		t.Fatal(fmt.Errorf("订阅消息：%w", err))
	}
}
