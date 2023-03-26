<!--
 * @Author: likfees
 * @Date: 2023-03-23 22:16:00
 * @LastEditors: likfees
 * @LastEditTime: 2023-03-26 13:38:56
-->
<template>
	<view :style="{height: windowHeight + 'px'}">
		<slot></slot>
		<u-toast ref="uToast"></u-toast>
	</view>
</template>

<script>
	export default {
		name:"pageWrapper",
		data() {
			return {
				windowHeight: 0,
			};
		},
		props: {
			subHeight: {
				type: Number,
                default: 0,
			}
		},
		mounted() {
			getApp().globalData.toast = this.$refs.uToast
			uni.getSystemInfo({
			success: (res) => {
				this.windowHeight = res.windowHeight - this.subHeight;
			},
		});
		}
	}
</script>

<style lang="scss">

</style>