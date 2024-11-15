package client

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/summer-gonner/go-zero-nacos/config"
	"log"
)

// InitNacosDiscoveryClient 初始化nacos 注册中心客户端
func InitNacosDiscoveryClient(nacosConf NacosConf) error {
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
	discovertClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ServerConfigs: serverConfig,
			ClientConfig:  &clientConfig,
		},
	)
	if err != nil {
		log.Printf("创建ncos注册中心客户端失败: %v", err)
	}
	res, err := discovertClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          nacosConf.Discovery.Ip,
		Port:        uint64(nacosConf.Discovery.Port),
		Weight:      1,
		Enable:      true,
		Healthy:     true,
		ServiceName: nacosConf.Discovery.ServiceName,
	})
	if err != nil {
		log.Printf("ncos注册服务失败: %v", err)
	}
	log.Printf("注册结果%v", res)
	return nil
}

// InitNacosConfigClient 初始化nacos 配置中心客户端
func InitNacosConfigClient(nacosConf NacosConf) (config.NacosConfigResult, error) {
	//配置nacos信息
	clientConfig := constant.ClientConfig{
		NamespaceId: nacosConf.Config.Namespace,
		TimeoutMs:   uint64(nacosConf.Config.TimeoutMs),
		Username:    nacosConf.Config.Username,
		Password:    nacosConf.Config.Password,
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
		log.Printf("创建nacos配置中心客户端失败: %v", err)
	}
	var contents []string
	for _, sharedConfig := range nacosConf.Config.SharedConfigs {
		content, _ := configClient.GetConfig(vo.ConfigParam{
			DataId: sharedConfig.DataID,
			Group:  sharedConfig.Group,
		})
		contents = append(contents, content)
	}
	var ncr config.NacosConfigResult
	ncr.Contents = contents
	return ncr, nil
}
