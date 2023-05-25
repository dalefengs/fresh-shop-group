package system

import (
	"errors"
	"fresh-shop/server/global"
	businessReq "fresh-shop/server/model/business/request"
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/system"
)

type SysConfigService struct {
}

// CreateSysConfig 创建SysConfig记录
// Author [likfees](https://github.com/likfees)
func (sysConfigService *SysConfigService) CreateSysConfig(sysConfig system.SysConfig) (err error) {
	err = global.DB.Create(&sysConfig).Error
	return err
}

// DeleteSysConfig 删除SysConfig记录
// Author [likfees](https://github.com/likfees)
func (sysConfigService *SysConfigService) DeleteSysConfig(sysConfig system.SysConfig) (err error) {
	err = global.DB.Delete(&sysConfig).Error
	return err
}

// DeleteSysConfigByIds 批量删除SysConfig记录
// Author [likfees](https://github.com/likfees)
func (sysConfigService *SysConfigService) DeleteSysConfigByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]system.SysConfig{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSysConfig 更新SysConfig记录
// Author [likfees](https://github.com/likfees)
func (sysConfigService *SysConfigService) UpdateSysConfig(sysConfig system.SysConfig) (err error) {
	err = global.DB.Save(&sysConfig).Error
	return err
}

// GetSysConfig 根据id获取SysConfig记录
// Author [likfees](https://github.com/likfees)
func (sysConfigService *SysConfigService) GetSysConfig(id uint) (sysConfig system.SysConfig, err error) {
	err = global.DB.Where("id = ?", id).First(&sysConfig).Error
	return
}

// GetSysConfigByName 根据 name 获取 SysConfig 记录
// Author [likfees](https://github.com/likfees)
func (sysConfigService *SysConfigService) GetSysConfigByName(name string) (sysConfig system.SysConfig, err error) {
	err = global.DB.Where("name = ?", name).First(&sysConfig).Error
	if err != nil {
		return
	}
	if *sysConfig.Status == 0 {
		return system.SysConfig{}, errors.New("配置参数已禁用")
	}
	return
}

// GetSysConfigInfoList 分页获取SysConfig记录
// Author [likfees](https://github.com/likfees)
func (sysConfigService *SysConfigService) GetSysConfigInfoList(info businessReq.SysConfigSearch) (list []system.SysConfig, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&system.SysConfig{})
	var sysConfigs []system.SysConfig
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Value != "" {
		db = db.Where("value LIKE ?", "%"+info.Value+"%")
	}
	if info.GroupType != "" {
		db = db.Where("group_type LIKE ?", "%"+info.GroupType+"%")
	}
	if info.Desc != "" {
		db = db.Where("desc LIKE ?", "%"+info.Desc+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&sysConfigs).Error
	return sysConfigs, total, err
}
