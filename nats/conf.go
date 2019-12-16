package nats

import "github.com/nats-io/stan.go"

type Config struct {
	clusterID     string
	clientID      string
	url           string
	connection    stan.Conn
}