<!--
 * @Author: dalefeng
 * @Date: 2023-04-23 12:02:57
 * @LastEditors: dalefeng
 * @LastEditTime: 2023-04-23 12:43:56
-->
<template>
    <view v-if="show">
        <view class="box" @click="showLogin">
            <view class="king-ml-20">
                登录查看专属价格及下单订货
            </view>
            <view class="king-mr-10">
                <u-button type="primary" :customStyle="buttonStyle" text="立即登录"></u-button>
            </view>
        </view>
        <loginPop :show="showLoginDialog" @close="hideLogin" @success="loginSuccess"> </loginPop>
        <u-toast style="z-index:9998;" ref="toast"></u-toast>
    </view>
</template>

<script>
import loginPop from '@/components/loginPop/loginPop.vue'
export default {
    components: {
        loginPop
    },
    props: {
        show: {
            type: Boolean,
            default: false
        }
    },
    data() {
        return {
            showLoginDialog: false,
            buttonStyle: {
                "height": "30px",
                "borderRadius": "14px"
            }
        }
    },
    methods: {
        // 显示登录框
        showLogin() {
            this.showLoginDialog = true
        },
        // 隐藏登录框
        hideLogin() {
            this.showLoginDialog = false
        },
        // 登录成功
        loginSuccess(u) {
            this.hideLogin();
            this.$emit("success")
            this.$message(this.$refs.toast).success("登录成功")
        },
    }
}
</script>

<style lang="scss">
.box {
    z-index: 9999;
    position: fixed;
    width: 100%;
    height: 50px;
    bottom: 50px;
    background-color: rgba(0, 0, 0, 0.7);
    display: flex;
    justify-content: space-between;
    align-items: center;
    color: white;
}
</style>
