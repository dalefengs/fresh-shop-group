package shop

import (
	"errors"
	"fmt"
	"fresh-shop/server/global"
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/shop"
	shopReq "fresh-shop/server/model/shop/request"
	"fresh-shop/server/utils"
	"strconv"
	"strings"
)

type GoodsService struct {
}

// CreateGoods 创建Goods记录
// Author [piexlmax](https://github.com/likfees)
func (goodsService *GoodsService) CreateGoods(from shopReq.GoodsSubmitFrom) (err error) {
	log := "创建商品 --- "

	goods := from.GoodsInfo
	var spec []shop.GoodsSpec
	var specItem []shop.GoodsSpecItem
	var specValue []shop.GoodsSpecValue
	goodsDesc := shop.GoodsDescription{ // 商品详情
		Details: from.Desc.Details,
		Notice:  from.Desc.Notice,
	}

	// 开始事物
	tx := global.DB.Begin()

	// 创建商品基本信息
	if err := tx.Create(&goods).Error; err != nil {
		tx.Rollback()
		global.SugarLog.Errorf(log+" 创建商品信息失败 goodsInfo: %#v, err: %s", goods, err.Error())
		return errors.New("创建商品信息失败")
	}
	goodsIdPointr := utils.Pointer(int(goods.ID))
	// 创建商品详情信息
	goodsDesc.GoodsId = goodsIdPointr
	if err := tx.Create(&goodsDesc).Error; err != nil {
		tx.Rollback()
		global.SugarLog.Errorf(log+" 创建商品详情信息失败 goodsDesc: %#v err: %s", goodsDesc, err.Error())
		return errors.New("创建商品详情信息失败")
	}

	// 创建商品图片信息
	if len(from.Images) > 0 {
		imgs := from.Images
		for k, _ := range imgs {
			imgs[k].GoodsId = goodsIdPointr
		}
		if err := tx.Create(&imgs).Error; err != nil {
			tx.Rollback()
			global.SugarLog.Errorf(log+" 创建商品图片列表失败 imgs: %#v, err: %s", imgs, err.Error())
			return errors.New("创建商品图片列表失败")
		}
	}

	// 多规格
	if *from.GoodsInfo.SpecType == 1 {
		// 创建规格
		// 处理规格项数据
		for _, s := range from.Spec {
			spec = append(spec, shop.GoodsSpec{
				GoodsId:       int(goods.ID),
				Title:         s.Name,
				IsUploadImage: s.IsUploadImage,
				Sort:          50,
				SpecId:        s.SpecId,
			})
		}
		if err := tx.Create(&spec).Error; err != nil {
			tx.Rollback()
			global.SugarLog.Errorf(log+"spec: %#v, err: %s", spec, err.Error())
			return errors.New("创建商品规格失败")
		}
		// 创建规格项
		// 处理规格项数据
		for _, item := range from.SpecItem {
			// 查找数据中的规格Id
			specId := findSpecIdByVirtualSpecId(&spec, item.SpecId)
			if specId == 0 {
				global.SugarLog.Errorf(log+"处理规格项数据失败, 查找 SpecId 异常, err: %s", err.Error())
				return errors.New("处理规格项数据失败")
			}
			specItem = append(specItem, shop.GoodsSpecItem{
				SpecId: &specId, // 此时记录的是数据库中真实的SpecId
				ImgUrl: item.ImgUrl,
				Item:   item.Name,
				ItemId: item.ItemId,
			})
		}
		if err := tx.Create(&specItem).Error; err != nil {
			tx.Rollback()
			global.SugarLog.Errorf(log+"specItem: %#v, err: %s", spec, err.Error())
			return errors.New("创建商品性规格项失败")
		}

		// 创建规格项明细
		// keys = 1,3
		for keys, value := range from.SpecValue {
			keyArr := strings.Split(keys, ",")
			if len(keyArr) == 0 {
				tx.Rollback()
				global.SugarLog.Errorf(log+"规格明细参数异常 keys: %s, err: %s", keys, err.Error())
				return errors.New("规格明细参数异常")
			}
			// 拼接 Key
			itemIdKey := "" // itemIdKey
			keyName := ""   // key的中文名 颜色:黑色,大小:S
			for _, key := range keyArr {
				itemId, itemName := findItemIdByVirtualItemId(&specItem, key)
				if itemId == 0 {
					global.SugarLog.Errorf(log+"处理规格明细数据失败, 查找 findItemIdByVirtualItemId 异常, err: %s", err.Error())
					return errors.New("处理规格明细数据失败")
				}
				specByItem := findSpecByDbItemId(&spec, &specItem, itemId)
				if specByItem == nil {
					global.SugarLog.Errorf(log+"处理规格明细数据失败, 查找 findSpecByDbItemId 异常, err: %s", err.Error())
					return errors.New("处理规格明细数据失败")
				}
				itemIdKey += strconv.Itoa(int(itemId)) + "_"
				keyName += fmt.Sprintf("%s:%s,", specByItem.Title, itemName)
			}
			// 删除末尾字符
			itemIdKey = strings.TrimRight(itemIdKey, "_")
			keyName = strings.TrimRight(keyName, ",")
			specValue = append(specValue, shop.GoodsSpecValue{
				GoodsId: goods.ID,
				ItemIds: itemIdKey,
				KeyName: keyName,
				Price:   value.Price,
				Store:   value.Store,
				Sort:    value.Sort,
			})
		}
		err = tx.Create(&specValue).Error
		if err != nil {
			tx.Rollback()
			global.SugarLog.Errorf(log+"specValue: %#v, err: %s", specValue, err.Error())
			return errors.New("创建商品性规格明细失败")
		}
	}

	// 提交事务
	tx.Commit()
	return nil
}

