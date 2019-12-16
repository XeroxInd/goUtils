package logs

import (
	"fmt"
	"testing"
	"time"
)

// Don't forget to set your SENTRY_DSN environment variable to provided DSN
// ex :
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

func TestInfof(t *testing.T) {
	Infof("test info message : %s, %d", "pouet", 42)
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

func TestFatal(t *testing.T) {
	Fatal("fatal error")
}

func TestFatalf(t *testing.T) {
	Fatalf("fatal error : %s", "kaboom")
}
