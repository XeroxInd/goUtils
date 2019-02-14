package s3

import (
	"crypto/rand"
	"log"
	"testing"
)

var c Config
var randValue = make([]byte, 1024)

func init() {
	c = Config{
		EndPoint:    "s3qa.libmed.fr",
		AccessKeyID: "15FHJF9E8RH4C0TXH4IT",
		SecretKey:   "a7/4e0qjIl+IPhYxhI2J0m5m+eWhPuoQFc9wr9qk",
		UseSSL:      true,
	}
	err := c.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = rand.Read(randValue)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func TestStatUnknown(t *testing.T) {
	_, err := c.Stat(S3Object{"unknown", "unknown"})
	if err != nil {
		t.Logf("ok, not found")
	} else {
		t.Errorf("must return an error")
	}
}

func TestPut(t *testing.T) {
	err := c.Put(S3Object{"known", "known"}, randValue)
	if err != nil {
		t.Errorf(err.Error())
	} else {
		t.Logf("ok")
	}
}

func TestStatKnown(t *testing.T) {
	_, err := c.Stat(S3Object{"known", "known"})
	if err != nil {
		t.Errorf("must not return an error, %s", err.Error())
	} else {
		t.Logf("ok, found")
	}
}
