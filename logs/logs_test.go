package logs

import (
	"fmt"
	"log"
	"testing"

	"github.com/getsentry/raven-go"
)

func init() {
	err := raven.SetDSN("http://ef51cc527b8c4a0c9de3a29425db9c2d@sentry.libmed.fr/8")
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
	Info("test info message", "info test")
}

func TestError(t *testing.T) {
	Error(fmt.Errorf("test error message"), "error test")
}
