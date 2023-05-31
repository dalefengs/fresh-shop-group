<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="用户名">
          <el-input v-model="searchInfo.username" placeholder="用户名" />
        </el-form-item>
        <el-form-item label="手机号">
          <el-input v-model="searchInfo.phone" placeholder="手机号" />
        </el-form-item>
        <el-form-item label="充值时间">
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间" />
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">增减账户</el-button>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
      >
        <el-table-column align="left" label="编号" prop="ID" width="120" />
        <el-table-column align="left" label="用户信息" prop="username" width="220">
          <template #default="scope">
            <div class="table-multi-line">
              <span>用户名：{{ scope.row.username }}</span> <br>
              <span>手机号：{{ scope.row.user.phone }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="账户类型" prop="groupId" width="120" >
          <template #default="scope">
            {{ groupComputed(scope.row.groupId) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="增减数额" prop="amount" width="120" />
        <el-table-column align="left" label="当前余额" prop="balance" width="120" />
        <el-table-column align="left" label="操作员" prop="adminName" width="120" />
        <el-table-column align="left" label="充值日期" width="280">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="备注" prop="remarks" />
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="增减账户">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户名:" prop="username">
          <el-input v-model="formData.username" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="账户类型" prop="groupId">
          <el-select v-model="formData.groupId" placeholder="请选择账户类型" style="width:100%" :clearable="true">
            <el-option v-for="(item, key) in groupOptions" :key="key" :label="item.nameCn" :value="item.ID" />
          </el-select>
        </el-form-item>
        <el-form-item label="增减数额:" prop="amount">
          <el-input-number v-model="formData.amount" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="备注:" prop="remarks">
          <el-input v-model="formData.remarks" :clearable="true" placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'SysRecharge'
}
</script>

<script setup>
import {
  getSysRechargeList,
  createSysRecharge,
} from '@/api/sysRecharge'

// 全量引入格式化工具 请按需保留
import { formatDate } from '@/utils/format'
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
import { getAccountGroupListAll } from '@/api/accountGroup'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  userId: 0,
  username: '',
  groupId: '',
  adminId: 0,
  adminName: '',
  remarks: '',
  amount: 0,
})

// 验证规则
const rule = reactive({
  amount: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
})

const elFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const groupOptions = ref({})

// 获取账户分组列表
const getGroupOptions = async() => {
  const res = await getAccountGroupListAll({ status: 1 })
  if (res.code !== 0) {
    ElMessage.error('获取账户配置异常')
    return
  }
  groupOptions.value = res.data
}
getGroupOptions()

const groupComputed = (groupId) => {
  for (const key in groupOptions.value) {
    console.log()
    if (groupOptions.value[key].ID === groupId) {
      return groupOptions.value[key].nameCn
    }
  }
  return ''
}

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getSysRechargeList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    userId: 0,
    username: '',
    groupId: null,
    remarks: '',
    amount: 0,
  }
}
// 弹窗确定
const enterDialog = async() => {
     elFormRef.value?.validate(async(valid) => {
       if (!valid) return
       const res = await createSysRecharge(formData.value)
       if (res.code === 0) {
         ElMessage({
           type: 'success',
           message: '增减账户成功'
         })
         closeDialog()
         getTableData()
       }
     })
}
</script>

<style>
</style>
