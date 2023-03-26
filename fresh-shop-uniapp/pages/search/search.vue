<!--
 * @Author: likfees
 * @Date: 2023-03-25 13:49:00
 * @LastEditors: likfees
 * @LastEditTime: 2023-03-26 17:10:09
-->
<template>
	<pageWrapper class="king-bg-white">
		<searchPage @search="searchClick" />
	</pageWrapper>
</template>

<script>
import searchPage from '@/components/searchPage/searchPage.vue'
export default {
	components: {
		searchPage
	},
	data() {
		return {
		}
	},
	methods: {
		// 点击了搜索
		searchClick(keyword) {
			console.log(keyword);
			// 跳转到商品列表页面
			uni.redirectTo({
				url: '/pages/goods/goods?keyword=' + keyword,
				query: {
					keyword: keyword
				}
			})
		},
	}
}
</script>

<style lang="scss">
body {
	background: #ffffff;
}
</style>
