<!--
 * @Author: dalefeng
 * @Date: 2023-03-23 15:52:23
 * @LastEditors: dalefeng
 * @LastEditTime: 2023-03-24 15:06:32
-->
<script>
export default {
	data() {
		return {
			globalData: {
				toast: Object,//全局提示，必须在页面初始化
			},
		}
	},
	onLaunch: function () {
		console.log('App Launch')
	},
	onShow: function () {
		console.log('App Show')
	},
	onHide: function () {
		console.log('App Hide')
	}
}
</script>

<style lang="scss">
@import url("style/king_style.css");
/*每个页面公共css */
body{
	background: #f7f7f7;
}
</style>
