<template>
  <div>
    <div class="gva-form-box">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="100px">
        <el-tabs v-model="tabsIndex" class="demo-tabs" @tab-change="tabsChange">
          <el-tab-pane label="基本信息" name="1">
            <el-form-item label="商品名称:" prop="name">
              <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
            </el-form-item>
            <el-form-item label="商品图片:" prop="img">
              <ImageList :file-list="goodsImages" @on-success="goodsImagesSuccessHandle" @update:fileList="updateFileListHandle" />
            </el-form-item>
            <el-form-item label="所属区域:" prop="goodsArea">
              <el-select v-model="formData.goodsArea" placeholder="请选择" :clearable="false">
                <el-option v-for="(item,key) in goodsAreaOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <!-- 积分商品只能是单规格， 普通商品才有多规格 -->
            <el-form-item v-if="formData.goodsArea === 0" label="规格类型:" prop="specType">
              <el-radio-group v-model="formData.specType" class="ml-4">
                <el-radio v-for="(item,key) in specTypeOptions" :key="key" :label="item.value" size="large">{{ item.label }}</el-radio>
              </el-radio-group>
            </el-form-item>
            <el-form-item label="分类:" prop="categoryId">
              <el-select v-model="formData.categoryId" filterable placeholder="Select">
                <el-option
                  v-for="item in categoryOptions"
                  :key="item.ID"
                  :label="item.title"
                  :value="item.ID"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="品牌:" prop="brandId">
              <el-select v-model="formData.brandId" filterable placeholder="Select">
                <el-option
                  v-for="item in brandOptions"
                  :key="item.ID"
                  :label="item.name"
                  :value="item.ID"
                />
              </el-select>
            </el-form-item>
            <el-form-item :label="formData.goodsArea === 1 ? '购买所需积分' : '商品价格:' " prop="price">
              <el-input-number v-model="formData.price" :precision="2" :clearable="true" />
            </el-form-item>
            <el-form-item label="最低购买数量:" prop="minCount">
              <el-input v-model.number="formData.minCount" :clearable="true" placeholder="请输入" />
            </el-form-item>
            <el-form-item label="商品单位:" prop="unit">
              <el-input v-model="formData.unit" :clearable="true" placeholder="请输入" />
            </el-form-item>
            <el-form-item label="商品重量(g):" prop="weight">
              <el-input v-model.number="formData.weight" :clearable="true" placeholder="请输入" />
            </el-form-item>
            <el-form-item label="库存:" prop="store">
              <el-input v-model.number="formData.store" :clearable="true" placeholder="请输入" />
            </el-form-item>
            <el-form-item label="排序:" prop="sort">
              <el-input v-model.number="formData.sort" :clearable="true" placeholder="请输入" />
            </el-form-item>
            <el-form-item label="上架状态:" prop="status">
              <el-radio-group v-model="formData.status" class="ml-4">
                <el-radio v-for="(item,key) in goodsStatus" :key="key" :label="item.value" size="large">{{ item.label }}</el-radio>
              </el-radio-group>
            </el-form-item>
            <el-form-item label="是否首页:" prop="isFirst">
              <el-radio-group v-model="formData.isFirst" class="ml-4">
                <el-radio v-for="(item,key) in whetherOptions" :key="key" :label="item.value" size="large">{{ item.label }}</el-radio>
              </el-radio-group>
            </el-form-item>
            <el-form-item label="是否热销:" prop="isHot">
              <el-radio-group v-model="formData.isHot" class="ml-4">
                <el-radio v-for="(item,key) in whetherOptions" :key="key" :label="item.value" size="large">{{ item.label }}</el-radio>
              </el-radio-group>
            </el-form-item>
            <el-form-item label="是否上新:" prop="isNew">
              <el-radio-group v-model="formData.isNew" class="ml-4">
                <el-radio v-for="(item,key) in whetherOptions" :key="key" :label="item.value" size="large">{{ item.label }}</el-radio>
              </el-radio-group>
            </el-form-item>
            <el-form-item label="商品详情:" prop="detail">
              <Editor ref="detailEditorRef" :content="goodsDesc.details" @update:content="handleUpdateDetail" />
            </el-form-item>
          </el-tab-pane>
          <!--    商品多规格      -->
          <!-- 积分商品只能是单规格， 普通商品才有多规格 -->
          <el-tab-pane v-if="formData.specType === 1 && formData.goodsArea === 0" label="商品规格" name="2">
            <el-form-item label="商品规格" prop="isNew">
              <el-card class="box-card" shadow="never" style="width: 100%;">
                <div v-for="(s, specId) in spec" :key="specId">
                  <div class="input-row">
                    <div style="width: 58px; padding-top: 8px">
                      <span>规格名</span>
                    </div>
                    <div class="input-wrap">
                      <div class="input-item">
                        <el-input v-model="spec[specId].name" placeholder="请输入规格名称" />
                        <el-icon class="input-close" @click="removeSpec(specId)">
                          <CircleClose />
                        </el-icon>
                      </div>
                    </div>
                  </div>
                  <!--                  <div class="input-row">
                    <div style="width: 58px; padding-top: 8px">
                      <span>上传图片</span>
                    </div>
                    <div class="input-wrap">
                      <div class="input-item" style="margin-left: 16px">
                        <el-radio-group v-model="spec[specId].isUploadImage" class="ml-4">
                          <el-radio :label="0" size="large">否</el-radio>
                          <el-radio :label="1" size="large">是</el-radio>
                        </el-radio-group>
                      </div>
                    </div>
                  </div>-->
                  <!-- 规格值 -->
                  <div class="input-row">
                    <div style="width: 58px; padding-top: 8px">
                      <span>规格值</span>
                    </div>
                    <div class="input-wrap">
                      <div v-for="(item, itemId) in specItem[specId]" :key="itemId" class="input-item-key">
                        <el-input v-model="specItem[specId][itemId].name" placeholder="规格值" @blur="blurSpecItemInput" />
                        <el-icon class="input-close" @click="removeSpecItem(specId, itemId)">
                          <CircleClose />
                        </el-icon>
                      </div>
                      <el-button type="primary" @click="addSpecItemInput(specId)">添加规格值</el-button>
                    </div>
                  </div>
                  <el-divider />
                </div>
                <div>
                  <el-button type="primary" @click="addSpecInput">添加商品规格</el-button>
                  <el-button type="danger" @click="clearSpecInput">清空全部规格</el-button>
                </div>
              </el-card>
            </el-form-item>
            <!-- 规格名细 -->
            <el-form-item v-if="Object.keys(spec).length > 0" label="规格名细" prop="isNew">
              <el-table :data="tableData" style="width: 100%" :span-method="spanTableMethod">
                <el-table-column
                  v-for="(item) in spec"
                  :key="item.specId"
                  :label="item.name"
                  width="120"
                  :prop="item.specId"
                />
                <el-table-column label="库存" width="180">
                  <template #default="scope">
                    <el-input v-model.number="specValue[scope.row.valueId].store" :clearable="true" placeholder="0" />
                  </template>
                </el-table-column>
                <el-table-column label="价格" width="180">
                  <template #default="scope">
                    <el-input v-model.number="specValue[scope.row.valueId].price" :clearable="true" placeholder="0" />
                  </template>
                </el-table-column>
                <el-table-column label="排序">
                  <template #default="scope">
                    <el-input v-model.number="specValue[scope.row.valueId].sort" :clearable="true" placeholder="0" />
                  </template>
                </el-table-column>
              </el-table>
            </el-form-item>
          </el-tab-pane>
        </el-tabs>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="danger" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Goods',
}
</script>

