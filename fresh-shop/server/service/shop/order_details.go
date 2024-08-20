package shop

import (
	"errors"
	"fresh-shop/server/global"
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/shop"
	shopReq "fresh-shop/server/model/shop/request"
	"gorm.io/gorm"
)

type OrderDetailsService struct {
}

// CreateOrderDetails 创建OrderDetails记录
// Author [dalefeng](https://github.com/dalefeng)
func (orderDetailsService *OrderDetailsService) CreateOrderDetails(orderDetails shop.OrderDetails) (err error) {
	err = global.DB.Create(&orderDetails).Error
	return err
}

// DeleteOrderDetails 删除OrderDetails记录
// Author [dalefeng](https://github.com/dalefeng)
func (orderDetailsService *OrderDetailsService) DeleteOrderDetails(orderDetails shop.OrderDetails) (err error) {
	err = global.DB.Delete(&orderDetails).Error
	return err
}

// DeleteOrderDetailsByIds 批量删除OrderDetails记录
// Author [dalefeng](https://github.com/dalefeng)
func (orderDetailsService *OrderDetailsService) DeleteOrderDetailsByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]shop.OrderDetails{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateOrderDetails 更新OrderDetails记录
// Author [dalefeng](https://github.com/dalefeng)
func (orderDetailsService *OrderDetailsService) UpdateOrderDetails(orderDetails shop.OrderDetails) (err error) {
	err = global.DB.Save(&orderDetails).Error
	return err
}

// GetOrderDetails 根据id获取OrderDetails记录
// Author [dalefeng](https://github.com/dalefeng)
func (orderDetailsService *OrderDetailsService) GetOrderDetails(id uint) (orderDetails shop.OrderDetails, err error) {
	err = global.DB.Where("id = ?", id).First(&orderDetails).Error
	return
}

// GetOrderDetailsInfoList 分页获取OrderDetails记录
// Author [dalefeng](https://github.com/dalefeng)
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

// RecentlyPurchasedGoodsList 分页获取近期购买的商品记录
// Author [dalefeng](https://github.com/dalefeng)
func (orderDetailsService *OrderDetailsService) RecentlyPurchasedGoodsList(info shopReq.GoodsSearch, userId uint) (goods []shop.Goods, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	sql := `select g.*, SUM(od.num) as pay_count from shop_order as o 
    left join shop_order_details as od on o.id = od.order_id
    left join shop_goods as g on od.goods_id = g.id
	where o.user_id = ? and g.goods_area = 0 and g.deleted_at is null`
	if info.Goods.Name != "" {
		sql += " and g.name LIKE '%" + info.Goods.Name + "%'"
	}
	countSql := `select count(1) from shop_order as o
    left join shop_order_details as od on o.id = od.order_id
    left join shop_goods as g on od.goods_id = g.id
	where o.user_id = ? and g.goods_area = 0 and g.deleted_at is null`
	if info.Goods.Name != "" {
		sql += " and g.name LIKE '%" + info.Goods.Name + "%'"
		countSql += " and g.name LIKE '%" + info.Goods.Name + "%'"
	}
	sql += " group by g.id order by g.id desc"
	countSql += " group by g.id order by g.id desc"
	err = global.DB.Raw(countSql, userId).Scan(&total).Error
	if err != nil {
		global.SugarLog.Errorw("查询失败!近期购买总数失败", "err", err)
		return
	}

	sql += " limit ? offset ?"
	err = global.DB.Raw(sql, userId, limit, offset).Scan(&goods).Error
	// 用户已经登录获取购物车数量
	if err == nil && userId > 0 {
		for key, item := range goods {
			var cart shop.Cart
			if !errors.Is(global.DB.Where("user_id = ? and goods_id = ?", userId, item.ID).First(&cart).Error, gorm.ErrRecordNotFound) {
				goods[key].GoodsCardId = cart.ID
				goods[key].CartNum = &cart.Num
			}
		}
	}

	// 获取商品图片
	for key, item := range goods {
		var goodsImage shop.GoodsImage
		if !errors.Is(global.DB.Where("goods_id = ?", item.ID).First(&goodsImage).Error, gorm.ErrRecordNotFound) {
			goods[key].Images = []shop.GoodsImage{goodsImage}
		}
	}
	return
}
