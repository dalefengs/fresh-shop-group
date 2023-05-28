import request from "@/utils/request"

// 订单支付
export const orderPay = (data, refs) => {
    return request({
        url: `/order/orderPay`,
        method: 'POST',
        loading: true,
        data
    }, refs)
}

// 获取默认列表
export const getOrderList = (data, loading, refs) => {
    return request({
        url: `/order/getUserOrderList`,
        method: 'GET',
        loading: loading,
        toLogin: true,
        data
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

// 用户订单数量统计
export const getOrderStatusCount = (data, refs) => {
    return request({
        url: `/order/findUserOrderStatus`,
        method: 'GET',
        loading: false,
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

