<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="用户名">
          <el-input v-model="searchInfo.username" placeholder="用户名" />
        </el-form-item>
        <el-form-item label="昵称">
          <el-input v-model="searchInfo.nickName" placeholder="昵称" />
        </el-form-item>
        <el-form-item label="手机号">
          <el-input v-model="searchInfo.phone" placeholder="手机号" />
        </el-form-item>
        <el-form-item label="联系人名称">
          <el-input v-model="searchInfo.contactName" placeholder="联系人名称" />
        </el-form-item>
        <el-form-item label="客户名称">
          <el-input v-model="searchInfo.customerName" placeholder="客户名称" />
        </el-form-item>
        <el-form-item label="推荐码">
          <el-input v-model="searchInfo.invitationCode" placeholder="推荐码" />
        </el-form-item>
        <el-form-item label="审核状态" prop="status">
          <el-select
            v-model="searchInfo.auditStatus"
            clearable
            placeholder="请选择"
            @clear="()=>{searchInfo.auditStatus=undefined}"
          >
            <el-option v-for="(item,key) in auditStatusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="getTableData">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="addUser">新增用户</el-button>
      </div>
      <el-table
        :data="tableData"
        @sort-change="sortChange"
      >
        <el-table-column align="left" label="ID" min-width="80" prop="ID" sortable="custom" />
        <el-table-column align="left" label="头像" min-width="75">
          <template #default="scope">
            <CustomPic style="margin-top:8px" :pic-src="scope.row.headerImg" />
          </template>
        </el-table-column>
        <el-table-column align="left" label="用户信息" min-width="220" prop="userName" style="line-height: 10px">
          <template #default="scope">
            <div class="table-multi-line">
              <span class="">用户名：{{ scope.row.userName }}</span><br>
              <span>用户昵称：{{ scope.row.nickName }}</span><br>
              <span>手机号：{{ scope.row.phone ? scope.row.phone : '无' }}</span><br>
              <span>邀请码：{{ scope.row.invitationCode }}</span><br>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="会员信息" min-width="220" prop="userName" style="line-height: 10px">
          <template #default="scope">
            <div class="table-multi-line">
              <div v-if="scope.row.auditStatus === 3">
                <span>原始联系人名称：{{ scope.row.originContactName ? scope.row.originContactName : '无' }}</span><br>
                <span>原始客户名称：{{ scope.row.originCustomerName ? scope.row.originCustomerName : '无' }}</span><br>
                修改联系人名称：<span
                  style="color: #039BE5"
                >{{ scope.row.changeContactName ? scope.row.changeContactName : '无' }}</span><br>
                修改客户名称：<span
                  style="color: #039BE5"
                >{{ scope.row.changeCustomerName ? scope.row.changeCustomerName : '无' }}</span><br>
                <span v-if="scope.row.auditRemark">审核备注：{{ scope.row.auditRemark ? scope.row.auditRemark : '无' }}</span><br>
              </div>
              <!--       默认状态       -->
              <div v-else>
                <span>联系人名称：{{ scope.row.originContactName ? scope.row.originContactName : '无' }}</span><br>
                <span>客户名称：{{ scope.row.originCustomerName ? scope.row.originCustomerName : '无' }}</span><br>
                状态：<span
                  :class="getStatusColorClass(scope.row.auditStatus)"
                >{{ filterDict(scope.row.auditStatus, auditStatusOptions) }}</span><br>
                <span v-if="scope.row.auditRemark">审核备注：{{ scope.row.auditRemark ? scope.row.auditRemark : '无' }}</span><br>
              </div>
            </div>
          </template>
        </el-table-column>
        <!--        <el-table-column align="left" label="账户信息" min-width="130" prop="point" sortable="custom">-->
        <!--          <template #default="scope">-->
        <!--            <div class="table-multi-line">-->
        <!--              <div v-for="ac in scope.row.account" :key="ac.ID">-->
        <!--                <span>{{ ac.group.nameCn }}：{{ ac.amount }}</span>-->
        <!--              </div>-->
        <!--            </div>-->
        <!--          </template>-->
        <!--        </el-table-column>-->
        <el-table-column align="left" label="用户角色" min-width="200">
          <template #default="scope">
            <el-cascader
              v-model="scope.row.authorityIds"
              :options="authOptions"
              :show-all-levels="false"
              collapse-tags
              :props="{ multiple:true,checkStrictly: true,label:'authorityName',value:'authorityId',disabled:'disabled',emitPath:false}"
              :clearable="false"
              :disabled="userStore.role.authorityId !== 888"
              @visible-change="(flag)=>{changeAuthority(scope.row,flag,0)}"
              @remove-tag="(removeAuth)=>{changeAuthority(scope.row,false,removeAuth)}"
            />
          </template>
        </el-table-column>
        <el-table-column align="left" label="状态" min-width="80">
          <template #default="scope">
            <el-switch
              v-model="scope.row.enable"
              inline-prompt
              :active-value="1"
              :inactive-value="2"
              :disabled="scope.row.isSuperAdmin"
              @change="()=>{switchEnable(scope.row)}"
            />
          </template>
        </el-table-column>
        <el-table-column align="left" label="登录信息" min-width="250" style="line-height: 10px">
          <template #default="scope">
            <div v-if="scope.row.loginTime" class="table-multi-line">
              <span>登录IP：{{ scope.row.loginIp }}</span><br>
              <span>最后登录时间：{{ scope.row.loginTime ? formatDate(scope.row.loginTime) : '无' }}</span><br>
            </div>
            <span v-else>从未登录</span>
          </template>
        </el-table-column>
        <el-table-column label="操作" min-width="250" fixed="right">
          <template #default="scope">
            <div v-if="!scope.row.isSuperAdmin">
              <el-popover v-model="scope.row.visible" placement="top" width="160">
                <p>确定要删除此用户吗</p>
                <div style="text-align: right; margin-top: 8px;">
                  <el-button type="primary" link @click="scope.row.visible = false">取消</el-button>
                  <el-button type="primary" @click="deleteUserFunc(scope.row)">确定</el-button>
                </div>
                <template #reference>
                  <el-button type="primary" link icon="delete">删除</el-button>
                </template>
              </el-popover>
              <el-button type="primary" link icon="edit" @click="openEdit(scope.row)">编辑</el-button>
              <el-button type="primary" link icon="magic-stick" @click="resetPasswordFunc(scope.row)">重置密码
              </el-button>
            </div>
          </template>
        </el-table-column>

      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog
      v-model="addUserDialog"
      class="user-dialog"
      title="用户"
      :show-close="false"
      :close-on-press-escape="false"
      :close-on-click-modal="false"
    >
      <div style="height:60vh;overflow:auto;padding:0 12px;">
        <el-form ref="userForm" :rules="rules" :model="userInfo" label-width="80px">
          <el-form-item v-if="dialogFlag === 'add'" label="用户名" prop="userName">
            <el-input v-model="userInfo.userName" />
          </el-form-item>
          <el-form-item v-if="dialogFlag === 'add'" label="密码" prop="password">
            <el-input v-model="userInfo.password" />
          </el-form-item>
          <el-form-item label="昵称" prop="nickName">
            <el-input v-model="userInfo.nickName" />
          </el-form-item>
          <el-form-item label="手机号" prop="phone">
            <el-input v-model="userInfo.phone" />
          </el-form-item>
          <el-form-item label="联系人名称" prop="originContactName">
            <el-input v-model="userInfo.originContactName" />
          </el-form-item>
          <el-form-item label="客户名称" prop="originCustomerName">
            <el-input v-model="userInfo.originCustomerName" />
          </el-form-item>
          <el-form-item v-if="userInfo.auditStatus === 3" label="修改联系人名称" prop="changeContactName">
            <el-input v-model="userInfo.changeContactName" />
          </el-form-item>
          <el-form-item v-if="userInfo.auditStatus === 3" label="修改客户名称" prop="changeCustomerName">
            <el-input v-model="userInfo.changeCustomerName" />
          </el-form-item>
          <el-form-item label="审核状态:" prop="auditStatus">
            <el-select v-model="userInfo.auditStatus" filterable placeholder="Select">
              <el-option
                v-for="item in auditStatusOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="审核备注" prop="auditRemark">
            <el-input v-model="userInfo.auditRemark" placeholder="审核备注" />
          </el-form-item>
          <!--          <el-form-item label="邮箱" prop="email">
            <el-input v-model="userInfo.email" />
          </el-form-item>-->
          <el-form-item v-if="userStore.role.authorityId === 888" label="用户角色" prop="authorityId">
            <el-cascader
              v-model="userInfo.authorityIds"
              style="width:100%"
              :options="authOptions"
              :show-all-levels="false"
              :props="{ multiple:true,checkStrictly: true,label:'authorityName',value:'authorityId',disabled:'disabled',emitPath:false}"
              :clearable="false"
            />
          </el-form-item>
          <el-form-item label="启用" prop="disabled">
            <el-switch
              v-model="userInfo.enable"
              inline-prompt
              :active-value="1"
              :inactive-value="2"
            />
          </el-form-item>
          <el-form-item label="头像" label-width="80px">
            <div style="display:inline-block" @click="openHeaderChange">
              <img
                v-if="userInfo.headerImg"
                alt="头像"
                class="header-img-box"
                :src="(userInfo.headerImg && userInfo.headerImg.slice(0, 4) !== 'http')?path+userInfo.headerImg:userInfo.headerImg"
              >
              <div v-else class="header-img-box">从媒体库选择</div>
            </div>
          </el-form-item>
        </el-form>

      </div>

      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeAddUserDialog">取 消</el-button>
          <el-button type="primary" @click="enterAddUserDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
    <ChooseImg ref="chooseImg" :target="userInfo" :target-key="`headerImg`" />
  </div>
