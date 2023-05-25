package shop

import (
	"errors"
	"fresh-shop/server/global"
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/shop"
	shopReq "fresh-shop/server/model/shop/request"
	"gorm.io/gorm"
)

type CartService struct {
}

// CreateCart 创建Cart记录
// Author [likfees](https://github.com/likfees)
func (cartService *CartService) CreateCart(cart shop.Cart) (err error) {
	var c shop.Cart
	var goods shop.Goods
	if errors.Is(global.DB.Where("id = ?", cart.GoodsId).First(&goods).Error, gorm.ErrRecordNotFound) {
		return errors.New("商品不存在")
	}
	if *goods.Store < cart.Num {
		return errors.New("商品库存不足")
	}
	// 记录不存在则创建
	cart.SpecItemId = 0
	if errors.Is(global.DB.Where("user_id = ? and goods_id = ?", cart.UserId, cart.GoodsId).First(&c).Error, gorm.ErrRecordNotFound) {

		if cart.Num > 0 {
			err = global.DB.Create(&cart).Error
		}
	} else {
		// 如果如果数量 == 0 则删除
		if cart.Num == 0 {
			err = global.DB.Delete(&c).Error
		} else {
			c.Num = cart.Num
			err = global.DB.Save(&c).Error
		}
	}
	return err
}

// DeleteCart 删除Cart记录
// Author [likfees](https://github.com/likfees)
func (cartService *CartService) DeleteCart(cart shop.Cart) (err error) {
	err = global.DB.Delete(&cart).Error
	return err
}

// DeleteCartByIds 批量删除Cart记录
// Author [likfees](https://github.com/likfees)
func (cartService *CartService) DeleteCartByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]shop.Cart{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCart 更新Cart记录
// Author [likfees](https://github.com/likfees)
func (cartService *CartService) UpdateCart(cart shop.Cart) (err error) {
	var dbC shop.Cart
	err = global.DB.Where("id = ?", cart.ID).First(&dbC).Error
	if err != nil {
		return err
	}
	dbC.Checked = cart.Checked
	err = global.DB.Save(&dbC).Error
	return err
}

// SelectAllChecked 全选 Cart记录
// Author [likfees](https://github.com/likfees)
func (cartService *CartService) SelectAllChecked(userId uint) (err error) {
	var ids []int
	err = global.DB.Model(&shop.Cart{}).Where("user_id = ?", userId).Pluck("id", &ids).Error
	if err != nil {
		return err
	}
	err = global.DB.Model(&shop.Cart{}).Where("id in ?", ids).Update("checked", 1).Error
	return err
}

// ClearAllChecked 取消全选 Cart记录
// Author [likfees](https://github.com/likfees)
func (cartService *CartService) ClearAllChecked(userId uint) (err error) {
	var ids []int
	err = global.DB.Model(&shop.Cart{}).Where("user_id = ?", userId).Pluck("id", &ids).Error
	if err != nil {
		return err
	}
	err = global.DB.Model(&shop.Cart{}).Where("id in ?", ids).Update("checked", 0).Error
	return err
}

// GetCart 根据id获取Cart记录
// Author [likfees](https://github.com/likfees)
func (cartService *CartService) GetCart(id uint) (cart shop.Cart, err error) {
	err = global.DB.Where("id = ?", id).First(&cart).Error
	return
}

// GetCartInfoList 获取全部 Cart记录
// Author [likfees](https://github.com/likfees)
func (cartService *CartService) GetCartInfoList(info shopReq.CartSearch, userId uint) (list []shop.Cart, total int64, err error) {
	// 创建db
	db := global.DB.Model(&shop.Cart{}).Where("user_id = ?", userId).Preload("Goods.Images")
	var carts []shop.Cart
	if info.Checked != nil {
		db = db.Where("checked = ?", *info.Checked)
	}
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Order("created_at desc").Find(&carts).Error
	return carts, total, err
}
