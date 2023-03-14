package system

import (
	"fresh-shop/server/global"
)

type JwtBlacklist struct {
	global.DbModel
	Jwt string `gorm:"type:text;comment:jwt"`
}
