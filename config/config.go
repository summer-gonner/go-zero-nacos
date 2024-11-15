package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"reflect"
)

type NacosConfigResult struct {
	Contents []string
}

func (nc NacosConfigResult) LoadConfig(v any) {
	// 解析配置
	if len(nc.Contents) > 0 {
		for _, content := range nc.Contents {
			// 确保传入的是一个指向结构体的指针
			if reflect.TypeOf(v).Kind() != reflect.Ptr {
				log.Printf("传入的参数 v 必须是结构体指针类型")
				return
			}

			err := yaml.Unmarshal([]byte(content), v) // 使用 YAML 解析库解析配置
			if err != nil {
				log.Printf("根据结构体加载配置失败 %v", err)
			}
		}
	}
}
