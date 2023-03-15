package account

type RouterGroup struct {
	AccountGroupRouter
	AccountRouter
	SysRechargeRouter
	UserFinanceTypeRouter
	UserFinanceCashRouter
}
