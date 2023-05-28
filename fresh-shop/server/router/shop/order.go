package shop

import (
	"fresh-shop/server/api/v1"
	"fresh-shop/server/middleware"
	"github.com/gin-gonic/gin"
)

type OrderRouter struct {
}

// InitOrderRouter 初始化 Order 路由信息
func (s *OrderRouter) InitOrderRouter(Router *gin.RouterGroup) {
	orderRouter := Router.Group("order").Use(middleware.OperationRecord())
	orderRouterWithoutRecord := Router.Group("order")
	var orderApi = v1.ApiGroupApp.ShopApiGroup.OrderApi
	{
		orderRouter.POST("createOrder", orderApi.CreateOrder)             // 创建待支付 Order
		orderRouter.POST("orderPay", orderApi.OrderPay)                   // 支付 Order, 返回微信支付所需要的参数
		orderRouter.POST("cancelOrder", orderApi.CancelOrder)             // 取消订单
		orderRouter.DELETE("deleteOrder", orderApi.DeleteOrder)           // 删除 Order
		orderRouter.DELETE("deleteOrderByIds", orderApi.DeleteOrderByIds) // 批量删除 Order
		orderRouter.PUT("updateOrder", orderApi.UpdateOrder)              // 更新 Order
	}
	{
		orderRouterWithoutRecord.GET("findOrder", orderApi.FindOrder)                     // 根据ID获取Order
		orderRouterWithoutRecord.GET("findUserOrderStatus", orderApi.FindUserOrderStatus) // 获取用户订单中数量
		orderRouterWithoutRecord.GET("getOrderList", orderApi.GetOrderList)               // 获取Order列表
		orderRouterWithoutRecord.GET("getUserOrderList", orderApi.GetUserOrderList)       // 根据登录用户获取Order列表
		orderRouterWithoutRecord.GET("orderStatus", orderApi.OrderStatus)                 // 获取订单状态 Order
	}
}
