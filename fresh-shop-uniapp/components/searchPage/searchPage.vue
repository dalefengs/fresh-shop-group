<!--
 * @Author: dalefeng
 * @Date: 2023-03-25 13:43:18
 * @LastEditors: dalefeng
 * @LastEditTime: 2023-03-26 17:11:39
-->
<template>
	<view class="search-panel">
		<!-- 检索框 -->
		<view class="search-panel-base">
			<view class="left">
				<image :src="searchIcon" class="search-icon"></image>
				<input class="search-input" v-model="keyword" type="text" :placeholder="placeholder" confirm-type="search" @confirm="search" />
				<image v-if="keyword" :src="delIcon1" class="search-del-icon" @click="clearKeyword"></image>
			</view>
			<view class="right">
				<text class="search-btn" @click="search" :style="'border: 1rpx solid '+borderColor+';color: '+textColor+';'">{{searchName?searchName:'搜索'}}</text>
			</view>
		</view>

		<!-- 其他 -->
		<view v-if="showMore" class="search-panel-more">
			<!-- 历史记录 -->
			<view class="search-more-title">
				<text>搜索历史</text>
				<image class="search-keyword-operate-icon" :src="delIcon2" @click="clearHistory"></image>
			</view>
			<view class="search-keywords" v-if="historyKeywords.length>0">
				<text class="search-keyword" v-for="(item,index) in historyKeywords" :key="index" @click="fillKeyword(item)">{{item}}</text>
			</view>
			<view v-else class="search-null">暂无</view>

			<!-- 热门搜索 -->
			<view class="search-more-title" v-if="hotKeywords.length>0">
				<text>热门搜索</text>
				<image class="search-keyword-operate-icon" :src="showHotKeywords ? eyeOpenIcon : eyeCloseIcon" @click="showHotKeywords = !showHotKeywords">
				</image>
			</view>
			<view class="search-keywords" v-if="hotKeywords.length>0 && showHotKeywords">
				<text class="search-keyword" v-for="(item,index) in hotKeywords" :key="index" @click="fillKeyword(item)">{{item}}</text>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		watch: {
			keyword(newVal, oldVal) {
				this.$emit('changeKeyword', newVal)
			}
		},
		data() {
			return {
				keyword: '',
				showHotKeywords: true,
				historyKeywords: [],
				delIcon1: 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAYAAACqaXHeAAAAAXNSR0IArs4c6QAABK9JREFUeF7tm49RFDEUxl8qECoQKhAqgHcNCBUIFXhUoFQgVuBRgdLAPajAswKxArWCON9O9mYvZHfzf7nDzOwMzOwl+X55SV5e3ioqXERkT2t9opQ6IiI8e0R0YJ5u649EhOcPEa201iul1AMz4/9iRZWoWUQg9B0RnRrRKc2siOieiG6ZGX9nLdkAYKSN6LljdHN1GhZyY2BksYxkAEb4eyKCcECoUSAeID6nTpEkACICM0dHagm34QLEnJlvY6lHARARLGJfzByPbTvn77BGXDIzpkhQCQYgImdG/FSj3icQ1gAI30IIBAEQkU9mroe0UfvdG2a+8m3UG4CIwOQvfCue+L0FM1/69MELgIh8JSKY/jYVLwijALZs5O0BGoUwCEBEsMVhj9/mAl8BPoqz9AIwqz1MfxfKed/u4ARg9vnvEzo4uaFjizx2+Ql9AOQZOTm5YNwzM9uVPQEgItjqsOXtYoGjtOgK2wBgDjY/d8j07UHEVDjsHqBsAB+J6MMuDn1H0zUzQ2dT1gBSR19r/VcphZPZAgERrfWNUuokB0yt9QPqNhGjM1P3q8i6N6ygCwANwNePLU/ml4hgvuHInFIQCdpwwTOsU1fMDB9nwwIw93HMjS37ruBEIoQn4tE5Y62/YzsKS2LmwzUAE8PDvp9SsM86Y3aREJziDQDEHLP0t5kCmVxe5z7bEg2E0Cve9DeHn9K4yC0A0ATV1DJ4+PCEMCY+17F8xczHKsN8sqGlQKglvu3zPgDgnJ/70BMDobZ4QDgHgFLOTwiEKcQDwDUAIIj4NnXy9/zeBwLZ+3y3rsIBmTu1XC7vc3lsMRCGwBcWT42HKSKpDpCP8YyGpuxKSos37a0AQPsoyPCON4RK4htJNQGgvVEINcVPAWBwtTdeXo4DlLfB1rSAUfGRbrO3WNeL2AUelVKvk2oZ/7G3+JoQtNY/amyDweJrQWi3wZKOkI+HB0eo9x7P8wA1boPuN+5KusI+4ttIj4/HmBpZciFoXOESh6EQ8W3HpoDQHIaQ6JASXrLJxoifCsJ+ExBZLpfIyXsTO5E6v0sRXxUCdoDZbHaULSSGFXU2myEv0FkCPbzB6ZDpALcREssWZHSpDxQ/agk5g7jre4EMDlFfWDwlhue0hNR1qzX/5izQcTpKXIykiO+1hFIXI0hqhluccuUEiHfmggW3TL1rQuCCizxAZH4hDxDRq+jkTHOFd9Be4vy/HO2OhEltT7GCwIGt+7o9+htrQGcteLkJEi2ETPts3eEdaa3PT+lNkjJfbMQuiM9NPHIXjryTpND7QoekqcCEpcl11oOXmyjZgVA1SJnZREajUaO5wmY6lIwaZda8rm5UvHMbHDjNbZMleIkPAmAsYRvWhMHkaHuAvaaA5S0iTW2RcGYoYvLGy7so+slMZ2E8MBCy5AGmEjF5hBBf/qMpyxouEpMWk7R3kzNjKwqeAnZDJjgx11rPa00LIxzrET6QSvqCNBlAZ1ogugyLAIgiV21a619KKQhHpChJeNvvbACsqYEYI2Ccpkabm/s7pRAQgejn+/H0gP8Ay2i/IkcSdfP5vG0lGF1EfJRSzefz5kHyZZaR7uvfP1aimGHUO86VAAAAAElFTkSuQmCC',
				delIcon2: 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAYAAACqaXHeAAAAAXNSR0IArs4c6QAAArdJREFUeF7tWstx20AMBUYjXWN3oBKSCiJXEKcDuwPrIHB0sn3ScHWgO0g6iFJBnAqSDhJXEOWs0SCDDJXQNEnvcrle2sQeJewHD8Au8AiEgQ8cuP6gAKgHDBwBDYGBO4BegsFDwBjzmpkvEfHUxduYeYOI10T03WWeq2xQAIwxZwDwwfVQJflzIvrouUbt9GAAiOUB4FtHB38TyhNCArABgHcdAfCZiJxCyHbfkABw6RDWVqzyHiIKctYgi4rixph7ALgq4Du/dx6gADi6cO88IH/SMgA4snWvSHJbAJjbPp3Wd4Ax5tczUP6A+ZaIjm0MYA1AmqZbRHxls2hsGWb+nSSJladaAyAhwMw3fQdBlEfEi85DILZVQ+1v7QGhDhB7XQUgtgVi79/aAwp5gejQ+O6maSqFTIaIcjOf1FV2IoeIn5j5p6yZJIkUVJWjuD8znzfJNoHsA0AxL6h9d7MsO9rtdiJ7GLWVnTHmFgDeHgTH4/HxfD6XxObBKOUlratFHwCsip31ej1j5i8FDb4S0axGqXsAIOLJYrGQ36oA+Lc/M18nSXLVJpwUgDaoyRzbYkU9QENA7wC9BPUV0GdQ8wBNhDQT/I+ApsJaC2gxpNWglsPKB1STHEqIFCkxZYReAiW2Wq2mo9HohyspKp+4JpPJtIEUfR6coCiepukVIl4CwB0AnNbR4nlL3S0iChNcK1em5HpNirblHB+bV+QkFYA+0OJNN/Zj1nT9v9xFFssDyn2A0tIqXZ1BW1tzml1adaQR8+9g5vdP/mmsgu93NWRX8ndENG27WOsvQ/lN3GU3aCsdfKwvG3oBkH/4lEbmrlpirUFwbYWpW9gLgMOieThcMPPsCXqIJJfY7Pf7m+VyKZ/RvUYnAHidIPJkBSCyAaJvrx4Q3QSRD6AeENkA0bcfvAf8AcIkGG42KSwxAAAAAElFTkSuQmCC',
				searchIcon: 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAFAAAABQCAYAAACOEfKtAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAyZpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDUuNi1jMTExIDc5LjE1ODMyNSwgMjAxNS8wOS8xMC0wMToxMDoyMCAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvIiB4bWxuczp4bXBNTT0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL21tLyIgeG1sbnM6c3RSZWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9zVHlwZS9SZXNvdXJjZVJlZiMiIHhtcDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIENDIDIwMTUgKFdpbmRvd3MpIiB4bXBNTTpJbnN0YW5jZUlEPSJ4bXAuaWlkOjU4MEVFRkU4RTlEMjExRUJBREJFRTIyQjJDMUE5NkQ3IiB4bXBNTTpEb2N1bWVudElEPSJ4bXAuZGlkOjU4MEVFRkU5RTlEMjExRUJBREJFRTIyQjJDMUE5NkQ3Ij4gPHhtcE1NOkRlcml2ZWRGcm9tIHN0UmVmOmluc3RhbmNlSUQ9InhtcC5paWQ6NTgwRUVGRTZFOUQyMTFFQkFEQkVFMjJCMkMxQTk2RDciIHN0UmVmOmRvY3VtZW50SUQ9InhtcC5kaWQ6NTgwRUVGRTdFOUQyMTFFQkFEQkVFMjJCMkMxQTk2RDciLz4gPC9yZGY6RGVzY3JpcHRpb24+IDwvcmRmOlJERj4gPC94OnhtcG1ldGE+IDw/eHBhY2tldCBlbmQ9InIiPz4L5MHuAAAMU0lEQVR42uxc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',
				eyeOpenIcon: 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAYAAACqaXHeAAAAAXNSR0IArs4c6QAABM5JREFUeF7tWU1OG0sQ7jLMbF9uEHKCwAkenCDmBC+cIHjhbrHCWaFuL4ATBE4QcoJnThDnBskNnK2xXdFn9URkmO6pxuMI2dMSEpL7r7766qvqGlJbPmjL7VctAC0DthyBNgS2nACtCLYh0IbAliPQhsCWE6DNAm0ItCGwZgQuLi72dnd33yql9heLxT4RvVJK4W+/dPRYKTVh5kmn08H/49ls9u3s7Oz7Oq+4lhBwzu0z839KqS4R7a1iADMDgDsiutVaA5hGR6MAWGth9GBVo0MWejAGxpjbplBoBIB1G142tkkgVgLAU/3zujwu8DJC4mSV0Hg2AMPh8JyZB4JL/p7CzN+I6DszF7G8FDyIoxdGRUTQjz0ignCKBhEN+v3+R9Hk0qRkAC4vL189PDz8X6HiVef/gIAx88gYc5dyQX9Ol5kPvZj+U7N+nGXZUa/Xm6SckwRAgvH3zHyVanTo4jh3Op2eKqVOiSgGRDIIYgAQ70qpTzHPg+KdTue03++PQsYMh8NDnyKxX1ELLPM+Ul1srXcAwu5DxMvj+Xx+LK0fRAB4D3wNiR0z/0T6M8Zc1XjxExF1YxRl5rs8z09iVPbiexPSCWSJPM8PJOEgAsA59zXkeS9s72NKnBA6BTa1VPZ7Qlf+DQA60lof1elBLQDOOdD+fdVGMD7P88M6pK21SJVRz1fsf621RtxHh7UWTEABVjVutNYnsQ2iAFhrITqXqxjvYx5ZI3kQ0VFME4oNa0BAnXATOjwIAB4xOzs7oD4eLn8MqeexKHI5ZIplavTsqKKyiAVeo0YBTZhkWfYmxNIgACHaQvBQuEhVNqAf91pr5PffwzmHzFEGYay1PpBQBw5DUVWVJgG0Mea4ap9KAKy1eMV9DhwcpVR5jXOOKxjUK2eMULhprWt1qtjfOQetgmY9Gcx8XFWXtACE6OWcQ3y+q/h9Mp/PD6QhYK0FLct1/ZMUVRUq0BpjTLlxUnnlmGYppb5orSuzUFQEQzGFqk1ad0dEEDFfvA9wuT80wVspFsHQ+wSalef5XrIIegUPpkEpCH8jDcZqFf9cTk+DwhwrYkIknGICL/J+TaF2a4ypLOKKg0UKG4jjYo/apkRNnq5S7NoK05fCyFRVoYM9n6RacRosT/SHwdDXAZfhDT7QWl+HXOpBuIqUrcXS6yzLBnWPodjLNKVQEzEAN/OFBrqzsU4NFL8neA4vGx3FXv5BhUruru45PJ1Oz4ko+EbAXovFoivNUmIAAEIClUfMfNNU9xbnzmazD8wMw5+U5gV1UjyfpAGPaZ0Agip6+p1OZ9Tv979IStpijj/nHREtW2Ixw/2a+yzLunUv0/IdkhjweLG1Fv3/8xSjkDoBiv/yA4CgHdCW4ouR8l+P8DFFVADhfGb+aIxJatA+mwGPDa7rzCSCkzxd0oyp2/TZDCgBgVwLD4SyRN09Un9HtxlZJ1jgSDdsBIDiMP8aWycQjRneSAiEUPYdZLAC4rUqK5bfFpRSaG+97I+jVYD4+mHZAsdXn0dqXm5+3Pv1+ES+bJMvFouxNJ9LKd9YFnjugS9tXaMa8NKMk9ynBUCC0ibPaRmwyd6V2NYyQILSJs9pGbDJ3pXY1jJAgtImz2kZsMneldjWMkCC0ibP+QUrPp1f+SC5rAAAAABJRU5ErkJggg==',
				eyeCloseIcon: 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAYAAACqaXHeAAAAAXNSR0IArs4c6QAAA8lJREFUeF7tmMFV20AQhncE0jWkgpgKQiqIqSCmgpAKgg9ePU7AibfywU4FmApCKoipAFIBpoKQq0GavOGt8pZ9knZkUABp92iPdme++WdnJBAdX9Dx+IUH4BXQcQK+BDouAH8J+hLwJdBxAr4EOi4A3wV8CfgS6DgBXwIcASilBkEQvEfEH1LKS84zL8EmSZItIcRXRJzHcXxa5JNTAXqTC/3wjRBi+zVA0H7/FEJskO8AsD0ajeY2BCcAyj4AfDcefPEQ7OC171+klLPaAOgBpdQlALx/DRCKgkfEX1EU9YfDISXvwXIqgKwnk8nGcrmcWxDor0Kqz1X/dYO/Lw2us2UQAOBwNBodcfdpyi5Jkl0hxIm5f1Xmczs2AIcS5mEY7hRJrKmAzX2VUhMA2LPOOg/DcODyqRaA/ACl1AwAPlsH3gDATtFN2xQELXnKOrU7c32TUtpACt1YCQDtlCTJlHqsvSsiTqMoOnKRfyyU8Xh8gIgU5H2byxcinsZxTOXAWisBKGkz5oE35FzZ8MHyrMRoPB73syw7AYBexT7sy7k2AEbwZjYWQojDpwBBgSPigRCiXxD4tRDinfU7C0ItAFVtZrlckuwOAeBNgYOkCLo3TutMkcfHx721tbVPpKaijCPiHyHEbhzHZ0mS0JT30SqHHfqvSnFsALoNXpiO2G2GHA6CgCDYF+SD8hBCnAHAgmZ02zkAoEz3ELFfJXNEPIqiaJrfNSVt2jm1sgDQ5re3tzRX/7ttq3osKUVnrQpE7auAMg4AszRNp/v7+1ReD1YZhDRNPxTZ08MsALr+CMD94gwYZKclvIeIuyWlwYVwrbvLzNVdiiCUvQixAehNFxQEN3gzMrtl6kwW3RVlQNh9nTYwIbj8ZSkgz2YQBFtRFM1dWbCjUEpdGfV8LaXskZN3d3f2ACPW19cvaX/zBQwRF3Ecb3LlktuRcl2DGRtA3cNze10GV0b5sAYVWzVpmm6W1fGqvrFL4DEHKKWohU0MAM7WRLbWhxi6d4ZxHNP0+aSrcQXY/VlKyT5TKUXvF/ldcS6lLBqCHgWE7cwqp+j2+dt4lr4pDrh72S9dYRi+rXv/uM5qFEBB+6wlY/tzXFU7cwVa9n+jAMw6ptaXZdlWnYvMar+1n+dAaRQAOaDf3vpBEMxdLanIYT1eD7IsO6sDjxP8f+kCXEeey65xBTxXYNxzPQAuqbbaeQW0NbPcuLwCuKTaaucV0NbMcuPyCuCSaqudV0BbM8uNyyuAS6qtdl4Bbc0sNy6vAC6pttp5BbQ1s9y4/gJ+T/JQ23CaiAAAAABJRU5ErkJggg==',
			}
		},
		created() {
			uni.getStorage({
				key: this.historyStoreKey,
				success: res => {
					this.historyKeywords = JSON.parse(res.data);
				},
				fail: err => {

				}
			})
		},
		props: {
			placeholder: {
				value: String,
				default: '请输入搜索内容'
			},
			historyStoreKey: {
				value: String,
				default: 'searchHistoryList'
			},
			borderColor: {
				value: String,
				default: '#2979ff'
			},
			textColor: {
				value: String,
				default: '#2979ff'
			},
			searchName: {
				value: String,
				default: '搜索'
			},
			showMore: {
				value: Boolean,
				default: true
			},
			maxHistoryKeywordNumber: {
				value: Number,
				default: 10
			},
			hotKeywords: {
				value: Array,
				default: () => []
			}
		},
		methods: {
			search() {
				this.saveKeyword()
				this.$emit('search', this.keyword)
			},
			fillKeyword(keyword) {
				this.keyword = keyword
				this.saveKeyword()
				this.$emit('fillKeyword', this.keyword)
			},
			clearKeyword() {
				this.keyword = ''
				this.$emit('clearKeyword');
			},
			clearHistory() {
				uni.showModal({
					title: '提示',
					content: '确定清空历史搜索记录?',
					success: res => {
						if (res.confirm) {
							uni.removeStorageSync(this.historyStoreKey)
							this.historyKeywords = []
							this.$emit('clearHistory');
						}
					}
				})
			},
			saveKeyword() {
				if (!this.keyword) {
					return false;
				}

				uni.getStorage({
					key: this.historyStoreKey,
					success: res => {
						let oldData = JSON.parse(res.data)
						let findIndex = oldData.indexOf(this.keyword)
						findIndex !== -1 && oldData.splice(findIndex, 1)
						oldData.unshift(this.keyword)
						//最多10个纪录
						oldData.length > this.maxHistoryKeywordNumber && oldData.pop()
						this.historyKeywords = oldData
						uni.setStorage({
							key: this.historyStoreKey,
							data: JSON.stringify(oldData)
						});
					},
					fail: err => {
						let oldData = [this.keyword]
						uni.setStorage({
							key: this.historyStoreKey,
							data: JSON.stringify(oldData)
						});
						this.historyKeywords = oldData
					}
				})
			}
		}
	}
