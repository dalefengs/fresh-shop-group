package account

import (
	"fresh-shop/server/api/v1"
	"fresh-shop/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysRechargeRouter struct {
}

// InitSysRechargeRouter 初始化 SysRecharge 路由信息
func (s *SysRechargeRouter) InitSysRechargeRouter(Router *gin.RouterGroup) {
	sysRechargeRouter := Router.Group("sysRecharge").Use(middleware.OperationRecord())
	sysRechargeRouterWithoutRecord := Router.Group("sysRecharge")
	var sysRechargeApi = v1.ApiGroupApp.AccountApiGroup.SysRechargeApi
	{
		sysRechargeRouter.POST("createSysRecharge", sysRechargeApi.CreateSysRecharge)             // 新建SysRecharge
		sysRechargeRouter.DELETE("deleteSysRecharge", sysRechargeApi.DeleteSysRecharge)           // 删除SysRecharge
		sysRechargeRouter.DELETE("deleteSysRechargeByIds", sysRechargeApi.DeleteSysRechargeByIds) // 批量删除SysRecharge
		sysRechargeRouter.PUT("updateSysRecharge", sysRechargeApi.UpdateSysRecharge)              // 更新SysRecharge
	}
	{
		sysRechargeRouterWithoutRecord.GET("findSysRecharge", sysRechargeApi.FindSysRecharge)       // 根据ID获取SysRecharge
		sysRechargeRouterWithoutRecord.GET("getSysRechargeList", sysRechargeApi.GetSysRechargeList) // 获取SysRecharge列表
	}
}
