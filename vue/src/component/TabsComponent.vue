<template>
  <div>
    <a-tabs :active-key="activeTabKey" @change="handleTabChange">
      <a-tab-pane key="/course" tab="Home">
        <router-view></router-view>
      </a-tab-pane>
      <a-tab-pane key="/person" tab="About">
        <router-view></router-view>
      </a-tab-pane>
    </a-tabs>
  </div>
</template>

<script>
import { Tabs } from 'ant-design-vue';

export default {
  components: {
    'a-tabs': Tabs,
    'a-tab-pane': Tabs.TabPane,
  },
  data() {
    return {
      activeTabKey: this.$route.path, // 初始激活的标签页根据当前路由确定
    };
  },
  watch: {
    $route(to, from) {
      this.activeTabKey = to.path;
    },
  },
  methods: {
    handleTabChange(key) {
      this.$router.push(key); // 切换标签页时，通过路由导航切换页面
    },
  },
};
</script>
