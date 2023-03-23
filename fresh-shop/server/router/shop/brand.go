package shop

import (
	"fresh-shop/server/api/v1"
	"fresh-shop/server/middleware"
	"github.com/gin-gonic/gin"
)

type BrandRouter struct {
}

// InitBrandRouter 初始化 Brand 路由信息
func (s *BrandRouter) InitBrandRouter(Router *gin.RouterGroup) {
	brandRouter := Router.Group("brand").Use(middleware.OperationRecord())
	var brandApi = v1.ApiGroupApp.ShopApiGroup.BrandApi
	{
		brandRouter.POST("createBrand", brandApi.CreateBrand)             // 新建Brand
		brandRouter.DELETE("deleteBrand", brandApi.DeleteBrand)           // 删除Brand
		brandRouter.DELETE("deleteBrandByIds", brandApi.DeleteBrandByIds) // 批量删除Brand
		brandRouter.PUT("updateBrand", brandApi.UpdateBrand)              // 更新Brand
	}
}

// InitBrandPublicRouter 初始化公开的 Brand 路由信息
func (s *BrandRouter) InitBrandPublicRouter(Router *gin.RouterGroup) {
	brandRouterWithoutRecord := Router.Group("brand")
	var brandApi = v1.ApiGroupApp.ShopApiGroup.BrandApi
	{
		brandRouterWithoutRecord.GET("findBrand", brandApi.FindBrand)             // 根据ID获取Brand
		brandRouterWithoutRecord.GET("getBrandList", brandApi.GetBrandList)       // 获取Brand列表
		brandRouterWithoutRecord.GET("getBrandListAll", brandApi.GetBrandListAll) // 获取所有Brand列表
	}
}
