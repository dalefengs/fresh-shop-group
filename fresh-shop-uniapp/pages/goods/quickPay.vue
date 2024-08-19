<template>
    <pageWrapper>
        <!-- 未登录	-->
        <view class="empty" v-if="!token">
            <u-empty mode="car" text="您还没有登陆哦" icon="http://cdn.uviewui.com/uview/empty/car.png">
                <view>
                    <u-button type="primary" :customStyle="toLoginStyle" text="去登录" @click="showLogin"></u-button>
                </view>
            </u-empty>
        </view>
        <view v-else class="king-bg-white">
            <!-- 吸顶标签栏 -->
            <u-sticky bgColor="#fff">
                <u-tabs
                        :list="tabsList"
                        lineWidth="50"
                        :activeStyle="{color: '#303133', fontWeight: 'bold',transform: 'scale(1.05)'}"
                        :inactiveStyle="{color: '#606266', transform: 'scale(1)'}"
                        itemStyle="height: 40px;"
                        :current="tabsIndex"
                        @change="tabsChange"
                ></u-tabs>
            </u-sticky>
            <view class="king-bg-white king-px-5 king-mt-5">
                <u-sticky>
                    <view class="king-p-5 king-bg-white" style="height: 35px;">
                        <u-search search-icon="search" v-model="keyword" placeholder="请输入商品名称" @custom="searchInputClick" @search="searchInputClick">
                        </u-search>
                    </view>
                </u-sticky>
            </view>
            <!-- 近期购买 -->
            <view v-if="tabsIndex === 0">
                <view class="empty" v-if="payGoodslist.length === 0">
                    <u-empty mode="history" icon="http://cdn.uviewui.com/uview/empty/history.png">
                        <view>
                            <u-button type="primary" :customStyle="btnStyle" text="去逛逛" @click="toHome"></u-button>
                        </view>
                    </u-empty>
                </view>
                <view v-else>
                    <scroll-view v-show="payGoodslist.length > 0" class="foods-wrapper" scroll-y
                                 :style="'height:' + windows_height + 'px'"
                                 scroll-with-animation="true" refresher-enabled="true" :refresher-threshold="70"
                                 :refresher-triggered="triggered" @refresherrefresh="onRefresh"
                                 @scrolltolower="scrollTolower"
                                 :scroll-anchoring="true">
                        <!-- 商品列表 -->
                        <GoodsList :vertical="true" :lists="payGoodslist" price-type="￥" :is-audit="isAudit" :is-add-cart="true" imgWidth="90px" imgHeight="90px" :show-pay-count="true" @updateCart="updatePayGoodsCart"></GoodsList>
                        <view class="king-py-20" @click="scrollTolower">
                            <u-loadmore :status="payLoadMore" loading-text="努力加载中，请喝杯茶" loadmore-text="上拉加载更多"
                                        nomore-text="实在是没有了"/>
                        </view>
                    </scroll-view>
                </view>
            </view>
            <!-- 收藏 -->
            <view v-if="tabsIndex === 1">
                <view class="empty" v-if="favoritesList.length === 0">
                    <u-empty mode="list" icon="http://cdn.uviewui.com/uview/empty/data.png">
                        <view>
                            <u-button type="primary" :customStyle="btnStyle" text="去逛逛" @click="toHome"></u-button>
                        </view>
                    </u-empty>
                </view>
                <view v-else>
                    <scroll-view v-show="favoritesList.length > 0" class="foods-wrapper" scroll-y
                                 :style="'height:' + windows_height + 'px'"
                                 scroll-with-animation="true" refresher-enabled="true" :refresher-threshold="70"
                                 :refresher-triggered="triggered" @refresherrefresh="onRefresh"
                                 @scrolltolower="scrollTolower"
                                 :scroll-anchoring="true">
                        <!-- 商品列表 -->
                        <GoodsList :vertical="true" :lists="favoritesList" price-type="￥" :is-audit="isAudit" :is-add-cart="true" imgWidth="90px" imgHeight="90px" @onGoodsLongClick="onGoodsLongClick" @updateCart="updateGoodsCart"></GoodsList>
                        <view class="king-py-20" @click="scrollTolower">
                            <u-loadmore :status="favLoadMore" loading-text="努力加载中，请喝杯茶" loadmore-text="上拉加载更多"
                                        nomore-text="实在是没有了"/>
                        </view>
                    </scroll-view>
                </view>
            </view>
        </view>
        <!-- 删除弹窗 -->
        <u-modal :show="showDeleteFav" :showCancelButton="true" :closeOnClickOverlay="true" @confirm="favoritesClick"
                 @cancel="hideDeleteFavDalog" @close="hideDeleteFavDalog">
            <text class="king-font-weight-100 king-font22">
                确定取消收藏吗？
            </text>
        </u-modal>
        <!-- 登录 -->
        <loginPop :show="showLoginDialog" @close="hideLogin" @success="loginSuccess"/>
        <Tabbar :tabsId="2"/>
        <u-toast ref="toast" style="z-index: 9999"></u-toast>
    </pageWrapper>
