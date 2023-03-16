package account

import (
	"fresh-shop/server/global"
	"fresh-shop/server/model/account"
	accountReq "fresh-shop/server/model/account/request"
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/common/response"
	"fresh-shop/server/service"
	"fresh-shop/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysRechargeApi struct {
}

var sysRechargeService = service.ServiceGroupApp.AccountServiceGroup.SysRechargeService

// CreateSysRecharge 创建SysRecharge
// @Tags SysRecharge
// @Summary 创建SysRecharge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body account.SysRecharge true "创建SysRecharge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysRecharge/createSysRecharge [post]
func (sysRechargeApi *SysRechargeApi) CreateSysRecharge(c *gin.Context) {
	var sysRecharge account.SysRecharge
	err := c.ShouldBindJSON(&sysRecharge)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	claims, err := utils.GetClaims(c)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if *sysRecharge.Amount == 0 {
		response.FailWithMessage("请输入正确的金额", c)
		return
	}
	if err := sysRechargeService.CreateSysRecharge(sysRecharge, claims); err != nil {
		global.Log.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("操作成功", c)
	}
}

// DeleteSysRecharge 删除SysRecharge
// @Tags SysRecharge
// @Summary 删除SysRecharge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body account.SysRecharge true "删除SysRecharge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysRecharge/deleteSysRecharge [delete]
func (sysRechargeApi *SysRechargeApi) DeleteSysRecharge(c *gin.Context) {
	var sysRecharge account.SysRecharge
	err := c.ShouldBindJSON(&sysRecharge)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sysRechargeService.DeleteSysRecharge(sysRecharge); err != nil {
		global.Log.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSysRechargeByIds 批量删除SysRecharge
// @Tags SysRecharge
// @Summary 批量删除SysRecharge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysRecharge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysRecharge/deleteSysRechargeByIds [delete]
func (sysRechargeApi *SysRechargeApi) DeleteSysRechargeByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sysRechargeService.DeleteSysRechargeByIds(IDS); err != nil {
		global.Log.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSysRecharge 更新SysRecharge
// @Tags SysRecharge
// @Summary 更新SysRecharge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body account.SysRecharge true "更新SysRecharge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysRecharge/updateSysRecharge [put]
func (sysRechargeApi *SysRechargeApi) UpdateSysRecharge(c *gin.Context) {
	var sysRecharge account.SysRecharge
	err := c.ShouldBindJSON(&sysRecharge)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Amount": {utils.NotEmpty()},
	}
	if err := utils.Verify(sysRecharge, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sysRechargeService.UpdateSysRecharge(sysRecharge); err != nil {
		global.Log.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSysRecharge 用id查询SysRecharge
// @Tags SysRecharge
// @Summary 用id查询SysRecharge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query account.SysRecharge true "用id查询SysRecharge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysRecharge/findSysRecharge [get]
func (sysRechargeApi *SysRechargeApi) FindSysRecharge(c *gin.Context) {
	var sysRecharge account.SysRecharge
	err := c.ShouldBindQuery(&sysRecharge)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if resysRecharge, err := sysRechargeService.GetSysRecharge(sysRecharge.ID); err != nil {
		global.Log.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resysRecharge": resysRecharge}, c)
	}
}

// GetSysRechargeList 分页获取SysRecharge列表
// @Tags SysRecharge
// @Summary 分页获取SysRecharge列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query accountReq.SysRechargeSearch true "分页获取SysRecharge列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysRecharge/getSysRechargeList [get]
func (sysRechargeApi *SysRechargeApi) GetSysRechargeList(c *gin.Context) {
	var pageInfo accountReq.SysRechargeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := sysRechargeService.GetSysRechargeInfoList(pageInfo); err != nil {
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
