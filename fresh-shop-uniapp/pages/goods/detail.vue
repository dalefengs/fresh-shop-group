<!--
 * @Author: likfees
 * @Date: 2023-04-23 14:20:03
 * @LastEditors: likfees
 * @LastEditTime: 2023-04-23 18:08:12
-->

<template>
    <pageWrapper>
        <u-swiper v-if="Object.keys(goods).length > 0" :list="goods.images" @change="e => currentImgIndex = e.current"
                  :autoplay="false" indicatorStyle="right: 20px" height="280">
            <view slot="indicator" class="indicator-num">
                <text class="indicator-num__text">{{ currentImgIndex + 1 }}/{{ goods.images.length }}</text>
            </view>
        </u-swiper>
        <!-- 商品头部 -->
        <view class="box" v-show="Object.keys(goods).length > 0">
            <view class="price-box">
                <view>
                    <text class="price">￥{{ goods.costPrice || '0' }}</text>
                    <text class="unit">/{{ goods.unit }}</text>
                    <text v-if="goods.price < goods.costPrice" class="del-price">￥{{
                            goods.price || '0'
                        }}
                    </text>
                </view>
                <view class="origin" v-show="goods.origin">
                    原产地：{{ goods.origin }}
                </view>
            </view>
            <view class="title">
                <text class="king-mr-10" v-if="goods.brand && goods.brand.name">{{ goods.brand.name }}</text>
                <text class="king-mr-10">{{ goods.name }}</text>
                <text class="king-mr-10" v-if="goods.weight">{{ goods.weight }}</text>
            </view>
            <view class="store">
                库存：
                <text v-if="goods.store > 100">充足</text>
                <text v-else-if="goods.store > 0 && goods.store < 100">{{ goods.store }}</text>
                <text v-else style="color: #fa3534;">缺货</text>
            </view>
        </view>
        <!-- 商品详情 -->
        <view class="box-desc" v-if="Object.keys(goods).length > 0">
            <view class="desc-title">商品详情</view>
            <view class="content">
                <u-parse :content="goods.desc.details" :tagStyle="tagStyle"></u-parse>
            </view>
        </view>
        <view class="tabbar">
            <view class="left">
                <view class="btn" @click="favoritesClick">
                    <view class="king-inline-block" style="height: 25px; margin-top:2px;">
                        <u-icon :name="goods.isFavorite ? 'heart-fill' : 'heart'"
                                :color="goods.isFavorite ? '#fa3534' : ''"
                                size="23"></u-icon>
                        <!--<u-icon name="" size="23"></u-icon>-->
                    </view>
                    <view>收藏</view>
                </view>
                <view class="btn" @click="toCartPage">
                    <view class="king-inline-block king-relative" style="height: 26px;">
                        <u-icon name="shopping-cart" size="28"></u-icon>
                        <view class="badge">
                            <u-badge max="99" :value="goods.cartTotalNum" shape="circle"></u-badge>
                        </view>
                    </view>
                    <view>购物车</view>
                </view>
            </view>
            <view class="right">
                <view v-if="goods.cartNum > 0" class="addCartBtnSelect">
                    <view class="symbol" @click="addCartClick(2)">
                        -
                    </view>
                    <view class="num">
                        {{ goods.cartNum }}
                    </view>
                    <view class="symbol" @click="addCartClick(1)">
                        +
                    </view>
                </view>
                <view v-else class="addCartBtn" :style="{'background-color': goods.store <= 0 || goods.store <= goods.minCount ? '#f29100' : '#2979ff'}" @click="addCartClick(1)">
                    <text v-if="goods.store <= 0 || goods.store <= goods.minCount"> 库存不足 </text>
                    <text v-else-if="goods.minCount > 1">最低 {{ goods.minCount }} 件起购</text>
                    <text v-else>加入购物车</text>
                </view>
            </view>
        </view>
        <loginSuspend :show="loginSuspendShow" @success="loginSuccess"></loginSuspend>
        <u-toast style="z-index:9998;" ref="toast"></u-toast>
    </pageWrapper>
</template>

<script>
import {getToken} from '@/store/storage.js'
import loginSuspend from '@/components/loginPop/loginSuspend.vue'
import {getGoodsInfo} from '@/api/goods.js'
import {favorites} from '@/api/favorites.js'
import config from '@/config/config.js'
import {addCart} from "@/api/cart";

