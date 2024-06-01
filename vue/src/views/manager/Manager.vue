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
                <el-menu-item index="/displaceline" class="submenu">
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
import axios from 'axios';
import { reactive, ref } from "vue";
import request from "@/utils/request";
const username = JSON.parse(localStorage.getItem('user') || '{}')
const token = JSON.parse(localStorage.getItem('token') || '{}')
//地图
var map, control;
var zoom = 17;
const data = reactive({
  tableData: ref([])
})
// 从后端获取数据
const fetchData = async () => {
  try {
    const response = await request.get('http://124.70.83.36:3000/api/v1/displacement/get', {
      headers: {
        Authorization: `Bearer ${token}` // 身份令牌
      },
    });
    data.tableData = response.data;
    console.log('传感器数据：', data.tableData);
  } catch (error) {
    console.error('Error fetching data:', error);
  }
};
function onLoad() {
  //初始化地图对象
  map = new T.Map("mapDiv");
  //设置显示地图的中心点和级别
  map.centerAndZoom(new T.LngLat(100.16090556,38.316356), zoom);
  //允许鼠标滚轮缩放地图
  map.enableScrollWheelZoom();

  var imageURL = "http://t0.tianditu.gov.cn/img_w/wmts?" +
    "SERVICE=WMTS&REQUEST=GetTile&VERSION=1.0.0&LAYER=img&STYLE=default&TILEMATRIXSET=w&FORMAT=tiles" +
    "&TILEMATRIX={z}&TILEROW={y}&TILECOL={x}&tk=9011ac4ded509603aed7dbec4fe8906b";
  //创建自定义图层对象
  var lay = new T.TileLayer(imageURL, { minZoom: 1, maxZoom: 18 });
  //将图层增加到地图上
  map.addLayer(lay);
  control = new T.Control.Zoom();
  //添加缩放平移控件
  map.addControl(control);
  // console.log('1111：', data.tableData);
  var filteredData79 = data.tableData.filter(item => item.sensor_name === '6079');
  var filteredData80 = data.tableData.filter(item => item.sensor_name === '6080');
  var filteredData81 = data.tableData.filter(item => item.sensor_name === '6081');
  var filteredData82 = data.tableData.filter(item => item.sensor_name === '6082');
  var filteredData83 = data.tableData.filter(item => item.sensor_name === '6083');
  // console.log('79：', filteredData79);
  var len = filteredData79.length-30
  var date = filteredData79[len+29].date;
  // console.log('nd:',date)
  var nd79 = filteredData79[len].nd;
  var ed79 = filteredData79[len].ed;
  var hd79 = filteredData79[len].hd;
  var nd80 = filteredData80[len].nd;
  var ed80 = filteredData80[len].ed;
  var hd80 = filteredData80[len].hd;
  var nd81 = filteredData81[len].nd;
  var ed81 = filteredData81[len].ed;
  var hd81 = filteredData81[len].hd;
  var nd82 = filteredData82[len].nd;
  var ed82 = filteredData82[len].ed;
  var hd82 = filteredData82[len].hd;
  var nd83 = filteredData83[len].nd;
  var ed83 = filteredData83[len].ed;
  var hd83 = filteredData83[len].hd;
  var nvel79 = filteredData79[len].nvel;
var evel79 = filteredData79[len].evel;
var hvel79 = filteredData79[len].hvel;

var nvel80 = filteredData80[len].nvel;
var evel80 = filteredData80[len].evel;
var hvel80 = filteredData80[len].hvel;

var nvel81 = filteredData81[len].nvel;
var evel81 = filteredData81[len].evel;
var hvel81 = filteredData81[len].hvel;

var nvel82 = filteredData82[len].nvel;
var evel82 = filteredData82[len].evel;
var hvel82 = filteredData82[len].hvel;

var nvel83 = filteredData83[len].nvel;
var evel83 = filteredData83[len].evel;
var hvel83 = filteredData83[len].hvel;

 
//数据
// var sensorData = [
//   {nd:}
// ]
 // 创建标记
 var markerData = [
  { sn: '21100200001271', number: 'JC01', type: 'GNSS', longitude: 100.16090556, latitude: 38.316356, nd: nd79, ed: ed79, hd: hd79, date: '2024-06-01', nvel: nvel79, evel: evel79, hvel: hvel79},
  { sn: '21100200001285', number: 'JC02', type: 'GNSS', longitude: 100.161530025, latitude: 38.3162559, nd: nd80, ed: ed80, hd: hd80, date: '2024-06-01', nvel: nvel80, evel: evel80,hvel: hvel80},
  { sn: '21100200001290', number: 'JC03', type: 'GNSS', longitude: 100.16090553, latitude: 38.315993, nd: nd81, ed: ed81, hd: hd81, date: '2024-06-01',nvel: nvel81, evel: evel81, hvel: hvel81},
  { sn: '21100200001270', number: 'JC04', type: 'GNSS', longitude: 100.16092616, latitude: 38.31568532, nd: nd82, ed: ed82, hd: hd82, date: '2024-06-01',nvel: nvel82, evel: evel82, hvel: hvel82 },
  { sn: '21100200001260', number: 'JC05', type: 'GNSS', longitude: 100.16110908, latitude: 38.31572915, nd: nd83, ed: ed83, hd: hd83, date: '2024-06-01',nvel: nvel83, evel: evel83, hvel: hvel83
 }
  // 可以继续添加更多标记数据
];
// console.log(markerData[0])
// 将标记添加到地图
markerData.forEach(function(data) {
  var marker = new T.Marker(new T.LngLat(data.longitude, data.latitude))
  map.addOverLay(marker);
  // 创建信息窗口内容
  var infoContent =
  "<div style='margin:10px 0px 10px 10px;'>设备SN: " + data.sn + "<br>设备编号: " + data.number + "<br>监测类型: " + data.type + "<br>采集日期: " + data.date + "<br>北方向位移: " + data.nd + "<br>东方向位移: " + data.ed + "<br>垂直方向位移: " + data.hd + "<br>北方向变化速率: " + data.nvel + "<br>东方向变化速率: " + data.evel + "<br>垂直方向变化速率: " + data.hvel +"</div>";

  // 创建信息窗口并设置内容
  var infoWindow = new T.InfoWindow();
  infoWindow.setContent(infoContent);

  // 监听标记点击事件
  marker.addEventListener("click", function() {
    marker.openInfoWindow(infoWindow);
  });
});


}
onMounted(async () => {
  await fetchData(); // 在组件挂载后立即执行 fetchData
  onLoad(); // 在组件挂载后初始化地图
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
</style>