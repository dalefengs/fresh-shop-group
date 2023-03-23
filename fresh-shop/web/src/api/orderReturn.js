import service from '@/utils/request'

// @Tags OrderReturn
// @Summary 创建OrderReturn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderReturn true "创建OrderReturn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderReturn/createOrderReturn [post]
export const createOrderReturn = (data) => {
  return service({
    url: '/orderReturn/createOrderReturn',
    method: 'post',
    data
  })
}

// @Tags OrderReturn
// @Summary 删除OrderReturn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderReturn true "删除OrderReturn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /orderReturn/deleteOrderReturn [delete]
export const deleteOrderReturn = (data) => {
  return service({
    url: '/orderReturn/deleteOrderReturn',
    method: 'delete',
    data
  })
}

// @Tags OrderReturn
// @Summary 删除OrderReturn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OrderReturn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /orderReturn/deleteOrderReturn [delete]
export const deleteOrderReturnByIds = (data) => {
  return service({
    url: '/orderReturn/deleteOrderReturnByIds',
    method: 'delete',
    data
  })
}

// @Tags OrderReturn
// @Summary 更新OrderReturn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderReturn true "更新OrderReturn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /orderReturn/updateOrderReturn [put]
export const updateOrderReturn = (data) => {
  return service({
    url: '/orderReturn/updateOrderReturn',
    method: 'put',
    data
  })
}

// @Tags OrderReturn
// @Summary 用id查询OrderReturn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OrderReturn true "用id查询OrderReturn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /orderReturn/findOrderReturn [get]
export const findOrderReturn = (params) => {
  return service({
    url: '/orderReturn/findOrderReturn',
    method: 'get',
    params
  })
}

// @Tags OrderReturn
// @Summary 分页获取OrderReturn列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取OrderReturn列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderReturn/getOrderReturnList [get]
export const getOrderReturnList = (params) => {
  return service({
    url: '/orderReturn/getOrderReturnList',
    method: 'get',
    params
  })
}
