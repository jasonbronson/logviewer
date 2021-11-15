<template>
  <el-aside class="sidebar">
    <p>Log Reader</p>
    <el-menu
      :default-active="activeMenuItem"
      class="el-menu"
      @open="handleOpen"
      @close="handleClose"
    >
      <router-link to="/">
        <el-menu-item index="1">
          <i class="el-icon-menu"></i>
          <span>Main</span>
        </el-menu-item>
      </router-link>
      <!-- <router-link to="/query">
        <el-menu-item index="2" :disabled="activateMenu">
          <i class="el-icon-menu"></i>
          <span>Query</span>
        </el-menu-item>
      </router-link>
      <router-link to="/create">
        <el-menu-item index="3" :disabled="activateMenu">
          <i class="el-icon-menu"></i>
          <span>Create</span>
        </el-menu-item>
      </router-link> -->
      <el-submenu index="4" :disabled="activateMenu">
        <template #title><i class="el-icon-menu"></i>Logs</template>
        <el-menu-item
          :index="index + 1"
          v-for="(log, index) in listLogs"
          @click="showTable(log.Name)"
        >
          <i class="el-icon-caret-right"></i>
          <span>{{ log.Name }}</span>
        </el-menu-item>
      </el-submenu>
    </el-menu>
  </el-aside>
</template>

<script>
export default {
  name: "HomeView",
  props: {
    menu: {
      required: true,
    },
  },
  data() {
    return {
      activeMenuItem: "1",
    };
  },
  computed: {
    listLogs() {
      let logs = this.$store.state.logs;
      if (logs) {
        logs.forEach((log, index) => {
          if (log.Name == this.$store.state.selectedTable) {
            this.activeMenuItem = index + 1;
          }
        });
      }
      return logs;
    },
  },
  methods: {
    showLog(log) {
      console.log("Select Log Name:", log);
      this.$store.commit("setSelectedLog", log);
      this.$router.push(`/log/${log}`);
    },
  },
};
</script>
<style scoped>
.sidebar {
}
</style>
