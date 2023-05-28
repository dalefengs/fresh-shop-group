<!--
 * @Author: likfees
 * @Date: 2023-03-23 22:16:00
 * @LastEditors: likfees
 * @LastEditTime: 2023-04-21 17:26:13
-->
<template>
    <view>
        <u-popup :show="show" @open="open" @close="close" :closeable="closeable">
            <view class="title">选择用户地址</view>
            <view class="box">
                <scroll-view :scroll-y="true" style="height: 45vh">
                    <view class="list" v-for="item in list" :key="item.ID" @tap="checkedAddress(item)">
                        <view class="radio">
                            <image src="../../static/select.png" v-if="item.ID == addressId" class="checked-image"
                                   mode=""></image>
                            <image src="../../static/not_select.png" v-else class="checked-image" mode=""
                            ></image>
                        </view>
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

<!--                    TODO 收货地址跳转    <view class="edit" @click="toUpdateAddress(item.ID)">
                            <u-icon name="edit-pen" size="30"></u-icon>
                        </view>-->
                    </view>
                </scroll-view>
            </view>
            <view class="btn">
                <u-button :customStyle="otherBtnStyle" @click="toCreateAddress">添加地址</u-button>
            </view>
        </u-popup>
        <u-toast style="z-index:9998;" ref="toast"></u-toast>
    </view>
</template>

<script>
import {getAddressList} from '@/api/address';

export default {
    name: "addressPop",
    data() {
        return {
            list: [], // 收货地址列表
            otherBtnStyle: {
                width: "80%",
                height: "50px",
                borderRadius: "30px",
                background: "#2979ff",
                color: "#fff",
            },
        };
    },
    props: {
        show: {
            type: Boolean,
            default: false,
        },
        closeable: {
            type: Boolean,
            default: true,
        },
        // 当前选择的收货地址 ID
        addressId: {
            type: Number,
            default: true,
        }
    },
    mounted() {
        uni.getSystemInfo({
            success: (res) => {
                this.windowHeight = res.windowHeight - this.subHeight;
            },
        });
    },
    methods: {
        async open() {
            //this.$emit("open");
            await this.getAddressListData()
        },
        // 选中时间
        checkedAddress(addressInfo) {
            this.$emit("checked", addressInfo)
        },
        // 获取收货地址
        async getAddressListData() {
            const res = await getAddressList(this.$refs.toast)
            if (res.code !== 0) {
                this.$message(this.$refs.toast).error(res.msg)
                return false
            }
            this.list = res.data.list
        },
        close() {
            this.$emit("close");
        },
        toCreateAddress() {
            uni.navigateTo({
                url: '/pages/address/addressForm?submit=1'
            })
        }
    }
}
</script>

<style lang="scss">

.box {
  height: 45vh;
  margin-bottom: 55px;
}

.title {
  margin: 12px 15px;
}

.list {
  padding: 10px 10px 10px 20px;
  background: #FFFFFF;
  margin-top: 10px;
  margin-bottom: 10px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 18px;
  border-bottom: 1px solid #dddddf;

  &:nth-last-child(1) {
    border: none;
  }

  .address {
    width: 80%;
  }

  .radio {
    width: 45px;
    height: 35px;

    .checked-image {
      width: 30px;
      height: 30px;
    }
  }

  .info {
    margin-top: 10px;
    color: #6a7076;
    display: flex;
    font-size: 17px;
  }

  .edit {
    width: 60px;
    display: flex;
    justify-content: center;
  }
}

.btn {
  width: 100%;
  display: flex;
  justify-content: center;
  position: absolute;
  bottom: 5px;
}
</style>
