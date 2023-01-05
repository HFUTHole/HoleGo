package fileservice

import (
	"github.com/bwmarrin/snowflake"
	"github.com/minio/minio-go/v6"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"hole/pkgs/config/logger"
	"hole/pkgs/exception"
	"io"
	"log"
)

const (
	ContentBucket = "content"
	AvatarBucket  = "avatar"
	TempBucket    = "temp"
	Location      = "cn-north-1"
)

var cfg *Config
var client *minio.Client

var ID, _ = snowflake.NewNode(1)

type Config struct {
	EndPoint        string
	AccessKeyID     string
	SecretAccessKey string
	Secure          bool
}

func InitConfig() {
	endPoint := viper.GetString("minio.end_point")
	if endPoint == "" {
		endPoint = "127.0.0.1:9000"
	}

	accessKey := viper.GetString("minio.access_key")
	secret := viper.GetString("minio.secret")
	secure := viper.GetBool("minio.secure")

	cfg = &Config{
		EndPoint:        endPoint,
		AccessKeyID:     accessKey,
		SecretAccessKey: secret,
		Secure:          secure,
	}
}

func Init() {
	InitConfig()
	var err error
	client, err = minio.New(
		cfg.EndPoint,
		cfg.AccessKeyID,
		cfg.SecretAccessKey,
		cfg.Secure,
	)

	if err != nil {
		panic("minio failed to start")
	}

	MakeBucket(ContentBucket)
	MakeBucket(AvatarBucket)
	MakeBucket(TempBucket)
}

func MakeBucket(bucket string) {
	err := client.MakeBucket(bucket, Location)
	if err != nil {
		exists, err := client.BucketExists(bucket)
		if err == nil && exists {
			log.Printf("We already own %s\n", bucket)
		} else {
			log.Fatalln(err)
		}
	}
	log.Printf("Successfully created %s\n", bucket)
}

func GetClient() *minio.Client {
	return client
}

func Exists(bucket string, id string) bool {
	_, err := client.StatObject(bucket, id, minio.StatObjectOptions{})
	if err != nil {
		return false
	}
	return true
}

func PutFile(reader io.Reader, size int64, contentType string) (string, error) {
	name := ID.Generate().String()
	n, err := client.PutObject(TempBucket, name, reader, size, minio.PutObjectOptions{ContentType: contentType})

	if err != nil || n != size {
		return "", err
	}

	return name, nil
}

func GetContent(id string) (*minio.Object, error) {
	object, err := client.GetObject(ContentBucket, id, minio.GetObjectOptions{})
	return object, err
}

func GetAvatar(id string) (*minio.Object, error) {
	object, err := client.GetObject(AvatarBucket, id, minio.GetObjectOptions{})
	return object, err
}

func GetTemp(id string) (*minio.Object, error) {
	object, err := client.GetObject(TempBucket, id, minio.GetObjectOptions{})
	return object, err
}

func copyFile(srcBucket string, dstBucket string, id string) error {
	src := minio.NewSourceInfo(srcBucket, id, nil)

	dst, err := minio.NewDestinationInfo(dstBucket, id, nil, nil)
	if err != nil {
		logger.GetLogger().Error("构建目标信息错误",
			zap.String("src", srcBucket),
			zap.String("dst", dstBucket),
			zap.String("id", id),
			zap.Error(err),
		)
		return &exception.ClientException{Msg: "构建目标文件信息错误"}
	}

	// Copy object call
	err = client.CopyObject(dst, src)
	if err != nil {
		logger.GetLogger().Error("拷贝文件错误",
			zap.String("src", srcBucket),
			zap.String("dst", dstBucket),
			zap.String("id", id),
			zap.Error(err),
		)
		return &exception.ClientException{Msg: "文件不存在"}
	}

	return nil
}

func CopyFileToContent(id string) error {
	return copyFile(TempBucket, ContentBucket, id)
}

func CopyFileToAvatar(id string) error {
	return copyFile(TempBucket, ContentBucket, id)
}

func DeleteFile(bucket string, id string) bool {
	err := client.RemoveObject(bucket, id)
	if err != nil {
		return false
	}
	return true
}
