import request from "@/utils/request"

// 订单支付
export const orderPay = (refs) => {
    return request({
        url: `/order/orderPay`,
        method: 'POST',
        loading: true,
        toLogin: true,
    }, refs)
}

// 获取默认列表
export const getOrderList = (refs) => {
    return request({
        url: `/order/getOrderList`,
        method: 'GET',
        loading: true,
        toLogin: true,
    }, refs)
}

// 获取订单
export const getOrderInfo = (data, refs) => {
    return request({
        url: `/order/findOrder`,
        method: 'GET',
        loading: true,
        data
    }, refs)
}

// 获取订单状态
export const getOrderStatus = (orderId, refs) => {
    return request({
        url: `/order/orderStatus`,
        method: 'GET',
        loading: false,
        data: {
            ID: orderId
        }
    }, refs)
}

// 创建订单
export const createOrder = (data, refs) => {
    return request({
        url: `/order/createOrder`,
        method: 'POST',
        loading: true,
        data
    }, refs)
}

// 更新订单
export const updateOrder = (data, refs) => {
    return request({
        url: `/order/updateOrder`,
        method: 'PUT',
        loading: true,
        data
    }, refs)
}

// 更新订单
export const deleteOrder = (data, refs) => {
    return request({
        url: `/order/deleteOrder`,
        method: 'DELETE',
        loading: true,
        data
    }, refs)
}

// 取消订单
export const cancelOrder = (data, refs) => {
    return request({
        url: `/order/cancelOrder`,
        method: 'POST',
        loading: true,
        data
    }, refs)
}


// 确认收货
export const confirmOrder = (data, refs) => {
    return request({
        url: `/order/confirmOrder`,
        method: 'POST',
        loading: true,
        data
    }, refs)
}

