package shop

import (
	"fresh-shop/server/global"
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/shop"
	shopReq "fresh-shop/server/model/shop/request"
)

type BrandService struct {
}

// CreateBrand 创建Brand记录
// Author [piexlmax](https://github.com/likfees)
func (brandService *BrandService) CreateBrand(brand shop.Brand) (err error) {
	err = global.DB.Create(&brand).Error
	return err
}

// DeleteBrand 删除Brand记录
// Author [piexlmax](https://github.com/likfees)
func (brandService *BrandService) DeleteBrand(brand shop.Brand) (err error) {
	err = global.DB.Delete(&brand).Error
	return err
}

// DeleteBrandByIds 批量删除Brand记录
// Author [piexlmax](https://github.com/likfees)
func (brandService *BrandService) DeleteBrandByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]shop.Brand{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateBrand 更新Brand记录
// Author [piexlmax](https://github.com/likfees)
func (brandService *BrandService) UpdateBrand(brand shop.Brand) (err error) {
	err = global.DB.Save(&brand).Error
	return err
}

// GetBrand 根据id获取Brand记录
// Author [piexlmax](https://github.com/likfees)
func (brandService *BrandService) GetBrand(id uint) (brand shop.Brand, err error) {
	err = global.DB.Where("id = ?", id).First(&brand).Error
	return
}

// GetBrandInfoList 分页获取Brand记录
// Author [piexlmax](https://github.com/likfees)
func (brandService *BrandService) GetBrandInfoList(info shopReq.BrandSearch) (list []shop.Brand, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&shop.Brand{})
	var brands []shop.Brand
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Order("sort asc").Find(&brands).Error
	return brands, total, err
}

// GetBrandListByCategoryId 分页获取Brand记录
// Author [piexlmax](https://github.com/likfees)
func (brandService *BrandService) GetBrandListByCategoryId(info shopReq.BrandSearch) (list []shop.Brand, err error) {
	var brandIds []int64
	if info.CategoryId > 0 {
		global.DB.Model(&shop.BrandCategory{}).Where("category_id = ?", info.CategoryId).Pluck("brand_id", &brandIds)
	}

	// 创建db
	db := global.DB.Model(&shop.Brand{})
	if info.CategoryId > 0 && len(brandIds) == 0 {
		return
	} else if info.CategoryId > 0 && len(brandIds) > 0 {
		db = db.Where("id in (?)", brandIds)
	}
	var brands []shop.Brand
	if err != nil {
		return
	}
	err = db.Order("sort asc").Find(&brands).Error
	return brands, err
}

// GetBrandInfoListAll 获取所有Brand列表
// Author [piexlmax](https://github.com/likfees)
func (brandService *BrandService) GetBrandInfoListAll() (list []shop.Brand, err error) {
	var brands []shop.Brand
	err = global.DB.Model(&shop.Brand{}).Find(&brands).Error
	return brands, err
}
