package s3

import (
	"os"
	"strconv"

	"fmt"
	"io/ioutil"

	"bytes"

	"github.com/minio/minio-go"
	"github.com/minio/minio-go/pkg/encrypt"
	"github.com/pkg/errors"
)

func (c *Config) Load() error {
	//Setup Minio connexion
	//Load ENV variables
	c.EndPoint = os.Getenv("S3_ENDPOINT")
	c.AccessKeyID = os.Getenv("S3_ACCESS_ID")
	c.SecretKey = os.Getenv("S3_SECRET_KEY")
	c.AesKeyPath = os.Getenv("S3_AES_KEY_PATH")
	useSSL := os.Getenv("S3_SSL")
	if len(c.EndPoint) == 0 || len(c.AccessKeyID) == 0 || len(c.SecretKey) == 0 || len(useSSL) == 0 {
		return errors.New("Some env variable aren't defined, please check it, S3_ENDPOINT, S3_ACCESS_ID, S3_SECRET_KEY and S3_SSL must be set")
	}
	var err error
	c.UseSSL, err = strconv.ParseBool(useSSL)
	if err != nil {
		return err
	}
	//Load Key value
	if c.AesKeyPath != "" {
		c.AesKeyValue, err = ioutil.ReadFile(c.AesKeyPath)
		if err != nil {
			return fmt.Errorf("error when trying to read AES key file : %s", err.Error())
		}
	}
	//Init client
	c.Client, err = minio.New(c.EndPoint, c.AccessKeyID, c.SecretKey, c.UseSSL)
	if err != nil {
		return fmt.Errorf("error when trying to initialize Minio client : %s", err.Error())
	}
	return err
}

func (c *Config) Put(s3object S3Object, content []byte) error {
	exist, err := c.Client.BucketExists(s3object.Bucket)
	if err != nil {
		return fmt.Errorf("unable to stat bucket : %s", err.Error())
	}
	if !exist {
		err = c.Client.MakeBucket(s3object.Bucket, "")
		if err != nil {
			return fmt.Errorf("unable to make missing bucket : %s", err.Error())
		}
	}

	encryption, err := encrypt.NewSSEC(c.AesKeyValue)
	if err != nil {
		return fmt.Errorf("unable to get AES key to encrypt file : %s", err.Error())
	}

	_, err = c.Client.PutObject(s3object.Bucket, s3object.Object, bytes.NewReader(content), int64(len(content)), minio.PutObjectOptions{
		ContentType:          "application/octet-stream",
		ServerSideEncryption: encryption,
	})
	if err != nil {
		return fmt.Errorf("error when trying to put object : %s", err.Error())
	}
	return nil

}

func (c *Config) Stat(object S3Object) (info minio.ObjectInfo, err error) {
	encryption, err := encrypt.NewSSEC(c.AesKeyValue)
	if err != nil {
		return minio.ObjectInfo{}, fmt.Errorf("unable to get AES key to encrypt file : %s", err.Error())
	}
	info, err = c.Client.StatObject(object.Bucket, object.Object, minio.StatObjectOptions{
		minio.GetObjectOptions{
			ServerSideEncryption: encryption,
		},
	})
	if err != nil {
		return minio.ObjectInfo{}, fmt.Errorf("unable to stat object : %s", err.Error())
	}
	return
}

func (c *Config) Get(s3object S3Object) (content []byte, err error) {
	obj, err := c.getObject(s3object, true)
	if err != nil {
		return nil, fmt.Errorf("error when trying to get object : %s", err.Error())
	}
	content, err = ioutil.ReadAll(obj)
	if err != nil {
		return nil, fmt.Errorf("error when trying to get object content : %s", err.Error())
	}
	return content, nil
}

func (c *Config) getObject(s3object S3Object, encrypted bool) (obj *minio.Object, err error) {
	if encrypted {
		encryption, err := encrypt.NewSSEC(c.AesKeyValue)
		if err != nil {
			return nil, fmt.Errorf("unable to get AES key to encrypt file : %s", err.Error())
		}
		obj, err = c.Client.GetObject(s3object.Bucket, s3object.Object, minio.GetObjectOptions{
			ServerSideEncryption: encryption,
		})
		if err != nil {
			return nil, fmt.Errorf("unable to get object : %s", err.Error())
		}
	} else {
		var err error
		obj, err = c.Client.GetObject(s3object.Bucket, s3object.Object, minio.GetObjectOptions{})
		if err != nil {
			return nil, fmt.Errorf("unable to get object : %s", err.Error())
		}
	}
	return obj, nil
}
