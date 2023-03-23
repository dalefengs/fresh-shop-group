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

type OrderDetailsApi struct {
}

var orderDetailsService = service.ServiceGroupApp.ShopServiceGroup.OrderDetailsService


// CreateOrderDetails 创建OrderDetails
// @Tags OrderDetails
// @Summary 创建OrderDetails
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.OrderDetails true "创建OrderDetails"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderDetails/createOrderDetails [post]
func (orderDetailsApi *OrderDetailsApi) CreateOrderDetails(c *gin.Context) {
	var orderDetails shop.OrderDetails
	err := c.ShouldBindJSON(&orderDetails)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderDetailsService.CreateOrderDetails(orderDetails); err != nil {
        global.Log.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOrderDetails 删除OrderDetails
// @Tags OrderDetails
// @Summary 删除OrderDetails
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.OrderDetails true "删除OrderDetails"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /orderDetails/deleteOrderDetails [delete]
func (orderDetailsApi *OrderDetailsApi) DeleteOrderDetails(c *gin.Context) {
	var orderDetails shop.OrderDetails
	err := c.ShouldBindJSON(&orderDetails)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderDetailsService.DeleteOrderDetails(orderDetails); err != nil {
        global.Log.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOrderDetailsByIds 批量删除OrderDetails
// @Tags OrderDetails
// @Summary 批量删除OrderDetails
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OrderDetails"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /orderDetails/deleteOrderDetailsByIds [delete]
func (orderDetailsApi *OrderDetailsApi) DeleteOrderDetailsByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderDetailsService.DeleteOrderDetailsByIds(IDS); err != nil {
        global.Log.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOrderDetails 更新OrderDetails
// @Tags OrderDetails
// @Summary 更新OrderDetails
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.OrderDetails true "更新OrderDetails"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /orderDetails/updateOrderDetails [put]
func (orderDetailsApi *OrderDetailsApi) UpdateOrderDetails(c *gin.Context) {
	var orderDetails shop.OrderDetails
	err := c.ShouldBindJSON(&orderDetails)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderDetailsService.UpdateOrderDetails(orderDetails); err != nil {
        global.Log.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOrderDetails 用id查询OrderDetails
// @Tags OrderDetails
// @Summary 用id查询OrderDetails
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shop.OrderDetails true "用id查询OrderDetails"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /orderDetails/findOrderDetails [get]
func (orderDetailsApi *OrderDetailsApi) FindOrderDetails(c *gin.Context) {
	var orderDetails shop.OrderDetails
	err := c.ShouldBindQuery(&orderDetails)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reorderDetails, err := orderDetailsService.GetOrderDetails(orderDetails.ID); err != nil {
        global.Log.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reorderDetails": reorderDetails}, c)
	}
}

// GetOrderDetailsList 分页获取OrderDetails列表
// @Tags OrderDetails
// @Summary 分页获取OrderDetails列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shopReq.OrderDetailsSearch true "分页获取OrderDetails列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderDetails/getOrderDetailsList [get]
func (orderDetailsApi *OrderDetailsApi) GetOrderDetailsList(c *gin.Context) {
	var pageInfo shopReq.OrderDetailsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := orderDetailsService.GetOrderDetailsInfoList(pageInfo); err != nil {
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
