package app

import "go.introduce/chapter02/code05/app/config"

// GetCver 获取版本信息
func GetCver() string {
	return config.Cver
}
