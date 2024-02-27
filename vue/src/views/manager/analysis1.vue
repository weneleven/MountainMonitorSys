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
              <el-menu-item >
                <el-icon><WarnTriangleFilled /></el-icon>
                <span>预测预警</span>
              </el-menu-item>
              <el-menu-item index="/upload">
                <el-icon><UploadFilled /></el-icon>
                <span>模拟数据发送</span>
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
                <el-select v-model="value1"  placeholder="Select" size="large">
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
            <div id="LC01" style="width: 1000px;height:400px;">
            </div>
            <div id="LC02" style="width: 1000px;height:400px;">
            </div>
          </div>
        </el-main>

      </el-container>
    </el-container>
  </div>
</template>
<script setup>
import { reactive, ref, onMounted, computed } from "vue";
import * as echarts from 'echarts';
// 项目切换数据
const value1 = ref('山西临汾古贤皮带机项目')
const options = [
  {
    value:'山西临汾古贤皮带机项目',
    label:'山西临汾古贤皮带机项目',
  },
]
onMounted(() => {
  const myChart1 = echarts.init(document.getElementById('LC01'));
  var option;
  function randomData() {
    now = new Date(+now + oneDay);
    value = value + Math.random() - 0.5;
    return {
      name: now.toString(),
      value: [
        [now.getFullYear(), now.getMonth() + 1, now.getDate()].join('/') + ' ' +
        [now.getHours(), now.getMinutes(), now.getSeconds()].join(':'),
        value.toFixed(1)
      ]
    };
  }
  let data = [];
  let now = new Date(2023, 12, 13);
  let oneDay = 10 * 1000;
  let value = Math.random();
  for (var i = 0; i < 1000; i++) {
    data.push(randomData());
  }
  option = {
    title: {
      text: 'LC01'
    },
    tooltip: {
      trigger: 'axis',
      formatter: function (params) {
        params = params[0];
        var date = new Date(params.name);
        return (
            date.getDate() +
            '/' +
            (date.getMonth() + 1) +
            '/' +
            date.getFullYear() +
            ' : ' +
            params.value[1]
        );
      },
      axisPointer: {
        animation: false
      }
    },
    xAxis: {
      type: 'time',
      splitLine: {
        show: false
      }
    },
    yAxis: {
      type: 'value',
      boundaryGap: [0, '100%'],
      splitLine: {
        show: false
      }
    },
    series: [
      {
        name: 'Fake Data',
        type: 'line',
        showSymbol: false,
        data: data
      }
    ]
  };
  setInterval(function () {
    for (var i = 0; i < 5; i++) {
      data.shift();
      data.push(randomData());
    }
    myChart1.setOption({
      series: [
        {
          data: data
        }
      ]
    });
  }, 1000);

  option && myChart1.setOption(option);
});

</script>