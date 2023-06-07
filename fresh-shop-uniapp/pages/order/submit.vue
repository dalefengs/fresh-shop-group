<template>
    <pageWrapper>
        <view class="select-address">
            <view class="address" v-if="shipmentType === '0'" @click="addressShow">
                <view class="icon">
                    <u-icon name="map" color="#2979ff" size="36"></u-icon>
                </view>
                <view v-if="addressId === 0" class="title">请选择收货地址</view>
                <view v-else>
                    <view class="title king-ellipsis2">{{ address.title + address.detail }}</view>
                    <view class="sub-title">
                        <text class="king-mr-10">{{ address.name }}{{ address.sex === 1 ? '先生' : '女士' }}</text>
                        <text>{{ address.mobile }}</text>
                        <view class="king-ml-10 king-inline-block">
                            <u-tag :text="address.lable" plain shape="circle" size="mini"></u-tag>
                        </view>
                    </view>
                </view>
            </view>
            <view class="food-type">
                <u-radio-group
                        v-model="shipmentType"
                        placement="row"
                        @change="shipmentTypeChange"
                >
                    <u-radio :customStyle="radioLeftStyle" size="24" iconSize="18" labelSize="18" label="配送"
                             name="0"></u-radio>
                    <u-radio :customStyle="radioRightStyle" size="24" iconSize="18" labelSize="18" label="自提"
                             name="1"></u-radio>
                </u-radio-group>
            </view>
        </view>
        <view class="goods-list">
            <view>订单详情</view>
            <view class="goods-info" v-for="cart in list" :key="cart.ID">
                <image :src="cart.goods.images ? cart.goods.images[0].url : ''"
                       class="goods-image king-radius10"
                       mode=""></image>
                <view class="goods-info-box">
                    <view class="goods-info-spec">
                        <text class="goods-name king-ellipsis2">{{ cart.goods.name }}</text>
                        <text class="spe">规格：{{ cart.goods.weight ? cart.goods.weight + 'g/' : '' }}{{
                                cart.goods.unit
                            }}
                        </text>
                        <text class="spe">
	                        单价：{{
                                cart.goods.price > 0 && cart.goods.price < cart.goods.costPrice ? cart.goods.price : cart.goods.costPrice
                            }}
	                        <text v-if="pointGoodsId === 0">元</text>
	                        <text v-if="pointGoodsId > 0"> 积分</text>
                        </text>
                    </view>
                    <view class="goods-box">
                        <text class="goods-price">
                            <text class="goods-symbol" v-if="pointGoodsId === 0">¥</text>
                            {{
                                cart.goods.price > 0 && cart.goods.price < cart.goods.costPrice ? cart.goods.price * cart.num : cart.goods.costPrice * cart.num
                            }}
	                          <text v-if="pointGoodsId > 0">积分</text>
                            <!--                            <text class="goods-unit"> / {{ cart.goods.unit }}</text>-->
                        </text>
                        <view class="goods-num-box">
                            <view class="goods-num">
                                <text>{{ cart.num }} {{ cart.goods.unit }}</text>
                            </view>
                        </view>
                    </view>
                </view>
            </view>
        </view>
        <view class="remark">
            <view>备注</view>
            <u--input placeholder="请输入备注信息" :customStyle="{'padding': '6px 0 0 0'}" border="bottom"
                      v-model="remark"
                      clearable
            ></u--input>
        </view>
        <view style="height: 70px"></view>
        <view class="statistics-box">
            <text class="total">合计：</text>
            <text class="text-color"><text v-if="pointGoodsId === 0">¥</text>
	            {{ total }}
	            <text class="text-color" style="font-size: 20px; margin-left: 3px" v-if="pointGoodsId > 0">积分</text>
            </text>
		        <view v-if="pointGoodsId > 0" class="btn" @tap="showPointPayDialog">
			        <text>积分支付</text>
		        </view>
            <view v-else class="btn" @tap="submit">
                <text>提交订单</text>
            </view>
        </view>
        <addressPop :show="showLoginDialog" @close="addressClose" :addressId="addressId"
                    @checked="addressChecked"></addressPop>
		    <u-modal :show="showPointPay" title="确认购买" @confirm="submit"
		             @cancel="() => showPointPay = false"
		             @close="() => showPointPay = false"
		             :showCancelButton="true"
		             :closeOnClickOverlay="true"
		             confirmText="支付">
			    <view class="king-center king-black6">
				    积分余额：{{ pointAmount }}<br>
				    您确认购买该商品么？
			    </view>
		    </u-modal>
        <u-toast ref="toast" style="z-index: 9999 !important"></u-toast>
    </pageWrapper>
</template>