export default {
    components: {
        loginSuspend
    },
    data() {
        return {
            id: 0, // 商品id
            goods: {},
            token: '',
            num: 0,
            loginSuspendShow: false, // 是否显示底部登录
            currentImgIndex: 0,
            tagStyle: {
                img: "width: 100%;vertical-align: bottom;"
            },
            addCartBtnStyle: {
                borderRadius: '20px',
            }
        }
    },
    onLoad(options) {
        if (options.id) {
            this.id = parseInt(options.id)
        } else {
            uni.showToast({
                title: '参数错误',
                icon: 'error',
                duration: 1500,
                success: () => {
                    setTimeout(() => {
                        uni.navigateBack({
                            delta: 1
                        })
                    }, 1500)
                }
            })
            return
        }
        const token = getToken()
        if (!token) {
            this.loginSuspendShow = true
        }
        this.getGoods()
    },
    onReady() {
    },
    mounted() {
    },
    methods: {
        // 登陆成功
        loginSuccess() {
            this.loginSuspendShow = false
        },
        getGoods() {
            const data = {
                ID: this.id,
            }
            getGoodsInfo(data, this.$refs.toast).then(res => {
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
                this.goods = res.data.regoods

            })
        },
        // 收藏商品
        favoritesClick() {
            favorites({
                goodsId: this.id
            }).then(res => {
                if (res.code !== 0) {
                    return false
                }
                this.goods.isFavorite = !this.goods.isFavorite
                this.$message(this.$refs.toast).success('收藏成功')
            })
        },
        // 添加购物车按钮
        // type 1增 2减
        addCartClick(type) {
            let currentNum = this.goods.cartNum
            let addNum = currentNum
            if (type === 1) {
                addNum++
            } else {
                addNum--
            }
            // 如果有最低购买的数量 且 当前商品没有达到超过的最低购买的数量
            if (this.goods.minCount > 0 && addNum < this.goods.minCount) {
                if (type === 1) {
                    addNum = this.goods.minCount
                } else {
                    addNum = 0
                }

            }
            // 添加购物车
            this.addCartReq(1, addNum)
        },

        // type 1增 2减
        // num 数量
        async addCartReq(type, num) {
            const data = {
                goodsId: this.id,
                specType: 0, // 单规格
                num: num
            }
            const res = await addCart(data)
            if (res.code === 0) {
                if (type === 1) {
                    this.$message(this.$refs.toast).success("添加购物车成功")
                    this.goods.cartTotalNum = this.goods.cartTotalNum + (num - this.goods.cartNum)
                } else {
                    this.goods.cartTotalNum = this.goods.cartTotalNum - (this.goods.cartNum - num)
                }
                this.goods.cartNum = num
            }
        },
        toCartPage() {
            uni.navigateTo({
                url: '/pages/cart/cart'
            })
        }

    }
}
</script>

<style lang="scss" scoped>
.indicator {
  @include flex(row);
  justify-content: center;

  &__dot {
    height: 6px;
    width: 6px;
    border-radius: 100px;
    background-color: rgba(255, 255, 255, 0.35);
    margin: 0 5px;
    transition: background-color 0.3s;

    &--active {
      background-color: #ffffff;
    }
  }
}

.indicator-num {
  padding: 2px 0;
  background-color: rgba(0, 0, 0, 0.35);
  border-radius: 100px;
  width: 35px;
  @include flex;
  justify-content: center;

  &__text {
    color: #FFFFFF;
    font-size: 12px;
  }
}

.box {
  margin: auto;
  background: white;
  margin-top: 10px;
  padding: 16px 14px;
  box-sizing: border-box;
}

.price-box {
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;

  .origin {
    font-size: 15px;
  }
}

.price {
  color: #fa3534;
  font-size: 24px;
  font-weight: 600;

  text {
    font-size: 15px;
  }
}

.unit {
  color: #666;
  font-size: 17px;
}

.del-price {
  color: #666;
  font-weight: 600;
  font-size: 15px;
  text-decoration: line-through;
  margin-left: 4px;
}

.sale-num {
  color: #888;
  font-size: 14px;
  font-weight: normal;
  margin-bottom: 2px;
}

.title {
  margin-top: 10px;
  font-size: 20px;
  font-weight: 600;
}

.store {
  margin-top: 8px;
  color: #888;
}

.box-desc {
  margin: auto;
  background: white;
  margin-top: 10px;
  padding: 16px 0;
  box-sizing: border-box;
}

.desc-title {
  font-size: 20px;
  font-weight: 600;
  display: flex;
  align-items: center;
  box-sizing: border-box;
  margin-bottom: 12px;
  margin-left: 10px;
}

.content {
  // 占位底部栏
  padding-bottom: 50px;
}

.tabbar {
  position: fixed;
  bottom: 0;
  background: white;
  width: 100%;
  height: 50px;
  box-shadow: 0px 7px 14px 0px rgba(0, 0, 0, 0.75);
  display: flex;
  align-items: center;
  justify-content: space-between;
  box-sizing: border-box;

  .left {
    width: 45%;
    display: flex;
    justify-content: space-between;

    .btn {
      width: 50%;
      text-align: center;
      margin: 0 auto;
      font-size: 12px;

      .badge {
        position: absolute;
        top: -1px;
        right: -4px;
      }
    }
  }

  .right {
    width: 55%;
    margin-right: 10px;
    font-size: 17px;
    padding: 0 10px;
    margin: 0 auto;

    .addCartBtn {
      width: 100%;
      height: 44px;
      line-height: 42px;
      text-align: center;
      border-radius: 22px;
      background-color: #2979ff;
      color: white;
      font-size: 18px;
    }

    .addCartBtnSelect {
      display: flex;
      justify-content: space-between;
      width: 100%;
      height: 44px;
      text-align: center;
      border-radius: 22px;
      background-color: #2979ff;
      color: white;
      font-size: 18px;
      box-sizing: border-box;

      .symbol {
        width: 32%;
        font-size: 32px;
        vertical-align: middle;
        display: flex;
        align-items: center;
        justify-content: center;
        margin-bottom: 5px;
      }

      .num {
        width: 36%;
        display: flex;
        align-items: center;
        justify-content: center;
        background: #236ce9;
      }
    }
  }
}
</style>
