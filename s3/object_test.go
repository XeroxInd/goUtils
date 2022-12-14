package s3

import (
	"crypto/rand"
	"log"
	"os"
	"testing"
)

var c Config
var randValue = make([]byte, 1024)

func init() {
	c = Config{}
	_ = os.Setenv("S3_ENDPOINT", "s3qa.libmed.fr")
	_ = os.Setenv("S3_ACCESS_ID", "15FHJF9E8RH4C0TXH4IT")
	_ = os.Setenv("S3_SECRET_KEY", "a7/4e0qjIl+IPhYxhI2J0m5m+eWhPuoQFc9wr9qk")
	_ = os.Setenv("S3_SSL", "true")
	_ = os.Setenv("S3_AES_KEY_PATH", "key.txt")
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

func TestExists(t *testing.T) {
	o := S3Object{"known", "known"}
	exists, err := c.ObjectExists(o, true)
	if err != nil {
		t.Errorf("%s/%s: %v\n", o.Bucket, o.Object, err.Error())
	} else if exists == false {
		t.Error("should be found and is not")
	} else {
		t.Log("ok")
	}
	exists, err = c.ObjectExists(S3Object{"unknown", "unknown"}, false)
	if err != nil {
		t.Errorf(err.Error())
	} else if exists == true {
		t.Error("should not be found and is")
	} else {
		t.Log("ok")
	}
	exists, err = c.ObjectExists(S3Object{"wololo", "wololo"}, false)
	if err != nil {
		t.Log("ok")
	} else if exists {
		t.Error("should be an error and is not")
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

func TestConfig_Stat(t *testing.T) {
	know := S3Object{"known", "known"}
	s, err := c.Stat(know)
	if err != nil {
		t.Errorf("stat must not return an error, %s", err.Error())
	}
	for k, v := range s.Metadata {
		log.Printf("meta : %s, value : %s\n", k, v)
	}

	meta := map[string]string{
		"string1": "value1",
		"int1":    "42",
	}
	err = c.SetMeta(know, meta, true)
	if err != nil {
		t.Errorf("must not return an error, %s", err.Error())
	}
	s2, err := c.Stat(know)
	if err != nil {
		t.Errorf("stat must not return an error, %s", err.Error())
	}
	log.Print("==== Meta ====")
	for k, v := range s2.Metadata {
		log.Printf("meta : %s, value : %s\n", k, v)
	}

	meta2 := map[string]string{
		"string2": "value2",
		"int2":    "43",
	}
	err = c.SetMeta(know, meta2, true)
	if err != nil {
		t.Errorf("must not return an error, %s", err.Error())
	}
	s3, err := c.Stat(know)
	if err != nil {
		t.Errorf("stat must not return an error, %s", err.Error())
	}
	log.Print("==== Meta ====")
	for k, v := range s3.Metadata {
		log.Printf("meta : %s, value : %s\n", k, v)
	}
}

func TestSetContractMeta(t *testing.T) {
	meta := map[string]string{
		"Hr_lowerlefty":          "104",
		"Hr_upperrightx":         "277",
		"Contractor_lowerleftx":  "317",
		"Contractor_upperrighty": "175",
		"Hr_lowerleftx":          "92",
		"Contractor_lowerlefty":  "104",
		"Contractor_pagenumber":  "1",
		"Hr_upperrighty":         "175",
		"Contractor_upperrightx": "502",
		"Hr_pagenumber":          "1",
	}
	err := c.SetMeta(S3Object{"residalya", "templates/summary_card_cover.html"}, meta, false)
	if err != nil {
		t.Errorf("stat must not return an error, %s", err.Error())
	}

	err = c.SetMeta(S3Object{"residalya", "templates/summary_card_activity_increase.html"}, meta, false)
	if err != nil {
		t.Errorf("stat must not return an error, %s", err.Error())
	}
}
