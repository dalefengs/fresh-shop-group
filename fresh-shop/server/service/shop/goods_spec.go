package shop

import (
	"fresh-shop/server/global"
	"fresh-shop/server/model/shop"
	"fresh-shop/server/model/common/request"
    shopReq "fresh-shop/server/model/shop/request"
)

type GoodsSpecService struct {
}

// CreateGoodsSpec 创建GoodsSpec记录
// Author [piexlmax](https://github.com/likfees)
func (goodsSpecService *GoodsSpecService) CreateGoodsSpec(goodsSpec shop.GoodsSpec) (err error) {
	err = global.DB.Create(&goodsSpec).Error
	return err
}

// DeleteGoodsSpec 删除GoodsSpec记录
// Author [piexlmax](https://github.com/likfees)
func (goodsSpecService *GoodsSpecService)DeleteGoodsSpec(goodsSpec shop.GoodsSpec) (err error) {
	err = global.DB.Delete(&goodsSpec).Error
	return err
}

// DeleteGoodsSpecByIds 批量删除GoodsSpec记录
// Author [piexlmax](https://github.com/likfees)
func (goodsSpecService *GoodsSpecService)DeleteGoodsSpecByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]shop.GoodsSpec{},"id in ?",ids.Ids).Error
	return err
}

// UpdateGoodsSpec 更新GoodsSpec记录
// Author [piexlmax](https://github.com/likfees)
func (goodsSpecService *GoodsSpecService)UpdateGoodsSpec(goodsSpec shop.GoodsSpec) (err error) {
	err = global.DB.Save(&goodsSpec).Error
	return err
}

// GetGoodsSpec 根据id获取GoodsSpec记录
// Author [piexlmax](https://github.com/likfees)
func (goodsSpecService *GoodsSpecService)GetGoodsSpec(id uint) (goodsSpec shop.GoodsSpec, err error) {
	err = global.DB.Where("id = ?", id).First(&goodsSpec).Error
	return
}

// GetGoodsSpecInfoList 分页获取GoodsSpec记录
// Author [piexlmax](https://github.com/likfees)
func (goodsSpecService *GoodsSpecService)GetGoodsSpecInfoList(info shopReq.GoodsSpecSearch) (list []shop.GoodsSpec, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.DB.Model(&shop.GoodsSpec{})
    var goodsSpecs []shop.GoodsSpec
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&goodsSpecs).Error
	return  goodsSpecs, total, err
}
