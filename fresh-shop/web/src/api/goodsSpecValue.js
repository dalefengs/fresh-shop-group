import service from '@/utils/request'

// @Tags GoodsSpecValue
// @Summary 创建GoodsSpecValue
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoodsSpecValue true "创建GoodsSpecValue"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goodsSpecValue/createGoodsSpecValue [post]
export const createGoodsSpecValue = (data) => {
  return service({
    url: '/goodsSpecValue/createGoodsSpecValue',
    method: 'post',
    data
  })
}

// @Tags GoodsSpecValue
// @Summary 删除GoodsSpecValue
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoodsSpecValue true "删除GoodsSpecValue"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goodsSpecValue/deleteGoodsSpecValue [delete]
export const deleteGoodsSpecValue = (data) => {
  return service({
    url: '/goodsSpecValue/deleteGoodsSpecValue',
    method: 'delete',
    data
  })
}

// @Tags GoodsSpecValue
// @Summary 删除GoodsSpecValue
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除GoodsSpecValue"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goodsSpecValue/deleteGoodsSpecValue [delete]
export const deleteGoodsSpecValueByIds = (data) => {
  return service({
    url: '/goodsSpecValue/deleteGoodsSpecValueByIds',
    method: 'delete',
    data
  })
}

// @Tags GoodsSpecValue
// @Summary 更新GoodsSpecValue
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoodsSpecValue true "更新GoodsSpecValue"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /goodsSpecValue/updateGoodsSpecValue [put]
export const updateGoodsSpecValue = (data) => {
  return service({
    url: '/goodsSpecValue/updateGoodsSpecValue',
    method: 'put',
    data
  })
}

// @Tags GoodsSpecValue
// @Summary 用id查询GoodsSpecValue
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.GoodsSpecValue true "用id查询GoodsSpecValue"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /goodsSpecValue/findGoodsSpecValue [get]
export const findGoodsSpecValue = (params) => {
  return service({
    url: '/goodsSpecValue/findGoodsSpecValue',
    method: 'get',
    params
  })
}

// @Tags GoodsSpecValue
// @Summary 分页获取GoodsSpecValue列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取GoodsSpecValue列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goodsSpecValue/getGoodsSpecValueList [get]
export const getGoodsSpecValueList = (params) => {
  return service({
    url: '/goodsSpecValue/getGoodsSpecValueList',
    method: 'get',
    params
  })
}