</template>

<script>
export default {
  name: 'User',
}
</script>

<script setup>

import {
  getUserList,
  setUserAuthorities,
  register,
  deleteUser,
} from '@/api/user'

import { getAuthorityList } from '@/api/authority'
import CustomPic from '@/components/customPic/index.vue'
import ChooseImg from '@/components/chooseImg/index.vue'
import { setUserInfo, resetPassword } from '@/api/user.js'

import { nextTick, ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { toSQLLine } from '@/utils/stringFun'
import { useUserStore } from '@/pinia/modules/user'
import { formatDate, getDictFunc, filterDict } from '@/utils/format'

const path = ref(import.meta.env.VITE_BASE_API + '/')
// 初始化相关
const setAuthorityOptions = (AuthorityData, optionsData) => {
  AuthorityData &&
  AuthorityData.forEach(item => {
    if (item.children && item.children.length) {
      const option = {
        authorityId: item.authorityId,
        authorityName: item.authorityName,
        children: [],
      }
      setAuthorityOptions(item.children, option.children)
      optionsData.push(option)
    } else {
      const option = {
        authorityId: item.authorityId,
        authorityName: item.authorityName,
      }
      optionsData.push(option)
    }
  })
}

const userStore = useUserStore()

const searchInfo = ref({})
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const data = {
    ...searchInfo.value,
    page: page.value,
    pageSize: pageSize.value,
  }
  const table = await getUserList(data)
  if (table.code === 0) {
    // 如果不是超级管理员
    if (userStore.role.authorityId !== 888) {
      table.data.list.forEach(item => {
        item.authorities.forEach(a => {
          // 判断是否是超级管理员
          if (a.authorityId === 888) {
            item.isSuperAdmin = true
          }
        })
      })
      console.log('table.data.list', table.data.list)
    }
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

const onReset = () => {
  searchInfo.value = {}
}

// 排序
const sortChange = ({ prop, order }) => {
  if (prop) {
    if (prop === 'ID') {
      prop = 'id'
    }
    searchInfo.value.orderKey = toSQLLine(prop)
    searchInfo.value.desc = order === 'descending'
  }
  getTableData()
}

watch(() => tableData.value, () => {
  setAuthorityIds()
})

const initPage = async() => {
  getTableData()
  const res = await getAuthorityList({ page: 1, pageSize: 999 })
  setOptions(res.data.list)
}

initPage()

const resetPasswordFunc = (row) => {
  ElMessageBox.confirm('是否将此用户密码重置为123456?', '警告', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  },
  ).then(async() => {
    const res = await resetPassword({
      ID: row.ID,
    })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: res.msg,
      })
    } else {
      ElMessage({
        type: 'error',
        message: res.msg,
      })
    }
  })
}
const setAuthorityIds = () => {
  tableData.value && tableData.value.forEach((user) => {
    user.authorityIds = user.authorities && user.authorities.map(i => {
      return i.authorityId
    })
  })
}

