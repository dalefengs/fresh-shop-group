package shop

import (
	"fresh-shop/server/global"
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/shop"
	shopReq "fresh-shop/server/model/shop/request"
	"fresh-shop/server/utils"
	"gorm.io/gorm"
)

type BrandCategoryService struct {
}

// CreateBrandCategory 创建BrandCategory记录
// Author [dalefeng](https://github.com/dalefeng)
func (brandCategoryService *BrandCategoryService) CreateBrandCategory(bc shopReq.BrandCategorySearch) (err error) {
	var data []shop.BrandCategory
	for _, brandId := range bc.BrandIds {
		data = append(data, shop.BrandCategory{
			CategoryId: bc.CategoryId,
			BrandId:    utils.Pointer(brandId),
		})
	}
	// 删除相关记录
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		txErr := tx.Delete(&shop.BrandCategory{}, "category_id = ?", bc.CategoryId).Error
		if txErr != nil {
			global.SugarLog.Errorf("删除 BrandCategory 相关记录失败，categoryId: %d", bc.CategoryId)
			return err
		}
		txErr = tx.Create(&data).Error
		if txErr != nil {
			global.SugarLog.Errorf("创建 BrandCategory 相关记录失败，categoryId: %d, data: %#v", bc.CategoryId, data)
			return err
		}
		return nil
	})
	return err
}

// DeleteBrandCategory 删除BrandCategory记录
// Author [dalefeng](https://github.com/dalefeng)
func (brandCategoryService *BrandCategoryService) DeleteBrandCategory(brandCategory shop.BrandCategory) (err error) {
	err = global.DB.Delete(&brandCategory).Error
	return err
}

// DeleteBrandCategoryByIds 批量删除BrandCategory记录
// Author [dalefeng](https://github.com/dalefeng)
func (brandCategoryService *BrandCategoryService) DeleteBrandCategoryByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]shop.BrandCategory{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateBrandCategory 更新BrandCategory记录
// Author [dalefeng](https://github.com/dalefeng)
func (brandCategoryService *BrandCategoryService) UpdateBrandCategory(brandCategory shop.BrandCategory) (err error) {
	err = global.DB.Save(&brandCategory).Error
	return err
}

// GetBrandCategory 根据id获取BrandCategory记录
// Author [dalefeng](https://github.com/dalefeng)
func (brandCategoryService *BrandCategoryService) GetBrandCategory(id uint) (brandCategory shop.BrandCategory, err error) {
	err = global.DB.Where("id = ?", id).First(&brandCategory).Error
	return
}

// GetBrandCategoryInfoList 分页获取BrandCategory记录
// Author [dalefeng](https://github.com/dalefeng)
func (brandCategoryService *BrandCategoryService) GetBrandCategoryInfoList(info shopReq.BrandCategorySearch) (list []shop.BrandCategory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&shop.BrandCategory{})
	var brandCategorys []shop.BrandCategory
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&brandCategorys).Error
	return brandCategorys, total, err
}
