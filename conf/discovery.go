package conf

type NacosDiscoveryConf struct {
	Ip                  string
	Port                int
	TimeoutMs           int
	Namespace           string
	Group               string
	ServiceName         string
	Username            string
	Password            string
	NotLoadCacheAtStart bool
}
