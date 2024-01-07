package business

import (
	"fresh-shop/server/global"
	"fresh-shop/server/model/business"
	businessReq "fresh-shop/server/model/business/request"
	"fresh-shop/server/model/common/request"
)

type BannerService struct {
}

// CreateBanner 创建Banner记录
// Author [dalefeng](https://github.com/dalefeng)
func (bannerService *BannerService) CreateBanner(banner business.Banner) (err error) {
	err = global.DB.Create(&banner).Error
	return err
}

// DeleteBanner 删除Banner记录
// Author [dalefeng](https://github.com/dalefeng)
func (bannerService *BannerService) DeleteBanner(banner business.Banner) (err error) {
	err = global.DB.Delete(&banner).Error
	return err
}

// DeleteBannerByIds 批量删除Banner记录
// Author [dalefeng](https://github.com/dalefeng)
func (bannerService *BannerService) DeleteBannerByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]business.Banner{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateBanner 更新Banner记录
// Author [dalefeng](https://github.com/dalefeng)
func (bannerService *BannerService) UpdateBanner(banner business.Banner) (err error) {
	err = global.DB.Save(&banner).Error
	return err
}

// GetBanner 根据id获取Banner记录
// Author [dalefeng](https://github.com/dalefeng)
func (bannerService *BannerService) GetBanner(id uint) (banner business.Banner, err error) {
	err = global.DB.Where("id = ?", id).First(&banner).Error
	return
}

// GetBannerInfoList 分页获取Banner记录
// Author [dalefeng](https://github.com/dalefeng)
func (bannerService *BannerService) GetBannerInfoList(info businessReq.BannerSearch) (list []business.Banner, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&business.Banner{})
	var banners []business.Banner
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Order("sort asc").Find(&banners).Error
	return banners, total, err
}
