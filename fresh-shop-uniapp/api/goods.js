import request from "@/utils/request"

export const getGoodsPageList = (data) => {
    data.status = 1
    return request({
        url: '/goods/getGoodsList',
        method: 'GET',
        data
    })
}