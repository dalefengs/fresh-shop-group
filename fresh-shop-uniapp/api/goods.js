import request from "@/utils/request"

export const getGoodsPageList = (data) => {
    data.status = 1
    if(!data.page) {
        data.page = 1
    }
    if(!data.pageSize) {
        data.pageSize = 5
    }
    return request({
        url: `/goods/getGoodsList`,
        method: 'GET',
        data
    })
}