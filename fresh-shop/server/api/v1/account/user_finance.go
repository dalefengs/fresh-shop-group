package account

import (
	"fresh-shop/server/global"
	"fresh-shop/server/model/account"
	accountReq "fresh-shop/server/model/account/request"
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/common/response"
	"fresh-shop/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserFinanceCashApi struct {
}

var userFinanceCashService = service.ServiceGroupApp.AccountServiceGroup.UserFinanceCashService

// CreateUserFinance 创建UserFinanceCash
// @Tags UserFinance
// @Summary 创建UserFinanceCash
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body account.UserFinance true "创建UserFinanceCash"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userFinanceCash/createUserFinanceCash [post]
func (userFinanceCashApi *UserFinanceCashApi) CreateUserFinance(c *gin.Context) {
	var userFinanceCash account.UserFinance
	err := c.ShouldBindJSON(&userFinanceCash)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userFinanceCashService.CreateUserFinanceCash(userFinanceCash); err != nil {
		global.Log.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteUserFinance 删除UserFinanceCash
// @Tags UserFinance
// @Summary 删除UserFinanceCash
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body account.UserFinance true "删除UserFinanceCash"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userFinanceCash/deleteUserFinanceCash [delete]
func (userFinanceCashApi *UserFinanceCashApi) DeleteUserFinance(c *gin.Context) {
	var userFinanceCash account.UserFinance
	err := c.ShouldBindJSON(&userFinanceCash)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userFinanceCashService.DeleteUserFinanceCash(userFinanceCash); err != nil {
		global.Log.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteUserFinanceByIds 批量删除UserFinanceCash
// @Tags UserFinance
// @Summary 批量删除UserFinanceCash
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserFinanceCash"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /userFinanceCash/deleteUserFinanceCashByIds [delete]
func (userFinanceCashApi *UserFinanceCashApi) DeleteUserFinanceByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userFinanceCashService.DeleteUserFinanceCashByIds(IDS); err != nil {
		global.Log.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateUserFinance 更新UserFinanceCash
// @Tags UserFinance
// @Summary 更新UserFinanceCash
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body account.UserFinance true "更新UserFinanceCash"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userFinanceCash/updateUserFinanceCash [put]
func (userFinanceCashApi *UserFinanceCashApi) UpdateUserFinance(c *gin.Context) {
	var userFinanceCash account.UserFinance
	err := c.ShouldBindJSON(&userFinanceCash)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userFinanceCashService.UpdateUserFinanceCash(userFinanceCash); err != nil {
		global.Log.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindUserFinance 用id查询UserFinanceCash
// @Tags UserFinance
// @Summary 用id查询UserFinanceCash
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query account.UserFinance true "用id查询UserFinanceCash"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userFinanceCash/findUserFinanceCash [get]
func (userFinanceCashApi *UserFinanceCashApi) FindUserFinance(c *gin.Context) {
	var userFinanceCash account.UserFinance
	err := c.ShouldBindQuery(&userFinanceCash)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reuserFinanceCash, err := userFinanceCashService.GetUserFinanceCash(userFinanceCash.ID); err != nil {
		global.Log.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reuserFinanceCash": reuserFinanceCash}, c)
	}
}

// GetUserFinanceList 分页获取UserFinanceCash列表
// @Tags UserFinance
// @Summary 分页获取UserFinanceCash列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query accountReq.UserFinanceSearch true "分页获取UserFinanceCash列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userFinanceCash/getUserFinanceCashList [get]
func (userFinanceCashApi *UserFinanceCashApi) GetUserFinanceList(c *gin.Context) {
	var pageInfo accountReq.UserFinanceSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if pageInfo.GroupNameEn == "" {
		response.FailWithMessage("账户参数异常", c)
		return
	}
	if list, total, err := userFinanceCashService.GetUserFinanceInfoList(pageInfo); err != nil {
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
