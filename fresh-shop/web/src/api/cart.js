import service from '@/utils/request'

// @Tags Cart
// @Summary 创建Cart
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cart true "创建Cart"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cart/createCart [post]
export const createCart = (data) => {
  return service({
    url: '/cart/createCart',
    method: 'post',
    data
  })
}

// @Tags Cart
// @Summary 删除Cart
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cart true "删除Cart"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cart/deleteCart [delete]
export const deleteCart = (data) => {
  return service({
    url: '/cart/deleteCart',
    method: 'delete',
    data
  })
}

// @Tags Cart
// @Summary 删除Cart
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Cart"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cart/deleteCart [delete]
export const deleteCartByIds = (data) => {
  return service({
    url: '/cart/deleteCartByIds',
    method: 'delete',
    data
  })
}

// @Tags Cart
// @Summary 更新Cart
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cart true "更新Cart"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cart/updateCart [put]
export const updateCart = (data) => {
  return service({
    url: '/cart/updateCart',
    method: 'put',
    data
  })
}

// @Tags Cart
// @Summary 用id查询Cart
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Cart true "用id查询Cart"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cart/findCart [get]
export const findCart = (params) => {
  return service({
    url: '/cart/findCart',
    method: 'get',
    params
  })
}

// @Tags Cart
// @Summary 分页获取Cart列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Cart列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cart/getCartList [get]
export const getCartList = (params) => {
  return service({
    url: '/cart/getCartList',
    method: 'get',
    params
  })
}
