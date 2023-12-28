<template>
  <div id="main" style="width: 500px;height:400px;"></div>
</template>

<script>
export default {
  name: "HelloWord",
  methods: {
    drawChart() {
      //2. 基于准备好的dom，初始化echarts实例
      var myChart = this.$echarts.init(document.getElementById('main'));
      //3. 指定图表的配置项和数据
      var option = {
        title: {
          text: 'ECharts 入门示例'
        },
        tooltip: {},
        legend: {    //图例组件
          data:['销量']
        },
        xAxis: {
          data: ["衬衫","羊毛衫","雪纺衫","裤子","高跟鞋","袜子"]
        },
        yAxis: {},
        series: [{
          name: '销量',
          type: 'bar',
          data: [5, 20, 36, 10, 10, 20]
        }]
      };
      //4.使用刚指定的配置项和数据显示图表。
      myChart.setOption(option);

    }
  },
  mounted(){
    this.drawChart();
  }
}
</script>

<style scoped>

</style>
