package shop

import (
	"errors"
	"fmt"
	"fresh-shop/server/global"
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/shop"
	shopReq "fresh-shop/server/model/shop/request"
	"fresh-shop/server/utils"
	"fresh-shop/server/utils/upload"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	_ "image/png"
	"mime/multipart"
	"strconv"
	"strings"
)

type GoodsService struct {
}

// 字段与单元格位置映射
var excelGoods = map[string]string{
	"name":         "",  // 商品名称
	"categoryName": "B", // 分类名称
	"brandName":    "C", // 品牌名称
	"costPrice":    "D", // 原价
	"price":        "E", // 优惠价
	"minCount":     "F", // 最小购买数量
	"origin":       "G", // 产地
	"unit":         "H", // 单位
	"weight":       "I", // 重量
	"store":        "J", // 库存
	"isHot":        "K", // 是否热销
	"isNew":        "L", // 是否上心
	"details":      "M", // 详情
	"image1":       "N", // 图片
	"image2":       "O", //
	"image3":       "P", //
	"image4":       "Q", //
	"image5":       "R", //
	"image6":       "S", //
}

// 字段与单元格位置映射
var excelGoodsIndex = map[string]int{
	"name":         0,  // 商品名称
	"categoryName": 1,  // 分类名称
	"brandName":    2,  // 品牌名称
	"costPrice":    3,  // 原价
	"price":        4,  // 优惠价
	"minCount":     5,  // 最小购买数量
	"origin":       6,  // 产地
	"unit":         7,  // 单位
	"weight":       8,  // 重量
	"store":        9,  // 库存
	"isHot":        10, // 是否热销
	"isNew":        10, // 是否上心
	"details":      12, // 详情
}

