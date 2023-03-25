package shop

import (
	"fresh-shop/server/global"
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/shop"
	shopReq "fresh-shop/server/model/shop/request"
)

type TagsService struct {
}

// CreateTags 创建Tags记录
// Author [piexlmax](https://github.com/likfees)
func (tagsService *TagsService) CreateTags(tags shop.Tags) (err error) {
	err = global.DB.Create(&tags).Error
	return err
}

// DeleteTags 删除Tags记录
// Author [piexlmax](https://github.com/likfees)
func (tagsService *TagsService) DeleteTags(tags shop.Tags) (err error) {
	err = global.DB.Delete(&tags).Error
	return err
}

// DeleteTagsByIds 批量删除Tags记录
// Author [piexlmax](https://github.com/likfees)
func (tagsService *TagsService) DeleteTagsByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]shop.Tags{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTags 更新Tags记录
// Author [piexlmax](https://github.com/likfees)
func (tagsService *TagsService) UpdateTags(tags shop.Tags) (err error) {
	err = global.DB.Save(&tags).Error
	return err
}

// GetTags 根据id获取Tags记录
// Author [piexlmax](https://github.com/likfees)
func (tagsService *TagsService) GetTags(id uint) (tags shop.Tags, err error) {
	err = global.DB.Where("id = ?", id).First(&tags).Error
	return
}

// GetTagsInfoList 分页获取Tags记录
// Author [piexlmax](https://github.com/likfees)
func (tagsService *TagsService) GetTagsInfoList(info shopReq.TagsSearch) (list []shop.Tags, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&shop.Tags{})
	var tagss []shop.Tags
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	OrderStr := "sort asc"
	orderMap := make(map[string]bool)
	orderMap["sort"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	err = db.Limit(limit).Offset(offset).Find(&tagss).Error
	return tagss, total, err
}
