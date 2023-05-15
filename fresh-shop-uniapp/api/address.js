import request from "@/utils/request"

// 获取默认列表
export const getAddressList = (refs) => {
    return request({
        url: `/userAddress/getUserAddressList`,
        method: 'GET',
        loading: true,
        toLogin: true,
    }, refs)
}

// 获取地址
export const getAddressInfo = (data, refs) => {
    return request({
        url: `/userAddress/findUserAddress`,
        method: 'GET',
        loading: true,
        data
    }, refs)
}

// 获取默认地址
export const getDefaultAddressInfo = (refs) => {
    return request({
        url: `/userAddress/findUserAddress`,
        method: 'GET',
        loading: false,
        data: {
            isDefault: 1
        }
    }, refs)
}

// 创建地址
export const createAddress = (data, refs) => {
    return request({
        url: `/userAddress/createUserAddress`,
        method: 'POST',
        loading: true,
        data
    }, refs)
}

// 更新地址
export const updateAddress = (data, refs) => {
    return request({
        url: `/userAddress/updateUserAddress`,
        method: 'PUT',
        loading: true,
        data
    }, refs)
}

// 更新地址
export const deleteAddress = (data, refs) => {
    return request({
        url: `/userAddress/deleteUserAddress`,
        method: 'DELETE',
        loading: true,
        data
    }, refs)
}
