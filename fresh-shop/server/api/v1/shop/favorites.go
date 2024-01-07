package shop

import (
	"fresh-shop/server/global"
	businessReq "fresh-shop/server/model/business/request"
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/common/response"
	"fresh-shop/server/model/shop"
	"fresh-shop/server/service"
	"fresh-shop/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FavoritesApi struct {
}

var favoritesService = service.ServiceGroupApp.ShopServiceGroup.FavoritesService

// Favorites 创建Favorites
// @Tags Favorites
// @Summary 创建Favorites
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.Favorites true "创建Favorites"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /favorites/favorites [post]
func (favoritesApi *FavoritesApi) Favorites(c *gin.Context) {
	userId := utils.GetUserID(c)
	var favorites shop.Favorites
	err := c.ShouldBindJSON(&favorites)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if favorites.GoodsId == nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	favorites.UserId = utils.Pointer(int(userId))
	if err := favoritesService.Favorites(favorites); err != nil {
		global.Log.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("操作失败", c)
	} else {
		response.OkWithMessage("操作成功", c)
	}
}

// DeleteFavorites 删除Favorites
// @Tags Favorites
// @Summary 删除Favorites
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.Favorites true "删除Favorites"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /favorites/deleteFavorites [delete]
func (favoritesApi *FavoritesApi) DeleteFavorites(c *gin.Context) {
	var favorites shop.Favorites
	err := c.ShouldBindJSON(&favorites)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := favoritesService.DeleteFavorites(favorites); err != nil {
		global.Log.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFavoritesByIds 批量删除Favorites
// @Tags Favorites
// @Summary 批量删除Favorites
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Favorites"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /favorites/deleteFavoritesByIds [delete]
func (favoritesApi *FavoritesApi) DeleteFavoritesByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := favoritesService.DeleteFavoritesByIds(IDS); err != nil {
		global.Log.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// FindFavorites 用id查询Favorites
// @Tags Favorites
// @Summary 用id查询Favorites
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.Favorites true "用id查询Favorites"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /favorites/findFavorites [get]
func (favoritesApi *FavoritesApi) FindFavorites(c *gin.Context) {
	var favorites shop.Favorites
	err := c.ShouldBindQuery(&favorites)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if refavorites, err := favoritesService.GetFavorites(favorites.ID); err != nil {
		global.Log.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refavorites": refavorites}, c)
	}
}

// GetFavoritesList 分页获取Favorites列表
// @Tags Favorites
// @Summary 分页获取Favorites列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.FavoritesSearch true "分页获取Favorites列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /favorites/getFavoritesList [get]
func (favoritesApi *FavoritesApi) GetFavoritesList(c *gin.Context) {
	var pageInfo businessReq.FavoritesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	pageInfo.UserId = utils.Pointer(int(userId))
	if list, total, err := favoritesService.GetFavoritesInfoList(pageInfo); err != nil {
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
