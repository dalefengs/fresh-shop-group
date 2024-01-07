/*
 * @Author: dalefeng
 * @Date: 2023-04-14 20:40:42
 * @LastEditors: dalefeng
 * @LastEditTime: 2023-04-18 17:57:36
 */
import request from "@/utils/request"

export const getWeChatOpenIdByCode = (data) => {
  return request({
    url: '/wechat/code2Session',
    data,
  })
}

export const wxLogin = (data) => {
  return request({
    url: '/base/loginWx',
    method: 'POST',
    loading: true,
    data,
  })
}
