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

type OrderReturnApi struct {
}

var orderReturnService = service.ServiceGroupApp.ShopServiceGroup.OrderReturnService


// CreateOrderReturn 创建OrderReturn
// @Tags OrderReturn
// @Summary 创建OrderReturn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.OrderReturn true "创建OrderReturn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderReturn/createOrderReturn [post]
func (orderReturnApi *OrderReturnApi) CreateOrderReturn(c *gin.Context) {
	var orderReturn shop.OrderReturn
	err := c.ShouldBindJSON(&orderReturn)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderReturnService.CreateOrderReturn(orderReturn); err != nil {
        global.Log.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOrderReturn 删除OrderReturn
// @Tags OrderReturn
// @Summary 删除OrderReturn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.OrderReturn true "删除OrderReturn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /orderReturn/deleteOrderReturn [delete]
func (orderReturnApi *OrderReturnApi) DeleteOrderReturn(c *gin.Context) {
	var orderReturn shop.OrderReturn
	err := c.ShouldBindJSON(&orderReturn)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderReturnService.DeleteOrderReturn(orderReturn); err != nil {
        global.Log.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOrderReturnByIds 批量删除OrderReturn
// @Tags OrderReturn
// @Summary 批量删除OrderReturn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OrderReturn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /orderReturn/deleteOrderReturnByIds [delete]
func (orderReturnApi *OrderReturnApi) DeleteOrderReturnByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderReturnService.DeleteOrderReturnByIds(IDS); err != nil {
        global.Log.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOrderReturn 更新OrderReturn
// @Tags OrderReturn
// @Summary 更新OrderReturn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.OrderReturn true "更新OrderReturn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /orderReturn/updateOrderReturn [put]
func (orderReturnApi *OrderReturnApi) UpdateOrderReturn(c *gin.Context) {
	var orderReturn shop.OrderReturn
	err := c.ShouldBindJSON(&orderReturn)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := orderReturnService.UpdateOrderReturn(orderReturn); err != nil {
        global.Log.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOrderReturn 用id查询OrderReturn
// @Tags OrderReturn
// @Summary 用id查询OrderReturn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shop.OrderReturn true "用id查询OrderReturn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /orderReturn/findOrderReturn [get]
func (orderReturnApi *OrderReturnApi) FindOrderReturn(c *gin.Context) {
	var orderReturn shop.OrderReturn
	err := c.ShouldBindQuery(&orderReturn)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reorderReturn, err := orderReturnService.GetOrderReturn(orderReturn.ID); err != nil {
        global.Log.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reorderReturn": reorderReturn}, c)
	}
}

// GetOrderReturnList 分页获取OrderReturn列表
// @Tags OrderReturn
// @Summary 分页获取OrderReturn列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shopReq.OrderReturnSearch true "分页获取OrderReturn列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderReturn/getOrderReturnList [get]
func (orderReturnApi *OrderReturnApi) GetOrderReturnList(c *gin.Context) {
	var pageInfo shopReq.OrderReturnSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := orderReturnService.GetOrderReturnInfoList(pageInfo); err != nil {
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
