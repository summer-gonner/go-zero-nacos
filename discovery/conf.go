package discovery

type NacosDiscoveryConf struct {
	Ip          string
	Port        int
	Namespace   string
	Group       string
	ServiceName string
	Username    string
	Password    string
}
