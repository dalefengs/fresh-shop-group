import service from '@/utils/request'

// @Tags AccountGroup
// @Summary 创建AccountGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountGroup true "创建AccountGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userAccountGroup/createAccountGroup [post]
export const createAccountGroup = (data) => {
  return service({
    url: '/userAccountGroup/createAccountGroup',
    method: 'post',
    data
  })
}

// @Tags AccountGroup
// @Summary 删除AccountGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountGroup true "删除AccountGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userAccountGroup/deleteAccountGroup [delete]
export const deleteAccountGroup = (data) => {
  return service({
    url: '/userAccountGroup/deleteAccountGroup',
    method: 'delete',
    data
  })
}

// 同步账户
export const syncAccountGroup = (data) => {
  return service({
    url: '/userAccountGroup/syncAccountGroup',
    method: 'POST',
    data
  })
}

// @Tags AccountGroup
// @Summary 删除AccountGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AccountGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userAccountGroup/deleteAccountGroup [delete]
export const deleteAccountGroupByIds = (data) => {
  return service({
    url: '/userAccountGroup/deleteAccountGroupByIds',
    method: 'delete',
    data
  })
}

// @Tags AccountGroup
// @Summary 更新AccountGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountGroup true "更新AccountGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userAccountGroup/updateAccountGroup [put]
export const updateAccountGroup = (data) => {
  return service({
    url: '/userAccountGroup/updateAccountGroup',
    method: 'put',
    data
  })
}

// @Tags AccountGroup
// @Summary 用id查询AccountGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AccountGroup true "用id查询AccountGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userAccountGroup/findAccountGroup [get]
export const findAccountGroup = (params) => {
  return service({
    url: '/userAccountGroup/findAccountGroup',
    method: 'get',
    params
  })
}

// @Tags AccountGroup
// @Summary 分页获取AccountGroup列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AccountGroup列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userAccountGroup/getAccountGroupList [get]
export const getAccountGroupList = (params) => {
  return service({
    url: '/userAccountGroup/getAccountGroupList',
    method: 'get',
    params
  })
}
