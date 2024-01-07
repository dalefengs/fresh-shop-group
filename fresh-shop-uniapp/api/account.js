/*
 * @Author: dalefeng
 * @Date: 2023-04-23 22:34:48
 * @LastEditors: dalefeng
 * @LastEditTime: 2023-04-23 22:34:49
 */
import request from "@/utils/request"

// 获取用户账户信息
export const getAccountInfo = (groupId, refs) => {
    const data = {
        ID: groupId,
    }
    return request({
        url: `/account/findUserAccount`,
        method: 'GET',
        loading: true,
        data
    },refs)
}

