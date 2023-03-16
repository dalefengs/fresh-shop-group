package shop

// BrandCategory 结构体
type BrandCategory struct {
	BrandId    *int `json:"brandId" form:"brandId" gorm:"column:brand_id;comment:品牌Id;size:20;"`
	CategoryId *int `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:分类Id;size:20;"`
}

// TableName BrandCategory 表名
func (BrandCategory) TableName() string {
	return "shop_brand_category"
}
