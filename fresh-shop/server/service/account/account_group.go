package account

import (
	"errors"
	"fresh-shop/server/global"
	"fresh-shop/server/model/account"
	accountReq "fresh-shop/server/model/account/request"
	"fresh-shop/server/model/common/request"
	sysModel "fresh-shop/server/model/system"
	"gorm.io/gorm"
)

type AccountGroupService struct {
}

// CreateAccountGroup 创建AccountGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (userAccountGroupService *AccountGroupService) CreateAccountGroup(userAccountGroup account.AccountGroup) (err error) {
	var accountGroup account.AccountGroup
	if !errors.Is(global.DB.Where("name_en = ?", userAccountGroup.NameEn).First(&accountGroup).Error, gorm.ErrRecordNotFound) {
		return errors.New("币种已存在！")
	}
	// 创建币种账户
	err = global.MustGetGlobalDBByDBName("freshShopMysql").Create(&userAccountGroup).Error
	// 为所有用户创建订单流水
	return err
}

// 同步初始化用户账户表
func syncUserAccount(groupId int64) error {
	// 查询所有用户
	var userList []sysModel.SysUser
	err := global.DB.Model(&sysModel.SysUser{}).Find(&userList).Error
	if err != nil {
		return err
	}
	if len(userList) == 0 {
		return errors.New("查询用户列表失败")
	}
	// 如果 user_account 表没有则创建
	//for u := range userList {
	//
	//}
	return nil
}

// DeleteAccountGroup 删除AccountGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (userAccountGroupService *AccountGroupService) DeleteAccountGroup(userAccountGroup account.AccountGroup) (err error) {
	err = global.MustGetGlobalDBByDBName("freshShopMysql").Delete(&userAccountGroup).Error
	return err
}

// DeleteAccountGroupByIds 批量删除AccountGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (userAccountGroupService *AccountGroupService) DeleteAccountGroupByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("freshShopMysql").Delete(&[]account.AccountGroup{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAccountGroup 更新AccountGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (userAccountGroupService *AccountGroupService) UpdateAccountGroup(userAccountGroup account.AccountGroup) (err error) {
	err = global.MustGetGlobalDBByDBName("freshShopMysql").Save(&userAccountGroup).Error
	return err
}

// GetAccountGroup 根据id获取AccountGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (userAccountGroupService *AccountGroupService) GetAccountGroup(id uint) (userAccountGroup account.AccountGroup, err error) {
	err = global.MustGetGlobalDBByDBName("freshShopMysql").Where("id = ?", id).First(&userAccountGroup).Error
	return
}

// GetAccountGroupInfoList 分页获取AccountGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (userAccountGroupService *AccountGroupService) GetAccountGroupInfoList(info accountReq.AccountGroupSearch) (list []account.AccountGroup, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("freshShopMysql").Model(&account.AccountGroup{})
	var userAccountGroups []account.AccountGroup
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.NameEn != "" {
		db = db.Where("name_en LIKE ?", "%"+info.NameEn+"%")
	}
	if info.NameCn != "" {
		db = db.Where("name_cn LIKE ?", "%"+info.NameCn+"%")
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&userAccountGroups).Error
	return userAccountGroups, total, err
}
