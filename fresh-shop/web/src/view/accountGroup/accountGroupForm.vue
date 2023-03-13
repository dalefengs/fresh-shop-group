<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="英文名:" prop="nameEn">
          <el-input v-model="formData.nameEn" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="中文名:" prop="nameCn">
          <el-input v-model="formData.nameCn" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="小数点位数:" prop="places">
          <el-input-number v-model="formData.places" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-select v-model="formData.status" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="同步状态:" prop="sync">
          <el-select v-model="formData.sync" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in account_syncOptions" :key="key" :label="item.label" :value="item.value" />
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
  name: 'AccountGroup'
}
</script>

<script setup>
import {
  createAccountGroup,
  updateAccountGroup,
  findAccountGroup
} from '@/api/accountGroup'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const statusOptions = ref([])
const account_syncOptions = ref([])
const formData = ref({
            nameEn: '',
            nameCn: '',
            places: 0,
            status: undefined,
            sync: undefined,
        })
// 验证规则
const rule = reactive({
               nameEn : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               nameCn : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               places : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findAccountGroup({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reuserAccountGroup
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    statusOptions.value = await getDictFunc('status')
    account_syncOptions.value = await getDictFunc('account_sync')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createAccountGroup(formData.value)
               break
             case 'update':
               res = await updateAccountGroup(formData.value)
               break
             default:
               res = await createAccountGroup(formData.value)
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
