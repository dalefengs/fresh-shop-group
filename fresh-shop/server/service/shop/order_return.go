package shop

import (
	"fresh-shop/server/global"
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/shop"
	shopReq "fresh-shop/server/model/shop/request"
)

type OrderReturnService struct {
}

// CreateOrderReturn 创建OrderReturn记录
// Author [dalefeng](https://github.com/dalefeng)
func (orderReturnService *OrderReturnService) CreateOrderReturn(orderReturn shop.OrderReturn) (err error) {
	err = global.DB.Create(&orderReturn).Error
	return err
}

// DeleteOrderReturn 删除OrderReturn记录
// Author [dalefeng](https://github.com/dalefeng)
func (orderReturnService *OrderReturnService) DeleteOrderReturn(orderReturn shop.OrderReturn) (err error) {
	err = global.DB.Delete(&orderReturn).Error
	return err
}

// DeleteOrderReturnByIds 批量删除OrderReturn记录
// Author [dalefeng](https://github.com/dalefeng)
func (orderReturnService *OrderReturnService) DeleteOrderReturnByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]shop.OrderReturn{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateOrderReturn 更新OrderReturn记录
// Author [dalefeng](https://github.com/dalefeng)
func (orderReturnService *OrderReturnService) UpdateOrderReturn(orderReturn shop.OrderReturn) (err error) {
	err = global.DB.Save(&orderReturn).Error
	return err
}

// GetOrderReturn 根据id获取OrderReturn记录
// Author [dalefeng](https://github.com/dalefeng)
func (orderReturnService *OrderReturnService) GetOrderReturn(id uint) (orderReturn shop.OrderReturn, err error) {
	err = global.DB.Where("id = ?", id).First(&orderReturn).Error
	return
}

// GetOrderReturnInfoList 分页获取OrderReturn记录
// Author [dalefeng](https://github.com/dalefeng)
func (orderReturnService *OrderReturnService) GetOrderReturnInfoList(info shopReq.OrderReturnSearch) (list []shop.OrderReturn, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&shop.OrderReturn{})
	var orderReturns []shop.OrderReturn
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.RefundStatus != nil {
		db = db.Where("refund_status = ?", info.RefundStatus)
	}
	if info.StartProcessTime != nil && info.EndProcessTime != nil {
		db = db.Where("process_time BETWEEN ? AND ? ", info.StartProcessTime, info.EndProcessTime)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&orderReturns).Error
	return orderReturns, total, err
}
