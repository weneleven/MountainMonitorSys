<template>
  <div class="common-layout">
    <el-container>
      <el-header>
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
      </el-header>
      <el-container>
        <el-aside width="200px" class="menu">
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
        </el-aside>

        <el-main>
          <div>
            <!--所属项目切换-->
            <el-form :label-position="left" label-width="120px"
              style="max-width: 560px; display: flex; align-items: center;">
              <el-icon style="margin-top: -26px;">
                <Search />
              </el-icon>
              <el-form-item label="切换传感器" style="flex: 1;margin-left: -28px;font-weight: bold;margin-top: -2px;">
                <el-select v-model="sensor" placeholder="Select" size="large" style="margin-top: -5px;"
                  @change="handleSelect">
                  <el-option v-for="item in sensors" :key="item.value" :label="item.label" :value="item.value" />
                </el-select>
              </el-form-item>
            </el-form>
            <hr style="margin-top: -4px;">
            <!-- 数据折线图           -->
            <div id="line" style="width: 1000px;height:460px;margin-top: 20px;">
            </div>
            <div id="line1" style="width: 1000px;height:460px;margin-top: 20px;">
            </div>
            <div id="line2" style="width: 1000px;height:460px;margin-top: 20px;">
            </div>
            <div id="line3" style="width: 1000px;height:460px;margin-top: 20px;">
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
const token = JSON.parse(localStorage.getItem('token') || '{}')

