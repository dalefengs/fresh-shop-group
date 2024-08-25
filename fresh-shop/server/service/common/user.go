package common

import (
	"errors"
	"fresh-shop/server/global"
	sysModel "fresh-shop/server/model/system"
	"gorm.io/gorm"
)

// GetUserInfo 获取用户信息
func GetUserInfo(userId uint) (user sysModel.SysUser, err error) {
	err = global.DB.Where("id = ?", userId).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user, err
	}
	return user, err
}
