import service from '@/utils/request'

// @Tags OrderReturnDetails
// @Summary 创建OrderReturnDetails
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderReturnDetails true "创建OrderReturnDetails"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderReturnDetails/createOrderReturnDetails [post]
export const createOrderReturnDetails = (data) => {
  return service({
    url: '/orderReturnDetails/createOrderReturnDetails',
    method: 'post',
    data
  })
}

// @Tags OrderReturnDetails
// @Summary 删除OrderReturnDetails
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderReturnDetails true "删除OrderReturnDetails"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /orderReturnDetails/deleteOrderReturnDetails [delete]
export const deleteOrderReturnDetails = (data) => {
  return service({
    url: '/orderReturnDetails/deleteOrderReturnDetails',
    method: 'delete',
    data
  })
}

// @Tags OrderReturnDetails
// @Summary 删除OrderReturnDetails
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OrderReturnDetails"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /orderReturnDetails/deleteOrderReturnDetails [delete]
export const deleteOrderReturnDetailsByIds = (data) => {
  return service({
    url: '/orderReturnDetails/deleteOrderReturnDetailsByIds',
    method: 'delete',
    data
  })
}

// @Tags OrderReturnDetails
// @Summary 更新OrderReturnDetails
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderReturnDetails true "更新OrderReturnDetails"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /orderReturnDetails/updateOrderReturnDetails [put]
export const updateOrderReturnDetails = (data) => {
  return service({
    url: '/orderReturnDetails/updateOrderReturnDetails',
    method: 'put',
    data
  })
}

// @Tags OrderReturnDetails
// @Summary 用id查询OrderReturnDetails
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OrderReturnDetails true "用id查询OrderReturnDetails"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /orderReturnDetails/findOrderReturnDetails [get]
export const findOrderReturnDetails = (params) => {
  return service({
    url: '/orderReturnDetails/findOrderReturnDetails',
    method: 'get',
    params
  })
}

// @Tags OrderReturnDetails
// @Summary 分页获取OrderReturnDetails列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取OrderReturnDetails列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderReturnDetails/getOrderReturnDetailsList [get]
export const getOrderReturnDetailsList = (params) => {
  return service({
    url: '/orderReturnDetails/getOrderReturnDetailsList',
    method: 'get',
    params
  })
}
