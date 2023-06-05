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

type OrderApi struct {
}

var orderService = service.ServiceGroupApp.ShopServiceGroup.OrderService

// CreateOrder 创建待支付订单 Order
// @Tags Order
// @Summary 创建Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.Order true "创建Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /order/createOrder [post]
func (orderApi *OrderApi) CreateOrder(c *gin.Context) {
	var order shop.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if order.AddressId == 0 && *order.ShipmentType == 0 {
		response.FailWithMessage("请选择收货地址", c)
		return
	}
	userId := utils.GetUserID(c)
	order.UserId = utils.Pointer(int(userId))
	userClaims := utils.GetUserInfo(c)
	if orderResp, err := orderService.CreateOrder(order, userClaims, c.ClientIP()); err != nil {
		global.Log.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(orderResp, c)
	}
}

// OrderPay 支付 Order, 返回微信支付所需要的参数
// @Tags Order
// @Summary 支付 Order, 返回微信支付所需要的参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.Order true "支付 Order, 返回微信支付所需要的参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /order/orderPay [post]
func (orderApi *OrderApi) OrderPay(c *gin.Context) {
	var order shop.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if order.ID == 0 {
		response.FailWithMessage("订单ID不能为空", c)
		return
	}
	userId := utils.GetUserID(c)
	order.UserId = utils.Pointer(int(userId))
	userClaims := utils.GetUserInfo(c)
	if orderResp, err := orderService.OrderPay(order, userClaims, c.ClientIP()); err != nil {
		global.Log.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(orderResp, c)
	}
}

// DeleteOrder 删除Order
// @Tags Order
// @Summary 删除Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.Order true "删除Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /order/deleteOrder [delete]
func (orderApi *OrderApi) DeleteOrder(c *gin.Context) {
	var order shop.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderService.DeleteOrder(order); err != nil {
		global.Log.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// CancelOrder 取消订单
// @Tags Order
// @Summary 取消订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.Order true "取消订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"取消成功"}"
// @Router /order/cancelOrder [delete]
func (orderApi *OrderApi) CancelOrder(c *gin.Context) {
	var order shop.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderService.CancelOrder(order); err != nil {
		global.Log.Error("取消失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("取消成功", c)
	}
}

// DeleteOrderByIds 批量删除Order
// @Tags Order
// @Summary 批量删除Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /order/deleteOrderByIds [delete]
func (orderApi *OrderApi) DeleteOrderByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderService.DeleteOrderByIds(IDS); err != nil {
		global.Log.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOrder 更新Order
// @Tags Order
// @Summary 更新Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.Order true "更新Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /order/updateOrder [put]
func (orderApi *OrderApi) UpdateOrder(c *gin.Context) {
	var order shop.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderService.UpdateOrder(order); err != nil {
		global.Log.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOrder 用id查询Order
// @Tags Order
// @Summary 用id查询Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shop.Order true "用id查询Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /order/findOrder [get]
func (orderApi *OrderApi) FindOrder(c *gin.Context) {
	var order shop.Order
	err := c.ShouldBindQuery(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reorder, err := orderService.GetOrder(order.ID); err != nil {
		global.Log.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reorder": reorder}, c)
	}
}

// FindUserOrderStatus 获取用户订单中数量
// @Tags Order
// @Summary 获取用户订单中数量
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shop.Order true "获取用户订单中数量"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /order/findUserOrderStatus [get]
func (orderApi *OrderApi) FindUserOrderStatus(c *gin.Context) {
	userId := utils.GetUserID(c)
	if reorder, err := orderService.FindUserOrderStatus(userId); err != nil {
		global.Log.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(reorder, c)
	}
}

// OrderStatus 获取订单状态
// @Tags Order
// @Summary 用id查询Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shop.Order true "获取订单状态"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /order/orderStatus [get]
func (orderApi *OrderApi) OrderStatus(c *gin.Context) {
	var order shop.Order
	err := c.ShouldBindQuery(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if order.ID == 0 {
		response.FailWithMessage("订单ID不能为空", c)
		return
	}
	if reorder, err := orderService.OrderStatus(order.ID); err != nil {
		global.Log.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(reorder, c)
	}
}

// GetOrderList 分页获取Order列表
// @Tags Order
// @Summary 分页获取Order列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shopReq.OrderSearch true "分页获取Order列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /order/getOrderList [get]
func (orderApi *OrderApi) GetOrderList(c *gin.Context) {
	var pageInfo shopReq.OrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := orderService.GetOrderInfoList(pageInfo); err != nil {
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

// GetUserOrderList 分页获取登录用户获取Order列表
// @Tags Order
// @Summary 页获取登录用户获取Order列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shopReq.OrderSearch true "页获取登录用户获取Order列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /order/getUserOrderList [get]
func (orderApi *OrderApi) GetUserOrderList(c *gin.Context) {
	var pageInfo shopReq.OrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	pageInfo.UserId = utils.Pointer(int(userId))
	if list, total, err := orderService.GetOrderInfoList(pageInfo); err != nil {
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
