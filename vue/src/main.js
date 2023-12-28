import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import ElementPlus from 'element-plus'
import zhCn from 'element-plus/dist/locale/zh-cn.mjs'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import * as echarts from 'echarts';

import '@/assets/css/global.css'

const app = createApp(App)

app.use(router)
app.use(ElementPlus, {
    locale: zhCn,
})
// 将ECharts挂载到Vue的原型上
app.config.globalProperties.$echarts = echarts;
app.mount('#app')

//使用element重要配置
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}