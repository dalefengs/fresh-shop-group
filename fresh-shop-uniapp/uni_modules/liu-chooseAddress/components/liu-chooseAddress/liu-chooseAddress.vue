<template>
	<view class="view-main">
		<map :latitude="latitude" :longitude="longitude" scale="16" show-location style="width: 100%; height: 30vh;">
			<image class="map-marker" :src="markerIcon"></image>
		</map>
		<view class="page-con">
			<view class="con-hd">
				<view class="choose-address" v-if="!addressInfo || !addressInfo.address" @click="chooseAddressView">
					选择收货地址 >
				</view>
				<view class="choose-address2" v-else>
					<view class="address2-l">
						<view class="address2-lhd">{{addressInfo.title}}</view>
						<view class="address2-lbd">{{addressInfo.address}}</view>
					</view>
					<view class="address2-r" @click="chooseAddressView">修改地址</view>
				</view>
			</view>
			<view class="con-bd">
				<view class="con-bd-item">
					<view class="con-bd-item-name">门牌号</view>
					<view class="con-bd-item-con">
						<input type="text" v-model="formData.detail" placeholder="详细地址，例如1层101室"
							placeholder-class="input-placeholder" />
					</view>
				</view>
				<view class="con-bd-item">
					<view class="con-bd-item-name">标签</view>
					<view class="con-bd-item-con">
						<view :class="formData.lable == item ? 'lable-item lable-item-on':'lable-item'"
							v-for="(item,index) in labelData" :key="index" @click="chooseLable(item)">{{item}}</view>
					</view>
				</view>
				<view class="con-bd-item">
					<view class="con-bd-item-name">联系人</view>
					<view class="con-bd-item-con con-bd-item-con2">
						<input type="text" v-model="formData.name" placeholder="请填写收货人姓名"
							placeholder-class="input-placeholder" />
						<view class="con-bd-sex-box" @click="chooseSex(1)">
							<image class="choose-sex-icon" :src="formData.sex == 1 ? sele2Icon : noseleIcon"
								mode=""></image>
							先生
						</view>
						<view class="con-bd-sex-box" @click="chooseSex(2)">
							<image class="choose-sex-icon" :src="formData.sex == 2 ? sele2Icon : noseleIcon"
								mode=""></image>
							女士
						</view>
					</view>
				</view>
				<view class="con-bd-item">
					<view class="con-bd-item-name">手机号</view>
					<view class="con-bd-item-con">
						<input type="number" maxlength="11" v-model="formData.mobile" placeholder="请填写收货人手机号码"
							placeholder-class="input-placeholder" />
					</view>
				</view>
				<view class="con-bd-agree" @click="changeIsDefault">
					<image class="agreeicon" :src="isDefault ? seleIcon : noseleIcon" mode=""></image>
					<view class="agreetext">设为默认地址</view>
				</view>
				<view class="con-ft-btn"
					:class="{'con-ft-btn1':addressInfo && addressInfo.address,'con-ft-btn3':!addressInfo || !addressInfo.address}"
					@click="submitAddress">
					保存地址</view>
				<view class="con-ft-btn con-ft-btn2" v-if="this.id" @click="deleteAddress">删除地址</view>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
                id: null,
				formData: {
					address: '',
                    detail: '',
					name: '',
					mobile: '',
					lable: '家',
					sex: 1, //1 先生 2 女士
				},
				labelData: ['家', '公司', '学校', '父母家'],
				addressInfo: {}, //选中的地址信息
				isDefault: false, //是否是默认地址
				longitude: null,
				latitude: null,
				markerIcon: require('../../static/marker.png'),
				seleIcon: require('../../static/sele.png'),
				sele2Icon: require('../../static/sele2.png'),
				noseleIcon: require('../../static/no-sele.png'),
			}
		},
		mounted() {
			this.getLocation()
		},
		methods: {
			//数据回显时使用此方法
			setData(obj) {
                this.id = obj.id
				this.formData.address = obj.address
				this.formData.name = obj.name
				this.formData.mobile = obj.mobile
				this.formData.lable = obj.lable
				this.formData.sex = obj.sex
				this.formData.detail = obj.detail
				this.addressInfo.title = obj.title
				this.addressInfo.address = obj.address
				this.addressInfo.latitude = obj.latitude
				this.addressInfo.longitude = obj.longitude
				this.isDefault = obj.isDefault ? true : false
			},
			//获取当前的地理位置
			getLocation() {
				uni.showLoading({
					title: '正在获取定位中...',
				})
				uni.getLocation({
					type: 'gcj02',
					isHighAccuracy: true,
					geocode: 'true',
					highAccuracyExpireTime: 3500,
					success: (res) => {
						uni.hideLoading();
						this.longitude = res.longitude;
						this.latitude = res.latitude;
					},
					fail: (res) => {
						if (res.errMsg == "getLocation:fail auth deny") {
							uni.showModal({
								content: '检测到您没打开获取信息功能权限，是否去设置打开？',
								confirmText: "确认",
								cancelText: '取消',
								success: (res) => {
									if (res.confirm) {
										uni.openSetting({
											success: (res) => {}
										})
									} else {
										return false;
									}
								}
							})
						}
					}
				})
			},
			//保存地址
			submitAddress() {
				let obj = {
					...this.formData,
					...this.addressInfo
				}
				obj.id = this.id
				obj.isDefault = this.isDefault ? 1 : 0
				this.$emit('submit', obj)
			},
			//删除地址
			deleteAddress() {
				this.$emit('detele', this.id)
			},
			//选择地址、修改编辑地址
			chooseAddressView() {
				uni.chooseLocation({
					success: res => {
						if (!res.name || !res.address || !res.latitude || !res.longitude) return
						let addressObj = {
							title: res.name,
							address: res.address,
							latitude: res.latitude,
							longitude: res.longitude
						}
						this.latitude = res.latitude
						this.longitude = res.longitude
						this.addressInfo = addressObj
					}
				});
			},
			//联系人性别选择
			chooseSex(type) {
				this.formData.sex = type
			},
			//标签选择
			chooseLable(item) {
				this.formData.lable = item
			},
			//点击隐私协议
			changeIsDefault() {
				this.isDefault = !this.isDefault
			}
		}
	}
