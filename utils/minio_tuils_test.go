package utils

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/gabriel-vasile/mimetype"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func TestUploadToMinio(t *testing.T) {
	// 创建MinIO客户端
	client, err := minio.New("localhost:9001", &minio.Options{
		Creds:  credentials.NewStaticV4("minio", "minio@123456", ""),
		Secure: false})
	if err != nil {
		log.Fatal(err)
	}
	// 打开本地视频文件
	if err != nil {
		log.Fatal(err)
	}
	filePath := "D:\\video\\2023-09-26 14-08-23.mp4"
	_, err = mimetype.DetectFile("D:\\video\\2023-09-26 14-08-23.mp4")
	// 创建一个上传对象
	objectName := "video.mp4"
	buckName := "test"
	_, err = client.FPutObject(context.Background(), buckName, objectName, filePath, minio.PutObjectOptions{ContentType: "video/map4"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("上传成功")
}
