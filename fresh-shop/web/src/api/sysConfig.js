import service from '@/utils/request'

// @Tags SysConfig
// @Summary 创建SysConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysConfig true "创建SysConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysConfig/createSysConfig [post]
export const createSysConfig = (data) => {
  return service({
    url: '/sysConfig/createSysConfig',
    method: 'post',
    data
  })
}

// @Tags SysConfig
// @Summary 删除SysConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysConfig true "删除SysConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysConfig/deleteSysConfig [delete]
export const deleteSysConfig = (data) => {
  return service({
    url: '/sysConfig/deleteSysConfig',
    method: 'delete',
    data
  })
}

// @Tags SysConfig
// @Summary 删除SysConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysConfig/deleteSysConfig [delete]
export const deleteSysConfigByIds = (data) => {
  return service({
    url: '/sysConfig/deleteSysConfigByIds',
    method: 'delete',
    data
  })
}

// @Tags SysConfig
// @Summary 更新SysConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysConfig true "更新SysConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysConfig/updateSysConfig [put]
export const updateSysConfig = (data) => {
  return service({
    url: '/sysConfig/updateSysConfig',
    method: 'put',
    data
  })
}

// @Tags SysConfig
// @Summary 用id查询SysConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SysConfig true "用id查询SysConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysConfig/findSysConfig [get]
export const findSysConfig = (params) => {
  return service({
    url: '/sysConfig/findSysConfig',
    method: 'get',
    params
  })
}

// @Tags SysConfig
// @Summary 分页获取SysConfig列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取SysConfig列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysConfig/getSysConfigList [get]
export const getSysConfigList = (params) => {
  return service({
    url: '/sysConfig/getSysConfigList',
    method: 'get',
    params
  })
}
