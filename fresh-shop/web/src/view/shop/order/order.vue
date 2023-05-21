<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="订单编号">
          <el-input v-model="searchInfo.orderSn" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="收货人姓名">
          <el-input v-model="searchInfo.shipmentName" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="收货人手机号">
          <el-input v-model="searchInfo.shipmentMobile" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="收获人地址">
          <el-input v-model="searchInfo.shipingAddress" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="支付方式" prop="payment">
          <el-select v-model="searchInfo.payment" clearable placeholder="请选择" @clear="()=>{searchInfo.payment=undefined}">
            <el-option v-for="(item,key) in paymentOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="订单状态" prop="status">
          <el-select v-model="searchInfo.status" clearable placeholder="请选择" @clear="()=>{searchInfo.status=undefined}">
            <el-option v-for="(item,key) in orderStatusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="发货时间">
          <el-date-picker v-model="searchInfo.startShipmentTime" type="datetime" placeholder="搜索条件（起）" />
          —
          <el-date-picker v-model="searchInfo.endShipmentTime" type="datetime" placeholder="搜索条件（止）" />

        </el-form-item>
        <el-form-item label="收货时间">

          <el-date-picker v-model="searchInfo.startReceiveTime" type="datetime" placeholder="搜索条件（起）" />
          —
          <el-date-picker v-model="searchInfo.endReceiveTime" type="datetime" placeholder="搜索条件（止）" />

        </el-form-item>
        <el-form-item label="取消时间">

          <el-date-picker v-model="searchInfo.startCancelTime" type="datetime" placeholder="搜索条件（起）" />
          —
          <el-date-picker v-model="searchInfo.endCancelTime" type="datetime" placeholder="搜索条件（止）" />

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
        <el-table-column align="left" label="编号" prop="ID" width="90" />
        <el-table-column align="left" label="用户id" prop="userId" width="120" />
        <el-table-column align="left" label="商品id" prop="goodsId" width="120" />
        <el-table-column align="left" label="订单编号" prop="orderSn" width="120" />
        <el-table-column align="left" label="所属区域" prop="goodsArea" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.goodsArea,goods_areaOptions) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="收货人姓名" prop="shipmentName" width="120" />
        <el-table-column align="left" label="收货人手机号" prop="shipmentMobile" width="120" />
        <el-table-column align="left" label="收货人地区编码" prop="shipmentArea" width="120" />
        <el-table-column align="left" label="收获人地址" prop="shipingAddress" width="120" />
        <el-table-column align="left" label="商品总数量" prop="num" width="120" />
        <el-table-column align="left" label="订单商品总金额" prop="total" width="120" />
        <el-table-column align="left" label="邮费" prop="postage" width="120" />
        <el-table-column align="left" label="实付金额" prop="finish" width="120" />
        <el-table-column align="left" label="支付方式" prop="payment" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.payment,paymentOptions) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="支付详情信息" prop="paymentInfo" width="120" />
        <el-table-column align="left" label="支付openId" prop="paymentOpenid" width="120" />
        <el-table-column align="left" label="支付流水订单号" prop="transationId" width="120" />
        <el-table-column align="left" label="留言" prop="remarks" width="120" />
        <el-table-column align="left" label="订单状态" prop="status" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.status,orderStatusOptions) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="取消状态" prop="statusCancel" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.statusCancel,GoodsStatusCancelOptions) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="退款状态" prop="statusRefund" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.statusRefund,OrderStatusRefundOptions) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="支付时间" width="180">
          <template #default="scope">{{ formatDate(scope.row.payTime) }}</template>
        </el-table-column>
        <el-table-column align="left" label="发货时间" width="180">
          <template #default="scope">{{ formatDate(scope.row.shipmentTime) }}</template>
        </el-table-column>
        <el-table-column align="left" label="收货时间" width="180">
          <template #default="scope">{{ formatDate(scope.row.receiveTime) }}</template>
        </el-table-column>
        <el-table-column align="left" label="取消时间" width="180">
          <template #default="scope">{{ formatDate(scope.row.cancelTime) }}</template>
        </el-table-column>
        <el-table-column align="left" label="创建日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="操作">
          <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateOrderFunc(scope.row)">变更</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="type === 'create' ? '新增操作' : '修改操作'">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户id:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品id:" prop="goodsId">
          <el-input v-model.number="formData.goodsId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="订单编号:" prop="orderSn">
          <el-input v-model="formData.orderSn" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="所属区域:" prop="goodsArea">
          <el-select v-model="formData.goodsArea" placeholder="请选择" style="width:100%" :clearable="true">
            <el-option v-for="(item,key) in goods_areaOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="收货人姓名:" prop="shipmentName">
          <el-input v-model="formData.shipmentName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="收货人手机号:" prop="shipmentMobile">
          <el-input v-model="formData.shipmentMobile" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="收货人地区编码:" prop="shipmentArea">
          <el-input v-model="formData.shipmentArea" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="收获人地址:" prop="shipingAddress">
          <el-input v-model="formData.shipingAddress" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品总数量:" prop="num">
          <el-input v-model.number="formData.num" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="订单商品总金额:" prop="total">
          <el-input-number v-model="formData.total" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="邮费:" prop="postage">
          <el-input-number v-model="formData.postage" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="实付金额:" prop="finish">
          <el-input-number v-model="formData.finish" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="支付方式:" prop="payment">
          <el-select v-model="formData.payment" placeholder="请选择" style="width:100%" :clearable="true">
            <el-option v-for="(item,key) in paymentOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="支付详情信息:" prop="paymentInfo">
          <el-input v-model="formData.paymentInfo" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="支付openId:" prop="paymentOpenid">
          <el-input v-model="formData.paymentOpenid" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="支付流水订单号:" prop="transationId">
          <el-input v-model="formData.transationId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="留言:" prop="remarks">
          <el-input v-model="formData.remarks" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="订单状态:" prop="status">
          <el-select v-model="formData.status" placeholder="请选择" style="width:100%" :clearable="true">
            <el-option v-for="(item,key) in orderStatusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="取消状态:" prop="statusCancel">
          <el-select v-model="formData.statusCancel" placeholder="请选择" style="width:100%" :clearable="true">
            <el-option v-for="(item,key) in GoodsStatusCancelOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="退款状态:" prop="statusRefund">
          <el-select v-model="formData.statusRefund" placeholder="请选择" style="width:100%" :clearable="true">
            <el-option v-for="(item,key) in OrderStatusRefundOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="支付时间:" prop="payTime">
          <el-date-picker v-model="formData.payTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
        </el-form-item>
        <el-form-item label="发货时间:" prop="shipmentTime">
          <el-date-picker v-model="formData.shipmentTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
        </el-form-item>
        <el-form-item label="收货时间:" prop="receiveTime">
          <el-date-picker v-model="formData.receiveTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
        </el-form-item>
        <el-form-item label="取消时间:" prop="cancelTime">
          <el-date-picker v-model="formData.cancelTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
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
  name: 'Order'
}
</script>

