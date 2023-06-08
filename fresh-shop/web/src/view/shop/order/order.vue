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
      <!--      <div class="gva-btn-list">
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
      </div>-->
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column align="left" label="基本信息" prop="orderInfo" width="300">
          <template #default="scope">
            <div class="table-multi-line">
              <div style="display:inline-block;">
                订单号：<span style="color: #4d70ff">{{ scope.row.orderSn }}</span>
              </div><br>
              <span>编号：{{ scope.row.ID }}</span><el-divider direction="vertical" />
              <span>订单类型：{{ filterDict(scope.row.goodsArea,goods_areaOptions) }}</span> <br>

              <div style="display:inline-block">
                订单状态：
                <!-- 判断订单是否取消 -->
                <span v-if="scope.row.statusCancel === 1" style="color: #000;font-weight: bold">已取消</span>
                <span v-else-if="scope.row.statusCancel === 2" style="color: #000;font-weight: bold">后台取消</span>
                <span v-else-if="scope.row.statusCancel === 3" style="color: #000;font-weight: bold">超时取消</span>
                <span v-else>
                  <!-- 判断是否是售后单 -->
                  <span v-if="scope.row.return.ID !== 0">
                    <span v-if="scope.row.return.refundStatus === 1">退款等待到账</span>
                    <span v-else-if="scope.row.return === -1" style="color:#fa3534;font-weight: bold">拒绝售后</span>
                    <span v-else-if="scope.row.return === 0" style="color:#ff9900;font-weight: bold">等待审核</span>
                    <span v-else-if="scope.row.return === 1 && scope.row.return.refundStatus === 2">售后完成</span>
                    <span v-else-if="scope.row.return === 1" style="font-weight: bold">审核通过</span>
                  </span>
                  <!-- 普通订单 -->
                  <span v-else>
                    <span v-if="scope.row.status === 0" style="color: #1677ff;font-weight: bold">等待支付</span>
                    <span v-else-if="scope.row.status === 1" style="color:#19be6b;font-weight: bold">备货中</span>
                    <span v-else-if="scope.row.status === 2" style="color: #13c2c2;font-weight: bold">已发货</span>
                    <span v-else-if="scope.row.status === 3" style="color: #3f6600;font-weight: bold">已完成</span>
                  </span>
                </span>
              </div>
              <el-divider direction="vertical" />
              <span>支付方式：
                <span v-if="scope.row.status === 0">无</span>
                <span v-else style="color: #52c41a">
                  {{ filterDict(scope.row.payment,paymentOptions) }}
                </span>
              </span> <br>

              <span>订单金额：{{ scope.row.total }} {{ scope.row.goodsArea === 1 ? '积分' : '元' }}</span><el-divider direction="vertical" />
              <span>实付金额：{{ scope.row.finish }} {{ scope.row.goodsArea === 1 ? '积分' : '元' }}</span><br>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="订单商品信息" prop="orderInfo" width="290">
          <template #default="scope">
            <div v-if="scope.row.details && scope.row.details.length === 1" class="table-multi-line single-goods">
              <div class="single-goods-img-box">
                <img
                  class="single-goods-img"
                  :src="(scope.row.details[0].goodsImage && scope.row.details[0].goodsImage.slice(0, 4) !== 'http') ? path + scope.row.details[0].goodsImage:scope.row.details[0].goodsImage"
                  @click="handlePictureCardPreview(scope.row.details[0].goodsImage)"
                >
              </div>
              <div>
                <div class="single-goods-title">
                  {{ scope.row.details[0].goodsName }}
                </div>
                <div>
                  <span>规格：{{ scope.row.details[0].specKeyName }}</span> <el-divider direction="vertical" />
                  <span class="king-ml-10">数量：{{ scope.row.details[0].num }}</span>
                </div>
                <div>
                  <span>商品单价：{{ scope.row.details[0].price }} {{ scope.row.goodsArea === 1 ? '积分' : '元' }}</span><el-divider direction="vertical" />
                  <span>商品总价：{{ scope.row.total }} {{ scope.row.goodsArea === 1 ? '积分' : '元' }}</span>
                </div>
              </div>
            </div>
            <div v-else-if="scope.row.details && scope.row.details.length > 1" class="multiple-goods">
              <div class="multiple-goods-img-box">
                <div v-for="(item,index) in scope.row.details" :key="index">
                  <img
                    v-if="index < 2"
                    class="multiple-goods-img"
                    :src="(item.goodsImage && item.goodsImage.slice(0, 4) !== 'http') ? path + item.goodsImage:item.goodsImage"
                    alt=""
                    @click="handlePictureCardPreview(item.goodsImage)"
                  >
                  <div v-if="index === 2" class="multiple-goods-img" @click="showOrderDetails(scope.row)">
                    ...
                  </div>
                </div>
                <div class="multiple-goods-img" @click="showOrderDetails(scope.row)">
                  查看详情
                </div>
              </div>
            </div>
            <div>订单留言：{{ scope.row.remarks ? scope.row.remarks : '无' }}</div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="收货人信息" prop="orderInfo" width="280">
          <template #default="scope">
            <div class="table-multi-line">
              <span>提货方式：{{ scope.row.shipmentType === 1 ? '自提': '配送' }}</span>
              <div v-if="scope.row.shipmentType === 1">
                <span>取货号：<span  style="color: #00afff">A{{ scope.row.pickUpNumber }}</span></span>
              </div>
              <div v-else>
                <span>姓名：{{ scope.row.shipmentName }}</span><el-divider direction="vertical" />
                <span>手机号：{{ scope.row.shipmentMobile }}</span><br>
                <span>地址：{{ scope.row.shipmentAddress }}</span><br>
              </div>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="时间" prop="paymentInfo" width="230">
          <template #default="scope">
            <div class="table-multi-line">
              <div>创建时间：{{ formatDate(scope.row.CreatedAt) }}</div>
              <div v-if="scope.row.status === 1">支付时间：{{ formatDate(scope.row.payTime) }}</div>
              <div v-if="scope.row.statusCancel === 1">取消时间：{{ formatDate(scope.row.cancelTime) }}</div>
              <div v-if="scope.row.status === 2">发货时间：{{ formatDate(scope.row.shipmentTime) }}</div>
              <div v-if="scope.row.status === 3">收货时间：{{ formatDate(scope.row.receiveTime) }}</div>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作">
          <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateOrderFunc(scope.row)">变更
            </el-button>
            <el-button
              v-if="scope.row.status === 1 && scope.row.statusCancel === 0 && scope.row.statusRefund === 0"
              type="primary"
              link
              icon="Van"
              class="table-button"
              @click="showOrderShipment(scope.row, 'create')"
            >发货
            </el-button>
            <el-button
              v-if="scope.row.status === 2 && scope.row.statusCancel === 0 && scope.row.statusRefund === 0"
              type="primary"
              link
              icon="Van"
              class="table-button"
              @click="showOrderShipment(scope.row, 'update')"
            >确认收货
            </el-button>
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
        <el-form-item label="订单编号:" prop="orderSn">
          <el-input v-model="formData.orderSn" disabled="disabled" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="所属区域:" prop="goodsArea">
          <el-select v-model="formData.goodsArea" placeholder="请选择" disabled="disabled" style="width:100%" :clearable="true">
            <el-option v-for="(item,key) in goods_areaOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="收货人姓名:" prop="shipmentName">
          <el-input v-model="formData.shipmentName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="收货人手机号:" prop="shipmentMobile">
          <el-input v-model="formData.shipmentMobile" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="收货人地址:" prop="shipmentAddress">
          <el-input v-model="formData.shipmentAddress" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品总数量:" prop="num">
          <el-input v-model.number="formData.num" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="订单总金额:" prop="total">
          <el-input-number v-model="formData.total" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <!--        <el-form-item label="邮费:" prop="postage">-->
        <!--          <el-input-number v-model="formData.postage" style="width:100%" :precision="2" :clearable="true" />-->
        <!--        </el-form-item>-->
        <el-form-item label="实付金额:" prop="finish">
          <el-input-number v-model="formData.finish" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="支付方式:" prop="payment">
          <el-select v-model="formData.payment" disabled="disabled" placeholder="请选择" style="width:100%" :clearable="true">
            <el-option v-for="(item,key) in paymentOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="支付详情信息:" prop="paymentInfo">
          <el-input v-model="formData.paymentInfo" disabled="disabled" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="支付流水订单号:" prop="transationId">
          <el-input v-model="formData.transationId" disabled="disabled" :clearable="true" placeholder="请输入" />
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
    <el-dialog v-model="dialogImgVisible" title="查看商品图片">
      <img style="width: 100%" :src="dialogImageUrl" alt="Preview Image">
    </el-dialog>
    <!--  查看订单详情   -->
    <el-dialog v-model="dialogDetailsVisible" width="80%" title="查看订单详情">
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableDetailData"
        row-key="ID"
      >
        <el-table-column align="center" label="商品编号" prop="ID" width="80" />
        <el-table-column align="left" label="商品图片" prop="goodsImage" width="160">
          <template #default="scope">
            <img
              style="width: 100%;object-fit:cover;"
              :src="(scope.row.goodsImage && scope.row.goodsImage.slice(0, 4) !== 'http') ? path + scope.row.goodsImage:scope.row.goodsImage"
              @click="handlePictureCardPreview(scope.row.goodsImage)"
            >
          </template>
        </el-table-column>
        <el-table-column align="left" label="商品名称" prop="goodsName" width="300" />
        <el-table-column align="left" label="详情信息" prop="goodsInfo">
          <template #default="scope">
            <div class="table-multi-line">
              <span>规格：{{ scope.row.specKeyName }}</span> <el-divider direction="vertical" />
              购买数量：<span style="color: #4d70ff">{{ scope.row.num }} {{ scope.row.unit }}</span> <br>
              购买价格：<span style="color: #19be6b">{{ scope.row.total }} 元</span>
            </div>
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>
    <!--  查看订单详情   -->
    <el-dialog v-model="dialogShipmentVisible" width="80%" title="订单发货">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="订单编号:" prop="orderSn">
          <el-input v-model="shipInfo.orderSn" disabled :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="订单数量:" prop="orderSn">
          <el-input v-model="shipInfo.num" disabled :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="订单金额:" prop="orderSn">
          <el-input v-model="shipInfo.total" disabled :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="选择送货人:" prop="deliveryId">
          <el-select v-model="shipmentformData.deliveryId" placeholder="请选择" style="width:100%" :clearable="true">
            <el-option label="自提" :value="0" />
            <el-option v-for="(item, key) in deliveryList" :key="key" :label="item.name" :value="item.ID" />
          </el-select>
        </el-form-item>
        <el-form-item v-if="shipmentformData.deliverName" label="送货人名称:" prop="deliveryName">
          <el-input v-model="shipmentformData.deliverName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item v-if="shipmentformData.deliverMobile" label="送货人手机号:" prop="deliveryMobile">
          <el-input v-model="shipmentformData.deliverMobile" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item v-if="shipmentformData.deliverName" label="预计送达时间:" prop="scheduledTime">
          <el-date-picker
            v-model="shipmentformData.scheduledTime"
            type="datetime"
            placeholder="请选择预计送达时间"
            format="YYYY-MM-DD HH:mm:ss"
            value-format="YYYY-MM-DD HH:mm:ss"
          />
        </el-form-item>
        <el-form-item v-if="shipmentType === 'update'" label="确认收货时间:" prop="receiptTime">
          <el-date-picker
            v-model="shipmentformData.receiptTime"
            type="datetime"
            placeholder="请选择确认收货时间"
            format="YYYY-MM-DD HH:mm:ss"
            value-format="YYYY-MM-DD HH:mm:ss"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeShipmentDialog">取 消</el-button>
          <el-button v-if="shipmentType === 'update'" type="primary" @click="enterShipmentDialog">确认送达</el-button>
          <el-button v-else type="primary" @click="enterShipmentDialog">确认发货</el-button>
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
  updateOrder,
  findOrder,
  getOrderList
} from '@/api/order'

