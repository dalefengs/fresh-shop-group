<template>
    <view class="common-car">
        <view class="empty" v-if="list.length === 0">
            <!-- 购物车未空 -->
            <u-empty mode="car" icon="http://cdn.uviewui.com/uview/empty/car.png">
                <view>
                    <u-button type="primary" :customStyle="btnStyle" text="去逛逛" @click="toHome"></u-button>
                </view>
            </u-empty>
        </view>
        <view class="shop-car" v-else>
            <scroll-view scroll-y="true"
                         :style="{ height: height + 'px' }" refresher-enabled="true"
                         :refresher-threshold="70"
                         :refresher-triggered="triggered" @refresherrefresh="onRefresh"
                         :scroll-anchoring="true">
                <view class="store-box">
                    <!--                <view class="store-header">
                                        <view class="store-header-left">
                                            <image src="../../static/select.png" v-if="isCheckAll" class="checked-image" mode=""
                                                   @tap="checkedAllGoods()"></image>
                                            <image src="../../static/not_select.png" v-else class="checked-image" mode=""
                                                   @tap="checkedCancleAllGoods()">
                                            </image>
                                            <text>启运冷鲜</text>
                                        </view>
                                        <text>编辑</text>
                                        &lt;!&ndash;                    <image src="../../static/youjiantou1.png" class="label" mode=""></image>&ndash;&gt;
                                    </view>-->
                    <view class="goodsInfo" v-for="(cart,index) in list" :key="index"
                          @longtap="showDeleteCartDalog(index)">
                        <view class="goodsInfo-left" @tap="checkedGoods(cart.ID, cart.checked, index)">
                            <image src="../../static/select.png" v-if="cart.checked == 1" class="checked-image"
                                   mode=""></image>
                            <image src="../../static/not_select.png" v-else class="checked-image" mode=""
                            ></image>
                        </view>
                        <view class="goodsInfo-right">
                            <image :src="cart.goods.images ? cart.goods.images[0].url : ''"
                                   class="goods-image king-radius10"
                                   mode=""></image>
                            <view class="goodsInfo-box">
                                <text class="goods-name">{{ cart.goods.name }}</text>
                                <!--                            <text class="spe">规格：{{ cart.goods.remark }}</text>-->
                                <view class="goods-box">
                                    <text class="goods-price">
                                        ¥{{ cart.goods.price > 0 ? cart.goods.price : cart.goods.costPrice }}
                                        <text class="goods-unit"> / {{ cart.goods.unit }}</text>
                                    </text>
                                    <view class="goods-num-box">
                                        <view class="goods-image"
                                              @tap="addCartReq(cart.goods.ID, index, 2, cart.num - 1)">
                                            <text>-</text>
                                        </view>
                                        <view class="goods-num">
                                            <text>{{ cart.num }}</text>
                                        </view>
                                        <view class="goods-image"
                                              @tap="addCartReq(cart.goods.ID,index,1, cart.num + 1)">
                                            <text>+</text>
                                        </view>
                                    </view>
                                </view>
                            </view>
                        </view>
                    </view>
                </view>
            </scroll-view>
            <view class="statistics-box">
                <view class="statistics">
                    <view class="statistics-left" v-if="isCheckAll" @tap="allCheck">
                        <image src="../../static/select.png" class="checked-image" mode="">
                        </image>
                        <text>全选</text>
                    </view>
                    <view class="statistics-left" v-else @tap="allCheck">
                        <image src="../../static/not_select.png" class="checked-image" mode="">
                        </image>
                        <text>全选</text>
                    </view>
                    <view class="statistics-right" v-if="isCut" @tap="accounts">
                        <text class="total">合计：</text>
                        <text class="text-color">¥{{ total }}</text>
                        <view class="btn">
                            <text>结算</text>
                        </view>
                    </view>
                    <view class="statistics-right" v-else @tap="delect">
                        <view class="btn">
                            <text>删除</text>
                        </view>
                    </view>
                </view>
            </view>
        </view>
        <!-- 删除弹窗 -->
        <u-modal :show="showDeleteCart" :showCancelButton="true" :closeOnClickOverlay="true" @confirm="deleteCart"
                 @cancel="hideDeleteCartDalog" @close="hideDeleteCartDalog">
            <text class="king-font-weight-100 king-font22">
                确定删除商品吗？
            </text>
        </u-modal>
        <u-toast ref="toast"></u-toast>
    </view>
