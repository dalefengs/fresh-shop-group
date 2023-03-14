package system

import (
	"fresh-shop/server/config"
	"fresh-shop/server/global"
	"fresh-shop/server/model/system"
	"fresh-shop/server/utils"
	"go.uber.org/zap"
)

//@author: [piexlmax](https://github.com/likfees)
//@function: GetSystemConfig
//@description: 读取配置文件
//@return: conf config.Server, err error

type SystemConfigService struct{}

func (systemConfigService *SystemConfigService) GetSystemConfig() (conf config.Server, err error) {
	return global.Config, nil
}

// @description   set system config,
//@author: [piexlmax](https://github.com/likfees)
//@function: SetSystemConfig
//@description: 设置配置文件
//@param: system model.System
//@return: err error

func (systemConfigService *SystemConfigService) SetSystemConfig(system system.System) (err error) {
	cs := utils.StructToMap(system.Config)
	for k, v := range cs {
		global.Viper.Set(k, v)
	}
	err = global.Viper.WriteConfig()
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: GetServerInfo
//@description: 获取服务器信息
//@return: server *utils.Server, err error

func (systemConfigService *SystemConfigService) GetServerInfo() (server *utils.Server, err error) {
	var s utils.Server
	s.Os = utils.InitOS()
	if s.Cpu, err = utils.InitCPU(); err != nil {
		global.Log.Error("func utils.InitCPU() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Ram, err = utils.InitRAM(); err != nil {
		global.Log.Error("func utils.InitRAM() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Disk, err = utils.InitDisk(); err != nil {
		global.Log.Error("func utils.InitDisk() Failed", zap.String("err", err.Error()))
		return &s, err
	}

	return &s, nil
}