<script setup>
import {
  createGoods,
  updateGoods,
  findGoods,
} from '@/api/goods'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
import { CircleClose } from '@element-plus/icons-vue'
import { getCategoryListAll } from '@/api/category'
import { getBrandListAll } from '@/api/brand'
import Editor from '@/components/quillEditor/editor.vue'
import ImageList from '@/components/upload/imageList.vue'

const route = useRoute()
const router = useRouter()
const tabsIndex = ref('1')
const id = ref(0)

const type = ref('')
const specTypeOptions = ref([])
const whetherOptions = ref([])
const goodsAreaOptions = ref([])
const goodsStatus = ref([])

const categoryOptions = ref([])
const brandOptions = ref([])

const formData = ref({
  name: '测试',
  categoryId: 1,
  brandId: 1,
  goodsArea: 0,
  specType: 0,
  unit: '件',
  price: 100,
  minCount: 1,
  weight: 200,
  store: 20,
  sort: 50,
  status: 1,
  isFirst: 0,
  isHot: 0,
  isNew: 0,
})
// 商品详情
const goodsDesc = ref({
  notice: '',
  details: ''
})

// 商品图片
const goodsImages = ref([])

// 验证规则
const rule = reactive({
  name: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  categoryId: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  brandId: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  goodsArea: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  specType: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  unit: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  price: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  minCount: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  store: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
})
const specVualId = ref(0) // 规格输入框数量
const specItemId = ref(1) // 全局规格项 id 累增
// const spec = ref([
//   {
//     specId: 1,
//     name: '颜色',
//     isUploadImage: 0,
//   },
//   {
//     specId: 2,
//     name: '大小',
//     isUploadImage: 0,
//   },
// ])
const spec = ref({})
// const specItem = {
//   spec_1: {
//     item_1: {
//       name: "红色"
//     }
//   },
//   spec_2: {
//     item_1: {
//       name: "大号"
//     }
//   }
// }
const specItem = ref({})
const specValue = ref({})
const tableData = ref([])

