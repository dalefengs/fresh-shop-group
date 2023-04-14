package main

import (
	"go.uber.org/zap"

	"fresh-shop/server/core"
	"fresh-shop/server/global"
	"fresh-shop/server/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       Swagger Example API
// @version                     0.0.1
// @description                 This is a sample Server pets
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	global.Viper = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.Log = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.Log)
	// 获取 zap SugarLog
	global.SugarLog = global.Log.Sugar()
	global.DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()
	initialize.Wechat()
	// 自动迁移数据表 本平台不使用
	//if global.DB != nil {
	//	initialize.RegisterTables() // 初始化表
	//	// 程序结束前关闭数据库链接
	//	db, _ := global.DB.DB()
	//	defer db.Close()
	//}
	core.RunWindowsServer()
}
