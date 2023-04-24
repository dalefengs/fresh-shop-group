package shop

import (
	"errors"
	"fresh-shop/server/global"
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/shop"
	shopReq "fresh-shop/server/model/shop/request"
	"fresh-shop/server/utils"
	"gorm.io/gorm"
)

type CartService struct {
}

// CreateCart 创建Cart记录
// Author [piexlmax](https://github.com/likfees)
func (cartService *CartService) CreateCart(cart shop.Cart) (err error) {
	var c shop.Cart
	// 记录不存在则创建
	cart.SpecItemId = utils.Pointer(0)
	if errors.Is(global.DB.Where("user_id = ? and goods_id = ?", cart.UserId, cart.GoodsId).First(&c).Error, gorm.ErrRecordNotFound) {
		if *cart.Num > 0 {
			err = global.DB.Create(&cart).Error
		}
	} else {
		// 如果如果数量 == 0 则删除
		if *cart.Num == 0 {
			err = global.DB.Delete(&c).Error
		} else {
			c.Num = cart.Num
			err = global.DB.Save(&c).Error
		}
	}
	return err
}

// DeleteCart 删除Cart记录
// Author [piexlmax](https://github.com/likfees)
func (cartService *CartService) DeleteCart(cart shop.Cart) (err error) {
	err = global.DB.Delete(&cart).Error
	return err
}

// DeleteCartByIds 批量删除Cart记录
// Author [piexlmax](https://github.com/likfees)
func (cartService *CartService) DeleteCartByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]shop.Cart{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCart 更新Cart记录
// Author [piexlmax](https://github.com/likfees)
func (cartService *CartService) UpdateCart(cart shop.Cart) (err error) {
	err = global.DB.Save(&cart).Error
	return err
}

// GetCart 根据id获取Cart记录
// Author [piexlmax](https://github.com/likfees)
func (cartService *CartService) GetCart(id uint) (cart shop.Cart, err error) {
	err = global.DB.Where("id = ?", id).First(&cart).Error
	return
}

// GetCartInfoList 分页获取Cart记录
// Author [piexlmax](https://github.com/likfees)
func (cartService *CartService) GetCartInfoList(info shopReq.CartSearch) (list []shop.Cart, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&shop.Cart{})
	var carts []shop.Cart
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&carts).Error
	return carts, total, err
}
