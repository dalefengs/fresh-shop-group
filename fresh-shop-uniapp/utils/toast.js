/*
 * @Author: dalefeng
 * @Date: 2023-03-23 22:41:48
 * @LastEditors: dalefeng
 * @LastEditTime: 2023-04-23 14:37:59
 */

var toast = null
const message = (ref) => {
    toast = ref
    return {
        success,
        error,
        info,
        none,
        warning,
        loading,
        hide,
        show,
    }
}

const success = (msg) => {
    return new Promise((resolve, reject) => {
        if (toast && Object.keys(toast).length > 0) {
            toast.show({
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
    return new Promise((resolve, reject) => {
        if (toast && Object.keys(toast).length > 0) {
            toast.show({
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
    return new Promise((resolve, reject) => {
        if (toast && Object.keys(toast).length > 0) {
            toast.show({
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
    return new Promise((resolve, reject) => {
        if (toast && Object.keys(toast).length > 0) {
            toast.show({
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
    return new Promise((resolve, reject) => {
        if (toast && Object.keys(toast).length > 0) {
            toast.show({
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
    return new Promise((resolve, reject) => {
        if (toast && Object.keys(toast).length > 0) {
            toast.show({
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
    try {
        uni.hideLoading()
        toast.hide()
    } catch (e) { }

}

const show = (options) => {
    return new Promise((resolve, reject) => {
        toast.show({
            ...options,
            complete: () => {
                resolve()
            }
        })
    })

}


export default {
    message,
}