// 传入数据库中真实 itemId 获取到 spec 对象
func findSpecByDbItemId(spec *[]shop.GoodsSpec, specItem *[]shop.GoodsSpecItem, itemId uint) *shop.GoodsSpec {
	for _, i := range *specItem {
		// 如果从前端传入的 specId 相等则 返回数据库中的 ID
		if i.ID == itemId {
			for _, s := range *spec {
				// 如果 id 相等
				// 此时的 item.SpecId 的地址为数据库真实的 specId
				if *i.SpecId == int(s.ID) {
					return &s
				}
			}
		}
	}
	return nil
}

// 如果从前端传入的 specId 相等则 返回数据库中的 ID
func findSpecIdByVirtualSpecId(spec *[]shop.GoodsSpec, specId string) int {
	for _, s := range *spec {
		// 如果从前端传入的 specId 相等则 返回数据库中的 ID
		if s.SpecId == specId {
			return int(s.ID)
		}
	}
	return 0
}

// 如果从前端传入的虚拟 itemId 相等则 返回数据库中的 ID和 规格项名称
func findItemIdByVirtualItemId(item *[]shop.GoodsSpecItem, itemId string) (uint, string) {
	for _, i := range *item {
		// 如果从前端传入的 specId 相等则 返回数据库中的 ID
		if i.ItemId == itemId {
			return i.ID, i.Item
		}
	}
	return 0, ""
}

// DeleteGoods 删除Goods记录
// Author [piexlmax](https://github.com/likfees)
func (goodsService *GoodsService) DeleteGoods(goods shop.Goods) (err error) {
	err = global.DB.Delete(&goods).Error
	return err
}

// DeleteGoodsByIds 批量删除Goods记录
// Author [piexlmax](https://github.com/likfees)
func (goodsService *GoodsService) DeleteGoodsByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]shop.Goods{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateGoods 更新Goods记录
// Author [piexlmax](https://github.com/likfees)
func (goodsService *GoodsService) UpdateGoods(goods shop.Goods) (err error) {
	err = global.DB.Save(&goods).Error
	return err
}

// GetGoods 根据id获取Goods记录
// Author [piexlmax](https://github.com/likfees)
func (goodsService *GoodsService) GetGoods(id uint) (goods shop.Goods, err error) {
	err = global.DB.Where("id = ?", id).
		Preload("Desc").
		Preload("Images").
		Preload("Spec").
		First(&goods).Error
	if err != nil {
		return goods, errors.New("获取商品详情失败")
	}
	if *goods.SpecType == 1 {
		for k, s := range goods.Spec {
			var specItem []shop.GoodsSpecItem
			err := global.DB.Where("spec_id = ?", s.ID).Find(&specItem).Error
			if err != nil {
				return shop.Goods{}, errors.New("获取商品规格失败")
			}
			goods.Spec[k].SpecItem = specItem
		}
		var specValue []shop.GoodsSpecValue
		err = global.DB.Where("goods_id = ?", goods.ID).Find(&specValue).Error
		if err != nil {
			return shop.Goods{}, errors.New("获取商品规格明细失败")
		}
		goods.SpecValue = specValue
	}
	return
}

// GetGoodsInfoList 分页获取Goods记录
// Author [piexlmax](https://github.com/likfees)
func (goodsService *GoodsService) GetGoodsInfoList(info shopReq.GoodsSearch) (list []shop.Goods, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	var goodss []shop.Goods
	db := global.DB.Model(&shop.Goods{}).Preload("Desc").Preload("Images").Preload("Category").Preload("Brand")
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.GoodsArea != nil {
		db = db.Where("goods_area = ?", info.GoodsArea)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	db = db.Order("sort asc, created_at desc")
	err = db.Limit(limit).Offset(offset).Find(&goodss).Error
	return goodss, total, err
}
