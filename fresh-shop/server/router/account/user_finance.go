package account

import (
	"fresh-shop/server/api/v1"
	"fresh-shop/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserFinanceCashRouter struct {
}

// InitUserFinanceCashRouter 初始化 UserFinance 路由信息
func (s *UserFinanceCashRouter) InitUserFinanceCashRouter(Router *gin.RouterGroup) {
	userFinanceCashRouter := Router.Group("userFinanceCash").Use(middleware.OperationRecord())
	userFinanceCashRouterWithoutRecord := Router.Group("userFinanceCash")
	var userFinanceCashApi = v1.ApiGroupApp.AccountApiGroup.UserFinanceCashApi
	{
		userFinanceCashRouter.POST("createUserFinanceCash", userFinanceCashApi.CreateUserFinance)             // 新建UserFinanceCash
		userFinanceCashRouter.DELETE("deleteUserFinanceCash", userFinanceCashApi.DeleteUserFinance)           // 删除UserFinanceCash
		userFinanceCashRouter.DELETE("deleteUserFinanceCashByIds", userFinanceCashApi.DeleteUserFinanceByIds) // 批量删除UserFinanceCash
		userFinanceCashRouter.PUT("updateUserFinanceCash", userFinanceCashApi.UpdateUserFinance)              // 更新UserFinanceCash
	}
	{
		userFinanceCashRouterWithoutRecord.GET("findUserFinanceCash", userFinanceCashApi.FindUserFinance)       // 根据ID获取UserFinanceCash
		userFinanceCashRouterWithoutRecord.GET("getUserFinanceCashList", userFinanceCashApi.GetUserFinanceList) // 获取UserFinanceCash列表
	}
}
