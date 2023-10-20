package initialize

import (
	"Chinese_Learning_App/global"

	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/minio/minio-go/v7"
)

func Minio() (*minio.Client, error) {
	return minio.New(global.CLA_CONFIG.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(global.CLA_CONFIG.Minio.AccessKey, global.CLA_CONFIG.Minio.SecretKey, ""),
		Secure: false})
}
