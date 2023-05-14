import service from '@/utils/request'

// @Tags Favorites
// @Summary 创建Favorites
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Favorites true "创建Favorites"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /favorites/createFavorites [post]
export const createFavorites = (data) => {
  return service({
    url: '/favorites/createFavorites',
    method: 'post',
    data
  })
}

// @Tags Favorites
// @Summary 删除Favorites
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Favorites true "删除Favorites"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /favorites/deleteFavorites [delete]
export const deleteFavorites = (data) => {
  return service({
    url: '/favorites/deleteFavorites',
    method: 'delete',
    data
  })
}

// @Tags Favorites
// @Summary 删除Favorites
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Favorites"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /favorites/deleteFavorites [delete]
export const deleteFavoritesByIds = (data) => {
  return service({
    url: '/favorites/deleteFavoritesByIds',
    method: 'delete',
    data
  })
}

// @Tags Favorites
// @Summary 更新Favorites
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Favorites true "更新Favorites"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /favorites/updateFavorites [put]
export const updateFavorites = (data) => {
  return service({
    url: '/favorites/updateFavorites',
    method: 'put',
    data
  })
}

// @Tags Favorites
// @Summary 用id查询Favorites
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Favorites true "用id查询Favorites"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /favorites/findFavorites [get]
export const findFavorites = (params) => {
  return service({
    url: '/favorites/findFavorites',
    method: 'get',
    params
  })
}

// @Tags Favorites
// @Summary 分页获取Favorites列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Favorites列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /favorites/getFavoritesList [get]
export const getFavoritesList = (params) => {
  return service({
    url: '/favorites/getFavoritesList',
    method: 'get',
    params
  })
}
