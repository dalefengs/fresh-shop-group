package account

import (
	"fresh-shop/server/global"
	"fresh-shop/server/model/account"
	accountReq "fresh-shop/server/model/account/request"
	"fresh-shop/server/model/common/request"
)

type UserFinanceTypeService struct {
}

// CreateUserFinanceType 创建UserFinanceType记录
// Author [likfees](https://github.com/likfees)
func (userFinanceTypeService *UserFinanceTypeService) CreateUserFinanceType(userFinanceType account.UserFinanceType) (err error) {
	err = global.DB.Create(&userFinanceType).Error
	return err
}

// DeleteUserFinanceType 删除UserFinanceType记录
// Author [likfees](https://github.com/likfees)
func (userFinanceTypeService *UserFinanceTypeService) DeleteUserFinanceType(userFinanceType account.UserFinanceType) (err error) {
	err = global.DB.Delete(&userFinanceType).Error
	return err
}

// DeleteUserFinanceTypeByIds 批量删除UserFinanceType记录
// Author [likfees](https://github.com/likfees)
func (userFinanceTypeService *UserFinanceTypeService) DeleteUserFinanceTypeByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]account.UserFinanceType{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateUserFinanceType 更新UserFinanceType记录
// Author [likfees](https://github.com/likfees)
func (userFinanceTypeService *UserFinanceTypeService) UpdateUserFinanceType(userFinanceType account.UserFinanceType) (err error) {
	err = global.DB.Save(&userFinanceType).Error
	return err
}

// GetUserFinanceType 根据id获取UserFinanceType记录
// Author [likfees](https://github.com/likfees)
func (userFinanceTypeService *UserFinanceTypeService) GetUserFinanceType(id uint) (userFinanceType account.UserFinanceType, err error) {
	err = global.DB.Where("id = ?", id).First(&userFinanceType).Error
	return
}

// GetUserFinanceTypeInfoList 分页获取UserFinanceType记录
// Author [likfees](https://github.com/likfees)
func (userFinanceTypeService *UserFinanceTypeService) GetUserFinanceTypeInfoList(info accountReq.UserFinanceTypeSearch) (list []account.UserFinanceType, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&account.UserFinanceType{}).
		Where("parent_id = 0").
		Preload("Children")
	var userFinanceTypes []account.UserFinanceType
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&userFinanceTypes).Error
	return userFinanceTypes, total, err
}

// GetUserFinanceTypeInfoListAll 获取所有父级数据
func (userFinanceTypeService *UserFinanceTypeService) GetUserFinanceTypeInfoListAll() ([]account.UserFinanceType, error) {
	var userFinanceTypes []account.UserFinanceType
	err := global.DB.Model(&account.UserFinanceType{}).Where("parent_id = 0").Find(&userFinanceTypes).Error
	return userFinanceTypes, err
}
