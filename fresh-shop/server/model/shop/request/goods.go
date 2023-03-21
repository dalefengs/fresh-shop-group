package request

import (
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/shop"
	"time"
)

type GoodsSearch struct {
	shop.Goods
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}

// GoodsSubmitFrom 提交表单数据
type GoodsSubmitFrom struct {
	GoodsInfo shop.Goods            `json:"goodsInfo" form:"goodsInfo"`
	Desc      shop.GoodsDescription `json:"desc" form:"desc"`
	Images    []shop.GoodsImage     `json:"images" form:"images"`
	Spec      []goodsSepc           `json:"spec" form:"spec"`
	SpecItem  []specItem            `json:"specItem" form:"specItem"`
	SpecValue map[string]specValue  `json:"specValue" form:"specValue"` // value_id => specItem
}

type goodsSepc struct {
	SpecId        string `json:"specId" from:"specId"`
	Name          string `json:"name" from:"name"`
	IsUploadImage int    `json:"isUploadImage" from:"isUploadImage"`
}

type specItem struct {
	ItemId string `json:"itemId" from:"itemId"`
	Name   string `json:"name" from:"name"`
	SpecId string `json:"specId" from:"specId"`
	ImgUrl string `json:"imgUrl" from:"imgUrl"`
}

type specValue struct {
	Price *float64 `json:"price" from:"price"`
	Sort  *int     `json:"sort" from:"sort"`
	Store *int     `json:"store" from:"store"`
}
