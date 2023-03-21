package shop

import (
	"fresh-shop/server/global"
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/shop"
)

type GoodsImageService struct {
}

// CreateGoodsImage 创建GoodsImage记录
// Author [piexlmax](https://github.com/likfees)
func (goodsImageService *GoodsImageService) CreateGoodsImage(goodsImage shop.GoodsImage) (err error) {
	err = global.DB.Create(&goodsImage).Error
	return err
}

// DeleteGoodsImage 删除GoodsImage记录
// Author [piexlmax](https://github.com/likfees)
func (goodsImageService *GoodsImageService) DeleteGoodsImage(goodsImage shop.GoodsImage) (err error) {
	err = global.DB.Delete(&goodsImage).Error
	return err
}

// DeleteGoodsImageByIds 批量删除GoodsImage记录
// Author [piexlmax](https://github.com/likfees)
func (goodsImageService *GoodsImageService) DeleteGoodsImageByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]shop.GoodsImage{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateGoodsImage 更新GoodsImage记录
// Author [piexlmax](https://github.com/likfees)
func (goodsImageService *GoodsImageService) UpdateGoodsImage(goodsImage shop.GoodsImage) (err error) {
	err = global.DB.Save(&goodsImage).Error
	return err
}
