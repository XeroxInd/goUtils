package logs

import (
	"fmt"
	"testing"
	"time"
)

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
	type args struct {
		msg      string
		category string
	}
	tests := []struct {
		name string
		args args
	}{
		struct {
			name string
			args args
		}{name: "test info", args: struct {
			msg      string
			category string
		}{msg: "test", category: "category test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Info(tt.args.msg, tt.args.category)
		})
	}
	time.Sleep(1 * time.Second)
}
