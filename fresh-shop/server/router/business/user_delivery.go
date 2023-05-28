package business

import (
	"fresh-shop/server/api/v1"
	"fresh-shop/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserDeliveryRouter struct {
}

// InitUserDeliveryRouter 初始化 UserDelivery 路由信息
func (s *UserDeliveryRouter) InitUserDeliveryRouter(Router *gin.RouterGroup) {
	userDeliveryRouter := Router.Group("userDelivery").Use(middleware.OperationRecord())
	userDeliveryRouterWithoutRecord := Router.Group("userDelivery")
	var userDeliveryApi = v1.ApiGroupApp.BusinessApiGroup.UserDeliveryApi
	{
		userDeliveryRouter.POST("createUserDelivery", userDeliveryApi.CreateUserDelivery)             // 新建UserDelivery
		userDeliveryRouter.DELETE("deleteUserDelivery", userDeliveryApi.DeleteUserDelivery)           // 删除UserDelivery
		userDeliveryRouter.DELETE("deleteUserDeliveryByIds", userDeliveryApi.DeleteUserDeliveryByIds) // 批量删除UserDelivery
		userDeliveryRouter.PUT("updateUserDelivery", userDeliveryApi.UpdateUserDelivery)              // 更新UserDelivery
	}
	{
		userDeliveryRouterWithoutRecord.GET("findUserDelivery", userDeliveryApi.FindUserDelivery)             // 根据ID获取UserDelivery
		userDeliveryRouterWithoutRecord.GET("getUserDeliveryList", userDeliveryApi.GetUserDeliveryList)       // 获取UserDelivery列表
		userDeliveryRouterWithoutRecord.GET("getUserDeliveryAllList", userDeliveryApi.GetUserDeliveryAllList) // 获取所有 UserDelivery列表
	}
}
