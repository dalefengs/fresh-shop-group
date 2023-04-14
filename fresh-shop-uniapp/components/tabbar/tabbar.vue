<!--
 * @Author: likfees
 * @Date: 2023-03-23 16:56:40
 * @LastEditors: likfees
 * @LastEditTime: 2023-04-13 21:32:22
-->
<template>
	<view>
		<u-tabbar :value="tabsId" @change="tabsChange" :fixed="true" :placeholder="true" :safeAreaInsetBottom="true">
			<u-tabbar-item v-for="(item, key) in tabs" :text="item.text" :icon="item.icon" :key="key" />
		</u-tabbar>
	</view>
</template>

<script>
export default {
	name: "tabbar",
	props: {
		tabsId: {
			type: Number,
			default: 0
		}
	},
	data() {
		return {
			tabs: [
				{
					text: "首页",
					icon: "home",
					url: "/pages/index/index"
				},
				{
					text: "分类",
					icon: "grid",
					url: "/pages/category/category"
				},
				{
					text: "购物车",
					icon: "shopping-cart",
					url: "/pages/cart/cart"
				},
				{
					text: "我的",
					icon: "account",
					url: "/pages/my/my"
				},
			],

		};
	},
	methods: {
		tabsChange(index) {
			uni.redirectTo({
				url: this.tabs[index].url
			})
		}
	}
}
</script>

<style lang="scss"></style>