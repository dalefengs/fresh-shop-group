package shop

import (
	"fresh-shop/server/global"
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/common/response"
	"fresh-shop/server/model/shop"
	shopReq "fresh-shop/server/model/shop/request"
	"fresh-shop/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GoodsSpecApi struct {
}

var goodsSpecService = service.ServiceGroupApp.ShopServiceGroup.GoodsSpecService

// CreateGoodsSpec 创建GoodsSpec
// @Tags GoodsSpec
// @Summary 创建GoodsSpec
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.GoodsSpec true "创建GoodsSpec"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goodsSpec/createGoodsSpec [post]
func (goodsSpecApi *GoodsSpecApi) CreateGoodsSpec(c *gin.Context) {
	var goodsSpec shop.GoodsSpec
	err := c.ShouldBindJSON(&goodsSpec)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := goodsSpecService.CreateGoodsSpec(goodsSpec); err != nil {
		global.Log.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteGoodsSpec 删除GoodsSpec
// @Tags GoodsSpec
// @Summary 删除GoodsSpec
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.GoodsSpec true "删除GoodsSpec"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goodsSpec/deleteGoodsSpec [delete]
func (goodsSpecApi *GoodsSpecApi) DeleteGoodsSpec(c *gin.Context) {
	var goodsSpec shop.GoodsSpec
	err := c.ShouldBindJSON(&goodsSpec)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := goodsSpecService.DeleteGoodsSpec(goodsSpec); err != nil {
		global.Log.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteGoodsSpecByIds 批量删除GoodsSpec
// @Tags GoodsSpec
// @Summary 批量删除GoodsSpec
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除GoodsSpec"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /goodsSpec/deleteGoodsSpecByIds [delete]
func (goodsSpecApi *GoodsSpecApi) DeleteGoodsSpecByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := goodsSpecService.DeleteGoodsSpecByIds(IDS); err != nil {
		global.Log.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateGoodsSpec 更新GoodsSpec
// @Tags GoodsSpec
// @Summary 更新GoodsSpec
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.GoodsSpec true "更新GoodsSpec"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /goodsSpec/updateGoodsSpec [put]
func (goodsSpecApi *GoodsSpecApi) UpdateGoodsSpec(c *gin.Context) {
	var goodsSpec shop.GoodsSpec
	err := c.ShouldBindJSON(&goodsSpec)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := goodsSpecService.UpdateGoodsSpec(goodsSpec); err != nil {
		global.Log.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindGoodsSpec 用id查询GoodsSpec
// @Tags GoodsSpec
// @Summary 用id查询GoodsSpec
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shop.GoodsSpec true "用id查询GoodsSpec"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /goodsSpec/findGoodsSpec [get]
func (goodsSpecApi *GoodsSpecApi) FindGoodsSpec(c *gin.Context) {
	var goodsSpec shop.GoodsSpec
	err := c.ShouldBindQuery(&goodsSpec)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if regoodsSpec, err := goodsSpecService.GetGoodsSpec(goodsSpec.ID); err != nil {
		global.Log.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"regoodsSpec": regoodsSpec}, c)
	}
}

// GetGoodsSpecList 分页获取GoodsSpec列表
// @Tags GoodsSpec
// @Summary 分页获取GoodsSpec列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shopReq.GoodsSpecSearch true "分页获取GoodsSpec列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goodsSpec/getGoodsSpecList [get]
func (goodsSpecApi *GoodsSpecApi) GetGoodsSpecList(c *gin.Context) {
	var pageInfo shopReq.GoodsSpecSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := goodsSpecService.GetGoodsSpecInfoList(pageInfo); err != nil {
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