</template>

<script>
import {addCart, updateCart, selectAllCart, clearSelectAllCart, deleteCartByIds} from "@/api/cart";

export default {
    name: "commonCar",
    data() {
        return {
            statisticsIndex: false,
            total: 0,
            isCut: true, // 是否编辑
            isCheckAll: false, // 是否全选
            showDeleteCart: false, // 删除购物车
            currentDeleteIndex: 0, // 当前长按删除的商品索引
            btnStyle: {
                width: '120px',
                marginTop: '20px',
                borderRadius: '20px',
            },
        }
    },
    props: {
        list: {
            type: [Array],
            default: []
        },
        height: {
            type: Number,
            default: 0
        },
        // 下拉刷新状态
        triggered: {
            type: Boolean,
            default: false
        }
    },
    watch: {
        list: {
            handler(newVal, oldVal) {
                let checkedAll = true
                if (newVal.length > 0) {
                    newVal.forEach(item => {
                        if (item.checked === 0) {
                            checkedAll = false
                        }
                    })
                    this.statistics()
                }
                this.isCheckAll = checkedAll
            }
        }
    },
    methods: {
        // 结算
        accounts() {
            uni.navigateTo({
                url: '/pages/order/submit'
            })
        },
        // 显示删除提示框
        showDeleteCartDalog(id) {
            this.currentDeleteIndex = id
            this.showDeleteCart = true
        },
        hideDeleteCartDalog() {
            this.showDeleteCart = false
        },
        async deleteCart() {
            let ids = []
            ids.push(this.list[this.currentDeleteIndex].ID)
            let data = {
                ids: ids
            }
            const res = await deleteCartByIds(data, this.$refs.toast)
            if (res.code !== 0) {
                return
            }
            this.$emit('deleteCart', this.currentDeleteIndex)
            this.hideDeleteCartDalog()
        },
        //商品选择
        async checkedGoods(id, checked, index) {
            checked = checked === 1 ? 0 : 1
            const data = {
                ID: id,
                checked: checked
            }
            const res = await updateCart(data, this.$refs.toast)
            if (res.code !== 0) {
                return
            }
            this.list[index].checked = checked
            // 是否全选
            let checkedAll = true
            this.list.forEach(item => {
                if (item.checked === 0) {
                    checkedAll = false
                }
            })
            this.isCheckAll = checkedAll
            this.statistics()
        },
        // type 1增 2减
        // num 数量
        async addCartReq(id, index, type, num) {
            const data = {
                goodsId: id,
                specType: 0, // 单规格
                num: num
            }
            if (num < this.list[index].goods.minCount) {
                this.$message(this.$refs.toast).error("商品数量不能小于 " + this.list[index].goods.minCount)
                return
            }
            if (num < 1) {
                this.$message(this.$refs.toast).error("商品数量不能小于 1")
                return
            }
            const res = await addCart(data)
            if (res.code === 0) {
                this.list[index].num = num
            }
            this.statistics()
        },
        //全选
        async allCheck() {
            if (this.isCheckAll) {
                const res = await clearSelectAllCart(this.$refs.toast)
                if (res.code !== 0) {
                    return
                }
                this.list.forEach(item => {
                    item.checked = 0
                })
                this.isCheckAll = false
            } else {
                const res = await selectAllCart(this.$refs.toast)
                if (res.code !== 0) {
                    return
                }
                this.list.forEach(item => {
                    item.checked = 1
                })
                this.isCheckAll = true
            }
            this.statistics()
        },
        //统计
        statistics() {
            let total = 0
            this.list.forEach(c => {
                if (c.checked !== 1) {
                    return
                }
                if (c.goods.price > 0) {
                    total += c.goods.price * c.num
                } else {
                    total += c.goods.constPrice
                }
            })
            this.total = total.toFixed(2)
        },
        cut() {
            this.isCut = !this.isCut
            this.statisticsIndex = true
            this.allCheck()
        },
        // 刷新
        onRefresh() {
            this.$emit('onRefresh')
        },
        toHome() {
            uni.navigateTo({
                url: '/pages/index/index'
            })
        }
    }
}
</script>

