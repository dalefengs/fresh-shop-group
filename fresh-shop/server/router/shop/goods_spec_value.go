package shop

import (
	"fresh-shop/server/api/v1"
	"fresh-shop/server/middleware"
	"github.com/gin-gonic/gin"
)

type GoodsSpecValueRouter struct {
}

// InitGoodsSpecValueRouter 初始化 GoodsSpecValue 路由信息
func (s *GoodsSpecValueRouter) InitGoodsSpecValueRouter(Router *gin.RouterGroup) {
	goodsSpecValueRouter := Router.Group("goodsSpecValue").Use(middleware.OperationRecord())
	goodsSpecValueRouterWithoutRecord := Router.Group("goodsSpecValue")
	var goodsSpecValueApi = v1.ApiGroupApp.ShopApiGroup.GoodsSpecValueApi
	{
		goodsSpecValueRouter.POST("createGoodsSpecValue", goodsSpecValueApi.CreateGoodsSpecValue)   // 新建GoodsSpecValue
		goodsSpecValueRouter.DELETE("deleteGoodsSpecValue", goodsSpecValueApi.DeleteGoodsSpecValue) // 删除GoodsSpecValue
		goodsSpecValueRouter.DELETE("deleteGoodsSpecValueByIds", goodsSpecValueApi.DeleteGoodsSpecValueByIds) // 批量删除GoodsSpecValue
		goodsSpecValueRouter.PUT("updateGoodsSpecValue", goodsSpecValueApi.UpdateGoodsSpecValue)    // 更新GoodsSpecValue
	}
	{
		goodsSpecValueRouterWithoutRecord.GET("findGoodsSpecValue", goodsSpecValueApi.FindGoodsSpecValue)        // 根据ID获取GoodsSpecValue
		goodsSpecValueRouterWithoutRecord.GET("getGoodsSpecValueList", goodsSpecValueApi.GetGoodsSpecValueList)  // 获取GoodsSpecValue列表
	}
}
