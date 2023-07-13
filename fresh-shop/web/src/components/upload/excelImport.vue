<template>
  <div>
    <el-upload
      :action="`${path}/goods/batchCreateGoodsByExcel`"
      :before-upload="checkFile"
      :headers="{ 'x-token': userStore.token }"
      :on-error="uploadError"
      :on-success="uploadSuccess"
      :show-file-list="false"
      class="upload-btn"
    >
      <el-button icon="Upload" ref="uploadBtnRef" type="primary" :loading="fullscreenLoading">Excel 导入商品</el-button>
    </el-upload>
  </div>
</template>

<script setup>

import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'

const emit = defineEmits(['on-success'])
const path = ref(import.meta.env.VITE_BASE_API)

const userStore = useUserStore()
const fullscreenLoading = ref(false)

const checkFile = (file) => {
  fullscreenLoading.value = true
  const isExcel = file.type === 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet'
  const isLt2M = file.size / 1024 / 1024 < 150
  if (!isExcel) {
    ElMessage.error('上传文件只能是 Excel 格式!')
    fullscreenLoading.value = false
  }
  if (!isLt2M) {
    ElMessage.error('上传文件大小不能超过 150MB')
    fullscreenLoading.value = false
  }
  return isExcel && isLt2M
}

const uploadSuccess = (res) => {
  emit('on-success', res)
  fullscreenLoading.value = false
}

const uploadError = () => {
  ElMessage({
    type: 'error',
    message: '导入失败'
  })
  fullscreenLoading.value = false
}

const uploadBtnRef = ref(null)

const clickUploadBtn = () => {
  uploadBtnRef.value.$el.click()
}

defineExpose({ clickUploadBtn })

</script>

<script>

export default {
  name: 'ExcelImport',
  methods: {

  }
}
</script>