import { getUserDeliveryAllList } from '@/api/userDelivery'
import { createOrderDelivery, findOrderDelivery, updateOrderDelivery } from '@/api/orderDelivery'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, watch } from 'vue'
import { useRouter } from 'vue-router'

const path = ref(import.meta.env.VITE_BASE_API + '/')

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
  shipmentAddress: '',
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
const route = useRouter()
watch(
  () => route.currentRoute.value.query.goodsArea,
  (n, o) => {
    console.log('监听路由 goodsArea', n)
    if (n) {
      getTableData()
    }
  }
)
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const tableDetailData = ref([])
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
  searchInfo.value.goodsArea = route.currentRoute.value.query.goodsArea
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

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')
const shipmentType = ref('')

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
    shipmentAddress: '',
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

// 关闭弹窗
const closeShipmentDialog = () => {
  dialogShipmentVisible.value = false
  shipmentformData.value = {
    ID: null,
    orderId: 0,
    deliveryId: 0,
    scheduledTime: null,
    deliverName: '',
    deliverMobile: '',
  }
}
// 弹窗确定
const enterShipmentDialog = async() => {
  if (shipmentformData.value.scheduledTime) {
    shipmentformData.value.scheduledTime = new Date(shipmentformData.value.scheduledTime)
  }
  if (shipmentformData.value.receiptTime) {
    shipmentformData.value.receiptTime = new Date(shipmentformData.value.receiptTime)
  }
  let res
  switch (shipmentType.value) {
    case 'create':
      res = await createOrderDelivery(shipmentformData.value)
      break
    case 'update':
      res = await updateOrderDelivery(shipmentformData.value)
      break
    default:
      res = await createOrderDelivery(shipmentformData.value)
      break
  }
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '操作成功',
    })
    closeShipmentDialog()
    getTableData()
  }
}

