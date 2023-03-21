package shop

import (
	"fresh-shop/server/global"
	"fresh-shop/server/model/shop"
	"fresh-shop/server/model/common/request"
    shopReq "fresh-shop/server/model/shop/request"
)

type GoodsSpecValueService struct {
}

// CreateGoodsSpecValue 创建GoodsSpecValue记录
// Author [piexlmax](https://github.com/likfees)
func (goodsSpecValueService *GoodsSpecValueService) CreateGoodsSpecValue(goodsSpecValue shop.GoodsSpecValue) (err error) {
	err = global.DB.Create(&goodsSpecValue).Error
	return err
}

// DeleteGoodsSpecValue 删除GoodsSpecValue记录
// Author [piexlmax](https://github.com/likfees)
func (goodsSpecValueService *GoodsSpecValueService)DeleteGoodsSpecValue(goodsSpecValue shop.GoodsSpecValue) (err error) {
	err = global.DB.Delete(&goodsSpecValue).Error
	return err
}

// DeleteGoodsSpecValueByIds 批量删除GoodsSpecValue记录
// Author [piexlmax](https://github.com/likfees)
func (goodsSpecValueService *GoodsSpecValueService)DeleteGoodsSpecValueByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]shop.GoodsSpecValue{},"id in ?",ids.Ids).Error
	return err
}

// UpdateGoodsSpecValue 更新GoodsSpecValue记录
// Author [piexlmax](https://github.com/likfees)
func (goodsSpecValueService *GoodsSpecValueService)UpdateGoodsSpecValue(goodsSpecValue shop.GoodsSpecValue) (err error) {
	err = global.DB.Save(&goodsSpecValue).Error
	return err
}

// GetGoodsSpecValue 根据id获取GoodsSpecValue记录
// Author [piexlmax](https://github.com/likfees)
func (goodsSpecValueService *GoodsSpecValueService)GetGoodsSpecValue(id uint) (goodsSpecValue shop.GoodsSpecValue, err error) {
	err = global.DB.Where("id = ?", id).First(&goodsSpecValue).Error
	return
}

// GetGoodsSpecValueInfoList 分页获取GoodsSpecValue记录
// Author [piexlmax](https://github.com/likfees)
func (goodsSpecValueService *GoodsSpecValueService)GetGoodsSpecValueInfoList(info shopReq.GoodsSpecValueSearch) (list []shop.GoodsSpecValue, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.DB.Model(&shop.GoodsSpecValue{})
    var goodsSpecValues []shop.GoodsSpecValue
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&goodsSpecValues).Error
	return  goodsSpecValues, total, err
}