const chooseImg = ref(null)
const openHeaderChange = () => {
  chooseImg.value.open()
}

const authOptions = ref([])
const auditStatusOptions = ref([])
const setOptions = async(authData) => {
  authOptions.value = []
  setAuthorityOptions(authData, authOptions.value)
  auditStatusOptions.value = await getDictFunc('auditStatus')
}

const deleteUserFunc = async(row) => {
  const res = await deleteUser({ id: row.ID })
  if (res.code === 0) {
    ElMessage.success('删除成功')
    row.visible = false
    await getTableData()
  }
}

// 弹窗相关
const userInfo = ref({
  username: '',
  password: '',
  nickName: '',
  headerImg: '',
  authorityId: '',
  originContactName: '',
  originCustomerName: '',
  changeCustomerName: '',
  changeContactName: '',
  auditStatus: 0,
  auditRemark: '',
  authorityIds: [1000],
  enable: 1,
})

const rules = ref({
  userName: [
    { message: '请输入用户名, 不填则随机生成', trigger: 'blur' },
    { min: 5, message: '最低5位字符', trigger: 'blur' },
  ],
  password: [
    { required: true, message: '请输入用户密码', trigger: 'blur' },
    { min: 6, message: '最低6位字符', trigger: 'blur' },
  ],
  nickName: [
    { required: true, message: '请输入用户昵称', trigger: 'blur' },
  ],
  originContactName: [
    { required: true, message: '请输入联系人名称', trigger: 'blur' },
  ],
  phone: [
    {
      required: true,
      pattern: /^1([38][0-9]|4[014-9]|[59][0-35-9]|6[2567]|7[0-8])\d{8}$/,
      message: '请输入合法手机号',
      trigger: 'blur',
    },
  ],
  email: [
    {
      pattern: /^([0-9A-Za-z\-_.]+)@([0-9a-z]+\.[a-z]{2,3}(\.[a-z]{2})?)$/g,
      message: '请输入正确的邮箱',
      trigger: 'blur',
    },
  ],
  authorityId: [
    { required: true, message: '请选择用户角色', trigger: 'blur' },
  ],
})
const userForm = ref(null)
const enterAddUserDialog = async() => {
  userInfo.value.authorityId = userInfo.value.authorityIds[0]
  userForm.value.validate(async valid => {
    if (valid) {
      const req = {
        ...userInfo.value,
      }
      if (dialogFlag.value === 'add') {
        const res = await register(req)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '创建成功' })
          await getTableData()
          closeAddUserDialog()
        }
      }
      if (dialogFlag.value === 'edit') {
        const res = await setUserInfo(req)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '编辑成功' })
          await getTableData()
          closeAddUserDialog()
        }
      }
    }
  })
}

