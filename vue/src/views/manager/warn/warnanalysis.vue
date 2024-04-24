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
            <div id="line" style="width: 1000px;height:460px;margin-top: 10px;">
            </div>
            <div id="line_bottom" style="width: 1000px;height:460px;margin-top: 10px;">
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
const sensor = ref('JC01')
const sensors = [
  {
    value: 'JC01',
    label: 'JC01',
  },
  {
    value: 'JC03',
    label: 'JC03'
  },
  {
    value: 'JC05',
    label: 'JC05'
  },
  {
    value: 'JC06',
    label: 'JC06'
  },
  {
    value: 'JC07',
    label: 'JC07'
  },
  {
    value: 'JC08',
    label: 'JC08'
  },
  {
    value: 'JC09',
    label: 'JC09'
  },
  {
    value: 'JC10',
    label: 'JC10'
  },
  {
    value: 'JC11',
    label: 'JC11'
  },
  {
    value: 'JC12',
    label: 'JC12'
  },
  {
    value: 'JC13',
    label: 'JC13'
  },
  {
    value: 'JC15',
    label: 'JC15'
  },
  {
    value: 'JC16',
    label: 'JC16'
  },
]
const data = reactive({
  tableData: ref([])
})
const handleSelect = () => {
  fetchData(sensor.value)
  ElMessage({
    message: sensor.value,
    type: 'success',
  })
}
onMounted(() => {
  fetchData(sensor.value);
});
// 从后端获取数据
const fetchData = async (sensorValue) => {
  try {
    const response = await request.get('http://124.70.83.36:3000/api/v1/warn/get', {
      headers: {
        Authorization: `Bearer ${token}` // 身份令牌
      },
    });
    data.tableData = response.data;
    var filteredData1 = data.tableData.filter(item => item.sensor_name === sensorValue);
    var collectTimes1 = filteredData1.map(item => item.today);
    var sumxy = 0, sumh = 0;
    var sumXyData1 = filteredData1.map(item => {
      sumxy += item.today_combinexy;
      return sumxy;
    });
    var sumHData1 = filteredData1.map(item => {
      sumh += item.today_combineh;
      return sumh;
    });
    var angleXy = filteredData1.map(item => item.xya)
    var angleH = filteredData1.map(item => item.ha)
    const myChart = echarts.init(document.getElementById('line'));
    var option = {
      title: {
        text: sensorValue + "：xy方向(左)、h方向(右)累计位移-时间曲线图",
        left: 'center',
      },
      grid: [
        { left: '5%', right: '50%', top: '10%', bottom: '10%' },
        { left: '60%', right: '5%', top: '10%', bottom: '10%' },
      ],
      tooltip: {
        trigger: 'axis'
      },
      toolbox: {
        feature: {
          saveAsImage: {}
        }
      },
      xAxis: [
        {
          gridIndex: 0,
          type: 'category',
          boundaryGap: false,
          data: collectTimes1,
          name: '日期',
          nameLocation: 'center',
          nameTextStyle: {
            borderWidth: 2,
            padding: 15
          }
        },
        {
          gridIndex: 1,
          type: 'category',
          boundaryGap: false,
          data: collectTimes1,
          name: '日期',
          nameLocation: 'center',
          nameTextStyle: {
            borderWidth: 2,
            padding: 15
          }
        }
      ],
      yAxis: [
        {
          gridIndex: 0,
          type: 'value',
          name: 'xy平面累积位移（mm）',
          nameLocation: 'middle',
          nameTextStyle: {
            borderWidth: 2,
            padding: 20
          }
        }, {
          gridIndex: 1,
          type: 'value',
          name: 'h平面累积位移（mm）',
          nameLocation: 'middle',
          nameTextStyle: {
            borderWidth: 2,
            padding: 20
          }
        }
      ],
      series: [
        {
          name: 'xy平面累积位移',
          type: 'line',
          smooth: true,
          stack: '总量',
          data: sumXyData1,
          xAxisIndex: 0,
          yAxisIndex: 0,
          markPoint: {
            data: [{
              type: 'max',
              name: '最大值'
            }, {
              type: 'min',
              name: '最小值'
            }]
          },
          markLine: {
            data: [{
              type: 'average',
              name: '平均值'
            }]
          }
        },
        {
          name: 'h平面累积位移',
          type: 'line',
          smooth: true,
          stack: '总量',
          data: sumHData1,
          xAxisIndex: 1,
          yAxisIndex: 1,
          markPoint: {
            data: [{
              type: 'max',
              name: '最大值'
            }, {
              type: 'min',
              name: '最小值'
            }]
          },
          markLine: {
            data: [{
              type: 'average',
              name: '平均值'
            }]
          }
        },
      ]
    };
    myChart.setOption(option);
    const myChart1 = echarts.init(document.getElementById('line_bottom'))
    var option1 = {
      title: {
        text: sensorValue + "：xy方向(左)、h方向(右)变形切线角-时间柱状图",
        left: 'center',
      },
      grid: [
        { left: '5%', right: '55%', top: '10%', bottom: '10%' },
        { left: '50%', right: '5%', top: '10%', bottom: '10%' },
      ],
      toolbox: {
        feature: {
          saveAsImage: {}
        }
      },
      xAxis: [
        {
          gridIndex: 0,
          type: 'category',
          data: collectTimes1,
          name: '日期',
          nameLocation: 'center',
          nameTextStyle: {
            borderWidth: 2,
            padding: 15
          }
        },
        {
          gridIndex: 1,
          type: 'category',
          data: collectTimes1,
          name: '日期',
          nameLocation: 'center',
          nameTextStyle: {
            borderWidth: 2,
            padding: 15
          }
        },
      ],
      yAxis: [
        {
          gridIndex: 0,
          type: 'value',
          name: 'xy平面变形切线角',
          nameLocation: 'middle',
          nameTextStyle: {
            borderWidth: 2,
            padding: 20
          }
        },
        {
          gridIndex: 1,
          type: 'value',
          name: 'h平面变形切线角',
          nameLocation: 'middle',
          nameTextStyle: {
            borderWidth: 2,
            padding: 20
          }
        },
      ],
      series: [
        {
          name: 'xy平面变形切线角',
          type: 'bar',
          barWidth: '40%',
          barGap: '50%',
          data: angleXy,
          xAxisIndex: 0,
          yAxisIndex: 0,
          label: {
            show: true, // 显示标签
            position: 'top', // 标签位置
            formatter: function (params) {
              return params.value.toFixed(2); // 标签内容，保留小数点后两位数
            }
          }
        },
        {
          name: 'h平面变形切线角',
          type: 'bar',
          data: angleH,
          xAxisIndex: 1,
          yAxisIndex: 1,
          label: {
            show: true, // 显示标签
            position: 'top', // 标签位置
            formatter: function (params) {
              return params.value.toFixed(2); // 标签内容，保留小数点后两位数
            }
          }
        },
      ]
    };
    myChart1.setOption(option1);
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