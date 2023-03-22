<template>
  <div>
    <el-upload
      v-model:file-list="fileListModel"
      list-type="picture-card"
      :action="`${path}/fileUploadAndDownload/upload`"
      :before-upload="checkFile"
      :headers="{ 'x-token': userStore.token }"
      :on-error="uploadError"
      :on-success="uploadSuccess"
      :on-preview="handlePictureCardPreview"
      :on-remove="handleRemove"
    >
      <el-icon><Plus /></el-icon>
    </el-upload>
    <el-dialog v-model="dialogVisible">
      <img style="width: 100%" :src="dialogImageUrl" alt="Preview Image" />
    </el-dialog>
  </div>
</template>

<script setup>

import { computed, ref, watch, defineProps, defineEmits } from 'vue'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'
import { Plus } from '@element-plus/icons-vue'

const emit = defineEmits(['on-success', 'update:fileList'])
const path = ref(import.meta.env.VITE_BASE_API)

const userStore = useUserStore()
const fullscreenLoading = ref(false)

const props = defineProps({
  fileList: {
    type: Array,
    default: () => []
  }
})

const fileListModel = computed({
  get: () => props.fileList,
  set: (val) => {
  },
})

watch(() => props.fileList, () =>{
  console.log('props.fileList', props.fileList)
}, { deep: true })

const checkFile = (file) => {
  fullscreenLoading.value = true
  const isJPG = file.type === 'image/jpeg'
  const isPng = file.type === 'image/png'
  const isLt2M = file.size / 1024 / 1024 < 2
  if (!isJPG && !isPng) {
    ElMessage.error('上传图片只能是 jpg或png 格式!')
    fullscreenLoading.value = false
  }
  if (!isLt2M) {
    ElMessage.error('未压缩的上传图片大小不能超过 500KB，请使用压缩上传')
    fullscreenLoading.value = false
  }
  return (isPng || isJPG) && isLt2M
}

const uploadSuccess = (res) => {
  const { data } = res
  if (data.file) {
    let url = data.file.url
    // 如果不是 http 开头的
    if (url.slice(0, 4) !== 'http') {
      url = path.value + '/' + url
    }
    emit('on-success', url, data.file.name)
  }
}

const uploadError = () => {
  ElMessage({
    type: 'error',
    message: '上传失败'
  })
  fullscreenLoading.value = false
}

const dialogImageUrl = ref('')
const dialogVisible = ref(false)

const handleRemove = (uploadFile, uploadFiles) => {
  const files = []
  uploadFiles.forEach(item => {
    let url = item.response.data.file.url
    // 如果不是 http 开头的
    if (url.slice(0, 4) !== 'http') {
      url = path.value +"/"+ url
    }
    files.push({
      name: item.name,
      url: url
    })
  })
  emit("update:fileList", files)
}

const handlePictureCardPreview = (uploadFile) => {
  dialogImageUrl.value = uploadFile.url
  dialogVisible.value = true
}

</script>

<script>

export default {
  name: 'ImageList',
  methods: {

  }
}
</script>