</template>


<script>
import Tabbar from '@/components/tabbar/tabbar.vue'
import loginPop from '@/components/loginPop/loginPop.vue'
import {getRecentlyPurchasedGoodsList, getRecentlyPurchasedGoodsListLoading} from '@/api/order';
import GoodsList from '@/components/goodsList/goodsList.vue'
import {getUser, getToken, setUser} from '@/store/storage.js'
import config from '@/config/config.js'
import {addCart, deleteCartByIds} from "@/api/cart";
import {getFavoritesListPage, getFavoritesListPageLoding, favorites} from "@/api/favorites";
import { getUserAuditStatus } from '@/api/user';
export default {
    name: "quickPay",
    components: {
        Tabbar,
        loginPop,
        GoodsList
    },
    data(){
        return {
            isAudit: false,
            token: '',
            tabsIndex: 0,
            keyword: '',
            windows_height: 0, //屏幕高度
            payGoodslist: [], // 近期购买商品
            favoritesList: [], // 收藏商品
            triggered: false, // 下拉刷新状态
            payLoadMore: 'loadmore', // 上拉加载状态
            favLoadMore: 'loadmore', // 上拉加载状态
            showLoginDialog: false, // 登录
            showDeleteFav: false, // 取消收藏模态框
            currentDeleteFavIndex: 0, // 当前取消收藏的ID
            payPage: {
                page: 1,
                pageSize:8,
                total: 0,
                isMore: true
            },
            favPage: {
                page: 1,
                pageSize:8,
                total: 0,
                isMore: true
            },
            tabsList: [{name: '近期购买'}, {name: '我的收藏'}],
            btnStyle: {
                width: '120px',
                marginTop: '20px',
                borderRadius: '20px',
            },
            toLoginStyle: {
                width: '120px',
                marginTop: '20px',
                borderRadius: '20px',
            },
        }
    },
    onLoad() {
        this.token = getToken()
        if (this.token) {
            let user = getUser()
            if (user) {
                if ( user.auditStatus === 1) {
                    this.isAudit = true
                }else {
                    getUserAuditStatus().then(res => {
                        if (res.data.auditStatus === 1) {
                            this.isAudit = true
                            user.auditStatus = 1
                            setUser(user)
                        }
                    })
                }
            }
        }

    },
    onReady() {
        const windowHeight = uni.getSystemInfoSync().windowHeight
        // 50 是 tabbar 的高度
        this.hh = windowHeight - 50
        this.navCount = Math.round(this.hh / 50)

        this.windows_height = Number(uni.getSystemInfoSync().windowHeight) - 45 - 89;
    },
    mounted() {},
    onShow() {
        // 登录的情况下获取
        if (this.token) {
            this.getRecentlyGoodsList(0)
        }
    },
    methods: {
        async getRecentlyGoodsList(type) {
            const data = {}
            if (this.keyword) {
                data.name = this.keyword
            }
            if (type === 0) {
                this.payPage.page = 1
                this.payPage.isMore = true
                this.payLoadMore = 'loadmore'
            }
            data.page = this.payPage.page
            data.pageSize = this.payPage.pageSize
            let res
            if (type === 0) {
                res = await getRecentlyPurchasedGoodsListLoading(data, this.$refs.toast)
            } else {
                res = await getRecentlyPurchasedGoodsList(data, this.$refs.toast)
            }
            if (res.code !== 0 && type === 0) {
                return this.$message(this.$refs.toast).error(res.msg)
            }else if (res.code !== 0) {
                return false
            }
            this.payPage.page++
            if (res.data.list == null || res.data.list.length === 0) {
                this.payPage.isMore = false
                this.payLoadMore = 'nomore'
            }
            // 进行赋值并计算是否还有下一页
            this.payPage.total = res.data.total
            // 如果没有更多数据，则将isMore设置为false
            if ((this.payPage.page - 1) * this.payPage.pageSize >= this.payPage.total || res.data.list == null || res.data.list.length === 0) {
                console.log("没有更多了");
                this.payPage.isMore = false
                this.payLoadMore = 'nomore'
            }
            res.data.list && res.data.list.forEach(item => {
                if (item.images !== null && item.images[0] && item.images[0].url.slice(0, 4) !== 'http') {
                    item.images[0].url = config.baseUrl + "/" + item.images[0].url
                }
                if (item.weight > 1000) {
                    item.weight = item.weight / 1000 + 'kg'
                } else {
                    item.weight = item.weight + 'g'
                }
                if (item.cartNum === null || item.cartNum === undefined) {
                    item.cartNum = 0
                }
            })
            if (type === 1) {
                this.payGoodslist = [...this.payGoodslist, ...res.data.list]
            } else {
                this.payGoodslist = res.data.list
            }
            return true
        },
        // 收藏商品列表
        async getFavoritesGoodsList(type) {
            const data = {}
            if (this.keyword) {
                data.name = this.keyword
            }
            if (type === 0) {
                this.favPage.page = 1
                this.favPage.isMore = true
                this.favLoadMore = 'loadmore'
            }
            data.page = this.favPage.page
            data.pageSize = this.favPage.pageSize
            let res
            if (type === 0) {
                res = await getFavoritesListPageLoding(data, this.$refs.toast)
            } else {
                res = await getFavoritesListPage(data, this.$refs.toast)
            }
            if (res.code !== 0 && type === 0) {
                return this.$message(this.$refs.toast).error(res.msg)
            }else if (res.code !== 0) {
                return false
            }
            this.favPage.page++
            if (res.data.list == null || res.data.list.length === 0) {
                this.favPage.isMore = false
                this.favLoadMore = 'nomore'
            }

            // 进行赋值并计算是否还有下一页
            this.favPage.total = res.data.total
            // 如果没有更多数据，则将isMore设置为false
            if ((this.favPage.page - 1) * this.favPage.pageSize >= this.favPage.total || res.data.list == null || res.data.list.length === 0) {
                console.log("没有更多了");
                this.favPage.isMore = false
                this.favLoadMore = 'nomore'
            }

            res.data.list && res.data.list.forEach(item => {
                if (item.images !== null && item.images[0] && item.images[0].url.slice(0, 4) !== 'http') {
                    item.images[0].url = config.baseUrl + "/" + item.images[0].url
                }
                if (item.weight > 1000) {
                    item.weight = item.weight / 1000 + 'kg'
                } else {
                    item.weight = item.weight + 'g'
                }
                if (item.cartNum === null || item.cartNum === undefined) {
                    item.cartNum = 0
                }
            })
            if (type === 1) {
                this.favoritesList = [...this.favoritesList, ...res.data.list]
            } else {
                this.favoritesList = res.data.list
            }
            return true
        },
        async updateGoodsCart(index, cardId, num){
            if (this.tabsIndex === 0) {
               await this.updatePayGoodsCart(index, cardId, num)
            } else {
                await this.updateFavGoodsCart(index, cardId, num)
            }
        },
        async updatePayGoodsCart(index, cardId, num){
            num = parseInt(num)
            console.log("updatePayGoodsCart", index, cardId, num)
            let originNum = this.payGoodslist[index].cartNum
            if (originNum > 0 && num < this.payGoodslist[index].minCount) {
                this.$message(this.$refs.toast).error("商品最低购买数量为 " + this.payGoodslist[index].minCount)
                num = 0
            }
            if ((originNum === null || originNum === undefined || originNum === 0) && num < this.payGoodslist[index].minCount) {
                num = this.payGoodslist[index].minCount
            }
            if (originNum !== 0 && num === 0) {
                let data = {
                    ids: [cardId]
                }
                const res = await deleteCartByIds(data, this.$refs.toast)
                if (res.code !== 0) {
                    return
                }
                this.payGoodslist[index].cartNum = 0
                return
            }

            this.payGoodslist[index].cartNum = num
            let success = await this.addCartReq(this.payGoodslist[index], index, num)
            if (!success) {
                this.payGoodslist[index].cartNum = originNum
            }
        },
        async updateFavGoodsCart(index, cardId, num){
            num = parseInt(num)
            console.log("updatePayGoodsCart", index, cardId, num)
            let originNum = this.favoritesList[index].cartNum
            if (originNum > 0 && num < this.favoritesList[index].minCount) {
                this.$message(this.$refs.toast).error("商品最低购买数量为 " + this.favoritesList[index].minCount)
                num = 0
            }
            if ((originNum === null || originNum === undefined || originNum === 0) && num < this.favoritesList[index].minCount) {
                num = this.favoritesList[index].minCount
            }
            if (originNum !== 0 && num === 0) {
                let data = {
                    ids: [cardId]
                }
                const res = await deleteCartByIds(data, this.$refs.toast)
                if (res.code !== 0) {
                    return
                }
                this.favoritesList[index].cartNum = 0
                return
            }

            this.favoritesList[index].cartNum = num
            let success = await this.addCartReq(this.favoritesList[index], index, num)
            if (!success) {
                this.favoritesList[index].cartNum = originNum
            }
        },
        hideDeleteFavDalog() {
            this.showDeleteFav = false
        },
        // 显示删除提示框
        showDeleteFavDalog(id) {
            this.currentDeleteFavIndex = id
            this.showDeleteFav = true
        },
        onGoodsLongClick(goods) {
            console.log("onGoodsLongClick", goods)
            this.showDeleteFavDalog(goods.ID)
        },
        // 收藏商品
        async favoritesClick() {
            const res = await favorites({
                goodsId: this.currentDeleteFavIndex
            })
            if (res.code !== 0) {
                this.$message(this.$refs.toast).error('取消收藏失败')
                this.hideDeleteFavDalog()
                return false
            }
            this.$message(this.$refs.toast).success('取消收藏成功')
            this.getFavoritesGoodsList(0)
            this.hideDeleteFavDalog()
        },
        searchInputClick(e) {
            console.log("searchInpuSearch",e)
            if (this.tabsIndex === 0) {
                this.getRecentlyGoodsList(0)
            } else {
                this.getFavoritesGoodsList(0)
            }
        },
        // 登录成功
        loginSuccess(u) {
            this.hideLogin();
            this.token = getToken()
            this.getRecentlyGoodsList(0)
            this.$message(this.$refs.toast).success("登录成功")
        },
        // 显示登录框
        showLogin() {
            this.showLoginDialog = true
        },
        // 隐藏登录框
        hideLogin() {
            this.showLoginDialog = false
        },
        tabsChange(e) {
            this.tabsIndex = e.index
            if (this.tabsIndex === 0) {
                this.getRecentlyGoodsList(0)
            } else {
                this.getFavoritesGoodsList(0)
            }
        },
        toHome() {
            uni.navigateTo({
                url: '/pages/index/index'
            })
        },
        // 下拉刷新
        async onRefresh() {
            this.triggered = true;
            if (this.tabsIndex === 0) {
                const b = await this.getRecentlyGoodsList(0)
                if (b) {
                    this.$message(this.$refs.toast).success("刷新成功")
                } else {
                    this.$message(this.$refs.toast).success("刷新失败")
                }
            } else {
                const b = await this.getFavoritesGoodsList(0)
                if (b) {
                    this.$message(this.$refs.toast).success("刷新成功")
                } else {
                    this.$message(this.$refs.toast).success("刷新失败")
                }
            }

            this.triggered = false;
        },
        async scrollTolower(e) {
            if (this.tabsIndex === 0) {
                // 如果是在加载中就不执行或没有更多时
                if (this.payLoadMore === 'loading' || !this.payPage.isMore) {
                    return
                }
                this.payLoadMore = 'loading'
                await this.getRecentlyGoodsList(1)
                // 如果还有更多
                if (this.payPage.isMore) {
                    this.payLoadMore = 'loadmore'
                } else {
                    this.payLoadMore = 'nomore'
                }
            } else {
                // 如果是在加载中就不执行或没有更多时
                if (this.favLoadMore === 'loading' || !this.favPage.isMore) {
                    return
                }
                this.favLoadMore = 'loading'
                await this.getFavoritesGoodsList(1)
                // 如果还有更多
                if (this.favPage.isMore) {
                    this.favLoadMore = 'loadmore'
                } else {
                    this.favLoadMore = 'nomore'
                }
            }
        },
        // type 1增 2减 暂时没用
        // num 数量
        async addCartReq(goodsInfo, index, num) {
            const data = {
                goodsId: goodsInfo.ID,
                specType: 0, // 单规格
                num: num
            }
            if (num < goodsInfo.minCount) {
                this.$message(this.$refs.toast).error("商品数量不能小于 " + goodsInfo.minCount)
                data.num = 0
                return false
            }
            if (num < 1) {
                this.$message(this.$refs.toast).error("商品数量不能小于 1")
                return false
            }
            const res = await addCart(data)
            if (res.code !== 0) {
                console.log("更新购物车失败")
                return false
            }
            return true
        },
    },

}
</script>

<style lang="scss">
page {
    background: #FFFFFF;
}

.u-tabs__wrapper__nav__item {
  flex: 1 !important;
  padding: 0 !important;
}


.empty {
    background: white;
    height: 70vh;
    display: flex;
    justify-content: center;
    align-items: center;
}
</style>
