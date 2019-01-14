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
		{name: "unknown", want: UNKNOWN},
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
		{name: "qa", args: struct{ c Stack }{c: 1}, want: "https://qa.libertymedical.fr/assets/mail_assets/logo_couleur_rvb_720_135_qa.png"},
		{name: "preprod", args: struct{ c Stack }{c: 2}, want: "https://preprod.libertymedical.fr/assets/mail_assets/logo_couleur_rvb_720_135_pp.png"},
		{name: "prod", args: struct{ c Stack }{c: 3}, want: "https://www.libertymedical.fr/assets/mail_assets/logo_couleur_rvb_720_135.png"},
		{name: "unknown", args: struct{ c Stack }{c: 4}, want: "https://www.libertymedical.fr/assets/mail_assets/logo_couleur_rvb_720_135.png"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStackLogo(tt.args.c); got != tt.want {
				t.Errorf("GetStackLogo() = %v, want %v", got, tt.want)
			}
		})
	}
}
