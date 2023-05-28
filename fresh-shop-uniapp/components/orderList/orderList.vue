<template>
    <view>
        <scroll-view   scroll-y="true"
                     :style="{ height: scrollViewHeight + 'px'}" refresher-enabled="true" :refresher-threshold="70"
                     :refresher-triggered="triggered" @refresherrefresh="onRefresh" @scrolltolower="scrollTolower"
                     :scroll-anchoring="true">
            <!-- 商品列表 -->
            <view v-if="list.length > 0">
                <view class="order" v-for="(order, index) in list" :key="index" @click="toOrderDetails(order.ID)">
                    <view class="order-status">
                        <text v-if="order.delivery && order.delivery.scheduledTime.charAt(0) !== '0'">预计送达: {{ order.delivery.scheduledTime | parseDate }}</text>
                        <view v-else>{{ order.CreatedAt | parseDate }}</view>
                        <view>
                            <!-- 判断订单是否取消 -->
                            <text v-if="order.statusCancel === 1">已取消</text>
                            <text v-else-if="order.statusCancel === 2">后台取消</text>
                            <text v-else-if="order.statusCancel === 3">超时取消</text>
                            <text v-else>
                                <!-- 判断是否是售后单 -->
                                <text v-if="order.return.ID !== 0">
                                    <text v-if="order.return.refundStatus === 1">退款等待到账</text>
                                    <text v-else-if="order.return === -1">拒绝售后</text>
                                    <text v-else-if="order.return === 0">等待审核</text>
                                    <text v-else-if="order.return === 1 && order.return.refundStatus === 2">售后完成</text>
                                    <text v-else-if="order.return === 1">审核通过</text>
                                </text>
                                <!-- 普通订单 -->
                                <text v-else>
                                    <text v-if="order.status === 0">等待支付</text>
                                    <text v-else-if="order.status === 1">备货中</text>
                                    <text v-else-if="order.status === 2">已发货</text>
                                    <text v-else-if="order.status === 3">已完成</text>
                                </text>
                            </text>
                        </view>
                    </view>
                    <!-- 订单内容 -->
                    <view v-if="order.details && order.details.length > 1" class="goods-list">
                        <view class="list-box" v-for="(item, index) in order.details" :key="index">
                            <image v-if="index <= 3" class="box-goods-img" :src="item.goodsImage"  ></image>
                            <view v-if="index === 4" class="box-goods-ellipsis">
                                ...
                            </view>
                        </view>
                    </view>
                    <view class="single-goods" v-else-if="order.details && order.details.length === 1">
                        <image class="single-goods-img" :src="order.details[0].goodsImage"></image>
                        <view class="single-goods-info">
                            <view class="single-goods-title">
                                <view class="king-ellipsis2 king-black">{{ order.details[0].goodsName }}</view>
                                <view class="single-goods-price">
                                    <text>¥ {{ order.total }}</text>
                                </view>
                            </view>
                            <view class="king-font12 king-black6">
                                <text >单价：¥{{ order.details[0].price }}</text>
                            </view>
                            <view class="king-font12 king-black6">
                                <text>规格：{{ order.details[0].specKeyName }}</text>
                                <text class="king-ml-10">数量：{{ order.details[0].num }}</text>
                            </view>
                        </view>
                    </view>
                    <view class="order-bottom">
                        <text>共 {{ order.num }} 件商品</text>
                        <text class="king-ml-10">实付 <text class="order-bottom-price">¥{{ order.finish }}</text></text>
                    </view>
                    <view class="order-btn">
                        <view class="btn" v-if="order.statusCancel === 0 && order.status === 0" @tap.stop="showCancelOrderDailog(order.ID)">取消订单</view>
                        <view class="btn" v-if="order.statusCancel > 0 || order.status === 3" @tap.stop="showDeleteOrderDailog(order.ID)">删除订单</view>