// BatchCreateGoodsByExcel 批量导入商品信息
// Author [likfees](https://github.com/likfees)
func (goodsService *GoodsService) BatchCreateGoodsByExcel(header *multipart.FileHeader) (err error) {
	oss := upload.NewOss()
	filePath, _, err := oss.UploadFile(header)
	if err != nil {
		global.SugarLog.Errorf("上传文件失败 %v", err)
		return errors.New("上传文件失败")
	}
	f, err := excelize.OpenFile(filePath)
	//f, err := excelize.OpenFile("D:\\Code\\fresh-shop-group\\fresh-shop\\server\\uploads\\file\\excel\\goodsImportTemplate.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(rows) <= 1 {
		global.SugarLog.Errorf("未在 Excel 表中查询到记录, len(rows) = %d", len(rows))
		return errors.New("未在 Excel 表中查询到记录")
	}
	txDB := global.DB.Begin()
	for key, row := range rows {
		rowIndex := key + 1
		if rowIndex <= 2 { // 前面两行跳过
			continue
		}
		fmt.Printf("正在遍历 Excel 第 %d 行: \n", rowIndex)
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		log := fmt.Sprintf("第 %d 行, ", rowIndex)
		name := row[excelGoodsIndex["name"]]
		categoryName := row[excelGoodsIndex["categoryName"]]
		brandName := row[excelGoodsIndex["brandName"]]
		costPrice := row[excelGoodsIndex["costPrice"]]
		price := row[excelGoodsIndex["price"]]
		minCount := row[excelGoodsIndex["minCount"]]
		origin := row[excelGoodsIndex["origin"]]
		unit := row[excelGoodsIndex["unit"]]
		weight := row[excelGoodsIndex["weight"]]
		store := row[excelGoodsIndex["store"]]
		isHot := row[excelGoodsIndex["isHot"]]
		isNew := row[excelGoodsIndex["isNew"]]
		details := row[excelGoodsIndex["details"]]

		err := paramCheckEmpty(rowIndex, name, categoryName, costPrice, unit, weight, store)
		if err != nil {
			txDB.Callback()
			global.SugarLog.Errorw(err.Error())
			return err
		}
		// 分类
		var category shop.Category
		categoryErr := txDB.Where("title = ?", categoryName).First(&category).Error
		if errors.Is(categoryErr, gorm.ErrRecordNotFound) {
			category = shop.Category{
				Pid:     utils.Pointer(0),
				Title:   categoryName,
				IsFirst: utils.Pointer(0),
				Sort:    utils.Pointer(50),
			}
			if categoryErr := txDB.Create(&category).Error; categoryErr != nil {
				txDB.Callback()
				global.SugarLog.Errorf(log+"分类创建失败 categoryName: %s, err:%v", categoryName, categoryErr)
				return errors.New(log + "分类创建失败")
			}
		} else if categoryErr != nil {
			txDB.Callback()
			global.SugarLog.Errorf(log+"分类查询失败 categoryName: %s, err:%v", categoryName, categoryErr)
			return errors.New(log + "分类查询失败")
		}

		// 品牌
		brand := shop.Brand{}
		if brandName != "" {
			brandErr := txDB.Where("name = ?", brandName).First(&brand).Error
			if errors.Is(brandErr, gorm.ErrRecordNotFound) {
				brand = shop.Brand{
					Name: brandName,
					Sort: utils.Pointer(50),
				}
				if brandErr = txDB.Create(&brand).Error; brandErr != nil {
					txDB.Callback()
					global.SugarLog.Errorf(log+"品牌创建失败 brandName: %s, err:%v", brandName, brandErr)
					return errors.New(log + "品牌创建失败")
				}
			} else if brandErr != nil {
				txDB.Callback()
				global.SugarLog.Errorf(log+"品牌查询失败 brandName: %s, err:%v", brandName, brandErr)
				return errors.New(log + "品牌查询失败")
			}
		}

		costPriceFloat, err := strconv.ParseFloat(costPrice, 64)
		if err != nil {
			txDB.Callback()
			global.SugarLog.Errorf(log+"转换原价价失败 costPrice: %s, err:%v", costPrice, err)
			return errors.New(log + "原价必须为数字")
		}
		priceFloat := 0.0
		if price != "" {
			priceFloat, err = strconv.ParseFloat(price, 64)
			if err != nil {
				txDB.Callback()
				global.SugarLog.Errorf(log+"转换优惠价失败 price: %s, err:%v", price, err)
				return errors.New(log + "优惠价格必须为数字")
			}
		}
		minCountInt := 0
		if minCount != "" {
			minCountInt, err = strconv.Atoi(minCount)
			if err != nil {
				txDB.Callback()
				global.SugarLog.Errorf(log+"转换最小购买数量失败 minCount: %s, err:%v", minCount, err)
				return errors.New(log + "最小购买数量必须为数字")
			}
		}
		weightInt, err := strconv.Atoi(weight)
		if err != nil {
			txDB.Callback()
			global.SugarLog.Errorf(log+"转换重量失败 weight: %s, err:%v", weight, err)
			return errors.New(log + "商品重量必须为数字")
		}

		isHotInt := 0
		if isHot != "" {
			isHotInt, err = strconv.Atoi(isHot)
			if err != nil {
				txDB.Callback()
				global.SugarLog.Errorf(log+"转换是否热销失败 isHot: %s, err:%v", isHot, err)
				return errors.New(log + "是否热销必须为数字")
			}
		}

		isNewInt := 0
		if isNew != "" {
			isNewInt, err = strconv.Atoi(isNew)
			if err != nil {
				txDB.Callback()
				global.SugarLog.Errorf(log+"转换是否上新失败 isNewInt: %s, err:%v", isNewInt, err)
				return errors.New(log + "是否上新必须为数字")
			}
		}

		goods := shop.Goods{
			GoodsArea:  utils.Pointer(0),
			SpecType:   utils.Pointer(0),
			Sort:       utils.Pointer(50),
			Name:       name,
			CategoryId: utils.Pointer(int(category.ID)),
			BrandId:    utils.Pointer(int(brand.ID)),
			CostPrice:  utils.Pointer(costPriceFloat),
			Price:      utils.Pointer(priceFloat),
			MinCount:   utils.Pointer(minCountInt),
			Weight:     utils.Pointer(weightInt),
			Unit:       unit,
			Origin:     origin,
			IsHot:      utils.Pointer(isHotInt),
			IsNew:      utils.Pointer(isNewInt),
		}

		if err := txDB.Create(&goods).Error; err != nil {
			txDB.Callback()
			global.SugarLog.Errorf(log+"创建商品信息失败 goods: %v, err:%v", goods, err)
			return errors.New(log + "创建商品信息失败")
		}

		// 商品详情
		goodsDetails := shop.GoodsDescription{
			GoodsId: utils.Pointer(int(goods.ID)),
			Details: details,
		}
		if err := txDB.Create(&goodsDetails).Error; err != nil {
			txDB.Callback()
			global.SugarLog.Errorf(log+"创建商品详情失败 goodsDetails: %v, err:%v", goodsDetails, err)
			return errors.New(log + "创建商品详情失败")
		}
		var images []shop.GoodsImage
		// 获取图片信息
		imgCell := []string{"N", "O", "P", "Q", "R", "S"} // 图片2 - 5单元格
		for _, c := range imgCell {
			getExcelGoodsImages(f, &images, int(goods.ID), c, rowIndex)
		}
		if len(images) > 0 {
			if err := txDB.Create(&images).Error; err != nil {
				txDB.Callback()
				global.SugarLog.Errorf(log+"创建商品图片信息失败 images: %v, err:%v", images, err)
				return errors.New(log + "创建商品图片信息失败")
			}
		}

		fmt.Println()
	}
	txDB.Commit()
	return
}

