import service from '@/utils/request'

// @Tags Brand
// @Summary 创建Brand
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Brand true "创建Brand"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /brand/createBrand [post]
export const createBrand = (data) => {
  return service({
    url: '/brand/createBrand',
    method: 'post',
    data
  })
}

// @Tags Brand
// @Summary 删除Brand
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Brand true "删除Brand"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /brand/deleteBrand [delete]
export const deleteBrand = (data) => {
  return service({
    url: '/brand/deleteBrand',
    method: 'delete',
    data
  })
}

// @Tags Brand
// @Summary 删除Brand
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Brand"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /brand/deleteBrand [delete]
export const deleteBrandByIds = (data) => {
  return service({
    url: '/brand/deleteBrandByIds',
    method: 'delete',
    data
  })
}

// @Tags Brand
// @Summary 更新Brand
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Brand true "更新Brand"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /brand/updateBrand [put]
export const updateBrand = (data) => {
  return service({
    url: '/brand/updateBrand',
    method: 'put',
    data
  })
}

// @Tags Brand
// @Summary 用id查询Brand
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Brand true "用id查询Brand"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /brand/findBrand [get]
export const findBrand = (params) => {
  return service({
    url: '/brand/findBrand',
    method: 'get',
    params
  })
}

// @Tags Brand
// @Summary 分页获取Brand列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Brand列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /brand/getBrandList [get]
export const getBrandList = (params) => {
  return service({
    url: '/brand/getBrandList',
    method: 'get',
    params
  })
}

// @Tags Brand
// @Summary 获取所有Brand列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "获取所有Brand列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /brand/getBrandListAll [get]
export const getBrandListAll = () => {
  return service({
    url: '/brand/getBrandListAll',
    method: 'get',
  })
}

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
