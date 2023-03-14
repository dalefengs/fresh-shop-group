package account

import (
	"fresh-shop/server/global"
	"fresh-shop/server/model/account"
	accountReq "fresh-shop/server/model/account/request"
	"fresh-shop/server/model/common/request"
)

type SysRechargeService struct {
}

// CreateSysRecharge 创建SysRecharge记录
// Author [piexlmax](https://github.com/likfees)
func (sysRechargeService *SysRechargeService) CreateSysRecharge(sysRecharge account.SysRecharge) (err error) {
	err = global.DB.Create(&sysRecharge).Error
	return err
}

// DeleteSysRecharge 删除SysRecharge记录
// Author [piexlmax](https://github.com/likfees)
func (sysRechargeService *SysRechargeService) DeleteSysRecharge(sysRecharge account.SysRecharge) (err error) {
	err = global.DB.Delete(&sysRecharge).Error
	return err
}

// DeleteSysRechargeByIds 批量删除SysRecharge记录
// Author [piexlmax](https://github.com/likfees)
func (sysRechargeService *SysRechargeService) DeleteSysRechargeByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]account.SysRecharge{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSysRecharge 更新SysRecharge记录
// Author [piexlmax](https://github.com/likfees)
func (sysRechargeService *SysRechargeService) UpdateSysRecharge(sysRecharge account.SysRecharge) (err error) {
	err = global.DB.Save(&sysRecharge).Error
	return err
}

// GetSysRecharge 根据id获取SysRecharge记录
// Author [piexlmax](https://github.com/likfees)
func (sysRechargeService *SysRechargeService) GetSysRecharge(id uint) (sysRecharge account.SysRecharge, err error) {
	err = global.DB.Where("id = ?", id).First(&sysRecharge).Error
	return
}

// GetSysRechargeInfoList 分页获取SysRecharge记录
// Author [piexlmax](https://github.com/likfees)
func (sysRechargeService *SysRechargeService) GetSysRechargeInfoList(info accountReq.SysRechargeSearch) (list []account.SysRecharge, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&account.SysRecharge{})
	var sysRecharges []account.SysRecharge
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&sysRecharges).Error
	return sysRecharges, total, err
}
