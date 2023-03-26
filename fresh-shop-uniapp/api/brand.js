/*
 * @Author: likfees
 * @Date: 2023-03-25 22:49:51
 * @LastEditors: likfees
 * @LastEditTime: 2023-03-25 22:51:15
 */
import request from "@/utils/request"

export const getBrandListAll = (data) => {
    return request({
        url: '/brand/getBrandListAll',
        method: 'GET',
        data
    })
}
