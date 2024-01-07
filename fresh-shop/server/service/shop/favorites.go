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
// Author [dalefeng](https://github.com/dalefeng)
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
// Author [dalefeng](https://github.com/dalefeng)
func (favoritesService *FavoritesService) DeleteFavorites(favorites shop.Favorites) (err error) {
	err = global.DB.Delete(&favorites).Error
	return err
}

// DeleteFavoritesByIds 批量删除Favorites记录
// Author [dalefeng](https://github.com/dalefeng)
func (favoritesService *FavoritesService) DeleteFavoritesByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]shop.Favorites{}, "id in ?", ids.Ids).Error
	return err
}

// GetFavorites 根据id获取Favorites记录
// Author [dalefeng](https://github.com/dalefeng)
func (favoritesService *FavoritesService) GetFavorites(id uint) (favorites shop.Favorites, err error) {
	err = global.DB.Where("id = ?", id).First(&favorites).Error
	return
}

//// GetFavoritesInfoList 分页获取Favorites记录
//// Author [dalefeng](https://github.com/dalefeng)
//func (favoritesService *FavoritesService) GetFavoritesInfoList(info businessReq.FavoritesSearch) (list []shop.Favorites, total int64, err error) {
//	limit := info.PageSize
//	offset := info.PageSize * (info.Page - 1)
//	// 创建db
//	db := global.DB.Model(&shop.Favorites{}).Preload("Goods")
//	var favoritess []shop.Favorites
//	// 如果有条件搜索 下方会自动创建搜索语句
//	if info.UserId != nil && *info.UserId != 0 {
//		db = db.Where("user_id = ?", info.UserId)
//	}
//	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
//		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
//	}
//	err = db.Count(&total).Error
//	if err != nil {
//		return
//	}
//
//	err = db.Limit(limit).Offset(offset).Find(&favoritess).Error
//	return favoritess, total, err
//}

// GetFavoritesInfoList 分页获取Favorites记录
// Author [dalefeng](https://github.com/dalefeng)
func (favoritesService *FavoritesService) GetFavoritesInfoList(info businessReq.FavoritesSearch) (goods []shop.Goods, total int64, err error) {
	userId := *info.UserId
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	if info.PageSize == 0 {
		info.PageSize = 10
	}
	if info.Page == 0 {
		info.Page = 1
	}

	countSql := `SELECT count(1) FROM shop_favorites f LEFT JOIN shop_goods g ON f.goods_id = g.id WHERE f.user_id = ? and f.deleted_at is null and g.deleted_at is null`
	if info.Name != "" {
		countSql += " AND g.name LIKE '%" + info.Name + "%'"
	}

	sql := `SELECT g.*, f.id as pay_count FROM shop_favorites f LEFT JOIN shop_goods g ON f.goods_id = g.id WHERE f.user_id = ? and f.deleted_at is null and g.deleted_at is null`
	if info.Name != "" {
		sql += " AND g.name LIKE '%" + info.Name + "%'"
	}
	sql = sql + " ORDER BY f.created_at DESC LIMIT ? OFFSET ?"

	err = global.DB.Raw(countSql, info.UserId).Scan(&total).Error
	if err != nil {
		return
	}

	err = global.DB.Raw(sql, info.UserId, limit, offset).Scan(&goods).Error
	if err != nil {
		return
	}
	// 用户已经登录获取购物车数量
	if err == nil && userId > 0 {
		for key, item := range goods {
			var cart shop.Cart
			if !errors.Is(global.DB.Where("user_id = ? and goods_id = ?", userId, item.ID).First(&cart).Error, gorm.ErrRecordNotFound) {
				goods[key].GoodsCardId = cart.ID
				goods[key].CartNum = &cart.Num
			}
		}
	}

	// 获取商品图片
	for key, item := range goods {
		var goodsImage shop.GoodsImage
		if !errors.Is(global.DB.Where("goods_id = ?", item.ID).First(&goodsImage).Error, gorm.ErrRecordNotFound) {
			goods[key].Images = []shop.GoodsImage{goodsImage}
		}
	}

	return
}
