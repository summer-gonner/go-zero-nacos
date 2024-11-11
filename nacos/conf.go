package nacos

import (
	"github.com/summer-gonner/go-zero-nacos/config"
	"github.com/summer-gonner/go-zero-nacos/discovery"
)

type NacosConf struct {
	Config    config.NacosConfigConf
	Discovery discovery.NacosDiscoveryConf
}
type Config struct {
	Server      Server
	Application Application
	Nacos       NacosConf
}

type Server struct {
	Port int
}
type Application struct {
	Name string
}
