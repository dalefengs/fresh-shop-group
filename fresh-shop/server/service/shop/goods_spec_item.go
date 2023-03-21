package shop

import (
	"fresh-shop/server/global"
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/shop"
)

type GoodsSpecItemService struct {
}

// CreateGoodsSpecItem 创建GoodsSpecItem记录
// Author [piexlmax](https://github.com/likfees)
func (goodsSpecItemService *GoodsSpecItemService) CreateGoodsSpecItem(goodsSpecItem shop.GoodsSpecItem) (err error) {
	err = global.DB.Create(&goodsSpecItem).Error
	return err
}

// DeleteGoodsSpecItem 删除GoodsSpecItem记录
// Author [piexlmax](https://github.com/likfees)
func (goodsSpecItemService *GoodsSpecItemService) DeleteGoodsSpecItem(goodsSpecItem shop.GoodsSpecItem) (err error) {
	err = global.DB.Delete(&goodsSpecItem).Error
	return err
}

// DeleteGoodsSpecItemByIds 批量删除GoodsSpecItem记录
// Author [piexlmax](https://github.com/likfees)
func (goodsSpecItemService *GoodsSpecItemService) DeleteGoodsSpecItemByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]shop.GoodsSpecItem{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateGoodsSpecItem 更新GoodsSpecItem记录
// Author [piexlmax](https://github.com/likfees)
func (goodsSpecItemService *GoodsSpecItemService) UpdateGoodsSpecItem(goodsSpecItem shop.GoodsSpecItem) (err error) {
	err = global.DB.Save(&goodsSpecItem).Error
	return err
}

// GetGoodsSpecItem 根据id获取GoodsSpecItem记录
// Author [piexlmax](https://github.com/likfees)
func (goodsSpecItemService *GoodsSpecItemService) GetGoodsSpecItem(id uint) (goodsSpecItem shop.GoodsSpecItem, err error) {
	err = global.DB.Where("id = ?", id).First(&goodsSpecItem).Error
	return
}