<!--                        <view class="btn" v-if="order.statusCancel === 0 && order.status === 2" @tap.stop="showConfirmOrderDailog(order.ID)">确认收货</view>-->
                        <view class="btn btn-pay" v-if="order.status === 0 && order.statusCancel === 0" @tap.stop="goPay(order.ID)">立即支付</view>
                    </view>
                </view>
                <view class="king-py-40" @click="scrollTolower">
                    <u-loadmore :status="loadMore" loading-text="努力加载中，请喝杯茶" loadmore-text="上拉加载更多" nomore-text="实在是没有了"/>
                </view>
            </view>
            <u-empty v-else :style="{ height: scrollViewHeight / 1.6 + 'px' }" width="220" height="220" textSize="16"
                     text="暂无订单" mode="data" icon="http://cdn.uviewui.com/uview/empty/data.png" />
            <u-toast style="z-index:9998;" ref="toast"></u-toast>
        </scroll-view>

        <u-modal :show="showCancelOrder" title="取消订单" @confirm="toCancelOrder"
                 @cancel="() => showCancelOrder = false"
                 @close="() => showCancelOrder = false"
                 :showCancelButton="true"
                 :closeOnClickOverlay="true">
            <text class="king-center">
                取消订单操作不可恢复,请谨慎操作！
            </text>
        </u-modal>
        <u-modal :show="showConfirmOrder" title="确认收货" @confirm="toConfirmOrder"
                 @cancel="() => showConfirmOrder = false"
                 @close="() => showConfirmOrder = false"
                 :showCancelButton="true"
                 :closeOnClickOverlay="true"
                 confirmText="确认收货">
            <text class="king-center">
                请您确认货物是否完整！
            </text>
        </u-modal>
        <u-modal :show="showDeleteOrder" title="删除订单" @confirm="toDeleteOrder"
                 @cancel="() => showDeleteOrder = false"
                 @close="() => showDeleteOrder = false"
                 :showCancelButton="true"
                 :closeOnClickOverlay="true"
                 confirmText="删除">
            <text class="king-center">
                此操作不可撤销，请谨慎操作！
            </text>
        </u-modal>
        <u-toast ref="toast" style="z-index: 9999"></u-toast>
    </view>
</template>

<script>
import config from '@/config/config.js'
import {getOrderList, confirmOrder, cancelOrder, orderPay, getOrderStatus, deleteOrder} from '@/api/order'
import UButton from "../../uni_modules/uview-ui/components/u-button/u-button";
export default {
    name: "orderList",
    components: {UButton},
    props: {
        status: { // 订单状态 null:全部 0:未支付 1:备货中 2:已发货 3:已完成 10售后
            type: String,
            default: 'null'
        }
    },
    data() {
        return {
            list: [],
            scrollViewHeight: 0,
            page: {
                page: 1,
                pageSize: 10,
                total: 0,
                isMore: true
            },
            triggered: false, // 下拉刷新状态
            loadMore: 'loadmore', // 上拉加载状态
            showConfirmOrder:false,
            showDeleteOrder:false,
            showCancelOrder:false,
            currentOperateOrderId: 0, // 当前操作的订单ID
        }

    },
    mounted() {
        // 设置商品列表高度为页面高度
        uni.getSystemInfo({
            success: (res) => {
                const windowHeight = res.windowHeight;
                this.scrollViewHeight = windowHeight - 50;
            },
        });
        this.getOrderListData(0)
    },
    methods: {
        // 订单提交
        async goPay(orderId) {
            const res = await orderPay({
                ID: orderId
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
            let count = 0
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
                count++
                if (count > 30) {
                    clearInterval(statusInterval); // 清除定时器
                    this.$message(this.$refs.toast).hide()
                    this.$message(this.$refs.toast).error("获取交易结果超时，请稍后查看")
                    return false;
                }
                if (res.data.status === 1) {
                    clearInterval(statusInterval); // 清除定时器
                    this.$message(this.$refs.toast).hide()
                    // 进行其他操作
                    this.$message(this.$refs.toast).success("支付成功").then(() => {
                        this.getOrderListData(0)
                    })
                }
            }, 1000);
        },
        // type = 1加载 其他为刷新
        async getOrderListData(type) {
            const data = {}
            if (this.status !== 'null') {
                data.status = parseInt(this.status)
            }
            if (type == 0) {
                this.page.page = 1
                this.page.isMore = true
                this.loadMore = 'loadmore'
            }
            data.page = this.page.page
            data.pageSize = this.page.pageSize
            let res = {}
            if (type == 1) { // 加载
                res = await getOrderList(data, false, this.$refs.toast)
            } else {
                res = await getOrderList(data, true, this.$refs.toast)
            }
            if (res.code !== 0) {
                return false
            }
            this.page.page++

            res.data.list.forEach((item) => {
                if (!item.details) {
                    return false
                }
                item.details.forEach((dItem, index) => {
                    if (dItem.goodsImage.slice(0, 4) !== 'http') {
                        item.details[index].goodsImage = config.baseUrl + "/" + dItem.goodsImage
                    }
                })
            })
            // 进行赋值并计算是否还有下一页
            this.page.total = res.data.total
            // 如果没有更多数据，则将isMore设置为false
            if ((this.page.page - 1) * this.page.pageSize >= this.page.total) {
                console.log("没有更多了");
                this.page.isMore = false
                this.loadMore = 'nomore'
            }
            if (type == 1) {
                this.list = [...this.list, ...res.data.list]
            } else {
                this.list = res.data.list
            }
            return true
        },
        // type = 1热门 2 上新
        async onRefresh() {
            this.triggered = true;
            const b = await this.getOrderListData(0)
            if (b) {
                this.$message(this.$refs.toast).success('刷新成功')
            } else {
                this.$message(this.$refs.toast).error('刷新失败')
            }
            this.triggered = false;
        },
        async scrollTolower(e) {
            // 如果是在加载中就不执行或没有更多时
            if (this.loadMore == 'loading' || !this.page.isMore) {
                return
            }
            // 设置状态为加载中
            this.loadMore = 'loading'
            await this.getOrderListData(1)
            // 如果还有更多
            if (this.page.isMore) {
                this.loadMore = 'loadmore'
            } else {
                this.loadMore = 'nomore'
            }
        },
        toOrderDetails(id) {
            uni.navigateTo({
                url: `/pages/order/detail?id=${id}`
            })
        },
        showCancelOrderDailog(id) {
            this.showCancelOrder = true
            this.currentOperateOrderId = id
        },
        showDeleteOrderDailog(id) {
            this.showDeleteOrder = true
            this.currentOperateOrderId = id
        },
        showConfirmOrderDailog(id) {
            this.showConfirmOrder = true
            this.currentOperateOrderId = id
        },
        // 取消订单
        async toCancelOrder() {
            const res = await cancelOrder({
                ID: this.currentOperateOrderId
            }, this.$refs.toast)
            if (res.code !== 0) {
                this.showCancelOrder = false
                return false
            }
            this.showCancelOrder = false
            await this.$message(this.$refs.toast).success('取消订单成功')
            await this.getOrderListData(0)
        },
        // 确认收货
        async toConfirmOrder() {
            const res = await confirmOrder({
                ID: this.currentOperateOrderId
            }, this.$refs.toast)
            if (res.code !== 0) {
                this.showConfirmOrder = false
                return false
            }
            this.showConfirmOrder = false
            await this.$message(this.$refs.toast).success('确认收货成功')
            await this.getOrderListData(0)
        },
        // 删除订单
        async toDeleteOrder() {
            const res = await deleteOrder({
                ID: this.currentOperateOrderId
            }, this.$refs.toast)
            if (res.code !== 0) {
                this.showDeleteOrder = false
                return false
            }
            this.showDeleteOrder = false
            await this.$message(this.$refs.toast).success('删除订单成功')
            await this.getOrderListData(0)
        },
    }
}
</script>

