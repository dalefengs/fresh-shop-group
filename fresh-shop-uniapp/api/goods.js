/*
 * @Author: likfees
 * @Date: 2023-03-24 20:48:21
 * @LastEditors: likfees
 * @LastEditTime: 2023-04-23 17:52:47
 */
import request from "@/utils/request"

export const getGoodsPageList = (data) => {
    data.status = 1
    if (!data.page) {
        data.page = 1
    }
    if (!data.pageSize) {
        data.pageSize = 5
    }
    return request({
        url: `/goods/getGoodsList`,
        method: 'GET',
        data
    })
}

export const getGoodsPageListLoading = (data) => {
    data.status = 1
    if (!data.page) {
        data.page = 1
    }
    if (!data.pageSize) {
        data.pageSize = 5
    }
    return request({
        url: `/goods/getGoodsList`,
        method: 'GET',
        loading: true,
        data
    })
}


// 获取商品详情
export const getGoodsInfo = (data) => {
    return request({
        url: `/goods/findGoods`,
        method: 'GET',
        data
    })
}

