package env

import (
	"os"
	"testing"
)

func TestGetEnvOrElse(t *testing.T) {
	os.Setenv("TEST", "OK")
	type args struct {
		env          string
		defaultValue string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test", args: struct {
			env          string
			defaultValue string
		}{env: "TEST", defaultValue: "KO"}, want: "OK"},
		{name: "TEST", args: struct {
			env          string
			defaultValue string
		}{env: "UNEXISTENT", defaultValue: "KO"}, want: "KO"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnvOrElse(tt.args.env, tt.args.defaultValue); got != tt.want {
				t.Errorf("GetEnvOrElse() = %v, want %v", got, tt.want)
			}
		})
	}
}
