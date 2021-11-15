<template>
  <el-header>
    <span>User</span>
    <el-dropdown>
      <i class="el-icon-setting"></i>
      <template #dropdown>
        <el-dropdown-menu>
          <el-dropdown-item>Settings</el-dropdown-item>
          <el-dropdown-item></el-dropdown-item>
          <el-dropdown-item>Logout</el-dropdown-item>
        </el-dropdown-menu>
      </template>
    </el-dropdown>
  </el-header>
</template>

<style scoped>
.el-header {
  text-align: right;
}
.el-header span {
  font-size: 18px;
  margin-right: 15px;
}
</style>
