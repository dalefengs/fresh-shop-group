package shop

import (
	"fresh-shop/server/global"
	"fresh-shop/server/model/shop"
	"fresh-shop/server/model/common/request"
    shopReq "fresh-shop/server/model/shop/request"
)

type GoodsDescriptionService struct {
}

// CreateGoodsDescription 创建GoodsDescription记录
// Author [piexlmax](https://github.com/likfees)
func (goodsDescriptionService *GoodsDescriptionService) CreateGoodsDescription(goodsDescription shop.GoodsDescription) (err error) {
	err = global.DB.Create(&goodsDescription).Error
	return err
}

// DeleteGoodsDescription 删除GoodsDescription记录
// Author [piexlmax](https://github.com/likfees)
func (goodsDescriptionService *GoodsDescriptionService)DeleteGoodsDescription(goodsDescription shop.GoodsDescription) (err error) {
	err = global.DB.Delete(&goodsDescription).Error
	return err
}

// DeleteGoodsDescriptionByIds 批量删除GoodsDescription记录
// Author [piexlmax](https://github.com/likfees)
func (goodsDescriptionService *GoodsDescriptionService)DeleteGoodsDescriptionByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]shop.GoodsDescription{},"id in ?",ids.Ids).Error
	return err
}

// UpdateGoodsDescription 更新GoodsDescription记录
// Author [piexlmax](https://github.com/likfees)
func (goodsDescriptionService *GoodsDescriptionService)UpdateGoodsDescription(goodsDescription shop.GoodsDescription) (err error) {
	err = global.DB.Save(&goodsDescription).Error
	return err
}

// GetGoodsDescription 根据id获取GoodsDescription记录
// Author [piexlmax](https://github.com/likfees)
func (goodsDescriptionService *GoodsDescriptionService)GetGoodsDescription(id uint) (goodsDescription shop.GoodsDescription, err error) {
	err = global.DB.Where("id = ?", id).First(&goodsDescription).Error
	return
}

// GetGoodsDescriptionInfoList 分页获取GoodsDescription记录
// Author [piexlmax](https://github.com/likfees)
func (goodsDescriptionService *GoodsDescriptionService)GetGoodsDescriptionInfoList(info shopReq.GoodsDescriptionSearch) (list []shop.GoodsDescription, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.DB.Model(&shop.GoodsDescription{})
    var goodsDescriptions []shop.GoodsDescription
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&goodsDescriptions).Error
	return  goodsDescriptions, total, err
}