const logout = () => {
  window.location.href = '/'; // 重定向到登录页面
}
// 项目切换数据
const sensor = ref('6079')
const sensors = [
  {
    value: '6079',
    label: 'JC01',
  },
  {
    value: '6080',
    label: 'JC02'
  },
  {
    value: '6081',
    label: 'JC03'
  },
  {
    value: '6082',
    label: 'JC04'
  },
  {
    value: '6083',
    label: 'JC05'
  },
]
const data = reactive({
  tableData: ref([])
})
const handleSelect = () => {
  fetchData(sensor.value)
  ElMessage({
    message: 'success',
    type: 'success',
  })
}
onMounted(() => {
  fetchData(sensor.value);
});
// 从后端获取数据
const fetchData = async (sensorValue) => {
  try {
    const response = await request.get('http://124.70.83.36:3000/api/v1/displacement/get', {
      headers: {
        Authorization: `Bearer ${token}` // 身份令牌
      },
    });
    data.tableData = response.data;
    var filteredData1 = data.tableData.filter(item => item.sensor_name === sensorValue);
    var collectTimes1 = filteredData1.map(item => item.date);
    var nd = filteredData1.map(item => item.nd);
    var ed = filteredData1.map(item => item.ed);
    var hd = filteredData1.map(item => item.hd);
    var len = nd.length;
    //预测未来30天
    var start = len-30;
    var end = len-1;
    //效果总览
    const myChart = echarts.init(document.getElementById('line'));
    var option = {
      title: {
        text: "位移预测图",
        // text: sensorValue + "：位移图",
        left: 'center',
      },
      grid: {
        left: 10,
        right: 10,
        bottom: 20,
        top: 30,
        containLabel: true
      },
      tooltip: {
        trigger: 'axis',
        axisPointer: {
          type: 'cross'
        },
        padding: [5, 10]
      },
      toolbox: {
        feature: {
          saveAsImage: {}
        }
      },
      xAxis: {
        type: 'category',
        boundaryGap: false,
        data: collectTimes1
      },
      yAxis: {
        type: 'value'
      },
      visualMap: {
        type: 'piecewise',
        show: false,
        dimension: 0,
        seriesIndex: 0,
        pieces: [
          {
            gt: start,
            lt: end,
            color: 'rgba(190, 184, 220, 0.4)'
          },
        ]
      },
      series: [
        {
          name: 'nd',
          itemStyle: {
            normal: {
              color: '#BE88DC',
              lineStyle: {
                color: '#BE88DC',
                width: 2
              }
            }
          },
          type: 'line',
          stack: '总量',
          data: nd,
          smooth: true,
          animationDuration: 2800,
          animationEasing: 'cubicInOut'
        },
        {
          name: 'ed',
          type: 'line',
          stack: '总量',
          data: ed,
          smooth: true,
          itemStyle: {
            normal: {
              color: '#FFBE7A',
              lineStyle: {
                color: '#FFBE7A',
                width: 2
              },
              areaStyle: {
                color: '#f3f8ff'
              }
            }
          },
          animationDuration: 2800,
          animationEasing: 'quadraticOut'
        },
        {
          name: 'hd',
          type: 'line',
          data: hd,
          smooth: true,
          itemStyle: {
            normal: {
              color: '#8ECFC9',
              lineStyle: {
                color: '#8ECFC9',
                width: 2
              },
              areaStyle: {
                color: '#f3f8ff'
              }
            },

          },
          markLine: {
            symbol: ['none', 'none'],
            data: [
              {
                xAxis: start,
                label: {
                  show: true,
                  position: 'end',
                  formatter: function (params) {
                    return 'forecast';
                  }
                }
              },
              {
                xAxis: end,
                label: { show: false }
              }
            ]
          },
          animationDuration: 2800,
          animationEasing: 'quadraticOut'
        }
      ]
    };
    myChart.setOption(option);
    //北方向位移
    const myChart1 = echarts.init(document.getElementById('line1'));
    var option1 = {
      title: {
        text: "北方向位移图",
        // text: sensorValue + "：北方向位移图",
        left: 'center',
      },
      grid: {
        left: 10,
        right: 10,
        bottom: 20,
        top: 30,
        containLabel: true
      },
      tooltip: {
        trigger: 'axis',
        axisPointer: {
          type: 'cross'
        },
        padding: [5, 10]
      },
      toolbox: {
        feature: {
          saveAsImage: {}
        }
      },
      xAxis: {
        type: 'category',
        boundaryGap: false,
        data: collectTimes1
      },
      yAxis: {
        type: 'value'
      },
      visualMap: {
        type: 'piecewise',
        show: false,
        dimension: 0,
        seriesIndex: 0,
        pieces: [
          {
            gt: start,
            lt: end,
            color: 'rgba(190, 184, 220, 0.4)'
          },
        ]
      },
      series: [
        {
          name: '北方向位移',
          lineStyle: {
            color: '#BE88DC',
            width: 5
          },
          markLine: {
            symbol: ['none', 'none'],
            data: [
              {
                xAxis: start,
                label: {
                  show: true,
                  position: 'end',
                  formatter: function (params) {
                    return 'forecast';
                  }
                }
              },
              {
                xAxis: end,
                label: { show: false }
              }
            ]
          },
          areaStyle: {
          },
          type: 'line',
          stack: '总量',
          data: nd,
          smooth: true,
          animationDuration: 2800,
          animationEasing: 'cubicInOut'
        },
      ]
    };
    myChart1.setOption(option1);
    //东方向位移
    const myChart2 = echarts.init(document.getElementById('line2'));
    var option2 = {
      title: {
        text: "东方向位移图",
        // text: sensorValue + "：东方向位移图",
        left: 'center',
      },
      grid: {
        left: 10,
        right: 10,
        bottom: 20,
        top: 30,
        containLabel: true
      },
      tooltip: {
        trigger: 'axis',
        axisPointer: {
          type: 'cross'
        },
        padding: [5, 10]
      },
      toolbox: {
        feature: {
          saveAsImage: {}
        }
      },
      xAxis: {
        type: 'category',
        boundaryGap: false,
        data: collectTimes1
      },
      yAxis: {
        type: 'value'
      },
      visualMap: {
        type: 'piecewise',
        show: false,
        dimension: 0,
        seriesIndex: 0,
        pieces: [
          {
            gt: start,
            lt: end,
            color: 'rgba(250, 127, 111, 0.4)'
          },
        ]
      },
      series: [
        {
          name: '东方向位移',
          lineStyle: {
            color: '#FFBE7A',
            width: 5
          },
          markLine: {
            symbol: ['none', 'none'],
            data: [
              {
                xAxis: start,
                label: {
                  show: true,
                  position: 'end',
                  formatter: function (params) {
                    return 'forecast';
                  }
                }
              },
              {
                xAxis: end,
                label: { show: false }
              }
            ]
          },
          areaStyle: {},
          type: 'line',
          stack: '总量',
          data: ed,
          smooth: true,
          animationDuration: 2800,
          animationEasing: 'cubicInOut'
        },
      ]
    };
    myChart2.setOption(option2);
    //北方向位移
    const myChart3 = echarts.init(document.getElementById('line3'));
    var option3 = {
      title: {
        text: "垂直方向位移图",
        // text: sensorValue + "：垂直方向位移图",
        left: 'center',
      },
      grid: {
        left: 10,
        right: 10,
        bottom: 20,
        top: 30,
        containLabel: true
      },
      tooltip: {
        trigger: 'axis',
        axisPointer: {
          type: 'cross'
        },
        padding: [5, 10]
      },
      toolbox: {
        feature: {
          saveAsImage: {}
        }
      },
      xAxis: {
        type: 'category',
        boundaryGap: false,
        data: collectTimes1
      },
      yAxis: {
        type: 'value'
      },
      visualMap: {
        type: 'piecewise',
        show: false,
        dimension: 0,
        seriesIndex: 0,
        pieces: [
          {
            gt: start,
            lt: end,
            color: 'rgba(142, 207, 201, 0.4)'
          },
        ]
      },
      series: [
        {
          name: '垂直方向位移',
          lineStyle: {
            color: '#8ECFC9',
            width: 5
          },
          markLine: {
            symbol: ['none', 'none'],
            data: [
              {
                xAxis: start,
                label: {
                  show: true,
                  position: 'end',
                  formatter: function (params) {
                    return 'forecast';
                  }
                }
              },
              {
                xAxis: end,
                label: { show: false }
              }
            ]
          },
          areaStyle: {},
          type: 'line',
          stack: '总量',
          data: hd,
          smooth: true,
          animationDuration: 2800,
          animationEasing: 'cubicInOut'
        },
      ]
    };
    myChart3.setOption(option3);
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

.label_style {
  font-size: 200px;
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