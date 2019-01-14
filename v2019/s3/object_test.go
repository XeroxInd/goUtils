package s3

import (
	"crypto/rand"
	"fmt"
	"github.com/minio/minio-go"
	"log"
	"testing"
)

var c Config
var randValue = make([]byte, 1024)

func init() {
	c = Config{
		EndPoint:    "s3qa.libmed.fr:9001",
		AccessKeyID: "1O1Z2MD6VST99199VEXP",
		SecretKey:   "SBHoFRI533JrT7ygSfu68HliBXz8+5X73H1sAHwR",
		UseSSL:      true,
	}
	err := c.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = rand.Read(randValue)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func BenchmarkConfig_Put(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c.Put(S3Object{Bucket: "test", Object: fmt.Sprintf("test_%d", b.N)}, randValue)
	}
}

func TestConfig_Put(t *testing.T) {
	type fields struct {
		EndPoint    string
		AccessKeyID string
		SecretKey   string
		UseSSL      bool
		AesKeyPath  string
		AesKeyValue []byte
		Client      *minio.Client
	}
	type args struct {
		s3object S3Object
		content  []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		struct {
			name    string
			fields  fields
			args    args
			wantErr bool
		}{name: "1k", fields: struct {
			EndPoint    string
			AccessKeyID string
			SecretKey   string
			UseSSL      bool
			AesKeyPath  string
			AesKeyValue []byte
			Client      *minio.Client
		}{
			EndPoint:    c.EndPoint,
			AccessKeyID: c.AccessKeyID,
			SecretKey:   c.SecretKey,
			UseSSL:      c.UseSSL,
			AesKeyPath:  c.AesKeyPath,
			AesKeyValue: c.AesKeyValue,
			Client:      c.Client},
			args: struct {
				s3object S3Object
				content  []byte
			}{s3object: struct {
				Bucket string
				Object string
			}{Bucket: "test", Object: "test1k.bin"}, content: randValue},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				EndPoint:    tt.fields.EndPoint,
				AccessKeyID: tt.fields.AccessKeyID,
				SecretKey:   tt.fields.SecretKey,
				UseSSL:      tt.fields.UseSSL,
				AesKeyPath:  tt.fields.AesKeyPath,
				AesKeyValue: tt.fields.AesKeyValue,
				Client:      tt.fields.Client,
			}
			if err := c.Put(tt.args.s3object, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("Config.Put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
