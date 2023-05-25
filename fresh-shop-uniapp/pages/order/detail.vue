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
                        <text v-if="order.statusCancel > 0">已取消</text>
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
                    <text v-else-if="order.status === 2">{{ order.shipmentTime | parseDate }}</text>
                    <text v-else-if="order.status === 3">{{ order.receiveTime | parseDate }}</text>
                </view>
            </view>
        </view>
        <view class="address">
            <view>
                <text class="king-mr-20">{{ order.shipmentName }}</text>
                <text>{{ order.shipmentMobile }}</text>
            </view>
            <view class="detail">{{ order.shipmentAddress }}</view>
        </view>

        <view class="goods-box">
            <view class="goods-text">共 {{ order.num }} 件商品</view>
            <view class="goods-list">
                <view v-for="item in order.details" :key="item.ID" class="goods">
                    <view class="king-flex">
                        <view class="goods-image-box">
                            <img class="goods-image" :src="item.goodsImage">
                        </view>
                        <view>
                            <view class="goods-name">{{ item.goodsName }}</view>
                            <view class="goods-spec">规格：{{ item.specKeyName }}</view>
                            <view class="goods-price">¥{{ item.total }}</view>
                        </view>
                    </view>
                    <view class="king-flex king-flex-row king-space-between-col goods-right ">
                        <view class="goods-num">x{{ item.num }}</view>
                        <!-- 付款发货收货状态且 订单没有取消 且 退款状态为未退款 且售后状态为未售后-->
                        <view class="king-mb-5" v-if="[1, 2, 3].includes(order.status) && order.statusCancel === 0 && order.statusRefund === 0">
<!--                            <u-button size="mini" type="info" text="退款"></u-button>-->
                        </view>
                    </view>
                </view>
            </view>
        </view>

        <view class="order">
            <view class="order-info">
                <text class="order-left">订单编号：</text>
                <text class="order-right">{{ order.orderSn }}</text>
            </view>
            <view class="order-info">
                <text class="order-left">支付状态：</text>
                <text class="order-right">{{ order.status > 0 ? '已支付' : '未支付' }}</text>
            </view>
            <view class="order-info">
                <text class="order-left">下单时间：</text>
                <text class="order-right">{{ order.CreatedAt | parseDate }}</text>
            </view>
            <view class="order-info" v-if="order.status > 0">
                <text class="order-left">支付时间：</text>
                <text class="order-right">{{ order.payTime | parseDate }}</text>
            </view>
            <view class="order-info">
                <text class="order-left" v-if="[2,3].includes(order.status)">发货时间：</text>
                <text class="order-right">{{ order.shipmentTime | parseDate }}</text>
            </view>
            <view class="order-info" v-if="order.status === 3">
                <text class="order-left">收货时间：</text>
                <text class="order-right">{{ order.receiveTime | parseDate }}</text>
            </view>
            <view class="order-info" v-if="order.statusCancel > 0">
                <text class="order-left">取消时间：</text>
                <text class="order-right">{{ order.cancelTime | parseDate }}</text>
            </view>
            <view class="order-info" v-if="order.statusCancel > 0">
                <text class="order-left">取消类型：</text>
                <text class="order-right">
                    <text v-if="order.statusCancel === 1">用户取消</text>
                    <text v-else-if="order.statusCancel === 2">后台取消</text>
                    <text v-else-if="order.statusCancel === 3">超时取消</text>
                </text>
            </view>
        </view>

        <view style="height: 50px"></view>
        <view class="menu">
            <!-- 订单状态未付款 且 取消状态0 且未售后-->
            <view class="menu-btn" v-if="order.status === 0 && order.statusCancel === 0 && order.statusRefund === 0 && (!order.return || !order.return.ID)">
                <u-button type="info" :customStyle="menuBtnStyle"  @click="cancelOrder">取消订单</u-button>
            </view>
            <!-- 订单状态已付款 且 取消状态0 且未售后 -->
            <view class="menu-btn" v-if="order.status > 0 && order.statusCancel === 0 && (!order.return || !order.return.ID)" >
                <u-button type="info" :customStyle="menuBtnStyle" @click="refundOrder">申请售后</u-button>
            </view>
            <view class="menu-btn" v-if="order.return && order.return.ID">
                <u-button type="primary" :customStyle="menuBtnStyle" @click="confirmOrder">售后详情</u-button>
            </view>
            <view class="menu-btn" v-if="order.status === 2">
                <u-button type="primary" :customStyle="menuBtnStyle" @click="confirmOrder">确认收货</u-button>
            </view>
            <view class="menu-btn" v-if="order.status === 0 && order.statusCancel === 0">
                <u-button type="primary" :customStyle="menuBtnStyle" @click="goPay">立即支付</u-button>
            </view>
        </view>
        <u-toast ref="toast"></u-toast>
    </pageWrapper>
</template>

<script>
import config from '@/config/config.js'
import {getOrderInfo, confirmOrder, cancelOrder, orderPay, getOrderStatus} from '@/api/order'
import {getToken} from '@/store/storage.js'

