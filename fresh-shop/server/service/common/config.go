package common

import (
	"errors"
	"fresh-shop/server/global"
	"fresh-shop/server/model/system"
)

// 公共获取配置信息方法

// ErrConfigDisabled 配置已经禁用错误
var ErrConfigDisabled = errors.New("该配置禁用")

// GetSysConfig 通过参数名获取配置参数信息
func GetSysConfig(name string) (string, error) {
	var config system.SysConfig
	if err := global.DB.Where("name = ?", name).First(&config).Error; err != nil {
		return "", err
	}
	if *config.Status == 0 {
		return "", ErrConfigDisabled
	}
	return config.Value, nil
}

// GetAllSysConfig 获取所有配置参数并返回一个Map
func GetAllSysConfig() map[string]string {
	result := make(map[string]string)
	var configs []system.SysConfig
	if err := global.DB.Find(&configs).Error; err != nil {
		return result
	}
	for _, config := range configs {
		if *config.Status == 1 {
			result[config.Name] = config.Value
		}
	}
	return result
}
