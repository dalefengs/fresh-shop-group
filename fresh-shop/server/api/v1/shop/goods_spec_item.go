package shop

import (
	"fresh-shop/server/global"
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/common/response"
	"fresh-shop/server/model/shop"
	"fresh-shop/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GoodsSpecItemApi struct {
}

var goodsSpecItemService = service.ServiceGroupApp.ShopServiceGroup.GoodsSpecItemService

// CreateGoodsSpecItem 创建GoodsSpecItem
// @Tags GoodsSpecItem
// @Summary 创建GoodsSpecItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.GoodsSpecItem true "创建GoodsSpecItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goodsSpecItem/createGoodsSpecItem [post]
func (goodsSpecItemApi *GoodsSpecItemApi) CreateGoodsSpecItem(c *gin.Context) {
	var goodsSpecItem shop.GoodsSpecItem
	err := c.ShouldBindJSON(&goodsSpecItem)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := goodsSpecItemService.CreateGoodsSpecItem(goodsSpecItem); err != nil {
		global.Log.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteGoodsSpecItem 删除GoodsSpecItem
// @Tags GoodsSpecItem
// @Summary 删除GoodsSpecItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.GoodsSpecItem true "删除GoodsSpecItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goodsSpecItem/deleteGoodsSpecItem [delete]
func (goodsSpecItemApi *GoodsSpecItemApi) DeleteGoodsSpecItem(c *gin.Context) {
	var goodsSpecItem shop.GoodsSpecItem
	err := c.ShouldBindJSON(&goodsSpecItem)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := goodsSpecItemService.DeleteGoodsSpecItem(goodsSpecItem); err != nil {
		global.Log.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteGoodsSpecItemByIds 批量删除GoodsSpecItem
// @Tags GoodsSpecItem
// @Summary 批量删除GoodsSpecItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除GoodsSpecItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /goodsSpecItem/deleteGoodsSpecItemByIds [delete]
func (goodsSpecItemApi *GoodsSpecItemApi) DeleteGoodsSpecItemByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := goodsSpecItemService.DeleteGoodsSpecItemByIds(IDS); err != nil {
		global.Log.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateGoodsSpecItem 更新GoodsSpecItem
// @Tags GoodsSpecItem
// @Summary 更新GoodsSpecItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.GoodsSpecItem true "更新GoodsSpecItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /goodsSpecItem/updateGoodsSpecItem [put]
func (goodsSpecItemApi *GoodsSpecItemApi) UpdateGoodsSpecItem(c *gin.Context) {
	var goodsSpecItem shop.GoodsSpecItem
	err := c.ShouldBindJSON(&goodsSpecItem)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := goodsSpecItemService.UpdateGoodsSpecItem(goodsSpecItem); err != nil {
		global.Log.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindGoodsSpecItem 用id查询GoodsSpecItem
// @Tags GoodsSpecItem
// @Summary 用id查询GoodsSpecItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shop.GoodsSpecItem true "用id查询GoodsSpecItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /goodsSpecItem/findGoodsSpecItem [get]
func (goodsSpecItemApi *GoodsSpecItemApi) FindGoodsSpecItem(c *gin.Context) {
	var goodsSpecItem shop.GoodsSpecItem
	err := c.ShouldBindQuery(&goodsSpecItem)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if regoodsSpecItem, err := goodsSpecItemService.GetGoodsSpecItem(goodsSpecItem.ID); err != nil {
		global.Log.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"regoodsSpecItem": regoodsSpecItem}, c)
	}
}
