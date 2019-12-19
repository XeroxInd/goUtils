package nats

import (
	"log"
	"testing"

	"github.com/nats-io/stan.go"
)

var c Config
var channel = "test"
var sub stan.Subscription

func init() {
	c = Config{
		clusterID: "libmed",
		clientID:  "test",
		url:       "nats://nats.qa.svc.cluster.local:4222",
	}
	c.Load()
}

func TestSubscribe(t *testing.T) {
	var err error
	sub, err = c.Subscribe(channel, func(m *stan.Msg) {
		log.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		t.Errorf("must not return an error, %s", err.Error())
	} else {
		t.Logf("ok")
	}
}

func TestPublish(t *testing.T) {
	err := c.Publish(channel, []byte("Hello World"))
	if err != nil {
		t.Errorf("must not return an error, %s", err.Error())
	} else {
		t.Logf("ok")
	}
}

func TestClose(t *testing.T) {
	err := c.ShutDown(sub)
	if err != nil {
		t.Errorf("must not return an error, %s", err.Error())
	} else {
		t.Logf("ok")
	}
}
