package shop

import (
	"fresh-shop/server/global"
	"fresh-shop/server/model/shop"
	"fresh-shop/server/model/common/request"
    shopReq "fresh-shop/server/model/shop/request"
)

type OrderDeliveryService struct {
}

// CreateOrderDelivery 创建OrderDelivery记录
// Author [piexlmax](https://github.com/likfees)
func (orderDeliveryService *OrderDeliveryService) CreateOrderDelivery(orderDelivery shop.OrderDelivery) (err error) {
	err = global.DB.Create(&orderDelivery).Error
	return err
}

// DeleteOrderDelivery 删除OrderDelivery记录
// Author [piexlmax](https://github.com/likfees)
func (orderDeliveryService *OrderDeliveryService)DeleteOrderDelivery(orderDelivery shop.OrderDelivery) (err error) {
	err = global.DB.Delete(&orderDelivery).Error
	return err
}

// DeleteOrderDeliveryByIds 批量删除OrderDelivery记录
// Author [piexlmax](https://github.com/likfees)
func (orderDeliveryService *OrderDeliveryService)DeleteOrderDeliveryByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]shop.OrderDelivery{},"id in ?",ids.Ids).Error
	return err
}

// UpdateOrderDelivery 更新OrderDelivery记录
// Author [piexlmax](https://github.com/likfees)
func (orderDeliveryService *OrderDeliveryService)UpdateOrderDelivery(orderDelivery shop.OrderDelivery) (err error) {
	err = global.DB.Save(&orderDelivery).Error
	return err
}

// GetOrderDelivery 根据id获取OrderDelivery记录
// Author [piexlmax](https://github.com/likfees)
func (orderDeliveryService *OrderDeliveryService)GetOrderDelivery(id uint) (orderDelivery shop.OrderDelivery, err error) {
	err = global.DB.Where("id = ?", id).First(&orderDelivery).Error
	return
}

// GetOrderDeliveryInfoList 分页获取OrderDelivery记录
// Author [piexlmax](https://github.com/likfees)
func (orderDeliveryService *OrderDeliveryService)GetOrderDeliveryInfoList(info shopReq.OrderDeliverySearch) (list []shop.OrderDelivery, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.DB.Model(&shop.OrderDelivery{})
    var orderDeliverys []shop.OrderDelivery
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
        if info.StartScheduledTime != nil && info.EndScheduledTime != nil {
            db = db.Where("scheduled_time BETWEEN ? AND ? ",info.StartScheduledTime,info.EndScheduledTime)
        }
    if info.DeliverName != "" {
        db = db.Where("deliver_name LIKE ?","%"+ info.DeliverName+"%")
    }
    if info.DeliverMobile != "" {
        db = db.Where("deliver_mobile LIKE ?","%"+ info.DeliverMobile+"%")
    }
        if info.StartReceiptTime != nil && info.EndReceiptTime != nil {
            db = db.Where("receipt_time BETWEEN ? AND ? ",info.StartReceiptTime,info.EndReceiptTime)
        }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&orderDeliverys).Error
	return  orderDeliverys, total, err
}
