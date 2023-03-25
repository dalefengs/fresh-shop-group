<!--
 * @Author: likfees
 * @Date: 2023-03-23 22:16:00
 * @LastEditors: likfees
 * @LastEditTime: 2023-03-24 12:40:29
-->
<template>
	<view>
		<slot></slot>
		<u-toast ref="uToast"></u-toast>
	</view>
</template>

<script>
	export default {
		name:"pageWrapper",
		data() {
			return {
				
			};
		},
		mounted() {
			getApp().globalData.toast = this.$refs.uToast
		}
	}
</script>

<style lang="scss">

</style>