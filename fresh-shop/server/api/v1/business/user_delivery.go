package business

import (
	"fresh-shop/server/global"
	"fresh-shop/server/model/business"
	businessReq "fresh-shop/server/model/business/request"
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/common/response"
	"fresh-shop/server/service"
	"fresh-shop/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserDeliveryApi struct {
}

var userDeliveryService = service.ServiceGroupApp.BusinessServiceGroup.UserDeliveryService

// CreateUserDelivery 创建UserDelivery
// @Tags UserDelivery
// @Summary 创建UserDelivery
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.UserDelivery true "创建UserDelivery"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userDelivery/createUserDelivery [post]
func (userDeliveryApi *UserDeliveryApi) CreateUserDelivery(c *gin.Context) {
	var userDelivery business.UserDelivery
	err := c.ShouldBindJSON(&userDelivery)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":   {utils.NotEmpty()},
		"Mobile": {utils.NotEmpty()},
		"Status": {utils.NotEmpty()},
	}
	if err := utils.Verify(userDelivery, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userDeliveryService.CreateUserDelivery(userDelivery); err != nil {
		global.Log.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteUserDelivery 删除UserDelivery
// @Tags UserDelivery
// @Summary 删除UserDelivery
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.UserDelivery true "删除UserDelivery"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userDelivery/deleteUserDelivery [delete]
func (userDeliveryApi *UserDeliveryApi) DeleteUserDelivery(c *gin.Context) {
	var userDelivery business.UserDelivery
	err := c.ShouldBindJSON(&userDelivery)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userDeliveryService.DeleteUserDelivery(userDelivery); err != nil {
		global.Log.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteUserDeliveryByIds 批量删除UserDelivery
// @Tags UserDelivery
// @Summary 批量删除UserDelivery
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserDelivery"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /userDelivery/deleteUserDeliveryByIds [delete]
func (userDeliveryApi *UserDeliveryApi) DeleteUserDeliveryByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userDeliveryService.DeleteUserDeliveryByIds(IDS); err != nil {
		global.Log.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateUserDelivery 更新UserDelivery
// @Tags UserDelivery
// @Summary 更新UserDelivery
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.UserDelivery true "更新UserDelivery"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userDelivery/updateUserDelivery [put]
func (userDeliveryApi *UserDeliveryApi) UpdateUserDelivery(c *gin.Context) {
	var userDelivery business.UserDelivery
	err := c.ShouldBindJSON(&userDelivery)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":   {utils.NotEmpty()},
		"Mobile": {utils.NotEmpty()},
		"Status": {utils.NotEmpty()},
	}
	if err := utils.Verify(userDelivery, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userDeliveryService.UpdateUserDelivery(userDelivery); err != nil {
		global.Log.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindUserDelivery 用id查询UserDelivery
// @Tags UserDelivery
// @Summary 用id查询UserDelivery
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.UserDelivery true "用id查询UserDelivery"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userDelivery/findUserDelivery [get]
func (userDeliveryApi *UserDeliveryApi) FindUserDelivery(c *gin.Context) {
	var userDelivery business.UserDelivery
	err := c.ShouldBindQuery(&userDelivery)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reuserDelivery, err := userDeliveryService.GetUserDelivery(userDelivery.ID); err != nil {
		global.Log.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reuserDelivery": reuserDelivery}, c)
	}
}

// GetUserDeliveryList 分页获取UserDelivery列表
// @Tags UserDelivery
// @Summary 分页获取UserDelivery列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.UserDeliverySearch true "分页获取UserDelivery列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userDelivery/getUserDeliveryList [get]
func (userDeliveryApi *UserDeliveryApi) GetUserDeliveryList(c *gin.Context) {
	var pageInfo businessReq.UserDeliverySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := userDeliveryService.GetUserDeliveryInfoList(pageInfo); err != nil {
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
