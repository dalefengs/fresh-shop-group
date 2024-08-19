package shop

import (
	"fresh-shop/server/global"
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/shop"
	shopReq "fresh-shop/server/model/shop/request"
)

type ShopFavoritesService struct {
}

// CreateShopFavorites 创建ShopFavorites记录
// Author [piexlmax](https://github.com/dalefeng)
func (shopFavoritesService *ShopFavoritesService) CreateShopFavorites(shopFavorites shop.ShopFavorites) (err error) {
	err = global.DB.Create(&shopFavorites).Error
	return err
}

// DeleteShopFavorites 删除ShopFavorites记录
// Author [piexlmax](https://github.com/dalefeng)
func (shopFavoritesService *ShopFavoritesService) DeleteShopFavorites(shopFavorites shop.ShopFavorites) (err error) {
	err = global.DB.Delete(&shopFavorites).Error
	return err
}

// DeleteShopFavoritesByIds 批量删除ShopFavorites记录
// Author [piexlmax](https://github.com/dalefeng)
func (shopFavoritesService *ShopFavoritesService) DeleteShopFavoritesByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]shop.ShopFavorites{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateShopFavorites 更新ShopFavorites记录
// Author [piexlmax](https://github.com/dalefeng)
func (shopFavoritesService *ShopFavoritesService) UpdateShopFavorites(shopFavorites shop.ShopFavorites) (err error) {
	err = global.DB.Save(&shopFavorites).Error
	return err
}

// GetShopFavorites 根据id获取ShopFavorites记录
// Author [piexlmax](https://github.com/dalefeng)
func (shopFavoritesService *ShopFavoritesService) GetShopFavorites(id uint) (shopFavorites shop.ShopFavorites, err error) {
	err = global.DB.Where("id = ?", id).First(&shopFavorites).Error
	return
}

// GetShopFavoritesInfoList 分页获取ShopFavorites记录
// Author [piexlmax](https://github.com/dalefeng)
func (shopFavoritesService *ShopFavoritesService) GetShopFavoritesInfoList(info shopReq.ShopFavoritesSearch) (list []shop.ShopFavorites, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&shop.ShopFavorites{})
	var shopFavoritess []shop.ShopFavorites
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&shopFavoritess).Error
	return shopFavoritess, total, err
}
