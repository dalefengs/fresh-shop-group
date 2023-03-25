<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
            <el-button type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="编号" prop="ID" width="80" />
        <el-table-column align="left" label="图片" prop="imgUrl" width="180">
          <template #default="scope">
            <img :src="(scope.row.imgUrl && scope.row.imgUrl.slice(0, 4) !== 'http') ? path + scope.row.imgUrl:scope.row.imgUrl" style="width: 100px;height: 100px" alt="">
          </template>
        </el-table-column>
        <el-table-column align="left" label="分类名" prop="title" width="120" />
        <el-table-column align="left" label="是否首页" prop="isFirst" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.isFirst,whetherOptions) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="关联品牌" min-width="120">
          <template #default="scope">
            <el-cascader
              v-model="scope.row.brandIds"
              :options="categoryOptions"
              :show-all-levels="false"
              :props="{ multiple:true, checkStrictly: true, label:'name', value:'ID', disabled:'disabled', emitPath:false}"
              :clearable="false"
              @visible-change="(flag)=>{changeBrands(scope.row,flag,0)}"
              @remove-tag="(removeBrandId)=>{changeBrands(scope.row,false, removeBrandId)}"
            />
          </template>
        </el-table-column>
        <el-table-column align="left" label="排序" prop="sort" width="120" />
        <el-table-column align="left" label="创建日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="操作">
          <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateCategoryFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 20, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" class="upload-img-dialog" :before-close="closeDialog" :title="type === 'create' ? '新增操作' : '修改操作'">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="分类名:" prop="title">
          <el-input v-model="formData.title" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="分类图片" label-width="80px">
          <div style="display:inline-block" @click="openImgUrlChange">
            <img
              v-if="formData.imgUrl"
              alt="头像"
              class="header-img-box"
              :src="(formData.imgUrl && formData.imgUrl.slice(0, 4) !== 'http') ? path + formData.imgUrl:formData.imgUrl"
            >
            <div v-else class="header-img-box">从媒体库选择</div>
          </div>
        </el-form-item>
        <el-form-item label="是否首页:" prop="isFirst">
          <el-radio-group v-model="formData.isFirst" class="ml-4">
            <el-radio v-for="(item,key) in whetherOptions" :key="key" :label="item.value" size="large">{{ item.label }}</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="排序:" prop="sort">
          <el-input v-model.number="formData.sort" :clearable="true" placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
    <ChooseImg ref="chooseImg" :target="formData" :target-key="`imgUrl`" />
  </div>
</template>

<script>
export default {
  name: 'Category'
}
</script>

<script setup>
import {
  createCategory,
  deleteCategory,
  deleteCategoryByIds,
  updateCategory,
  findCategory,
  getCategoryList
} from '@/api/category'
import ChooseImg from '@/components/chooseImg/index.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, nextTick, watch } from 'vue'
import { getBrandListAll, createBrandCategory } from '@/api/brand'

const path = ref(import.meta.env.VITE_BASE_API + '/')
// 自动化生成的字典（可能为空）以及字段
const whetherOptions = ref([])
const formData = ref({
  pid: 0,
  title: '',
  imgUrl: '',
  sort: 50,
  isFirst: 0,
})

// 验证规则
const rule = reactive({
  title: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
})

const elFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(20)
const tableData = ref([])
const searchInfo = ref({})
const chooseImg = ref(null)

watch(() => tableData.value, () => {
  setBrandIds()
})

// 为每行设置 ids
const setBrandIds = () => {
  tableData.value && tableData.value.forEach(item => {
    console.log(item)
    item.brandIds = item.brands.map(i => {
      return i.ID
    })
  })
  console.log(tableData.value)
}

const categoryOptions = ref([])
// 获取所有品牌列表
const setCategoryOptions = async() => {
  // 获取到所有品牌列表
  const res = await getBrandListAll()
  if (res.code !== 0) {
    ElMessage.error('获取品牌列表失败')
    return
  }
  const list = []
  res.data.forEach(item => {
    list.push({
      name: item.name,
      ID: item.ID
    })
  })
  categoryOptions.value = list
}

const tempBrand = {}
const changeBrands = async(row, flag, removeBrandId) => {
  console.log(flag, removeBrandId)
  if (flag) {
    if (!removeBrandId) {
      // 设置缓存，修改失败时恢复
      tempBrand[row.ID] = [...row.brandIds]
    }
    return
  }
  await nextTick()
  const res = await createBrandCategory({
    categoryId: row.ID,
    brandIds: row.brandIds
  })
  if (res.code === 0) {
    ElMessage({ type: 'success', message: '品牌设置成功' })
  } else {
    if (!removeBrandId) {
      row.brandIds = [...tempBrand[row.ID]]
      delete tempBrand[row.ID]
    } else {
      // row.brandIds = [removeBrandId, ...row.brandIds]
    }
  }
}

const openImgUrlChange = () => {
  chooseImg.value.open()
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
  const table = await getCategoryList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
  whetherOptions.value = await getDictFunc('whether')
  await setCategoryOptions()
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
    deleteCategoryFunc(row)
  })
}

// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据'
    })
    return
  }
  multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
  const res = await deleteCategoryByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateCategoryFunc = async(row) => {
  const res = await findCategory({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.recategory
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteCategoryFunc = async(row) => {
  const res = await deleteCategory({ ID: row.ID })
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
    pid: 0,
    title: '',
    imgUrl: '',
    sort: 50,
    isFirst: 0,
  }
}
// 弹窗确定
const enterDialog = async() => {
     elFormRef.value?.validate(async(valid) => {
       if (!valid) return
       let res
       switch (type.value) {
         case 'create':
           res = await createCategory(formData.value)
           break
         case 'update':
           res = await updateCategory(formData.value)
           break
         default:
           res = await createCategory(formData.value)
           break
       }
       if (res.code === 0) {
         ElMessage({
           type: 'success',
           message: '操作成功'
         })
         closeDialog()
         getTableData()
       }
     })
}

</script>

<style lang="scss">
.upload-img-dialog {
  .header-img-box {
    width: 200px;
    height: 200px;
    border: 1px dashed #ccc;
    border-radius: 20px;
    text-align: center;
    line-height: 200px;
    cursor: pointer;
  }

  .avatar-uploader .el-upload:hover {
    border-color: #409eff;
  }

  .avatar-uploader-icon {
    border: 1px dashed #d9d9d9 !important;
    border-radius: 6px;
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
  }

  .avatar {
    width: 178px;
    height: 178px;
    display: block;
  }
}
</style>
