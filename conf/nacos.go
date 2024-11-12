package conf

type NacosConf struct {
	Config    NacosConfigConf
	Discovery NacosDiscoveryConf
}
type NacosConfig struct {
	Nacos NacosConf
}
