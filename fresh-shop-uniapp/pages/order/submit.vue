<template>
    <pageWrapper>
        <view>
            <view class="statistics-box">
                <text class="total">合计：</text>
                <text class="text-color">¥{{ total }}</text>
                <view class="btn" @tap="submit">
                    <text>提交订单</text>
                </view>
            </view>
        </view>
        <u-toast ref="toast"></u-toast>
    </pageWrapper>
</template>

<script>
import {getCheckedCartList} from "@/api/cart";
import config from '@/config/config.js'
import {getToken} from '@/store/storage.js'

export default {
    data() {
        return {
            token: '',
            list: [],
            total: 0,
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
    },
    methods: {
        async getCartListData() {
            const res = await getCheckedCartList()
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
        }
    }
}
</script>

<style lang="scss" scoped>

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
