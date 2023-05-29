<template>
    <pageWrapper>
        <!-- 未登录	-->
        <view class="empty" v-if="!token">
            <u-empty mode="car" text="您还没有登陆哦" icon="http://cdn.uviewui.com/uview/empty/car.png">
                <view>
                    <u-button type="primary" :customStyle="toLoginStyle" text="去登录" @click="showLogin"></u-button>
                </view>
            </u-empty>
        </view>
        <!--  已登录   -->
        <view v-else>
            <!-- 吸顶标签栏 -->
            <u-sticky bgColor="#fff">
                <u-tabs
                        :list="tabsList"
                        lineWidth="50"
                        :activeStyle="{color: '#303133', fontWeight: 'bold',transform: 'scale(1.05)'}"
                        :inactiveStyle="{color: '#606266', transform: 'scale(1)'}"
                        itemStyle="padding:0 15px; height: 40px;"
                        :current="tabsIndex"
                        @change="tabsChange"
                ></u-tabs>
            </u-sticky>
            <!-- 购物车 -->
            <view v-if="tabsIndex === 0">
                <!-- 购物车列表 -->
                <shopCart :list="list" :height="scrollViewHeight" :triggered="triggered"
                          @onRefresh="onRefresh" @delect="delectCart" @update="updateCart" @accounts="accounts" @deleteCart="deleteCartByIndex"/>
            </view>
            <!-- 快速下单 -->
            <view v-if="tabsIndex === 1">

            </view>
        </view>

        <Tabbar :tabsId="2"/>
        <!-- 登录 -->
        <loginPop :show="showLoginDialog" @close="hideLogin" @success="loginSuccess"/>
        <u-toast ref="toast" style="z-index: 9999"></u-toast>
    </pageWrapper>
</template>

<script>
import Tabbar from '@/components/tabbar/tabbar.vue'
import shopCart from '@/components/shopCart/shopCart.vue'
import loginPop from '@/components/loginPop/loginPop.vue'
import {getToken} from '@/store/storage.js'
import {getCartList} from '@/api/cart';
import config from '@/config/config.js'

export default {
    components: {
        Tabbar,
        loginPop,
        shopCart
    },
    data() {
        return {
            token: '',
            tabsIndex: 0,
            list: [], // 购物车列表
            showLoginDialog: false, // 登录
            scrollViewHeight: 0,
            triggered: false,
            toLoginStyle: {
                width: '120px',
                marginTop: '20px',
                borderRadius: '20px',
            },
            tabsList: [{
                name: '购物车',
            },
                // {
                //     name: '快速下单',
                // }
            ]
        };
    },
    onLoad() {
        this.token = getToken()

    },
    onShow() {
        // 登录的情况下获取
        if (this.token) {
            this.getCartListData()
        }
    },
    mounted() {
        // 设置商品列表高度为页面高度
        uni.getSystemInfo({
            success: (res) => {
                const windowHeight = res.windowHeight;
                this.scrollViewHeight = windowHeight - 130;
            },
        });
    },
    methods: {
        // 更新购物车解决微信小程序选中问题
        updateCart(list) {
            this.list = list
        },
        deleteCartByIndex(index) {
            const l = JSON.parse(JSON.stringify(this.list))
            l.splice(index, 1)
            this.list = l
        },
        async getCartListData() {
            const res = await getCartList(this.$refs.toast)
            // 授权过期
            if (res.code === 401) {
                this.token = ''
                return false
            }
            res.data.list.forEach(item => {
                if (item.goods.images && item.goods.images[0].url.slice(0, 4) !== 'http') {
                    item.goods.images[0].url = config.baseUrl + "/" + item.goods.images[0].url
                }
            })
            this.list = res.data.list
            console.log(this.list)
            return true
        },
        // 登录成功
        loginSuccess(u) {
            this.hideLogin();
            this.token = getToken()
            this.$message(this.$refs.toast).success("登录成功")
            this.getCartListData()
        },
        // 显示登录框
        showLogin() {
            this.showLoginDialog = true
        },
        // 隐藏登录框
        hideLogin() {
            this.showLoginDialog = false
        },
        tabsChange(e) {
            this.tabsIndex = e.index
        },
        delectCart(e) {
            console.log('delectCart', e)
        },
        accounts(e) {
            console.log('accounts', e);
        },
        // 刷新
        async onRefresh() {
            this.triggered = true;
            const b = await this.getCartListData()
            if (b) {
                this.$message(this.$refs.toast).success('刷新成功')
            } else {
                this.$message(this.$refs.toast).error('刷新失败')
            }
            this.triggered = false;
        },
    }
}
</script>

<style lang="scss">

page {
  background: #FFFFFF;
}

.empty {
  background: white;
  height: 70vh;
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
