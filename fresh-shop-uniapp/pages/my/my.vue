<!--
 * @Author: likfees
 * @Date: 2023-04-14 13:36:30
 * @LastEditors: likfees
 * @LastEditTime: 2023-04-23 14:12:03
-->
<template>
    <pageWrapper>
        <view class="header" v-bind:class="{ 'status': isH5Plus }" v-if="token">
            <view class="userinfo">
                <view class="face">
                    <image :src="user.headerImg"></image>
                </view>
                <view class="info">
                    <view class="username" @click="() => showChangeNickName = true">{{ user.nickName }}</view>
<!--                    <view class="integral" @click="copyInvitationCode(user.invitationCode)">-->
<!--                        <text style="margin-right: 4px;">推荐码：{{ user.invitationCode ? user.invitationCode : '' }}</text>-->
<!--                        <u-icon name="file-text" color="#fff"></u-icon>-->
<!--                    </view>-->
                    <view class="integral">
                      积分：{{ point ? point : 0 }}
                    </view>
                </view>
            </view>
            <view class="setting">
<!--                <image src="../../static/my/setting.png"></image>-->
            </view>
        </view>
        <!-- 未登录  -->
        <view v-else class="header" v-bind:class="{ 'status': isH5Plus }">
            <view class="userinfo" @click="showLogin">
                <view class="face">
                    <image src="../../static/my/face.png"></image>
                </view>
                <view class="login-btn">登录 / 注册</view>
            </view>

        </view>
        <view class="orders">
            <view class="box">
                <view class="title">我的订单</view>
                <view class="king-flex">
                    <view class="label" v-for="(row, index) in orderTypeLise" :key="row.name" hover-class="hover"
                          @tap="toOrderType(index)">
                        <view class="icon">
                            <view class="badge" v-if="orderStatusCount[row.badge] > 0">{{ orderStatusCount[row.badge] }}</view>
                            <image class="icon-img" :src="'../../static/my/' + row.icon"></image>
                        </view>
                        <view style="width: 100%; text-align: center;">{{ row.name }}</view>
                    </view>
                </view>
            </view>
        </view>
        <!-- 菜单 -->
        <view class="list" v-for="(list, index) in severList" :key="index">
            <view v-for="(li, key) in list" :key="li.name">
                <!-- 条件判断未登录 不显示菜单的条件  -->
                <view v-if="!severList[index][key].isLogin || severList[index][key].isLogin && token" class="li"
                      @click="menuClick(index, key)" v-bind:class="{ 'noborder': key == list.length - 1 }"
                      hover-class="hover">
                    <view class="icon">
                        <image :src="'../../static/my/list/' + li.icon"></image>
                    </view>
                    <view class="text">{{ severList[index][key].name }}</view>
                    <image class="to" src="../../static/my/to.png"></image>
                </view>
            </view>

        </view>
        <!-- 登录 -->
        <loginPop :show="showLoginDialog" @close="hideLogin" @success="loginSuccess"/>
        <!-- 拨号 -->
        <u-modal :show="showPhoneDialog" showCancelButton closeOnClickOverlay @confirm="callPhone"
                 @cancel="() => showPhoneDialog = false" @close="close" confirmText="拨号">
            <view>联系电话：{{ relationPhone }}</view>
        </u-modal>
        <Tabbar :tabsId="3"/>
        <u-toast ref="toast" style="z-index: 9999"></u-toast>
        <u-modal :show="showChangeNickName" title="修改昵称" @confirm="changeNickName"
                 @cancel="() => showChangeNickName = false"
                 @close="() => showChangeNickName = false"
                 :showCancelButton="true"
                 :closeOnClickOverlay="true"
                 confirmText="修改">

            <view class="king-center">
                <u--input
                        placeholder="请输入昵称"
                        border="surround"
                        v-model="user.nickName"
                ></u--input>
            </view>
        </u-modal>
    </pageWrapper>
</template>
<script>
import Tabbar from '@/components/tabbar/tabbar.vue'
import loginPop from '@/components/loginPop/loginPop.vue'
import {getUserInfo, setSelfInfo} from "@/api/user";
import {getOrderStatusCount} from "@/api/order";
import {getUser, getToken, setUser, setToken, setOpenId} from '@/store/storage.js'

