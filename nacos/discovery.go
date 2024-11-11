package nacos

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"log"
)

func registerService(namingClient naming_client.INamingClient) {
	//// 定义服务实例
	//instance := vo.RegisterInstanceParam{
	//	Ip:          nacosDiscoveryConfig.Ip,
	//	Port:        uint64(nacosDiscoveryConfig.Port), // 服务运行的端口
	//	ServiceName: nacosDiscoveryConfig.ServiceName,  // 服务名称
	//	Weight:      1,                                 // 服务的权重
	//	ClusterName: "DEFAULT",                         // 集群名称
	//	GroupName:   nacosDiscoveryConfig.Group,        // 分组名称
	//}
	//
	//// 将服务注册到 Nacos
	//result, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
	//	Ip:          instance.Ip,
	//	Port:        instance.Port,
	//	ServiceName: instance.ServiceName,
	//	Weight:      instance.Weight,
	//	ClusterName: instance.ClusterName,
	//	GroupName:   instance.GroupName,
	//})
	//if err != nil {
	//	log.Fatalf("Failed to register service to Nacos: %v", err)
	//}
	//
	//fmt.Println("Service registered successfully.", result)
}

func deregisterService(namingClient naming_client.INamingClient) {
	result, err := namingClient.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          "127.0.0.1",
		Port:        8080,
		ServiceName: "my-service",
		Cluster:     "DEFAULT",
		GroupName:   "DEFAULT_GROUP",
	})
	if err != nil {
		log.Fatalf("Failed to deregister service from Nacos: %v", err)
	}

	fmt.Println("Service deregistered successfully.", result)
}
