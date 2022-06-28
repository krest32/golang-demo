package service

import "demo-basic/blog/code/mapper"

func GetAllApps() []mapper.Apps {
	apps := mapper.GetApps()
	return apps
}

