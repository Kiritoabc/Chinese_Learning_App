package config

type Minio struct {
	Endpoint   string `mapstructure:",endpoint" yaml:"endpoint" json:"endpoint"`
	AccessKey  string `mapstructure:"accessKey" yaml:"accessKey" json:"accessKey"`
	SecretKey  string `mapstructure:"secretKey" yaml:"secretKey" json:"secretKey"`
	BucketName string `mapstructure:"bucketName" yaml:"bucketName" json:"bucketName"`
	// Region is the region of the minio server
	Region string `mapstructure:"region" yaml:"region" json:"region"`
	UseSSL bool   `mapstructure:"useSSL" yaml:"useSSL" json:"useSSL"`
}
