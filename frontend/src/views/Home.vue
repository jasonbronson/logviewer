<template>
  <el-form ref="form" :model="form">
    <el-form-item>
      <el-button type="primary" @click="reloadLogs"
        >{{ btntext }} Log Files</el-button
      >
    </el-form-item>
  </el-form>
</template>

<script>
import api from "@/api";
import { mapMutations, mapState } from "vuex";
export default {
  name: "Home",
  components: {},
  data() {
    return {};
  },
  computed: {
    ...mapState({
      logs: (state) => state.log.logs,
    }),
    btntext() {
      return this.logs.length > 0 ? "Refresh" : "Load";
    },
  },
  methods: {
    ...mapMutations(["setLogs"]),
    async reloadLogs() {
      let payload = {
        logName: "",
        pageLimit: "",
        pageOffset: "",
      }
      try {
        var response = await this.$store.dispatch('getLogs', payload)
        this.setLogs(response.logs);
        this.$notify({
          title: "Success",
          message: "Getting logs succeeded",
        });
        if (response.logs[0].Name) {
          let logName = response.logs[0].Name;
          console.log("Selected Log:", logName);
          this.$store.commit("setSelectedLog", logName);
          this.$router.push(`/log/${logName}`);
        }
      } catch(err) {
        this.$notify.error({
          title: "Error",
          message: "Getting logs failed",
        });
      }
      
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