// 查看图片大图
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

const order = ref('')
const dialogDetailsVisible = ref(false)
const showOrderDetails = (row) => {
  order.value = row
  tableDetailData.value = row.details
  dialogDetailsVisible.value = true
}

const shipInfo = ref('')
const shipmentformData = ref({
  orderId: 0,
  deliveryId: 0,
  scheduledTime: null,
  deliverName: '',
  deliverMobile: '',
  receiptTime: null,
})

// 监听选择送货人信息
watch(() => shipmentformData.value.deliveryId, (newVal, oldVal) => {
  shipmentformData.value.deliverName = ''
  shipmentformData.value.deliverMobile = ''
  deliveryList.value.forEach(item => {
    if (item.ID === newVal) {
      shipmentformData.value.deliverName = item.name
      shipmentformData.value.deliverMobile = item.mobile
    }
  })
  if (shipmentType.value !== 'update') {
    // 设置预计送达时间 - 1小时后
    // 获取到一小时后的时间 格式为 yyyy-MM-dd HH:mm:ss
    const date = new Date()
    date.setHours(date.getHours() + 1)
    console.log(123123, formatDate(date))
    shipmentformData.value.scheduledTime = formatDate(date)
    console.log('watch, oldVal, newVal', oldVal, newVal)
  }
})

