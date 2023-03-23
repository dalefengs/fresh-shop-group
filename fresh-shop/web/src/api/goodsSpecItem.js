import service from '@/utils/request'

// @Tags GoodsSpecItem
// @Summary 创建GoodsSpecItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoodsSpecItem true "创建GoodsSpecItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goodsSpecItem/createGoodsSpecItem [post]
export const createGoodsSpecItem = (data) => {
  return service({
    url: '/goodsSpecItem/createGoodsSpecItem',
    method: 'post',
    data
  })
}

// @Tags GoodsSpecItem
// @Summary 删除GoodsSpecItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoodsSpecItem true "删除GoodsSpecItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goodsSpecItem/deleteGoodsSpecItem [delete]
export const deleteGoodsSpecItem = (data) => {
  return service({
    url: '/goodsSpecItem/deleteGoodsSpecItem',
    method: 'delete',
    data
  })
}

// @Tags GoodsSpecItem
// @Summary 删除GoodsSpecItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除GoodsSpecItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goodsSpecItem/deleteGoodsSpecItem [delete]
export const deleteGoodsSpecItemByIds = (data) => {
  return service({
    url: '/goodsSpecItem/deleteGoodsSpecItemByIds',
    method: 'delete',
    data
  })
}

// @Tags GoodsSpecItem
// @Summary 更新GoodsSpecItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoodsSpecItem true "更新GoodsSpecItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /goodsSpecItem/updateGoodsSpecItem [put]
export const updateGoodsSpecItem = (data) => {
  return service({
    url: '/goodsSpecItem/updateGoodsSpecItem',
    method: 'put',
    data
  })
}

// @Tags GoodsSpecItem
// @Summary 用id查询GoodsSpecItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.GoodsSpecItem true "用id查询GoodsSpecItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /goodsSpecItem/findGoodsSpecItem [get]
export const findGoodsSpecItem = (params) => {
  return service({
    url: '/goodsSpecItem/findGoodsSpecItem',
    method: 'get',
    params
  })
}

// @Tags GoodsSpecItem
// @Summary 分页获取GoodsSpecItem列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取GoodsSpecItem列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goodsSpecItem/getGoodsSpecItemList [get]
export const getGoodsSpecItemList = (params) => {
  return service({
    url: '/goodsSpecItem/getGoodsSpecItemList',
    method: 'get',
    params
  })
}
