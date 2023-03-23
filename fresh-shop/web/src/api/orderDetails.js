import service from '@/utils/request'

// @Tags OrderDetails
// @Summary 创建OrderDetails
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderDetails true "创建OrderDetails"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderDetails/createOrderDetails [post]
export const createOrderDetails = (data) => {
  return service({
    url: '/orderDetails/createOrderDetails',
    method: 'post',
    data
  })
}

// @Tags OrderDetails
// @Summary 删除OrderDetails
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderDetails true "删除OrderDetails"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /orderDetails/deleteOrderDetails [delete]
export const deleteOrderDetails = (data) => {
  return service({
    url: '/orderDetails/deleteOrderDetails',
    method: 'delete',
    data
  })
}

// @Tags OrderDetails
// @Summary 删除OrderDetails
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OrderDetails"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /orderDetails/deleteOrderDetails [delete]
export const deleteOrderDetailsByIds = (data) => {
  return service({
    url: '/orderDetails/deleteOrderDetailsByIds',
    method: 'delete',
    data
  })
}

// @Tags OrderDetails
// @Summary 更新OrderDetails
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderDetails true "更新OrderDetails"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /orderDetails/updateOrderDetails [put]
export const updateOrderDetails = (data) => {
  return service({
    url: '/orderDetails/updateOrderDetails',
    method: 'put',
    data
  })
}

// @Tags OrderDetails
// @Summary 用id查询OrderDetails
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OrderDetails true "用id查询OrderDetails"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /orderDetails/findOrderDetails [get]
export const findOrderDetails = (params) => {
  return service({
    url: '/orderDetails/findOrderDetails',
    method: 'get',
    params
  })
}

// @Tags OrderDetails
// @Summary 分页获取OrderDetails列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取OrderDetails列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderDetails/getOrderDetailsList [get]
export const getOrderDetailsList = (params) => {
  return service({
    url: '/orderDetails/getOrderDetailsList',
    method: 'get',
    params
  })
}