export default {
    components: {
        Tabbar,
        loginPop
    },
    data() {
        return {
            //#ifdef APP-PLUS
            isH5Plus: true,
            //#endif
            //#ifndef APP-PLUS
            isH5Plus: false,
            //#endif
            user: {},
	          point: 0, // 积分余额
            orderStatusCount: {},
            token: '',
            showLoginDialog: false, // 登录
            showPhoneDialog: false, // 拨号
            relationPhone: "13527326320", // 联系人电话
            showChangeNickName: false, // 修改昵称模态框
            orderTypeLise: [
                //name-标题 icon-图标 badge-角标
                {name: '全部订单', icon: 'qrbu.png', badge: 0, status: null},
                {name: '待付款', icon: 'fukr.png', badge: 'unpaid', status: 0},
                {name: '备货中', icon: 'faho.png', badge: 'delivered', status: 1},
                {name: '待收货', icon: 'uzho.png', badge: 'shipped', status: 2},
                // {name: '售后', icon: 'uzhz.png', badge: 0, status: 10}
            ],
            severList: [
                [
                    { name: '积分兑换', icon: 'point_shop.png', handle: this.toPointGoods, isLogin: true},
                    // { name: '积分明细', icon: 'finance.png', handle: this.showPhone },
                    {name: '收货地址', icon: 'address.png', handle: this.toAddress},
                    {name: '清除缓存', icon: 'clear.png', handle: this.clearStorage},
                    {name: '联系我们', icon: 'relation.png', handle: this.showPhone},
                    {name: '退出登录', icon: 'logout.png', handle: this.logout, isLogin: true},
                ]
            ],
        };
    },
    onLoad() {
        //加载
        this.init();
    },
    methods: {
        async init() {
            this.token = getToken()
            //如果登录了，则获取用户信息
            if (this.token) {
                const res = await getUserInfo()
                if (res.code === 401) {
                  this.user = ""
                  this.point = 0
                  this.token = ""
                }
                this.user = res.data.userInfo
	              this.point = res.data.point
                setUser(this.user)
                // 获取订单状态数量
                const orderStatusRes = await getOrderStatusCount()
                this.orderStatusCount = orderStatusRes.data
                console.log('this.orderStatusCount', this.orderStatusCount)
            }
        },
        // 登录成功
        async loginSuccess(u) {
            console.log("loginSuccess", u);
            this.hideLogin();
            this.token = getToken()
            this.user = u
            // 获取积分
            const res = await getUserInfo()
            this.point = res.data.point
            this.$message(this.$refs.toast).success("登录成功")
        },
        async changeNickName() {
            const res = await setSelfInfo({nickName: this.user.nickName})
            if (res.code !== 0){
                return false
            }
            this.showChangeNickName = false
            this.$message(this.$refs.toast).success("修改成功")
        },
        toPointGoods(){
						uni.navigateTo({
							url: '/pages/goods/pointGoods'
						})
        },
        toAddress() {
            uni.navigateTo({
                url: '/pages/address/address'
            })
        },
        // 显示登录框
        showLogin() {
            this.showLoginDialog = true
        },
        // 隐藏登录框
        hideLogin() {
            this.showLoginDialog = false
        },
        // 菜单栏点击
        menuClick(index, key) {
            this.severList[index][key].handle()
        },
        // 显示手机号
        showPhone() {
            this.showPhoneDialog = true
        },
        // 拨打电话
        callPhone() {
            uni.makePhoneCall({
                phoneNumber: this.relationPhone,
                success: (result) => {
                },
                fail: (error) => {
                }
            })
        },
        // 点击退出登录
        logout() {
            uni.showModal({
                title: '退出登录',
                content: '是否确认退出登录',
                showCancel: true,
                success: ({confirm, cancel}) => {
                    if (confirm) {
                        this.token = ''
                        this.user = ''
                        setUser('')
                        setToken('')
                        setOpenId('')
                        this.$message(this.$refs.toast).success("退出登录")
                    }
                }
            })
        },
        // 复制邀请码
        copyInvitationCode(code) {
            uni.setClipboardData({
                data: code,
                success: () => {
                    this.$message(this.$refs.toast).success("复制推荐码成功")
                }
            })
        },
        // 清除缓存
        clearStorage() {
            uni.showModal({
                title: '清除缓存',
                content: '请谨慎操作',
                showCancel: true,
                success: ({confirm, cancel}) => {
                    if (confirm) {
                        this.token = ''
                        this.user = ''
                        uni.clearStorage()
                        uni.clearStorageSync()
                        this.$message(this.$refs.toast).success("清除缓存成功")
                    }
                }
            })

        },
        //用户点击订单类型
        toOrderType(index) {
            uni.navigateTo({
                url: '/pages/order/list?status=' + this.orderTypeLise[index].status
            })
        },
        //用户点击列表项
        toPage(list_i, li_i) {
            uni.showToast({title: this.severList[list_i][li_i].name});
        }
    }
}
</script>

