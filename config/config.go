package config

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/configor"
)

type Config struct {
	ServerIP   string `yaml:"server_ip"`
	ServerPort string `yaml:"server_port"`

	EnableLogMonitor bool   `yaml:"enable_log_monitor"`
	LogMonitorUrl    string `yaml:"log_monitor_url"`

	WxCorpID     string `yaml:"wx_corp_id"`
	WxCorpSecret string `yaml:"wx_corp_secret"`
	WxCorpAgent  int    `yaml:"wx_corp_agent"`

	MonitorTime      int `yaml:"monitor_time"`      //监控最近多长时间的错误(分钟:1-30)
	MonitorThreshold int `yaml:"monitor_threshold"` //阈值
	MonitorCD        int `yaml:"monitor_cooldown"`  //监控冷却时间
}

var Cfg *Config

func Init(confPath string) {
	var path string
	path = confPath

	Cfg = &Config{}
	if err := configor.Load(Cfg, path); err != nil {
		panic(err)
	}

	cfgjson, _ := json.MarshalIndent(Cfg, "", "  ")
	fmt.Printf("[ConfigInit] %s\n", string(cfgjson))
}