// 初始化方法
const init = async() => {
  const queryId = route.query.id
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (queryId && queryId !== '新增商品') {
    id.value = queryId
    const res = await findGoods({ ID: id.value })
    if (res.code === 0) {
      formData.value = res.data.regoods
      type.value = 'update'
      initEditGoods()
    }
  } else {
    type.value = 'create'
  }
  specTypeOptions.value = await getDictFunc('spec_type')
  whetherOptions.value = await getDictFunc('whether')
  goodsAreaOptions.value = await getDictFunc('goods_area')
  goodsStatus.value = await getDictFunc('goods_status')
  const categroyRes = await getCategoryListAll()
  // 获取分类
  if (categroyRes.code === 0) {
    categoryOptions.value = categroyRes.data
  } else {
    ElMessage.error('获取分类列表失败')
    return
  }
  // 获取所有品牌列表
  const brandRes = await getBrandListAll()
  if (brandRes.code === 0) {
    brandOptions.value = brandRes.data
  } else {
    ElMessage.error('获取品牌列表失败')
    return
  }
}

init()

// 编辑商品时初始化
const initEditGoods = () => {
  // 初始化图片
  goodsImages.value = formData.value.images

  // 初始化详情
  goodsDesc.value.details = formData.value.desc.details

  // 多规格时初始化
  if (formData.value.specType === 1) {
    initEditGoodsSpec()
  }
}

// 编辑商品时规格的初始化
const initEditGoodsSpec = () => {
  spec.value = {}
  specItem.value = {}
  specValue.value = {}
  // 初始化规格
  const formSpec = formData.value.spec
  specVualId.value = formSpec[formSpec.length - 1].ID + 1000 ?? 0
  specItemId.value = formSpec[formSpec.length - 1].specItem[0].ID + 1000
  formSpec.forEach(s => {
    spec.value[s.ID] = {
      specId: s.ID + '',
      name: s.title,
      isUploadImage: s.isUploadImage
    }
    // 初始化规格项
    s.specItem.forEach(i => {
      if (!specItem.value[i.specId]) {
        specItem.value[i.specId] = {}
      }
      specItem.value[i.specId][i.ID] = {
        specId: i.specId + '',
        itemId: i.ID + '',
        name: i.item
      }
    })
    // 生成规格明细
    cartesianProductTableData()
    // 对规格明细的值进行填充
    tableData.value.forEach((t, key) => {
      formData.value.specValue.forEach(v => {
        const itemIds = v.itemIds.replaceAll('_', ',')
        if (t.valueId === itemIds) {
          specValue.value[itemIds].price = v.price
          specValue.value[itemIds].store = v.store
          specValue.value[itemIds].sort = v.sort
        }
      })
    })
    console.log('specValue', specValue.value)
  })
}

// 递归绩笛卡尔积
const cartesianProductOf = (...args) => {
  // 将不定长参数转换为数组并调用 reduce 方法, 返回笛卡尔积结果
  const data = Array.prototype.reduce.call(args, (a, b) => {
    const ret = []
    // 计算笛卡尔积
    a.forEach(aa => {
      b.forEach(bb => {
        // 合并数组
        ret.push(aa.concat([bb]))
      })
    })
    // 返回笛卡尔积结果
    return ret
  }, [[]]) // 初始值是一个空数组
  // 将笛卡尔积中的每个数组转换成包含所有键值对的对象
  const result = data.map(arr => {
    let specId = ''
    let valueId = ''
    arr.forEach(item => {
      if (item.specId) {
        specId += ',' + item.specId
      }
      if (item.itemId) {
        valueId += ',' + item.itemId
      }
    })
    // 删除前面逗号
    specId = specId.replace(/^,+/g, '')
    valueId = valueId.replace(/^,+/g, '')
    // 初始化 specValue
    if (!specValue.value[valueId]) {
      specValue.value[valueId] = {
        sort: 50
      }
    }
    // 返回
    const obj = Object.assign({}, ...arr)
    obj.specId = specId
    obj.valueId = valueId
    return obj
  })
  const retainArr = result.map(item => item.valueId)
  specValue.value = retainKey(specValue.value, retainArr)
  return result
}

// 生成规格明细表格内容
const cartesianProductTableData = () => {
  const keys = Object.keys(specItem.value) // 获取到 specItem = ['1', '2']
  const valueList = keys.map(key => {
    const items = specItem.value[key]
    return Object.values(items).map((item) => {
      return {
        [key]: item.name,
        specId: item.specId,
        itemId: item.itemId
      }
    })
  })
  const result = cartesianProductOf(...valueList)
  tableData.value = result
  console.log('result', result)
}

