<template>
  <div>
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
//地图设置
const initBaiduMap = () => {
  // 创建地图实例
  const map = new BMapGL.Map('baidu-map');
  // 设置地图中心点，这里以北京为例
  const point = new BMapGL.Point(116.404, 39.915);
  map.centerAndZoom(point, 6); // 设置缩放级别
  // 启用滚轮缩放
  map.enableScrollWheelZoom(true);
  map.setMapStyleV2({
    styleId: '5b94ec7ba18095a1b8db2ed051cd7052'
  });
  var cityCtrl = new BMapGL.CityListControl();  // 添加城市列表控件
  map.addControl(cityCtrl);
  var point_shanxi = new BMapGL.Point(111.52,36.09);
  var point_liquan= new BMapGL.Point(108.42,34.48);
  var point_yunnan= new BMapGL.Point(99.40,23.52);
  var marker1 = new BMapGL.Marker(point_shanxi);        // 创建标注
  var marker2 = new BMapGL.Marker(point_liquan);
  var marker3 = new BMapGL.Marker(point_yunnan);
  map.addOverlay(marker1);                     // 将标注添加到地图中
  map.addOverlay(marker2);
  map.addOverlay(marker3);

// 创建信息窗口对象
  var opts = {
    width: 250, // 信息窗口宽度
    height: 150, // 信息窗口高度
    title: "项目信息", // 信息窗口标题
  };
  var infoWindow1 = new BMapGL.InfoWindow(
      `<p>项目名称：山西临汾古贤皮带机项目</p>`,
      opts
  );
// 监听事件
  marker1.addEventListener("click", function () {
    map.openInfoWindow(infoWindow1, point_shanxi);
  });
  var infoWindow2 = new BMapGL.InfoWindow(
      `<p>项目名称：231019东庄水利枢纽工程地灾监测</p>`,
      opts
  );

// 监听事件
  marker2.addEventListener("click", function () {
    map.openInfoWindow(infoWindow2, point_liquan);
  });
  var infoWindow3 = new BMapGL.InfoWindow(
      `<p>项目名称：230626云南耿马灌区项目安全监测</p>`,
      opts
  );

// 监听事件
  marker3.addEventListener("click", function () {
    map.openInfoWindow(infoWindow3, point_yunnan);
  });

};

onMounted(() => {
  // 在组件挂载后初始化百度地图
  initBaiduMap();
});

const logout = () => {
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