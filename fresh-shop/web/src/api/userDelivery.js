import service from '@/utils/request'

// @Tags UserDelivery
// @Summary 创建UserDelivery
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserDelivery true "创建UserDelivery"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userDelivery/createUserDelivery [post]
export const createUserDelivery = (data) => {
  return service({
    url: '/userDelivery/createUserDelivery',
    method: 'post',
    data
  })
}

// @Tags UserDelivery
// @Summary 删除UserDelivery
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserDelivery true "删除UserDelivery"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userDelivery/deleteUserDelivery [delete]
export const deleteUserDelivery = (data) => {
  return service({
    url: '/userDelivery/deleteUserDelivery',
    method: 'delete',
    data
  })
}

// @Tags UserDelivery
// @Summary 删除UserDelivery
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserDelivery"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userDelivery/deleteUserDelivery [delete]
export const deleteUserDeliveryByIds = (data) => {
  return service({
    url: '/userDelivery/deleteUserDeliveryByIds',
    method: 'delete',
    data
  })
}

// @Tags UserDelivery
// @Summary 更新UserDelivery
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserDelivery true "更新UserDelivery"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userDelivery/updateUserDelivery [put]
export const updateUserDelivery = (data) => {
  return service({
    url: '/userDelivery/updateUserDelivery',
    method: 'put',
    data
  })
}

// @Tags UserDelivery
// @Summary 用id查询UserDelivery
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.UserDelivery true "用id查询UserDelivery"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userDelivery/findUserDelivery [get]
export const findUserDelivery = (params) => {
  return service({
    url: '/userDelivery/findUserDelivery',
    method: 'get',
    params
  })
}

// @Tags UserDelivery
// @Summary 分页获取UserDelivery列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取UserDelivery列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userDelivery/getUserDeliveryList [get]
export const getUserDeliveryList = (params) => {
  return service({
    url: '/userDelivery/getUserDeliveryList',
    method: 'get',
    params
  })
}
