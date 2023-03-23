package shop

import (
	"fresh-shop/server/global"
	"fresh-shop/server/model/shop"
	"fresh-shop/server/model/common/request"
    shopReq "fresh-shop/server/model/shop/request"
)

type OrderService struct {
}

// CreateOrder 创建Order记录
// Author [piexlmax](https://github.com/likfees)
func (orderService *OrderService) CreateOrder(order shop.Order) (err error) {
	err = global.DB.Create(&order).Error
	return err
}

// DeleteOrder 删除Order记录
// Author [piexlmax](https://github.com/likfees)
func (orderService *OrderService)DeleteOrder(order shop.Order) (err error) {
	err = global.DB.Delete(&order).Error
	return err
}

// DeleteOrderByIds 批量删除Order记录
// Author [piexlmax](https://github.com/likfees)
func (orderService *OrderService)DeleteOrderByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]shop.Order{},"id in ?",ids.Ids).Error
	return err
}

// UpdateOrder 更新Order记录
// Author [piexlmax](https://github.com/likfees)
func (orderService *OrderService)UpdateOrder(order shop.Order) (err error) {
	err = global.DB.Save(&order).Error
	return err
}

// GetOrder 根据id获取Order记录
// Author [piexlmax](https://github.com/likfees)
func (orderService *OrderService)GetOrder(id uint) (order shop.Order, err error) {
	err = global.DB.Where("id = ?", id).First(&order).Error
	return
}

// GetOrderInfoList 分页获取Order记录
// Author [piexlmax](https://github.com/likfees)
func (orderService *OrderService)GetOrderInfoList(info shopReq.OrderSearch) (list []shop.Order, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.DB.Model(&shop.Order{})
    var orders []shop.Order
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.OrderSn != "" {
        db = db.Where("order_sn LIKE ?","%"+ info.OrderSn+"%")
    }
    if info.ShipmentName != "" {
        db = db.Where("shipment_name LIKE ?","%"+ info.ShipmentName+"%")
    }
    if info.ShipmentMobile != "" {
        db = db.Where("shipment_mobile LIKE ?","%"+ info.ShipmentMobile+"%")
    }
    if info.ShipingAddress != "" {
        db = db.Where("shiping_address LIKE ?","%"+ info.ShipingAddress+"%")
    }
    if info.Payment != nil {
        db = db.Where("payment = ?",info.Payment)
    }
    if info.Status != nil {
        db = db.Where("status = ?",info.Status)
    }
        if info.StartShipmentTime != nil && info.EndShipmentTime != nil {
            db = db.Where("shipment_time BETWEEN ? AND ? ",info.StartShipmentTime,info.EndShipmentTime)
        }
        if info.StartReceiveTime != nil && info.EndReceiveTime != nil {
            db = db.Where("receive_time BETWEEN ? AND ? ",info.StartReceiveTime,info.EndReceiveTime)
        }
        if info.StartCancelTime != nil && info.EndCancelTime != nil {
            db = db.Where("cancel_time BETWEEN ? AND ? ",info.StartCancelTime,info.EndCancelTime)
        }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&orders).Error
	return  orders, total, err
}
