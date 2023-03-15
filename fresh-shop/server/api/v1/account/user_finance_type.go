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

type UserFinanceTypeApi struct {
}

var userFinanceTypeService = service.ServiceGroupApp.AccountServiceGroup.UserFinanceTypeService

// CreateUserFinanceType 创建UserFinanceType
// @Tags UserFinanceType
// @Summary 创建UserFinanceType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body account.UserFinanceType true "创建UserFinanceType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userFinanceType/createUserFinanceType [post]
func (userFinanceTypeApi *UserFinanceTypeApi) CreateUserFinanceType(c *gin.Context) {
	var userFinanceType account.UserFinanceType
	err := c.ShouldBindJSON(&userFinanceType)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(userFinanceType, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userFinanceTypeService.CreateUserFinanceType(userFinanceType); err != nil {
		global.Log.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteUserFinanceType 删除UserFinanceType
// @Tags UserFinanceType
// @Summary 删除UserFinanceType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body account.UserFinanceType true "删除UserFinanceType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userFinanceType/deleteUserFinanceType [delete]
func (userFinanceTypeApi *UserFinanceTypeApi) DeleteUserFinanceType(c *gin.Context) {
	var userFinanceType account.UserFinanceType
	err := c.ShouldBindJSON(&userFinanceType)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userFinanceTypeService.DeleteUserFinanceType(userFinanceType); err != nil {
		global.Log.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteUserFinanceTypeByIds 批量删除UserFinanceType
// @Tags UserFinanceType
// @Summary 批量删除UserFinanceType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserFinanceType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /userFinanceType/deleteUserFinanceTypeByIds [delete]
func (userFinanceTypeApi *UserFinanceTypeApi) DeleteUserFinanceTypeByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userFinanceTypeService.DeleteUserFinanceTypeByIds(IDS); err != nil {
		global.Log.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateUserFinanceType 更新UserFinanceType
// @Tags UserFinanceType
// @Summary 更新UserFinanceType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body account.UserFinanceType true "更新UserFinanceType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userFinanceType/updateUserFinanceType [put]
func (userFinanceTypeApi *UserFinanceTypeApi) UpdateUserFinanceType(c *gin.Context) {
	var userFinanceType account.UserFinanceType
	err := c.ShouldBindJSON(&userFinanceType)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(userFinanceType, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userFinanceTypeService.UpdateUserFinanceType(userFinanceType); err != nil {
		global.Log.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindUserFinanceType 用id查询UserFinanceType
// @Tags UserFinanceType
// @Summary 用id查询UserFinanceType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query account.UserFinanceType true "用id查询UserFinanceType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userFinanceType/findUserFinanceType [get]
func (userFinanceTypeApi *UserFinanceTypeApi) FindUserFinanceType(c *gin.Context) {
	var userFinanceType account.UserFinanceType
	err := c.ShouldBindQuery(&userFinanceType)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reuserFinanceType, err := userFinanceTypeService.GetUserFinanceType(userFinanceType.ID); err != nil {
		global.Log.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reuserFinanceType": reuserFinanceType}, c)
	}
}

// GetUserFinanceTypeList 分页获取UserFinanceType列表
// @Tags UserFinanceType
// @Summary 分页获取UserFinanceType列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query accountReq.UserFinanceTypeSearch true "分页获取UserFinanceType列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userFinanceType/getUserFinanceTypeList [get]
func (userFinanceTypeApi *UserFinanceTypeApi) GetUserFinanceTypeList(c *gin.Context) {
	var pageInfo accountReq.UserFinanceTypeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := userFinanceTypeService.GetUserFinanceTypeInfoList(pageInfo); err != nil {
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

// GetUserFinanceTypeListAll 获取所有父级数据
func (userFinanceTypeApi *UserFinanceTypeApi) GetUserFinanceTypeListAll(c *gin.Context) {
	if list, err := userFinanceTypeService.GetUserFinanceTypeInfoListAll(); err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}
