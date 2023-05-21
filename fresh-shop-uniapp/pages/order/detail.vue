<template>
    <pageWrapper>
        <view class="status">
            <view class="status-icon">
                <u-icon name="car-fill" color="#fff" size="28"></u-icon>
            </view>
            <view class="status-text">
                <view>
                    <view>
                        <text>待付款</text>
                        <text>备货中</text>
                        <text>配送中</text>
                        <text>已完成</text>
                        <text>已取消</text>
                    </view>
                    <view>
                        <text>售后中</text>
                        <text>拒绝售后</text>
                        <text>已退款</text>
                        <text>退款失败</text>
                    </view>
                </view>
                <view>
                    <text>{{ order.payTime }}</text>
                    <text>{{ order.shipmentTime }}</text>
                    <text>{{ order.shipmentTime }}</text>
                    <text>{{ order.receiveTime }}</text>
                    <text>{{ order.cancelTime }}</text>
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

export default {
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
</style>
