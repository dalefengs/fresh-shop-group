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

type GoodsApi struct {
}

var goodsService = service.ServiceGroupApp.ShopServiceGroup.GoodsService

// CreateGoods 创建Goods
// @Tags Goods
// @Summary 创建Goods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.Goods true "创建Goods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goods/createGoods [post]
func (goodsApi *GoodsApi) CreateGoods(c *gin.Context) {
	var goods shopReq.GoodsSubmitFrom
	err := c.ShouldBindJSON(&goods)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := goodsService.CreateGoods(goods); err != nil {
		global.Log.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteGoods 删除Goods
// @Tags Goods
// @Summary 删除Goods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.Goods true "删除Goods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goods/deleteGoods [delete]
func (goodsApi *GoodsApi) DeleteGoods(c *gin.Context) {
	var goods shop.Goods
	err := c.ShouldBindJSON(&goods)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := goodsService.DeleteGoods(goods); err != nil {
		global.Log.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteGoodsByIds 批量删除Goods
// @Tags Goods
// @Summary 批量删除Goods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Goods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /goods/deleteGoodsByIds [delete]
func (goodsApi *GoodsApi) DeleteGoodsByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := goodsService.DeleteGoodsByIds(IDS); err != nil {
		global.Log.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateGoods 更新Goods
// @Tags Goods
// @Summary 更新Goods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.Goods true "更新Goods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /goods/updateGoods [put]
func (goodsApi *GoodsApi) UpdateGoods(c *gin.Context) {
	var goods shop.Goods
	err := c.ShouldBindJSON(&goods)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":       {utils.NotEmpty()},
		"CategoryId": {utils.NotEmpty()},
		"BrandId":    {utils.NotEmpty()},
		"GoodsArea":  {utils.NotEmpty()},
		"SpecType":   {utils.NotEmpty()},
		"Unit":       {utils.NotEmpty()},
		"Price":      {utils.NotEmpty()},
		"MinCount":   {utils.NotEmpty()},
		"Store":      {utils.NotEmpty()},
	}
	if err := utils.Verify(goods, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := goodsService.UpdateGoods(goods); err != nil {
		global.Log.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindGoods 用id查询Goods
// @Tags Goods
// @Summary 用id查询Goods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shop.Goods true "用id查询Goods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /goods/findGoods [get]
func (goodsApi *GoodsApi) FindGoods(c *gin.Context) {
	var goods shop.Goods
	err := c.ShouldBindQuery(&goods)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if regoods, err := goodsService.GetGoods(goods.ID); err != nil {
		global.Log.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"regoods": regoods}, c)
	}
}

// GetGoodsList 分页获取Goods列表
// @Tags Goods
// @Summary 分页获取Goods列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shopReq.GoodsSearch true "分页获取Goods列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goods/getGoodsList [get]
func (goodsApi *GoodsApi) GetGoodsList(c *gin.Context) {
	var pageInfo shopReq.GoodsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := goodsService.GetGoodsInfoList(pageInfo); err != nil {
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
