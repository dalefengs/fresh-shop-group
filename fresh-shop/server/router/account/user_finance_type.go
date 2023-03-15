package account

import (
	"fresh-shop/server/api/v1"
	"fresh-shop/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserFinanceTypeRouter struct {
}

// InitUserFinanceTypeRouter 初始化 UserFinanceType 路由信息
func (s *UserFinanceTypeRouter) InitUserFinanceTypeRouter(Router *gin.RouterGroup) {
	userFinanceTypeRouter := Router.Group("userFinanceType").Use(middleware.OperationRecord())
	userFinanceTypeRouterWithoutRecord := Router.Group("userFinanceType")
	var userFinanceTypeApi = v1.ApiGroupApp.AccountApiGroup.UserFinanceTypeApi
	{
		userFinanceTypeRouter.POST("createUserFinanceType", userFinanceTypeApi.CreateUserFinanceType)             // 新建UserFinanceType
		userFinanceTypeRouter.DELETE("deleteUserFinanceType", userFinanceTypeApi.DeleteUserFinanceType)           // 删除UserFinanceType
		userFinanceTypeRouter.DELETE("deleteUserFinanceTypeByIds", userFinanceTypeApi.DeleteUserFinanceTypeByIds) // 批量删除UserFinanceType
		userFinanceTypeRouter.PUT("updateUserFinanceType", userFinanceTypeApi.UpdateUserFinanceType)              // 更新UserFinanceType
	}
	{
		userFinanceTypeRouterWithoutRecord.GET("findUserFinanceType", userFinanceTypeApi.FindUserFinanceType)             // 根据ID获取UserFinanceType
		userFinanceTypeRouterWithoutRecord.GET("getUserFinanceTypeList", userFinanceTypeApi.GetUserFinanceTypeList)       // 获取UserFinanceType列表
		userFinanceTypeRouterWithoutRecord.GET("getUserFinanceTypeListAll", userFinanceTypeApi.GetUserFinanceTypeListAll) // 获取所有 UserFinanceType 列表
	}
}
