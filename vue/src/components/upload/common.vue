<template>
  <div>
    <el-upload
        :action="'http://124.70.83.36:3000/api/v1/data/upload'"
        :on-error="uploadError"
        :on-success="uploadSuccess"
        :show-file-list="false"
        class="upload-btn"
    >
      <el-button type="primary">普通上传</el-button>
    </el-upload>
  </div>
</template>

<script setup>

import { ref } from 'vue'
import { ElMessage } from 'element-plus'

defineOptions({
  name: 'UploadCommon',
})

const emit = defineEmits(['on-success'])
const path = ref(import.meta.env.VITE_BASE_API)

const fullscreenLoading = ref(false)


const uploadSuccess = (res) => {
  const { data } = res
  if (data.url) {
    emit('on-success', data.url)
  }
}

const uploadError = () => {
  ElMessage({
    type: 'error',
    message: '上传失败'
  })
  fullscreenLoading.value = false
}

</script>

