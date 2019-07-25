package logs

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/getsentry/raven-go"
)

func init() {
	err := raven.SetDSN("http://91a6b7f47e3a4419a2f66e02c87d4592:7ed4e739e65748ffb3d214439167f7a6@sentry.libmed.fr/3")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func TestPrettyPrint(t *testing.T) {
	type Test struct {
		name string
		age  int
	}
	var test = map[string]Test{
		"test1": Test{name: "tintin", age: 32},
		"test2": Test{name: "hadock", age: 54},
	}
	s, _ := PrettyPrint(test)
	fmt.Printf("%v", s)
}

func TestInfo(t *testing.T) {
	Info("test info message")
	time.Sleep(3 * time.Second)
	InfoWithTags("test info message with tags", map[string]string{"tag1": "toto", "tag2": "titi"})
	time.Sleep(3 * time.Second)
}

func TestError(t *testing.T) {
	Error(fmt.Errorf("test error message"))
	time.Sleep(3 * time.Second)
	ErrorWithTags(fmt.Errorf("test error message with tags"), map[string]string{"tag1": "toto", "tag2": "titi"})
	time.Sleep(3 * time.Second)
}
