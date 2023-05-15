<template>
    <pageWrapper>
        <view class="select-address" @click="addressShow">
            <view class="icon">
                <u-icon name="map" color="#2979ff" size="40"></u-icon>
            </view>
            <view v-if="addressId === 0" class="title">请选择收货地址</view>
            <view v-else>
                <view class="title">{{ address.title + address.detail }}</view>
                <view class="sub-title">
                    <text class="king-mr-10">{{ address.name }}{{ address.sex === 1 ? '先生' : '女士' }}</text>
                    <text>{{ address.mobile }}</text>
                    <view class="king-ml-10 king-inline-block">
                        <u-tag :text="address.lable" plain shape="circle" size="mini"></u-tag>
                    </view>
                </view>
            </view>
        </view>
        <view class="statistics-box">
            <text class="total">合计：</text>
            <text class="text-color">¥{{ total }}</text>
            <view class="btn" @tap="submit">
                <text>提交订单</text>
            </view>
        </view>
        <addressPop :show="showLoginDialog" @close="addressClose" :addressId="addressId"
                    @checked="addressChecked"></addressPop>
        <u-toast ref="toast"></u-toast>
    </pageWrapper>
</template>

<script>
import {getCheckedCartList} from "@/api/cart";
import config from '@/config/config.js'
import {getToken} from '@/store/storage.js'
import {getDefaultAddressInfo} from '@/api/address';
import addressPop from '@/components/addressPop/addressPop'

export default {
    components: {
        addressPop,
    },
    data() {
        return {
            token: '',
            list: [],
            addressId: 0, // 用户地址ID
            address: {}, // 用户收货地址
            type: 0, // 提货方式 0运输 1自提
            total: 0,
            showLoginDialog: false,

        }
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
        this.getCartListData()
        this.getAddressInfo()
    },
    methods: {
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
                if (item.goods.price > 0) { // 如果设置了优惠价格
                    this.total += item.goods.price * item.num
                } else { // 没设置优惠价格
                    this.total += item.goods.costPrice * item.num
                }
            })
            this.list = res.data.list
        },
        // 获取用户默认收货地址
        getAddressInfo() {
            getDefaultAddressInfo(this.$refs.toast).then(res => {
                this.address = res.data
                this.addressId = res.data.ID
            })
        }
    }
}
</script>

<style lang="scss" scoped>

.select-address {
  display: flex;
  align-items: center;
  height: 120px;
  margin: 0 10px;
  background: #fff;
  border-radius: 10px;
  padding: 0 20px;

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
</style>
