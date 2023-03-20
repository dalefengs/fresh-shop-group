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

type GoodsImageApi struct {
}

var goodsImageService = service.ServiceGroupApp.ShopServiceGroup.GoodsImageService

// CreateGoodsImage 创建GoodsImage
// @Tags GoodsImage
// @Summary 创建GoodsImage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.GoodsImage true "创建GoodsImage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goodsImage/createGoodsImage [post]
func (goodsImageApi *GoodsImageApi) CreateGoodsImage(c *gin.Context) {
	var goodsImage shop.GoodsImage
	err := c.ShouldBindJSON(&goodsImage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := goodsImageService.CreateGoodsImage(goodsImage); err != nil {
		global.Log.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteGoodsImage 删除GoodsImage
// @Tags GoodsImage
// @Summary 删除GoodsImage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.GoodsImage true "删除GoodsImage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goodsImage/deleteGoodsImage [delete]
func (goodsImageApi *GoodsImageApi) DeleteGoodsImage(c *gin.Context) {
	var goodsImage shop.GoodsImage
	err := c.ShouldBindJSON(&goodsImage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := goodsImageService.DeleteGoodsImage(goodsImage); err != nil {
		global.Log.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteGoodsImageByIds 批量删除GoodsImage
// @Tags GoodsImage
// @Summary 批量删除GoodsImage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除GoodsImage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /goodsImage/deleteGoodsImageByIds [delete]
func (goodsImageApi *GoodsImageApi) DeleteGoodsImageByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := goodsImageService.DeleteGoodsImageByIds(IDS); err != nil {
		global.Log.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateGoodsImage 更新GoodsImage
// @Tags GoodsImage
// @Summary 更新GoodsImage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.GoodsImage true "更新GoodsImage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /goodsImage/updateGoodsImage [put]
func (goodsImageApi *GoodsImageApi) UpdateGoodsImage(c *gin.Context) {
	var goodsImage shop.GoodsImage
	err := c.ShouldBindJSON(&goodsImage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := goodsImageService.UpdateGoodsImage(goodsImage); err != nil {
		global.Log.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}
