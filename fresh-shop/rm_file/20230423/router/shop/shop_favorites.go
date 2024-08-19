package shop

import (
	"fresh-shop/server/api/v1"
	"fresh-shop/server/middleware"
	"github.com/gin-gonic/gin"
)

type ShopFavoritesRouter struct {
}

// InitShopFavoritesRouter 初始化 ShopFavorites 路由信息
func (s *ShopFavoritesRouter) InitShopFavoritesRouter(Router *gin.RouterGroup) {
	shopFavoritesRouter := Router.Group("shopFavorites").Use(middleware.OperationRecord())
	shopFavoritesRouterWithoutRecord := Router.Group("shopFavorites")
	var shopFavoritesApi = v1.ApiGroupApp.ShopApiGroup.ShopFavoritesApi
	{
		shopFavoritesRouter.POST("createShopFavorites", shopFavoritesApi.CreateShopFavorites)             // 新建ShopFavorites
		shopFavoritesRouter.DELETE("deleteShopFavorites", shopFavoritesApi.DeleteShopFavorites)           // 删除ShopFavorites
		shopFavoritesRouter.DELETE("deleteShopFavoritesByIds", shopFavoritesApi.DeleteShopFavoritesByIds) // 批量删除ShopFavorites
		shopFavoritesRouter.PUT("updateShopFavorites", shopFavoritesApi.UpdateShopFavorites)              // 更新ShopFavorites
	}
	{
		shopFavoritesRouterWithoutRecord.GET("findShopFavorites", shopFavoritesApi.FindShopFavorites)       // 根据ID获取ShopFavorites
		shopFavoritesRouterWithoutRecord.GET("getShopFavoritesList", shopFavoritesApi.GetShopFavoritesList) // 获取ShopFavorites列表
	}
}
