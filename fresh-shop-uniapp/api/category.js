/*
 * @Author: likfees
 * @Date: 2023-03-24 14:39:31
 * @LastEditors: likfees
 * @LastEditTime: 2023-03-24 14:41:07
 */
import request from "@/utils/request"

export const getCategoryListAll = (data) => {
    return request({
    url: '/category/getCategoryListAll',
    method: 'GET',
    data
  })
}

export const getHomeCategoryList = () => {
    return request({
        url: '/category/getCategoryList',
        method: 'GET',
        data: {
            page: 1,
            pageSize: 8,
            isFirst: 1
        }
    })
}