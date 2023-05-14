import service from '@/utils/request'

// @Tags UserAddress
// @Summary 创建UserAddress
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserAddress true "创建UserAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userAddress/createUserAddress [post]
export const createUserAddress = (data) => {
  return service({
    url: '/userAddress/createUserAddress',
    method: 'post',
    data
  })
}

// @Tags UserAddress
// @Summary 删除UserAddress
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserAddress true "删除UserAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userAddress/deleteUserAddress [delete]
export const deleteUserAddress = (data) => {
  return service({
    url: '/userAddress/deleteUserAddress',
    method: 'delete',
    data
  })
}

// @Tags UserAddress
// @Summary 删除UserAddress
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userAddress/deleteUserAddress [delete]
export const deleteUserAddressByIds = (data) => {
  return service({
    url: '/userAddress/deleteUserAddressByIds',
    method: 'delete',
    data
  })
}

// @Tags UserAddress
// @Summary 更新UserAddress
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserAddress true "更新UserAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userAddress/updateUserAddress [put]
export const updateUserAddress = (data) => {
  return service({
    url: '/userAddress/updateUserAddress',
    method: 'put',
    data
  })
}

// @Tags UserAddress
// @Summary 用id查询UserAddress
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.UserAddress true "用id查询UserAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userAddress/findUserAddress [get]
export const findUserAddress = (params) => {
  return service({
    url: '/userAddress/findUserAddress',
    method: 'get',
    params
  })
}

// @Tags UserAddress
// @Summary 分页获取UserAddress列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取UserAddress列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userAddress/getUserAddressList [get]
export const getUserAddressList = (params) => {
  return service({
    url: '/userAddress/getUserAddressList',
    method: 'get',
    params
  })
}
