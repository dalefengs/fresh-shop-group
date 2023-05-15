package system

import (
	"fresh-shop/server/global"
	businessReq "fresh-shop/server/model/business/request"
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/common/response"
	"fresh-shop/server/model/system"
	"fresh-shop/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysConfigApi struct {
}

var sysConfigService = service.ServiceGroupApp.SystemServiceGroup.SysConfigService

// CreateSysConfig 创建SysConfig
// @Tags SysConfig
// @Summary 创建SysConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SysConfig true "创建SysConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysConfig/createSysConfig [post]
func (sysConfigApi *SysConfigApi) CreateSysConfig(c *gin.Context) {
	var sysConfig system.SysConfig
	err := c.ShouldBindJSON(&sysConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sysConfigService.CreateSysConfig(sysConfig); err != nil {
		global.Log.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSysConfig 删除SysConfig
// @Tags SysConfig
// @Summary 删除SysConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SysConfig true "删除SysConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysConfig/deleteSysConfig [delete]
func (sysConfigApi *SysConfigApi) DeleteSysConfig(c *gin.Context) {
	var sysConfig system.SysConfig
	err := c.ShouldBindJSON(&sysConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sysConfigService.DeleteSysConfig(sysConfig); err != nil {
		global.Log.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSysConfigByIds 批量删除SysConfig
// @Tags SysConfig
// @Summary 批量删除SysConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysConfig/deleteSysConfigByIds [delete]
func (sysConfigApi *SysConfigApi) DeleteSysConfigByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sysConfigService.DeleteSysConfigByIds(IDS); err != nil {
		global.Log.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSysConfig 更新SysConfig
// @Tags SysConfig
// @Summary 更新SysConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SysConfig true "更新SysConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysConfig/updateSysConfig [put]
func (sysConfigApi *SysConfigApi) UpdateSysConfig(c *gin.Context) {
	var sysConfig system.SysConfig
	err := c.ShouldBindJSON(&sysConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sysConfigService.UpdateSysConfig(sysConfig); err != nil {
		global.Log.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSysConfig 用id查询SysConfig
// @Tags SysConfig
// @Summary 用id查询SysConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.SysConfig true "用id查询SysConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysConfig/findSysConfig [get]
func (sysConfigApi *SysConfigApi) FindSysConfig(c *gin.Context) {
	var sysConfig system.SysConfig
	err := c.ShouldBindQuery(&sysConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if resysConfig, err := sysConfigService.GetSysConfig(sysConfig.ID); err != nil {
		global.Log.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resysConfig": resysConfig}, c)
	}
}

// FindSysConfigByName 用 name 查询SysConfig
// @Tags SysConfig
// @Summary 用 name 查询SysConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.SysConfig true "用 name 查询 SysConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysConfig/findSysConfigByName [get]
func (sysConfigApi *SysConfigApi) FindSysConfigByName(c *gin.Context) {
	var sysConfig system.SysConfig
	err := c.ShouldBindQuery(&sysConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if resysConfig, err := sysConfigService.GetSysConfigByName(sysConfig.Name); err != nil {
		global.Log.Error("查询失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(resysConfig, c)
	}
}

// GetSysConfigList 分页获取SysConfig列表
// @Tags SysConfig
// @Summary 分页获取SysConfig列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.SysConfigSearch true "分页获取SysConfig列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysConfig/getSysConfigList [get]
func (sysConfigApi *SysConfigApi) GetSysConfigList(c *gin.Context) {
	var pageInfo businessReq.SysConfigSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := sysConfigService.GetSysConfigInfoList(pageInfo); err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
