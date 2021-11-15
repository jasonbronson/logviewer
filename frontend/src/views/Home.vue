<template>
  <el-form ref="form" :model="form">
    <el-form-item>
      <el-button type="primary" @click="reloadLogs">Reload Log Files</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
import api from "@/api";
import { mapMutations } from "vuex";
export default {
  name: "Home",
  components: {},
  data() {
    return {};
  },
  computed: {},
  methods: {
    ...mapMutations(["setLogs"]),
    reloadLogs() {
      api.logs.getLogs().then((response) => {
        let success = false;
        if (response.status === 200) {
          this.setLogs(response.data.logs);
          success = true;
          this.$notify({
            title: "Success",
            message: "Getting logs succeeded",
          });
        }
        if (!success) {
          this.$notify.error({
            title: "Error",
            message: "Getting logs failed",
          });
        }
      });
    },
    // checkDatabase() {
    //   this.axios.get("/databaseinfo").then((response) => {
    //     if (response.status === 200) {
    //       console.log("db connection info ", response.data);
    //       if (response.data.db_name) {
    //         this.dblocation = response.data.db_name;
    //         this.status = " Connected";
    //         this.tableCount = response.data.table_count;
    //       }
    //     }
    //   });
    // },
  },
  mounted() {
    this.dblocation = this.$store.getters.getDatabase;
  },
};
</script>
<style scoped>
.dblocation {
}
</style>
