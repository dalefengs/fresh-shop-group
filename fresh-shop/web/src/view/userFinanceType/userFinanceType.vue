<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
      </div>
      <!-- 由于此处菜单跟左侧列表一一对应所以不需要分页 pageSize默认999 -->
      <el-table :data="tableData" row-key="ID">
        <el-table-column align="left" label="ID" width="100" prop="ID" />
        <el-table-column align="left" label="流水类型" prop="name" width="220" />
        <el-table-column align="left" label="备注" prop="remarks" width="220" />
        <el-table-column align="left" label="创建日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="操作">
          <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateUserFinanceTypeFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="type === 'create' ? '新增操作' : '修改操作'">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="父级类型" prop="status">
          <el-select v-model="formData.parentId" placeholder="请选择父级类型" style="width:100%" :clearable="true">
            <el-option label="顶级类型" :value="0" />
            <el-option v-for="(item,key) in topFinanceList" :key="key" :label="item.name" :value="item.ID" />
          </el-select>
        </el-form-item>
        <el-form-item label="流水类型:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
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
  name: 'UserFinanceType'
}
</script>

<script setup>
import {
  createUserFinanceType,
  deleteUserFinanceType,
  updateUserFinanceType,
  findUserFinanceType,
  getUserFinanceTypeList, getUserFinanceTypeListAll,
} from '@/api/userFinanceType'

// 全量引入格式化工具 请按需保留
import { formatDate } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  parentId: 0,
  name: '',
  remarks: '',
})

// 验证规则
const rule = reactive({
  name: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
})

const elFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(30)
const tableData = ref([])
const searchInfo = ref({})
const topFinanceList = ref({})
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
  const table = await getUserFinanceTypeList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteUserFinanceTypeFunc(row)
  })
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateUserFinanceTypeFunc = async(row) => {
  const res = await findUserFinanceType({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reuserFinanceType
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteUserFinanceTypeFunc = async(row) => {
  const res = await deleteUserFinanceType({ ID: row.ID })
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

// 获取顶级流水类型列表
const getTopLevelTypeList = async()=> {
  const res = await getUserFinanceTypeListAll()
  if (res.code !== 0) {
    ElMessage.error('获取顶级流水类型失败')
    return
  }
  topFinanceList.value = res.data
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  getTopLevelTypeList()
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    parentId: 0,
    name: '',
    remarks: '',
  }
}
// 弹窗确定
const enterDialog = async() => {
     elFormRef.value?.validate(async(valid) => {
       if (!valid) return
       let res
       switch (type.value) {
         case 'create':
           res = await createUserFinanceType(formData.value)
           break
         case 'update':
           res = await updateUserFinanceType(formData.value)
           break
         default:
           res = await createUserFinanceType(formData.value)
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
