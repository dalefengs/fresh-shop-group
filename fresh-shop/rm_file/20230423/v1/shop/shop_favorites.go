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

type ShopFavoritesApi struct {
}

var shopFavoritesService = service.ServiceGroupApp.ShopServiceGroup.ShopFavoritesService

// CreateShopFavorites 创建ShopFavorites
// @Tags ShopFavorites
// @Summary 创建ShopFavorites
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.ShopFavorites true "创建ShopFavorites"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shopFavorites/createShopFavorites [post]
func (shopFavoritesApi *ShopFavoritesApi) CreateShopFavorites(c *gin.Context) {
	var shopFavorites shop.ShopFavorites
	err := c.ShouldBindJSON(&shopFavorites)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := shopFavoritesService.CreateShopFavorites(shopFavorites); err != nil {
		global.Log.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteShopFavorites 删除ShopFavorites
// @Tags ShopFavorites
// @Summary 删除ShopFavorites
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.ShopFavorites true "删除ShopFavorites"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shopFavorites/deleteShopFavorites [delete]
func (shopFavoritesApi *ShopFavoritesApi) DeleteShopFavorites(c *gin.Context) {
	var shopFavorites shop.ShopFavorites
	err := c.ShouldBindJSON(&shopFavorites)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := shopFavoritesService.DeleteShopFavorites(shopFavorites); err != nil {
		global.Log.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteShopFavoritesByIds 批量删除ShopFavorites
// @Tags ShopFavorites
// @Summary 批量删除ShopFavorites
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ShopFavorites"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /shopFavorites/deleteShopFavoritesByIds [delete]
func (shopFavoritesApi *ShopFavoritesApi) DeleteShopFavoritesByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := shopFavoritesService.DeleteShopFavoritesByIds(IDS); err != nil {
		global.Log.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateShopFavorites 更新ShopFavorites
// @Tags ShopFavorites
// @Summary 更新ShopFavorites
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.ShopFavorites true "更新ShopFavorites"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /shopFavorites/updateShopFavorites [put]
func (shopFavoritesApi *ShopFavoritesApi) UpdateShopFavorites(c *gin.Context) {
	var shopFavorites shop.ShopFavorites
	err := c.ShouldBindJSON(&shopFavorites)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := shopFavoritesService.UpdateShopFavorites(shopFavorites); err != nil {
		global.Log.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindShopFavorites 用id查询ShopFavorites
// @Tags ShopFavorites
// @Summary 用id查询ShopFavorites
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shop.ShopFavorites true "用id查询ShopFavorites"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /shopFavorites/findShopFavorites [get]
func (shopFavoritesApi *ShopFavoritesApi) FindShopFavorites(c *gin.Context) {
	var shopFavorites shop.ShopFavorites
	err := c.ShouldBindQuery(&shopFavorites)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reshopFavorites, err := shopFavoritesService.GetShopFavorites(shopFavorites.ID); err != nil {
		global.Log.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reshopFavorites": reshopFavorites}, c)
	}
}

// GetShopFavoritesList 分页获取ShopFavorites列表
// @Tags ShopFavorites
// @Summary 分页获取ShopFavorites列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shopReq.ShopFavoritesSearch true "分页获取ShopFavorites列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shopFavorites/getShopFavoritesList [get]
func (shopFavoritesApi *ShopFavoritesApi) GetShopFavoritesList(c *gin.Context) {
	var pageInfo shopReq.ShopFavoritesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := shopFavoritesService.GetShopFavoritesInfoList(pageInfo); err != nil {
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
