package shop

import (
	"fresh-shop/server/api/v1"
	"fresh-shop/server/middleware"
	"github.com/gin-gonic/gin"
)

type OrderDeliveryRouter struct {
}

// InitOrderDeliveryRouter 初始化 OrderDelivery 路由信息
func (s *OrderDeliveryRouter) InitOrderDeliveryRouter(Router *gin.RouterGroup) {
	orderDeliveryRouter := Router.Group("orderDelivery").Use(middleware.OperationRecord())
	orderDeliveryRouterWithoutRecord := Router.Group("orderDelivery")
	var orderDeliveryApi = v1.ApiGroupApp.ShopApiGroup.OrderDeliveryApi
	{
		orderDeliveryRouter.POST("createOrderDelivery", orderDeliveryApi.CreateOrderDelivery)   // 新建OrderDelivery
		orderDeliveryRouter.DELETE("deleteOrderDelivery", orderDeliveryApi.DeleteOrderDelivery) // 删除OrderDelivery
		orderDeliveryRouter.DELETE("deleteOrderDeliveryByIds", orderDeliveryApi.DeleteOrderDeliveryByIds) // 批量删除OrderDelivery
		orderDeliveryRouter.PUT("updateOrderDelivery", orderDeliveryApi.UpdateOrderDelivery)    // 更新OrderDelivery
	}
	{
		orderDeliveryRouterWithoutRecord.GET("findOrderDelivery", orderDeliveryApi.FindOrderDelivery)        // 根据ID获取OrderDelivery
		orderDeliveryRouterWithoutRecord.GET("getOrderDeliveryList", orderDeliveryApi.GetOrderDeliveryList)  // 获取OrderDelivery列表
	}
}
