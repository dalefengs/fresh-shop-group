<template>
    <pageWrapper>
        <view class="status">
            <view class="status-icon">
                <u-icon name="car-fill" color="#fff" size="28"></u-icon>
            </view>
            <view class="status-text">
                <view>
                    <view v-if="order.return && order.return.ID">
                        <text v-if="order.return.status === 0">售后中</text>
                        <text v-else-if="order.return.status === -1">拒绝售后</text>
                        <text v-else-if="order.return.refundStatus === 1">退款中</text>
                        <text v-else-if="order.return.refundStatus === 2">已退款</text>
                        <text v-else-if="order.return.refundStatus === -1">退款失败</text>
                    </view>
                    <view v-else>
                        <text v-if="statusCancel > 0">已取消</text>
                        <text v-else-if="order.status === 0">待付款</text>
                        <text v-else-if="order.status === 1">备货中</text>
                        <text v-else-if="order.status === 2">配送中</text>
                        <text v-else-if="order.status === 3">已完成</text>
                    </view>
                </view>
                <view class="king-mt-5">
                    <!-- TODO 售后时间 -->
                    <text v-if="order.statusCancel > 0">{{ order.cancelTime }}</text>
                    <text v-else-if="order.status === 1">{{ order.payTime | parseDate }}</text>
                    <text v-else-if="order.status === 2">{{ order.shipmentTime | parseDate  }}</text>
                    <text v-else-if="order.status === 3">{{ order.receiveTime | parseDate }}</text>
                </view>
            </view>
        </view>
        <view class="address">
            <view>
                <text class="king-mr-20">{{order.shipmentName}}</text>
                <text>{{order.shipmentMobile}}</text>
            </view>
            <view class="detail">{{ order.shipmentAddress }}</view>
        </view>

        <view class="goods-box">
            <view class="goods-text">共 {{order.num}} 件商品</view>
            <view class="goods-list">
                <view v-for="item in order.details" :key="item.ID" class="goods">
                    <view class="goods-image">
                        <img :src="item.goodsImage">
                    </view>
                    <view>
                        <view class="goods-name">{{item.goodsName}}</view>
                        <view>规格：{{  }}</view>
                    </view>
                </view>
            </view>
        </view>
        <u-toast ref="toast"></u-toast>
    </pageWrapper>
</template>

<script>
import config from '@/config/config.js'
import {getOrderInfo} from '@/api/order'
import {getToken} from '@/store/storage.js'
import UImage from "../../uni_modules/uview-ui/components/u--image/u--image.vue";

export default {
    components: {UImage},
    data() {
        return {
            token: '',
            orderId: '',
            order: {},
        }
    },
    onLoad(options) {
        this.orderId = options.id
        this.token = getToken()
        if (!this.token) {
            this.$refs.toast.show('请先登录').then(res => {
                uni.redirectTo({
                    url: '/pages/my/my'
                })
            })
            return false
        }
        this.getOrderInfoData()
    },
    methods: {
        async getOrderInfoData() {
            const res = await getOrderInfo({
                ID: this.orderId
            })
            if (res.code !== 0) {
                this.$refs.toast.error(res.msg)
                return false
            }
            res.data.reorder.details.forEach(item => {
                console.log('item', item)
                if (item.goodsImage && item.goodsImage.slice(0, 4) !== 'http') {
                    item.goodsImage = config.baseUrl + "/" + item.goodsImage
                }
            })
            this.order = res.data.reorder
            console.log('this.order', this.order)
        }
    }
}
</script>

<style lang="scss" scoped>

.status {
  width: 100%;
  height: 150px;
  background: #2979ff;
  display: flex;
  align-items: center;
  color: #FFFFFF;
  .status-icon {
    margin: 0 20px;
    padding: 10px;
    border-radius: 50%; /* 将边框设置为圆形 */
    border: 2px solid #fff; /* 设置边框样式 */
  }
  .status-text {
    display: flex;
    flex-direction: column;
  }
}

.address {
  margin-top: 5px;
  background: #FFFFFF;
  padding: 10px 10px;
  .detail {
    color: #6a7076;
    margin-top: 5px;
    font-size: 16px;
  }
}
</style>
