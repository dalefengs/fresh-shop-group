package shop

import (
	"fresh-shop/server/global"
)

// Goods 结构体
type Goods struct {
	global.DbModel
	Name       string           `json:"name" form:"name" gorm:"column:name;comment:商品名称;size:255;"`
	CategoryId *int             `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:分类id;size:20;"`
	BrandId    *int             `json:"brandId" form:"brandId" gorm:"column:brand_id;comment:品牌Id;size:20;"`
	GoodsArea  *int             `json:"goodsArea" form:"goodsArea" gorm:"column:goods_area;default:0;comment:所属区域(0普通商品 1积分商城 );"`
	SpecType   *int             `json:"specType" form:"specType" gorm:"column:spec_type;default:0;comment:规格类型(0单规格 1多规格);"`
	Unit       string           `json:"unit" form:"unit" gorm:"column:unit;comment:商品单位(盒、件、瓶、克等);size:20;"`
	Price      *float64         `json:"price" form:"price" gorm:"column:price;comment:商品价格;size:10;"`
	MinCount   *int             `json:"minCount" form:"minCount" gorm:"column:min_count;default:1;comment:最低购买数量;size:10;"`
	Weight     *int             `json:"weight" form:"weight" gorm:"column:weight;default:0;comment:商品重量（g）;size:10;"`
	Store      *int             `json:"store" form:"store" gorm:"column:store;default:0;comment:库存;size:10;"`
	Sale       *int             `json:"sale" form:"sale" gorm:"column:sale;default:0;comment:所有规格的总销量;size:10;"`
	Sort       *int             `json:"sort" form:"sort" gorm:"column:sort;default:50;comment:排序;size:10;"`
	Status     *int             `json:"status" form:"status" gorm:"column:status;default:1;comment:状态(0 下架 1上架 );"`
	IsFirst    *int             `json:"isFirst" form:"isFirst" gorm:"column:is_first;default:0;comment:是否首页(0否 1是);"`
	IsHot      *int             `json:"isHot" form:"isHot" gorm:"column:is_hot;default:0;comment:是否热销(0否 1是);"`
	IsNew      *int             `json:"isNew" form:"isNew" gorm:"column:is_new;default:0;comment:是否上新(0否 1是);"`
	Desc       GoodsDescription `json:"desc"`
	Images     []GoodsImage     `json:"images"`
	Spec       []GoodsSpec      `json:"spec"`
	SpecValue  []GoodsSpecValue `json:"specValue"`
	Category   Category         `json:"category"`
	Brand      Brand            `json:"brand"`
}

// TableName Goods 表名
func (Goods) TableName() string {
	return "shop_goods"
}
