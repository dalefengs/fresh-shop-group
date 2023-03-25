<!--
 * @Author: likfees
 * @Date: 2023-03-23 15:52:23
 * @LastEditors: likfees
 * @LastEditTime: 2023-03-25 11:09:15
-->
<template>
	<pageWrapper>
		<u-sticky class="king-bg-white">
			<view class="king-p-10 king-bg-white" style="height: 35px;">
				<u-search search-icon="scan" disabled :show-action="false" placeholder="日照香炉生紫烟" @click="searchClick">
				</u-search>
			</view>
		</u-sticky>
		<!-- 轮播图 -->
		<view class="king-p-20">
			<u-swiper :list="banner" keyName="imgUrl" indicator indicatorMode="line" :height="180" circular
				bgColor="#ffffff"></u-swiper>
		</view>
		<!-- 首页分类 -->
		<view class="king-bg-white king-mx-20 king-radius10 king-py-20" style="min-height: 170px">
			<u-grid :border="false" col="4">
				<u-grid-item v-for="c in category" :key="c.ID" class="king-my-5">
					<u--image width="50" height="50" :src="c.imgUrl" shape="circle"></u--image>
					<text class="grid-text king-my-10">{{ c.title }}</text>
				</u-grid-item>
			</u-grid>
		</view>
		<view class="king-m-20 king-radius10">
			<!-- 导航栏目 -->
			<u-sticky offset-top="45">
				<view class="king-bg-white king-pb-10 king-radius10" style="height: 50px;">
					<u-row customStyle="margin-bottom: 10px; height: 50px">
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
		<view class="king-bg-white king-mx-20 king-mb-20 king-radius10">
			<!-- 列表 -->
			<view>
				<swiper :style="{ height: swiperHeight + 'px' }" :current="goodsTabsId" @change="onChangeGoodsTabs">
					<!-- 热销商品  -->
					<swiper-item>
						<scroll-view :scroll-top="hotScrollTop" scroll-y="true" @scroll="hotScrollTopHandle"
							:style="{ height: swiperHeight + 'px' }" refresher-enabled="true" :refresher-threshold="70"
							:refresher-triggered="hotTriggered"
							@refresherrefresh="onRefresh">
							<!-- 商品列表 -->
							<GoodsList :lists="goodsHotArr" price-type="$"></GoodsList>
							<u-back-top :scroll-top="hotScrollTop" @click="toTop"></u-back-top>
						</scroll-view>
					</swiper-item>
					<!-- 新品上市  -->
					<swiper-item>
						<scroll-view :scroll-top="newScrollTop" scroll-y="true" class="king-scroll-row"
							@scroll="newScrollTopHandle" :style="{ height: swiperHeight + 'px' }" refresher-enabled="true"
							:refresher-threshold="70" :refresher-triggered="newTriggered"
							@refresherrefresh="onRefresh">
							<!-- 商品列表 -->
							<GoodsList :lists="goodsNewArr" price-type="$"></GoodsList>
							<u-back-top :scroll-top="newScrollTop" @click="toTop"></u-back-top>
						</scroll-view>
					</swiper-item>
				</swiper>
			</view>
		</view>

		<Tabbar />
	</pageWrapper>
</template>

<script>
import Tabbar from '@/components/tabbar/tabbar.vue'
import GoodsList from '@/components/goodsList/goodsList.vue'
import config from '@/config/config.js'
import { getBannerList } from '@/api/banner.js'
import { getHomeCategoryList } from '@/api/category.js'
import { getGoodsPageList } from '@/api/goods.js'
export default {
	components: {
		Tabbar,
		GoodsList
	},
	data() {
		return {
			goodsTabsId: 0, // 商品标签切换
			swiperHeight: 0, // 商品栏目整体高度 页面大小
			goodsListHeight: 0, // 子栏栏目商品
			hotScrollTop: 0,
			newScrollTop: 0,
			hotTriggered: false, // 下拉刷新状态
			newTriggered: false, // 下拉刷新状态
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
	},
	mounted() {
		// 设置商品列表高度为页面高度
		uni.getSystemInfo({
			success: (res) => {
				const windowHeight = res.windowHeight;
				this.swiperHeight = windowHeight - 180;
			},
		});
	},
	methods: {
		// 搜索框点击跳转到搜索页面
		searchClick() {
			console.log('跳转')
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
		// type = 0热门 1 上新
		async getGoodsListData(type) {
			const data = {}
			if(type == 0) {
				data.isHot = 1
			} else if(type == 1) {
				data.isNew  = 1
			} else {
				return false
			}
			const res = await getGoodsPageList(data)
			if (res.code !== 0) {
				return false
			}
			res.data.list.forEach(item => {
				if (item.images[0].url.slice(0, 4) !== 'http') {
					item.images[0].url = config.baseUrl + "/" + item.images[0].url
				}
			})
			console.log(res);
			if (type == 0) {
				this.goodsHotArr = res.data.list
			} else if (type == 1) {
				this.goodsNewArr = res.data.list
			} else {
				return false
			}
			return true
		},
		// 热门列表滑动距离
		hotScrollTopHandle(e) {
			this.hotScrollTop = e.detail.scrollTop

		},
		// 上新列表滑动距离
		newScrollTopHandle(e) {
			this.newScrollTop = e.detail.scrollTop

		},
		// 返回列表的顶部 
		toTop() {
			this.hotScrollTop = 0
			this.newScrollTop = 0
		},
		// 下拉刷新
		// type = 1热门 2 上新
		async onRefresh() {
			if(this.goodsTabsId == 0) {
				this.hotTriggered = true;
			}else if(this.goodsTabsId == 1) {
				this.newTriggered = true;
 			} else {
				return
			}
			const b = await this.getGoodsListData(this.goodsTabsId)
			if (b) {
				this.$message.success("刷新成功")
			} else {
				this.$message.success("刷新失败")
			}
			this.hotTriggered = false;
			this.newTriggered = false;
		}
	},
}
</script>

<style lang="scss">
.goods-tabs {
	margin: 0 auto;
	text-align: center;
	height: 20px;
	background: #ffffff;

	.goods-tabs-active {
		width: 66px;
		border-bottom: 4px solid #2979ff;
		padding-bottom: 4px;
		display: block;
	}
}
</style>
