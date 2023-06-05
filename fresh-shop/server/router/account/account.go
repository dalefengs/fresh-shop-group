package account

import (
	"fresh-shop/server/api/v1"
	"fresh-shop/server/middleware"
	"github.com/gin-gonic/gin"
)

type AccountRouter struct {
}

// InitAccountRouter 初始化 Account 路由信息
func (s *AccountRouter) InitAccountRouter(Router *gin.RouterGroup) {
	accountRouter := Router.Group("account").Use(middleware.OperationRecord())
	accountRouterWithoutRecord := Router.Group("account")
	var accountApi = v1.ApiGroupApp.AccountApiGroup.AccountApi
	{
		accountRouter.POST("createAccount", accountApi.CreateAccount)             // 新建Account
		accountRouter.DELETE("deleteAccount", accountApi.DeleteAccount)           // 删除Account
		accountRouter.DELETE("deleteAccountByIds", accountApi.DeleteAccountByIds) // 批量删除Account
		accountRouter.PUT("updateAccount", accountApi.UpdateAccount)              // 更新Account
	}
	{
		accountRouterWithoutRecord.GET("findAccount", accountApi.FindAccount)         // 根据ID获取Account
		accountRouterWithoutRecord.GET("findUserAccount", accountApi.FindUserAccount) // 获取当前用户指定 Account
		accountRouterWithoutRecord.GET("getAccountList", accountApi.GetAccountList)   // 获取Account列表
	}
}