func getExcelGoodsImages(f *excelize.File, list *[]shop.GoodsImage, goodsId int, cell string, rowIndex int) {
	log := fmt.Sprintf("正在获取%s%d图片", cell, rowIndex)
	// 如果做线上图片服务器需要配置接口使用 oss 包
	localOss := upload.Local{}
	img1, err := f.GetPictures("Sheet1", fmt.Sprintf("%s%d", cell, rowIndex))
	if err != nil {
		global.SugarLog.Errorf(log+"获取图片信息失败 N%d, err:%v", rowIndex, err)
		return
	}
	if len(img1) == 0 {
		return
	}
	filePath, _, uploadErr := localOss.UploadFileByBytes(&img1[0].File, img1[0].Extension)
	if uploadErr != nil {
		global.SugarLog.Errorf(log+"上传图片信息失败 %s%d, err:%v", cell, rowIndex, uploadErr)
		return
	}
	img := shop.GoodsImage{
		GoodsId: utils.Pointer(goodsId),
		Url:     filePath,
		Sort:    utils.Pointer(50),
	}
	*list = append(*list, img)
}

// 验证是否为空
func paramCheckEmpty(rowIndex int, name, categoryName, constPrice, uint, weight, store string) error {
	msg := ""
	if name == "" {
		msg = msg + "商品名称、"
	}
	if categoryName == "" {
		msg = msg + "商品分类、"
	}
	if constPrice == "" {
		msg = msg + "商品原价、"
	}
	if uint == "" {
		msg = msg + "商品单位、"
	}
	if weight == "" {
		msg = msg + "商品重量、"
	}
	if store == "" {
		msg = msg + "商品库存、"
	}
	if msg != "" {
		msg = strings.TrimRight(msg, "、")
		return fmt.Errorf("第 %d 行, %s 不能为空", rowIndex, msg)
	}
	return nil
}

