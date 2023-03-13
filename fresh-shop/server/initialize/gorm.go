package initialize

import (
	"os"

	"fresh-shop/server/global"
	"fresh-shop/server/model/example"
	"fresh-shop/server/model/system"

	"fresh-shop/server/model/account"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.Config.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	case "oracle":
		return GormOracle()
	case "mssql":
		return GormMssql()
	default:
		return GormMysql()
	}
}

func RegisterTables() {
	freshShopMysql := global.GetGlobalDBByDBName("freshShopMysql")
	db := global.DB
	err := db.AutoMigrate(

		system.SysApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		system.SysDictionary{},
		system.SysOperationRecord{},
		system.SysAutoCodeHistory{},
		system.SysDictionaryDetail{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		system.SysAuthorityBtn{},
		system.SysAutoCode{},

		example.ExaFile{},
		example.ExaCustomer{},
		example.ExaFileChunk{},
		example.ExaFileUploadAndDownload{},
	)
	if err != nil {
		global.Log.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.Log.Info("register table success")
	freshShopMysql.AutoMigrate(account.AccountGroup{}, account.AccountGroup{})
}
