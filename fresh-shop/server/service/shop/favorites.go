package shop

import (
	"errors"
	"fresh-shop/server/global"
	businessReq "fresh-shop/server/model/business/request"
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/shop"
	"gorm.io/gorm"
)

type FavoritesService struct {
}

// Favorites 创建Favorites记录
// Author [likfees](https://github.com/likfees)
func (favoritesService *FavoritesService) Favorites(favorites shop.Favorites) (err error) {
	var f shop.GoodsFavorites
	err = global.DB.Where("user_id = ? and goods_id = ?", favorites.UserId, favorites.GoodsId).First(&f).Error
	// 如果没有就保存，有就删除
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = global.DB.Save(&favorites).Error
	} else {
		err = global.DB.Delete(&f).Error
	}
	return err
}

// DeleteFavorites 删除Favorites记录
// Author [likfees](https://github.com/likfees)
func (favoritesService *FavoritesService) DeleteFavorites(favorites shop.Favorites) (err error) {
	err = global.DB.Delete(&favorites).Error
	return err
}

// DeleteFavoritesByIds 批量删除Favorites记录
// Author [likfees](https://github.com/likfees)
func (favoritesService *FavoritesService) DeleteFavoritesByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]shop.Favorites{}, "id in ?", ids.Ids).Error
	return err
}

// GetFavorites 根据id获取Favorites记录
// Author [likfees](https://github.com/likfees)
func (favoritesService *FavoritesService) GetFavorites(id uint) (favorites shop.Favorites, err error) {
	err = global.DB.Where("id = ?", id).First(&favorites).Error
	return
}

// GetFavoritesInfoList 分页获取Favorites记录
// Author [likfees](https://github.com/likfees)
func (favoritesService *FavoritesService) GetFavoritesInfoList(info businessReq.FavoritesSearch) (list []shop.Favorites, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&shop.Favorites{})
	var favoritess []shop.Favorites
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&favoritess).Error
	return favoritess, total, err
}
