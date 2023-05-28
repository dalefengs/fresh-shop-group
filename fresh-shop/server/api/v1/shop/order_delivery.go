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

type OrderDeliveryApi struct {
}

var orderDeliveryService = service.ServiceGroupApp.ShopServiceGroup.OrderDeliveryService

// CreateOrderDelivery 创建OrderDelivery
// @Tags OrderDelivery
// @Summary 创建OrderDelivery
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.OrderDelivery true "创建OrderDelivery"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderDelivery/createOrderDelivery [post]
func (orderDeliveryApi *OrderDeliveryApi) CreateOrderDelivery(c *gin.Context) {
	var orderDelivery shop.OrderDelivery
	err := c.ShouldBindJSON(&orderDelivery)
	if err != nil {
		global.SugarLog.Errorf("解析json 失败 %s \n", err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	if orderDelivery.OrderId == nil && *orderDelivery.OrderId == 0 {
		global.SugarLog.Errorf("发货失败! 订单id参数错误 orderId: %d", *orderDelivery.OrderId)
		response.FailWithMessage("参数错误", c)
	}
	if err := orderDeliveryService.CreateOrderDelivery(orderDelivery); err != nil {
		global.Log.Error("发货失败!", zap.Error(err))
		response.FailWithMessage("发货失败", c)
	} else {
		response.OkWithMessage("发货成功", c)
	}
}

// DeleteOrderDelivery 删除OrderDelivery
// @Tags OrderDelivery
// @Summary 删除OrderDelivery
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.OrderDelivery true "删除OrderDelivery"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /orderDelivery/deleteOrderDelivery [delete]
func (orderDeliveryApi *OrderDeliveryApi) DeleteOrderDelivery(c *gin.Context) {
	var orderDelivery shop.OrderDelivery
	err := c.ShouldBindJSON(&orderDelivery)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderDeliveryService.DeleteOrderDelivery(orderDelivery); err != nil {
		global.Log.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOrderDeliveryByIds 批量删除OrderDelivery
// @Tags OrderDelivery
// @Summary 批量删除OrderDelivery
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OrderDelivery"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /orderDelivery/deleteOrderDeliveryByIds [delete]
func (orderDeliveryApi *OrderDeliveryApi) DeleteOrderDeliveryByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderDeliveryService.DeleteOrderDeliveryByIds(IDS); err != nil {
		global.Log.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOrderDelivery 更新OrderDelivery
// @Tags OrderDelivery
// @Summary 更新OrderDelivery
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.OrderDelivery true "更新OrderDelivery"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /orderDelivery/updateOrderDelivery [put]
func (orderDeliveryApi *OrderDeliveryApi) UpdateOrderDelivery(c *gin.Context) {
	var orderDelivery shop.OrderDelivery
	err := c.ShouldBindJSON(&orderDelivery)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderDeliveryService.UpdateOrderDelivery(orderDelivery); err != nil {
		global.Log.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOrderDelivery 用id查询OrderDelivery
// @Tags OrderDelivery
// @Summary 用id查询OrderDelivery
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shop.OrderDelivery true "用id查询OrderDelivery"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /orderDelivery/findOrderDelivery [get]
func (orderDeliveryApi *OrderDeliveryApi) FindOrderDelivery(c *gin.Context) {
	var orderDelivery shop.OrderDelivery
	err := c.ShouldBindQuery(&orderDelivery)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reorderDelivery, err := orderDeliveryService.GetOrderDelivery(orderDelivery.ID, *orderDelivery.OrderId); err != nil {
		global.Log.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reorderDelivery": reorderDelivery}, c)
	}
}

// GetOrderDeliveryList 分页获取OrderDelivery列表
// @Tags OrderDelivery
// @Summary 分页获取OrderDelivery列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shopReq.OrderDeliverySearch true "分页获取OrderDelivery列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderDelivery/getOrderDeliveryList [get]
func (orderDeliveryApi *OrderDeliveryApi) GetOrderDeliveryList(c *gin.Context) {
	var pageInfo shopReq.OrderDeliverySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := orderDeliveryService.GetOrderDeliveryInfoList(pageInfo); err != nil {
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
