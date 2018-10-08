package s3

import "github.com/minio/minio-go"

type Config struct {
	EndPoint    string
	AccessKeyID string
	SecretKey   string
	UseSSL      bool
	AesKeyPath  string
	AesKeyValue []byte
	Client      *minio.Client
}

type S3Object struct {
	bucket string
	object string
}
