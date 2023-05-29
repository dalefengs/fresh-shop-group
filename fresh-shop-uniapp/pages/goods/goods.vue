<template>
	<pageWrapper>
		<view class="king-bg-white">
			<!-- 检索框 -->
			<view class="search-panel-base king-px-10">
				<view class="left">
					<image :src="searchIcon" class="search-icon"></image>
					<input class="search-input" v-model="keyword" type="text" confirm-type="search" @confirm="searchConfirm"
						placeholder="请输入商品名称" />
					<image v-if="keyword" :src="delIcon" class="search-del-icon" @click="clearKeyword"></image>
				</view>
				<view class="right">
					<text class="search-btn" @click="search" style="border: 1rpx solid; color: #2979ff">搜索</text>
				</view>
			</view>
			<filterDropdown style="z-index: 1000" ref="dropdown" :itemArr="itemArr" @finish="finish" @clear="clearWhere"></filterDropdown>
		</view>
		<view class="king-mx-10 king-my-5 king-bg-white king-radius10">
			<scroll-view v-if="goodsArr.length > 0" :scroll-top="scrollTop" scroll-y="true" @scroll="scrollTopHandle"
				:style="{ height: scrollViewHeight + 'px' }" refresher-enabled="true" :refresher-threshold="70"
				:refresher-triggered="triggered" @refresherrefresh="onRefresh" @scrolltolower="scrollTolower"
				:scroll-anchoring="true">
				<!-- 商品列表 -->
				<GoodsList :lists="goodsArr" price-type="￥"></GoodsList>
				<view class="king-py-40" @click="scrollTolower">
					<u-loadmore :status="loadMore" loading-text="努力加载中，请喝杯茶" loadmore-text="上拉加载更多" nomore-text="实在是没有了" />
				</view>
			</scroll-view>
			<u-empty v-else :style="{ height: scrollViewHeight / 1.6 + 'px' }" width="220" height="220" textSize="16"
				text="暂无商品" mode="data" icon="http://cdn.uviewui.com/uview/empty/data.png" />
				<!-- 返回顶部 -->
			<u-back-top :scroll-top="scrollTop" @click="toTop"></u-back-top>
		</view>
        <u-toast style="z-index:9998;" ref="toast"></u-toast>
	</pageWrapper>
</template>

