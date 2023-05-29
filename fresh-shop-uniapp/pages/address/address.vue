<template>
    <pageWrapper>
        <view class="box" v-for="item in list" :key="item.ID" @longtap="showDeleteAddressDialog(item.ID)">
            <view class="address">
                <text>
                    {{ item.address }}{{ item.title }}{{ item.detail }}
                </text>
                <view class="info">
                    <text>{{ item.name }}{{ item.sex === 1 ? '先生' : '女士' }}</text>
                    <text class="king-ml-10">{{ item.mobile }}</text>
                    <view class="king-ml-10 king-inline-block" v-if="item.isDefault === 1">
                        <u-tag text="默认" plain shape="circle" size="mini"></u-tag>
                    </view>
                    <view class="king-ml-10 king-inline-block">
                        <u-tag :text="item.lable" plain shape="circle" size="mini"></u-tag>
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
        <u-modal :show="showDeleteAddress" title="删除收货地址" @confirm="toDeleteAddress"
                 @cancel="() => showDeleteAddress = false"
                 @close="() => showDeleteAddress = false"
                 :showCancelButton="true"
                 :closeOnClickOverlay="true"
                 confirmText="删除">
            <text class="king-center">
                此操作不可撤销，请谨慎操作！
            </text>
        </u-modal>
        <u-toast ref="toast" style="z-index: 9999"></u-toast>
    </pageWrapper>
</template>

<script>
import {getAddressList, deleteAddress} from '@/api/address';
import {getToken} from '@/store/storage.js'

export default {
    data() {
        return {
            token: '',
            showDeleteAddress: false,
            currentDeleteId : 0,
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
        showDeleteAddressDialog(id) {
            this.currentDeleteId = id
            this.showDeleteAddress = true
        },
        async toDeleteAddress() {
            const res = await deleteAddress({ID: this.currentDeleteId}, this.$refs.toast)
            if (res.code !== 0) {
                return  false
            }
            this.showDeleteAddress = false
            await this.$message(this.$refs.toast).success("删除成功")
            await this.getAddressListData()
        },
        async getAddressListData() {
            const res = await getAddressList(this.$refs.toast)
            if (res.code !== 0) {
                // this.$message(this.$refs.toast).error(res.msg)
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
  font-size: 18px;
  .address {
    width: 85%;
  }
  .info {
    margin-top: 10px;
    color: #6a7076;
    display: flex;
    align-items: center;
    font-size: 17px;
  }

  .edit {
    width: 60px;
    display: flex;
    justify-content: center;
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
