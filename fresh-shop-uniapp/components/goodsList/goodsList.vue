<template>
	<view>
		<view class="goods--box" v-if="!vertical">
			<view class="goods--item" v-for="(item, index) in lists" :key="index" @click="goodsClick(item)">
                <view class="item-cover-mask" v-if="item.store <= 0">
                    <view class="item-cover-mask-text">补货中</view>
                </view>
                <img v-if="item.images[0]" class="item-img" :src="item.images.length > 0 ? item.images[0].url : ''" alt="">
                <img v-else class="item-img" src="/static/nopicture.jpg" alt="">
				<view class="item-content">
					<view class="title">{{ item.name }}</view>
					<view class="bottom-txt">
						<view>
							<text class="price king-font17" v-if="isPoint">{{ item.costPrice || '0' }}积分</text>
							<text class="price" v-else>{{ priceType }}{{ item.price > 0 && item.price < item.costPrice ? item.price : item.costPrice || '0' }}</text>
							<text v-if="item.price > 0 && item.price < item.costPrice && !isPoint" class="del-price">{{ priceType }}{{ item.costPrice || '0'
							}}</text>
						</view>
						<view class="sale-num"><text>已售 {{ item.sale }}</text></view>
					</view>
				</view>
			</view>
		</view>
        <!--  每个商品一行   -->
		<view class="goodsv--box" v-else>
			<view class="goodsv--item" v-for="(item, index) in lists" :key="index" @click="goodsClick(item)">
                <view class="item-cover-mask" v-if="item.store <= 0">
                    <view class="item-cover-mask-text">补货中</view>
                </view>
				<view class="item-cover" v-if="item.images && item.images[0]"
					:style="{ backgroundImage: 'url(' + (item.images.length > 0 ? item.images[0].url : '') + ')' }" />
				<image class="item-cover" v-else src="/static/nopicture.jpg" />

				<view class="item-content">
					<text class="title">{{ item.name }}</text>
                    <view class="store">
                        <text>库存：{{ item.store }}</text>
                    </view>
					<view class="bottom-txt">
						<view>
							<text class="price">{{ priceType }}{{ item.price > 0 && item.price < item.costPrice ? item.price : item.costPrice| '0' }}</text>
							<text class="unit">/{{ item.unit }}</text>
							<text v-if="item.price > 0 && item.price < item.costPrice" class="del-price">{{ priceType }}{{ item.costPrice || '0'
							}}</text>
						</view>
						<view class="sale-num"><text>已售 {{ item.sale }}</text></view>
						<view class="add-cart-button" v-if="item.store > 0">
							<u-icon name="shopping-cart" color="#ffffff" size="24"></u-icon>
                            <view class="badge">
                                <u-badge max="99" :value="item.cartNum" shape="circle"></u-badge>
                            </view>
						</view>
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
		},
		// 样式
		vertical: {
			type: Boolean,
			default: false
		},
		// 禁用商品点击默认跳转， 自定义监听 @onGoods
		disableJump: {
			type: Boolean,
			default: false
		},
		// 是否是积分商品
		isPoint: {
			type: Boolean,
			default: false
		}
	},
	mounted() {
	},
	methods: {
		// 商品点击 默认跳转详情
		goodsClick(goods) {
			if (this.disableJump) {
				this.$emit("onGoods", goods)
			} else {
				uni.navigateTo({
					url: `/pages/goods/detail?id=${goods.ID}`
				})
			}
		}
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

.ellipsis-2-v {
  max-width: 80%;
  display: -webkit-box;
  overflow: hidden;
  text-overflow: ellipsis;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 2;
}

.store {
  color: #888;
  font-size: 12px;
  font-weight: normal;
  margin-left: 3px;
}

.bottom-txt {
  display: flex;
  justify-content: space-between;
  align-items: flex-end;

  .price {
    color: #fa3534;
    font-size: 16px;
    font-weight: 600;

    text {
      font-size: 12px;
    }
  }

  .unit {
    color: #666;
    font-size: 12px;
  }

  .del-price {
    color: #666;
    font-weight: 600;
    font-size: 11px;
    text-decoration: line-through;
    margin-left: 4px;
  }

  .sale-num {
    color: #888;
    font-size: 12px;
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
    border-radius: $radius;
    overflow: hidden;
    margin: $margin;
    position: relative;


    .item-img {
      width: 100%;
      border-radius: 10px;
	    height: 180px;
    }

    .item-cover-mask {
      width: 100%;
      height: 100%;
      flex-shrink: 0;
      z-index: 9999;
      background-color: rgba(173, 171, 171, 0.5);
      position: absolute;
      display: flex;
      align-items: center;
      justify-content: center;
      border-radius: 10px;

      .item-cover-mask-text {
        font-size: 14px;
        padding: 6px 12px;
        background-color: rgba(0, 0, 0, 0.7);
        color: #fff0f0;
        border-radius: 20px;
      }
    }

    .item-content {
      padding: $padding - 10;

      & > .title {
        font-size: 16px;
        line-height: 24px;
        height: 30px;
        color: #333;
        @extend .ellipsis-1;
      }

    }
  }
}


.goodsv--box {
  display: flex;
  flex-wrap: wrap;
  padding: $padding/3;

  .goodsv--item {
    background: white;
    width: calc(100% - #{$margin*2});
    background: white;
    border-radius: $radius;
    overflow: hidden;
    margin: $margin;
    display: flex;
    padding: 8px 1px;

    .item-cover-mask {
      width: 75px;
      height: 75px;
      flex-shrink: 0;
      background-color: rgba(173, 171, 171, 0.5);
      position: absolute;
      display: flex;
      align-items: center;
      justify-content: center;
      border-radius: 10px;
      margin: 10px 0 0 10px;

      .item-cover-mask-text {
        font-size: 12px;
        padding: 2px 6px;
        background-color: rgba(0, 0, 0, 0.7);
        color: #fff0f0;
        border-radius: 10px;
      }
    }


    .item-cover {
      width: 75px;
      height: 75px;
      background-position: center;
      background-size: cover;
      display: inline-block;
      border-radius: 8px;
      margin-top: 10px;
      margin-left: 10px;
    }

    .item-content {
      display: flex;
      flex-direction: column;
      justify-content: space-between;
      padding: 8px 10px 8px 5px;
      flex: 1;

      & > .title {
        font-size: 28 rpx;
        line-height: 40 rpx;
        height: 80 rpx;
        color: #333;
        @extend .ellipsis-2-v;
      }

      .bottom-txt {
        justify-content: normal;

        .sale-num {
          margin-left: 6px;
        }
      }


    }
  }
}

.add-cart-button {
  position: absolute;
  border-radius: 100%;
  background: #2979ff;
  color: white;
  right: 20px;
  cursor: pointer;
  padding: 2px 2px 2px 1px;

  .badge {
    position: absolute;
    top: -6px;
    right: -6px;
  }
}
</style>
