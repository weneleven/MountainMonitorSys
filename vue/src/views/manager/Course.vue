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
            <!-- 搜索部分 -->
            <div class="card" style="margin-bottom: 10px; display: flex; justify-content: space-between; align-items: center;">
              <div>
                <el-input v-model="searchQuery" style="width: 300px; margin-right: 10px;" placeholder="请输入项目名称"></el-input>
                <el-button type="primary" @click="search">查询</el-button>
                <el-button type="info" @click="resetSearch">重置</el-button>
              </div>
              <div>
                <el-button type="primary" icon="plus" @click="addproject">新增项目</el-button>
              </div>
            </div>
            <div class="card" style="margin-bottom: 10px">

              <el-table :data="pagedTableData" stripe>
                <el-table-column label="项目ID" prop="ID"></el-table-column>
                <el-table-column label="项目名称" prop="ProjectName"></el-table-column>
                <el-table-column label="项目简称" prop="ProjectShortName"></el-table-column>
                <el-table-column label="项目类型" prop="ProjectType"></el-table-column>
                <el-table-column label="项目位置" prop="Location"></el-table-column>
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
                <el-table-column label="离线推送频率(s)" prop="OfflinePushFrequency"></el-table-column>
                <el-table-column label="地图比例" prop="MapScale"></el-table-column>
                <el-table-column label="短信签名" prop="SmsSignature"></el-table-column>
                <el-table-column label="备注信息" prop="Remark"></el-table-column>
                <el-table-column label="项目简介" prop="ProjectDescription"></el-table-column>
                <el-table-column label="观察者" prop="ObserverName"></el-table-column>
                <el-table-column label="校验者" prop="ValidatorName"></el-table-column>
                <el-table-column label="计算者" prop="CalculatorName"></el-table-column>
                <el-table-column label="审核者" prop="ReviewerName"></el-table-column>
                <el-table-column label="检测单位名称" prop="InspectionUnitName"></el-table-column>
                <el-table-column label="用户ID" prop="UserID"></el-table-column>
                <el-table-column label="操作"  min-width="160"
                                 fixed="right">
                  <template #default="scope">
                    <!--            删除操作-->
                    <el-popover
                        v-model:visible="scope.row.visible"
                        placement="top"
                        width="160"
                    >
                      <p>确定要删除此项目吗</p>
                      <div style="text-align: right; margin-top: 8px;">
                        <el-button
                            type="primary"
                            link
                            @click="scope.row.visible = false"
                        >取消</el-button>
                        <el-button
                            type="primary"
                            @click="deleteProjectFunc(scope.row)"
                        >确定</el-button>
                      </div>
                      <template #reference>
                        <el-button
                            type="primary"
                            link
                            icon="delete"
                        >删除</el-button>
                      </template>
                    </el-popover>
                    <!--           编辑操作 -->
                    <el-button
                        type="primary"
                        link
                        icon="edit"
                        @click="openEdit(scope.row)"
                    >编辑</el-button>
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
            <!--新增项目选项卡-->
            <el-dialog
                v-model="addprojectDialog"
                title="项目"
                :show-close="false"
                :close-on-press-escape="false"
                :close-on-click-modal="false"
            >
              <div style="height:60vh;overflow:auto;padding:0 12px;">
                <el-form
                    ref="projectForm"
                    :rules="rules"
                    :model="projectInfo"
                    label-width="80px"
                >
                  <el-form-item
                      label="项目名称"
                      prop="project_name"
                  >
                    <el-input v-model="projectInfo.project_name" />
                  </el-form-item>
                  <el-form-item
                      label="项目简称"
                      prop="project_short_name"
                  >
                    <el-input v-model="projectInfo.project_short_name" />
                  </el-form-item>
                  <el-form-item
                      label="项目类型"
                      prop="project_type"
                  >
                    <el-input v-model="projectInfo.project_type" />
                  </el-form-item>
                  <el-form-item
                      label="项目位置"
                      prop="location"
                  >
                    <el-input v-model="projectInfo.location" />
                  </el-form-item>
                  <el-form-item
                      label="所属单位"
                      prop="inspection_unit_name"
                  >
                    <el-input v-model="projectInfo.inspection_unit_name" />
                  </el-form-item>
                </el-form>
              </div>

              <template #footer>
                <div class="dialog-footer">
                  <el-button @click="closeAddprojectDialog">取 消</el-button>
                  <el-button
                      type="primary"
                      @click="enterAddProjectDialog()"
                  >确 定</el-button>
                </div>
              </template>
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
import {ElMessage, ElMessageBox} from "element-plus";
import { useRoute } from 'vue-router';
import axios from "axios";

