package config

import (
	"strings"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type EmailSmtpConfig struct {
	Host       string `mapstructure:"host"`        // SMTP 服务器地址
	Port       int    `mapstructure:"port"`        // SMTP 服务器端口
	Email      string `mapstructure:"email"`       // 发件邮箱
	AuthCode   string `mapstructure:"auth_code"`   // SMTP 授权码
	SenderName string `mapstructure:"sender_name"` // 发件人显示名称
}
type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Pass string `mapstructure:"pass"`
	DB   int    `mapstructure:"db"`
}
type EtcdConfig struct {
	Host         string `mapstructure:"host"` // etcd 服务地址
	ServerName   string `mapstructure:"server_name"`
	ServeAddress string `mapstructure:"address"`
}

type Config struct {
	Server    ServerConfig
	EmailSmtp EmailSmtpConfig
	Redis     RedisConfig
	Etcd      EtcdConfig
}

var Cfg Config

func InitConfig() *Config {
	nacosConfig, err := GetNacosConfig()
	if err != nil {
		panic(err)
	}

	// 配置文件类型
	viper.SetConfigType("yaml")

	// 读取配置内容
	err = viper.ReadConfig(strings.NewReader(nacosConfig))
	if err != nil {
		panic(err)
	}

	// 将读取到的配置解码到 Config 结构体
	if err := viper.Unmarshal(&Cfg); err != nil {
		panic(err)
	}

	InitRedis() // 初始化redis

	return &Cfg
}
