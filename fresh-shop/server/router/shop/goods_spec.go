package shop

import (
	"fresh-shop/server/api/v1"
	"fresh-shop/server/middleware"
	"github.com/gin-gonic/gin"
)

type GoodsSpecRouter struct {
}

// InitGoodsSpecRouter 初始化 GoodsSpec 路由信息
func (s *GoodsSpecRouter) InitGoodsSpecRouter(Router *gin.RouterGroup) {
	goodsSpecRouter := Router.Group("goodsSpec").Use(middleware.OperationRecord())
	goodsSpecRouterWithoutRecord := Router.Group("goodsSpec")
	var goodsSpecApi = v1.ApiGroupApp.ShopApiGroup.GoodsSpecApi
	{
		goodsSpecRouter.POST("createGoodsSpec", goodsSpecApi.CreateGoodsSpec)   // 新建GoodsSpec
		goodsSpecRouter.DELETE("deleteGoodsSpec", goodsSpecApi.DeleteGoodsSpec) // 删除GoodsSpec
		goodsSpecRouter.DELETE("deleteGoodsSpecByIds", goodsSpecApi.DeleteGoodsSpecByIds) // 批量删除GoodsSpec
		goodsSpecRouter.PUT("updateGoodsSpec", goodsSpecApi.UpdateGoodsSpec)    // 更新GoodsSpec
	}
	{
		goodsSpecRouterWithoutRecord.GET("findGoodsSpec", goodsSpecApi.FindGoodsSpec)        // 根据ID获取GoodsSpec
		goodsSpecRouterWithoutRecord.GET("getGoodsSpecList", goodsSpecApi.GetGoodsSpecList)  // 获取GoodsSpec列表
	}
}
