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
        <el-button type="primary" icon="Download" @click="importGoodsShowClick">Excel 导入</el-button>
        <el-button type="primary" icon="Upload" @click="exportGoodsShowClick">导出所有商品</el-button>
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
              <span>产地：{{ scope.row.origin ? scope.row.origin : '无' }}</span><br>
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
              <div v-if="scope.row.goodsArea === 1">
                商品价格：<span style="color: #f56c6c; font-weight: bold">{{ scope.row.costPrice }} 积分</span>
              </div>
              <div v-else>
                商品价格：<span style="color: #f56c6c; font-weight: bold">{{ scope.row.costPrice }} 元</span><el-divider direction="vertical" />
                优惠价格: <span v-if="scope.row.price > 0" style="color: #f56c6c; font-weight: bold">{{ scope.row.price }} 元</span>
                <span v-else style="color: #f56c6c; font-weight: bold">无优惠</span>
              </div>
              <span>商品库存：{{ scope.row.store }}</span><el-divider direction="vertical" />
              <span>商品单位：{{ scope.row.unit }}</span><br>
              <span>最低购买数量：{{ scope.row.minCount }} {{ scope.row.unit }}</span><el-divider direction="vertical" />
              <span>总销量：{{ scope.row.sale }}</span><br>
              <span>商品重量：{{ scope.row.weight }}g</span><el-divider direction="vertical" />
              <span /><br>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="状态信息" prop="status" width="120">
          <template #default="scope">
            <div class="table-multi-line">
              <span>状态：{{ filterDict(scope.row.status, goodsStatusOptions) }}</span><br>
              <!--              <span>是否首页：{{ filterDict(scope.row.isFirst,whetherOptions) }}</span><br>-->
              <span v-if="scope.row.goodsArea === 0">是否热销：{{ filterDict(scope.row.isHot,whetherOptions) }}</span><br>
              <span v-if="scope.row.goodsArea === 0">是否上新：{{ filterDict(scope.row.isNew,whetherOptions) }}</span><br>
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

    <el-drawer v-model="importGoodsShow">
      <template #header>
        <h4>Excel 导入商品</h4>
      </template>
      <template #default>
        <pre class="notice">
注意事项：
1. *号是必填选项
2. 图片必须在表格内
3. 表格第一、第二行不允许修改
4. 请勿使用 WPS 进行编辑 Excel 表格
        </pre>
        <div>
          <el-button icon="Download" type="primary" @click="importGoodsDownload">下载导入模版</el-button>
        </div>
        <div style="margin-top: 20px">
          <ExcelImport
            class="upload-btn"
            @on-success="importGoodsSuccess"
          />
        </div>
        <div v-if="importGoodsMsg">
          <div v-for="(msg, index) in importGoodsMsg" :key="index" style="margin-top: 5px">
            <el-alert :title="msg" type="error" />
          </div>
        </div>
      </template>
      <template #footer>
        <div style="flex: auto">
          <el-button @click="importGoodsShowClose">取消</el-button>
        </div>
      </template>
    </el-drawer>
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
  getGoodsList,
  exportGoods
} from '@/api/goods'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import ExcelImport from '@/components/upload/excelImport.vue'

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

// 导入商品
const importGoodsShow = ref(false)
const importGoodsShowClick = () => {
  importGoodsShow.value = true
}
const importGoodsShowClose = () => {
  importGoodsShow.value = false
}
const importGoodsDownload = () => {
  console.log('importGoodsDownload')
  const url = 'goapi/uploads/file/excel/goodsImportTemplate.xlsx'
  window.location.href = url
}

const importGoodsMsg = ref([])
const importGoodsSuccess = (data) => {
  console.log('importGoodsSuccess', data)
  if (data.code !== 0) {
    importGoodsMsg.value.push(data.msg)
    return
  }
  importGoodsMsg.value = []
  ElMessage({
    type: 'success',
    message: '商品导入成功'
  })
  importGoodsShow.value = true
  getTableData()
}

const exportGoodsShowClick = async() => {
  ElMessage({
    type: 'warning',
    message: '正在导出，请耐心等待，不要切换页面。',
    duration: 50000,
  })
  const res = await exportGoods({})
  if (res.code && res.code !== 0) {
    ElMessage({
      type: 'error',
      message: res.msg
    })
  }
  ElMessage({
    type: 'success',
    message: '导出完成请保存文件!'
  })
  fileDownload(res, '启运冻品-商品信息列表.xlsx')
}

const fileDownload = (res, filename) => {
  const blob = new Blob([res.data]) // 将返回的数据通过Blob的构造方法，创建Blob对象
  if ('msSaveOrOpenBlob' in navigator) {
    window.navigator.msSaveOrOpenBlob(blob, filename) // 针对浏览器
  } else {
    const elink = document.createElement('a') // 创建a标签
    elink.download = filename
    elink.style.display = 'none'
    // 创建一个指向blob的url，这里就是点击可以下载文件的根结
    elink.href = URL.createObjectURL(blob)
    document.body.appendChild(elink)
    elink.click()
    URL.revokeObjectURL(elink.href) // 移除链接
    document.body.removeChild(elink) // 移除a标签
  }
}

</script>

<style lang="scss">
.notice {
  color: #ed1515;
  font-size: 18px;
}
</style>
