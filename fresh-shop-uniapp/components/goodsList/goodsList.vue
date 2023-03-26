<template>
	<view>
		<view class="goods--box">
			<view class="goods--item" v-for="(item, index) in lists" :key="index" @click="$emit('onGoods', item)">
				<view class="item-cover"
					:style="{ backgroundImage: 'url(' + (item.images.length > 0 ? item.images[0].url : '') + ')' }" />
				<view class="item-content">
					<view class="title">{{ item.name }}</view>
					<view class="bottom-txt">
						<view>
							<text class="price">{{ priceType }}{{ item.costPrice || '0' }}</text>
							<text v-if="item.price < item.costPrice" class="del-price">{{ priceType }}{{ item.price || '0' }}</text>
						</view>
						<view class="sale-num"><text>已售 {{ item.sale }}</text></view>
					</view>
				</view>
			</view>
		</view>
	</view>
</template>
<script>
export default {
	props: {
		lists: {
			type: Array,
			required: true,
			default: () => {
				return []
			}
		},
		priceType: {
			type: String,
			required: false,
			default: "￥"
		}
	},
	mounted() {
	}
}
</script>
<style lang="scss" scoped>
$padding: 30rpx;
$margin: 10rpx;
$radius: 20rpx;

.ellipsis-1 {
	max-width: 100%;
	overflow: hidden;
	white-space: nowrap;
	text-overflow: ellipsis;
}

.ellipsis-2 {
	max-width: 100%;
	display: -webkit-box;
	overflow: hidden;
	text-overflow: ellipsis;
	-webkit-box-orient: vertical;
	-webkit-line-clamp: 2;
}

.bottom-txt {
	display: flex;
	justify-content: space-between;
	align-items: flex-end;

	.price {
		color: #d4282d;
		font-size: 32rpx;
		font-weight: 600;

		text {
			font-size: 24rpx;
		}
	}

	.del-price {
		color: #666;
		font-weight: 600;
		font-size: 22rpx;
		text-decoration: line-through;
		margin-left: 4px;
	}

	.sale-num {
		color: #888;
		font-size: 24rpx;
		font-weight: normal;
		margin-bottom: 2px;
	}
}

.goods--box {
	display: flex;
	flex-wrap: wrap;
	padding: $padding/3;

	.goods--item {
		background: white;
		width: calc(50% - #{$margin*2});
		background: white;
		border-radius: $radius;
		overflow: hidden;
		margin: $margin;

		.item-cover {
			width: 100%;
			padding-top: 100%;
			background-position: center;
			background-size: cover;
		}

		.item-content {
			padding: $padding - 10;

			&>.title {
				font-size: 28rpx;
				line-height: 40rpx;
				height: 80rpx;
				color: #333;
				@extend .ellipsis-2;
			}

		}
	}
}</style>
