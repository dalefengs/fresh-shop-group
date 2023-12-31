/*
 * @Author: likfees
 * @Date: 2023-04-23 22:34:48
 * @LastEditors: likfees
 * @LastEditTime: 2023-04-23 22:34:49
 */
import request from "@/utils/request"

// 操作购物车
export const addCart = (data) => {
    return request({
        url: `/cart/createCart`,
        method: 'POST',
        data
    })
}

// 获取所有购物车
export const getCartList = (refs) => {
    return request({
        url: `/cart/getCartList`,
        method: 'GET',
        loading: true,
    },refs)
}

// 获取所有已经选中购物车
export const getCheckedCartList = (refs) => {
    const data = {
        checked: 1
    }
    return request({
        url: `/cart/getCartList`,
        method: 'GET',
        loading: true,
        data
    },refs)
}

// 更新购物车
export const updateCart = (data, refs) => {
    return request({
        url: `/cart/updateCart`,
        method: 'PUT',
        loading: false,
        data
    },refs)
}

// 全选购物车
export const selectAllCart = (refs) => {
    return request({
        url: `/cart/selectAllChecked`,
        method: 'POST',
        loading: false,
    },refs)
}

// 全选购物车
export const selectGoodsSingeChecked = (data, refs) => {
    return request({
        url: `/cart/selectGoodsSingeChecked`,
        method: 'POST',
        loading: false,
        data
    },refs)
}

// 取消全选购物车
export const clearSelectAllCart = (refs) => {
    return request({
        url: `/cart/clearAllChecked`,
        method: 'POST',
        loading: false,
    },refs)
}

// 删除购物车
export const deleteCartByIds = (data, refs) => {
    return request({
        url: `/cart/deleteCartByIds`,
        method: 'DELETE',
        loading: false,
        data,
    },refs)
}
