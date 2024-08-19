/*
 * @Author: dalefeng
 * @Date: 2023-04-23 22:34:48
 * @LastEditors: dalefeng
 * @LastEditTime: 2023-04-23 22:34:49
 */
import request from "@/utils/request"

// 获取所有已经选中购物车
export const getUserInfo = (refs) => {
    return request({
        url: `/user/getUserInfo`,
        method: 'GET',
        loading: false,
    },refs)
}

// 获取所有已经选中购物车
export const getUserAuditStatus= (refs) => {
    return request({
        url: `/user/getAuditStatus`,
        method: 'GET',
        loading: false,
    },refs)
}

// 设置自身用户信息
export const setSelfInfo = (data, refs) => {
    return request({
        url: `/user/setSelfInfo`,
        method: 'PUT',
        loading: true,
        data
    },refs)
}

// 设置自身用户信息
export const getUnlimitedQRCodeImg = (refs) => {
    return request({
        url: `/wechat/getUnlimitedQRCode`,
        method: 'GET',
        loading: true,
		responseType: 'arraybuffer'
    },refs)
}
