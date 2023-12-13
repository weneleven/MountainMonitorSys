<template>
  <div>
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

    <div style="display: flex">
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
            <el-menu-item index="/manage">
              <span>项目地图</span>
            </el-menu-item>
          </el-sub-menu>
          <el-menu-item index="/course">
            <el-icon><Document /></el-icon>
            <span>项目信息</span>
          </el-menu-item>
          <el-menu-item index="/person">
            <el-icon><Tools /></el-icon>
            <span>传感器信息</span>
          </el-menu-item>
          <el-menu-item index="/person">
            <el-icon><Document /></el-icon>
            <span>数据查看</span>
          </el-menu-item>
          <el-menu-item index="/person">
            <el-icon><DataLine /></el-icon>
            <span>数据分析</span>
          </el-menu-item>
          <el-menu-item index="login" @click="logout">
            <el-icon><SwitchButton /></el-icon>
            <span>退出系统</span>
          </el-menu-item>
        </el-menu>
      </div>

      <div style="flex: 1; width: 0; background-color: #f8f8ff; padding: 10px">
        <!-- 百度地图容器 -->
        <div id="baidu-map" style="width: 100%; height: 100vh;"></div>
        <router-view />
      </div>
    </div>

  </div>
</template>

<script setup>
import { useRoute } from 'vue-router'
const $route = useRoute()
console.log($route.path)

import { onMounted } from 'vue';

const username = JSON.parse(localStorage.getItem('user') || '{}')
// 百度地图初始化
const initBaiduMap = () => {
  // 创建地图实例
  const map = new BMap.Map('baidu-map');

  // 设置地图中心点，这里以北京为例
  const point = new BMap.Point(116.404, 39.915);
  map.centerAndZoom(point, 5); // 设置缩放级别

  // 启用滚轮缩放
  map.enableScrollWheelZoom(true);
};

onMounted(() => {
  // 在组件挂载后初始化百度地图
  initBaiduMap();
});

const logout = () => {
  //localStorage.removeItem('student-user')
  window.location.href = '/'; // 重定向到登录页面
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
</style>