// CreateGoods 创建Goods记录
// Author [likfees](https://github.com/likfees)
func (goodsService *GoodsService) CreateGoods(form shopReq.GoodsSubmitFrom) (err error) {
	log := "创建商品 --- "

	goods := form.GoodsInfo
	var spec []shop.GoodsSpec
	var specItem []shop.GoodsSpecItem
	var specValue []shop.GoodsSpecValue
	goodsDesc := shop.GoodsDescription{ // 商品详情
		Details: form.Desc.Details,
		Notice:  form.Desc.Notice,
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
	if len(form.Images) > 0 {
		imgs := form.Images
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
	if *form.GoodsInfo.SpecType == 1 {
		// 创建规格
		// 处理规格项数据
		for _, s := range form.Spec {
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
		for _, item := range form.SpecItem {
			// 查找数据中的规格Id
			specId := findSpecIdByVirtualSpecId(&spec, item.SpecId)
			if specId == 0 {
				tx.Rollback()
				global.SugarLog.Errorf(log+"处理规格项数据失败, 查找 SpecId 异常, err: %s", err.Error())
				return errors.New("处理规格项数据失败")
			}
			specItem = append(specItem, shop.GoodsSpecItem{
				SpecId:  &specId, // 此时记录的是数据库中真实的SpecId
				GoodsId: goods.ID,
				ImgUrl:  item.ImgUrl,
				Item:    item.Name,
				ItemId:  item.ItemId,
			})
		}
		if err := tx.Create(&specItem).Error; err != nil {
			tx.Rollback()
			global.SugarLog.Errorf(log+"specItem: %#v, err: %s", spec, err.Error())
			return errors.New("创建商品性规格项失败")
		}

		// 创建规格项明细
		// keys = 1,3
		for keys, value := range form.SpecValue {
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
					tx.Rollback()
					global.SugarLog.Errorf(log+"处理规格明细数据失败, 查找 findItemIdByVirtualItemId 异常, err: %s", err.Error())
					return errors.New("处理规格明细数据失败")
				}
				specByItem := findSpecByDbItemId(&spec, &specItem, itemId)
				if specByItem == nil {
					tx.Rollback()
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
				GoodsId:   goods.ID,
				ItemIds:   itemIdKey,
				KeyName:   keyName,
				Price:     value.Price,
				CostPrice: value.CostPrice,
				Store:     value.Store,
				Sort:      value.Sort,
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

// UpdateGoods 更新Goods记录
// Author [likfees](https://github.com/likfees)
func (goodsService *GoodsService) UpdateGoods(form shopReq.GoodsSubmitFrom) (err error) {
	log := "更新商品 --- "

	goods := form.GoodsInfo
	var dbGoods shop.Goods
	var spec []shop.GoodsSpec
	var specItem []shop.GoodsSpecItem
	// 查看商品数据
	err = global.DB.Where("id = ?", goods.ID).Preload("Spec").Preload("Desc").First(&dbGoods).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		global.SugarLog.Errorf(log+"商品不存在: id: %d, err: %s", goods.ID, err.Error())
		return errors.New("商品不存在")
	}

	// 处理商品详情编辑数据
	goodsDesc := dbGoods.Desc
	goodsDesc.Details = form.Desc.Details
	goodsDesc.Notice = form.Desc.Notice

	// region 处理图片 ID 数据
	var dbUrls []string    // DB中商品图片列表
	var formUrls []string  // 表单中商品图片列表
	var deleteUrl []string // 需要删除的图片列表
	var createUrl []string // 需要添加的图片列表
	for _, i := range form.Images {
		formUrls = append(formUrls, i.Url)
	}
	if err := global.DB.Model(&shop.GoodsImage{}).Where("goods_id = ?", goods.ID).Pluck("url", &dbUrls).Error; err != nil {
		global.SugarLog.Errorf(log+"获取商品图片数据失败: id: %d, err: %s", goods.ID, err.Error())
		return errors.New("获取商品图片数据失败")
	}
	if len(dbUrls) == 0 { // 如果数据库中没有数据，则全部插入
		createUrl = formUrls
	} else {
		// 求差集 数据库中有的，表单中没有的， 则为需要删除的记录
		deleteUrl = utils.SliceDifference(dbUrls, formUrls)
		// 求差集 表单中有的，数据库中没有的， 则为添加的记录
		createUrl = utils.SliceDifference(formUrls, dbUrls)
	}
	// endregion

	// region 处理规格 ID 数据
	var dbSpecId []int               // DB中规格ID列表
	var formSpecId []int             // 表单中规格ID列表
	var deleteSpecId []int           // 需要删除的规格ID列表
	var createSpecId []int           // 需要添加的规格ID列表
	var unionSpecId []int            // 需要编辑的规格Id列表
	var createSpec []shop.GoodsSpec  // 需要添加的规格列表
	var unionSpec []shop.GoodsSpec   // 需要编辑的规格列表
	for _, s := range dbGoods.Spec { // 得到数据库中现有的 ID 列表
		dbSpecId = append(dbSpecId, int(s.ID))
	}
	for _, s := range form.Spec { // 得到表单传递数据
		id, err := strconv.Atoi(s.SpecId)
		if err != nil {
			global.SugarLog.Errorf(log+"处理规格表单数据失败, 类型转换异常: SpecId: %s, err: %s", s.SpecId, err.Error())
			return errors.New("处理规格表单数据失败")
		}
		formSpecId = append(formSpecId, id)
	}
	// 求差集
	deleteSpecId = utils.SliceDifference(dbSpecId, formSpecId)
	createSpecId = utils.SliceDifference(formSpecId, dbSpecId)
	// 求并集 修改
	unionSpecId = utils.SliceUnionSet(dbSpecId, formSpecId)
	// endregion

	// region 处理规格项 ID 数据
	var dbItemId []int                  // DB中规格项ID列表
	var formItemId []int                // 表单中规格项ID列表
	var deleteItemId []int              // 需要删除的规格项ID列表
	var createItemId []int              // 需要添加的规格项ID列表
	var unionItemId []int               // 需要编辑的规格项Id列表
	var createItem []shop.GoodsSpecItem // 需要添加的规格项列表
	var unionItem []shop.GoodsSpecItem  // 需要编辑的规格项列表

	err = global.DB.Model(&shop.GoodsSpecItem{}).Where("goods_id = ?", goods.ID).Pluck("id", &dbItemId).Error
	if err != nil {
		global.SugarLog.Errorf(log+"获取商品规格ID列表失败 err: %s", err.Error())
		return errors.New("处理商品规格失败")
	}
	for _, i := range form.SpecItem {
		id, err := strconv.Atoi(i.ItemId)
		if err != nil {
			global.SugarLog.Errorf(log+"处理规格项表单数据失败, 类型转换异常: ItemId: %s, err: %s", i.ItemId, err.Error())
			return errors.New("处理规格表单数据失败")
		}
		formItemId = append(formItemId, id)
	}
	// 求差集
	deleteItemId = utils.SliceDifference(dbItemId, formItemId)
	createItemId = utils.SliceDifference(formItemId, dbItemId)
	// 求并集
	unionItemId = utils.SliceUnionSet(dbItemId, formItemId)
	// endregion

	// region 处理规格明细 ID 数据
	var dbValueId []string                // DB中规格明细ID列表
	var formValueId []string              // 表单中规格明细ID列表
	var deleteValueId []string            // 需要删除的规格明细ID列表
	var createValueId []string            // 需要添加的规格明细ID列表
	var unionValueId []string             // 需要编辑的规格明细Id列表
	var createValue []shop.GoodsSpecValue // 需要添加的规格明细列表
	var unionValue []shop.GoodsSpecValue  // 需要编辑的规格明细列表

	err = global.DB.Model(&shop.GoodsSpecValue{}).Where("goods_id = ?", goods.ID).Pluck("item_ids", &dbValueId).Error
	if err != nil {
		global.SugarLog.Errorf(log+"获取商品规明细 item_ids 列表失败 err: %s", err.Error())
		return errors.New("处理商品规格明细失败")
	}
	for itemIds, _ := range form.SpecValue {
		itemIds = strings.ReplaceAll(itemIds, ",", "_")
		formValueId = append(formValueId, itemIds)
	}
	// 求差集
	deleteValueId = utils.SliceDifference(dbValueId, formValueId)
	createValueId = utils.SliceDifference(formValueId, dbValueId)
	// 求并集
	unionValueId = utils.SliceUnionSet(dbValueId, formValueId)
	// endregion

	// 开始事务
	tx := global.DB.Begin()

	// 更新商品基本信息
	if err := tx.Save(&goods).Error; err != nil {
		tx.Rollback()
		global.SugarLog.Errorf(log+" 更新商品信息失败 goodsInfo: %#v, err: %s", goods, err.Error())
		return errors.New("更新商品信息失败")
	}
	goodsIdPointr := utils.Pointer(int(goods.ID))

	// 更新商品详情信息
	goodsDesc.GoodsId = goodsIdPointr
	if err := tx.Save(&goodsDesc).Error; err != nil {
		tx.Rollback()
		global.SugarLog.Errorf(log+" 更新商品详情信息失败 goodsDesc: %#v err: %s", goodsDesc, err.Error())
		return errors.New("更新商品详情信息失败")
	}

	// 更新商品图片信息
	if len(deleteUrl) > 0 { // 删除图片信息
		err := tx.Where("url in ?", deleteUrl).Delete(&shop.GoodsImage{}).Error
		if err != nil {
			tx.Rollback()
			global.SugarLog.Errorf(log+"删除图片信息失败  deleteUrl: %#v, err: %s", deleteUrl, err.Error())
			return errors.New("处理商品图片数据失败")
		}
	}
	if len(createUrl) > 0 { // 添加图片信息
		var imgs []shop.GoodsImage
		// 通过 url 查找到对应的对象
		for _, f := range form.Images {
			for _, url := range createUrl {
				if f.Url == url {
					imgs = append(imgs, shop.GoodsImage{
						GoodsId: goodsIdPointr,
						Url:     f.Url,
						Name:    f.Name,
						Type:    f.Type,
					})
				}
			}
		}
		if err := tx.Create(&imgs).Error; err != nil {
			tx.Rollback()
			global.SugarLog.Errorf(log+" 更新商品图片列表失败 imgs: %#v, err: %s", imgs, err.Error())
			return errors.New("更新商品图片列表失败")
		}
	}

	// region update 多规格
	if *form.GoodsInfo.SpecType == 1 {

		// region 处理规格数据
		if len(deleteSpecId) > 0 {
			err := tx.Where("id in ?", deleteSpecId).Delete(&shop.GoodsSpec{}).Error
			if err != nil {
				tx.Rollback()
				global.SugarLog.Errorf(log+"删除规格失败  deleteSpecId: %#v, err: %s", deleteSpecId, err.Error())
				return errors.New("处理规格数据失败")
			}
		}
		if len(createSpecId) > 0 {
			// 处理规格项数据
			for _, id := range createSpecId {
				for _, f := range form.Spec {
					if f.SpecId == strconv.Itoa(id) {
						createSpec = append(createSpec, shop.GoodsSpec{
							GoodsId:       int(goods.ID),
							Title:         f.Name,
							IsUploadImage: f.IsUploadImage,
							Sort:          50,
							SpecId:        f.SpecId,
						})
					}
				}
			}
			if err := tx.Create(&createSpec).Error; err != nil {
				tx.Rollback()
				global.SugarLog.Errorf(log+"创建商品规格失败: createSpec %#v, err: %s", createSpec, err.Error())
				return errors.New("创建商品规格失败")
			}
			// 将 创建和编辑数组合并到 Spec
			spec = append(spec, createSpec...)
		}

		// 编辑更新规格
		if len(unionSpecId) > 0 {
			for _, id := range unionSpecId {
				for _, f := range form.Spec {
					if f.SpecId == strconv.Itoa(id) {
						s := shop.GoodsSpec{
							GoodsId:       int(goods.ID),
							Title:         f.Name,
							IsUploadImage: f.IsUploadImage,
							Sort:          50,
							SpecId:        f.SpecId,
						}
						s.ID = uint(id)
						unionSpec = append(unionSpec, s)
					}
				}
			}
			if err := tx.Save(&unionSpec).Error; err != nil {
				tx.Rollback()
				global.SugarLog.Errorf(log+"更新商品规格失败 unionSpec: %#v, err: %s", unionSpec, err.Error())
				return errors.New("更新商品规格失败")
			}
			spec = append(spec, unionSpec...)
		}

		// endregion

		// region 处理规格项数据
		if len(deleteItemId) > 0 {
			err := tx.Where("id in ?", deleteItemId).Delete(&shop.GoodsSpecItem{}).Error
			if err != nil {
				tx.Rollback()
				global.SugarLog.Errorf(log+"删除规格项失败  deleteItemId: %#v, err: %s", deleteItemId, err.Error())
				return errors.New("处理规格项数据失败")
			}
		}
		// 创建规格项
		if len(createItemId) > 0 {
			for _, id := range createItemId {
				for _, item := range form.SpecItem {
					if item.ItemId == strconv.Itoa(id) {
						// 查找数据中的规格Id
						specId := findSpecIdByVirtualSpecId(&spec, item.SpecId)
						if specId == 0 {
							tx.Rollback()
							global.SugarLog.Errorf(log+"创建规格项 处理规格项数据失败, 查找 SpecId 异常, err: %s", err.Error())
							return errors.New("处理规格项数据失败")
						}
						createItem = append(createItem, shop.GoodsSpecItem{
							SpecId:  &specId, // 此时记录的是数据库中真实的SpecId
							GoodsId: goods.ID,
							ImgUrl:  item.ImgUrl,
							Item:    item.Name,
							ItemId:  item.ItemId,
						})
					}
				}
			}
			if err := tx.Create(&createItem).Error; err != nil {
				tx.Rollback()
				global.SugarLog.Errorf(log+"创建商品性规格项失败 createItem: %#v, err: %s", createItem, err.Error())
				return errors.New("创建商品性规格项失败")
			}
			specItem = append(specItem, createItem...)
		}
		// 编辑规格项
		if len(unionItemId) > 0 {
			for _, id := range unionItemId {
				for _, item := range form.SpecItem {
					if item.ItemId == strconv.Itoa(id) {
						// 查找数据中的规格Id
						specId := findSpecIdByVirtualSpecId(&spec, item.SpecId)
						if specId == 0 {
							tx.Rollback()
							global.SugarLog.Errorf(log+"编辑规格项 处理规格项数据失败, 查找 SpecId 异常, err: %s", err.Error())
							return errors.New("处理规格项数据失败")
						}
						i := shop.GoodsSpecItem{
							SpecId:  &specId, // 此时记录的是数据库中真实的SpecId
							ImgUrl:  item.ImgUrl,
							Item:    item.Name,
							ItemId:  item.ItemId,
							GoodsId: goods.ID,
						}
						i.ID = uint(id)
						unionItem = append(unionItem, i)
					}
				}
			}
			if err := tx.Save(&unionItem).Error; err != nil {
				tx.Rollback()
				global.SugarLog.Errorf(log+"更新商品性规格项失败 unionItem: %#v, err: %s", unionItem, err.Error())
				return errors.New("更新商品性规格项失败")
			}
			specItem = append(specItem, unionItem...)
		}
		// endregion

		// region 处理规格项明细
		if len(deleteValueId) > 0 { // 删除
			err := tx.Where("item_ids in ?", deleteValueId).Delete(&shop.GoodsSpecValue{}).Error
			if err != nil {
				tx.Rollback()
				global.SugarLog.Errorf(log+"删除规格明细失败  deleteValueId: %#v, err: %s", deleteValueId, err.Error())
				return errors.New("处理规格明细失败")
			}
		}
		if len(createValueId) > 0 {
			for _, id := range createValueId {
				ids := strings.ReplaceAll(id, "_", ",")
				for keys, value := range form.SpecValue {
					if ids != keys {
						continue
					}
					keyArr := strings.Split(keys, ",")
					if len(keyArr) == 0 {
						tx.Rollback()
						global.SugarLog.Errorf(log+"创建 -- 规格明细参数异常 keys: %s, err: %s", keys, err.Error())
						return errors.New("规格明细参数异常")
					}
					// 拼接 Key
					itemIdKey := "" // itemIdKey
					keyName := ""   // key的中文名 颜色:黑色,大小:S
					for _, key := range keyArr {
						itemId, itemName := findItemIdByVirtualItemId(&specItem, key)
						if itemId == 0 {
							tx.Rollback()
							global.SugarLog.Errorf(log+"创建 -- 处理规格明细数据失败, 查找 findItemIdByVirtualItemId 异常, err: %s", err.Error())
							return errors.New("处理规格明细数据失败")
						}
						specByItem := findSpecByDbItemId(&spec, &specItem, itemId)
						if specByItem == nil {
							tx.Rollback()
							global.SugarLog.Errorf(log+"创建 -- 处理规格明细数据失败, 查找 findSpecByDbItemId 异常, err: %s", err.Error())
							return errors.New("处理规格明细数据失败")
						}
						itemIdKey += strconv.Itoa(int(itemId)) + "_"
						keyName += fmt.Sprintf("%s:%s,", specByItem.Title, itemName)
					}
					// 删除末尾字符
					itemIdKey = strings.TrimRight(itemIdKey, "_")
					keyName = strings.TrimRight(keyName, ",")
					createValue = append(createValue, shop.GoodsSpecValue{
						GoodsId:   goods.ID,
						ItemIds:   itemIdKey,
						KeyName:   keyName,
						CostPrice: value.CostPrice,
						Price:     value.Price,
						Store:     value.Store,
						Sort:      value.Sort,
					})
				}
			}
			err = tx.Create(&createValue).Error
			if err != nil {
				tx.Rollback()
				global.SugarLog.Errorf(log+"创建 -- createValue: %#v, err: %s", createValue, err.Error())
				return errors.New("更新商品性规格明细失败")
			}
		}

		if len(unionValueId) > 0 {
			for _, id := range unionValueId {
				ids := strings.ReplaceAll(id, "_", ",")
				for keys, value := range form.SpecValue {
					if ids != keys {
						continue
					}
					keyArr := strings.Split(keys, ",")
					if len(keyArr) == 0 {
						tx.Rollback()
						global.SugarLog.Errorf(log+"更新 -- 规格明细参数异常 keys: %s, err: %s", keys, err.Error())
						return errors.New("规格明细参数异常")
					}
					// 拼接 Key
					itemIdKey := "" // itemIdKey
					keyName := ""   // key的中文名 颜色:黑色,大小:S
					for _, key := range keyArr {
						itemId, itemName := findItemIdByVirtualItemId(&specItem, key)
						if itemId == 0 {
							tx.Rollback()
							global.SugarLog.Errorf(log + "更新 -- 处理规格明细数据失败, 查找 findItemIdByVirtualItemId 异常")
							return errors.New("处理规格明细数据失败")
						}
						specByItem := findSpecByDbItemId(&spec, &specItem, itemId)
						if specByItem == nil {
							tx.Rollback()
							global.SugarLog.Errorf(log + "更新 -- 处理规格明细数据失败, 查找 findSpecByDbItemId 异常")
							return errors.New("处理规格明细数据失败")
						}
						itemIdKey += strconv.Itoa(int(itemId)) + "_"
						keyName += fmt.Sprintf("%s:%s,", specByItem.Title, itemName)
					}
					// 删除末尾字符
					itemIdKey = strings.TrimRight(itemIdKey, "_")
					keyName = strings.TrimRight(keyName, ",")
					unionValue = append(unionValue, shop.GoodsSpecValue{
						GoodsId:   goods.ID,
						ItemIds:   itemIdKey,
						KeyName:   keyName,
						CostPrice: value.CostPrice,
						Price:     value.Price,
						Store:     value.Store,
						Sort:      value.Sort,
					})
				}
			}

			for _, u := range unionValue {
				err := tx.Model(&shop.GoodsSpecValue{}).Where("item_ids = ?", u.ItemIds).Updates(u).Error
				if err != nil {
					tx.Rollback()
					global.SugarLog.Errorf(log+"更新 -- unionValue: %#v, err: %s", unionValue, err.Error())
					return errors.New("更新商品性规格明细失败")
				}
			}

		}
		// endregion
	}
	// endregion
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
// Author [likfees](https://github.com/likfees)
func (goodsService *GoodsService) DeleteGoods(goods shop.Goods) (err error) {
	err = global.DB.Delete(&goods).Error
	return err
}

// DeleteGoodsByIds 批量删除Goods记录
// Author [likfees](https://github.com/likfees)
func (goodsService *GoodsService) DeleteGoodsByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]shop.Goods{}, "id in ?", ids.Ids).Error
	return err
}

// GetGoods 根据id获取Goods记录
// Author [likfees](https://github.com/likfees)
func (goodsService *GoodsService) GetGoods(id, userId uint) (goods shop.Goods, err error) {
	err = global.DB.Where("id = ?", id).
		Preload("Desc").
		Preload("Images").
		Preload("Spec").
		Preload("Brand").
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
	// 查询商品是否收藏
	var f shop.GoodsFavorites
	if errors.Is(global.DB.Where("user_id = ? and goods_id = ?", userId, id).First(&f).Error, gorm.ErrRecordNotFound) {
		goods.IsFavorite = false
	} else {
		goods.IsFavorite = true
	}
	var cart shop.Cart   // 当前商品购物车中存在的数量
	var cartTotalNum int // 当前用户购物车总存在的数量
	// 查询当前用户和当前商品购物车数量 直接查询 num 字段
	global.DB.Model(shop.Cart{}).Where("user_id = ?", userId).Pluck("SUM(num) as cartNum", &cartTotalNum)
	cartErr := global.DB.Model(shop.Cart{}).Where("user_id = ? and goods_id = ?", userId, id).First(&cart).Error
	if cartErr != nil {
		goods.CartNum = utils.Pointer(0)
	} else {
		goods.CartNum = &cart.Num
	}
	goods.CartTotalNum = &cartTotalNum
	return
}

// GetGoodsInfoList 分页获取Goods记录
// Author [likfees](https://github.com/likfees)
func (goodsService *GoodsService) GetGoodsInfoList(info shopReq.GoodsSearch, userId uint) (list []shop.Goods, total int64, err error) {
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
	if info.BrandId != nil {
		db = db.Where("brand_id = ?", info.BrandId)
	}
	if info.CategoryId != nil {
		db = db.Where("category_id = ?", info.CategoryId)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.IsNew != nil {
		db = db.Where("is_new = ?", info.IsNew)
	}
	if info.IsHot != nil {
		db = db.Where("is_hot = ?", info.IsHot)
	}
	if info.IsFirst != nil {
		db = db.Where("is_first = ?", info.IsFirst)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	db = db.Order("sort asc, created_at desc")
	err = db.Limit(limit).Offset(offset).Order("sort asc").Find(&goodss).Error
	// 用户已经登录
	if err == nil && userId > 0 {
		for key, item := range goodss {
			var cart shop.Cart
			if !errors.Is(global.DB.Where("user_id = ? and goods_id = ?", userId, item.ID).First(&cart).Error, gorm.ErrRecordNotFound) {
				goodss[key].CartNum = &cart.Num
			}
		}
	}
	return goodss, total, err
}