<style lang="scss" scoped>

.order {
  margin: 10px 10px;
  background-color: #fff;
  padding: 8px 10px;
  border-radius: 10px;

  .order-status {
    display: flex;
    justify-content: space-between;
    margin: 5px 0 10px 0;
    font-size: 14px;
    color: #4d4d4d;
  }
  .goods-list {
    display: flex;
    background: #ededed;
    padding: 15px 10px;
    border-radius: 6px;

    .list-box {
      width: 20%;
      display: flex;
      justify-content: center;
      align-items: center;
      .box-goods-img {
        width: 60px;
        height: 60px;
        border-radius: 10px;
      }
      .box-goods-ellipsis {
        font-size: 18px;
        color: #333333;
      }
    }
  }

  .single-goods {
    display: flex;
    background: #ededed;
    padding: 15px 10px;
    border-radius: 6px;

    .single-goods-img {
      width: 65px;
      height: 65px;
      border-radius: 10px;
      flex-shrink: 0;
    }

    .single-goods-info {
      margin-left: 8px;
      width: 100%;

      .single-goods-title {
        font-size: 14px;
        display: flex;
        justify-content: space-between;

        .single-goods-price {
          width: 50px;
          flex-shrink: 0;
          text-align: right;
        }
      }
    }

  }

  .order-bottom {
    margin-top: 5px;
    padding: 5px 0;
    display: flex;
    justify-content: flex-end;
    align-items: center;
    font-size: 15px;
    color: #4d4d4d;

    .order-bottom-price {
      font-size: 16px;
      color: #000;
      margin-left: 3px;
    }
  }

  .order-btn {
    display: flex;
    justify-content: flex-end;
    padding: 6px 0;

    .btn {
      padding: 5px 12px;
      border-radius: 20px;
      border: 1px solid #d5d4d4;
      font-size: 14px;
      margin-left: 6px;
    }
    .btn-pay {
      background-color: #3c9cff;
      color: #FFFFFF;
    }
  }

}


.item-single {
  margin: 5px 10px;
  padding: 8px;
  background-color: #fff;
  width: 100%;
  height: 100px;
}
</style>
