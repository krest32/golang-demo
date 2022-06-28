package ini

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

/**
流程总文件
*/
func Parse(iniFilePath string) *ini.File {
	cfg, err := ini.Load(iniFilePath)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
		return nil
	}
	return cfg
}

func GetValue(iniFilePath,section,key string) *ini.Key {
	cfg, err := ini.Load(iniFilePath)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	return cfg.Section(section).Key(key)
}
