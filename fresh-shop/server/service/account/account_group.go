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
// Author [dalefeng](https://github.com/dalefeng)
func (userAccountGroupService *AccountGroupService) CreateAccountGroup(userAccountGroup account.AccountGroup) (err error) {
	var accountGroup account.AccountGroup
	if !errors.Is(global.DB.Where("name_en = ?", userAccountGroup.NameEn).First(&accountGroup).Error, gorm.ErrRecordNotFound) {
		return errors.New("币种已存在！")
	}
	// 创建币种账户
	err = global.DB.Create(&userAccountGroup).Error
	return err
}

// ChangeSyncStatus 更新账户同步状态
func ChangeSyncStatus(groupId uint, status int) {
	// 更新同步状态
	var acg account.AccountGroup
	err := global.DB.Where("id = ?", groupId).First(&acg).Update("sync", status).Error
	if err != nil {
		global.SugarLog.Errorf("同步初始化用户账户表失败 --- 更新同步状态失败 groupId: %d, status: %d error: %s", groupId, status, err.Error())
	}
}

// SyncUserAccount 同步初始化用户账户表
func SyncUserAccount(groupId uint) {
	// 更新同步中状态
	ChangeSyncStatus(groupId, 1)
	// 查询所有没有账户的用户
	var userList []sysModel.SysUser
	err := global.DB.Table("sys_users").
		Where("id not in (?)",
			global.DB.Table("user_account").
				Select("user_id").
				Where("sys_users.id = user_account.user_id and group_id = ? and deleted_at is null", groupId)).
		Find(&userList).Error
	if err != nil {
		global.SugarLog.Infof("同步初始化用户账户 --- 查询用户失败：%d, err: %s", len(userList), err.Error())
		ChangeSyncStatus(groupId, -1)
		return
	}
	if len(userList) == 0 {
		global.SugarLog.Infof("同步初始化用户账户完成 --- 没有需要同步的用户 groupId:%d", groupId)
		ChangeSyncStatus(groupId, 2)
		return
	}
	global.SugarLog.Infof("同步初始化用户账户 --- 未同步用户数量：%d", len(userList))
	// 如果 user_account 表没有则创建
	for _, u := range userList {
		acc := account.Account{
			GroupId: &groupId,
			UserId:  &u.ID,
		}
		global.SugarLog.Infof("同步初始化用户账户 --- 创建账户，userId：%d, groupId: %d", u.ID, groupId)
		err := global.DB.Create(&acc).Error
		if err != nil {
			global.SugarLog.Error("同步初始化用户账户 --- 创建账户失败，userId：%d, groupId: %d, err:", u.ID, groupId, err.Error())
			ChangeSyncStatus(groupId, -1)
			return
		}
	}
	global.SugarLog.Infof("同步初始化用户账户 --- 初始化完成， 同步用户数为：%d", len(userList))
	ChangeSyncStatus(groupId, 2)
}

// DeleteAccountGroup 删除AccountGroup记录
// Author [dalefeng](https://github.com/dalefeng)
func (userAccountGroupService *AccountGroupService) DeleteAccountGroup(userAccountGroup account.AccountGroup) (err error) {
	err = global.DB.Delete(&userAccountGroup).Error
	return err
}

// DeleteAccountGroupByIds 批量删除AccountGroup记录
// Author [dalefeng](https://github.com/dalefeng)
func (userAccountGroupService *AccountGroupService) DeleteAccountGroupByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]account.AccountGroup{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAccountGroup 更新AccountGroup记录
// Author [dalefeng](https://github.com/dalefeng)
func (userAccountGroupService *AccountGroupService) UpdateAccountGroup(userAccountGroup account.AccountGroup) (err error) {
	err = global.DB.Save(&userAccountGroup).Error
	return err
}

// GetAccountGroup 根据id获取AccountGroup记录
// Author [dalefeng](https://github.com/dalefeng)
func (userAccountGroupService *AccountGroupService) GetAccountGroup(id uint) (userAccountGroup account.AccountGroup, err error) {
	err = global.DB.Where("id = ?", id).First(&userAccountGroup).Error
	return
}

// GetAccountGroupInfoList 分页获取AccountGroup记录
// Author [dalefeng](https://github.com/dalefeng)
func (userAccountGroupService *AccountGroupService) GetAccountGroupInfoList(info accountReq.AccountGroupSearch) (list []account.AccountGroup, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&account.AccountGroup{})
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

// GetAccountGroupInfoListAll 获取所有AccountGroup记录
// Author [dalefeng](https://github.com/dalefeng)
func (userAccountGroupService *AccountGroupService) GetAccountGroupInfoListAll(g account.AccountGroup) (list []account.AccountGroup, err error) {
	var userAccountGroups []account.AccountGroup
	db := global.DB.Model(&account.AccountGroup{})
	if g.Status != nil {
		db = db.Where("status = ?", g.Status)
	}
	err = db.Find(&userAccountGroups).Error
	return userAccountGroups, err
}

// SyncAccountGroup 同步用户账户
func (userAccountGroupService *AccountGroupService) SyncAccountGroup(group account.AccountGroup) error {
	var userAccountGroup account.AccountGroup
	err := global.DB.Where("id = ?", group.ID).First(&userAccountGroup).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("记录不存在")
	}
	if *userAccountGroup.Sync == 1 {
		return errors.New("后台同步中，请勿重复提交")
	}
	go SyncUserAccount(userAccountGroup.ID)
	return err
}
