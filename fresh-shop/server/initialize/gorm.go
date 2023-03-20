package initialize

import (
	"os"

	"fresh-shop/server/global"
	"fresh-shop/server/model/file"
	"fresh-shop/server/model/system"

	"fresh-shop/server/model/account"
	"fresh-shop/server/model/business"
	"fresh-shop/server/model/shop"
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

		file.ExaFile{},
		file.ExaFileChunk{},
		file.ExaFileUploadAndDownload{}, account.SysRecharge{}, account.UserFinanceType{}, account.UserFinance{}, business.Banner{}, shop.Category{}, shop.Brand{}, shop.BrandCategory{}, shop.Tags{}, shop.Goods{}, shop.GoodsDescription{}, shop.GoodsImage{}, shop.GoodsSpec{}, shop.GoodsSpecItem{}, shop.GoodsSpecValue{},
	)
	if err != nil {
		global.Log.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.Log.Info("register table success")
	freshShopMysql.AutoMigrate(account.AccountGroup{}, account.Account{})
}
