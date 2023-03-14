<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户id:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="币种id:" prop="groupId">
          <el-input v-model.number="formData.groupId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="可用余额:" prop="amount">
          <el-input-number v-model="formData.amount" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="冻结余额:" prop="freezeAmount">
          <el-input-number v-model="formData.freezeAmount" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="锁仓金额:" prop="lockAmount">
          <el-input-number v-model="formData.lockAmount" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="累计入账数额:" prop="inAmount">
          <el-input-number v-model="formData.inAmount" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="累计出账数额:" prop="outAmount">
          <el-input-number v-model="formData.outAmount" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="账户状态(0禁用 1启用):" prop="status">
          <el-select v-model="formData.status" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Account'
}
</script>

<script setup>
import {
  createAccount,
  updateAccount,
  findAccount
} from '@/api/account'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const statusOptions = ref([])
const formData = ref({
            userId: 0,
            groupId: 0,
            amount: 0,
            freezeAmount: 0,
            lockAmount: 0,
            inAmount: 0,
            outAmount: 0,
            status: undefined,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findAccount({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reaccount
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    statusOptions.value = await getDictFunc('status')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createAccount(formData.value)
               break
             case 'update':
               res = await updateAccount(formData.value)
               break
             default:
               res = await createAccount(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
