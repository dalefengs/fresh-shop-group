import service from '@/utils/request'

// @Tags GoodsSpec
// @Summary 创建GoodsSpec
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoodsSpec true "创建GoodsSpec"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goodsSpec/createGoodsSpec [post]
export const createGoodsSpec = (data) => {
  return service({
    url: '/goodsSpec/createGoodsSpec',
    method: 'post',
    data
  })
}

// @Tags GoodsSpec
// @Summary 删除GoodsSpec
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoodsSpec true "删除GoodsSpec"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goodsSpec/deleteGoodsSpec [delete]
export const deleteGoodsSpec = (data) => {
  return service({
    url: '/goodsSpec/deleteGoodsSpec',
    method: 'delete',
    data
  })
}

// @Tags GoodsSpec
// @Summary 删除GoodsSpec
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除GoodsSpec"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goodsSpec/deleteGoodsSpec [delete]
export const deleteGoodsSpecByIds = (data) => {
  return service({
    url: '/goodsSpec/deleteGoodsSpecByIds',
    method: 'delete',
    data
  })
}

// @Tags GoodsSpec
// @Summary 更新GoodsSpec
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoodsSpec true "更新GoodsSpec"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /goodsSpec/updateGoodsSpec [put]
export const updateGoodsSpec = (data) => {
  return service({
    url: '/goodsSpec/updateGoodsSpec',
    method: 'put',
    data
  })
}

// @Tags GoodsSpec
// @Summary 用id查询GoodsSpec
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.GoodsSpec true "用id查询GoodsSpec"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /goodsSpec/findGoodsSpec [get]
export const findGoodsSpec = (params) => {
  return service({
    url: '/goodsSpec/findGoodsSpec',
    method: 'get',
    params
  })
}

// @Tags GoodsSpec
// @Summary 分页获取GoodsSpec列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取GoodsSpec列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goodsSpec/getGoodsSpecList [get]
export const getGoodsSpecList = (params) => {
  return service({
    url: '/goodsSpec/getGoodsSpecList',
    method: 'get',
    params
  })
}
