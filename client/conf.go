package client

import (
	"github.com/summer-gonner/go-zero-nacos/config"
	"github.com/summer-gonner/go-zero-nacos/discovery"
)

type NacosConf struct {
	Config    config.NacosConfigConf
	Discovery discovery.NacosDiscoveryConf
}
