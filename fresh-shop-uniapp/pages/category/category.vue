<template>
	<pageWrapper>
		<u-sticky class="king-bg-white">
			<view class="king-p-5 king-bg-white" style="height: 35px;">
				<u-search search-icon="scan" disabled :show-action="false" placeholder="日照香炉生紫烟" @click="searchClick">
				</u-search>
			</view>
		</u-sticky>
		<view class="content king-mt-5">
			<scroll-view class="menu-wrapper" scroll-y :scroll-top="left_scroll" scroll-with-animation="true">
				<view class="left-content">
					<view style="position: relative;" v-for="(item, index) in categoryList" :key="index" ref="menuList"
						@click="select(item.ID)" :class="{ 'current': currentIndex == item.ID }">
						<view class="menu-item">{{ item.title }}</view>
						<!--<text class="allcount" v-if="getAllCount >= item.count && item.count > 0">{{ item.count }}</text>-->
					</view>
				</view>
			</scroll-view>
			<!-- 右侧滚动 -->
			<view style="width: 78%;">
				<view v-if="brandList.length > 0" class="king-pt-10 king-px-20 king-bg-white">
					<u-scroll-list :indicator="brandList.length > 4" indicatorColor="#fff0f0"
						indicatorActiveColor="#00A0DC">
						<view class="scroll-list" style="flex-direction: row;">
							<view class="scroll-list__goods-item" v-for="(item, index) in brandList" :key="index">
								<image class="scroll-list__goods-item__image" :src="item.logo" mode=""></image>
								<br>
								<text class="scroll-list__goods-item__text">{{ item.name }}</text>
							</view>
						</view>
					</u-scroll-list>
				</view>
				<scroll-view v-if="goodsArr.length > 0" class="foods-wrapper" scroll-y :style="'height:' + windows_height + 'px'"
					:scroll-top="foodSTop" @scroll="myscroll" scroll-with-animation="true">
					<!-- 商品列表 -->
					<GoodsList :vertical="true" :lists="goodsArr" price-type="￥"></GoodsList>
				</scroll-view>
				<u-empty v-else :style="{ height: scrollViewHeight / 1.6 + 'px' }" width="220" height="220" textSize="16"
					text="暂无商品" mode="data" icon="http://cdn.uviewui.com/uview/empty/data.png" />
			</view>

			<shopcart :goods="goods" @add="addCart" @dec="decreaseCart" @delAll="delAll"></shopcart>
		</view>
		<Tabbar :tabsId="1" />
	</pageWrapper>
</template>