<script>
import filterDropdown from '@/components/filterDropdown/filterDropdown.vue';
import config from '@/config/config.js'
import { getBrandListAll } from '@/api/brand.js'
import { getCategoryListAll } from '@/api/category.js'
import { getTagsListAll } from '@/api/tags.js'
import { getGoodsPageList, getGoodsPageListLoading } from '@/api/goods.js'
import GoodsList from '@/components/goodsList/goodsList.vue'
export default {
	components: {
		filterDropdown,
		GoodsList
	},
	data() {
		return {
			searchIcon: 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAFAAAABQCAYAAACOEfKtAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAyZpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDUuNi1jMTExIDc5LjE1ODMyNSwgMjAxNS8wOS8xMC0wMToxMDoyMCAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvIiB4bWxuczp4bXBNTT0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL21tLyIgeG1sbnM6c3RSZWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9zVHlwZS9SZXNvdXJjZVJlZiMiIHhtcDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIENDIDIwMTUgKFdpbmRvd3MpIiB4bXBNTTpJbnN0YW5jZUlEPSJ4bXAuaWlkOjU4MEVFRkU4RTlEMjExRUJBREJFRTIyQjJDMUE5NkQ3IiB4bXBNTTpEb2N1bWVudElEPSJ4bXAuZGlkOjU4MEVFRkU5RTlEMjExRUJBREJFRTIyQjJDMUE5NkQ3Ij4gPHhtcE1NOkRlcml2ZWRGcm9tIHN0UmVmOmluc3RhbmNlSUQ9InhtcC5paWQ6NTgwRUVGRTZFOUQyMTFFQkFEQkVFMjJCMkMxQTk2RDciIHN0UmVmOmRvY3VtZW50SUQ9InhtcC5kaWQ6NTgwRUVGRTdFOUQyMTFFQkFEQkVFMjJCMkMxQTk2RDciLz4gPC9yZGY6RGVzY3JpcHRpb24+IDwvcmRmOlJERj4gPC94OnhtcG1ldGE+IDw/eHBhY2tldCBlbmQ9InIiPz4L5MHuAAAMU0lEQVR42uxc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',
			delIcon: 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAYAAACqaXHeAAAAAXNSR0IArs4c6QAABK9JREFUeF7tm49RFDEUxl8qECoQKhAqgHcNCBUIFXhUoFQgVuBRgdLAPajAswKxArWCON9O9mYvZHfzf7nDzOwMzOwl+X55SV5e3ioqXERkT2t9opQ6IiI8e0R0YJ5u649EhOcPEa201iul1AMz4/9iRZWoWUQg9B0RnRrRKc2siOieiG6ZGX9nLdkAYKSN6LljdHN1GhZyY2BksYxkAEb4eyKCcECoUSAeID6nTpEkACICM0dHagm34QLEnJlvY6lHARARLGJfzByPbTvn77BGXDIzpkhQCQYgImdG/FSj3icQ1gAI30IIBAEQkU9mroe0UfvdG2a+8m3UG4CIwOQvfCue+L0FM1/69MELgIh8JSKY/jYVLwijALZs5O0BGoUwCEBEsMVhj9/mAl8BPoqz9AIwqz1MfxfKed/u4ARg9vnvEzo4uaFjizx2+Ql9AOQZOTm5YNwzM9uVPQEgItjqsOXtYoGjtOgK2wBgDjY/d8j07UHEVDjsHqBsAB+J6MMuDn1H0zUzQ2dT1gBSR19r/VcphZPZAgERrfWNUuokB0yt9QPqNhGjM1P3q8i6N6ygCwANwNePLU/ml4hgvuHInFIQCdpwwTOsU1fMDB9nwwIw93HMjS37ruBEIoQn4tE5Y62/YzsKS2LmwzUAE8PDvp9SsM86Y3aREJziDQDEHLP0t5kCmVxe5z7bEg2E0Cve9DeHn9K4yC0A0ATV1DJ4+PCEMCY+17F8xczHKsN8sqGlQKglvu3zPgDgnJ/70BMDobZ4QDgHgFLOTwiEKcQDwDUAIIj4NnXy9/zeBwLZ+3y3rsIBmTu1XC7vc3lsMRCGwBcWT42HKSKpDpCP8YyGpuxKSos37a0AQPsoyPCON4RK4htJNQGgvVEINcVPAWBwtTdeXo4DlLfB1rSAUfGRbrO3WNeL2AUelVKvk2oZ/7G3+JoQtNY/amyDweJrQWi3wZKOkI+HB0eo9x7P8wA1boPuN+5KusI+4ttIj4/HmBpZciFoXOESh6EQ8W3HpoDQHIaQ6JASXrLJxoifCsJ+ExBZLpfIyXsTO5E6v0sRXxUCdoDZbHaULSSGFXU2myEv0FkCPbzB6ZDpALcREssWZHSpDxQ/agk5g7jre4EMDlFfWDwlhue0hNR1qzX/5izQcTpKXIykiO+1hFIXI0hqhluccuUEiHfmggW3TL1rQuCCizxAZH4hDxDRq+jkTHOFd9Be4vy/HO2OhEltT7GCwIGt+7o9+htrQGcteLkJEi2ETPts3eEdaa3PT+lNkjJfbMQuiM9NPHIXjryTpND7QoekqcCEpcl11oOXmyjZgVA1SJnZREajUaO5wmY6lIwaZda8rm5UvHMbHDjNbZMleIkPAmAsYRvWhMHkaHuAvaaA5S0iTW2RcGYoYvLGy7so+slMZ2E8MBCy5AGmEjF5hBBf/qMpyxouEpMWk7R3kzNjKwqeAnZDJjgx11rPa00LIxzrET6QSvqCNBlAZ1ogugyLAIgiV21a619KKQhHpChJeNvvbACsqYEYI2Ccpkabm/s7pRAQgejn+/H0gP8Ay2i/IkcSdfP5vG0lGF1EfJRSzefz5kHyZZaR7uvfP1aimGHUO86VAAAAAElFTkSuQmCC',
			type: 0, // 筛选类型 0:排序 1:分类 2:品牌 3:标签
			sortType: 0, // 排序 id
			brandId: 0, // 品牌 id
			categoryId: 0, // 分类 id
			keyword: '', // 关键字
			tagsIds: '', // 标签 id
			brandList: [], // 品牌列表
			categoryList: [], // 分类列表
			itemArr: [
				[
					{ text: '智能排序', value: 0 },
					{ text: '最新', value: 1 },
					{ text: '销量', value: 2 },
				],
				[
					{ text: '全部分类', value: 0 },
				],
				[
					{ text: '全部品牌', value: 0 },
				]
			],
			goodsArr: [], // 商品列表
			scrollViewHeight: 0,
			scrollTop: 0,
			page: {
				page: 1,
				pageSize: 6,
				total: 0,
				isMore: true
			},
			triggered: false, // 下拉刷新状态
			loadMore: 'loadmore', // 上拉加载状态
			scrollTimer: false, // 定时器id
		};
	},
	onLoad(option) {
		if (option.keyword) {
			this.keyword = option.keyword
		}
		if (option.type) {
			this.type = option.type
		}
		if (option.sortType) {
			this.sortType = option.sortType
		}
		if (option.brandId) {
			this.brandId = option.brandId
		}
		if (option.categoryId) {
			this.categoryId = option.categoryId
		}
		if (option.tagsId) {
			this.tagsId = option.tagsId
		}

		this.init()
	},
	mounted() {
		// 设置商品列表高度为页面高度
		uni.getSystemInfo({
			success: (res) => {
				const windowHeight = res.windowHeight;
				this.scrollViewHeight = windowHeight - 110;
			},
		});

		setTimeout(() => {
			if (this.categoryId) {
				let title = ''
				this.categoryList.forEach(item => {
					if (item.ID == this.categoryId) {
						title = item.title
					}
				})
				if (title) {
					this.$refs.dropdown.updateTitle(1, title)
				}
			}

			if (this.brandId) {
				let title = ''
				this.brandList.forEach(item => {
					if (item.ID == this.brandId) {
						title = item.name
					}
				})
				if (title) {
					this.$refs.dropdown.updateTitle(2, title)
				}
			}
		}, 1000);


	},
	methods: {
		// 用户选择了筛选条件
		finish(item) {
			console.log(item);
			// $index筛选类型 0:排序 1:分类 2:品牌 3:标签
			this.type = item.$index
			if (item.$index == 0) { // 排序
				this.sortType = item.value
			} else if (item.$index == 1) { // 分类
				this.categoryId = item.value
			} else if (item.$index == 2) { // 品牌
				this.brandId = item.value
			} else if (item.$index == 3) { // 标签
				this.tagsIds = item.value.map(i => {
					return i.value
				})
				console.log("ids", this.tagsIds);
			}
			this.getGoodsListData(0)
		},
		async init() {
			const category = []
			const brand = []
			const tags = []
			// 获取banner
			const resCategory = await getCategoryListAll();
			this.categoryList = resCategory.data
			resCategory.data.forEach(item => {
				category.push({
					text: item.title,
					value: item.ID
				})
			})
			this.itemArr[1] = [...this.itemArr[1], ...category]
			const resBrand = await getBrandListAll();
			this.brandList = resBrand.data
			resBrand.data.forEach(item => {
				brand.push({
					text: item.name,
					value: item.ID
				})
			})
			this.itemArr[2] = [...this.itemArr[2], ...brand]
			const resTags = await getTagsListAll();
			resTags.data.forEach(item => {
				tags.push({
					text: item.name,
					value: item.ID
				})
			})
			this.itemArr[3] = tags
			await this.getGoodsListData(0)
		},

		// type = 1加载 其他为刷新
		async getGoodsListData(type) {
			const data = {}
			if (this.keyword) {
				data.name = this.keyword
			}
			if (this.categoryId) {
				data.categoryId = this.categoryId
			}
			if (this.brandId) {
				data.brandId = this.brandId
			}
			if (this.tagsIds) {
				data.tagsIds = this.tagsIds.toString()
			}
			if (this.sortType) {
				data.sortType = this.sortType
			}

			if (type == 0) {
				this.page.page = 1
				this.page.isMore = true
				this.loadMore = 'loadmore'
			}
			data.page = this.page.page
			data.pageSize = this.page.pageSize
            let res = {}
            if (type == 1) { // 加载
                res = await getGoodsPageList(data)
            }else {
                res = await getGoodsPageListLoading(data, this.$refs.toast)
            }
			if (res.code !== 0) {
				return false
			}
			this.page.page++

			res.data.list.forEach(item => {
				if (item.images[0] && item.images[0].url.slice(0, 4) !== 'http') {
					item.images[0].url = config.baseUrl + "/" + item.images[0].url
				}
			})
			// 进行赋值并计算是否还有下一页
			this.page.total = res.data.total
			// 如果没有更多数据，则将isMore设置为false
			if ((this.page.page - 1) * this.page.pageSize >= this.page.total) {
				console.log("没有更多了");
				this.page.isMore = false
				this.loadMore = 'nomore'
			}
			if (type == 1) {
				this.goodsArr = [...this.goodsArr, ...res.data.list]
			} else {
				this.goodsArr = res.data.list
			}
			return true
		},
		// 热门列表滑动距离
		scrollTopHandle(e) {
			if (this.scrollTimer) {
				clearTimeout(this.scrollTimer)
			}
			// 解决赋值小程序会发生抖动
			this.scrollTimer = setTimeout(() => {
				this.scrollTop = e.detail.scrollTop.toFixed(0)
				console.log('q', this.scrollTop);
			}, 500)

		},
		// 返回列表的顶部
		toTop() {
			this.scrollTop = 0
		},
		// 下拉刷新
		// type = 1热门 2 上新
		async onRefresh() {
			this.triggered = true;
			const b = await this.getGoodsListData(0)
            if (b) {
                this.$message(this.$refs.toast).success('刷新成功')
            } else {
                this.$message(this.$refs.toast).error('刷新失败')
            }
			this.triggered = false;
		},
		async scrollTolower(e) {
			// 如果是在加载中就不执行或没有更多时
			if (this.loadMore == 'loading' || !this.page.isMore) {
				return
			}
			// 设置状态为加载中
			this.loadMore = 'loading'
			await this.getGoodsListData(1)
			// 如果还有更多
			if (this.page.isMore) {
				this.loadMore = 'loadmore'
			} else {
				this.loadMore = 'nomore'
			}
		},
		// 清空搜索关键字
		clearKeyword() {
			this.keyword = ''
			this.getGoodsListData(0)
		},
		// 点击搜索
		search() {
			this.getGoodsListData(0)
		},
		// 清空所有条件
		clearWhere() {
			this.categoryId = 0
			this.brandId = 0
			this.tagsIds = []
			this.getGoodsListData(0)
		},
		// 回车搜索
		searchConfirm() {
			this.getGoodsListData(0)
		}
	}
}
</script>

<style lang="scss" scoped>
.search-panel-base {
	height: 100rpx;
	display: flex;
	justify-content: space-between;
	align-items: center;

	.left {
		flex: 1;
		height: 68rpx;
		background-color: #e8e8e8;
		border-radius: 30rpx;
		display: flex;
		align-items: center;
		justify-content: space-between;
		padding: 0 20rpx;
		border: 1rpx solid #f2f2f2;

		.search-icon {
			flex: none;
			width: 30rpx;
			height: 30rpx;
		}

		.search-input {
			flex: 1;
			height: 100%;
			padding: 0 20rpx;
		}

		.search-del-icon {
			width: 35rpx;
			height: 35rpx;
		}
	}

	.right {
		margin-left: 30rpx;
		height: 60rpx;
		line-height: 60rpx;

		.search-btn {
			border-radius: 30rpx;
			padding: 4px 16px;
		}
	}
}
</style>