</script>

<style lang="scss" scoped>
	.search-panel {
		padding: 10rpx 20rpx;

		.search-panel-base {
			height: 100rpx;
			display: flex;
			justify-content: space-between;
			align-items: center;

			.left {
				flex: 1;
				height: 68rpx;
				background-color: #e8e8e8;
				border-radius: 30rpx;
				display: flex;
				align-items: center;
				justify-content: space-between;
				padding: 0 20rpx;
				border: 1rpx solid #f2f2f2;

				.search-icon {
					flex: none;
					width: 30rpx;
					height: 30rpx;
				}

				.search-input {
					flex: 1;
					height: 100%;
					padding: 0 20rpx;
				}

				.search-del-icon {
					width: 35rpx;
					height: 35rpx;
				}
			}

			.right {
				margin-left: 30rpx;
				height: 60rpx;
				line-height: 60rpx;

				.search-btn {
					border-radius: 30rpx;
					padding: 7rpx 26rpx;
				}
			}
		}

		.search-panel-more {
			.search-more-title {
				height: 80rpx;
				display: flex;
				align-items: center;
				justify-content: space-between;
				font-weight: 600;
				font-size: 32rpx;
				margin-bottom: 20rpx;

				.search-keyword-operate-icon {
					width: 40rpx;
					height: 40rpx;
				}
			}

			.search-null {
				text-align: center;
				color: #8799a3;
			}

			.search-keywords {
				display: flex;
				flex-wrap: wrap;
				justify-content: flex-start;

				.search-keyword {
					line-height: 60rpx;
					height: 64rpx;
					background-color: #e8e8e8;
					padding: 0rpx 30rpx;
					border-radius: 30rpx;
					margin: 10rpx 25rpx 15rpx 0;
					color: #8799a3;
				}
			}
		}
	}
</style>
