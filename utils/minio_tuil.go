package utils

import (
	"Chinese_Learning_App/global"
	"context"
	"log"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func UploadToMinio(objectName string, file *multipart.FileHeader, contentType string) error {
	// 初使化 minio client对象
	ctx := context.Background()
	minioClient, err := minio.New(global.CLA_CONFIG.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(global.CLA_CONFIG.Minio.AccessKey, global.CLA_CONFIG.Minio.SecretKey, ""),
		Secure: false})
	if err != nil {
		return err
	}
	bucketName := "test"
	location := "us-east-1"
	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}
	src, err1 := file.Open()
	if err1 != nil {
		return err1
	}
	// 获取图片文件的信息
	defer src.Close()

	// 上传:
	_, err = minioClient.PutObject(ctx, bucketName, objectName, src, file.Size, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}
