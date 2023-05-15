package system

import (
	"fresh-shop/server/api/v1"
	"fresh-shop/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysConfigRouter struct {
}

// InitSysConfigRouter 初始化 SysConfig 路由信息
func (s *SysConfigRouter) InitSysConfigRouter(Router *gin.RouterGroup) {
	sysConfigRouter := Router.Group("sysConfig").Use(middleware.OperationRecord())
	sysConfigRouterWithoutRecord := Router.Group("sysConfig")
	var sysConfigApi = v1.ApiGroupApp.SystemApiGroup.SysConfigApi
	{
		sysConfigRouter.POST("createSysConfig", sysConfigApi.CreateSysConfig)             // 新建SysConfig
		sysConfigRouter.DELETE("deleteSysConfig", sysConfigApi.DeleteSysConfig)           // 删除SysConfig
		sysConfigRouter.DELETE("deleteSysConfigByIds", sysConfigApi.DeleteSysConfigByIds) // 批量删除SysConfig
		sysConfigRouter.PUT("updateSysConfig", sysConfigApi.UpdateSysConfig)              // 更新SysConfig
	}
	{
		sysConfigRouterWithoutRecord.GET("findSysConfig", sysConfigApi.FindSysConfig)       // 根据ID获取SysConfig
		sysConfigRouterWithoutRecord.GET("getSysConfigList", sysConfigApi.GetSysConfigList) // 获取SysConfig列表
	}
}

// InitSysConfigPublicRouter 初始化公开的 SysConfig 路由信息
func (s *SysConfigRouter) InitSysConfigPublicRouter(Router *gin.RouterGroup) {
	var sysConfigApi = v1.ApiGroupApp.SystemApiGroup.SysConfigApi
	sysConfigRouterWithoutRecord := Router.Group("sysConfig")
	{
		sysConfigRouterWithoutRecord.GET("findSysConfigByName", sysConfigApi.FindSysConfigByName) // 根据 key 获取SysConfig
	}
}
