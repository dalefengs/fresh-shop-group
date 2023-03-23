import service from '@/utils/request'

// @Tags OrderDelivery
// @Summary 创建OrderDelivery
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderDelivery true "创建OrderDelivery"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderDelivery/createOrderDelivery [post]
export const createOrderDelivery = (data) => {
  return service({
    url: '/orderDelivery/createOrderDelivery',
    method: 'post',
    data
  })
}

// @Tags OrderDelivery
// @Summary 删除OrderDelivery
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderDelivery true "删除OrderDelivery"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /orderDelivery/deleteOrderDelivery [delete]
export const deleteOrderDelivery = (data) => {
  return service({
    url: '/orderDelivery/deleteOrderDelivery',
    method: 'delete',
    data
  })
}

// @Tags OrderDelivery
// @Summary 删除OrderDelivery
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OrderDelivery"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /orderDelivery/deleteOrderDelivery [delete]
export const deleteOrderDeliveryByIds = (data) => {
  return service({
    url: '/orderDelivery/deleteOrderDeliveryByIds',
    method: 'delete',
    data
  })
}

// @Tags OrderDelivery
// @Summary 更新OrderDelivery
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderDelivery true "更新OrderDelivery"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /orderDelivery/updateOrderDelivery [put]
export const updateOrderDelivery = (data) => {
  return service({
    url: '/orderDelivery/updateOrderDelivery',
    method: 'put',
    data
  })
}

// @Tags OrderDelivery
// @Summary 用id查询OrderDelivery
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OrderDelivery true "用id查询OrderDelivery"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /orderDelivery/findOrderDelivery [get]
export const findOrderDelivery = (params) => {
  return service({
    url: '/orderDelivery/findOrderDelivery',
    method: 'get',
    params
  })
}

// @Tags OrderDelivery
// @Summary 分页获取OrderDelivery列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取OrderDelivery列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderDelivery/getOrderDeliveryList [get]
export const getOrderDeliveryList = (params) => {
  return service({
    url: '/orderDelivery/getOrderDeliveryList',
    method: 'get',
    params
  })
}
