package shop

import (
	"fresh-shop/server/global"
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/shop"
	shopReq "fresh-shop/server/model/shop/request"
)

type OrderDetailsService struct {
}

// CreateOrderDetails 创建OrderDetails记录
// Author [likfees](https://github.com/likfees)
func (orderDetailsService *OrderDetailsService) CreateOrderDetails(orderDetails shop.OrderDetails) (err error) {
	err = global.DB.Create(&orderDetails).Error
	return err
}

// DeleteOrderDetails 删除OrderDetails记录
// Author [likfees](https://github.com/likfees)
func (orderDetailsService *OrderDetailsService) DeleteOrderDetails(orderDetails shop.OrderDetails) (err error) {
	err = global.DB.Delete(&orderDetails).Error
	return err
}

// DeleteOrderDetailsByIds 批量删除OrderDetails记录
// Author [likfees](https://github.com/likfees)
func (orderDetailsService *OrderDetailsService) DeleteOrderDetailsByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]shop.OrderDetails{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateOrderDetails 更新OrderDetails记录
// Author [likfees](https://github.com/likfees)
func (orderDetailsService *OrderDetailsService) UpdateOrderDetails(orderDetails shop.OrderDetails) (err error) {
	err = global.DB.Save(&orderDetails).Error
	return err
}

// GetOrderDetails 根据id获取OrderDetails记录
// Author [likfees](https://github.com/likfees)
func (orderDetailsService *OrderDetailsService) GetOrderDetails(id uint) (orderDetails shop.OrderDetails, err error) {
	err = global.DB.Where("id = ?", id).First(&orderDetails).Error
	return
}

// GetOrderDetailsInfoList 分页获取OrderDetails记录
// Author [likfees](https://github.com/likfees)
func (orderDetailsService *OrderDetailsService) GetOrderDetailsInfoList(info shopReq.OrderDetailsSearch) (list []shop.OrderDetails, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&shop.OrderDetails{})
	var orderDetailss []shop.OrderDetails
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.GoodsName != "" {
		db = db.Where("goods_name LIKE ?", "%"+info.GoodsName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&orderDetailss).Error
	return orderDetailss, total, err
}
