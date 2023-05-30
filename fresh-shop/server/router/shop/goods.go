package shop

import (
	"fresh-shop/server/api/v1"
	"fresh-shop/server/middleware"
	"github.com/gin-gonic/gin"
)

type GoodsRouter struct {
}

// InitGoodsRouter 初始化 Goods 路由信息
func (s *GoodsRouter) InitGoodsRouter(Router *gin.RouterGroup) {
	goodsRouter := Router.Group("goods").Use(middleware.OperationRecord())
	var goodsApi = v1.ApiGroupApp.ShopApiGroup.GoodsApi
	{
		goodsRouter.POST("createGoods", goodsApi.CreateGoods)                         // 新建Goods
		goodsRouter.DELETE("deleteGoods", goodsApi.DeleteGoods)                       // 删除Goods
		goodsRouter.DELETE("deleteGoodsByIds", goodsApi.DeleteGoodsByIds)             // 批量删除Goods
		goodsRouter.PUT("updateGoods", goodsApi.UpdateGoods)                          // 更新Goods
		goodsRouter.POST("batchCreateGoodsByExcel", goodsApi.BatchCreateGoodsByExcel) // 批量导入商品信息
	}
}

// InitGoodsPublicRouter 初始化公开的 Goods 路由信息
func (s *GoodsRouter) InitGoodsPublicRouter(Router *gin.RouterGroup) {
	goodsRouterWithoutRecord := Router.Group("goods")
	var goodsApi = v1.ApiGroupApp.ShopApiGroup.GoodsApi
	{
		goodsRouterWithoutRecord.GET("findGoods", goodsApi.FindGoods)       // 根据ID获取Goods
		goodsRouterWithoutRecord.GET("getGoodsList", goodsApi.GetGoodsList) // 获取Goods列表
	}
}
