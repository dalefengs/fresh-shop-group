import service from '@/utils/request'

// @Tags Tags
// @Summary 创建Tags
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Tags true "创建Tags"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tags/createTags [post]
export const createTags = (data) => {
  return service({
    url: '/tags/createTags',
    method: 'post',
    data
  })
}

// @Tags Tags
// @Summary 删除Tags
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Tags true "删除Tags"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tags/deleteTags [delete]
export const deleteTags = (data) => {
  return service({
    url: '/tags/deleteTags',
    method: 'delete',
    data
  })
}

// @Tags Tags
// @Summary 删除Tags
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Tags"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tags/deleteTags [delete]
export const deleteTagsByIds = (data) => {
  return service({
    url: '/tags/deleteTagsByIds',
    method: 'delete',
    data
  })
}

// @Tags Tags
// @Summary 更新Tags
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Tags true "更新Tags"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tags/updateTags [put]
export const updateTags = (data) => {
  return service({
    url: '/tags/updateTags',
    method: 'put',
    data
  })
}

// @Tags Tags
// @Summary 用id查询Tags
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Tags true "用id查询Tags"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tags/findTags [get]
export const findTags = (params) => {
  return service({
    url: '/tags/findTags',
    method: 'get',
    params
  })
}

// @Tags Tags
// @Summary 分页获取Tags列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Tags列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tags/getTagsList [get]
export const getTagsList = (params) => {
  return service({
    url: '/tags/getTagsList',
    method: 'get',
    params
  })
}
