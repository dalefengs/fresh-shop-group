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

type AccountGroupApi struct {
}

var userAccountGroupService = service.ServiceGroupApp.AccountServiceGroup.AccountGroupService

// SyncAccountGroup 同步用户账户
func (ag *AccountGroupApi) SyncAccountGroup(c *gin.Context) {
	var userAccountGroup account.AccountGroup
	err := c.ShouldBindJSON(&userAccountGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err = userAccountGroupService.SyncAccountGroup(userAccountGroup); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("后台同步中！", c)
	}

}

// CreateAccountGroup 创建AccountGroup
// @Tags AccountGroup
// @Summary 创建AccountGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body account.AccountGroup true "创建AccountGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userAccountGroup/createAccountGroup [post]
func (ag *AccountGroupApi) CreateAccountGroup(c *gin.Context) {
	var userAccountGroup account.AccountGroup
	err := c.ShouldBindJSON(&userAccountGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"NameEn": {utils.NotEmpty()},
		"NameCn": {utils.NotEmpty()},
		"Places": {utils.NotEmpty()},
		"Status": {utils.NotEmpty()},
	}
	if err := utils.Verify(userAccountGroup, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userAccountGroupService.CreateAccountGroup(userAccountGroup); err != nil {
		global.Log.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAccountGroup 删除AccountGroup
// @Tags AccountGroup
// @Summary 删除AccountGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body account.AccountGroup true "删除AccountGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userAccountGroup/deleteAccountGroup [delete]
func (ag *AccountGroupApi) DeleteAccountGroup(c *gin.Context) {
	var userAccountGroup account.AccountGroup
	err := c.ShouldBindJSON(&userAccountGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userAccountGroupService.DeleteAccountGroup(userAccountGroup); err != nil {
		global.Log.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAccountGroupByIds 批量删除AccountGroup
// @Tags AccountGroup
// @Summary 批量删除AccountGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AccountGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /userAccountGroup/deleteAccountGroupByIds [delete]
func (ag *AccountGroupApi) DeleteAccountGroupByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userAccountGroupService.DeleteAccountGroupByIds(IDS); err != nil {
		global.Log.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAccountGroup 更新AccountGroup
// @Tags AccountGroup
// @Summary 更新AccountGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body account.AccountGroup true "更新AccountGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userAccountGroup/updateAccountGroup [put]
func (ag *AccountGroupApi) UpdateAccountGroup(c *gin.Context) {
	var userAccountGroup account.AccountGroup
	err := c.ShouldBindJSON(&userAccountGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"NameEn": {utils.NotEmpty()},
		"NameCn": {utils.NotEmpty()},
		"Places": {utils.NotEmpty()},
		"Status": {utils.NotEmpty()},
	}
	if err := utils.Verify(userAccountGroup, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userAccountGroupService.UpdateAccountGroup(userAccountGroup); err != nil {
		global.Log.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAccountGroup 用id查询AccountGroup
// @Tags AccountGroup
// @Summary 用id查询AccountGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query account.AccountGroup true "用id查询AccountGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userAccountGroup/findAccountGroup [get]
func (ag *AccountGroupApi) FindAccountGroup(c *gin.Context) {
	var userAccountGroup account.AccountGroup
	err := c.ShouldBindQuery(&userAccountGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reuserAccountGroup, err := userAccountGroupService.GetAccountGroup(userAccountGroup.ID); err != nil {
		global.Log.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reuserAccountGroup": reuserAccountGroup}, c)
	}
}

// GetAccountGroupList 分页获取AccountGroup列表
// @Tags AccountGroup
// @Summary 分页获取AccountGroup列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query accountReq.AccountGroupSearch true "分页获取AccountGroup列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userAccountGroup/getAccountGroupList [get]
func (ag *AccountGroupApi) GetAccountGroupList(c *gin.Context) {
	var pageInfo accountReq.AccountGroupSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := userAccountGroupService.GetAccountGroupInfoList(pageInfo); err != nil {
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

// GetAccountGroupListAll 获取所有 AccountGroup 列表
// @Tags AccountGroup
// @Summary 获取所有 AccountGroup 列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userAccountGroup/getAccountGroupListAll [get]
func (ag *AccountGroupApi) GetAccountGroupListAll(c *gin.Context) {
	var g account.AccountGroup
	err := c.ShouldBindQuery(&g)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, err := userAccountGroupService.GetAccountGroupInfoListAll(g); err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}
