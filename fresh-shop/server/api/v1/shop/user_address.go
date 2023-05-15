package shop

import (
	"fresh-shop/server/global"
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/common/response"
	"fresh-shop/server/model/shop"
	shopReq "fresh-shop/server/model/shop/request"
	"fresh-shop/server/service"
	"fresh-shop/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserAddressApi struct {
}

var userAddressService = service.ServiceGroupApp.ShopServiceGroup.UserAddressService

// CreateUserAddress 创建UserAddress
// @Tags UserAddress
// @Summary 创建UserAddress
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.UserAddress true "创建UserAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userAddress/createUserAddress [post]
func (userAddressApi *UserAddressApi) CreateUserAddress(c *gin.Context) {
	var userAddress shop.UserAddress
	err := c.ShouldBindJSON(&userAddress)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userAddress.UserId = utils.Pointer(int(utils.GetUserID(c)))
	if address, err := userAddressService.CreateUserAddress(userAddress); err != nil {
		global.Log.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithData(address, c)
	}
}

// DeleteUserAddress 删除UserAddress
// @Tags UserAddress
// @Summary 删除UserAddress
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.UserAddress true "删除UserAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userAddress/deleteUserAddress [delete]
func (userAddressApi *UserAddressApi) DeleteUserAddress(c *gin.Context) {
	var userAddress shop.UserAddress
	err := c.ShouldBindJSON(&userAddress)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userAddress.UserId = utils.Pointer(int(utils.GetUserID(c)))
	if err := userAddressService.DeleteUserAddress(userAddress); err != nil {
		global.Log.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteUserAddressByIds 批量删除UserAddress
// @Tags UserAddress
// @Summary 批量删除UserAddress
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /userAddress/deleteUserAddressByIds [delete]
func (userAddressApi *UserAddressApi) DeleteUserAddressByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userAddressService.DeleteUserAddressByIds(IDS); err != nil {
		global.Log.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateUserAddress 更新UserAddress
// @Tags UserAddress
// @Summary 更新UserAddress
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.UserAddress true "更新UserAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userAddress/updateUserAddress [put]
func (userAddressApi *UserAddressApi) UpdateUserAddress(c *gin.Context) {
	var userAddress shop.UserAddress
	err := c.ShouldBindJSON(&userAddress)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userAddress.UserId = utils.Pointer(int(utils.GetUserID(c)))
	if err := userAddressService.UpdateUserAddress(userAddress); err != nil {
		global.Log.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindUserAddress 用id查询UserAddress
// @Tags UserAddress
// @Summary 用id查询UserAddress
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shop.UserAddress true "用id查询UserAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userAddress/findUserAddress [get]
func (userAddressApi *UserAddressApi) FindUserAddress(c *gin.Context) {
	var userAddress shop.UserAddress
	err := c.ShouldBindQuery(&userAddress)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reuserAddress, err := userAddressService.GetUserAddress(userAddress.ID); err != nil {
		global.Log.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reuserAddress": reuserAddress}, c)
	}
}

// FindUserDefaultAddress 查询用户默认 UserAddress
// @Tags UserAddress
// @Summary 用id查询UserAddress
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shop.UserAddress true "用id查询UserAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userAddress/findUserAddress [get]
func (userAddressApi *UserAddressApi) FindUserDefaultAddress(c *gin.Context) {
	var userAddress shop.UserAddress
	err := c.ShouldBindQuery(&userAddress)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if reuserAddress, err := userAddressService.GetUserDeafultAddress(userId); err != nil {
		global.Log.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(reuserAddress, c)
	}
}

// GetUserAddressList 分页获取UserAddress列表
// @Tags UserAddress
// @Summary 分页获取UserAddress列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shopReq.UserAddressSearch true "分页获取UserAddress列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userAddress/getUserAddressList [get]
func (userAddressApi *UserAddressApi) GetUserAddressList(c *gin.Context) {
	var pageInfo shopReq.UserAddressSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	pageInfo.UserId = utils.Pointer(int(userId))
	if list, err := userAddressService.GetUserAddressInfoList(pageInfo); err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: list,
		}, "获取成功", c)
	}
}
