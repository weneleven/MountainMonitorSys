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
const username = JSON.parse(localStorage.getItem('user') || '{}')
//地图

function onLoad() {
    var map = new BMapGL.Map("mapDiv");
    var point = new BMapGL.Point(110.691386273, 35.803706204);
    map.centerAndZoom(point, 17);
    map.enableScrollWheelZoom(true);
    map.enableKeyboard();
    var scaleCtrl = new BMapGL.ScaleControl();  // 添加比例尺控件
    map.addControl(scaleCtrl);
    var zoomCtrl = new BMapGL.ZoomControl();  // 添加缩放控件
    map.addControl(zoomCtrl);
    map.setMapType(BMAP_SATELLITE_MAP);  //使用卫星地图
    var marker = new BMapGL.Marker(point);        // 创建标注   
    map.addOverlay(marker);                     // 将标注添加到地图中
    // 创建自定义覆盖物
    var customOverlay = new BMapGL.CustomOverlay(createDOM, {
        point: new BMapGL.Point(110.68272, 35.80558),
        opacity: 0.5,
        offsetY: -60,
        properties: {
            title: '项目总览',
            text: ['项目名称: 230505黄藏寺左岸BXT5变形监测','项目地址：祁连县八宝镇','项目联系人：王宏飞 15890198888'],
            imgSrc: 'https://bj.bcebos.com/v1/mapopen-pub-jsapigl/assets/images/gugong.png'
        }
    });
    // 将自定义覆盖物添加到地图上
    map.addOverlay(customOverlay);
    customOverlay.addEventListener('click', function (e) {
        var data = e.target.properties.title;
        alert(data);
    });
    //警告覆盖物
    var warnOverlay = new BMapGL.CustomOverlay(createwarnDOM, {
        point: new BMapGL.Point(110.68272, 35.8000),
        opacity: 0.5,
        offsetY: -60,
        properties: {
            title: '预警信息',
            text: ['当天警告记录0条',],
            imgSrc: 'https://bj.bcebos.com/v1/mapopen-pub-jsapigl/assets/images/gugong.png'
        }
    });
    // 将自定义覆盖物添加到地图上
    map.addOverlay(customOverlay);
    map.addOverlay(warnOverlay);
}
//
function createDOM() {
    var div = document.createElement('div');
    div.style.zIndex = BMapGL.Overlay.getZIndex(this.point.lat);
    div.style.backgroundColor = 'rgba(0, 0, 255, 0.5)'; // 蓝色背景，设置透明度为0.5
    div.style.color = '#fff'; // 白色文字颜色
    div.style.height = '300px';
    div.style.width = '500px';
    div.style.padding = '2px';
    div.style.lineHeight = '50px';
    div.style.whiteSpace = 'nowrap';
    div.style.MozUserSelect = 'none';
    div.style.fontSize = '12px';
    div.style.borderRadius = '10px';
    div.style.display = 'flex';
    div.style.alignItems = 'flex-start';
div.style.justifyContent = 'space-between'; // 左右两端对齐
    div.style.flexDirection = 'column';

    var title = document.createElement('div');
    title.style.display = 'block';
    title.style.lineHeight = '16px';
    title.style.fontSize = '24px';
    title.style.fontWeight = '700';
    title.style.padding = '20px';
    div.appendChild(title);
    title.appendChild(document.createTextNode(this.properties.title));

    var span1 = document.createElement('span');
    span1.style.wordWrap = 'break-word';
    span1.style.lineHeight = '10px';
    span1.style.wordWrap = 'break-word';
    span1.style.fontSize = '18px';
    span1.style.whiteSpace = 'normal';
    span1.style.padding = '30px';
    span1.style.color = '#fff';
    span1.style.fontWeight = '400'; // 加粗字体
    span1.appendChild(document.createTextNode(this.properties.text[0]));
    div.appendChild(span1);

    var span2 = document.createElement('span');
    span2.style.wordWrap = 'break-word';
    span2.style.lineHeight = '10px';
    span2.style.wordWrap = 'break-word';
    span2.style.fontSize = '18px';
    span2.style.whiteSpace = 'normal';
    span2.style.padding = '30px';
    span2.style.color = '#fff';
    span2.style.fontWeight = '400'; // 加粗字体
    span2.appendChild(document.createTextNode(this.properties.text[1]));
    div.appendChild(span2);


    var span3 = document.createElement('span');
    span3.style.wordWrap = 'break-word';
    span3.style.lineHeight = '10px';
    span3.style.wordWrap = 'break-word';
    span3.style.fontSize = '18px';
    span3.style.whiteSpace = 'normal';
    span3.style.padding = '30px';
    span3.style.color = '#fff';
    span3.style.fontWeight = '400'; // 加粗字体
    span3.appendChild(document.createTextNode(this.properties.text[2]));
    div.appendChild(span3);

    // let img = document.createElement('img');
    // img.style.width = '120px';
    // img.src = this.properties.imgSrc;
    // div.appendChild(img);

    div.onmouseover = function () {
        this.style.backgroundColor = 'skyblue'; this.style.color = '#fff';
        span.style.color = '#fff';
    };

    div.onmouseout = function () {
        this.style.backgroundColor = 'rgba(0, 0, 255, 0.5)';
        this.style.color = '#fff';
        span.style.color = '#fff';
    };
    return div;
}
function createwarnDOM() {
    var div = document.createElement('div');
    div.style.zIndex = BMapGL.Overlay.getZIndex(this.point.lat);
    div.style.backgroundColor = 'rgba(0, 0, 255, 0.5)'; // 蓝色背景，设置透明度为0.5
    div.style.color = '#fff'; // 白色文字颜色
    div.style.height = '100px';
    div.style.width = '500px';
    div.style.padding = '2px';
    div.style.lineHeight = '50px';
    div.style.whiteSpace = 'nowrap';
    div.style.MozUserSelect = 'none';
    div.style.fontSize = '12px';
    div.style.borderRadius = '10px';
    div.style.display = 'flex';
    div.style.alignItems = 'flex-start';
div.style.justifyContent = 'space-between'; // 左右两端对齐
    div.style.flexDirection = 'column';

    var title = document.createElement('div');
    title.style.display = 'block';
    title.style.lineHeight = '16px';
    title.style.fontSize = '24px';
    title.style.fontWeight = '700';
    title.style.padding = '20px';
    div.appendChild(title);
    title.appendChild(document.createTextNode(this.properties.title));

    var span1 = document.createElement('span');
    span1.style.wordWrap = 'break-word';
    span1.style.lineHeight = '10px';
    span1.style.wordWrap = 'break-word';
    span1.style.fontSize = '18px';
    span1.style.whiteSpace = 'normal';
    span1.style.padding = '20px';
    span1.style.color = '#fff';
    span1.style.fontWeight = '400'; // 加粗字体
    span1.appendChild(document.createTextNode(this.properties.text[0]));
    div.appendChild(span1);

    
    // let img = document.createElement('img');
    // img.style.width = '120px';
    // img.src = this.properties.imgSrc;
    // div.appendChild(img);

    div.onmouseover = function () {
        this.style.backgroundColor = 'skyblue'; this.style.color = '#fff';
        span1.style.color = '#fff';
    };

    div.onmouseout = function () {
        this.style.backgroundColor = 'rgba(0, 0, 255, 0.5)';
        this.style.color = '#fff';
        span1.style.color = '#fff';
    };
    return div;
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