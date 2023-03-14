import service from '@/utils/request'

// @Tags SysRecharge
// @Summary 创建SysRecharge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysRecharge true "创建SysRecharge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysRecharge/createSysRecharge [post]
export const createSysRecharge = (data) => {
  return service({
    url: '/sysRecharge/createSysRecharge',
    method: 'post',
    data
  })
}

// @Tags SysRecharge
// @Summary 删除SysRecharge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysRecharge true "删除SysRecharge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysRecharge/deleteSysRecharge [delete]
export const deleteSysRecharge = (data) => {
  return service({
    url: '/sysRecharge/deleteSysRecharge',
    method: 'delete',
    data
  })
}

// @Tags SysRecharge
// @Summary 删除SysRecharge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysRecharge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysRecharge/deleteSysRecharge [delete]
export const deleteSysRechargeByIds = (data) => {
  return service({
    url: '/sysRecharge/deleteSysRechargeByIds',
    method: 'delete',
    data
  })
}

// @Tags SysRecharge
// @Summary 更新SysRecharge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysRecharge true "更新SysRecharge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysRecharge/updateSysRecharge [put]
export const updateSysRecharge = (data) => {
  return service({
    url: '/sysRecharge/updateSysRecharge',
    method: 'put',
    data
  })
}

// @Tags SysRecharge
// @Summary 用id查询SysRecharge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SysRecharge true "用id查询SysRecharge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysRecharge/findSysRecharge [get]
export const findSysRecharge = (params) => {
  return service({
    url: '/sysRecharge/findSysRecharge',
    method: 'get',
    params
  })
}

// @Tags SysRecharge
// @Summary 分页获取SysRecharge列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取SysRecharge列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysRecharge/getSysRechargeList [get]
export const getSysRechargeList = (params) => {
  return service({
    url: '/sysRecharge/getSysRechargeList',
    method: 'get',
    params
  })
}
