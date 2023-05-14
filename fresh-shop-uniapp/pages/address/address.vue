<template>
    <pageWrapper>
        <view class="box" v-for="item in list" :key="item.ID">
            <view class="address">
                <text>
                    {{ item.address }}{{ item.title }}{{ item.detail }}
                </text>
                <view class="info">
                    <text>{{ item.name }}{{ item.sex === 1 ? '先生' : '女士' }}</text>
                    <text class="king-ml-10">{{ item.mobile }}</text>
                    <view class="king-ml-10 king-inline-block" v-if="item.isDefault === 1">
                        <u-tag text="默认" plain shape="circle"></u-tag>
                    </view>
                </view>
            </view>
            <view class="edit" @click="toUpdateAddress(item.ID)">
                <u-icon name="edit-pen" size="30"></u-icon>
            </view>
        </view>
        <view style="height: 50px"></view>
        <view class="btn-box" @click="toCreateAddress">
            <view class="btn">添加收货地址</view>
        </view>
        <u-toast ref="toast"></u-toast>
    </pageWrapper>
</template>

<script>
import {getAddressList} from '@/api/address';
import {getToken} from '@/store/storage.js'

export default {
    data() {
        return {
            token: '',
            list: [],
        }
    },
    onLoad() {
        this.token = getToken()
        if (!this.token) {
            this.$message(this.$refs.toast).error("请请先登录").then(() => {
                uni.redirectTo({
                    url: '/pages/my/my'
                })
            })
            return
        }
    },
    onShow() {
        this.getAddressListData()
    },
    methods: {
        async getAddressListData() {
            const res = await getAddressList()
            if (res.code !== 0) {
                this.$message(this.$refs.toast).error(res.msg)
                return false
            }
            this.list = res.data.list
        },
        toCreateAddress() {
            uni.navigateTo({
                url: '/pages/address/addressForm'
            })
        },
        toUpdateAddress(id) {
            uni.navigateTo({
                url: '/pages/address/addressForm?id=' + id
            })
        }
    }
}
</script>

<style lang="scss" scoped>

.box {
  padding: 20px 10px 20px 20px;
  background: #FFFFFF;
  margin-top: 10px;
  margin-bottom: 10px;
  display: flex;
  justify-content: space-between;
  align-items: center;

  .info {
    margin-top: 10px;
    color: #6a7076;
  }

  .edit {
    width: 60px;
  }
}

.btn-box {
  position: fixed;
  bottom: 10px;
  height: 50px;
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;

  .btn {
    width: 70%;
    height: 50px;
    color: #fff;
    background-color: #2979ff;
    border-radius: 30px;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 20px;
  }
}
</style>