<style lang="scss" scoped>
.empty {
  background: white;
  height: 70vh;
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
}

.checked-image {
  width: 30px;
  height: 30px;
}


.common-car {
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  margin: 0 auto;
}

.shop-car {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 6px;

  .store-box {
    margin-bottom: 10px;
    display: flex;
    flex-direction: column;
    align-items: center;
    background-color: #FFFFFF;

    .store-header {
      width: 100%;
      height: 39px;
      padding: 0 15px;
      display: flex;
      flex-direction: row;
      align-items: center;
      justify-content: space-between;
      font-size: 14px;
      font-weight: 400;
      color: #313133;

      .store-header-left {
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: space-between;
      }

      text {
        font-size: 14px;
        font-weight: 400;
        color: #313133;
        margin-left: 10px;
      }

      .label {
        width: 7px;
        height: 12px;
        margin-left: 5px;
        margin-top: 5px;
      }
    }

    .goodsInfo {
      width: 100%;
      height: 123px;
      display: flex;
      flex-direction: row;
      align-items: center;
      border-bottom: 1px solid #EDEDED;

      &:nth-last-child(1) {
        border: none;
      }

      .goodsInfo-left {
        width: 45px;
        height: 35px;
      }

      .goodsInfo-right {
        height: 123px;
        margin-left: 10px;
        display: flex;
        flex-direction: row;
        align-items: center;

        .goods-image {
          width: 70px;
          height: 70px;
        }

        .goodsInfo-box {
          margin-left: 12px;
          display: flex;
          flex-direction: column;
          align-items: center;

          .goods-name {
            width: 214px;
            font-size: 16px;
            font-weight: 400;
            color: #313133;
          }

          .spe {
            width: 214px;
            margin-top: 5px;
            font-size: 12px;
            font-weight: 400;
            color: #919298;
          }

          .goods-box {
            width: 214px;
            margin-top: 9px;
            display: flex;
            flex-direction: row;
            align-items: center;
            justify-content: space-between;

            .goods-price {
              font-size: 22px;
              font-weight: 400;
              color: rgba(242, 18, 18, 1);
              margin-right: 2px;

              .goods-unit {
                font-size: 14px;
                color: #6a7076;
                margin-left: 2px;
              }
            }

            .goods-num-box {
              width: 88px;
              height: 23px;
              display: flex;
              flex-direction: row;
              align-items: center;

              .goods-image {
                width: 23px;
                height: 23px;
                text-align: center;
                line-height: 23px;
                border: 1px solid #CFCFCF;
                font-size: 26px;
              }

              .goods-num {
                width: 38px;
                height: 23px;
                text-align: center;
                line-height: 22px;
                font-size: 17px;
                font-weight: 400;
                color: #666666;
                border-top: 1px solid #CFCFCF;
                border-bottom: 1px solid #CFCFCF;
              }
            }
          }
        }
      }
    }
  }

  .statistics-box {
    display: flex;
    flex-direction: column;
    align-items: center;
    background-color: #FFFFFF;
    position: fixed;
    bottom: 50px;
    border-bottom: 1px solid #e0dfdf;

    .statistics {
      height: 60px;
      display: flex;
      flex-direction: row;
      align-items: center;
      justify-content: space-between;

      .statistics-left {
        width: 70px;
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: space-between;

        image {
          width: 30px;
          height: 30px;
        }

        text {
          font-size: 16px;
          font-weight: 400;
          color: #313133;
        }
      }

      .statistics-right {
        width: 300px;
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: flex-end;

        .total {
          color: #313133;
        }

        .btn {
          width: 109px;
          height: 46px;
          line-height: 46px;
          background: #2979ff;
          text-align: center;
          font-size: 15px;
          font-weight: 500;
          color: white;
          margin-left: 20px;
          border-radius: 40px;
        }

        text {
          font-size: 17px;
          font-weight: 400;
          color: white;
        }

        .text-color {
          font-size: 22px;
          color: rgba(242, 18, 18, 1);
        }
      }
    }
  }
}
</style>
