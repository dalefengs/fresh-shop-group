<!--
 * @Author: dalefeng
 * @Date: 2023-03-23 16:56:40
 * @LastEditors: dalefeng
 * @LastEditTime: 2023-04-14 20:56:15
-->
<template>
	<view>
		<u-tabbar :value="tabsId" @change="tabsChange" :fixed="true" :placeholder="true" :safeAreaInsetBottom="true">
			<u-tabbar-item v-for="(item, key) in tabs" :text="item.text" :icon="item.icon" :key="key" style="z-index: 9999"/>
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
					text: "快速下单",
					icon: "star-fill",
					url: "/pages/goods/quickPay"
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