<script setup>
import {
  createOrder,
  deleteOrder,
  deleteOrderByIds,
  updateOrder,
  findOrder,
  getOrderList
} from '@/api/order'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const paymentOptions = ref([])
const orderStatusOptions = ref([])
const GoodsStatusCancelOptions = ref([])
const OrderStatusRefundOptions = ref([])
const goods_areaOptions = ref([])
const formData = ref({
  userId: 0,
  goodsId: 0,
  orderSn: '',
  goodsArea: undefined,
  shipmentName: '',
  shipmentMobile: '',
  shipmentArea: '',
  shipingAddress: '',
  num: 0,
  total: 0,
  postage: 0,
  finish: 0,
  payment: undefined,
  paymentInfo: '',
  paymentOpenid: '',
  transationId: '',
  remarks: '',
  status: undefined,
  statusCancel: undefined,
  statusRefund: undefined,
  payTime: new Date(),
  shipmentTime: new Date(),
  receiveTime: new Date(),
  cancelTime: new Date(),
})

// 验证规则
const rule = reactive({
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
  const table = await getOrderList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
  paymentOptions.value = await getDictFunc('payment')
  orderStatusOptions.value = await getDictFunc('orderStatus')
  GoodsStatusCancelOptions.value = await getDictFunc('GoodsStatusCancel')
  OrderStatusRefundOptions.value = await getDictFunc('OrderStatusRefund')
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
    deleteOrderFunc(row)
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
  const res = await deleteOrderByIds({ ids })
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
const updateOrderFunc = async(row) => {
  const res = await findOrder({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reorder
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteOrderFunc = async(row) => {
  const res = await deleteOrder({ ID: row.ID })
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
    userId: 0,
    goodsId: 0,
    orderSn: '',
    goodsArea: undefined,
    shipmentName: '',
    shipmentMobile: '',
    shipmentArea: '',
    shipingAddress: '',
    num: 0,
    total: 0,
    postage: 0,
    finish: 0,
    payment: undefined,
    paymentInfo: '',
    paymentOpenid: '',
    transationId: '',
    remarks: '',
    status: undefined,
    statusCancel: undefined,
    statusRefund: undefined,
    payTime: new Date(),
    shipmentTime: new Date(),
    receiveTime: new Date(),
    cancelTime: new Date(),
  }
}
// 弹窗确定
const enterDialog = async() => {
     elFormRef.value?.validate(async(valid) => {
       if (!valid) return
       let res
       switch (type.value) {
         case 'create':
           res = await createOrder(formData.value)
           break
         case 'update':
           res = await updateOrder(formData.value)
           break
         default:
           res = await createOrder(formData.value)
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

<style>
</style>