export default {
    data() {
        return {
            token: '',
            orderId: '',
            order: {},
            menuBtnStyle: {
                'border-radius': '20px',
            }
        }
    },
    onLoad(options) {
        this.orderId = parseInt(options.id)
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
        // 获取订单信息
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
        },
        // 订单提交
        async goPay() {
            const res = await orderPay({
                ID: this.orderId
            }, this.$refs.toast)
            if (res.code !== 0) {
                return false
            }
            if (!res.data.pay) {
                this.$message(this.$refs.toast).error("交易失败，请重试")
                return false
            }
            this.toPay(res.data.pay, res.data.order)
        },
        // 发起微信支付
        toPay(pay, order) {
            this.$message(this.$refs.toast).loading('正在支付中...')
            const payment = {
                provider: 'wxpay', // 服务提供商，通过 uni.getProvider 获取。
                timeStamp: pay.timestamp,
                nonceStr: pay.nonceStr,
                orderInfo: pay.order,
                package: 'prepay_id=' + pay.prePayId,
                signType: pay.signType,
                paySign: pay.paySign,
                success: res => {
                    console.log('success', res)
                    this.$message(this.$refs.toast).hide()
                    this.paySuccess(order.ID)
                },
                fail: res => {
                    console.log('fail', res)
                    this.$message(this.$refs.toast).hide()
                    if (res.errMsg === 'requestPayment:fail cancel') {
                        this.$message(this.$refs.toast).error("取消支付")
                        return false
                    }
                    this.$message(this.$refs.toast).error("支付失败")
                }
            }
            console.log('payment', payment)
            uni.requestPayment(payment)
        },
        // 支付成功回调
        paySuccess(orderId) {
            this.$message(this.$refs.toast).loading('正在获取支付结果...')
            let errCount = 0
            const statusInterval = setInterval(async () => {
                const res = await getOrderStatus(orderId);
                if (res.code !== 0) {
                    errCount ++
                    // 只允许重试 30 次 30秒
                    if (errCount > 30) {
                        clearInterval(statusInterval); // 清除定时器
                        this.$message(this.$refs.toast).hide()
                        this.$message(this.$refs.toast).error("获取交易结果超时，请稍后查看")
                    }
                    return false;
                }
                if (res.data.status === 1) {
                    clearInterval(statusInterval); // 清除定时器
                    this.$message(this.$refs.toast).hide()
                    // 进行其他操作
                    this.$message(this.$refs.toast).success("支付成功").then(() => {
                        this.getOrderInfoData()
                    })
                }
            }, 1000);
        },
    }
}
</script>

<style lang="scss" scoped>

.status {
  width: 100%;
  height: 120px;
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
  margin-top: 6px;
  background: #FFFFFF;
  padding: 16px 10px;

  .detail {
    color: #6a7076;
    margin-top: 5px;
    font-size: 16px;
  }
}

.goods-box {
  margin-top: 6px;
  background: #FFFFFF;
  padding: 10px 10px;

  .goods-text {
    color: #333333;
    font-size: 16px;
    padding-bottom: 5px;
    margin-bottom: 10px;
    border-bottom: 1px solid #F2F2F2;
  }

  .goods-list {
    .goods {
      display: flex;
      justify-content: space-between;
      border-bottom: 1px solid #F2F2F2;
      padding: 8px 0;
      &:last-child {
        border-bottom: none;
      }

      .goods-image-box {
        width: 80px;
        height: 80px;
        margin-right: 10px;
        flex-shrink: 0;

        .goods-image {
          width: 100%;
          height: 100%;
          border-radius: 5px;
        }
      }

      .goods-name {
        font-size: 15px;
        width: 95%;
        overflow: hidden;
        text-overflow: ellipsis;
        display: -webkit-box;
        // 控制行数
        -webkit-line-clamp: 2;
        -webkit-box-orient: vertical;
      }

      .goods-spec {
        color: #6a7076;
        margin-top: 2px;
        font-size: 14px;
      }

      .goods-price {
        color: #FF0000;
        font-size: 14px;
        margin-top: 2px;
      }

      .goods-num {
        text-align: right;
        font-size: 14px;
        font-weight: bold;
        color: #767a82;
      }
    }
  }
}

.order {
  background: #FFFFFF;
  padding: 10px 10px;
  margin-top: 6px;

  .order-info {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 3px 0;

    .order-left {
      color: #333333;
      font-size: 15px;
    }

    .order-right {
      color: #6a7076;
      font-size: 15px;
    }
  }
}

.menu {
  position: fixed;
  width: 100%;
  height: 50px;
  background: #FFFFFF;
  box-shadow: 0px 2px 10px 0px rgba(0, 0, 0, 0.05);
  bottom: 0;
  display: flex;
  align-items: center;
  justify-content: flex-end;
  padding-right: 10px;
  box-sizing: border-box;

  .menu-btn {
    margin-left: 5px;
  }
}
</style>
