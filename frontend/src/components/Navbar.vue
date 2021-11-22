<template>
  <el-header>
    <span>User</span>
    <el-dropdown>
      <i class="el-icon-setting"></i>
      <template #dropdown>
        <el-dropdown-menu>
          <el-dropdown-item>Settings</el-dropdown-item>
          <el-dropdown-item></el-dropdown-item>
          <el-dropdown-item @click="logout">Logout</el-dropdown-item>
        </el-dropdown-menu>
      </template>
    </el-dropdown>
  </el-header>
</template>
<script>
import { localstorage } from "../services/storage/localStorageService";
export default({
  methods: {
    logout() {
      localStorage.clear("token_user");
      localstorage.clearAuth();
      var aValue = localStorage.getItem("token_user");
      this.$store.commit("setAuthenticated", aValue);
      this.$router.push("/login");
    }
  },
})
</script>


<style scoped>
.el-header {
  text-align: right;
}
.el-header span {
  font-size: 18px;
  margin-right: 15px;
}
</style>
