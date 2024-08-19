import service from '@/utils/request'

// @Tags ShopFavorites
// @Summary 创建ShopFavorites
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShopFavorites true "创建ShopFavorites"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shopFavorites/createShopFavorites [post]
export const createShopFavorites = (data) => {
  return service({
    url: '/shopFavorites/createShopFavorites',
    method: 'post',
    data
  })
}

// @Tags ShopFavorites
// @Summary 删除ShopFavorites
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShopFavorites true "删除ShopFavorites"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shopFavorites/deleteShopFavorites [delete]
export const deleteShopFavorites = (data) => {
  return service({
    url: '/shopFavorites/deleteShopFavorites',
    method: 'delete',
    data
  })
}

// @Tags ShopFavorites
// @Summary 删除ShopFavorites
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ShopFavorites"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shopFavorites/deleteShopFavorites [delete]
export const deleteShopFavoritesByIds = (data) => {
  return service({
    url: '/shopFavorites/deleteShopFavoritesByIds',
    method: 'delete',
    data
  })
}

// @Tags ShopFavorites
// @Summary 更新ShopFavorites
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShopFavorites true "更新ShopFavorites"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /shopFavorites/updateShopFavorites [put]
export const updateShopFavorites = (data) => {
  return service({
    url: '/shopFavorites/updateShopFavorites',
    method: 'put',
    data
  })
}

// @Tags ShopFavorites
// @Summary 用id查询ShopFavorites
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ShopFavorites true "用id查询ShopFavorites"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /shopFavorites/findShopFavorites [get]
export const findShopFavorites = (params) => {
  return service({
    url: '/shopFavorites/findShopFavorites',
    method: 'get',
    params
  })
}

// @Tags ShopFavorites
// @Summary 分页获取ShopFavorites列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ShopFavorites列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shopFavorites/getShopFavoritesList [get]
export const getShopFavoritesList = (params) => {
  return service({
    url: '/shopFavorites/getShopFavoritesList',
    method: 'get',
    params
  })
}
