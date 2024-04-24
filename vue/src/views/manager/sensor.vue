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
            <el-form
                :label-position="left"
                label-width="120px"
                style="max-width: 560px"
            >
              <el-form-item label="切换项目">
              <el-select v-model="value"  placeholder="Select" size="large">
                <el-option
                    v-for="item in options"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value"
                />
              </el-select>
              </el-form-item>
            </el-form>
<!--传感器信息展示-->
            <div class="card" style="margin-bottom: 10px">

              <el-table :data="pagedTableData" stripe>
                <el-table-column label="设备SN" prop="sensor_sn"></el-table-column>
                <el-table-column label="设备名称" prop="sensor_name"></el-table-column>
                <el-table-column label="经度" prop="longitude"></el-table-column>
                <el-table-column label="纬度" prop="latitude"></el-table-column>
                <el-table-column label="地址" prop="address"></el-table-column>
                <el-table-column label="预警启用" align="center">
                  <template #default="scope">
                    <el-icon-check v-if="scope.row.AlertsEnabled" class="el-icon-check"></el-icon-check>
                    <el-icon-close v-else class="el-icon-close"></el-icon-close>
                  </template>
                </el-table-column>
                <el-table-column label="自动推送预警" align="center">
                  <template #default="scope">
                    <el-icon-check v-if="scope.row.AutoAlert" class="el-icon-check"></el-icon-check>
                    <el-icon-close v-else class="el-icon-close"></el-icon-close>
                  </template>
                </el-table-column>
                <el-table-column label="离线推送" align="center">
                  <template #default="scope">
                    <el-icon-check v-if="scope.row.OfflinePushEnabled" class="el-icon-check"></el-icon-check>
                    <el-icon-close v-else class="el-icon-close"></el-icon-close>
                  </template>
                </el-table-column>
                <el-table-column label="采集间隔" prop="acquisition_interval"></el-table-column>
                <el-table-column label="归档状态" prop="archive_status"></el-table-column>
                <el-table-column label="是否报警" prop="is_warning"></el-table-column>
                <el-table-column label="报警间隔" prop="warning_interval"></el-table-column>
                <el-table-column label="备注" prop="commit"></el-table-column>
                <el-table-column label="操作" align="center" width="160">
                  <template v-slot="scope">
                    <el-button type="primary" @click="handleEdit">编辑</el-button>
                  </template>
                </el-table-column>
              </el-table>
            </div>

            <div class="card">
              <el-pagination
                  background
                  layout="prev, pager, next"
                  :page-size="pageSize"
                  :current-page.sync="currentPage"
                  :total="totalItems"
                  @current-change="handlePageChange"
              />
            </div>

            <el-dialog title="修改数据" width="40%" v-model="data.formVisible" :close-on-click-modal="false" destroy-on-close>
              <el-form ref="formRef" :model="data.form" label-width="150px">
                <el-form-item label="项目ID" prop="ID"  required>
                  <el-input v-model="data.form.ID" />
                </el-form-item>

                <el-form-item label="项目名称" prop="ProjectName"  >
                  <el-input v-model="data.form.ProjectName" />
                </el-form-item>

                <el-form-item label="项目简称" prop="ProjectShortName">
                  <el-input v-model="data.form.ProjectShortName" type="textarea" />
                </el-form-item>

                <el-form-item label="项目类型" prop="projectType">
                  <el-radio-group v-model="data.form.ProjectType">
                    <el-radio label="边坡" />
                    <el-radio label="水库" />
                    <el-radio label="桥梁" />
                  </el-radio-group>
                </el-form-item>

                <el-form-item label="项目位置" prop="Location">
                  <el-input v-model="data.form.Location" type="textarea" />
                </el-form-item>

                <el-form-item label="自动推送预警" prop="AutoAlert">
                  <el-switch v-model="data.form.AutoAlert" />
                </el-form-item>

                <el-form-item label="预警启用" prop="AlertsEnabled">
                  <el-switch v-model="data.form.AlertsEnabled" />
                </el-form-item>

                <el-form-item label="离线推送" prop="OfflinePushEnabled">
                  <el-switch v-model="data.form.OfflinePushEnabled" />
                </el-form-item>


                <el-form-item label="离线推送频率(s)" prop="offlineFrequency">
                  <el-input-number v-model="data.form.OfflinePushFrequency" :min="0" :max="86400" />
                </el-form-item>


                <el-form-item label="地图比例" prop="mapScale">
                  <el-input-number v-model="data.form.MapScale" :min="0" />
                </el-form-item>

                <el-form-item label="短信签名" prop="smsSignature">
                  <el-input v-model="data.form.SmsSignature" />
                </el-form-item>

                <el-form-item label="项目简介" prop="ProjectDescription">
                  <el-input v-model="data.form.ProjectDescription" type="textarea" />
                </el-form-item>

                <el-form-item label="备注" prop="Remark">
                  <el-input v-model="data.form.Remark" type="textarea" />
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
import request from "@/utils/request";
import { reactive, ref, onMounted, computed } from "vue";
import { ElMessageBox } from "element-plus";
import { useRoute } from 'vue-router';

const $route = useRoute()
console.log($route.path)


const username = JSON.parse(localStorage.getItem('user') || '{}')
const token = JSON.parse(localStorage.getItem('token')||'{}')

const logout = () => {
  window.location.href = '/'; // 重定向到登录页面
}
// 项目切换数据
const value = ref('山西临汾古贤皮带机项目')
const options = [
  {
    value:'山西临汾古贤皮带机项目',
    label:'山西临汾古贤皮带机项目',
  },
]
// 示例数据
const totalItems = ref(100); // 可根据后端接口修改，使总数据数由后端提供
const currentPage = ref(1);
const pageSize = ref(2);
const searchQuery = ref('');


const data = reactive({
  pageNum: 1,
  formVisible: false,
  form: {},
  //测试数据
  tableData: ref([])
})
onMounted(() => {
  fetchData(); // 在组件挂载时调用fetchData函数获取数据
});
// 从后端获取数据
const fetchData = async () => {
  try {
    const response = await request.get('http://124.70.83.36:3000/api/v1/projects', {
      headers: {
        Authorization: `Bearer ${token}` // 身份令牌
      },
    });
    //  data.tableData = response.data;
    // 映射后端返回的字段名到前端表单项的字段名
    const mappedData = response.data.map(item => ({
      ID:item.ID,
      ProjectName: item.project_name,
      ProjectShortName:item.project_short_name,
      ProjectType:item.project_type,
      MapScale:item.map_scale,
      Location:item.location,
      OfflinePushFrequency:item.offline_push_frequency,
      InspectionUnitName:item.inspection_unit_name,
      // 其他映射项...
    }));
    data.tableData = mappedData;
  } catch (error) {
    console.error('Error fetching data:', error);
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
        ID:response.data.ID,
        ProjectName: response.data.project_name,
        ProjectShortName: response.data.project_short_name,
        ProjectType: response.data.project_type,
        MapScale: response.data.map_scale,
        Location: response.data.location,
        OfflinePushFrequency: response.data.offline_push_frequency,
        InspectionUnitName: response.data.inspection_unit_name,
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
