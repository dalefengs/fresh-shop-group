/*
 * @Author: likfees
 * @Date: 2023-03-24 20:48:21
 * @LastEditors: likfees
 * @LastEditTime: 2023-04-23 22:34:54
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

// refs 为 toast refs
export const getGoodsPageListLoading = (data, refs) => {
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
    }, refs)
}


// 获取商品详情
export const getGoodsInfo = (data, refs) => {
    return request({
        url: `/goods/findGoods`,
        method: 'GET',
        loading: true,
        data
    }, refs)
}