<style lang="scss">
.header {
  &.status {
    padding-top: var(--status-bar-height);
  }

  background-color: #2979ff;
  width: 92%;
  height: 50vw;
  padding: 0 4%;
  display: flex;
  align-items: center;

  .userinfo {
    width: 90%;
    display: flex;
    position: relative;
    bottom: 16px;
    align-items: center;

    .login-btn {
      font-size: 19px;
      margin: 15px 20px;
      color: #fff;
      cursor: pointer;
    }

    .face {
      flex-shrink: 0;
      width: 20vw;
      height: 20vw;
      margin-top: 6px;

      image {
        width: 100%;
        height: 100%;
        border-radius: 100%
      }
    }

    .info {
      flex-flow: wrap;
      padding-left: 15px;

      .username {
        width: 100%;
        color: #fff;
        font-size: 20px;
        margin-left: 5px;
      }

      .integral {
        display: flex;
        align-items: center;
        padding: 0 10px;
        height: 20px;
        color: #fff;
        background-color: rgba(0, 0, 0, 0.1);
        border-radius: 10px;
        font-size: 12px;
        margin-top: 5px;
      }
    }
  }

  .setting {
    flex-shrink: 0;
    width: 6vw;
    height: 6vw;
    position: relative;
    bottom: 16px;

    image {
      width: 100%;
      height: 100%
    }
  }
}

.hover {
  background-color: #eee
}

.orders {
  background-color: #2979ff;
  width: 92%;
  height: 11vw;
  padding: 0 4%;
  margin-bottom: calc(20vw + 20px);
  display: flex;
  align-items: flex-start;
  border-radius: 0 0 100% 100%;
  margin-top: -1px;

  .box {
    width: 98%;
    padding: 0 1%;
    height: 34vw;
    background-color: #fefefe;
    border-radius: 12px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.15);
    position: relative;
    top: -35px;

    //display: flex;
    //align-items: center;
    //justify-content: center;
    .title {
      margin: 12px 24px;
    }

    .label {
      display: flex;
      align-items: center;
      justify-content: center;
      flex-flow: wrap;
      width: 100%;
      height: 16vw;
      color: #666666;
      font-size: 13px;

      .icon {
        position: relative;
        width: 40px;
        height: 40px;
        margin: 0 1vw;

        .icon-img {
          width: 40px;
          height: 40px;
        }

        .badge {
          position: absolute;
          width: 4vw;
          height: 4vw;
          background-color: #ec6d2c;
          top: 0vw;
          right: 0vw;
          border-radius: 100%;
          font-size: 10px;
          color: #fff;
          display: flex;
          align-items: center;
          justify-content: center;
          z-index: 10;
        }

        image {
          width: 7vw;
          height: 7vw;
          z-index: 9;
        }
      }
    }
  }
}

.list {
  width: 92%;
  margin: 0 auto;
  border-radius: 14px;
  //border-bottom: solid 13px #f1f1f1;
  background: #fff;
  position: relative;
  top: -35px;

  .li {
    width: 92%;
    height: 50px;
    padding: 0 4%;
    border-bottom: solid 1px #e7e7e7;
    display: flex;
    align-items: center;

    &.noborder {
      border-bottom: 0
    }

    .icon {
      flex-shrink: 0;
      width: 25px;
      height: 25px;

      image {
        width: 25px;
        height: 25px
      }
    }

    .text {
      padding-left: 10px;
      width: 100%;
      color: #666
    }

    .to {
      flex-shrink: 0;
      width: 20px;
      height: 20px
    }
  }
}
</style>
