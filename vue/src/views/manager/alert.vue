<template>
    <div class="common-layout">
        <el-container>
            <el-header>
                <div
                    style="height: 60px; background-color:#ffffff; display: flex; align-items: center; border-bottom: 1px solid #ddd">
                    <div style="flex: 1">
                        <div style="padding-left: 20px; display: flex; align-items: center">
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
                <el-aside width="200px">
                    <div style="width: 200px; border-right: 1px solid #ddd; min-height: calc(100vh - 60px)">
                        <el-menu router style="border: none" :default-active="$route.path"
                            :default-openeds="['/home', '2']">
                            <el-menu-item index="/home">
                                <el-icon>
                                    <HomeFilled />
                                </el-icon>
                                <span>系统首页</span>
                            </el-menu-item>
                            <el-sub-menu index="2">
                                <template #title>
                                    <el-icon>
                                        <MapLocation />
                                    </el-icon>
                                    <span>数据视图</span>
                                </template>
                                <el-menu-item index="/manager">
                                    <span>项目地图</span>
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
                            <el-menu-item index="/person">
                                <el-icon>
                                    <DataLine />
                                </el-icon>
                                <span>数据分析</span>
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
                        <div class="card" style="margin-bottom: 10px">
                            <el-form :label-position="left" label-width="120px" style="max-width: 560px">
                                <el-form-item label="切换项目" style="margin-bottom: 8px;">
                                    <el-select v-model="value" placeholder="Select" size="large">
                                        <el-option v-for="item in options" :key="item.value" :label="item.label"
                                            :value="item.value" />
                                    </el-select>
                                </el-form-item>
                            </el-form>
                        </div>

                        <div class="card" style="display: flex; align-items: center; margin-bottom: 10px;">

                            <!-- 左侧容器，包含 demo-progress -->
                            <div class="left-container" style="margin-right: 10px; margin-left: 10px;">
                                <div v-if="hasAlarms" class="demo-progress" style="margin-left: 50px; margin-right: 50px;">
                                    <el-progress type="circle" :percentage="alarmPercentage"
                                        :status="alarmStatus"></el-progress>
                                </div>
                                <div v-else style="margin-left: 50px; margin-right: 50px;">
                                    <el-progress type="circle" :percentage="100" status="success" />
                                </div>

                                <div v-if="hasAlarms" style="margin-top: 10px; margin-left: 10px;">
                                    <el-alert :title="`当前有${tableData.length}条待处理预警！`" type="warning" :closable="false"
                                        show-icon />
                                </div>
                                <div v-else style="margin-top: 10px; margin-left: 10px;">
                                    <el-alert title="当前无待处理预警信息" type="success" :closable="false" show-icon />
                                </div>
                            </div>



                            <!-- 右侧容器，包含折线图 -->
                            <div class="right-container" style="width: 1200px; margin-left: 30px;">
                                <div id="main" style="width: 100%; height: 420px;"></div>
                            </div>

                        </div>
                        <div class="card" style="margin-bottom: 10px">
                            <el-table :data="pagedTableData" stripe>
                                <el-table-column label="报警传感器" prop="SensorSN"></el-table-column>
                                <el-table-column label="报警名称" prop="AlertName"></el-table-column>
                                <el-table-column label="报警字段" prop="AlertParameter"></el-table-column>
                                <el-table-column label="X值阈值" prop="VXThreshold"></el-table-column>
                                <el-table-column label="Y值阈值" prop="VYThreshold"></el-table-column>
                                <el-table-column label="H值阈值" prop="VHThreshold"></el-table-column>
                                <el-table-column label="X方向累计阈值" prop="SumXThreshold"></el-table-column>
                                <el-table-column label="Y方向累计阈值" prop="SumYThreshold"></el-table-column>
                                <el-table-column label="H方向累计阈值" prop="SumHThreshold"></el-table-column>
                                <el-table-column label="报警次数" prop="Frequency"></el-table-column>
                                <el-table-column label="操作" align="center" width="160">
                                    <template v-slot="scope">
                                        <el-button type="primary" @click="handleEdit">编辑</el-button>
                                    </template>
                                </el-table-column>
                            </el-table>
                        </div>

                        <div class="card">
                            <el-pagination background layout="prev, pager, next" :page-size="pageSize"
                                :current-page.sync="currentPage" :total="totalItems" @current-change="handlePageChange" />
                        </div>

                        <el-dialog title="修改数据" width="40%" v-model="data.formVisible" :close-on-click-modal="false"
                            destroy-on-close>
                            <el-form ref="formRef" :model="data.form" label-width="150px">
                                <el-form-item label="报警传感器" prop="SensorSN" required>
                                    <el-input v-model="data.form.SensorSN" />
                                </el-form-item>

                                <el-form-item label="报警名称" prop="AlertName">
                                    <el-input v-model="data.form.AlertName" />
                                </el-form-item>

                                <el-form-item label="报警字段" prop="AlertParameter">
                                    <el-input v-model="data.form.AlertParameter" type="textarea" />
                                </el-form-item>

                                <el-form-item label="X值阈值" prop="VXThreshold">
                                    <el-input-number v-model="data.form.VXThreshold" :min="0" />
                                </el-form-item>

                                <el-form-item label="Y值阈值" prop="VYThreshold">
                                    <el-input-number v-model="data.form.VYThreshold" :min="0" />
                                </el-form-item>

                                <el-form-item label="H值阈值" prop="VHThreshold">
                                    <el-input-number v-model="data.form.VHThreshold" :min="0" />
                                </el-form-item>

                                <el-form-item label="X方向累计阈值" prop="SumXThreshold">
                                    <el-input-number v-model="data.form.SumXThreshold" :min="0" />
                                </el-form-item>

                                <el-form-item label="Y方向累计阈值" prop="SumYThreshold">
                                    <el-input-number v-model="data.form.SumYThreshold" :min="0" />
                                </el-form-item>

                                <el-form-item label="H方向累计阈值" prop="SumHThreshold">
                                    <el-input-number v-model="data.form.SumHThreshold" :min="0" />
                                </el-form-item>

                                <el-form-item label="报警次数" prop="Frequency">
                                    <el-input-number v-model="data.form.Frequency" :min="1" />
                                </el-form-item>

                                <!-- 操作按钮 -->
                                <el-form-item>
                                    <el-button type="primary" @click="submitEdit">保存更改</el-button>
                                    <el-button @click="resetEditForm">取消</el-button>
                                </el-form-item>
                            </el-form>
                        </el-dialog>
                    </div>
                </el-main>
            </el-container>
        </el-container>
    </div>