const addUserDialog = ref(false)
const closeAddUserDialog = () => {
  userForm.value.resetFields()
  userInfo.value.headerImg = ''
  userInfo.value.authorityIds = []
  addUserDialog.value = false
}

const dialogFlag = ref('add')

const addUser = () => {
  dialogFlag.value = 'add'
  addUserDialog.value = true
}

const tempAuth = {}
const changeAuthority = async(row, flag, removeAuth) => {
  if (flag) {
    if (!removeAuth) {
      tempAuth[row.ID] = [...row.authorityIds]
    }
    return
  }
  await nextTick()
  const res = await setUserAuthorities({
    ID: row.ID,
    authorityIds: row.authorityIds,
  })
  if (res.code === 0) {
    ElMessage({ type: 'success', message: '角色设置成功' })
  } else {
    if (!removeAuth) {
      row.authorityIds = [...tempAuth[row.ID]]
      delete tempAuth[row.ID]
    } else {
      row.authorityIds = [removeAuth, ...row.authorityIds]
    }
  }
}

const openEdit = (row) => {
  dialogFlag.value = 'edit'
  userInfo.value = JSON.parse(JSON.stringify(row))
  addUserDialog.value = true
}

const switchEnable = async(row) => {
  userInfo.value = JSON.parse(JSON.stringify(row))
  await nextTick()
  const req = {
    ...userInfo.value,
  }
  const res = await setUserInfo(req)
  if (res.code === 0) {
    ElMessage({ type: 'success', message: `${req.enable === 2 ? '禁用' : '启用'}成功` })
    await getTableData()
    userInfo.value.headerImg = ''
    userInfo.value.authorityIds = []
  }
}

const getStatusColorClass = (auditStatus) => {
  switch (auditStatus) {
    case 1:
      return 'green-text' // 定义绿色样式的 class
    case 2:
    case 3:
      return 'yellow-text' // 定义黄色样式的 class
    case 4:
      return 'dark-red-text' // 定义暗红色样式的 class
    default:
      return 'default-text' // 默认样式
  }
}
</script>

<style lang="scss">
.user-dialog {
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

.nickName {
  display: flex;
  justify-content: flex-start;
  align-items: center;
}

.pointer {
  cursor: pointer;
  font-size: 16px;
  margin-left: 2px;
}

/* 根据需要定义样式 */
.green-text {
  color: #66BB6A;
}

.yellow-text {
  color: #FB8C00;
}

.dark-red-text {
  color: #F4511E;
}

.default-text {
  color: #78909C;
}
</style>