</script>

<style scoped>
	.view-main {
		width: 100%;
		height: 100%;
		background: #f3f3f3;
	}

	.map-marker {
		width: 44rpx;
		height: 62rpx;
		position: absolute;
		top: 15vh;
		left: 50%;
		transform: translate(-50%, -50%);
	}

	.map-tip {
		font-size: 20rpx;
		color: #b6b6b6;
		line-height: 42rpx;
		padding: 0 20rpx;
		position: absolute;
		left: 50%;
		top: 30rpx;
		box-shadow: 0px 1px 10px 1px rgba(153, 153, 153, 0.34);
		background-color: #fff;
		border-radius: 4px;
		transform: translateX(-50%);
	}

	.poiss-box {
		width: 100%;
		padding: 20rpx 0;
		background-color: #FFFFFF;
		font-size: 26rpx;
		position: fixed;
		bottom: 0;
		left: 0;
	}

	.poiss-item {
		display: flex;
		align-items: center;
		justify-content: space-between;
		width: calc(100% - 48rpx);
		height: 100rpx;
		padding: 0 24rpx;
		border-bottom: 1px solid #f9f9f9;
	}

	.poiss-item-l {
		width: calc(100% - 184rpx);
	}

	.poiss-item-lhd {
		font-size: 32rpx;
		color: #333333;
		line-height: 48rpx;
		text-overflow: ellipsis;
		white-space: nowrap;
		overflow: hidden;
	}

	.poiss-item-lbd {
		font-size: 28rpx;
		color: #b6b6b6;
		line-height: 42rpx;
		text-overflow: ellipsis;
		white-space: nowrap;
		overflow: hidden;
	}

	.poiss-item-r {
		width: 36rpx;
		height: 36rpx;
		margin: 0 30rpx;
	}

	.primary-btn {
		width: 688rpx;
		height: 88rpx;
		margin: 0 auto;
		margin-top: 28rpx;
		background: #3579FF;
		border-radius: 44rpx;
		font-size: 32rpx;
		color: #FFFFFF;
		line-height: 88rpx;
		text-align: center;
	}

	.page-con {
		width: 702rpx;
		height: 75vh;
		position: fixed;
		bottom: 0;
		left: 24rpx;
	}

	.con-hd {
		display: flex;
		align-items: center;
		width: 702rpx;
		height: 128rpx;
		margin: 0 auto;
		background: #FFFFFF;
		box-shadow: 0rpx 4rpx 8rpx 0rpx rgba(88, 102, 123, 0.2);
		border-radius: 24rpx;
	}

	.choose-address {
		display: flex;
		align-items: center;
		justify-content: center;
		width: 92%;
		height: 64rpx;
		margin: 0 auto;
		border-radius: 12rpx;
		border: 2rpx solid #3579FF;
		font-size: 32rpx;
		color: #3579FF;
	}

	.choose-address2 {
		display: flex;
		align-items: center;
		justify-content: space-between;
		width: calc(100% - 48rpx);
		padding: 0 24rpx;
	}

	.address2-l {
		width: calc(100% - 184rpx);
	}

	.address2-lhd {
		font-size: 32rpx;
		font-weight: bold;
		color: #333333;
		line-height: 48rpx;
		text-overflow: ellipsis;
		white-space: nowrap;
		overflow: hidden;
	}

	.address2-lbd {
		font-size: 28rpx;
		color: #999999;
		line-height: 42rpx;
		text-overflow: ellipsis;
		white-space: nowrap;
		overflow: hidden;
	}

	.address2-r {
		width: 144rpx;
		height: 60rpx;
		background: #FFFFFF;
		border-radius: 12rpx;
		border: 2rpx solid #3579FF;
		font-size: 28rpx;
		color: #3579FF;
		line-height: 60rpx;
		text-align: center;
	}

	.con-bd {
		width: 654rpx;
		height: calc(90vh - 160rpx);
		padding: 8rpx 24rpx 0;
		margin-top: 24rpx;
		background: #FFFFFF;
		border-radius: 24rpx 24rpx 0rpx 0rpx;
	}

	.con-bd-item {
		display: flex;
		align-items: center;
		width: 100%;
		height: 110rpx;
	}

	.con-bd-item-name {
		width: 120rpx;
		height: 110rpx;
		text-align: left;
		font-size: 28rpx;
		font-weight: bold;
		color: #333333;
		line-height: 110rpx;
	}

	.con-bd-item-con {
		display: flex;
		align-items: center;
		width: calc(100% - 120rpx);
		height: 108rpx;
		border-bottom: 2rpx solid #F7F7F7;
	}

	.con-bd-item-con input {
		width: 100%;
		height: 60rpx;
	}

	.lable-item {
		width: 88rpx;
		height: 48rpx;
		margin-right: 18rpx;
		background: #FFFFFF;
		border-radius: 6rpx;
		border: 2rpx solid #E3E3E3;
		font-size: 24rpx;
		color: #333333;
		line-height: 48rpx;
		text-align: center;
	}

	.lable-item-on {
		background: #3579FF !important;
		border: 2rpx solid #3579FF;
		color: #FFFFFF !important;
	}

	.con-bd-item-con2 input {
		width: 50%;
		height: 60rpx;
	}

	.con-bd-sex-box {
		display: flex;
		align-items: center;
		justify-content: center;
		margin-left: 16rpx;
		font-size: 28rpx;
	}

	.choose-sex-icon {
		width: 36rpx;
		height: 36rpx;
		margin-right: 12rpx;
	}

	.con-bd-agree {
		display: flex;
		align-items: center;
		margin-top: 26rpx;
		width: 100%;
	}

	.agreeicon {
		width: 36rpx;
		height: 36rpx;
		margin-right: 30rpx;
	}

	.agreetext {
		display: flex;
		align-items: center;
		font-size: 28rpx;
		color: #999999;
	}

	.con-ft-btn {
		width: 644rpx;
		height: 88rpx;
		margin-top: 36rpx;
		border-radius: 44rpx;
		font-size: 32rpx;
		line-height: 88rpx;
		text-align: center;
	}

	.con-ft-btn:last-child {
		margin-top: 15rpx;
	}

	.con-ft-btn1 {
		background: #3579FF;
		color: #FFFFFF;
	}

	.con-ft-btn2 {
		background: #FFFFFF;
		border: 2rpx solid #DEDEDE;
		color: #666666;
	}

	.con-ft-btn3 {
		background: #DEDEDE;
		border: 2rpx solid #DEDEDE;
		font-size: 32rpx;
		color: #FFFFFF;
	}

	.input-placeholder {
		font-size: 28rpx;
		color: #999999;
	}
</style>
