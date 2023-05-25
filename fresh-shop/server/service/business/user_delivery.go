package business

import (
	"fresh-shop/server/global"
	"fresh-shop/server/model/business"
	businessReq "fresh-shop/server/model/business/request"
	"fresh-shop/server/model/common/request"
)

type UserDeliveryService struct {
}

// CreateUserDelivery 创建UserDelivery记录
// Author [likfees](https://github.com/likfees)
func (userDeliveryService *UserDeliveryService) CreateUserDelivery(userDelivery business.UserDelivery) (err error) {
	err = global.DB.Create(&userDelivery).Error
	return err
}

// DeleteUserDelivery 删除UserDelivery记录
// Author [likfees](https://github.com/likfees)
func (userDeliveryService *UserDeliveryService) DeleteUserDelivery(userDelivery business.UserDelivery) (err error) {
	err = global.DB.Delete(&userDelivery).Error
	return err
}

// DeleteUserDeliveryByIds 批量删除UserDelivery记录
// Author [likfees](https://github.com/likfees)
func (userDeliveryService *UserDeliveryService) DeleteUserDeliveryByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.UserDelivery{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateUserDelivery 更新UserDelivery记录
// Author [likfees](https://github.com/likfees)
func (userDeliveryService *UserDeliveryService) UpdateUserDelivery(userDelivery business.UserDelivery) (err error) {
	err = global.DB.Save(&userDelivery).Error
	return err
}

// GetUserDelivery 根据id获取UserDelivery记录
// Author [likfees](https://github.com/likfees)
func (userDeliveryService *UserDeliveryService) GetUserDelivery(id uint) (userDelivery business.UserDelivery, err error) {
	err = global.DB.Where("id = ?", id).First(&userDelivery).Error
	return
}

// GetUserDeliveryInfoList 分页获取UserDelivery记录
// Author [likfees](https://github.com/likfees)
func (userDeliveryService *UserDeliveryService) GetUserDeliveryInfoList(info businessReq.UserDeliverySearch) (list []business.UserDelivery, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&business.UserDelivery{})
	var userDeliverys []business.UserDelivery
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Mobile != "" {
		db = db.Where("mobile LIKE ?", "%"+info.Mobile+"%")
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["deliverCount"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	err = db.Limit(limit).Offset(offset).Find(&userDeliverys).Error
	return userDeliverys, total, err
}
