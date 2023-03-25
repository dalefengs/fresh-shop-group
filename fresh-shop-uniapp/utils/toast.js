/*
 * @Author: likfees
 * @Date: 2023-03-23 22:41:48
 * @LastEditors: likfees
 * @LastEditTime: 2023-03-24 13:44:48
 */


const success = (msg) => {
    const t = getApp().globalData.toast
    return new Promise((resolve, reject) => {
        if (t) {
            t.show({
                type: 'success',
                message: msg,
                duration: 1500,
                complete: () => {
                    resolve()
                }
            })
        } else {
            uni.showToast({
                title: msg,
                icon: 'success',
                duration: 1500,
                complete: () => {
                    resolve()
                }
            })
        }
    })
}

const error = (msg) => {
    const t = getApp().globalData.toast
    return new Promise((resolve, reject) => {
        if (t) {
            t.show({
                type: 'error',
                message: msg,
                duration: 1500,
                complete: () => {
                    resolve()
                }
            })
        } else {
            uni.showToast({
                title: msg,
                icon: 'error',
                duration: 1500,
                complete: () => {
                    resolve()
                }
            })
        }
    })
}

const warning = (msg) => {
    const t = getApp().globalData.toast
    return new Promise((resolve, reject) => {
        if (t) {
            t.show({
                type: 'warning',
                message: msg,
                duration: 1500,
                complete: () => {
                    resolve()
                }
            })
        } else {
            uni.showToast({
                title: msg,
                icon: 'none',
                duration: 1500,
                complete: () => {
                    resolve()
                }
            })
        }
    })
}

const info = (msg) => {
    const t = getApp().globalData.toast
    return new Promise((resolve, reject) => {
        if (t) {
            t.show({
                type: 'primary',
                message: msg,
                duration: 1500,
                complete: () => {
                    resolve()
                }
            })
        } else {
            uni.showToast({
                title: msg,
                icon: 'none',
                duration: 1500,
                complete: () => {
                    resolve()
                }
            })
        }
    })
}

const none = (msg) => {
    const t = getApp().globalData.toast
    return new Promise((resolve, reject) => {
        if (t) {
            t.show({
                type: 'default',
                message: msg,
                duration: 1500,
                complete: () => {
                    resolve()
                }
            })
        } else {
            uni.showToast({
                title: msg,
                icon: 'none',
                duration: 1500,
                complete: () => {
                    resolve()
                }
            })
        }
    })
}


const loading = (msg) => {
    const t = getApp().globalData.toast
    return new Promise((resolve, reject) => {
        if (t) {
            t.show({
                type: 'loading',
                message: msg,
                duration: 200000,
                complete: () => {
                    resolve()
                }
            })
        } else {
            uni.showLoading({
                title: msg
            });
        }
    })
}


const hide = () => {
    const t = getApp().globalData.toast
    if (t) {
        t.hide()
    } else {
        uni.hideLoading();
    }
}

const show = (options) => {
    const t = getApp().globalData.toast
    return new Promise((resolve, reject) => {
        t.show({
            ...options,
            complete: () => {
                resolve()
            }
        })
    })
    
}


export default {
    success,
    error,
    info,
    none,
    warning,
    loading,
    hide,
    show,
}