const dialogShipmentVisible = ref(false)
const deliveryList = ref({})
// type 0创建 1编辑
const showOrderShipment = async(row, type) => {
  shipInfo.value = row
  dialogShipmentVisible.value = true
  shipmentformData.value.orderId = row.ID
  shipmentType.value = type
  const res = await getUserDeliveryAllList({ status: 1 })
  if (res.code !== 0) {
    this.$message.error('获取送货人列表失败')
    return false
  }
  deliveryList.value = res.data.list
  if (type === 'update') {
    const dres = await findOrderDelivery({
      orderId: row.ID,
    })
    if (dres.code !== 0) {
      this.$message.error('获取发货信息失败')
      return false
    }
    const info = dres.data.reorderDelivery
    shipmentformData.value.ID = info.ID
    shipmentformData.value.orderId = info.orderId
    shipmentformData.value.deliveryId = info.deliveryId
    setTimeout(() => {
      shipmentformData.value.deliverMobile = info.deliverMobile
      shipmentformData.value.deliverName = info.deliverName
      if (info.scheduledTime && info.scheduledTime.charAt(0) !== '0') {
        shipmentformData.value.scheduledTime = formatDate(info.scheduledTime)
      }
      if (info.receiptTime && info.receiptTime.charAt(0) !== '0') {
        shipmentformData.value.receiptTime = formatDate(info.receiptTime)
      }
    }, 500)
  }
}
</script>

<style lang="scss">
.single-goods {
  display: flex;
  align-items: center;

  .single-goods-title {
    overflow: hidden;
    text-overflow: ellipsis;
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
  }
}
.single-goods-img-box {
  margin-right: 10px;
  .single-goods-img {
    width: 85px;
    height: 85px;
    border-radius: 10px;
  }
}

.multiple-goods {
  .multiple-goods-img-box {
    display: flex;
    align-items: center;

    .multiple-goods-img {
      width: 60px;
      height: 60px;
      border-radius: 10px;
      margin-right: 6px;
      text-align: center;
      font-size: 14px;
      font-weight: bold;
      cursor: pointer;
    }
  }
}

</style>
