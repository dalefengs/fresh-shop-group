<!--
 * @Author: likfees
 * @Date: 2023-03-23 22:16:00
 * @LastEditors: likfees
 * @LastEditTime: 2023-04-21 17:26:13
-->
<template>
    <u-popup :show="show" @open="open" @close="close" :closeable="closeable">
        <view class="box">
            <view class="title">用户登录</view>
            <view class="btn">
                <u-button :customStyle="otherBtnStyle" open-type="getPhoneNumber"
                    @getphonenumber="getPhoneNumber">微信用户一键登录</u-button>
            </view>
            <view class="other-btn" @click="otherBtnClick">其他号码登录/注册</view>
        </view>
    </u-popup>
</template>

<script>
import toast from '@/utils/toast.js'
import { setToken, setExpires, setUser } from '@/store/storage.js'
import { getWeChatOpenIdByCode, wxLogin } from '@/api/login.js'
export default {
    name: "loginPop",
    data() {
        return {
            openid: '',
            sessionKey: '',
            otherBtnStyle: {
                width: "80%",
                height: "50px",
                borderRadius: "30px",
                background: "#2979ff",
                color: "#fff",
            },
        };
    },
    props: {
        show: {
            type: Boolean,
            default: false,
        },
        closeable: {
            type: Boolean,
            default: true,
        }
    },
    mounted() {
        getApp().globalData.toast = this.$refs.uToast
        uni.getSystemInfo({
            success: (res) => {
                this.windowHeight = res.windowHeight - this.subHeight;
            },
        });
    },
    methods: {

        //用户授权登陆允许后，返回encryptedData, iv参数
        getPhoneNumber(e) {
            //拿到参数后进一步去解密....
            // 授权通过后轮询等待获取sessionKey响应成功
            if (['getPhoneNumber:ok'].includes(e?.detail?.errMsg)) {
                console.log('getPhoneNumber e', e);
                // 发起登录请求
                this.login(e.detail.encryptedData, e.detail.iv)
            } else {
                toast.error("获取手机号失败")
            }
        },
        login(encryptedData, iv) {
            if (!this.sessionKey) {
                toast.warn("请稍后在试试")
                return
            }
            wxLogin({
                encryptedData,
                iv,
                sessionKey: this.sessionKey,
                openid: this.openId
            }).then(res => {
                console.log('wxLogin res', res);
                setToken(res.data.token)
                setExpires(res.data.expiresAt)
                setUser(res.data.user)
                this.$emit("success", res.data.user)
            })
        },
        // 初始化调用wx.login获取登陆凭证code（用户无感知）
        loginGetJsCode() {
            return new Promise((resolve, reject) => {
                uni.login({
                    success(res) {
                        resolve(res);
                    },
                    fail(err) {
                        reject(err);
                    }
                });
            });
        },
        async open() {
            //this.$emit("open");
            let res = await this.loginGetJsCode()
            console.log("login res", res);
            if (res.code && ['login:ok'].includes(res.errMsg)) {
                //后端通过code调用微信API返回openid/unionid/session_key参数， 存储起来,前端无法直接调用微信相关API
                let openIdRes = await getWeChatOpenIdByCode({ js_code: res.code })
                console.log('openIdRes', openIdRes);
                if (openIdRes.data.errcode == 0) {
                    this.openId = openIdRes.data.openid
                    this.sessionKey = openIdRes.data.session_key
                }
            } else {
                //异常处理，再次发起请求或者抛出异常
                toast.error("获取 sessionKey 失败")
            }
        },
        close() {
            this.$emit("close");
        },
        otherBtnClick() {
            toast.info("开发中, 请使用微信登录")
        }
    }
}
</script>

<style lang="scss">
.box {
    height: 30vh;
}

.title {
    margin: 12px 15px;
}

.btn {
    height: 50%;
    display: flex;
    justify-content: center;
    align-items: center;
}

.other-btn {
    width: 100%;
    text-align: center;
}
</style>