// 在对象中保留多个指定的键并删除其他所有键
const retainKey = (obj, keys) => {
  Object.keys(obj).forEach((k) => {
    if (!keys.includes(k)) {
      delete obj[k]
    }
  })
  return obj
}

// 当规格值失去焦点时，笛卡尔积计算列表数据
const blurSpecItemInput = () => {
  cartesianProductTableData()
}

// 添加规格输入框
const addSpecInput = () => {
  const specId = specVualId.value + 1 + ''
  // 初始化规格双向绑定数据
  spec.value[specId] = {
    specId: specId,
    isUploadImage: 0,
  }
  // 初始化规格值
  specItem.value[specId] = {}
  // 清空规格明细
  // specValue.value = {}
  // 输入框加 1
  specVualId.value++
}

// 添加规格值输入框
const addSpecItemInput = (specId) => {
  // 初始化规格值保存数据
  const itemId = (++specItemId.value) + ''
  specItem.value[specId][itemId] = {
    specId: specId,
    itemId: itemId,
  }
}

// 删除规格
const removeSpec = (specId) => {
  delete spec.value[specId]
  delete specItem.value[specId]
  cartesianProductTableData()
  ElMessage.success('删除规格项成功')
}

// 删除规格值
const removeSpecItem = (specId, itemId) => {
  delete specItem.value[specId][itemId]
  cartesianProductTableData()
  ElMessage.success('删除规格值成功')
}

// 清空规格值
const clearSpecInput = () => {
  specVualId.value = 0
  spec.value = {}
  specItem.value = {}
  cartesianProductTableData()
  ElMessage.success('清空成功')
}

// TODO 合并单元格子
const spanTableMethod = ({
  row,
  column,
  rowIndex,
  columnIndex,
}) => {
}

const elFormRef = ref()

const tabsChange = (index, e) => {
  console.log('tabIndex,', index)
}

const goodsImagesSuccessHandle = (url, name) => {
  goodsImages.value.push({
    url, name,
  })
}

// 更新图片文件列表
const updateFileListHandle = (files) => {
  goodsImages.value = []
  files.forEach(item => {
    goodsImages.value.push({
      url: item.url,
      name: item.name,
    })
  })
}

const detailEditorRef = ref(null)
// 保存按钮
const save = async() => {
  console.log('sava orgin spec', spec.value)
  console.log('sava orgin specItem', specItem.value)
  console.log('sava orgin specValue', specValue.value)
  // 如果是积分商品 必定是单规格
  if (formData.value.goodsArea === 1) {
    formData.value.specType = 0
  }
  if (goodsImages.value.length === 0) {
    ElMessage.error('请上传商品图片')
    return
  }
  elFormRef.value?.validate(async(valid) => {
    if (!valid) {
      ElMessage.error('您有必填项未填写完成，请检查')
      return
    }
    const specData = []
    for (const sId in spec.value) {
      specData.push(spec.value[sId])
    }
    const specItemData = []
    for (const sId in specItem.value) {
      for (const itemId in specItem.value[sId]) {
        specItemData.push(specItem.value[sId][itemId])
      }
    }
    const goodsInfo = formData.value
    goodsInfo.spec = null
    goodsInfo.specValue = null
    goodsInfo.images = null
    goodsInfo.desc = null
    const data = {
      goodsInfo: formData.value,
      spec: specData,
      specItem: specItemData,
      specValue: specValue.value,
      desc: goodsDesc.value,
      images: goodsImages.value
    }
    let res
    switch (type.value) {
      case 'create':
        res = await createGoods(data)
        break
      case 'update':
        res = await updateGoods(data)
        break
      default:
        res = await createGoods(data)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '操作成功',
      })
      if (type.value === 'update') {
        await init()
      } else {
        router.go(-1)
      }
    }
  })
}

// 接收子组件更新字段信息
const handleUpdateDetail = (val) => {
  goodsDesc.value.details = val
}

// 返回按钮
const back = () => {
  router.go(-1)
}

</script>

<style lang="scss">
.el-form-item__label {
  width: 120px !important;
}
.input-row {
  display: flex;
  margin-bottom: 20px;

  .input-wrap {
    flex: 1;
  }

  .input-item {
    width: 100%;
    display: inline-block;
    position: relative;
  }

  .input-item-key {
    display: inline-block;
    position: relative;

  }

  .el-input__wrapper {
    margin: 8px 8px 8px 8px;
  }

  .input-close {
    position: absolute;
    color: #e5e2e2;
    font-size: 26px;
    right: 0;
    top: -2px;
    background: #ffffff;
    cursor: pointer;
  }

  .input-close:hover {
    color: #b2b2b2;
  }
}

</style>