</template>
  
  
  
<script setup>
import { Check } from '@element-plus/icons-vue'
import request from "@/utils/request";
import { reactive, ref, onMounted, computed, toRefs } from "vue";
import { ElMessageBox } from "element-plus";
import { useRoute } from 'vue-router';
import * as echarts from 'echarts'


const $route = useRoute()
console.log($route.path)
//演示数据，实际可与表格绑定
const opinionData = ref(["500", "400", "900", "800", "300", "900", "200"]);
const charts = ref(null);

const drawLine = (id) => {
    charts.value = echarts.init(document.getElementById(id));
    charts.value.setOption({
        title: {
            left: '3%',
            top: '5%',
            text: "传感器报警情况",
        },
        tooltip: {
            trigger: 'axis',
            formatter: params => {
                const data = params[0].data;
                return `报警数量: ${data}`;
            },
        },
        legend: {
            align: 'right',
            left: '3%',
            top: '15%',
            data: ['近一周'],
        },
        grid: {
            top: '30%',
            left: '5%',
            right: '5%',
            bottom: '5%',
            containLabel: true,
        },
        toolbox: {
            feature: {
                saveAsImage: {},
            },
        },
        xAxis: {
            type: 'category',
            boundaryGap: true,
            axisTick: {
                alignWithLabel: true,
            },
            data: ["SN-01", "SN-02", "SN-03", "SN-04", "SN-05", "SN-06", "SN-07"],
        },
        yAxis: {
            type: 'value',
            boundaryGap: true,
            splitNumber: 4,
            interval: 250,
        },
        series: [{
            name: '近一周',
            type: 'line',
            stack: '总量',
            areaStyle: {
                color: {
                    type: 'linear',
                    x: 0,
                    y: 0,
                    x2: 0,
                    y2: 1,
                    colorStops: [{
                        offset: 0, color: 'rgb(255,200,213)',
                    }, {
                        offset: 1, color: '#ffffff',
                    }],
                    global: false,
                },
            },
            itemStyle: {
                color: 'rgb(255,96,64)',
                lineStyle: {
                    color: 'rgb(255,96,64)',
                },
            },
            data: opinionData.value,
        }],
    });
};

// 在组件挂载后初始化图表
onMounted(() => {
    drawLine('main');
});


const username = JSON.parse(localStorage.getItem('user') || '{}')
const token = JSON.parse(localStorage.getItem('token') || '{}')

const logout = () => {
    window.location.href = '/'; // 重定向到登录页面
}
// 项目切换数据
const value = ref('山西临汾古贤皮带机项目')
const options = [
    {
        value: '山西临汾古贤皮带机项目',
        label: '山西临汾古贤皮带机项目',
    },
]
// 示例数据
const totalItems = ref(2); // 可根据后端接口修改，使总数据数由后端提供
const currentPage = ref(1);
const pageSize = ref(2);
const searchQuery = ref('');
//测试数据，可删除
const testData1 = {
    ID: 1,
    SensorSN: 'SN123',
    AlertName: '测试报警1',
    AlertParameter: '参数1',
    VXThreshold: 10.5,
    VYThreshold: 20.2,
    VHThreshold: 30.0,
    SumXThreshold: 100.8,
    SumYThreshold: 200.3,
    SumHThreshold: 300.6,
    Frequency: 3,
};

