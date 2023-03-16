package common

import (
	"errors"
	"fmt"
	"fresh-shop/server/global"
	"fresh-shop/server/model/account"
	sysModel "fresh-shop/server/model/system"
	"fresh-shop/server/utils"
	"gorm.io/gorm"
	"math"
)

// AccountUnifyDeduction 账户统一扣减
// groupId 账户类型
// finance 入账数据，需要填写完整
func AccountUnifyDeduction(groupId int, finance account.UserFinance) error {
	if finance.FeeAmount == nil {
		finance.FeeAmount = utils.Pointer(0.0)
	}
	if finance.Balance == nil {
		finance.Balance = utils.Pointer(0.0)
	}
	log := fmt.Sprintf("账户统一扣减 --- userId: %d, group: %d, typeId: %d, amount: %f, feeAmount: %f, optionType: %d",
		finance.UserId, groupId, finance.TypeId, *finance.Amount, *finance.FeeAmount, finance.OptionType)
	var user sysModel.SysUser
	if errors.Is(global.DB.Where("username = ?", finance.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		global.SugarLog.Errorf(log + " 用户不存在")
		return errors.New("用户不存在")
	}
	if user.Enable != 1 {
		global.SugarLog.Errorf(log+" 用户状态异常，user.enable: %d", user.Enable)
		return errors.New("用户状态异常")
	}
	if *finance.Amount == 0 {
		global.SugarLog.Errorf(log + " 变动数额为 0")
		return errors.New("变动数额不能为 0")
	}
	if finance.TypeId == nil || *finance.TypeId == 0 {
		global.SugarLog.Errorf(log + " 流水类型错误")
		return errors.New("流水类型错误")
	}
	// 获取账户配置
	var group account.AccountGroup
	if errors.Is(global.DB.Where("id = ?", groupId).First(&group).Error, gorm.ErrRecordNotFound) {
		global.SugarLog.Errorf(log + " 账户配置不存在")
		return errors.New("账户配置不存在")
	}
	// 获取该用户的账户信息
	accountInfo, err := GetUserAccountInfo(*finance.UserId, groupId)
	// 如果账户不存在，应该创建
	if err != nil {
		global.SugarLog.Errorf(log + " 获取账户信息失败")
		return err
	}
	if *accountInfo.Status != 1 {
		global.SugarLog.Errorf(log+" 账户异常，account.status: %d", *accountInfo.Status)
		return errors.New("账户异常")
	}
	// 手续费
	if *finance.FeeAmount > 0 {
		finance.IsFee = utils.Pointer(1)
	} else {
		finance.IsFee = utils.Pointer(0)
	}

	switch *finance.OptionType {
	case 0: // 操作余额
		if *finance.Amount < 0 && *accountInfo.Amount < *finance.Amount+*finance.FeeAmount { //  减操作且余额不足
			global.SugarLog.Errorf(log+" %s不足，当前%s为：%f", group.NameCn, group.NameCn, *accountInfo.Amount)
			return errors.New(group.NameCn + "不足")
		}
		*accountInfo.Amount += *finance.Amount
		*finance.Balance = *accountInfo.Amount
	case 1: // 操作冻结
		if *finance.Amount < 0 && *accountInfo.FreezeAmount < *finance.Amount+*finance.FeeAmount { //  减操作且余额不足
			global.SugarLog.Errorf(log+"冻结%s不足，当前%s为：%f", group.NameCn, group.NameCn, *accountInfo.Amount)
			return errors.New("冻结" + group.NameCn + "不足")
		}
		*accountInfo.FreezeAmount += *finance.Amount
		*finance.Balance = *accountInfo.FreezeAmount
	case 2: // 操作锁仓
		if *finance.Amount < 0 && *accountInfo.LockAmount < *finance.Amount+*finance.FeeAmount { //  减操作且余额不足
			global.SugarLog.Errorf(log+"锁仓%s不足，当前%s为：%f", group.NameCn, group.NameCn, *accountInfo.Amount)
			return errors.New("锁仓" + group.NameCn + "不足")
		}
		*accountInfo.LockAmount += *finance.Amount
		*finance.Balance = *accountInfo.LockAmount
	}
	// 增加累计
	if *finance.Amount > 0 {
		*accountInfo.InAmount += *finance.Amount
	} else {
		*accountInfo.OutAmount += math.Abs(*finance.Amount)
	}
	// 开始事务
	subTx := global.DB.Begin()
	// 保存用户账户数据
	err = subTx.Save(accountInfo).Error
	if err != nil {
		subTx.Rollback()
		global.SugarLog.Errorf(log+" 更新用户账户失败, accountInfo: %#v, err: %s", accountInfo, err.Error())
		return errors.New("更新用户账户失败")
	}
	// 创建流水记录
	err = subTx.Table("user_finance_" + group.NameEn).Create(&finance).Error
	if err != nil {
		subTx.Rollback()
		global.SugarLog.Errorf(log+" 创建账户流水失败, finance: %#v, err: %s", finance, err.Error())
		return errors.New("创建账户流水失败")
	}
	// TODO 增加团队累计金额
	subTx.Commit() // 提交事务
	return nil
}

// GetUserAccountInfo 获取用户币种信息
func GetUserAccountInfo(userId, groupId int) (*account.Account, error) {
	// 获取该用户的账户信息
	var userAcount account.Account
	if errors.Is(global.DB.Where("user_id = ? and group_id = ?", userId, groupId).First(&userAcount).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("账户配置不存在")
	}
	return &userAcount, nil
}