<script>
import {getCheckedCartList} from "@/api/cart";
import config from '@/config/config.js'
import {getToken} from '@/store/storage.js'
import {getDefaultAddressInfo} from '@/api/address';
import {createOrder, getOrderStatus} from '@/api/order';
import addressPop from '@/components/addressPop/addressPop'
import {getAccountInfo} from "@/api/account.js";
import {getGoodsInfo} from '@/api/goods.js'

export default {
    components: {
        addressPop,
    },
    data() {
        return {
            token: '',
            list: [],
	          pointGoodsId: 0,
            addressId: 0, // 用户地址ID
            address: {}, // 用户收货地址
            total: 0,
            showLoginDialog: false,
            shipmentType: '0', // 配送方式 1配送 2自提
            remark: "", // 备注
            radioLeftStyle: {
                "width": "50%",
                "justify-content": "center",
                "border-right": "1px solid #eee",
            },
            radioRightStyle: {
                "width": "50%",
                "justify-content": "center",
            },
		        showPointPay: false, // 积分购买提示框
		        pointAmount: 0, // 积分数量
        }
    },
		onLoad(options) {
			if (options.pointGoodsId) {
				this.pointGoodsId = parseInt(options.pointGoodsId)
			}
			console.log('pointGoodsId', this.pointGoodsId)
		},
    mounted() {
        this.token = getToken()
        if (!this.token) {
            this.$message(this.$refs.toast).error("请请先登录").then(() => {
                uni.redirectTo({
                    url: '/pages/my/my'
                })
            })
            return
        }
				if (this.pointGoodsId > 0) {
					  this.getPointGoodsData()
				}else {
					  this.getCartListData()
				}
        this.getAddressInfo()
    },
    methods: {
	    // 显示积分支付
	    async showPointPayDialog() {
		    this.showPointPay = true
		    const res = await getAccountInfo(2, this.$refs.toast) // 积分
		    if (res.code !== 0){
			    this.showPointPay = false
			    return false
		    }
		    this.pointAmount = res.data.account.amount
	    },
        // 订单提交
        async submit() {
						const data = {
							remarks: this.remark,
							addressId: this.addressId,
							shipmentType: parseInt(this.shipmentType)
						}
						if (this.pointGoodsId) {
							data.pointGoodsId = this.pointGoodsId
						}
            const res = await createOrder(data, this.$refs.toast)
            if (res.code !== 0) {
							  this.showPointPay = false
                return false
            }
						// 如果是积分商品
						if (this.pointGoodsId > 0) {
							this.showPointPay = false
							await this.$message(this.$refs.toast).success("兑换成功")
							uni.redirectTo({
								url: '/pages/order/detail?id=' + res.data.order.ID
							})
							return true
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
                        this.$message(this.$refs.toast).error("取消支付").then(() => {
                            uni.redirectTo({
                                url: '/pages/order/detail?id=' + order.ID
                            })
                        })
                        return false
                    }
                    this.$message(this.$refs.toast).error("支付失败").then(() => {
                        uni.redirectTo({
                            url: '/pages/order/detail?id=' + order.ID
                        })
                    })
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
                    // 只允许错误重试 30 次 30秒
                    if (errCount > 30) {
                        clearInterval(statusInterval); // 清除定时器
                        this.$message(this.$refs.toast).hide()
                        this.$message(this.$refs.toast).error("获取交易结果超时，请稍后查看").then(() => {
                            uni.redirectTo({
                                url: '/pages/order/detail?id=' + orderId
                            })
                        })
                    }
                    return false;
                }
                count++
                // 只允许重试获取状态 30 次 30秒
                if (count > 30) {
                    clearInterval(statusInterval); // 清除定时器
                    this.$message(this.$refs.toast).hide()
                    this.$message(this.$refs.toast).error("获取交易结果超时，请稍后查看").then(() => {
                        uni.redirectTo({
                            url: '/pages/order/detail?id=' + orderId
                        })
                    })
                    return
                }
                if (res.data.status === 1) {
                    clearInterval(statusInterval); // 清除定时器
                    this.$message(this.$refs.toast).hide()
                    // 进行其他操作
                    this.$message(this.$refs.toast).success("支付成功").then(() => {
                        uni.redirectTo({
                            url: '/pages/order/detail?id=' + orderId
                        })
                    })
                }
            }, 1000);
        },
        // 地址选择
        addressChecked(addressInfo) {
            this.address = addressInfo
            this.addressId = addressInfo.ID
            this.addressClose()
        },
        // 显示地址
        addressShow() {
            this.showLoginDialog = true
        },
        addressClose() {
            this.showLoginDialog = false
        },
	      async getPointGoodsData() {
		      const data = {
			      ID: this.pointGoodsId,
		      }
		      const res = await getGoodsInfo(data, this.$refs.toast)
		      res.data.regoods.images.forEach((item, index) => {
			      if (item.url.slice(0, 4) !== 'http') {
				      res.data.regoods.images[index].url = config.baseUrl + "/" + item.url
			      }
		      })
		      if (res.data.regoods.weight > 1000) {
			      res.data.regoods.weight = res.data.regoods.weight / 1000 + 'kg'
		      } else {
			      res.data.regoods.weight = res.data.regoods.weight + 'g'
		      }
					const g = res.data.regoods
					const c = {
						goodsId: g.ID,
						specType: 0,
						num: 1,
						goods: g
					}
					this.list.push(c)
		      this.total = g.costPrice.toFixed(2)
	      },
        async getCartListData() {
            const res = await getCheckedCartList(this.$refs.toast)
            if (res.code !== 0) {
                this.$message(this.$refs.toast).error(res.msg)
                return
            }
            if (!res.data.list || res.data.list.length === 0) {
                this.$message(this.$refs.toast).error("购物车中无数据").then(() => {
                    uni.redirectTo({
                        url: '/pages/cart/cart'
                    })
                })
                return
            }
            res.data.list.forEach(item => {
                if (item.goods.images && item.goods.images[0].url.slice(0, 4) !== 'http') {
                    item.goods.images[0].url = config.baseUrl + "/" + item.goods.images[0].url
                }
                // 计算总价格

                if (item.goods.price > 0 && item.goods.price < item.goods.costPrice) { // 如果设置了优惠价格
                    this.total += item.goods.price * item.num
                } else { // 没设置优惠价格
                    console.log(this.total, item.goods.costPrice, item.num, this.total + item.goods.costPrice * item.num)
                    this.total += item.goods.costPrice * item.num
                }
            })
            this.total = this.total.toFixed(2)
            this.list = res.data.list
        },
        // 获取用户默认收货地址
        getAddressInfo() {
            getDefaultAddressInfo(this.$refs.toast).then(res => {
                this.address = res.data
                this.addressId = res.data.ID
            })
        },
		    // 选择框切换
		    shipmentTypeChange(e) {
			    console.log('shipmentTypeChange', e)
			    if (e === '1') {
				    this.address = {}
				    this.addressId = 0
			    }else {
						this.getAddressInfo()
			    }
		    }
    }
}
</script>

