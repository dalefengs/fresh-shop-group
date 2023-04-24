/*
 * @Author: likfees
 * @Date: 2023-04-23 22:34:48
 * @LastEditors: likfees
 * @LastEditTime: 2023-04-23 22:34:49
 */
import request from "@/utils/request"

// 收藏
export const addCart = (data) => {
    return request({
        url: `/cart/createCart`,
        method: 'POST',
        data
    })
}
