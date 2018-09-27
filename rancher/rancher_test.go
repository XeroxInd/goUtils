package rancher

import (
	"reflect"
	"testing"
)

func TestGetStack(t *testing.T) {
	tests := []struct {
		name string
		want Stack
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStackLogo(t *testing.T) {
	type args struct {
		c Stack
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStackLogo(tt.args.c); got != tt.want {
				t.Errorf("GetStackLogo() = %v, want %v", got, tt.want)
			}
		})
	}
}
