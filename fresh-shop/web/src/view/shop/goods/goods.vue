<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="商品名称">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="所属区域" prop="goodsArea">
          <el-select v-model="searchInfo.goodsArea" clearable placeholder="请选择" @clear="()=>{searchInfo.goodsArea=undefined}">
            <el-option v-for="(item,key) in goods_areaOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
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
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openGoodsFrom('新增商品')">新增</el-button>
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
        @sort-change="sortChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="编号" prop="ID" width="90" />
        <el-table-column align="left" label="商品名称" prop="name" width="120" />
        <el-table-column align="left" label="分类id" prop="categoryId" width="120" />
        <el-table-column align="left" label="品牌Id" prop="brandId" width="120" />
        <el-table-column align="left" label="所属区域" prop="goodsArea" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.goodsArea,goods_areaOptions) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="规格类型" prop="specType" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.specType,spec_typeOptions) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="商品单位" prop="unit" width="120" />
        <el-table-column sortable align="left" label="商品价格" prop="price" width="120" />
        <el-table-column align="left" label="最低购买数量" prop="minCount" width="120" />
        <el-table-column align="left" label="商品重量（g）" prop="weight" width="120" />
        <el-table-column sortable align="left" label="库存" prop="store" width="120" />
        <el-table-column sortable align="left" label="总销量" prop="sale" width="120" />
        <el-table-column align="left" label="排序" prop="soft" width="120" />
        <el-table-column align="left" label="状态(0 下架 1上架 )" prop="status" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.status) }}</template>
        </el-table-column>
        <el-table-column sortable align="left" label="是否首页(0否 1是)" prop="isFrist" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.isFrist,whetherOptions) }}
          </template>
        </el-table-column>
        <el-table-column sortable align="left" label="是否热销(0否 1是)" prop="isHot" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.isHot,whetherOptions) }}
          </template>
        </el-table-column>
        <el-table-column sortable align="left" label="是否上新(0否 1是)" prop="isNew" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.isNew,whetherOptions) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="创建日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="操作">
          <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="openGoodsFrom(scope.row.ID)">变更</el-button>
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
  </div>
</template>

<script>
export default {
  name: 'Goods'
}
</script>

<script setup>
import {
  deleteGoods,
  deleteGoodsByIds,
  getGoodsList
} from '@/api/goods'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import { useRouter } from 'vue-router'

// 自动化生成的字典（可能为空）以及字段
const spec_typeOptions = ref([])
const whetherOptions = ref([])
const goods_areaOptions = ref([])

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 排序
const sortChange = ({ prop, order }) => {
  searchInfo.value.sort = prop
  searchInfo.value.order = order
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
  if (searchInfo.value.status === '') {
    searchInfo.value.status = null
  }
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
  const table = await getGoodsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
  spec_typeOptions.value = await getDictFunc('spec_type')
  whetherOptions.value = await getDictFunc('whether')
  goods_areaOptions.value = await getDictFunc('goods_area')
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
    deleteGoodsFunc(row)
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
  const res = await deleteGoodsByIds({ ids })
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

// 删除行
const deleteGoodsFunc = async(row) => {
  const res = await deleteGoods({ ID: row.ID })
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

const router = useRouter()
// 打开商品表单页
const openGoodsFrom = (id) => {
  console.log(id)
  router.push({ name: 'goodsFrom', query: { id: id }})
}

</script>

<style>
</style>
