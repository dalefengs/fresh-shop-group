package shop

import (
	"fresh-shop/server/global"
)

// GoodsSpecValue 结构体
type GoodsSpecValue struct {
	global.DbModel
	GoodsId   uint     `json:"goodsId" form:"goodsId" gorm:"column:goods_id;comment:商品id 1_2;size:20;"`
	ItemIds   string   `json:"itemIds" form:"itemIds" gorm:"column:item_ids;comment:规格项id 1_2;size:200;"`
	KeyName   string   `json:"keyName" form:"keyName" gorm:"column:key_name;comment:规格中文键名;size:500;"`
	Price     *float64 `json:"price" form:"price" gorm:"column:price;comment:优惠价格;size:10;"`
	CostPrice *float64 `json:"costPrice" form:"costPrice" gorm:"column:cost_price;default:0;comment:原价;size:10;"`
	Store     *int     `json:"store" form:"store" gorm:"column:store;default:50;comment:库存;size:10;"`
	Sale      *int     `json:"sale" form:"sale" gorm:"column:sale;default:50;comment:销量;size:10;"`
	Sort      *int     `json:"sort" form:"sort" gorm:"column:sort;default:50;comment:排序;size:10;"`
}

// TableName GoodsSpecValue 表名
func (GoodsSpecValue) TableName() string {
	return "shop_goods_spec_value"
}
