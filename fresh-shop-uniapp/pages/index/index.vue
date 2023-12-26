<!--
 * @Author: likfees
 * @Date: 2023-03-23 15:52:23
 * @LastEditors: likfees
 * @LastEditTime: 2023-04-23 14:22:39
-->
<template>
	<pageWrapper>
		<u-sticky class="king-bg-white">
			<view class="king-p-5 king-bg-white" style="height: 35px;">
<!--				<u-search search-icon="scan" disabled :show-action="false" placeholder="请输入商品名称" @click="searchClick">-->
				<u-search search-icon="search" disabled :show-action="false" placeholder="请输入商品名称" @click="searchClick">
				</u-search>
			</view>
		</u-sticky>
		<!-- 轮播图 -->
		<view class="king-p-5">
			<u-swiper :list="banner" keyName="imgUrl" indicator indicatorMode="line" :height="150" circular
				bgColor="#ffffff"></u-swiper>
		</view>
		<!-- 首页分类 -->
		<view class="king-bg-white king-mx-10 king-radius10 king-pb-5 king-pt-10" style="min-height: 80px">
			<u-grid :border="false" col="4">
				<u-grid-item v-for="c in category" :key="c.ID" @click="toGoodsByCategory(c.ID)">
					<u--image width="45" height="45" :src="c.imgUrl" shape="circle"></u--image>
					<text class="grid-text king-my-5">{{ c.title }}</text>
				</u-grid-item>
			</u-grid>
		</view>
		<view class="king-mx-10 king-my-5 king-radius10">
			<!-- 导航栏目 -->
			<u-sticky offset-top="50">
				<view class="king-bg-white king-radius10" style="height: 40px;">
					<u-row customStyle="height: 28px">
						<u-col span="6">
							<view class="goods-tabs" @click="changeGoodsTabs(0)">
								<text>热销商品</text>
								<u-transition :show="goodsTabsId == 0" mode="fade-right" duration="200">
									<view style="margin: 0 auto;">
										<view class="goods-tabs-active"></view>
									</view>
								</u-transition>
							</view>
						</u-col>
						<u-col span="6">
							<view class="goods-tabs" @click="changeGoodsTabs(1)">
								<text>新品上市</text>
								<u-transition :show="goodsTabsId == 1" mode="fade" duration="200">
									<view style="margin: 0 auto;">
										<text class="goods-tabs-active"></text>
									</view>
								</u-transition>
							</view>
						</u-col>
					</u-row>
				</view>
			</u-sticky>
		</view>
		<!-- 商品列表 -->
		<view class="king-bg-white king-mx-10 king-mb-10 king-radius10">
			<!-- 列表 -->
			<view>
				<swiper :style="{ height: swiperHeight + 'px' }" :current="goodsTabsId" @change="onChangeGoodsTabs">
					<!-- 热销商品  -->
					<swiper-item>
						<scroll-view :scroll-top="hotScrollTop" scroll-y="true" @scroll="hotScrollTopHandle"
							:style="{ height: swiperHeight + 'px' }" refresher-enabled="true" :refresher-threshold="70"
							:refresher-triggered="hotTriggered" @refresherrefresh="onRefresh"
							@scrolltolower="hotScrollTolower" :scroll-anchoring="true">
							<!-- 商品列表 -->
							<GoodsList :lists="goodsHotArr" price-type="￥" @onGoods="toGoodsInfo"></GoodsList>
							<view class="king-py-40" @click="hotScrollTolower">
								<u-loadmore :status="hotLoadMore" loading-text="努力加载中，请喝杯茶" loadmore-text="上拉加载更多"
									nomore-text="实在是没有了" />
							</view>
						</scroll-view>
					</swiper-item>
					<!-- 新品上市  -->
					<swiper-item>
						<scroll-view :scroll-top="newScrollTop" scroll-y="true" @scroll="newScrollTopHandle"
							:style="{ height: swiperHeight + 'px' }" refresher-enabled="true" :refresher-threshold="70"
							:refresher-triggered="newTriggered" @refresherrefresh="onRefresh"
							@scrolltolower="newScrollTolower" :scroll-anchoring="true">
							<!-- 商品列表 -->
							<GoodsList :lists="goodsNewArr" price-type="￥"></GoodsList>
							<view class="king-py-40" @click="newScrollTolower">
								<u-loadmore :status="newLoadMore" loading-text="努力加载中，请喝杯茶" loadmore-text="上拉加载更多"
									nomore-text="实在是没有了" />
							</view>
						</scroll-view>
					</swiper-item>
				</swiper>
			</view>
		</view>
		<loginSuspend :show="loginSuspendShow" @success="loginSuccess"></loginSuspend>
		<Tabbar :tabsId="0" />
        <u-toast style="z-index:9998;" ref="toast"></u-toast>
	</pageWrapper>
