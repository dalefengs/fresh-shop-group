/*
 * @Author: likfees
 * @Date: 2023-04-23 22:34:48
 * @LastEditors: likfees
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