<style lang="scss" scoped>
page {
  padding: 0 10px;
}

.select-address {
  background: #fff;
  border-radius: 10px;
  padding: 20px 20px;
  display: flex;
  justify-content: center;
  flex-direction: column;

  .address {
    display: flex;
    align-items: center;
    padding-bottom: 14px;
    border-bottom: 1px solid #F2F2F2;
		margin-bottom: 20px;
  }

  .food-type {
    margin: 0px 0 2px 0;
  }

  .icon {
    margin-right: 20px;
  }

  .title {
    font-size: 18px;
    font-weight: 550;
  }

  .sub-title {
    margin-top: 8px;
    font-size: 17px;
    color: #454749;
    display: flex;
    align-items: center;
  }
}

.remark {
  margin-top: 20px;
  padding: 14px 20px;
  color: #454749;
  background: white;
  border-radius: 10px;

  .input {

  }
}

.goods-list {
  background: white;
  border-radius: 10px;
  padding: 20px 20px 10px 20px;
  margin-top: 10px;

  .goods-info {
    display: flex;
    flex-direction: row;
    align-items: center;
    padding: 20px 0;
    border-bottom: 1px solid #F2F2F2;

    &:last-child {
      border-bottom: none;
    }

    .goods-image {
      width: 75px;
      height: 75px;
    }

    .goods-info-box {
      width: 80%;
      display: flex;
      align-items: center;
      justify-content: space-between;

      .goods-info-spec {
        margin-left: 12px;
        display: flex;
        flex-direction: column;
        font-size: 18px;
      }

      .goods-name {
        font-size: 16px;
        font-weight: 400;
        color: #313133;
      }

      .spe {
        margin-top: 5px;
        font-size: 16px;
        font-weight: 400;
        color: #919298;
      }

      .goods-box {
        margin-top: 9px;
        display: flex;
        flex-direction: column;
        align-items: flex-end;

        .goods-symbol {
          font-size: 18px;
          font-weight: 400;
          color: rgb(50, 38, 38);
          margin-right: 2px;
        }

        .goods-price {
          font-size: 20px;
          font-weight: 400;
          color: rgb(50, 38, 38);
          margin-right: 2px;

          .goods-unit {
            font-size: 14px;
            color: #6a7076;
            margin-left: 2px;
          }
        }

        .goods-num-box {
          .goods-num {
            text-align: center;
            font-size: 16px;
            font-weight: 400;
            color: #666666;
          }
        }
      }
    }
  }
}

.statistics-box {
  width: 100%;
  height: 60px;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: flex-end;
  background-color: #FFFFFF;
  position: fixed;
  bottom: 0;

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
    margin: 0 20px;
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

.u-radio-group {
  justify-content: center;
}

</style>
