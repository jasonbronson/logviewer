<template>
  <el-tabs type="card" @tab-click="handleClick">
    <el-tab-pane label="Structure"
      ><structure :columns="columns" :sql="sql"
    /></el-tab-pane>
    <el-tab-pane label="Content"
      ><content :tablecontent="tablecontent"
    /></el-tab-pane>
    <el-tab-pane label="Import"><import /></el-tab-pane>
    <el-tab-pane label="Drop"><drop /></el-tab-pane>
  </el-tabs>
</template>

<script>
import { ref } from "vue";
import Content from "@/components/Content";
import api from "@/api";

export default {
  name: "Table",
  components: {
    Content,
  },
  data() {
    return {
      tabledata: [],
      sql: "",
      columns: [],
      tablecontent: [],
      pageOffset: 0,
      pageLimit: 10,
    };
  },
  computed: {
    getSelectedTable() {
      return this.$store.state.selectedTable;
    },
  },
  methods: {
    handleClick(tab, event) {
      console.log(event.target.outerText);
      if (event.target.outerText === "Content") {
        this.loadContent();
      }
    },
    loadContent() {
      api.table
        .getTableContent(this.getSelectedTable, this.pageOffset, this.pageLimit)
        .then((res) => {
          console.log(res);
          this.tablecontent = res.data;
        })
        .catch((err) => {
          console.log(err);
        });
    },
    changeTable() {
      let tables = this.$store.state.tables;
      tables.forEach((item) => {
        if (item.Name == this.$store.state.selectedTable) {
          this.sql = item.SQL;
        }
      });
      api.structure
        .getColumns(this.$store.state.selectedTable)
        .then((response) => {
          if (response.status == 200) {
            this.columns = response.data;
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
    // createdatabase() {
    //   this.axios
    //     .post("/table/create?table=" + this.createtablename)
    //     .then((response) => {
    //       console.log(response);
    //       if (response.data === "Success") {
    //         this.createtablename = "";
    //         this.loadData();
    //         //throw success msg
    //       } else {
    //         //throw error msg
    //       }
    //     });
    // },
  },
  watch: {
    getSelectedTable(newValue, oldValue) {
      console.log(oldValue, newValue);
      this.changeTable();
    },
  },
  mounted() {
    this.changeTable();
  },
};
</script>
<style scoped></style>
