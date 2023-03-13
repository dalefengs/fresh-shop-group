package account

import (
	"fresh-shop/server/api/v1"
	"fresh-shop/server/middleware"
	"github.com/gin-gonic/gin"
)

type AccountGroupRouter struct {
}

// InitAccountGroupRouter 初始化 AccountGroup 路由信息
func (s *AccountGroupRouter) InitAccountGroupRouter(Router *gin.RouterGroup) {
	userAccountGroupRouter := Router.Group("userAccountGroup").Use(middleware.OperationRecord())
	userAccountGroupRouterWithoutRecord := Router.Group("userAccountGroup")
	var userAccountGroupApi = v1.ApiGroupApp.AccountApiGroup.AccountGroupApi
	{
		userAccountGroupRouter.POST("createAccountGroup", userAccountGroupApi.CreateAccountGroup)             // 新建AccountGroup
		userAccountGroupRouter.DELETE("deleteAccountGroup", userAccountGroupApi.DeleteAccountGroup)           // 删除AccountGroup
		userAccountGroupRouter.DELETE("deleteAccountGroupByIds", userAccountGroupApi.DeleteAccountGroupByIds) // 批量删除AccountGroup
		userAccountGroupRouter.PUT("updateAccountGroup", userAccountGroupApi.UpdateAccountGroup)              // 更新AccountGroup
	}
	{
		userAccountGroupRouterWithoutRecord.GET("findAccountGroup", userAccountGroupApi.FindAccountGroup)       // 根据ID获取AccountGroup
		userAccountGroupRouterWithoutRecord.GET("getAccountGroupList", userAccountGroupApi.GetAccountGroupList) // 获取AccountGroup列表
	}
}
