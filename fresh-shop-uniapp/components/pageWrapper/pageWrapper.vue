<!--
 * @Author: likfees
 * @Date: 2023-03-23 22:16:00
 * @LastEditors: likfees
 * @LastEditTime: 2023-04-23 13:47:08
-->
<template>
	<view :style="{ height: windowHeight + 'px' }">
		<slot></slot>
	</view>
</template>

<script>
export default {
	name: "pageWrapper",
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
		uni.getSystemInfo({
			success: (res) => {
				this.windowHeight = res.windowHeight - this.subHeight;
			},
		});
	}
}
</script>

<style lang="scss"></style>
