<template>
  <div>
    <div
      style="height: 60px; background-color:#ffffff; display: flex; align-items: center; border-bottom: 1px solid #ddd">
      <div style="flex: 1;">
        <div style="display: flex; align-items: center">
          <img src="@/assets/imgs/newlogo.jpg" alt="" style="width: 40px">
          <div style="font-weight: bold; font-size: 24px; margin-left: 5px">山体滑坡监测系统</div>
        </div>
      </div>
      <div style="width: fit-content; padding-right: 10px; display: flex; align-items: center;">
        <img src="https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png" alt=""
          style="width: 40px; height: 40px">
        <div style="margin-left: 5px">欢迎您，{{ username }}</div>
      </div>
    </div>
    <div style="display: flex" class="menu">
      <div style="width: 200px; border-right: 1px solid #ddd; min-height: calc(100vh - 60px);">
        <el-menu router class="menu" style="border: none;" :default-active="$route.path"
          :default-openeds="['/home', '2']">
          <el-menu-item index="/home">
            <el-icon>
              <HomeFilled />
            </el-icon>
            <span>系统首页</span>
          </el-menu-item>
          <el-sub-menu index="2">
            <template #title>
              <el-icon class="submenu-title">
                <MapLocation />
              </el-icon>
              <span class="submenu-title">数据视图</span>
            </template>
            <el-menu-item index="/infowindow" class="submenu">
              <span>项目地图</span>
            </el-menu-item>
            <el-menu-item index="/manager" class="submenu">
              <span>传感器位点</span>
            </el-menu-item>
          </el-sub-menu>
          <el-menu-item index="/course">
            <el-icon>
              <Document />
            </el-icon>
            <span>项目信息</span>
          </el-menu-item>
          <el-menu-item index="/sensor">
            <el-icon>
              <Tools />
            </el-icon>
            <span>传感器信息</span>
          </el-menu-item>
          <el-menu-item index="/dataview">
            <el-icon>
              <Document />
            </el-icon>
            <span>数据查看</span>
          </el-menu-item>
          <el-menu-item index="/analysis">
            <el-icon>
              <DataLine />
            </el-icon>
            <span>数据分析</span>
          </el-menu-item>
          <el-sub-menu index="2" :collapse="true">
            <template #title>
              <el-icon class="submenu-title">
                <WarnTriangleFilled />
              </el-icon>
              <span class="submenu-title">预测预警</span>
            </template>
            <el-menu-item index="/warnview" class="submenu">
              <span>预警信息</span>
            </el-menu-item>
            <el-menu-item index="warnanalysis" class="submenu">
              <span>预警分析</span>
            </el-menu-item>
            <el-menu-item index="/dashboard" class="submenu">
              <span>预测</span>
            </el-menu-item>
          </el-sub-menu>
          <el-menu-item index="/upload">
            <el-icon>
              <UploadFilled />
            </el-icon>
            <span>模拟数据发送</span>
          </el-menu-item>
          <el-menu-item index="login" @click="logout">
            <el-icon>
              <SwitchButton />
            </el-icon>
            <span>退出系统</span>
          </el-menu-item>
        </el-menu>
      </div>

      <div style="flex: 1; width: 0; background-color: #f8f8ff; padding: 10px;position: relative;">
        <!-- 折线图 -->
        <el-row style="background:#fff;padding:16px 16px 0;margin-bottom:32px;">
          <line-chart :chart-data="lineChartData" />
        </el-row>
        <!-- 其他 -->
        <el-row :gutter="32">
          <el-col :xs="24" :sm="24" :lg="8">
            <div class="chart-wrapper">
              <h1>radder</h1>
            </div>
          </el-col>
          <el-col :xs="24" :sm="24" :lg="8">
            <div class="chart-wrapper">
              <h1>pie</h1>
            </div>
          </el-col>
          <el-col :xs="24" :sm="24" :lg="8">
            <div class="chart-wrapper">
              <bar-chart />
            </div>
          </el-col>
        </el-row>
        <el-row>
          <el-col :xs="{ span: 24 }" :sm="{ span: 24 }" :md="{ span: 24 }" :lg="{ span: 12 }" :xl="{ span: 12 }"
            style="padding-right:8px;margin-bottom:30px;">
            <h1>1</h1>
          </el-col>
          <el-col :xs="{ span: 24 }" :sm="{ span: 12 }" :md="{ span: 12 }" :lg="{ span: 6 }" :xl="{ span: 6 }"
            style="margin-bottom:30px;">
            <h1>2</h1>
          </el-col>
          <el-col :xs="{ span: 24 }" :sm="{ span: 12 }" :md="{ span: 12 }" :lg="{ span: 6 }" :xl="{ span: 6 }"
            style="margin-bottom:30px;">
            <h1>3</h1>
          </el-col>
        </el-row>
      </div>
    </div>

  </div>
</template>

<script>
import BarChart from './components/BarChart.vue'
import LineChart from './components/LineChart.vue'
const lineChartData = {
  newVisitis: {
    expectedData: [100, 120, 161, 134, 105, 160, 165],
    actualData: [120, 82, 91, 154, 162, 140, 145]
  },
  messages: {
    expectedData: [200, 192, 120, 144, 160, 130, 140],
    actualData: [180, 160, 151, 106, 145, 150, 130]
  },
  purchases: {
    expectedData: [80, 100, 121, 104, 105, 90, 100],
    actualData: [120, 90, 100, 138, 142, 130, 130]
  },
  shoppings: {
    expectedData: [130, 140, 141, 142, 145, 150, 160],
    actualData: [120, 82, 91, 154, 162, 140, 130]
  }
}
export default {
  name: 'DashboardAdmin',
  components: {
    BarChart,
    LineChart,
  },
  data() {
    return {
      lineChartData: lineChartData.newVisitis
    }
  },
}
</script>

<style scoped>
.el-menu-item.is-active {
  background-color: #dcede9 !important;
}

.el-menu-item:hover {
  color: #11A983;
}

.el-table th {
  color: #333;
}

.menu {
  background-color: #2a3852;
}

.submenu {
  background-color: #232c45;
}

.el-menu-item {
  color: white;
}

.submenu-title {
  color: white;
  /* 设置标题的颜色 */
}

.chart-wrapper {
  background: #fff;
  padding: 16px 16px 0;
  margin-bottom: 32px;
}
</style>