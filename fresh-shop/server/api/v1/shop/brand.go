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

type BrandApi struct {
}

var brandService = service.ServiceGroupApp.ShopServiceGroup.BrandService

// CreateBrand 创建Brand
// @Tags Brand
// @Summary 创建Brand
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.Brand true "创建Brand"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /brand/createBrand [post]
func (brandApi *BrandApi) CreateBrand(c *gin.Context) {
	var brand shop.Brand
	err := c.ShouldBindJSON(&brand)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := brandService.CreateBrand(brand); err != nil {
		global.Log.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBrand 删除Brand
// @Tags Brand
// @Summary 删除Brand
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.Brand true "删除Brand"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /brand/deleteBrand [delete]
func (brandApi *BrandApi) DeleteBrand(c *gin.Context) {
	var brand shop.Brand
	err := c.ShouldBindJSON(&brand)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := brandService.DeleteBrand(brand); err != nil {
		global.Log.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBrandByIds 批量删除Brand
// @Tags Brand
// @Summary 批量删除Brand
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Brand"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /brand/deleteBrandByIds [delete]
func (brandApi *BrandApi) DeleteBrandByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := brandService.DeleteBrandByIds(IDS); err != nil {
		global.Log.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBrand 更新Brand
// @Tags Brand
// @Summary 更新Brand
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.Brand true "更新Brand"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /brand/updateBrand [put]
func (brandApi *BrandApi) UpdateBrand(c *gin.Context) {
	var brand shop.Brand
	err := c.ShouldBindJSON(&brand)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := brandService.UpdateBrand(brand); err != nil {
		global.Log.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBrand 用id查询Brand
// @Tags Brand
// @Summary 用id查询Brand
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shop.Brand true "用id查询Brand"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /brand/findBrand [get]
func (brandApi *BrandApi) FindBrand(c *gin.Context) {
	var brand shop.Brand
	err := c.ShouldBindQuery(&brand)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rebrand, err := brandService.GetBrand(brand.ID); err != nil {
		global.Log.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebrand": rebrand}, c)
	}
}

// GetBrandList 分页获取Brand列表
// @Tags Brand
// @Summary 分页获取Brand列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shopReq.BrandSearch true "分页获取Brand列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /brand/getBrandList [get]
func (brandApi *BrandApi) GetBrandList(c *gin.Context) {
	var pageInfo shopReq.BrandSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := brandService.GetBrandInfoList(pageInfo); err != nil {
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

// GetBrandListAll 获取所有Brand列表
// @Tags Brand
// @Summary 获取所有Brand列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shopReq.BrandSearch true "获取所有Brand列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /brand/getBrandListAll [get]
func (brandApi *BrandApi) GetBrandListAll(c *gin.Context) {
	if list, err := brandService.GetBrandInfoListAll(); err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}
