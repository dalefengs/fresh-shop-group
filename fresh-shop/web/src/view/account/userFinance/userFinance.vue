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
        <el-form-item label="创建时间">
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
      <el-tabs v-model="groupNameEn" class="demo-tabs" @tab-change="tabsChange">
        <el-tab-pane v-for="(item) in accountGroupList" :key="item.ID" :label="item.nameCn" :name="item.nameEn" />
      </el-tabs>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
      >
        <el-table-column align="left" label="编号" prop="ID" width="120" />
        <el-table-column align="left" label="用户信息" prop="user" width="200">
          <template #default="scope">
            <div class="table-multi-line">
              <span>用户名：{{ scope.row.user.userName }}</span><br>
              <span>手机号：{{ scope.row.user.phone }}</span><br>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="流水类型" prop="typeId" width="120">
          <template #default="scope">
            {{ scope.row.financeType.name }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作类型" prop="optionType" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.optionType,finance_option_typeOptions) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="流水信息" width="200">
          <template #default="scope">
            <div class="table-multi-line">
              <span>变动数额：{{ scope.row.amount }}</span><br>
              <span>变动后余额：{{ scope.row.balance }}</span><br>
              <span>手续费：{{ scope.row.feeAmount }}</span><br>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="来源信息" width="120">
          <template #default="scope">
            <div class="table-multi-line">
              <span>来源Id: {{ scope.row.fromId ? scope.row.fromId : '无' }}</span><br>
              <span>来源用户Id: {{ scope.row.fromUserId ? scope.row.fromUserId : '无' }}</span><br>
              <span>来源用户名: {{ scope.row.fromName ? scope.row.fromName : '无' }}</span><br>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="备注" prop="remarks" width="120" />
        <el-table-column align="left" label="创建日期">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
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
  </div>
</template>

<script>
export default {
  name: 'UserFinanceCash'
}
</script>

<script setup>
import {
  getUserFinanceCashList
} from '@/api/userFinance'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, filterDict } from '@/utils/format'
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
import { getAccountGroupList } from '@/api/accountGroup'

// 自动化生成的字典（可能为空）以及字段
const finance_option_typeOptions = ref([])

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const groupNameEn = ref()
const accountGroupList = ref({})

// 获取账户列表
const getAccountGroupTab = async() => {
  const res = await getAccountGroupList({ status: 1 })
  accountGroupList.value = res.data.list
  if (accountGroupList.value.length === 0) {
    ElMessage.error('获取账户配置失败')
    return
  }
  groupNameEn.value = accountGroupList.value[0].nameEn
}
getAccountGroupTab()

const tabsChange = (name) => {
  groupNameEn.value = name
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
  const table = await getUserFinanceCashList({ page: page.value, pageSize: pageSize.value, groupNameEn: groupNameEn.value, ...searchInfo.value })
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
  finance_option_typeOptions.value = await getDictFunc('finance_option_type')
}

// 获取需要的字典 可能为空 按需保留
setOptions()

</script>

<style>
</style>
