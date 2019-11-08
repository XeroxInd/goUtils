package nats

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"os"
)

func (c *Config) Load() {
	var err error
	if c.clusterID == "" {
		c.clusterID = os.Getenv("NATS_CLUSTER_ID", )
	}
	if c.clientID == "" {
		c.clientID = os.Getenv("NATS_CLIENT_ID", )
	}
	if c.url == "" {
		c.url = os.Getenv("NATS_URL")
	}
	if len(c.clientID) == 0 || len(c.clusterID) == 0 || len(c.url) == 0 {
		log.Fatal("Some env variable aren't defined, please check it, NATS_CLUSTER_ID, NATS_CLIENT_ID and NATS_URL must be set")
	}
	c.connection, err = stan.Connect(
		c.clusterID,
		c.clientID,
		stan.NatsURL(c.url),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("connection to NATS server lost, reason: %v\nThen I quit: \"Au revoir...\"", reason)
		}))
	if err != nil {
		log.Fatalf("unable to connect to NATS Streaming server (%s): %s", c.url, err.Error())
	}
	log.Printf("connected to NATS Streaming server (%s), with clientID : %s\n", c.url, c.clientID)
}

func (c *Config) Subscribe(subject string, handler stan.MsgHandler, opts ...stan.SubscriptionOption) (sub stan.Subscription, err error) {
	log.Printf("Subscribing to subject '%s'\n", subject)
	sub, err = c.connection.Subscribe(subject, handler, opts...)
	if err != nil {
		log.Print(fmt.Errorf("unable to subscribe: %s", err.Error()), "nats")
	}
	return
}

func (c *Config) Publish(subject string, message []byte) (err error) {
	return c.connection.Publish(subject, message)
}

func (c *Config) ShutDown(subs ...stan.Subscription) (err error){
	log.Print("closing NATS connection...")
	for _, s := range subs {
		err = s.Close()
		if err != nil {
			log.Print(fmt.Errorf("error when close NATS subscription: %s", err.Error()), "nats")
		}
	}
	err = c.connection.Close()
	if err != nil {
		log.Print(fmt.Errorf("unable to close NATS connection : %s", err.Error()), "nats")
	}
	return
}
