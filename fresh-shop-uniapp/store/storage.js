const TOKEN = "token"
const EXPIRES = "expires"
const USER = "user"


// 设置 token
export function setToken(token) {
    uni.setStorageSync(TOKEN, token)
}

export function getToken() {
    return uni.getStorageSync(TOKEN)
}


// 设置过期时间
export function setExpires(expires) {
    uni.setStorageSync(EXPIRES, expires)
}

export function getExpires() {
   return uni.getStorageSync(EXPIRES)
}


// 设置用户信息
export function setUser(user) {
    uni.setStorageSync(USER, user)
}

export function getUser() {
    return uni.getStorageSync(USER)
}