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
        <div id="mapDiv" style="width: 100%; height: 100vh;"></div>
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
//地图
var map,control,marker;
var zoom = 16;
function onLoad() {
  //初始化地图对象
  map = new T.Map("mapDiv");
  //设置显示地图的中心点和级别
  map.centerAndZoom(new T.LngLat(110.691386273,35.803706204), zoom);
  //允许鼠标滚轮缩放地图
  map.enableScrollWheelZoom();

  var imageURL = "http://t0.tianditu.gov.cn/img_w/wmts?" +
      "SERVICE=WMTS&REQUEST=GetTile&VERSION=1.0.0&LAYER=img&STYLE=default&TILEMATRIXSET=w&FORMAT=tiles" +
      "&TILEMATRIX={z}&TILEROW={y}&TILECOL={x}&tk=9011ac4ded509603aed7dbec4fe8906b";
  //创建自定义图层对象
  var lay = new T.TileLayer(imageURL, {minZoom: 1, maxZoom: 18});
  //将图层增加到地图上
  map.addLayer(lay);
  control = new T.Control.Zoom();
  //添加缩放平移控件
  map.addControl(control);
  //创建标注对象
  marker = new T.Marker(new T.LngLat(110.691386273,35.803706204));
  map.addOverLay(marker);
  var infoWin1 = new T.InfoWindow();
  var sContent =
      "<div style='margin:0px;'>" +
      "<div style='margin:10px 10px; '>" +
      "<div style='margin:10px 0px 10px 10px;'>设备SN: 21100200001109 <br>设备名称: JC07 <br>监测类型: GNSS" +
      "</div>" +
      "</div>";
  infoWin1.setContent(sContent);
  marker.addEventListener("click", function () {
    marker.openInfoWindow(infoWin1);
  });// 将标注添加到地图中
}
onMounted(() => {
  // 在组件挂载后初始化地图
  onLoad()
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