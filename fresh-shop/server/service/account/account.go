package account

import (
	"fresh-shop/server/global"
	"fresh-shop/server/model/account"
	accountReq "fresh-shop/server/model/account/request"
	"fresh-shop/server/model/common/request"
)

type AccountService struct {
}

// CreateAccount 创建Account记录
// Author [likfees](https://github.com/likfees)
func (accountService *AccountService) CreateAccount(account account.Account) (err error) {
	err = global.DB.Create(&account).Error
	return err
}

// DeleteAccount 删除Account记录
// Author [likfees](https://github.com/likfees)
func (accountService *AccountService) DeleteAccount(account account.Account) (err error) {
	err = global.DB.Delete(&account).Error
	return err
}

// DeleteAccountByIds 批量删除Account记录
// Author [likfees](https://github.com/likfees)
func (accountService *AccountService) DeleteAccountByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]account.Account{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAccount 更新Account记录
// Author [likfees](https://github.com/likfees)
func (accountService *AccountService) UpdateAccount(account account.Account) (err error) {
	err = global.DB.Save(&account).Error
	return err
}

// GetAccount 根据id获取Account记录
// Author [likfees](https://github.com/likfees)
func (accountService *AccountService) GetAccount(id uint) (account account.Account, err error) {
	err = global.DB.Where("id = ?", id).First(&account).Error
	return
}

// GetAccountInfoList 分页获取Account记录
// Author [likfees](https://github.com/likfees)
func (accountService *AccountService) GetAccountInfoList(info accountReq.AccountSearch) (list []account.Account, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//var ac account.Account
	// 创建db
	db := global.DB.Model(&account.Account{}).
		Joins("User").
		//Preload("User").
		Preload("Group")
	var accounts []account.Account
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("user_account.created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.GroupId != nil {
		db = db.Where("user_account.group_id = ?", info.GroupId)
	}
	if info.Username != "" {
		db = db.Where("User.userName like ?", "%"+info.Username+"%")
	}
	if info.Phone != "" {
		db = db.Where("User.phone like ?", "%"+info.Phone+"%")
	}
	err = db.Limit(limit).Offset(offset).Find(&accounts).Error
	if err != nil {
		return
	}

	// 解决 count 也附带 limit 的问题
	dbCount := global.DB.Debug().Model(&account.Account{}).
		Joins("User").
		//Preload("User").
		Preload("Group")
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		dbCount = dbCount.Where("user_account.created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.GroupId != nil {
		dbCount = dbCount.Where("user_account.group_id = ?", info.GroupId)
	}
	if info.Username != "" {
		dbCount = dbCount.Where("User.userName like ?", "%"+info.Username+"%")
	}
	if info.Phone != "" {
		dbCount = dbCount.Where("User.phone like ?", "%"+info.Phone+"%")
	}
	err = dbCount.Count(&total).Error
	if err != nil {
		return
	}
	return accounts, total, err
}
