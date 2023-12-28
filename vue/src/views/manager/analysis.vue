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
<!-- 数据折线图           -->
            <div id="LC02" style="width: 1000px;height:400px;">
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
import * as echarts from 'echarts';


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
const data = reactive({
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
    // 过滤数据，只获取特定 sensor_name 的数据
    var targetSensorName = "JC02";
    var filteredData = data.tableData.filter(item => item.sensor_name === targetSensorName);
    //整理数据
    var collectTimes = filteredData.map(item => item.collect_time);
    var sumXData = filteredData.map(item => item.sum_x);
    var sumYData = filteredData.map(item => item.sum_y);
    var sumHData = filteredData.map(item => item.sum_h);
    const myChart = echarts.init(document.getElementById('LC02'));
    var option = {
      title: {
        text: "LC02"
      },
      tooltip: {
        trigger: 'axis'
      },
      legend: {    //图例组件
        data:['sum_x','sum_y','sum_h']
      },
      toolbox: {
        feature: {
          saveAsImage: {}
        }
      },
      xAxis: {
        type: 'category',
        boundaryGap: false,
        data: collectTimes
      },
      yAxis: {
        type: 'value'
      },
      series:  [
        {
          name: 'sum_x',
          type: 'line',
          stack: '总量',
          data: sumXData
        },
        {
          name: 'sum_y',
          type: 'line',
          stack: '总量',
          data: sumYData
        },
        {
          name: 'sum_h',
          type: 'line',
          stack: '总量',
          data: sumHData
        }
      ]
    };
    //4.使用刚指定的配置项和数据显示图表。
    myChart.setOption(option);
  } catch (error) {
    console.error('Error fetching data:', error);
  }
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
