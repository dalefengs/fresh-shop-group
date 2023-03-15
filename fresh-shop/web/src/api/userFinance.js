import service from '@/utils/request'

// @Tags UserFinanceCash
// @Summary 创建UserFinanceCash
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserFinanceCash true "创建UserFinanceCash"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userFinanceCash/createUserFinanceCash [post]
export const createUserFinanceCash = (data) => {
  return service({
    url: '/userFinanceCash/createUserFinanceCash',
    method: 'post',
    data
  })
}

// @Tags UserFinanceCash
// @Summary 删除UserFinanceCash
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserFinanceCash true "删除UserFinanceCash"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userFinanceCash/deleteUserFinanceCash [delete]
export const deleteUserFinanceCash = (data) => {
  return service({
    url: '/userFinanceCash/deleteUserFinanceCash',
    method: 'delete',
    data
  })
}

// @Tags UserFinanceCash
// @Summary 删除UserFinanceCash
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserFinanceCash"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userFinanceCash/deleteUserFinanceCash [delete]
export const deleteUserFinanceCashByIds = (data) => {
  return service({
    url: '/userFinanceCash/deleteUserFinanceCashByIds',
    method: 'delete',
    data
  })
}

// @Tags UserFinanceCash
// @Summary 更新UserFinanceCash
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserFinanceCash true "更新UserFinanceCash"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userFinanceCash/updateUserFinanceCash [put]
export const updateUserFinanceCash = (data) => {
  return service({
    url: '/userFinanceCash/updateUserFinanceCash',
    method: 'put',
    data
  })
}

// @Tags UserFinanceCash
// @Summary 用id查询UserFinanceCash
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.UserFinanceCash true "用id查询UserFinanceCash"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userFinanceCash/findUserFinanceCash [get]
export const findUserFinanceCash = (params) => {
  return service({
    url: '/userFinanceCash/findUserFinanceCash',
    method: 'get',
    params
  })
}

// @Tags UserFinanceCash
// @Summary 分页获取UserFinanceCash列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取UserFinanceCash列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userFinanceCash/getUserFinanceCashList [get]
export const getUserFinanceCashList = (params) => {
  return service({
    url: '/userFinanceCash/getUserFinanceCashList',
    method: 'get',
    params
  })
}
