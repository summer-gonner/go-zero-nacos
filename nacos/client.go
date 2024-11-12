package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
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
func InitNacosConfigClient(nacosConf conf.NacosConf) error {
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
	_, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ServerConfigs: serverConfig,
			ClientConfig:  &clientConfig,
		},
	)
	if err != nil {
		log.Fatalf("Failed to create Nacos config client: %v", err)
	}
	//registerService(nacosClient, nacosConf)
	return nil
}
