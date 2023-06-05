package account

import (
	"fresh-shop/server/global"
	"fresh-shop/server/model/account"
	accountReq "fresh-shop/server/model/account/request"
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/common/response"
	"fresh-shop/server/service"
	"fresh-shop/server/service/common"
	"fresh-shop/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AccountApi struct {
}

var accountService = service.ServiceGroupApp.AccountServiceGroup.AccountService

// CreateAccount 创建Account
// @Tags Account
// @Summary 创建Account
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body account.Account true "创建Account"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /account/createAccount [post]
func (accountApi *AccountApi) CreateAccount(c *gin.Context) {
	var ac account.Account
	err := c.ShouldBindJSON(&ac)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := accountService.CreateAccount(ac); err != nil {
		global.Log.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAccount 删除Account
// @Tags Account
// @Summary 删除Account
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body account.Account true "删除Account"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /account/deleteAccount [delete]
func (accountApi *AccountApi) DeleteAccount(c *gin.Context) {
	var ac account.Account
	err := c.ShouldBindJSON(&ac)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := accountService.DeleteAccount(ac); err != nil {
		global.Log.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAccountByIds 批量删除Account
// @Tags Account
// @Summary 批量删除Account
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Account"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /account/deleteAccountByIds [delete]
func (accountApi *AccountApi) DeleteAccountByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := accountService.DeleteAccountByIds(IDS); err != nil {
		global.Log.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAccount 更新Account
// @Tags Account
// @Summary 更新Account
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body account.Account true "更新Account"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /account/updateAccount [put]
func (accountApi *AccountApi) UpdateAccount(c *gin.Context) {
	var ac account.Account
	err := c.ShouldBindJSON(&ac)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := accountService.UpdateAccount(ac); err != nil {
		global.Log.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAccount 用id查询Account
// @Tags Account
// @Summary 用id查询Account
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query account.Account true "用id查询Account"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /account/findAccount [get]
func (accountApi *AccountApi) FindAccount(c *gin.Context) {
	var ac account.Account
	err := c.ShouldBindQuery(&ac)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reaccount, err := accountService.GetAccount(ac.ID); err != nil {
		global.Log.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reaccount": reaccount}, c)
	}
}

// FindUserAccount 获取当前用户指定 Account 信息
// @Tags Account
// @Summary 获取当前用户指定 Account
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query account.Account true "获取当前用户指定 Account"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /account/findUserAccount [get]
func (accountApi *AccountApi) FindUserAccount(c *gin.Context) {
	var ac account.Account
	err := c.ShouldBindQuery(&ac)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	accInfo, err := common.GetUserAccountInfo(int(userId), int(ac.ID))
	if err != nil {
		global.Log.Error("查询账户失败!", zap.Error(err))
		response.FailWithMessage("查询账户失败", c)
	} else {
		response.OkWithData(gin.H{"account": accInfo}, c)
	}
}

// GetAccountList 分页获取Account列表
// @Tags Account
// @Summary 分页获取Account列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query accountReq.AccountSearch true "分页获取Account列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /account/getAccountList [get]
func (accountApi *AccountApi) GetAccountList(c *gin.Context) {
	var pageInfo accountReq.AccountSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := accountService.GetAccountInfoList(pageInfo); err != nil {
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
