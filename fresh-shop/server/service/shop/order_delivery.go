package shop

import (
	"errors"
	"fresh-shop/server/global"
	"fresh-shop/server/model/business"
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/shop"
	shopReq "fresh-shop/server/model/shop/request"
	"fresh-shop/server/utils"
	"gorm.io/gorm"
)

type OrderDeliveryService struct {
}

// CreateOrderDelivery 创建OrderDelivery记录
// Author [likfees](https://github.com/likfees)
func (orderDeliveryService *OrderDeliveryService) CreateOrderDelivery(orderDelivery shop.OrderDelivery) (err error) {
	var order shop.Order
	err = global.DB.Where("id = ? and status = 1 and status_cancel = 0", orderDelivery.OrderId).First(&order).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		global.SugarLog.Errorf("获取订单信息失败 orderId:%d, error: %v", order.ID, err)
		return err
	}
	order.Status = utils.Pointer(2)
	order.ShipmentTime = &orderDelivery.ScheduledTime
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if txErr := tx.Save(&order).Error; txErr != nil {
			return txErr
		}
		err = tx.Create(&orderDelivery).Error
		if err != nil {
			return err
		}
		return nil
	})
	return
}

// DeleteOrderDelivery 删除OrderDelivery记录
// Author [likfees](https://github.com/likfees)
func (orderDeliveryService *OrderDeliveryService) DeleteOrderDelivery(orderDelivery shop.OrderDelivery) (err error) {
	err = global.DB.Delete(&orderDelivery).Error
	return err
}

// DeleteOrderDeliveryByIds 批量删除OrderDelivery记录
// Author [likfees](https://github.com/likfees)
func (orderDeliveryService *OrderDeliveryService) DeleteOrderDeliveryByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]shop.OrderDelivery{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateOrderDelivery 更新OrderDelivery记录
// Author [likfees](https://github.com/likfees)
func (orderDeliveryService *OrderDeliveryService) UpdateOrderDelivery(orderDelivery shop.OrderDelivery) (err error) {
	var order shop.Order
	err = global.DB.Where("id = ? and status = 2 and status_cancel = 0", orderDelivery.OrderId).First(&order).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		global.SugarLog.Errorf("获取订单信息失败 orderId:%d, error: %v", orderDelivery.OrderId, err)
		return err
	}
	var deliver business.UserDelivery
	if orderDelivery.DeliveryId != nil && *orderDelivery.DeliveryId > 0 {
		err = global.DB.Where("id = ?", orderDelivery.DeliveryId).First(&deliver).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			global.SugarLog.Errorf("获取送货员信息失败 DeliveryId:%d, error: %v", orderDelivery.DeliveryId, err)
			return err
		}
		deliver.DeliverCount = utils.Pointer(*deliver.DeliverCount + 1)
	}
	if orderDelivery.ReceiptTime != nil {
		order.Status = utils.Pointer(3)
		order.ReceiveTime = orderDelivery.ReceiptTime
	}
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if txErr := tx.Save(&order).Error; txErr != nil {
			return txErr
		}
		// 保存收货人信息
		if orderDelivery.DeliveryId != nil && *orderDelivery.DeliveryId > 0 {
			if txErr := tx.Save(&deliver).Error; txErr != nil {
				return txErr
			}
		}
		err = tx.Save(&orderDelivery).Error
		if err != nil {
			return err
		}
		return nil
	})
	return
}

// GetOrderDelivery 根据id获取OrderDelivery记录
// Author [likfees](https://github.com/likfees)
func (orderDeliveryService *OrderDeliveryService) GetOrderDelivery(id uint, orderId int) (orderDelivery shop.OrderDelivery, err error) {
	db := global.DB
	if id != 0 {
		db = db.Where("id = ?", id)
	} else if orderId != 0 {
		db = db.Where("order_id = ?", orderId)
	} else {
		return orderDelivery, errors.New("参数异常")
	}
	err = db.First(&orderDelivery).Error
	return
}

// GetOrderDeliveryInfoList 分页获取OrderDelivery记录
// Author [likfees](https://github.com/likfees)
func (orderDeliveryService *OrderDeliveryService) GetOrderDeliveryInfoList(info shopReq.OrderDeliverySearch) (list []shop.OrderDelivery, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&shop.OrderDelivery{})
	var orderDeliverys []shop.OrderDelivery
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.StartScheduledTime != nil && info.EndScheduledTime != nil {
		db = db.Where("scheduled_time BETWEEN ? AND ? ", info.StartScheduledTime, info.EndScheduledTime)
	}
	if info.DeliverName != "" {
		db = db.Where("deliver_name LIKE ?", "%"+info.DeliverName+"%")
	}
	if info.DeliverMobile != "" {
		db = db.Where("deliver_mobile LIKE ?", "%"+info.DeliverMobile+"%")
	}
	if info.StartReceiptTime != nil && info.EndReceiptTime != nil {
		db = db.Where("receipt_time BETWEEN ? AND ? ", info.StartReceiptTime, info.EndReceiptTime)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&orderDeliverys).Error
	return orderDeliverys, total, err
}
