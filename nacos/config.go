package nacos

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func getConfigFromNacos(iClient config_client.IConfigClient, dataId string, group string) (string, error) {
	content, err := iClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
	if err != nil {
		return "", fmt.Errorf("failed to get config from nacos: %v", err)
	}
	return content, nil
}
