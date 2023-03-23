package shop

import (
	"fresh-shop/server/api/v1"
	"fresh-shop/server/middleware"
	"github.com/gin-gonic/gin"
)

type CategoryRouter struct {
}

// InitCategoryRouter 初始化 Category 路由信息
func (s *CategoryRouter) InitCategoryRouter(Router *gin.RouterGroup) {
	categoryRouter := Router.Group("category").Use(middleware.OperationRecord())
	var categoryApi = v1.ApiGroupApp.ShopApiGroup.CategoryApi
	{
		categoryRouter.POST("createCategory", categoryApi.CreateCategory)             // 新建Category
		categoryRouter.DELETE("deleteCategory", categoryApi.DeleteCategory)           // 删除Category
		categoryRouter.DELETE("deleteCategoryByIds", categoryApi.DeleteCategoryByIds) // 批量删除Category
		categoryRouter.PUT("updateCategory", categoryApi.UpdateCategory)              // 更新Category
	}
}

// InitCategoryPublicRouter 初始化公开的 Category 路由信息
func (s *CategoryRouter) InitCategoryPublicRouter(Router *gin.RouterGroup) {
	categoryRouterWithoutRecord := Router.Group("category")
	var categoryApi = v1.ApiGroupApp.ShopApiGroup.CategoryApi
	{
		categoryRouterWithoutRecord.GET("findCategory", categoryApi.FindCategory)             // 根据ID获取Category
		categoryRouterWithoutRecord.GET("getCategoryList", categoryApi.GetCategoryList)       // 获取Category列表
		categoryRouterWithoutRecord.GET("getCategoryListAll", categoryApi.GetCategoryListAll) // 获取所有Category列表
	}
}
