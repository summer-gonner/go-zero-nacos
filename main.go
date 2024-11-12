package main

import (
	"flag"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	conf2 "github.com/summer-gonner/go-zero-nacos/conf"
	"log"
	"strings"
)

var configFile = flag.String("f", "etc/config.yaml", "the config file")

// 创建 Nacos 客户端并获取配置内容
func InitNacosConfigClient(nacosConf conf2.NacosConfig) (viper.Viper, error) {
	// 配置 Nacos 客户端
	clientConfig := constant.ClientConfig{
		NamespaceId: nacosConf.Nacos.Discovery.Namespace,
		TimeoutMs:   uint64(nacosConf.Nacos.Discovery.TimeoutMs),
		Username:    nacosConf.Nacos.Discovery.Username,
		Password:    nacosConf.Nacos.Discovery.Password,
	}

	// 配置 Nacos 服务地址
	serverConfig := []constant.ServerConfig{
		{
			IpAddr: nacosConf.Nacos.Discovery.Ip,
			Port:   uint64(nacosConf.Nacos.Discovery.Port),
		},
	}

	// 创建 Nacos 配置客户端
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ServerConfigs: serverConfig,
			ClientConfig:  &clientConfig,
		},
	)
	if err != nil {
		log.Fatalf("Failed to create Nacos config client: %v", err)
		return viper.Viper{}, err
	}

	// 创建 viper 实例
	v := viper.New()

	// 获取并加载每个共享配置
	for _, sharedConfig := range nacosConf.Nacos.Config.SharedConfigs {
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
	return viper.Viper{}, fmt.Errorf("no valid config found in Nacos")
}

// 从 Nacos 获取配置内容
func getConfig(configClient clients.IConfigClient, dataId, group string) (io.Reader, error) {
	// 获取配置内容
	content, err := configClient.GetConfig(dataId, group)
	if err != nil {
		return nil, err
	}

	// 返回一个字节流 (io.Reader)
	return strings.NewReader(content), nil
}

func main() {
	// 假设配置已通过某种方式加载到 nacosConf 中
	var nacosConf NacosConf
	// 配置的初始化部分（这部分根据实际需要初始化 nacosConf）

	// 获取 Nacos 配置
	v, err := InitNacosConfigClient(nacosConf)
	if err != nil {
		log.Fatalf("Error initializing Nacos config: %v", err)
	}

	// 使用返回的 Viper 配置实例
	fmt.Println("Config loaded successfully")

	// 通过 Viper 获取某个配置项的值
	appName := v.GetString("app.name")
	fmt.Println("App Name:", appName)
}