</template>

<script>
import Tabbar from '@/components/tabbar/tabbar.vue'
import loginSuspend from '@/components/loginPop/loginSuspend.vue'
import GoodsList from '@/components/goodsList/goodsList.vue'
import config from '@/config/config.js'
import { getBannerList } from '@/api/banner.js'
import { getHomeCategoryList } from '@/api/category.js'
import { getGoodsPageList } from '@/api/goods.js'
import { getToken } from '@/store/storage.js'
export default {
	components: {
		Tabbar,
		GoodsList,
		loginSuspend
	},
	data() {
		return {
			loginSuspendShow: false, // 是否显示底部登录
			goodsTabsId: 0, // 商品标签切换
			swiperHeight: 1000, // 商品栏目整体高度 页面大小
			hotScrollTop: 0,
			newScrollTop: 0,
			hotTriggered: false, // 下拉刷新状态
			newTriggered: false, // 下拉刷新状态
			hotLoadMore: 'loadmore', // 上拉加载状态
			newLoadMore: 'loadmore', // 上拉加载状态
			hotPage: {
				page: 1,
				pageSize: 12,
				total: 0, // 总条数
				isMore: true // 是否还有更多
			},
			newPage: {
				page: 1,
				pageSize: 12,
				total: 0, // 总条数
				isMore: true // 是否还有更多
			},
			banner: [],
			category: [],
			goodsHotArr: [],
			goodsNewArr: [],
		}
	},
	onLoad() {
		this._freshing = false;
		this.getBanner();
		this.getHomeCategory();
		this.getGoodsListData(0)
		this.getGoodsListData(1)
		// 如果为登录状态
		const t = getToken()
		if (!t) {
			this.loginSuspendShow = true
		}
	},
	mounted() {
		// 设置商品列表高度为页面高度
		uni.getSystemInfo({
			success: (res) => {
				const windowHeight = res.windowHeight;
				this.swiperHeight = windowHeight - 155;
			},
		});
	},
	methods: {
		//分享好友
		onShareAppMessage() {
		      return {
		        title: '启运冻品',  // 分享标题
		        path: '/pages/index/index',  // 分享路径，注意要写正确的页面路径
		        imageUrl: '/static/qiyun_logo.png',// 分享图片的本地路径
		      }
		},
		//分享到朋友圈
		onShareTimeline() {
			return {
				title: '启运冻品',
				link: '/pages/index/index',
				imageUrl: '/static/qiyun_logo.png',
			}
		},
		// 搜索框点击跳转到搜索页面
		searchClick() {
			console.log('跳转')
			uni.navigateTo({
				url: '/pages/search/search'
			})
		},
		// 获取轮播图
		getBanner() {
			getBannerList().then(res => {
				res.data.list.forEach(item => {
					if (item.imgUrl.slice(0, 4) !== 'http') {
						item.imgUrl = config.baseUrl + "/" + item.imgUrl
					}
				})
				this.banner = res.data.list;
			})
		},
		// 获取首页分类
		getHomeCategory() {
			getHomeCategoryList().then(res => {
				res.data.list.forEach(item => {
					if (item.imgUrl.slice(0, 4) !== 'http') {
						item.imgUrl = config.baseUrl + "/" + item.imgUrl
					}
				})
				this.category = res.data.list
				console.log('category', this.category);
			})
		},
		// 切换标签页
		onChangeGoodsTabs(e) {
			this.goodsTabsId = e.detail.current
		},
		// 点击切换标签页
		changeGoodsTabs(id) {
			this.goodsTabsId = id
		},
		// 获取商品列表
		// type = 1加载 其他为刷新
		async getGoodsListData(tabId, type) {
			const data = {
				goodsArea: 0
			}
			if (type == 0) {
				this.hotPage.page = 1
				this.newPage.page = 1
				this.hotPage.isMore = true
				this.newPage.isMore = true
				this.hotLoadMore = 'loadmore'
				this.newLoadMore = 'loadmore'
			}
			if (tabId == 0) {
				data.isHot = 1
				data.page = this.hotPage.page
				data.pageSize = this.hotPage.pageSize
				this.hotPage.page++
			} else if (tabId == 1) {
				data.isNew = 1
				data.page = this.newPage.page
				data.pageSize = this.newPage.pageSize
				this.newPage.page++
			} else {
				return false
			}
			const res = await getGoodsPageList(data)
			if (res.code !== 0) {
				return false
			}
			res.data.list.forEach(item => {
				if (item.images[0] && item.images[0].url.slice(0, 4) !== 'http') {
					item.images[0].url = config.baseUrl + "/" + item.images[0].url
				}
			})
			console.log(res);
			// 进行赋值并计算是否还有下一页
			if (tabId == 0) { // 热门商品
				this.hotPage.total = res.data.total
				// 如果没有更多数据，则将isMore设置为false
				if ((this.hotPage.page - 1) * this.hotPage.pageSize >= this.hotPage.total) {
					console.log("没有更多了");
					this.hotPage.isMore = false
					this.hotLoadMore = 'nomore'
				}
				if (type == 1) {
					this.goodsHotArr = [...this.goodsHotArr, ...res.data.list]
				} else {
					this.goodsHotArr = res.data.list
				}
			} else if (tabId == 1) { // 新品上市
				this.newPage.total = res.data.total
				// 如果没有更多数据，则将isMore设置为false
				console.log('(this.newPage.page - 1) * this.newPage.pageSize >= this.newPage.total', (this.newPage.page - 1) * this.newPage.pageSize, this.newPage.total);
				if ((this.newPage.page - 1) * this.newPage.pageSize >= this.newPage.total) {
					this.newPage.isMore = false
					this.newLoadMore = 'nomore'
				}
				if (type == 1) {
					this.goodsNewArr = [...this.goodsNewArr, ...res.data.list]
				} else {
					this.goodsNewArr = res.data.list
				}
			} else {
				return false
			}
			return true
		},
		// 热门列表滑动距离
		hotScrollTopHandle(e) {
			// 会出现抖动
			//this.hotScrollTop = e.detail.scrollTop.toFixed(0)
		},
		// 上新列表滑动距离
		newScrollTopHandle(e) {
			//this.newScrollTop = e.detail.scrollTop.toFixed(0)
		},
		// 返回列表的顶部
		toTop() {
			this.hotScrollTop = 0
			this.newScrollTop = 0
		},
		// 下拉刷新
		// type = 1热门 2 上新
		async onRefresh() {
			if (this.goodsTabsId == 0) {
				this.hotTriggered = true;
			} else if (this.goodsTabsId == 1) {
				this.newTriggered = true;
			} else {
				return
			}
			const b = await this.getGoodsListData(this.goodsTabsId, 0)
			if (b) {
                this.$message(this.$refs.toast).success("刷新成功")
			} else {
                this.$message(this.$refs.toast).success("刷新失败")
			}
			this.hotTriggered = false;
			this.newTriggered = false;
		},
		async hotScrollTolower(e) {
			// 如果是在加载中就不执行或没有更多时
			if (this.hotLoadMore == 'loading' || !this.hotPage.isMore) {
				return
			}
			// 设置状态为加载中
			this.hotLoadMore = 'loading'
			await this.getGoodsListData(this.goodsTabsId, 1)
			// 如果还有更多
			if (this.hotPage.isMore) {
				this.hotLoadMore = 'loadmore'
			} else {
				this.hotLoadMore = 'nomore'
			}
		},
		async newScrollTolower(e) {
			// 如果是在加载中就不执行或没有更多时
			if (this.newLoadMore == 'loading' || !this.newPage.isMore) {
				return
			}
			// 设置状态为加载中
			this.newLoadMore = 'loading'
			await this.getGoodsListData(this.goodsTabsId, 1)
			// 如果还有更多
			if (this.newPage.isMore) {
				this.newLoadMore = 'loadmore'
			} else {
				this.newLoadMore = 'nomore'
			}
		},
		// 跳转商品列表
		toGoodsByCategory(categoryId) {
			uni.navigateTo({
				url: `/pages/goods/goods?categoryId=` + categoryId
			})
		},
		// 跳转商品详情
		toGoodsInfo(goods) {
			uni.navigateTo({
				url: `/pages/goods/goodsInfo?id=` + goods.ID
			})
		},
		// 登陆成功
		loginSuccess() {
			this.loginSuspendShow = false
		}
	},
}
</script>

<style lang="scss">
.goods-tabs {
	margin: 0 auto;
	text-align: center;
	height: 16px;
	background: #ffffff;

	.goods-tabs-active {
		width: 66px;
		border-bottom: 4px solid #2979ff;
		padding-bottom: 4px;
		display: block;
	}
}
</style>
