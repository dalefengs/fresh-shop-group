<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="商品名称">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="所属区域" prop="goodsArea">
          <el-select v-model="searchInfo.goodsArea" clearable placeholder="请选择" @clear="()=>{searchInfo.goodsArea=undefined}">
            <el-option v-for="(item,key) in goodsAreaOptions" :key="key" :label="item.label" :value="item.value" />
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
        <el-table-column align="left" label="商品图片" prop="images" width="100">
          <template #default="scope">
            <img v-if="scope.row.images[0]" :src="(scope.row.images[0].url && scope.row.images[0].url.slice(0, 4) !== 'http') ? path + scope.row.images[0].url:scope.row.images[0].url" style="width: 100%" @click="handlePictureCardPreview(scope.row.images[0].url)">
            <span v-else>暂无图片</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="基本信息" prop="name" width="220">
          <template #default="scope">
            <div class="table-multi-line">
              <span>编号：{{ scope.row.ID }}</span><br>
              <div style="display:inline-block;">
                商品名称：<span style="color: #4d70ff">{{ scope.row.name }}</span>
              </div><br>
              <span>所属分类：{{ scope.row.category.title }}</span><br>
              <span>所属品牌：{{ scope.row.brand.name }}</span><br>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="商品类型" prop="goodsArea" width="160">
          <template #default="scope">
            <div class="table-multi-line">
              <span>商品区域：{{ filterDict(scope.row.goodsArea, goodsAreaOptions) }}</span><br>
              <span>规格类型：{{ filterDict(scope.row.specType,specTypeOptions) }}</span><br>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="详情信息" prop="name" width="250">
          <template #default="scope">
            <div class="table-multi-line">
              <div >
                商品价格：<span style="color: #f56c6c; font-weight: bold">{{ scope.row.costPrice }} 元</span><el-divider direction="vertical" />
                优惠价格: <span v-if="scope.row.price > 0" style="color: #f56c6c; font-weight: bold">{{ scope.row.price }} 元</span>
                <span v-else style="color: #f56c6c; font-weight: bold">无优惠</span>
              </div>
              <span>商品库存：{{ scope.row.price }}</span><el-divider direction="vertical" />
              <span>商品单位：{{ scope.row.unit }}</span><br>
              <span>最低购买数量：{{ scope.row.minCount }} {{ scope.row.unit }}</span><el-divider direction="vertical" />
              <span>总销量：{{ scope.row.sale }}</span><br>
              <span>商品重量：{{ scope.row.weight }}g</span><br>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="状态信息" prop="status" width="120">
          <template #default="scope">
            <div class="table-multi-line">
              <span>状态：{{ filterDict(scope.row.status, goodsStatusOptions) }}</span><br>
              <!--              <span>是否首页：{{ filterDict(scope.row.isFirst,whetherOptions) }}</span><br>-->
              <span>是否热销：{{ filterDict(scope.row.isHot,whetherOptions) }}</span><br>
              <span>是否上新：{{ filterDict(scope.row.isNew,whetherOptions) }}</span><br>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="其他">
          <template #default="scope">
            <div class="table-multi-line">
              <span>排序：{{ scope.row.sort }}</span><br>
              <span>日期：{{ formatDate(scope.row.CreatedAt) }}</span><br>
            </div>
          </template>
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
    <el-dialog v-model="dialogImgVisible">
      <img style="width: 100%" :src="dialogImageUrl" alt="Preview Image">
    </el-dialog>
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
import { getDictFunc, formatDate, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const path = ref(import.meta.env.VITE_BASE_API + '/')

// 自动化生成的字典（可能为空）以及字段
const specTypeOptions = ref([])
const whetherOptions = ref([])
const goodsAreaOptions = ref([])
const goodsStatusOptions = ref([])

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
  specTypeOptions.value = await getDictFunc('spec_type')
  whetherOptions.value = await getDictFunc('whether')
  goodsAreaOptions.value = await getDictFunc('goods_area')
  goodsStatusOptions.value = await getDictFunc('goods_status')
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

const dialogImageUrl = ref('')
const dialogImgVisible = ref(false)
const handlePictureCardPreview = (url) => {
  console.log(url)
  if (url.slice(0, 4) !== 'http') {
    url = path.value + '/' + url
  }
  dialogImageUrl.value = url
  dialogImgVisible.value = true
}

</script>

<style>
</style>
