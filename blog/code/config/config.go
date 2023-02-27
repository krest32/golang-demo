package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

// 配置实体类
type Config struct {
	boot_port           string
	datasource_userName string
	datasource_password string
	datasource_ip       string
	datasource_port     string
	datasource_dbName   string
}

// 配置文件地址
var iniFilePath string
var config Config
var concfg *ini.File

// 返回Config指针
func ParseConfig(iniFile string) *Config {

	iniFilePath = iniFile
	cfg, err := ini.Load(iniFilePath)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
		return nil
	} else {
		concfg = cfg
		config.boot_port = GetBootPort()
		config.datasource_dbName = GetDatasourceDBName()
		config.datasource_ip = GetDatasourceIP()
		config.datasource_userName = GetDatasourceUserName()
		config.datasource_password = GetDatasourcePassword()
		config.datasource_port = GetDatasourcePort()

		return &config
	}
}

func GetDatasourceUserName() string {
	return GetValue("datasource", "userName").String()
}

func GetDatasourcePassword() string {
	return GetValue("datasource", "password").String()
}

func GetDatasourceIP() string {
	return GetValue("datasource", "ip").String()
}

func GetDatasourcePort() string {
	return GetValue("datasource", "port").String()
}

func GetDatasourceDBName() string {
	return GetValue("datasource", "dbName").String()
}

func GetBootPort() string {
	return GetValue("boot", "port").String()
}

func GetValue(section, key string) *ini.Key {
	return concfg.Section(section).Key(key)
}
