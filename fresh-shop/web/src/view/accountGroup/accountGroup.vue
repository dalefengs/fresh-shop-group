<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="英文名">
          <el-input v-model="searchInfo.nameEn" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="中文名">
          <el-input v-model="searchInfo.nameCn" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="searchInfo.status" clearable placeholder="请选择" @clear="()=>{searchInfo.status=undefined}">
            <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column align="left" label="编号" prop="ID" width="120" />
        <el-table-column align="left" label="英文名" prop="nameEn" width="120" />
        <el-table-column align="left" label="中文名" prop="nameCn" width="120" />
        <el-table-column align="left" label="小数位数" prop="places" width="120" />
        <el-table-column align="left" label="状态" prop="status" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.status,statusOptions) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="同步状态" prop="sync" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.sync,account_syncOptions) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作">
          <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateAccountGroupFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="edit" :disabled="scope.row.sync === 1" class="table-button" @click="syncAccountGroupOption(scope.row)">同步账户</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
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
        <el-form-item label="中文名:" prop="nameCn">
          <el-input v-model="formData.nameCn" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="英文名:" prop="nameEn">
          <el-input v-model="formData.nameEn" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="小数位数:" prop="places">
          <el-input-number v-model="formData.places" style="width:100%" :clearable="true" />
        </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-select v-model="formData.status" placeholder="请选择" style="width:100%" :clearable="true">
            <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <!--        <el-form-item label="同步状态:" prop="sync">
          <el-select v-model="formData.sync" placeholder="请选择" style="width:100%" :clearable="true">
            <el-option v-for="(item,key) in account_syncOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>-->
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
  name: 'AccountGroup'
}
</script>

<script setup>
import {
  createAccountGroup,
  deleteAccountGroup,
  updateAccountGroup,
  findAccountGroup,
  getAccountGroupList,
  syncAccountGroup
} from '@/api/accountGroup'

// 全量引入格式化工具 请按需保留
import { getDictFunc, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const statusOptions = ref([])
const account_syncOptions = ref([])
const formData = ref({
  nameEn: '',
  nameCn: '',
  places: 4,
  status: undefined,
})

// 验证规则
const rule = reactive({
  nameEn: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  nameCn: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  places: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  status: [{
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
  const table = await getAccountGroupList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
  statusOptions.value = await getDictFunc('status')
  account_syncOptions.value = await getDictFunc('account_sync')
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteAccountGroupFunc(row)
  })
}

// 删除行
const syncAccountGroupOption = (row) => {
  ElMessageBox.confirm('确定要同步账户吗?', '提示', {
    confirmButtonText: '开始同步',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    syncAccountGroupFunc(row)
  })
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateAccountGroupFunc = async(row) => {
  const res = await findAccountGroup({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reuserAccountGroup
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteAccountGroupFunc = async(row) => {
  const res = await deleteAccountGroup({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

//  同步账户
const syncAccountGroupFunc = async(row) => {
  const res = await syncAccountGroup({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '正在同步中，请耐心等待...'
    })
    getTableData()
  }
}

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
    nameEn: '',
    nameCn: '',
    places: 0,
    status: undefined,
    sync: undefined,
  }
}
// 弹窗确定
const enterDialog = async() => {
     elFormRef.value?.validate(async(valid) => {
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
         closeDialog()
         getTableData()
       }
     })
}
</script>

<style>
</style>
