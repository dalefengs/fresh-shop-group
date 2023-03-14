package system

import "fresh-shop/server/global"

type SysBaseMenuBtn struct {
	global.DbModel
	Name          string `json:"name" gorm:"comment:按钮关键key"`
	Desc          string `json:"desc" gorm:"按钮备注"`
	SysBaseMenuID uint   `json:"sysBaseMenuID" gorm:"comment:菜单ID"`
}
