package env

import (
	"os"
	"testing"
)

func TestGetEnvOrElse(t *testing.T) {
	_ = os.Setenv("STRING", "OK")
	type args struct {
		env          string
		defaultValue string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "know string", args: args{env: "STRING", defaultValue: "KO"}, want: "OK"},
		{name: "unknown string", args: args{env: "UNKNOWN_STRING", defaultValue: "KO"}, want: "KO"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnvOrElse(tt.args.env, tt.args.defaultValue); got != tt.want {
				t.Errorf("GetEnvOrElse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetIntEnvOrElse(t *testing.T) {
	_ = os.Setenv("INT", "42")
	_ = os.Setenv("UN_PARSABLE_INT", "4.6692")
	type args struct {
		env          string
		defaultValue int
	}
	tests := []struct {
		name      string
		args      args
		wantValue int
		wantErr   bool
	}{
		{name: "know int", args: args{env: "INT", defaultValue: 0}, wantValue: 42, wantErr: false},
		{name: "unknown int", args: args{env: "UNKNOWN_INT", defaultValue: 0}, wantValue: 0, wantErr: false},
		{name: "un parsable int", args: args{env: "UN_PARSABLE_INT", defaultValue: 0}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, err := GetIntEnvOrElse(tt.args.env, tt.args.defaultValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIntEnvOrElse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotValue != tt.wantValue {
				t.Errorf("GetIntEnvOrElse() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestGetFloatEnvOrElse(t *testing.T) {
	_ = os.Setenv("FLOAT", "4.6692")
	_ = os.Setenv("UN_PARSABLE_FLOAT", "4,6692xc")

	type args struct {
		env          string
		defaultValue float64
	}
	tests := []struct {
		name      string
		args      args
		wantValue float64
		wantErr   bool
	}{
		{name: "know float", args: args{env: "FLOAT", defaultValue: 0.1}, wantValue: 4.6692, wantErr: false},
		{name: "unknown float", args: args{env: "UNKNOWN_FLOAT", defaultValue: 0.1}, wantValue: 0.1, wantErr: false},
		{name: "un parsable float", args: args{env: "UN_PARSABLE_FLOAT", defaultValue: 0.1}, wantValue: 0.1, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, err := GetFloatEnvOrElse(tt.args.env, tt.args.defaultValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFloatEnvOrElse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotValue != tt.wantValue {
				t.Errorf("GetFloatEnvOrElse() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestGetBoolEnvOrElse(t *testing.T) {
	_ = os.Setenv("BOOL_SHORT", "t")
	_ = os.Setenv("BOOL_LOWER", "true")
	_ = os.Setenv("BOOL_UPPER", "TRUE")
	_ = os.Setenv("UN_PARSABLE_BOOL", "pouet")

	type args struct {
		env          string
		defaultValue bool
	}
	tests := []struct {
		name      string
		args      args
		wantValue bool
		wantErr   bool
	}{
		{name: "know short bool", args: args{env: "BOOL_SHORT", defaultValue: false}, wantValue: true, wantErr: false},
		{name: "know lower bool", args: args{env: "BOOL_LOWER", defaultValue: false}, wantValue: true, wantErr: false},
		{name: "know upper bool", args: args{env: "BOOL_UPPER", defaultValue: false}, wantValue: true, wantErr: false},
		{name: "unknown bool", args: args{env: "UNKNOWN_BOOL", defaultValue: true}, wantValue: true, wantErr: false},
		{name: "un parsable bool", args: args{env: "UN_PARSABLE_BOOL", defaultValue: true}, wantValue: true, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, err := GetBoolEnvOrElse(tt.args.env, tt.args.defaultValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBoolEnvOrElse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotValue != tt.wantValue {
				t.Errorf("GetBoolEnvOrElse() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}
