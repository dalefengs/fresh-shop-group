package shop

import (
	"fresh-shop/server/api/v1"
	"fresh-shop/server/middleware"
	"github.com/gin-gonic/gin"
)

type TagsRouter struct {
}

// InitTagsRouter 初始化 Tags 路由信息
func (s *TagsRouter) InitTagsRouter(Router *gin.RouterGroup) {
	tagsRouter := Router.Group("tags").Use(middleware.OperationRecord())
	var tagsApi = v1.ApiGroupApp.ShopApiGroup.TagsApi
	{
		tagsRouter.POST("createTags", tagsApi.CreateTags)             // 新建Tags
		tagsRouter.DELETE("deleteTags", tagsApi.DeleteTags)           // 删除Tags
		tagsRouter.DELETE("deleteTagsByIds", tagsApi.DeleteTagsByIds) // 批量删除Tags
		tagsRouter.PUT("updateTags", tagsApi.UpdateTags)              // 更新Tags
	}
}

// InitTagsPublicRouter 初始化公开的 Tags 路由信息
func (s *TagsRouter) InitTagsPublicRouter(Router *gin.RouterGroup) {
	tagsRouterWithoutRecord := Router.Group("tags")
	var tagsApi = v1.ApiGroupApp.ShopApiGroup.TagsApi
	{
		tagsRouterWithoutRecord.GET("findTags", tagsApi.FindTags)       // 根据ID获取Tags
		tagsRouterWithoutRecord.GET("getTagsList", tagsApi.GetTagsList) // 获取Tags列表
	}
}
