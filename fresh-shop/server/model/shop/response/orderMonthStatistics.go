package response

// 当月总订单数：5 单当月未结订单数：5 单当月已结订单数：5 单
// 当月总金额：￥ 0.00当月未结总金额：￥ 0.00当月已结总金额：￥ 0.00
type OrderMonthStatistics struct {
	UserId uint `json:"userId"`

	MonthTotal  int `json:"monthTotal"`  // 当月总订单数
	MonthPaid   int `json:"monthPaid"`   // 当月已结订单数
	MonthUnpaid int `json:"monthUnpaid"` // 当月未结订单数

	SettlementTotal  float64 `json:"settlementTotal"`  // 当月总金额
	SettlementUnpaid float64 `json:"settlementUnpaid"` // 当月未结总金额
	SettlementPaid   float64 `json:"settlementPaid"`   // 当月已结总金额
	Month            string  `json:"month"`
}
