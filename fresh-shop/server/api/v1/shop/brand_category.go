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

type BrandCategoryApi struct {
}

var brandCategoryService = service.ServiceGroupApp.ShopServiceGroup.BrandCategoryService

// CreateBrandCategory 创建BrandCategory
// @Tags BrandCategory
// @Summary 创建BrandCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.BrandCategory true "创建BrandCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /brandCategory/createBrandCategory [post]
func (brandCategoryApi *BrandCategoryApi) CreateBrandCategory(c *gin.Context) {
	var brandCategory shopReq.BrandCategorySearch
	err := c.ShouldBindJSON(&brandCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if brandCategory.CategoryId == nil || *brandCategory.CategoryId == 0 {
		global.Log.Error("参数异常!", zap.Error(err))
		response.FailWithMessage("参数异常", c)
	}
	if err := brandCategoryService.CreateBrandCategory(brandCategory); err != nil {
		global.Log.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// DeleteBrandCategory 删除BrandCategory
// @Tags BrandCategory
// @Summary 删除BrandCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.BrandCategory true "删除BrandCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /brandCategory/deleteBrandCategory [delete]
func (brandCategoryApi *BrandCategoryApi) DeleteBrandCategory(c *gin.Context) {
	var brandCategory shop.BrandCategory
	err := c.ShouldBindJSON(&brandCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := brandCategoryService.DeleteBrandCategory(brandCategory); err != nil {
		global.Log.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBrandCategoryByIds 批量删除BrandCategory
// @Tags BrandCategory
// @Summary 批量删除BrandCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BrandCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /brandCategory/deleteBrandCategoryByIds [delete]
func (brandCategoryApi *BrandCategoryApi) DeleteBrandCategoryByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := brandCategoryService.DeleteBrandCategoryByIds(IDS); err != nil {
		global.Log.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBrandCategory 更新BrandCategory
// @Tags BrandCategory
// @Summary 更新BrandCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.BrandCategory true "更新BrandCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /brandCategory/updateBrandCategory [put]
func (brandCategoryApi *BrandCategoryApi) UpdateBrandCategory(c *gin.Context) {
	var brandCategory shop.BrandCategory
	err := c.ShouldBindJSON(&brandCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := brandCategoryService.UpdateBrandCategory(brandCategory); err != nil {
		global.Log.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}
