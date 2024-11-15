package main

import (
	"github.com/emicklei/go-restful/v3/log"
	"github.com/summer-gonner/go-zero-nacos/client"
	conf2 "github.com/zeromicro/go-zero/core/conf"
)

type Config struct {
	Nacos client.NacosConf
}
type K8s struct {
	IP   string `yaml:"IP"`
	Port string `yaml:"Port"`
}

type AppConfig struct {
	Server      Server
	Application Application
}

type Server struct {
	Port string `yaml:"Port"`
}
type Application struct {
	Name string `yaml:"Name"`
}

var configFile = "etc/config.yaml"

func main() {
	var cfg Config
	conf2.MustLoad(configFile, &cfg)
	configClient, err := client.InitNacosConfigClient(cfg.Nacos)
	if err != nil {
		log.Printf("InitNacosConfigClient failed, err:%v\n", err)
	}
	var app AppConfig
	configClient.LoadConfig(&app)
	_ = client.InitNacosDiscoveryClient(cfg.Nacos)
	if err != nil {
		return
	}
}
