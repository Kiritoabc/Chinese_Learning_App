package utils

import (
	"Chinese_Learning_App/global"
	"context"
	"log"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
)

// UploadToMinio uploads a file to Minio with the given object name and content type.
// It returns an error if there was an issue during the upload.
func UploadToMinio(bucketName, objectName string, file *multipart.FileHeader, contentType string) error {
	// 初使化 minio client对象
	ctx := context.Background()
	minioClient := global.CLA_Minio_Client
	location := "us-east-1"
	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
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

// UploadVideoToMinio 该方法将视频文件上传到MinIO。
// 参数说明：
// filePath：视频文件的路径。
// bucketName：MinIO存储桶的名称。
// objectName：MinIO对象（文件）的名称。
// minioClient：MinIO客户端实例。
// 返回值：
// error：上传过程中发生的错误。如果上传成功，则返回nil。
// 示例：
func UploadVideoToMinio(filePath, bucketName, objectName string, minioClient *minio.Client) error {
	_, err := minioClient.FPutObject(context.Background(), bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: "video/map4"})
	if err != nil {
		// 处理创建客户端时发生的错误
		global.CLA_LOG.Error(err.Error())
	}
	return nil
}
