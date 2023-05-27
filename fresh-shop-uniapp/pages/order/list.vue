<template>
    <pageWrapper>
        <view class="tabs">
            <u-tabs :list="statusList" :current="tabsCurrentIndex" @click="clickTabs" :scrollable="false"></u-tabs>
        </view>
        <swiper class="swiper" :current-item-id="currentStatus" @change="swiperChange">
            <!-- 全部订单 -->
            <swiper-item item-id="null">
                <orderList status="null" v-if="currentStatus === 'null' || swiperLazyShow['null'] "></orderList>
            </swiper-item>
            <!-- 待付款订单 -->
            <swiper-item item-id="0">
                <orderList status="0" v-if="currentStatus === '0' || swiperLazyShow['0'] "></orderList>
            </swiper-item>
            <!-- 备货中 -->
            <swiper-item item-id="1">
                <orderList status="1" v-if="currentStatus === '1' || swiperLazyShow['1'] "></orderList>
            </swiper-item>
            <!-- 待收货 -->
            <swiper-item item-id="2">
                <orderList status="2" v-if="currentStatus === '2' || swiperLazyShow['2'] "></orderList>
            </swiper-item>
            <!-- 售后订单 -->
            <swiper-item item-id="10">
                <orderList status="10" v-if="currentStatus === '10' || swiperLazyShow['10'] "></orderList>
            </swiper-item>
        </swiper>
        <u-toast ref="toast" style="z-index: 9999"></u-toast>
    </pageWrapper>
</template>

<script>
import config from '@/config/config.js'
import {getOrderList, confirmOrder, cancelOrder, orderPay, getOrderStatus} from '@/api/order'
import {getToken} from '@/store/storage.js'
import orderList from "@/components/orderList/orderList";

export default {
    components: {
        orderList
    },
    data() {
        return {
            token: '',
            list: [],
            statusList: [{
                name: '全部订单',
                status: 'null',
            }, {
                name: '未付款',
                status: '0',
            }, {
                name: '备货中',
                status: '1',
            }, {
                name: '待收货',
                status: '2',
            }, {
                name: '待售后',
                status: '10',
            }],
            tabsCurrentIndex: 0,
            currentStatus: '1', // 当前订单状态 null:全部 0:未支付 1:备货中 2:已发货 3:已完成
            swiperLazyShow: { // 用来控制swiper的懒加载，
                'null': false,
                '0': false,
                '1': false,
                '2': false,
                '10': false
            }
        }
    },
    onLoad(options) {
        this.currentStatus = options.status ? '' + options.status : 'null'
        this.swiperLazyShow[this.currentStatus] = true
        this.token = getToken()
        if (!this.token) {
            this.$refs.toast.show('请先登录').then(res => {
                uni.redirectTo({
                    url: '/pages/my/my'
                })
            })
            return false
        }
    },
    methods: {
        clickTabs(e) {
            console.log('clickTabs', e)
            this.currentStatus = e.status
            this.tabsCurrentIndex = e.index
            this.swiperLazyShow[e.status] = true
        },
        // 订单列表左右切换
        swiperChange(e) {
            const itemId = e.detail.currentItemId
            this.currentStatus = itemId
            this.swiperLazyShow[itemId] = true
            this.statusList.forEach((item, index) => {
                if (item.status === itemId) {
                    this.tabsCurrentIndex = index
                }
            })
        }
    }
}
</script>

<style lang="scss" scoped>
.tabs {
  background: #FFFFFF;
  height: 45px;
}

.swiper {
  //background: #FFFFFF;
  height: calc(100% - 45px);
}

</style>
