package request

import (
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/shop"
	"time"
)

type FavoritesSearch struct {
	shop.Favorites
	Name           string     `json:"name" form:"name"`
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
