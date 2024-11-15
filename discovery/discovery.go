package discovery

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"log"
)

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