const testData2 = {
    ID: 2,
    SensorSN: 'SN456',
    AlertName: '测试报警2',
    AlertParameter: '参数2',
    VXThreshold: 15.2,
    VYThreshold: 25.0,
    VHThreshold: 35.5,
    SumXThreshold: 120.1,
    SumYThreshold: 210.6,
    SumHThreshold: 310.4,
    Frequency: 5,
};

const data = reactive({
    pageNum: 1,
    formVisible: false,
    form: {},
    //测试数据
    tableData: [testData1, testData2],
});

const { tableData } = toRefs(data);

const hasAlarms = computed(() => {
    return tableData.value.length > 0;
});

const alarmPercentage = computed(() => {
    return 70;
});

const alarmStatus = computed(() => {
    return hasAlarms.value ? "warning" : "success";
});

onMounted(() => {
    fetchData(); // 在组件挂载时调用fetchData函数获取数据
});
// 从后端获取数据
const fetchData = async () => {
    try {
        const response = await request.get('http://api', {//获取报警列表api
            headers: {
                Authorization: `Bearer ${token}`, // 您的身份验证令牌
            },
        });

        const mappedData = response.data.map(item => ({
            ID: item.ID,
            SensorSN: item.sensor_sn,
            AlertName: item.alert_name,
            AlertParameter: item.alert_parameter,
            VXThreshold: item.vx_threshold,
            VYThreshold: item.vy_threshold,
            VHThreshold: item.vh_threshold,
            SumXThreshold: item.sum_x_threshold,
            SumYThreshold: item.sum_y_threshold,
            SumHThreshold: item.sum_h_threshold,
            Frequency: item.frequency,
        }));

        data.tableData = mappedData;
    } catch (error) {
        console.error('获取数据时出错:', error);
    }
};

// 根据名称获取项目
const search = async () => {
    try {
        const response = await request.get('http://124.70.83.36:3000/api/v1/project/getByName', {
            params: { name: searchQuery.value },
            headers: {
                Authorization: `Bearer ${token}` // 身份令牌
            },
        });
        //返回值检查
        if (response.data) {
            // 提取数据并赋值给 tableData
            data.tableData = [{
                ID: item.ID,
                SensorSN: item.sensor_sn,
                AlertName: item.alert_name,
                AlertParameter: item.alert_parameter,
                VXThreshold: item.vx_threshold,
                VYThreshold: item.vy_threshold,
                VHThreshold: item.vh_threshold,
                SumXThreshold: item.sum_x_threshold,
                SumYThreshold: item.sum_y_threshold,
                SumHThreshold: item.sum_h_threshold,
                Frequency: item.frequency,
            }];
        } else {
            data.tableData = ref([])
            console.log(data.tableData)
            console.error('Error searching: Response data is null or undefined.');
        }

    } catch (error) {
        console.error('Error searching:', error);
    }
};

// 重置搜索
const resetSearch = () => {
    searchQuery.value = '';
    fetchData(); // 重置后重新获取数据
};

// 计算当前页的数据
const pagedTableData = computed(() => {
    const start = (currentPage.value - 1) * pageSize.value;
    const end = start + pageSize.value;
    return data.tableData?.slice(start, end);
});

// 处理页码变化
const handlePageChange = (newPage) => {
    currentPage.value = newPage;
};


request.get('/').then(res => {
    console.log(res)
})

const handleEdit = (row) => {
    let form = JSON.parse(JSON.stringify(row))
    data.formVisible = true
}

const submitEdit = async () => {
    try {
        await request.put('/api/', data.form);//更新接口
        data.formVisible = false;
        // 重新获取更新后的数据列表
        fetchData();
    } catch (error) {
        console.error('Error updating project:', error);
    }
};

const resetEditForm = () => {
    data.form = {}; // 清空表单数据
    data.formVisible = false;
};




</script>
  
<style scoped>
.el-icon-check {
    color: green;
}

.el-icon-close {
    color: red;
}

.percentage-value {
    display: block;
    margin-top: 10px;
    font-size: 28px;
}

.percentage-label {
    display: block;
    margin-top: 10px;
    font-size: 12px;
}

.demo-progress .el-progress--line {
    margin-bottom: 15px;
    width: 350px;
}

.demo-progress .el-progress--circle {
    margin-right: 15px;
}

.el-alert {
    margin: 20px 0 0;
}

.el-alert:first-child {
    margin: 0;
}
</style>
  