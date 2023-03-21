<template>
  <div style="width: 100%">
    <CommonUpload ref="uploadImageRef" @on-success="uploadSuccess" style="display: none" />
    <QuillEditor ref="editorRef" v-model:content="editorContent" content-type="html" theme="snow" :options="options" />
    <div>内容：{{ content }}</div>
  </div>
</template>

<script>
export default {
  components: 'Editor',
}
</script>
<script setup>
import { QuillEditor, Quill } from '@vueup/vue-quill'
import CommonUpload from '@/components/upload/common.vue'
import '@vueup/vue-quill/dist/vue-quill.snow.css'
import { computed, reactive, ref, toRaw } from 'vue'
import { ElMessage } from 'element-plus'

const props = defineProps({
  content: {
    type: String,
    default: ''
  },
})
const emit = defineEmits(['update:content'])

const editorContent = computed({
  get: () => props.content,
  set: (val) => {
    emit('update:content', val)
  }
})

const uploadImageRef = ref(null)

const options = reactive({
  theme: 'snow',
  debug: 'warn',
  modules: {
    // 工具栏配置
    toolbar: {
      container: [
        ['bold', 'italic', 'underline', 'strike'], // 加粗 斜体 下划线 删除线
        ['blockquote', 'code-block'], // 引用  代码块
        [{ list: 'ordered' }, { list: 'bullet' }], // 有序、无序列表
        [{ indent: '-1' }, { indent: '+1' }], // 缩进
        [{ size: ['small', false, 'large', 'huge'] }], // 字体大小
        [{ header: [1, 2, 3, 4, 5, 6, false] }], // 标题
        [{ color: [] }, { background: [] }], // 字体颜色、字体背景颜色
        [{ align: [] }], // 对齐方式
        ['clean'], // 清除文本格式
        ['link', 'image'], // 链接、图片、视频
      ],
      handlers: {
        // 重写图片上传事件
        image: function(value) {
          console.log(233333333333333, value)
          if (value) {
            // 调用图片上传
            console.log(uploadImageRef.value)
            uploadImageRef.value.clickUploadBtn()
          } else {
            Quill.format('image', true)
          }
        },
      },
    },
  },
  placeholder: '请输入内容',
  readOnly: props.readOnly,
})

const editorRef = ref(null)
const path = ref(import.meta.env.VITE_BASE_API + '/')
// 上传图片成功进行处理
const uploadSuccess = (url) => {
  console.log('uploadSuccess orgin url:', url)
  if (!url || url === '') {
    ElMessage.error('获取图片地址失败')
    return
  }
  // 如果不是 http 开头的
  if (url.slice(0, 4) !== 'http') {
    url = path.value + '/' + url
  }
  // 获取富文本实例
  const quill = toRaw(editorRef.value).getQuill()
  // 获取光标位置
  const length = quill.selection.savedRange.index
  // 插入图片，res为服务器返回的图片链接地址
  quill.insertEmbed(length, 'image', url)
  // 调整光标到最后
  quill.setSelection(length + 1)
}

// 获取文本内容
const getContent = () => {
  return props.content
}

defineExpose({
  getContent
})
</script>

<style>
.ql-editor {
  width: 1000px;
  min-height: 300px;
}
</style>
