package nats

import "github.com/nats-io/go-nats-streaming"

type Config struct {
	clusterID     string
	clientID      string
	url           string
	connection    stan.Conn
}