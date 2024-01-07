package shop

import (
	"fresh-shop/server/api/v1"
	"fresh-shop/server/middleware"
	"github.com/gin-gonic/gin"
)

type OrderDetailsRouter struct {
}

// InitOrderDetailsRouter 初始化 OrderDetails 路由信息
func (s *OrderDetailsRouter) InitOrderDetailsRouter(Router *gin.RouterGroup) {
	orderDetailsRouter := Router.Group("orderDetails").Use(middleware.OperationRecord())
	orderDetailsRouterWithoutRecord := Router.Group("orderDetails")
	var orderDetailsApi = v1.ApiGroupApp.ShopApiGroup.OrderDetailsApi
	{
		orderDetailsRouter.POST("createOrderDetails", orderDetailsApi.CreateOrderDetails)             // 新建OrderDetails
		orderDetailsRouter.DELETE("deleteOrderDetails", orderDetailsApi.DeleteOrderDetails)           // 删除OrderDetails
		orderDetailsRouter.DELETE("deleteOrderDetailsByIds", orderDetailsApi.DeleteOrderDetailsByIds) // 批量删除OrderDetails
		orderDetailsRouter.PUT("updateOrderDetails", orderDetailsApi.UpdateOrderDetails)              // 更新OrderDetails
	}
	{
		orderDetailsRouterWithoutRecord.GET("recentlyPurchasedGoods", orderDetailsApi.RecentlyPurchasedGoods) // 近期购买的商品列表
		orderDetailsRouterWithoutRecord.GET("findOrderDetails", orderDetailsApi.FindOrderDetails)             // 根据ID获取OrderDetails
		orderDetailsRouterWithoutRecord.GET("getOrderDetailsList", orderDetailsApi.GetOrderDetailsList)       // 获取OrderDetails列表
	}
}
