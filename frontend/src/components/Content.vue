<template>
  <div>
    <p>Table Content</p>
    <el-table :data="tablecontent" stripe style="width: 100%">
      <el-table-column prop="created_at" label="Date" width="180"> </el-table-column>
      <el-table-column prop="id" label="Name" width="180"> </el-table-column>
      <el-table-column prop="address" label="Address"> </el-table-column>
    </el-table>
  </div>
</template>
<script>
import Vue from "vue";

export default {
  name: "Content",
  props: {
    tablecontent: Object,
  },
  mounted(){
    this.$store.commit('setLogger', JSON.stringify(this.tablecontent))
  }
};
</script>