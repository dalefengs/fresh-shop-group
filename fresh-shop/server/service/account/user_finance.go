package account

import (
	"fresh-shop/server/global"
	"fresh-shop/server/model/account"
	accountReq "fresh-shop/server/model/account/request"
	"fresh-shop/server/model/common/request"
)

type UserFinanceCashService struct {
}

// CreateUserFinanceCash 创建UserFinanceCash记录
// Author [likfees](https://github.com/likfees)
func (userFinanceCashService *UserFinanceCashService) CreateUserFinanceCash(userFinanceCash account.UserFinance) (err error) {
	err = global.DB.Create(&userFinanceCash).Error
	return err
}

// DeleteUserFinanceCash 删除UserFinanceCash记录
// Author [likfees](https://github.com/likfees)
func (userFinanceCashService *UserFinanceCashService) DeleteUserFinanceCash(userFinanceCash account.UserFinance) (err error) {
	err = global.DB.Delete(&userFinanceCash).Error
	return err
}

// DeleteUserFinanceCashByIds 批量删除UserFinanceCash记录
// Author [likfees](https://github.com/likfees)
func (userFinanceCashService *UserFinanceCashService) DeleteUserFinanceCashByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]account.UserFinance{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateUserFinanceCash 更新UserFinanceCash记录
// Author [likfees](https://github.com/likfees)
func (userFinanceCashService *UserFinanceCashService) UpdateUserFinanceCash(userFinanceCash account.UserFinance) (err error) {
	err = global.DB.Save(&userFinanceCash).Error
	return err
}

// GetUserFinanceCash 根据id获取UserFinanceCash记录
// Author [likfees](https://github.com/likfees)
func (userFinanceCashService *UserFinanceCashService) GetUserFinanceCash(id uint) (userFinanceCash account.UserFinance, err error) {
	err = global.DB.Where("id = ?", id).First(&userFinanceCash).Error
	return
}

// GetUserFinanceInfoList 分页获取UserFinanceCash记录
// Author [likfees](https://github.com/likfees)
func (userFinanceCashService *UserFinanceCashService) GetUserFinanceInfoList(info accountReq.UserFinanceSearch) (list []account.UserFinance, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Table("user_finance_" + info.GroupNameEn).
		Joins("User").
		Preload("FinanceType")
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.UserSearch.Username != "" {
		db = db.Where("User.username like ?", "%"+info.UserSearch.Username+"%")
	}
	if info.UserSearch.Phone != "" {
		db = db.Where("User.phone like ?", "%"+info.UserSearch.Phone+"%")
	}
	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&list).Error

	dbCount := global.DB.Table("user_finance_" + info.GroupNameEn).
		Joins("User")
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		dbCount = dbCount.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.UserSearch.Username != "" {
		dbCount = dbCount.Where("User.username like ?", "%"+info.UserSearch.Username+"%")
	}
	if info.UserSearch.Phone != "" {
		dbCount = dbCount.Where("User.phone like ?", "%"+info.UserSearch.Phone+"%")
	}
	err = dbCount.Count(&total).Error
	if err != nil {
		return
	}
	return list, total, err
}
