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
// Author [piexlmax](https://github.com/piexlmax)
func (accountService *AccountService) CreateAccount(account account.Account) (err error) {
	err = global.MustGetGlobalDBByDBName("freshShopMysql").Create(&account).Error
	return err
}

// DeleteAccount 删除Account记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountService *AccountService) DeleteAccount(account account.Account) (err error) {
	err = global.MustGetGlobalDBByDBName("freshShopMysql").Delete(&account).Error
	return err
}

// DeleteAccountByIds 批量删除Account记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountService *AccountService) DeleteAccountByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("freshShopMysql").Delete(&[]account.Account{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAccount 更新Account记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountService *AccountService) UpdateAccount(account account.Account) (err error) {
	err = global.MustGetGlobalDBByDBName("freshShopMysql").Save(&account).Error
	return err
}

// GetAccount 根据id获取Account记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountService *AccountService) GetAccount(id uint) (account account.Account, err error) {
	err = global.MustGetGlobalDBByDBName("freshShopMysql").Where("id = ?", id).First(&account).Error
	return
}

// GetAccountInfoList 分页获取Account记录
// Author [piexlmax](https://github.com/piexlmax)
func (accountService *AccountService) GetAccountInfoList(info accountReq.AccountSearch) (list []account.Account, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("freshShopMysql").Model(&account.Account{})
	var accounts []account.Account
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.UserId != nil {
		db = db.Where("user_id = ?", info.UserId)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&accounts).Error
	return accounts, total, err
}
