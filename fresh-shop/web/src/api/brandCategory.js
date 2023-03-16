import service from '@/utils/request'

// @Tags BrandCategory
// @Summary 创建BrandCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BrandCategory true "创建BrandCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /brandCategory/createBrandCategory [post]
export const createBrandCategory = (data) => {
  return service({
    url: '/brandCategory/createBrandCategory',
    method: 'post',
    data
  })
}

// @Tags BrandCategory
// @Summary 删除BrandCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BrandCategory true "删除BrandCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /brandCategory/deleteBrandCategory [delete]
export const deleteBrandCategory = (data) => {
  return service({
    url: '/brandCategory/deleteBrandCategory',
    method: 'delete',
    data
  })
}

// @Tags BrandCategory
// @Summary 删除BrandCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BrandCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /brandCategory/deleteBrandCategory [delete]
export const deleteBrandCategoryByIds = (data) => {
  return service({
    url: '/brandCategory/deleteBrandCategoryByIds',
    method: 'delete',
    data
  })
}

// @Tags BrandCategory
// @Summary 更新BrandCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BrandCategory true "更新BrandCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /brandCategory/updateBrandCategory [put]
export const updateBrandCategory = (data) => {
  return service({
    url: '/brandCategory/updateBrandCategory',
    method: 'put',
    data
  })
}

// @Tags BrandCategory
// @Summary 用id查询BrandCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BrandCategory true "用id查询BrandCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /brandCategory/findBrandCategory [get]
export const findBrandCategory = (params) => {
  return service({
    url: '/brandCategory/findBrandCategory',
    method: 'get',
    params
  })
}

// @Tags BrandCategory
// @Summary 分页获取BrandCategory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BrandCategory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /brandCategory/getBrandCategoryList [get]
export const getBrandCategoryList = (params) => {
  return service({
    url: '/brandCategory/getBrandCategoryList',
    method: 'get',
    params
  })
}
