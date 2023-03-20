package shop

import (
	"fresh-shop/server/global"
    "fresh-shop/server/model/shop"
    "fresh-shop/server/model/common/request"
    shopReq "fresh-shop/server/model/shop/request"
    "fresh-shop/server/model/common/response"
    "fresh-shop/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type GoodsDescriptionApi struct {
}

var goodsDescriptionService = service.ServiceGroupApp.ShopServiceGroup.GoodsDescriptionService


// CreateGoodsDescription 创建GoodsDescription
// @Tags GoodsDescription
// @Summary 创建GoodsDescription
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.GoodsDescription true "创建GoodsDescription"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goodsDescription/createGoodsDescription [post]
func (goodsDescriptionApi *GoodsDescriptionApi) CreateGoodsDescription(c *gin.Context) {
	var goodsDescription shop.GoodsDescription
	err := c.ShouldBindJSON(&goodsDescription)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := goodsDescriptionService.CreateGoodsDescription(goodsDescription); err != nil {
        global.Log.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteGoodsDescription 删除GoodsDescription
// @Tags GoodsDescription
// @Summary 删除GoodsDescription
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.GoodsDescription true "删除GoodsDescription"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goodsDescription/deleteGoodsDescription [delete]
func (goodsDescriptionApi *GoodsDescriptionApi) DeleteGoodsDescription(c *gin.Context) {
	var goodsDescription shop.GoodsDescription
	err := c.ShouldBindJSON(&goodsDescription)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := goodsDescriptionService.DeleteGoodsDescription(goodsDescription); err != nil {
        global.Log.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteGoodsDescriptionByIds 批量删除GoodsDescription
// @Tags GoodsDescription
// @Summary 批量删除GoodsDescription
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除GoodsDescription"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /goodsDescription/deleteGoodsDescriptionByIds [delete]
func (goodsDescriptionApi *GoodsDescriptionApi) DeleteGoodsDescriptionByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := goodsDescriptionService.DeleteGoodsDescriptionByIds(IDS); err != nil {
        global.Log.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateGoodsDescription 更新GoodsDescription
// @Tags GoodsDescription
// @Summary 更新GoodsDescription
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.GoodsDescription true "更新GoodsDescription"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /goodsDescription/updateGoodsDescription [put]
func (goodsDescriptionApi *GoodsDescriptionApi) UpdateGoodsDescription(c *gin.Context) {
	var goodsDescription shop.GoodsDescription
	err := c.ShouldBindJSON(&goodsDescription)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := goodsDescriptionService.UpdateGoodsDescription(goodsDescription); err != nil {
        global.Log.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindGoodsDescription 用id查询GoodsDescription
// @Tags GoodsDescription
// @Summary 用id查询GoodsDescription
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shop.GoodsDescription true "用id查询GoodsDescription"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /goodsDescription/findGoodsDescription [get]
func (goodsDescriptionApi *GoodsDescriptionApi) FindGoodsDescription(c *gin.Context) {
	var goodsDescription shop.GoodsDescription
	err := c.ShouldBindQuery(&goodsDescription)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if regoodsDescription, err := goodsDescriptionService.GetGoodsDescription(goodsDescription.ID); err != nil {
        global.Log.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"regoodsDescription": regoodsDescription}, c)
	}
}

// GetGoodsDescriptionList 分页获取GoodsDescription列表
// @Tags GoodsDescription
// @Summary 分页获取GoodsDescription列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shopReq.GoodsDescriptionSearch true "分页获取GoodsDescription列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goodsDescription/getGoodsDescriptionList [get]
func (goodsDescriptionApi *GoodsDescriptionApi) GetGoodsDescriptionList(c *gin.Context) {
	var pageInfo shopReq.GoodsDescriptionSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := goodsDescriptionService.GetGoodsDescriptionInfoList(pageInfo); err != nil {
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
