package business

import (
	"fresh-shop/server/api/v1"
	"fresh-shop/server/middleware"
	"github.com/gin-gonic/gin"
)

type BannerRouter struct {
}

// InitBannerRouter 初始化 Banner 路由信息
func (s *BannerRouter) InitBannerRouter(Router *gin.RouterGroup) {
	bannerRouter := Router.Group("banner").Use(middleware.OperationRecord())
	var bannerApi = v1.ApiGroupApp.BusinessApiGroup.BannerApi
	{
		bannerRouter.POST("createBanner", bannerApi.CreateBanner)             // 新建Banner
		bannerRouter.DELETE("deleteBanner", bannerApi.DeleteBanner)           // 删除Banner
		bannerRouter.DELETE("deleteBannerByIds", bannerApi.DeleteBannerByIds) // 批量删除Banner
		bannerRouter.PUT("updateBanner", bannerApi.UpdateBanner)              // 更新Banner
	}
}

// InitBannerPublicRouter 初始化公共的 Banner 路由信息
func (s *BannerRouter) InitBannerPublicRouter(Router *gin.RouterGroup) {
	bannerRouterWithoutRecord := Router.Group("banner")
	var bannerApi = v1.ApiGroupApp.BusinessApiGroup.BannerApi
	{
		bannerRouterWithoutRecord.GET("findBanner", bannerApi.FindBanner)       // 根据ID获取Banner
		bannerRouterWithoutRecord.GET("getBannerList", bannerApi.GetBannerList) // 获取Banner列表
	}
}
