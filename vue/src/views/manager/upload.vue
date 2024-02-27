<template>
  <div class="common-layout">
    <el-container>
      <el-header>
        <div style="height: 60px; background-color:#ffffff; display: flex; align-items: center; border-bottom: 1px solid #ddd">
          <div style="flex: 1">
            <div style="padding-left: 20px; display: flex; align-items: center">
              <img src="@/assets/imgs/logo.png" alt="" style="width: 40px">
              <div style="font-weight: bold; font-size: 24px; margin-left: 5px">山体滑坡监测系统</div>
            </div>
          </div>
          <div style="width: fit-content; padding-right: 10px; display: flex; align-items: center;">
            <img src="https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png" alt="" style="width: 40px; height: 40px">
            <div style="margin-left: 5px">欢迎您，{{username}}</div>
          </div>
        </div>
      </el-header>
      <el-container>
        <el-aside width="200px">
          <div style="width: 200px; border-right: 1px solid #ddd; min-height: calc(100vh - 60px)">
            <el-menu
                router
                style="border: none"
                :default-active="$route.path"
                :default-openeds="['/home', '2']"
            >
              <el-menu-item index="/home">
                <el-icon><HomeFilled /></el-icon>
                <span>系统首页</span>
              </el-menu-item>
              <el-sub-menu index="2">
                <template #title>
                  <el-icon><MapLocation /></el-icon>
                  <span>数据视图</span>
                </template>
                <el-menu-item index="/manager">
                  <span>项目地图</span>
                </el-menu-item>
              </el-sub-menu>
              <el-menu-item index="/course">
                <el-icon><Document /></el-icon>
                <span>项目信息</span>
              </el-menu-item>
              <el-menu-item index="/sensor">
                <el-icon><Tools /></el-icon>
                <span>传感器信息</span>
              </el-menu-item>
              <el-menu-item index="/dataview">
                <el-icon><Document /></el-icon>
                <span>数据查看</span>
              </el-menu-item>
              <el-menu-item index="/analysis">
                <el-icon><DataLine /></el-icon>
                <span>数据分析</span>
              </el-menu-item>
              <el-menu-item >
                <el-icon><WarnTriangleFilled /></el-icon>
                <span>预测预警</span>
              </el-menu-item>
              <el-menu-item index="/upload">
                <el-icon><UploadFilled /></el-icon>
                <span>模拟数据发送</span>
              </el-menu-item>
              <el-menu-item index="login" @click="logout">
                <el-icon><SwitchButton /></el-icon>
                <span>退出系统</span>
              </el-menu-item>
            </el-menu>
          </div>
        </el-aside>
        <el-main>
          <div v-loading.fullscreen.lock="fullscreenLoading">
            <div class="gva-table-box">
              <div class="gva-btn-list">
                <upload-common
                    v-model:imageCommon="imageCommon"
                    @on-success="getTableData"
                />
                <el-input
                    v-model="search.keyword"
                    class="keyword"
                    placeholder="请输入文件名或备注"
                />
                <el-button
                    type="primary"
                    icon="search"
                    @click="getTableData"
                >查询</el-button>
              </div>

              <el-table :data="tableData" style="margin-top: 30px;">
                <el-table-column
                    align="left"
                    label="日期"
                    prop="UpdatedAt"
                    width="180"
                >
                </el-table-column>
                <el-table-column
                    align="left"
                    label="文件名/备注"
                    prop="file_name"
                    width="180"
                >
                </el-table-column>
                <el-table-column
                    align="left"
                    label="链接"
                    prop="url"
                    min-width="300"
                >
                </el-table-column>
                <el-table-column
                    align="left"
                    label="标签"
                    prop="extension"
                    width="100"
                >
                </el-table-column>
                <el-table-column
                    align="left"
                    label="操作"
                    width="160"
                >
                  <template #default="scope">
                    <el-button
                        icon="download"
                        type="primary"
                        link
                        @click="downloadFile(scope.row)"
                    >下载</el-button>
                    <el-button
                        icon="delete"
                        type="primary"
                        link
                        @click="deleteFileFunc(scope.row)"
                    >删除</el-button>
                  </template>
                </el-table-column>
              </el-table>
              <div class="gva-pagination">
                <el-pagination
                    :current-page="page"
                    :page-size="pageSize"
                    :page-sizes="[10, 30, 50, 100]"
                    :style="{ float: 'right', padding: '20px' }"
                    :total="total"
                    layout="total, sizes, prev, pager, next, jumper"
                    @current-change="handleCurrentChange"
                    @size-change="handleSizeChange"
                />
              </div>
            </div>
          </div>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script setup>
import request from "@/utils/request";
import { reactive, ref, onMounted, computed } from "vue";
import { ElMessageBox } from "element-plus";
import { useRoute } from 'vue-router';
import UploadCommon from '@/components/upload/common.vue'


const $route = useRoute()
console.log($route.path)
const username = JSON.parse(localStorage.getItem('user') || '{}')
const imageCommon = ref('')
const search = ref({})
const tableData = ref([])
//分页
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const token = JSON.parse(localStorage.getItem('token')||'{}')
const getTableData = async() => {
  try {
    const response = await request.get('http://124.70.83.36:3000/api/v1/data/getfileinfo', {
      headers: {
        Authorization: `Bearer ${token}` // 身份令牌
      },
    });
    const mappedData = response.data.map(item => ({
      UpdatedAt: item.UpdatedAt,
      file_name: item.file_name,
      url: item.url,
      extension: item.extension,
    }));
    tableData.value = mappedData
    total.value = response.total
}catch (error) {
    console.error('Error fetching data:', error);
  }
}
onMounted(() => {
  getTableData()
});
const editFileNameFunc = async(row) => {
  ElMessageBox.prompt('请输入文件名或者备注', '编辑', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    inputPattern: /\S/,
    inputErrorMessage: '不能为空',
    inputValue: row.name
  }).then(async({ value }) => {
    row.name = value
    // console.log(row)
    const res = await editFileName(row)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '编辑成功!',
      })
      getTableData()
    }
  }).catch(() => {
    ElMessage({
      type: 'info',
      message: '取消修改'
    })
  })
}
const downloadFile = (row) => {

}
const deleteFileFunc = async(row) => {

}
// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

</script>

<style scope>
.gva-btn-list {
  display: flex; /* 使用 Flex 布局 */
  gap: 20px;
  align-items: center; /* 垂直居中对齐子元素 */
}

.el-icon-check {
  color: green;
}

.el-icon-close {
  color: red;
}
</style>
