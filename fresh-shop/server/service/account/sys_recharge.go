package account

import (
	"errors"
	"fresh-shop/server/global"
	"fresh-shop/server/model/account"
	accountReq "fresh-shop/server/model/account/request"
	"fresh-shop/server/model/common/request"
	sysModel "fresh-shop/server/model/system"
	request2 "fresh-shop/server/model/system/request"
	"fresh-shop/server/service/common"
	"fresh-shop/server/utils"
	"gorm.io/gorm"
)

type SysRechargeService struct {
}

// CreateSysRecharge 创建SysRecharge记录
// 还未支持 冻结，锁定余额操作
// Author [piexlmax](https://github.com/likfees)
func (sysRechargeService *SysRechargeService) CreateSysRecharge(recharge account.SysRecharge, claims *request2.CustomClaims) (err error) {
	var user sysModel.SysUser
	if errors.Is(global.DB.Where("username = ?", recharge.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户不存在，请检查用户名")
	}
	var acInfo account.Account
	if errors.Is(global.DB.Where("user_id = ? and group_id = ?", user.ID, recharge.GroupId).First(&acInfo).Error, gorm.ErrRecordNotFound) {
		return errors.New("账户不存在")
	}
	finance := account.UserFinance{
		TypeId:     utils.Pointer(8),
		Username:   user.Username,
		UserId:     utils.Pointer(int(user.ID)),
		OptionType: utils.Pointer(0), // 余额操作
		Amount:     recharge.Amount,
		Remarks:    recharge.Remarks,
		FromId:     claims.Id,
		FromName:   claims.Username,
		FromUserId: utils.Pointer(int(claims.ID)),
	}
	recharge.UserId = utils.Pointer(int(user.ID))
	// 获取管理员身份
	recharge.AdminId = utils.Pointer(int(claims.ID))
	recharge.AdminName = claims.Username
	total := *acInfo.Amount + *recharge.Amount
	recharge.Balance = &total
	tx := global.DB.Begin()
	// 检查点
	tx.SavePoint("recharge")
	// 插入充值记录
	err = tx.Create(&recharge).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// 生成流水
	err = common.AccountUnifyDeduction(*recharge.GroupId, finance)
	if err != nil {
		// 回滚到检查点
		tx.RollbackTo("recharge")
		return err
	}
	tx.Commit()
	return nil
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
	db := global.DB.Model(&account.SysRecharge{}).
		Joins("User").
		Joins("Account").
		Preload("Group").
		Where("Account.group_id = sys_recharge.group_id")
	var sysRecharges []account.SysRecharge
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
	err = db.Limit(limit).Offset(offset).Order("created_at desc").Find(&sysRecharges).Error
	if err != nil {
		return
	}
	err = db.Count(&total).Error
	return sysRecharges, total, err
}
