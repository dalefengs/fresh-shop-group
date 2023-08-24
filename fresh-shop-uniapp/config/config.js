/*
 * @Author: likfees
 * @Date: 2023-03-24 14:09:29
 * @LastEditors: likfees
 * @LastEditTime: 2023-04-23 14:06:02
 */

export default {
    // H5 使用代理请求, 正式环境请修改 manifest.json => h5 的请求 url
    // #ifdef H5
    baseUrl: '/api',
    // #endif
    // 其他使用 url
    // #ifndef H5
     // baseUrl: 'http://localhost:48888',
    baseUrl: 'https://qiyundongpin.cn/goapi',
    // #endif
}
