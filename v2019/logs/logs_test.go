package logs

import (
	"fmt"
	"testing"
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
