package config

// todo：配置文件暂时没有
type Config struct {
	Zap Zap `mapstructure:"zap" json:"zap" yaml:"zap"`
}
