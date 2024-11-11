package main

import (
	"flag"
	"github.com/summer-gonner/go-zero-nacos/discovery"
	"github.com/summer-gonner/go-zero-nacos/nacos"
	"github.com/zeromicro/go-zero/core/conf"
	"log"
)

var configFile = flag.String("f", "etc/config.yaml", "the config file")

func main() {
	var cfg nacos.Config
	conf.MustLoad(*configFile, &cfg)
	log.Printf("路径%s\n", *configFile)
	log.Printf("结果%s\n", cfg.Application.Name)
	discovery.InitNacosDiscoveryClient(cfg.Nacos.Discovery)
}
