<template>
    <pageWrapper>
        <liu-chooseAddress ref="chooseAddress" @submit="submitAddress" @detele="deteleAddress">
        </liu-chooseAddress>
        <u-toast ref="toast"></u-toast>
    </pageWrapper>
</template>

<script>
import {getToken} from '@/store/storage.js'
import {getAddressInfo, createAddress, updateAddress, deleteAddress} from '@/api/address';
import ChooseAddress from '@/uni_modules/liu-chooseAddress/components/liu-chooseAddress/liu-chooseAddress.vue'

export default {
    components: {
        ChooseAddress,
    },
    data() {
        return {
            token: '',
            id: 0,
            address: {}
        }
    },
    onLoad(options) {
        this.token = getToken()
        if (!this.token) {
            this.$message(this.$refs.toast).error("请请先登录").then(() => {
                uni.redirectTo({
                    url: '/pages/my/my'
                })
            })
            return
        }
        this.id = options.id
        if (this.id) {
            this.getAddressInfoData()
        }
    },
    methods: {
        async getAddressInfoData() {
            const res = await getAddressInfo({
                ID: this.id
            })
            if (res.code !== 0) {
                this.$message(this.$refs.toast).error(res.msg)
                return false
            }
            this.address = res.data.reuserAddress
            this.$refs.chooseAddress.setData(this.address)
        },
        //添加地址成功回调
        async submitAddress(data) {
            //拿到地址信息、表单数据、是否默认地址等
            console.log('添加调地址信息：', data)
            if (!data.longitude || !data.latitude) {
                this.$message(this.$refs.toast).error("请选择地址")
                return false
            }
            if (!data.address) {
                this.$message(this.$refs.toast).error("请选择地址")
                return false
            }
            if (!data.detail) {
                this.$message(this.$refs.toast).error("请输入详细地址")
                return false
            }
            if (!data.name) {
                this.$message(this.$refs.toast).error("请输入姓名")
                return false
            }
            if (!data.mobile) {
                this.$message(this.$refs.toast).error("请输入手机号")
                return false
            }
            let res
            if (this.id) {
                data.ID = parseInt(this.id)
                res = await updateAddress(data, this.$refs.toast)
            } else {
                res = await createAddress(data, this.$refs.toast)
            }
            if (res.code !== 0) {
                this.$message(this.$refs.toast).error(res.msg)
                return false
            }
            this.$message(this.$refs.toast).success("操作成功").then(() => {
                uni.navigateBack()
            })
        },
        //删除地址回调
        async deteleAddress() {
            const res = await deleteAddress({
                ID: this.id
            }, this.$refs.toast)
            if (res.code !== 0) {
                this.$message(this.$refs.toast).error(res.msg)
                return false
            }
            this.$message(this.$refs.toast).success("删除成功").then(() => {
                uni.navigateBack()
            })
        },
        //地址回显调用此方法
        editAddress() {
            let obj = {
                detailAddress: '',
                userName: '',
                userPhone: '',
                lableName: '',
                sexType: '',
                title: '',
                address: '',
                latitude: '',
                longitude: '',
                agreeState: ''
            }
            this.$refs.chooseAddress.setData(obj)
        }
    }
}
</script>

<style scoped>

</style>
