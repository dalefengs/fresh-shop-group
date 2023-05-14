package shop

import (
	"fresh-shop/server/api/v1"
	"fresh-shop/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserAddressRouter struct {
}

// InitUserAddressRouter 初始化 UserAddress 路由信息
func (s *UserAddressRouter) InitUserAddressRouter(Router *gin.RouterGroup) {
	userAddressRouter := Router.Group("userAddress").Use(middleware.OperationRecord())
	var userAddressApi = v1.ApiGroupApp.ShopApiGroup.UserAddressApi
	{
		userAddressRouter.POST("createUserAddress", userAddressApi.CreateUserAddress)             // 新建UserAddress
		userAddressRouter.DELETE("deleteUserAddress", userAddressApi.DeleteUserAddress)           // 删除UserAddress
		userAddressRouter.DELETE("deleteUserAddressByIds", userAddressApi.DeleteUserAddressByIds) // 批量删除UserAddress
		userAddressRouter.PUT("updateUserAddress", userAddressApi.UpdateUserAddress)              // 更新UserAddress
		userAddressRouter.GET("findUserAddress", userAddressApi.FindUserAddress)                  // 根据ID获取UserAddress
		userAddressRouter.GET("getUserAddressList", userAddressApi.GetUserAddressList)            // 获取UserAddress列表
	}
}
