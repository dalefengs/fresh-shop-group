import service from '@/utils/request'

// @Tags Account
// @Summary 创建Account
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Account true "创建Account"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /account/createAccount [post]
export const createAccount = (data) => {
  return service({
    url: '/account/createAccount',
    method: 'post',
    data
  })
}

// @Tags Account
// @Summary 删除Account
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Account true "删除Account"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /account/deleteAccount [delete]
export const deleteAccount = (data) => {
  return service({
    url: '/account/deleteAccount',
    method: 'delete',
    data
  })
}

// @Tags Account
// @Summary 删除Account
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Account"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /account/deleteAccount [delete]
export const deleteAccountByIds = (data) => {
  return service({
    url: '/account/deleteAccountByIds',
    method: 'delete',
    data
  })
}

// @Tags Account
// @Summary 更新Account
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Account true "更新Account"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /account/updateAccount [put]
export const updateAccount = (data) => {
  return service({
    url: '/account/updateAccount',
    method: 'put',
    data
  })
}

// @Tags Account
// @Summary 用id查询Account
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Account true "用id查询Account"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /account/findAccount [get]
export const findAccount = (params) => {
  return service({
    url: '/account/findAccount',
    method: 'get',
    params
  })
}

// @Tags Account
// @Summary 分页获取Account列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Account列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /account/getAccountList [get]
export const getAccountList = (params) => {
  return service({
    url: '/account/getAccountList',
    method: 'get',
    params
  })
}
