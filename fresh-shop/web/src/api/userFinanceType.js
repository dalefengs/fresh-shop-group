import service from '@/utils/request'

// @Tags UserFinanceType
// @Summary 创建UserFinanceType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserFinanceType true "创建UserFinanceType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userFinanceType/createUserFinanceType [post]
export const createUserFinanceType = (data) => {
  return service({
    url: '/userFinanceType/createUserFinanceType',
    method: 'post',
    data
  })
}

// @Tags UserFinanceType
// @Summary 删除UserFinanceType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserFinanceType true "删除UserFinanceType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userFinanceType/deleteUserFinanceType [delete]
export const deleteUserFinanceType = (data) => {
  return service({
    url: '/userFinanceType/deleteUserFinanceType',
    method: 'delete',
    data
  })
}

// @Tags UserFinanceType
// @Summary 删除UserFinanceType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserFinanceType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userFinanceType/deleteUserFinanceType [delete]
export const deleteUserFinanceTypeByIds = (data) => {
  return service({
    url: '/userFinanceType/deleteUserFinanceTypeByIds',
    method: 'delete',
    data
  })
}

// @Tags UserFinanceType
// @Summary 更新UserFinanceType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserFinanceType true "更新UserFinanceType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userFinanceType/updateUserFinanceType [put]
export const updateUserFinanceType = (data) => {
  return service({
    url: '/userFinanceType/updateUserFinanceType',
    method: 'put',
    data
  })
}

// @Tags UserFinanceType
// @Summary 用id查询UserFinanceType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.UserFinanceType true "用id查询UserFinanceType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userFinanceType/findUserFinanceType [get]
export const findUserFinanceType = (params) => {
  return service({
    url: '/userFinanceType/findUserFinanceType',
    method: 'get',
    params
  })
}

// @Tags UserFinanceType
// @Summary 分页获取UserFinanceType列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取UserFinanceType列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userFinanceType/getUserFinanceTypeList [get]
export const getUserFinanceTypeList = (params) => {
  return service({
    url: '/userFinanceType/getUserFinanceTypeList',
    method: 'get',
    params
  })
}

// @Tags UserFinanceType
// @Summary 获取所有顶级 UserFinanceType 列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userFinanceType/getUserFinanceTypeList [get]
export const getUserFinanceTypeListAll = () => {
  return service({
    url: '/userFinanceType/getUserFinanceTypeListAll',
    method: 'get',
  })
}
