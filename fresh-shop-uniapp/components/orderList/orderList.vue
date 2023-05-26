<template>
    <view>
        <scroll-view   scroll-y="true"
                     :style="{ height: scrollViewHeight + 'px'}" refresher-enabled="true" :refresher-threshold="70"
                     :refresher-triggered="triggered" @refresherrefresh="onRefresh" @scrolltolower="scrollTolower"
                     :scroll-anchoring="true">
            <!-- 商品列表 -->
            <view v-if="list.length > 0">
                <view class="order" v-for="(order, index) in list" :key="index">
                    <view class="order-status">
                        <view>{{ order.CreatedAt | parseDate }}</view>
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
                                    <text v-else-if="order.return === 1">审核通过</text>
                                    <text v-else-if="order.return === 1 && order.return.refundStatus === 2">售后完成</text>
                                </text>
                                <!-- 普通订单 -->
                                <text v-else>
                                    <text v-if="order.status === 0">等待支付</text>
                                    <text v-else-if="order.status === 1">备货中</text>
                                    <text v-else-if="order.status === 2">已发货</text>
                                    <text v-else-if="order.status === 3">已发货</text>
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
                        <view class="btn" v-if="order.statusCancel === 0 && order.status === 0" @click="() => showCancelOrder = true">取消订单</view>
                        <view class="btn" v-if="order.statusCancel > 0 || order.status === 3" @click="() => showDeletelOrder = true">删除订单</view>
                        <view class="btn" v-if="order.statusCancel === 0 || order.status === 3" @click="() => showConfirmlOrder = true">确认收货</view>
                        <view class="btn btn-pay" v-if="order.status === 0 && order.statusCancel === 0" @click="goPay(order.ID)">立即支付</view>
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
    </view>
</template>

<script>
import config from '@/config/config.js'
import {getOrderList, confirmOrder, cancelOrder, orderPay, getOrderStatus} from '@/api/order'
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
            console.log(res);
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