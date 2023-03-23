package shop

import (
	"fresh-shop/server/api/v1"
	"fresh-shop/server/middleware"
	"github.com/gin-gonic/gin"
)

type OrderReturnRouter struct {
}

// InitOrderReturnRouter 初始化 OrderReturn 路由信息
func (s *OrderReturnRouter) InitOrderReturnRouter(Router *gin.RouterGroup) {
	orderReturnRouter := Router.Group("orderReturn").Use(middleware.OperationRecord())
	orderReturnRouterWithoutRecord := Router.Group("orderReturn")
	var orderReturnApi = v1.ApiGroupApp.ShopApiGroup.OrderReturnApi
	{
		orderReturnRouter.POST("createOrderReturn", orderReturnApi.CreateOrderReturn)   // 新建OrderReturn
		orderReturnRouter.DELETE("deleteOrderReturn", orderReturnApi.DeleteOrderReturn) // 删除OrderReturn
		orderReturnRouter.DELETE("deleteOrderReturnByIds", orderReturnApi.DeleteOrderReturnByIds) // 批量删除OrderReturn
		orderReturnRouter.PUT("updateOrderReturn", orderReturnApi.UpdateOrderReturn)    // 更新OrderReturn
	}
	{
		orderReturnRouterWithoutRecord.GET("findOrderReturn", orderReturnApi.FindOrderReturn)        // 根据ID获取OrderReturn
		orderReturnRouterWithoutRecord.GET("getOrderReturnList", orderReturnApi.GetOrderReturnList)  // 获取OrderReturn列表
	}
}
