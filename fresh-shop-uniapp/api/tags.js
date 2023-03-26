import request from "@/utils/request"

export const getTagsListAll = () => {
    return request({
        url: '/tags/getTagsListAll',
        data: {}
    })
}