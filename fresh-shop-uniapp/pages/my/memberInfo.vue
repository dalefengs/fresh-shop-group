<template>
	<pageWrapper>
		<view class="box1">
			<view class=user-info>
				<view class="face">
					<image :src="user.headerImg"></image>
				</view>
			</view>
			<view class="username">{{ user.nickName }}</view>
			<view class="title">启运冻品欢迎您！填写会员信息加入我们！</view>
		</view>
		<view class="box2">
			<view class="con-bd-item">
				<view class="con-bd-item-name"><span style="color: red; margin-right: 2px">* </span>联系人</view>
				<view class="con-bd-item-con">
					<input type="text" v-model="formData.originContactName" placeholder="请输入您的姓名"
						placeholder-class="input-placeholder" />
				</view>
			</view>
			<view class="con-bd-item">
				<view class="con-bd-item-name">客户名称</view>
				<view class="con-bd-item-con">
					<input type="text" v-model="formData.originCustomerName" placeholder="请输入您所在公司/店铺名称"
						placeholder-class="input-placeholder" />
				</view>
			</view>
			<view class="con-bd-item">
				<view class="con-bd-item-name">客户类型</view>
				<view class="con-bd-item-con">
					<u-radio-group v-model="customerTypeValue" placement="row">
						<u-radio :customStyle="{marginRight: '14px'}" v-for="(item, index) in customerTypeList" :key="index"
							:label="item.name" :name="item.name"  @change="customerTypeChange(item.id)">
						</u-radio>
					</u-radio-group>
				</view>
			</view>
			<view v-if="auditStatus !== 0" class="con-bd-item">
				<view class="con-bd-item-name">审核状态</view>
				<view class="con-bd-item-con">
					<span v-if="auditStatus === 2 || auditStatus === 3">审核中</span>
					<span v-else-if="auditStatus === 4">不通过</span>
					<span v-else-if="auditStatus === 1">通过</span>
				</view>
			</view>
			<view v-if="auditStatus === 4" class="con-bd-item">
				<view class="con-bd-item-name">审核信息</view>
				<view class="con-bd-item-con">
					<input type="text" v-model="user.auditRemark" disabled="disabled" placeholder="无"
						placeholder-class="input-placeholder" />
				</view>
			</view>
			<view v-if="auditStatus !== 0" class="con-bd-item">
				<view class="con-bd-item-name">申请时间</view>
				<view class="con-bd-item-con">
					<text>{{user.applyTime | parseDate}}</text>
				</view>
			</view>
		</view>
		<view class="btn-box" @click="submit">
			<view v-if="auditStatus === 1" class="btn">修改</view>
			<view v-else-if="[2,3].includes(auditStatus)" class="btn">等待审核</view>
			<view v-else-if="auditStatus === 4" class="btn">重新提交</view>
			<view v-else-if="auditStatus === 0" class="btn">提交</view>
		</view>
		<u-toast style="z-index:9998;" ref="toast"></u-toast>
	</pageWrapper>

</template>

