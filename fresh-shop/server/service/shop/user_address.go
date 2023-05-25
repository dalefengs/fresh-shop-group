package shop

import (
	"errors"
	"fresh-shop/server/global"
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/shop"
	shopReq "fresh-shop/server/model/shop/request"
	"gorm.io/gorm"
)

type UserAddressService struct {
}

// CreateUserAddress 创建UserAddress记录
// Author [likfees](https://github.com/likfees)
func (userAddressService *UserAddressService) CreateUserAddress(userAddress shop.UserAddress) (address shop.UserAddress, err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if *userAddress.IsDefault == 1 {
			// 更新所有默认值为 0
			err = tx.Model(&shop.UserAddress{}).Where("user_id = ? and is_default = 1", userAddress.UserId).Update("is_default", 0).Error
			if err != nil {
				return err
			}
		}
		err = tx.Create(&userAddress).Error
		return err
	})
	return userAddress, err
}

// DeleteUserAddress 删除UserAddress记录
// Author [likfees](https://github.com/likfees)
func (userAddressService *UserAddressService) DeleteUserAddress(userAddress shop.UserAddress) (err error) {
	err = global.DB.Delete(&userAddress).Error
	return err
}

// DeleteUserAddressByIds 批量删除UserAddress记录
// Author [likfees](https://github.com/likfees)
func (userAddressService *UserAddressService) DeleteUserAddressByIds(ids request.IdsReq) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id in ?", ids.Ids).Delete(&shop.UserAddress{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateUserAddress 更新UserAddress记录
// Author [likfees](https://github.com/likfees)
func (userAddressService *UserAddressService) UpdateUserAddress(userAddress shop.UserAddress) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if *userAddress.IsDefault == 1 {
			// 更新所有默认值为 0
			err = tx.Model(&shop.UserAddress{}).Where("user_id = ? and is_default = 1", userAddress.UserId).Update("is_default", 0).Error
			if err != nil {
				return err
			}
		}
		err = tx.Save(&userAddress).Error
		return err
	})

	return err
}

// GetUserAddress 根据id获取UserAddress记录
// Author [likfees](https://github.com/likfees)
func (userAddressService *UserAddressService) GetUserAddress(id uint) (userAddress shop.UserAddress, err error) {
	err = global.DB.Where("id = ?", id).First(&userAddress).Error
	return
}

// GetUserDeafultAddress 根据用户id获取默认UserAddress记录
// Author [likfees](https://github.com/likfees)
func (userAddressService *UserAddressService) GetUserDeafultAddress(userId uint) (userAddress shop.UserAddress, err error) {
	err = global.DB.Where("user_id = ? and is_default = 1", userId).First(&userAddress).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userAddress, nil
		}
		return
	}
	return
}

// GetUserAddressInfoList 获取UserAddress记录
// Author [likfees](https://github.com/likfees)
func (userAddressService *UserAddressService) GetUserAddressInfoList(info shopReq.UserAddressSearch) (list []shop.UserAddress, err error) {
	// 创建db
	db := global.DB.Model(&shop.UserAddress{}).Where("user_id = ?", info.UserId)
	var userAddresss []shop.UserAddress
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Order("is_default desc").Find(&userAddresss).Error
	return userAddresss, err
}
