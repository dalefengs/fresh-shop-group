package shop

import (
	"fresh-shop/server/global"
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/shop"
	shopReq "fresh-shop/server/model/shop/request"
)

type CategoryService struct {
}

// CreateCategory 创建Category记录
// Author [piexlmax](https://github.com/likfees)
func (categoryService *CategoryService) CreateCategory(category shop.Category) (err error) {
	err = global.DB.Create(&category).Error
	return err
}

// DeleteCategory 删除Category记录
// Author [piexlmax](https://github.com/likfees)
func (categoryService *CategoryService) DeleteCategory(category shop.Category) (err error) {
	err = global.DB.Delete(&category).Error
	return err
}

// DeleteCategoryByIds 批量删除Category记录
// Author [piexlmax](https://github.com/likfees)
func (categoryService *CategoryService) DeleteCategoryByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]shop.Category{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCategory 更新Category记录
// Author [piexlmax](https://github.com/likfees)
func (categoryService *CategoryService) UpdateCategory(category shop.Category) (err error) {
	err = global.DB.Save(&category).Error
	return err
}

// GetCategory 根据id获取Category记录
// Author [piexlmax](https://github.com/likfees)
func (categoryService *CategoryService) GetCategory(id uint) (category shop.Category, err error) {
	err = global.DB.Where("id = ?", id).First(&category).Error
	return
}

// GetCategoryInfoList 分页获取Category记录
// Author [piexlmax](https://github.com/likfees)
func (categoryService *CategoryService) GetCategoryInfoList(info shopReq.CategorySearch) (list []shop.Category, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	var categorys []shop.Category
	// 创建db
	db := global.DB.Model(&shop.Category{}).Preload("Brands")
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&categorys).Error
	return categorys, total, err
}

// GetCategoryInfoListAll 获取所有 Category列表
// Author [piexlmax](https://github.com/likfees)
func (categoryService *CategoryService) GetCategoryInfoListAll() (list []shop.Category, err error) {
	var categorys []shop.Category
	err = global.DB.Model(&shop.Category{}).Preload("Brands").Find(&categorys).Error
	return categorys, err
}
