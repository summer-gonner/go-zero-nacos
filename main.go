package main

import (
	"flag"
	conf2 "github.com/summer-gonner/go-zero-nacos/conf"
	"github.com/summer-gonner/go-zero-nacos/nacos"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/config.yaml", "the config file")

func main() {
	var cfg conf2.NacosConfig
	conf.MustLoad(*configFile, &cfg)
	//log.Printf("结果%s\n", cfg.Application.Name)
	err := nacos.InitNacosDiscoveryClient(cfg.Nacos)
	if err != nil {
		return
	}
}