const $route = useRoute()
console.log($route.path)


const username = JSON.parse(localStorage.getItem('user') || '{}')
const token = JSON.parse(localStorage.getItem('token')||'{}')

const logout = () => {
  window.location.href = '/'; // 重定向到登录页面
}
//新增相关
const projectInfo = ref({
  project_name: '',
  project_short_name: '',
  project_type: '',
  location: '',
  inspection_unit_name:'',
})
const id = ref()
const rules = ref({
  project_name: [
    { required: true, message: '请输入项目名称', trigger: 'blur' },
    { min: 1, message: '最低1位字符', trigger: 'blur' }
  ],
  project_type: [
    { required: true, message: '请输入项目类型', trigger: 'blur' },
    { min: 1, message: '最低1位字符', trigger: 'blur' }
  ],
  location: [
    { required: true, message: '请输入项目位置', trigger: 'blur' }
  ]
})
const projectForm = ref(null)
const addprojectDialog = ref(false)
const enterAddProjectDialog = async() => {
  if(dialogFlag.value === 'add'){
  try {
    const response = await axios.post('http://124.70.83.36:3000/api/v1/project/add', projectInfo.value, {
      headers: {
        Authorization: `Bearer ${token}`
      },
    });
    console.log('项目添加成功:', response.data);
    ElMessage({ type: 'success', message: '创建成功' })
    closeAddprojectDialog()
    fetchData()
  } catch (error) {
    console.error('项目添加失败:', error);
  }}
  if(dialogFlag.value === 'edit'){
    try {
      const response = await axios.put(`http://124.70.83.36:3000/api/v1/project/updateProject?id=${id.value}`, projectInfo.value, {
        headers: {
          Authorization: `Bearer ${token}`
        },
      });
      if (response.status === 200) {
        console.log('项目更新成功:', response.data);
        ElMessage({ type: 'success', message: '更新成功' })
        closeAddprojectDialog()
        fetchData()
      } else {
        console.error('项目更新失败:', response.data);
      }
    } catch (error) {
      console.error('更新请求出错:', error);
    }
  }
}
const closeAddprojectDialog = () => {
  projectForm.value.resetFields()
  addprojectDialog.value = false
}
const dialogFlag = ref('add')
const addproject = () => {
  dialogFlag.value = 'add'
  addprojectDialog.value = true
}
const deleteProjectFunc = async (row) => {
  try {
    const response = await axios.delete(`http://124.70.83.36:3000/api/v1/project/deleteById?id=${row.ID}`, {
      headers: {
        Authorization: `Bearer ${token}`
      },
    });
    if (response.status === 200) {
      console.log('项目删除成功:', response.data);
      ElMessage({ type: 'success', message: '删除成功' });
      fetchData()
    } else {
      console.error('项目删除失败:', response.data);
      ElMessage({ type: 'error', message: '删除失败' });
    }
  } catch (error) {
    console.error('删除请求出错:', error);
    ElMessage({ type: 'error', message: '删除请求出错' });
  }
};
const openEdit = (row) => {
  dialogFlag.value = 'edit'
  projectInfo.value = JSON.parse(JSON.stringify(row))
  addprojectDialog.value = true
  id.value = row.ID
}
// 示例数据
const totalItems = ref(100); // 可根据后端接口修改，使总数据数由后端提供
const currentPage = ref(1);
const pageSize = ref(5);
const searchQuery = ref('');
const data = reactive({
  pageNum: 1,
  formVisible: false,
  form: {},
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
