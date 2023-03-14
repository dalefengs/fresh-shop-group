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
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <el-tabs v-model="groupId" class="demo-tabs" @tab-change="tabsChange">
        <el-tab-pane v-for="(item) in accountGroupList" :label="item.nameCn" :name="item.ID" :key="item.ID"/>
      </el-tabs>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="用户信息" prop="user" width="200">
          <template #default="scope">
            <div class="table-multi-line">
              <span>用户名：{{ scope.row.user.userName }}</span><br>
              <span>手机号：{{ scope.row.user.phone }}</span><br>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="可用余额" prop="amount" width="120" />
        <el-table-column align="left" label="冻结余额" prop="freezeAmount" width="120" />
        <el-table-column align="left" label="锁仓金额" prop="lockAmount" width="120" />
        <el-table-column align="left" label="累计入账数额" prop="inAmount" width="120" />
        <el-table-column align="left" label="累计出账数额" prop="outAmount" width="120" />
        <el-table-column align="left" label="账户状态" prop="status" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.status, statusOptions) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="操作">
          <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateAccountFunc(scope.row)">变更
            </el-button>
          </template>
        </el-table-column>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="变更操作">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="账户状态" prop="status">
          <el-select v-model="formData.status" placeholder="请选择" style="width:100%" :clearable="true">
            <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
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
  name: 'Account',
}
</script>

<script setup>
import {
  createAccount,
  updateAccount,
  findAccount,
  getAccountList,
} from '@/api/account'
import { getAccountGroupList } from '@/api/accountGroup'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, filterDict } from '@/utils/format'
import { ElMessage } from 'element-plus'
import { ref } from 'vue'

// 自动化生成的字典（可能为空）以及字段
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
const groupId = ref()

const elFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const accountGroupList = ref({})

// 获取账户列表
const getAccountGroupTab = async() => {
  const res = await getAccountGroupList()
  accountGroupList.value = res.data.list
  if (accountGroupList.value.length === 0) {
    ElMessage.error('获取账户配置失败')
    return
  }
  groupId.value = accountGroupList.value[0].ID
}
getAccountGroupTab()

const tabsChange = (name) => {
  groupId.value = name
  getTableData()
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
  const table = await getAccountList({
    page: page.value,
    pageSize: pageSize.value,
    groupId: groupId.value, ...searchInfo.value,
  })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {
  statusOptions.value = await getDictFunc('status')
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateAccountFunc = async(row) => {
  const res = await findAccount({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reaccount
    dialogFormVisible.value = true
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    userId: 0,
    groupId: 0,
    amount: 0,
    freezeAmount: 0,
    lockAmount: 0,
    inAmount: 0,
    outAmount: 0,
    status: undefined,
  }
}
// 弹窗确定
const enterDialog = async() => {
  elFormRef.value?.validate(async(valid) => {
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
        message: '创建/更改成功',
      })
      closeDialog()
      getTableData()
    }
  })
}
</script>

<style>
</style>
