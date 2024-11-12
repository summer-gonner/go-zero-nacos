package nacos

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"io"
	"strings"
)

func getConfig(iClient config_client.IConfigClient, dataId string, group string) (io.Reader, error) {
	// 获取配置内容
	content, err := iClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
	if err != nil {
		return nil, err
	}
	exists, err := checkConfigExists(iClient, dataId, group)
	if exists {
		// 返回一个字节流 (io.Reader)
		return strings.NewReader(content), nil
	} else {
		return nil, nil
	}
}

// 判断指定的配置是否存在
func checkConfigExists(iClient config_client.IConfigClient, dataId, group string) (bool, error) {
	// 获取配置内容
	content, err := iClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
	if err != nil {
		return false, fmt.Errorf("failed to get config from Nacos: %v", err)
	}

	// 如果返回内容为空，则表示配置不存在
	if content == "" {
		return false, nil
	}

	// 如果返回内容不为空，则表示配置存在
	return true, nil
}
func interfaceToIo(content interface{}) (io.Reader, error) {
	// 尝试将 interface{} 转换为 io.Reader
	reader, ok := content.(io.Reader)
	if !ok {
		return nil, fmt.Errorf("value does not implement io.Reader")
	}

	// 如果转换成功，返回 reader
	return reader, nil
}
