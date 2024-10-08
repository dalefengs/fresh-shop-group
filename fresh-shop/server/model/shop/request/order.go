package request

import (
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/shop"
	"time"
)

type OrderSearch struct {
	shop.Order
	StartCreatedAt    *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt      *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	StartShipmentTime *time.Time `json:"startShipmentTime" form:"startShipmentTime"`
	EndShipmentTime   *time.Time `json:"endShipmentTime" form:"endShipmentTime"`
	StartReceiveTime  *time.Time `json:"startReceiveTime" form:"startReceiveTime"`
	EndReceiveTime    *time.Time `json:"endReceiveTime" form:"endReceiveTime"`
	StartCancelTime   *time.Time `json:"startCancelTime" form:"startCancelTime"`
	EndCancelTime     *time.Time `json:"endCancelTime" form:"endCancelTime"`
	SettlementMonth   *time.Time `json:"settlementMonth" form:"settlementMonth"`
	UserPhone         string     `json:"userPhone" form:"userPhone"`
	request.PageInfo
}

type BatchSettlement struct {
	UserId          uint       `json:"userId" form:"userId"`
	Payment         *int       `json:"payment" form:"payment"`
	SettlementMonth *time.Time `json:"settlementMonth" form:"settlementMonth"`
}
