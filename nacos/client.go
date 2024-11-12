package nacos

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	viper2 "github.com/spf13/viper"
	"github.com/summer-gonner/go-zero-nacos/conf"
	"log"
)

func InitNacosDiscoveryClient(nacosConf conf.NacosConf) error {
	//配置nacos信息
	clientConfig := constant.ClientConfig{
		NamespaceId:         nacosConf.Discovery.Namespace,
		TimeoutMs:           uint64(nacosConf.Discovery.TimeoutMs),
		NotLoadCacheAtStart: nacosConf.Discovery.NotLoadCacheAtStart,
		Username:            nacosConf.Discovery.Username,
		Password:            nacosConf.Discovery.Password,
	}
	//配置nacos服务地址
	serverConfig := []constant.ServerConfig{
		{
			IpAddr: nacosConf.Discovery.Ip,
			Port:   uint64(nacosConf.Discovery.Port),
		},
	}
	nacosDiscoveryClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ServerConfigs: serverConfig,
			ClientConfig:  &clientConfig,
		},
	)
	if err != nil {
		log.Fatalf("Failed to create Nacos discovery client: %v", err)
	}
	registerService(nacosDiscoveryClient)
	return nil
}
func InitNacosConfigClient(nacosConf conf.NacosConf) (viper2.Viper, error) {
	//配置nacos信息
	clientConfig := constant.ClientConfig{
		NamespaceId:         nacosConf.Discovery.Namespace,
		TimeoutMs:           uint64(nacosConf.Discovery.TimeoutMs),
		NotLoadCacheAtStart: nacosConf.Discovery.NotLoadCacheAtStart,
		Username:            nacosConf.Discovery.Username,
		Password:            nacosConf.Discovery.Password,
	}
	//配置nacos服务地址
	serverConfig := []constant.ServerConfig{
		{
			IpAddr: nacosConf.Discovery.Ip,
			Port:   uint64(nacosConf.Discovery.Port),
		},
	}
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ServerConfigs: serverConfig,
			ClientConfig:  &clientConfig,
		},
	)
	if err != nil {
		log.Fatalf("Failed to create Nacos config client: %v", err)
	}
	// 创建 viper 实例
	v := viper2.New()

	// 获取并加载每个共享配置
	for _, sharedConfig := range nacosConf.Config.SharedConfigs {
		// 从 Nacos 获取配置内容
		configContent, err := getConfig(configClient, sharedConfig.DataID, sharedConfig.Group)
		if err != nil {
			log.Printf("Failed to get config: %v, 文件名为: {%v}, Group 为: {%v}", err, sharedConfig.DataID, sharedConfig.Group)
			continue
		}

		// 使用 Viper 加载配置
		err = v.ReadConfig(configContent)
		if err != nil {
			log.Printf("Failed to read config into Viper: %v", err)
			continue
		}

		// 配置加载成功，返回 Viper 实例
		return *v, nil
	}

	// 如果没有找到有效的配置文件，返回错误
	return viper2.Viper{}, fmt.Errorf("no valid config found in Nacos")
}
