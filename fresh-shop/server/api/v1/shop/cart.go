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

type CartApi struct {
}

var cartService = service.ServiceGroupApp.ShopServiceGroup.CartService

// CreateCart 添加购物车 Cart
// @Tags Cart
// @Summary 创建Cart
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.Cart true "创建Cart"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cart/createCart [post]
func (cartApi *CartApi) CreateCart(c *gin.Context) {
	var cart shop.Cart
	err := c.ShouldBindJSON(&cart)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if cart.GoodsId == nil || cart.Num == 0 {
		global.Log.Error("参数错误!", zap.Error(err))
		response.FailWithMessage("参数错误", c)
	}
	userId := utils.GetUserID(c)
	cart.UserId = utils.Pointer(int(userId))
	if err := cartService.CreateCart(cart); err != nil {
		global.Log.Error(err.Error(), zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// DeleteCart 删除Cart
// @Tags Cart
// @Summary 删除Cart
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.Cart true "删除Cart"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cart/deleteCart [delete]
func (cartApi *CartApi) DeleteCart(c *gin.Context) {
	var cart shop.Cart
	err := c.ShouldBindJSON(&cart)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := cartService.DeleteCart(cart); err != nil {
		global.Log.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCartByIds 批量删除Cart
// @Tags Cart
// @Summary 批量删除Cart
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Cart"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cart/deleteCartByIds [delete]
func (cartApi *CartApi) DeleteCartByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := cartService.DeleteCartByIds(IDS); err != nil {
		global.Log.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCart 更新Cart
// @Tags Cart
// @Summary 更新Cart
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.Cart true "更新Cart"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cart/updateCart [put]
func (cartApi *CartApi) UpdateCart(c *gin.Context) {
	var cart shop.Cart
	err := c.ShouldBindJSON(&cart)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := cartService.UpdateCart(cart); err != nil {
		global.Log.Error("更新失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// SelectAllChecked 全选 Cart
// @Tags Cart
// @Summary 全选 Cart
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.Cart true "全选 Cart"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"全选成功"}"
// @Router /cart/selectAllChecked [post]
func (cartApi *CartApi) SelectAllChecked(c *gin.Context) {
	userId := utils.GetUserID(c)
	if err := cartService.SelectAllChecked(userId); err != nil {
		global.Log.Error("全选失败!", zap.Error(err))
		response.FailWithMessage("全选失败", c)
	} else {
		response.OkWithMessage("全选成功", c)
	}
}

// ClearAllChecked 取消全选 Cart
// @Tags Cart
// @Summary 取消全选 Cart
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.Cart true "取消全选 Cart"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"取消全选成功"}"
// @Router /cart/clearAllChecked [post]
func (cartApi *CartApi) ClearAllChecked(c *gin.Context) {
	userId := utils.GetUserID(c)
	if err := cartService.ClearAllChecked(userId); err != nil {
		global.Log.Error("取消全选失败!", zap.Error(err))
		response.FailWithMessage("取消全选失败", c)
	} else {
		response.OkWithMessage("取消全选成功", c)
	}
}

// FindCart 用id查询Cart
// @Tags Cart
// @Summary 用id查询Cart
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shop.Cart true "用id查询Cart"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cart/findCart [get]
func (cartApi *CartApi) FindCart(c *gin.Context) {
	var cart shop.Cart
	err := c.ShouldBindQuery(&cart)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recart, err := cartService.GetCart(cart.ID); err != nil {
		global.Log.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recart": recart}, c)
	}
}

// GetCartList 分页获取Cart列表
// @Tags Cart
// @Summary 分页获取Cart列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shopReq.CartSearch true "分页获取Cart列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cart/getCartList [get]
func (cartApi *CartApi) GetCartList(c *gin.Context) {
	var pageInfo shopReq.CartSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if list, total, err := cartService.GetCartInfoList(pageInfo, userId); err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:  list,
			Total: total,
		}, "获取成功", c)
	}
}
