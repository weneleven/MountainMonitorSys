<template>
  <div class="common-layout">
    <el-container>
      <el-header>
        <div style="height: 60px; background-color:#ffffff; display: flex; align-items: center; border-bottom: 1px solid #ddd">
          <div style="flex: 1">
            <div style="padding-left: 20px; display: flex; align-items: center">
              <img src="@/assets/imgs/newlogo.jpg" alt="" style="width: 40px">
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
              <el-menu-item index="login" @click="logout">
                <el-icon><SwitchButton /></el-icon>
                <span>退出系统</span>
              </el-menu-item>
            </el-menu>
          </div>
        </el-aside>
        <el-main>
          <div>
            <!--所属项目切换-->
            <el-form
                :label-position="left"
                label-width="120px"
                style="max-width: 560px"
            >
              <el-form-item label="切换项目">
                <el-select v-model="value"  placeholder="Select" size="large">
                  <el-option
                      v-for="item in options"
                      :key="item.value"
                      :label="item.label"
                      :value="item.value"
                  />
                </el-select>
              </el-form-item>
            </el-form>
            <!--采集数据展示-->
            <div class="card" style="margin-bottom: 10px">

              <el-table :data="pagedTableData" stripe>
                <el-table-column label="设备名称" prop="sensor_name"></el-table-column>
                <el-table-column label="采集时间" prop="collect_time"></el-table-column>
                <el-table-column label="X坐标值(m)" prop="value_x"></el-table-column>
                <el-table-column label="Y坐标值(m)" prop="value_y"></el-table-column>
                <el-table-column label="H坐标值(m)" prop="value_h"></el-table-column>
                <el-table-column label="X方向累计(mm)" prop="sum_x"></el-table-column>
                <el-table-column label="Y方向累计(mm)" prop="sum_y"></el-table-column>
                <el-table-column label="H方向累计(mm)" prop="sum_h"></el-table-column>
                <el-table-column label="X方向本次(mm)" prop="now_x"></el-table-column>
                <el-table-column label="Y方向本次(mm)" prop="now_y"></el-table-column>
                <el-table-column label="H方向本次(mm)" prop="now_h"></el-table-column>
              </el-table>
            </div>
      <!--分页-->
            <div class="card">
              <el-pagination
                  background
                  layout="prev, pager, next"
                  :page-size="pageSize"
                  :current-page.sync="currentPage"
                  :total="totalItems"
                  @current-change="handlePageChange"
              />
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

const $route = useRoute()
console.log($route.path)

const username = JSON.parse(localStorage.getItem('user') || '{}')
const token = JSON.parse(localStorage.getItem('token')||'{}')

const logout = () => {
  window.location.href = '/'; // 重定向到登录页面
}
// 项目切换数据
const value = ref('山西临汾古贤皮带机项目')
const options = [
  {
    value:'山西临汾古贤皮带机项目',
    label:'山西临汾古贤皮带机项目',
  },
]
// 示例数据
const totalItems = ref(188); // 可根据后端接口修改，使总数据数由后端提供
const currentPage = ref(1);
const pageSize = ref(10);

const data = reactive({
  pageNum: 1,
  formVisible: false,
  form: {},
  //测试数据
  tableData: ref([])
})
onMounted(() => {
  fetchData(); // 在组件挂载时调用fetchData函数获取数据
});
// 从后端获取数据
const fetchData = async () => {
  try {
    const response = await request.get('http://124.70.83.36:3000/api/v1/data/get', {
      headers: {
        Authorization: `Bearer ${token}` // 身份令牌
      },
    });
    data.tableData = response.data;
  } catch (error) {
    console.error('Error fetching data:', error);
  }
};

// 计算当前页的数据
const pagedTableData = computed(() => {
  const start = (currentPage.value - 1) * pageSize.value;
  const end = start + pageSize.value;
  return data.tableData?.slice(start, end);
});

// 处理页码变化
const handlePageChange = (newPage) => {
  currentPage.value = newPage;
};


request.get('/').then(res => {
  console.log(res)
})
</script>

<style scoped>

.el-icon-check {
  color: green;
}

.el-icon-close {
  color: red;
}
</style>
