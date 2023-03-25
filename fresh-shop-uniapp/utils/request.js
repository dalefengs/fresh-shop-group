// uniapp request 请求
import toast from '@/utils/toast.js'
import config from '@/config/config.js'


/**
 * 封装 Uniapp 的请求工具类
 * @param {Object} options 请求配置参数
 * @param {string} options.url 请求的地址
 * @param {string} [options.method='GET'] 请求的方法
 * @param {Object} [options.data] 请求的数据
 * @param {Object} [options.header] 请求的头部信息
 * @param {boolean} [options.showLoading=false] 是否显示加载动画
 * @param {boolean} [options.showError=true] 是否显示错误提示
 * @param {boolean} [options.autoLogin=false] 是否自动跳转到登录页
 * @returns {Promise} Promise 对象
 */
const request = (options) => {
	options.method = options.method || 'GET'
	options.loading = options.loading === undefined ? false : options.loading // 默认不显示 loading
	//options.showError = options.showError === undefined ? true : options.showError
	options.toLogin = options.toLogin === undefined ? false : options.toLogin // 默认不跳转到登录页

	if (options.loading) {
		toast.loading('加载中...')
	}

	// 添加请求头
	options.header = Object.assign({
		'content-type': 'application/json',
		'x-token': uni.getStorageSync('token')
	}, options.header)

	options.url = config.baseUrl + options.url

	// 发起请求
	return new Promise((resolve, reject) => {
		uni.request({
			...options,
			success: (res) => {
				// 隐藏加载动画
				if (options.loading) {
					toast.hide()
				}
				// 如果响应头部中返回了新的 token，则重新写入本地存储中
				if (res.header['new-token']) {
					uni.setStorageSync('token', res.header['new-token'])
				}

				// 根据响应状态码处理响应结果
				switch (res.statusCode) {
					case 200:
						resolve(res.data)
						break
					case 401:
						if (options.toLogin) {
							// 如果需要自动跳转到登录页，则跳转到登录页
							uni.reLaunch({
								url: '/pages/login/login'
							})
						} else {
							reject(res)
						}
						break
					case 500:
						toast.error('服务器异常，请稍后重试！')
						reject(res)
						break
					default:
						toast.error('请求失败，请稍后重试！')
						reject(res)
				}

				resolve(res)
			},
			fail: (err) => {
				// 隐藏加载动画
				if (options.loading) {
					toast.hide()
				}
				console.log('uni.request fail', err)
				toast.error('服务器异常，请稍后重试！')
				reject(err)
			},
		})
	})
}

export default request