<script>
	import {
		getToken,
		setUser,
		getUser,
		getRole,
		setFirstEntry,
		setRole
	} from "@/store/storage";
	import {
		setSelfInfo
	} from "@/api/user";
	import {
		getUserAuditStatus
	} from "@/api/user";

	export default {
		data() {
			return {
				token: '',
				user: null,
				role: {},
				auditStatus: 0,
				formData: {
					AuthorityIds: [1000]
				},
				customerTypeList: [
					{
						name: '零售',
						id: 1000,
						disabled: false
					},
					{
						name: '月结',
						id: 1001,
						disabled: false
					}
				],
				customerTypeValue: '零售'
			}
		},
		onLoad() {
			this.token = getToken()
			if (!this.token) {
				this.$message(this.$refs.toast).error("请请先登录").then(() => {
					uni.redirectTo({
						url: '/pages/my/my'
					})
				})
				return
			}
			setFirstEntry(true) // 已经进入
			this.user = getUser()
			this.role = getRole()
			if (this.user) {
				if (this.user.auditStatus === 1) {
					this.auditStatus = 1
					this.formData.originContactName = this.user.originContactName
					this.formData.originCustomerName = this.user.originCustomerName
				} else {
					getUserAuditStatus().then(res => {
						this.auditStatus = res.data.auditStatus
						this.user.auditStatus = res.data.auditStatus
						this.user.auditRemark = res.data.auditRemark
						this.user.applyTime = res.data.applyTime
						if (this.auditStatus === 3) {
							this.formData.originContactName = this.user.changeContactName
							this.formData.originCustomerName = this.user.changeCustomerName
						} else {
							this.formData.originContactName = this.user.originContactName
							this.formData.originCustomerName = this.user.originCustomerName
						}
						setUser(this.user)
					})
				}
				// 去除末尾客户
				this.customerTypeValue = this.role.authorityName.replace("客户", "")
			}


		},
		methods: {
			customerTypeChange(id) {
				this.formData.AuthorityIds = [id]
			},
			async submit() {
				if (this.auditStatus === 2) {
					this.$message(this.$refs.toast).error("审核中,请勿重复提交")
					return false
				}

				if (this.formData.originContactName === "") {
					this.$message(this.$refs.toast).error("请输入联系人姓名")
					return false
				}

				// 修改原有信息
				if (this.auditStatus === 1) {
					this.formData.auditStatus = 3
					this.formData.changeContactName = this.formData.originContactName
					this.formData.changeCustomerName = this.formData.originCustomerName
					this.formData.originContactName = this.user.originContactName
					this.formData.originCustomerName = this.user.originCustomerName
				} else {
					this.formData.auditStatus = 2
				}
				const res = await setSelfInfo(this.formData)
				if (res.code !== 0) {
					this.$message(this.$refs.toast).error("提交失败！")
					return false
				}
				this.auditStatus = this.formData.auditStatus
				this.user.originContactName = this.formData.originContactName
				this.user.originCustomerName = this.formData.originCustomerName
				this.user.auditStatus = this.auditStatus
				this.user.applyTime = this.getNowFormatDate()
				setUser(this.user)
				this.role.authorityName = this.customerTypeValue + "客户"
				this.role.authorityId = this.formData.AuthorityIds[0]
				setRole(this.role)
				this.$message(this.$refs.toast).success("提交成功，请耐心等待审核！")
			},
			// 获取当前时间
			getNowFormatDate() {
				var date = new Date();
				var year = date.getFullYear();
				var month = date.getMonth() + 1;
				var strDate = date.getDate();
				var hour = date.getHours();
				var minute = date.getMinutes();
				var second = date.getSeconds();
				if (month >= 1 && month <= 9) {
					month = "0" + month;
				}
				if (strDate >= 0 && strDate <= 9) {
					strDate = "0" + strDate;
				}
				if (hour >= 0 && hour <= 9) {
					hour = "0" + hour;
				}
				if (minute >= 0 && minute <= 9) {
					minute = "0" + minute;
				}
				if (second >= 0 && second <= 9) {
					second = "0" + second;
				}
				var currentdate = year + '-' + month + '-' + strDate +
					" " + hour + ':' + minute + ':' + second;
				return currentdate;
			}
		}

	}
</script>

<style scoped lang="scss">
	.box1 {
		margin: 10px 10px 12px;
		border-radius: 10px;
		box-shadow: 0px 0px 20px #f4f3f3;
		background: #FFFFFF;
		padding: 20px 10px 20px;
	}

	.box2 {
		margin: 10px 10px 12px;
		border-radius: 10px;
		box-shadow: 0px 0px 20px #f4f3f3;
		background: #FFFFFF;
		padding: 20px 10px 30px;
	}

	.user-info {
		display: flex;
		align-items: center;
		width: 100%;
		height: 100%;
		justify-content: center;

		.face {
			flex-shrink: 0;
			width: 20vw;
			height: 20vw;

			image {
				width: 20vw;
				height: 100%;
				border-radius: 100%
			}
		}
	}

	.username {
		font-size: 17px;
		width: 100%;
		text-align: center;
		margin: 16px 0 12px 0;
	}

	.title {
		font-size: 17px;
		width: 100%;
		text-align: center;
	}

	.con-bd-item {
		display: flex;
		align-items: center;
		width: 100%;
		height: 55px;
	}

	.con-bd-item-name {
		width: 110px;
		height: 110px;
		font-size: 16px;
		color: #333333;
		line-height: 110px;
		text-align: right;
		margin-right: 15px;
	}

	.con-bd-item-con {
		display: flex;
		align-items: center;
		width: calc(100% - 60px);
		height: 54px;
		border-bottom: 1px solid #F7F7F7;
	}

	.con-bd-item-con input {
		width: 100%;
		height: 30px;
	}

	.btn-box {
		position: fixed;
		bottom: 10px;
		height: 50px;
		width: 100%;
		display: flex;
		align-items: center;
		justify-content: center;

		.btn {
			width: 90%;
			height: 50px;
			color: #fff;
			background-color: #2979ff;
			border-radius: 30px;
			display: flex;
			align-items: center;
			justify-content: center;
			font-size: 20px;
		}
	}
</style>