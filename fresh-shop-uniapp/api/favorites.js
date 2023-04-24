/*
 * @Author: likfees
 * @Date: 2023-04-23 17:53:00
 * @LastEditors: likfees
 * @LastEditTime: 2023-04-23 18:03:49
 */

import request from "@/utils/request"

// 收藏
export const favorites = (data) => {
    return request({
        url: `/favorites/favorites`,
        method: 'POST',
        data
    })
}


// 收藏列表
export const getFavoritesListPage = (data) => {
    if (!data.page) {
        data.page = 1
    }
    if (!data.pageSize) {
        data.pageSize = 5
    }
    return request({
        url: `/favorites/getFavoritesList`,
        method: 'GET',
        loading: true,
        data
    })
}