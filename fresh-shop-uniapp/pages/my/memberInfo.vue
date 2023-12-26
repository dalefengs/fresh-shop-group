<template>
    <pageWrapper>
        <view class="box">

        </view>
        <view class="box">
            <view class="con-bd-item">
                <view class="con-bd-item-name"><span style="color: red; margin-right: 2px">* </span>联系人名称</view>
                <view class="con-bd-item-con">
                    <input type="text" v-model="formData.originContactName" placeholder="请输入您的姓名"
                           placeholder-class="input-placeholder" />
                </view>
            </view>
            <view class="con-bd-item">
                <view class="con-bd-item-name">客户名称</view>
                <view class="con-bd-item-con">
                    <input type="text" v-model="formData.originCustomerName" placeholder="请输入您所在公司/店铺名称"
                           placeholder-class="input-placeholder" />
                </view>
            </view>
        </view>
    </pageWrapper>

</template>

<script>
import { getToken, setUser, getUser } from "@/store/storage";

export default {
    data() {
        return {
            token: '',
            user: null,
            formData: {}
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
        this.user = getUser()
        this.formData.ID = this.user.ID
        // TODO
    },
}
</script>

<style scoped lang="scss">
.box {
  margin: 10px 10px 12px;
  height: 100px;
  border-radius: 10px;
  box-shadow: 0px 0px 20px #f4f3f3;
  background: #FFFFFF;
  padding: 20px 15px 30px;
}
.con-bd-item {
  display: flex;
  align-items: center;
  width: 100%;
  height: 110rpx;
}

.con-bd-item-name {
  width: 120px;
  height: 110px;
  text-align: left;
  font-size: 16px;
  color: #333333;
  line-height: 110px;
}

.con-bd-item-con {
  display: flex;
  align-items: center;
  width: calc(100% - 120rpx);
  height: 108rpx;
  border-bottom: 2rpx solid #F7F7F7;
}

.con-bd-item-con input {
  width: 100%;
  height: 60rpx;
}
</style>
