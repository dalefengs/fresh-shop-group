const TOKEN = "token"
const EXPIRES = "expires"
const USER = "user"
const ROLE = "role"
const OPENID = "openid"
const FirstEntry = "firstEntry"

const SettlmentInfo = "settlmentInfo"


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

// 设置用户角色信息
export function setRole(role) {
    uni.setStorageSync(ROLE, role)
}

export function getRole() {
    return uni.getStorageSync(ROLE)
}

// 设置 openid
export function setOpenId(openId) {
    uni.setStorageSync(OPENID, openId)
}

export function getOpenId() {
    return uni.getStorageSync(OPENID)
}


export function getFirstEntry() {
    return uni.getStorageSync(FirstEntry)
}
// 设置 openid
export function setFirstEntry(bool) {
    uni.setStorageSync(FirstEntry, bool)
}


// 设置结算状态
export function getSettlmentInfo() {
    return uni.getStorageSync(SettlmentInfo)
}
export function setSettlmentInfo(data) {
    uni.setStorageSync(SettlmentInfo, data)
}
