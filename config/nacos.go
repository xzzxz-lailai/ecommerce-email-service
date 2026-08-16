package config

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func GetNacosConfig() (string, error) {
	// 配置 Nacos 服务信息
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: "127.0.0.1",
			Port:   8848,
		},
	}

	// 客户端配置
	clientConfig := constant.ClientConfig{
		NamespaceId:         "",   // 当 namespace 是 public 时，此处填空字符串
		TimeoutMs:           5000, // 请求 Nacos 的超时时间（毫秒）
		NotLoadCacheAtStart: true, // 启动时是否不读取本地缓存
		LogDir:              "./tmp/nacos/log",
		CacheDir:            "./tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// 创建动态配置客户端
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		panic(err)
	}

	// 从 Nacos 拉取配置
	nacosConfig, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "email-service.yaml",
		Group:  "DEFAULT_GROUP",
	})
	if err != nil {
		panic(err)
	}

	return nacosConfig, nil
}