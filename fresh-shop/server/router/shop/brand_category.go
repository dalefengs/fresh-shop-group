package shop

import (
	"fresh-shop/server/api/v1"
	"fresh-shop/server/middleware"
	"github.com/gin-gonic/gin"
)

type BrandCategoryRouter struct {
}

// InitBrandCategoryRouter 初始化 BrandCategory 路由信息
func (s *BrandCategoryRouter) InitBrandCategoryRouter(Router *gin.RouterGroup) {
	brandCategoryRouter := Router.Group("brandCategory").Use(middleware.OperationRecord())
	var brandCategoryApi = v1.ApiGroupApp.ShopApiGroup.BrandCategoryApi
	{
		brandCategoryRouter.POST("createBrandCategory", brandCategoryApi.CreateBrandCategory) // 新建BrandCategory
	}
}
