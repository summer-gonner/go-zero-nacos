package conf

type NacosConfigConf struct {
	Ip            string
	Port          int
	Username      string
	Password      string
	TimeoutMs     int
	Namespace     string
	FileExtension string
	SharedConfigs []SharedConfig
}

type SharedConfig struct {
	DataID  string
	Group   string
	Refresh bool
}
