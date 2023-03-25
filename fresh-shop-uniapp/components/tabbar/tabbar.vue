<template>
	<view>
		<u-tabbar :value="tabsId" @change="tabsChange" :fixed="true" :placeholder="true" :safeAreaInsetBottom="true">
			<u-tabbar-item text="首页" icon="home"></u-tabbar-item>
			<u-tabbar-item text="分类" icon="grid"></u-tabbar-item>
			<!-- <u-tabbar-item text="快速下单" icon="photo" ></u-tabbar-item> -->
			<u-tabbar-item text="购物车" icon="shopping-cart"></u-tabbar-item>
			<u-tabbar-item text="我的" icon="account"></u-tabbar-item>
		</u-tabbar>
	</view>
</template>

<script>
	export default {
		name:"tabbar",
		data() {
			return {
				tabsId: 0
			};
		},
		methods: {
			tabsChange(index) {
				this.tabsId = index
			}
		}
	}
</script>

<style lang="scss">

</style>