<script>
import Tabbar from '@/components/tabbar/tabbar.vue'
import shopcart from '@/components/goodsList/shopcart.vue'
import {
	getGoodsPageListLoading
} from '@/api/goods.js'
import {
	getCategoryListAll
} from '@/api/category.js'
import {
	getBrandListByCategoryId
} from '@/api/brand.js'
import config from '@/config/config.js'
import GoodsList from '@/components/goodsList/goodsList.vue'
export default {
	components: {
		Tabbar,
		shopcart,
		GoodsList
	},
	data() {
		return {
			categoryList: [{
				ID: 0,
				title: "全部分类",
			}],
			brandList: [],
			goods: [],
			goodsArr: [],
			page: {
				page: 1,
				pageSize: 20,
				total: 0,
				isMore: true
			},
			windows_height: 0, //屏幕高度
			foodSTop: 0, //右侧的滑动值
			currentIndex: 0, // 左边栏默认ID, 分类ID
			brandId: 0, // 选择的品牌 ID
			last: 0,
			right_height: [], //右侧每个内容的高度集合
			select_index: 0,
			left_height: 0, //左侧总高度
			left_scroll: 0, //左侧滑动值
			catrgoryList: [],
			timeout: null
		}

	},
	onLoad() {
		this.init()
	},
	onReady() {
		var windowHeight = uni.getSystemInfoSync().windowHeight
		console.log(windowHeight)
		// 50 是 tabbar 的高度
		this.hh = windowHeight - 50
		this.navCount = Math.round(this.hh / 50)

		this.windows_height = Number(uni.getSystemInfoSync().windowHeight) - 55 - 110 - 50;
		setTimeout(() => {
			this.getHeightList();
		}, 1000)
	},
	computed: {
		getList() {
			let result = [];
			this.goods.forEach((good) => {

				good.foods.forEach((food) => {

					if (food.count) {
						result.push(food)
					}
				})
			})
			return result

		},
		// 获得购物车所有商品数量
		getAllCount: function (item) {
			let result = [];
			let count = 0;

			for (let i = 0; i < this.goods.length; i++) {
				count = 0;
				this.goods[i].foods.forEach((food) => {
					if (food.count >= 0) {
						count += food.count
						Vue.set(this.goods[i], 'count', count)
					}
				})
				result.push(count)
			}

			result.sort(function (a, b) {
				return a - b;
			})
			count = result[result.length - 1]
			return count;
		},

	},
	methods: {
		async getBrandListData() {
			// 获取品牌列表
			const brandRes = await getBrandListByCategoryId({
				categoryId: this.currentIndex
			})
			brandRes.data.forEach(item => {
				if (item.logo && item.logo.slice(0, 4) !== 'http') {
					item.logo = config.baseUrl + "/" + item.logo
				}
			})

			this.brandList = brandRes.data
		},
		// type = 1加载 其他为刷新
		async getGoodsListData(type) {
			const data = {}
			if (this.keyword) {
				data.name = this.keyword
			}
			if (this.currentIndex) {
				data.categoryId = this.currentIndex
			}
			if (this.brandId) {
				data.brandId = this.brandId
			}
			if (this.tagsIds) {
				data.tagsIds = this.tagsIds
			}

			if (type == 0) {
				this.page.page = 1
				this.page.isMore = true
				this.loadMore = 'loadmore'
			}
			data.page = this.page.page
			data.pageSize = this.page.pageSize
			const res = await getGoodsPageListLoading(data)
			if (res.code !== 0) {
				return false
			}
			this.page.page++

			res.data.list.forEach(item => {
				if (item.images[0] && item.images[0].url.slice(0, 4) !== 'http') {
					item.images[0].url = config.baseUrl + "/" + item.images[0].url
				}
			})
			console.log(res);
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
		async init() {
			const res = await getCategoryListAll()
			this.categoryList.push(...res.data)
			this.getBrandListData()

			this.getGoodsListData(0)
		},

		// 点击侧边栏
		select(categoryId) {
			this.currentIndex = categoryId;
			// 切换分类清空品牌筛选数据
			this.brandId = 0
			this.getBrandListData()
			this.getGoodsListData(0)
		},

		// 设置点击侧边栏右边滚动的高度
		setScrollH(index) {
			var that = this;
			let height = 0;
			var query = uni.createSelectorQuery();
			let foods = query.selectAll('.foods');

			//this.$nextTick(() => {
			//	foods.fields({
			//		size: true
			//	}, data => {
			//		if (index == 0) {
			//			that.foodSTop = 0;
			//		}
			//		for (let i = 0; i < index; i++) {

			//			height += parseInt(data[i].height);
			//			// console.log('fh', foods);
			//			that.foodSTop = height;
			//			// console.log(that.foodSTop)
			//		}

			//	}).exec();
			//})

		},
		// 搜索框点击跳转到搜索页面
		searchClick() {
			console.log('跳转')
			uni.navigateTo({
				url: '/pages/search/search'
			})
		},
		addCart(item) {
			if (item.count >= 0) {
				item.count++
				this.goods.forEach((good) => {
					good.foods.forEach((food) => {
						// 根据名字添加购物车,实际环境根据需要更改
						if (item.name == food.name)
							food.count = item.count
					})
				})
			} else {
				console.log('add')
				this.goods.forEach((good) => {
					good.foods.forEach((food) => {
						if (item.name == food.name)
							Vue.set(food, 'count', 1)
					})
				})
			}
		},
		decreaseCart(item) {
			if (item.count) {
				item.count--
				this.goods.forEach((good) => {
					good.foods.forEach((food) => {
						if (item.name == food.name)
							food.count = item.count
					})
				})
			}
		},
		// 清空购物车
		delAll() {
			this.goods.forEach((good) => {
				good.foods.forEach((food) => {
					if (food.count) {
						food.count = 0
					}
				})
			})
		},
		getHeightList() {
			let _this = this;
			let selectorQuery = uni.createSelectorQuery().in(this);
			selectorQuery.select('.left-content').boundingClientRect(function (rects) {
				_this.left_height = rects.height;
			});
			selectorQuery.selectAll('.foods').boundingClientRect(function (rects) {
				_this.right_height = rects.map((item) => item.top);
			}).exec();
		},
		myscroll(e) {
			//引入节流	
			let right_content_height = e.detail.scrollHeight - this.windows_height;
			if (right_content_height == e.detail.scrollTop) {
				return;
			}
			let scroll_top = e.detail.scrollTop + 110; //110是banner图的高度
			//判断当前的scrollTop在哪个区间内;
			let now = +new Date();
			if (now - this.last > 100) {
				this.last = now;
				let active_index = this.right_height.findIndex((value, index, arr) => value <= scroll_top &&
					scroll_top < arr[index + 1]);
				this.currentIndex = active_index;
				if (this.left_height - this.windows_height) {
					//如果有超出部分
					let diff = this.left_height - this.windows_height
					this.left_scroll = Math.round((active_index * diff) / (this.goods.length - 1))
				}
			}
		}
	}
}
</script>

<style lang="scss">
.content {
	display: flex;
	position: absolute;
	top: 45px;
	bottom: 115px;
	width: 100%;
	overflow: hidden;
}

.current {
	position: relative;
	z-index: 0;
	background-color: #fff;
	color: #00A0DC;
	border-left: 5px solid #00A0DC;
}

.menu-wrapper {
	text-align: center;
	width: 22%;
	display: flex;
	flex-direction: column;
	background: #f3f5f7;

}

.menu-item {
	width: 100%;
	display: flex;
	align-items: center;
	justify-content: center;
	height: 50px;
}

.allcount {
	position: absolute;
	right: 6px;
	top: 8px;
	display: inline-block;
	padding: 0 4px;
	font-size: 9px;
	line-height: 16px;
	font-weight: 400;
	border-radius: 50%;
	background-color: #f01414;
	color: #ffffff;
}

.foods-wrapper {
	margin-left: 4px;
	width: 100%;
}

.food,
.food-btm,
.content {
	display: flex;
	flex-direction: row;
}

.food-title {
	padding: 2px 0;
}

.food-info {
	margin-left: 10px;
	margin-right: 16px;
	display: flex;
	flex-direction: column;
	flex: 2;
}

.food-btm {
	justify-content: space-between;
}

.food-price {
	color: #00A0DC;
	font-size: 16px;
}

.scroll-list {
	display: flex;
	flex-direction: column;

	&__goods-item {
		margin-right: 20px;

		&__image {
			width: 60px;
			height: 60px;
			border-radius: 4px;
		}

		&__text {
			color: #00A0DC;
			text-align: center;
			font-size: 12px;
			margin-top: 5px;
		}
	}

	&__show-more {
		background-color: #fff0f0;
		border-radius: 3px;
		padding: 3px 6px;
		@include flex(column);
		align-items: center;

		&__text {
			font-size: 12px;
			width: 12px;
			color: #00A0DC;
			line-height: 16px;
		}
	}
}
</style>