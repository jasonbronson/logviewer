<template>
  <div class="content-page">
    <div class="content-body">
      <!-- <el-tabs type="card" @tab-click="handleClick">
        <el-tab-pane label="Structure"
          ><structure :columns="columns" :sql="sql"
        /></el-tab-pane>
        <el-tab-pane label="Content"
          ><content :tablecontent="tablecontent"
        /></el-tab-pane>
        <el-tab-pane label="Import"><import /></el-tab-pane>
        <el-tab-pane label="Drop"><drop /></el-tab-pane>
      </el-tabs> -->
      <div v-for="(content, idx) in logcontent" :key="idx" class="content-item">
        <span>{{ content }}</span>
      </div>
    </div>
    <div class="search">
      <input
        placeholder="Enter text for searching..."
        v-model="searchKey"
        class="search-input"
      />
      <button @click="searchHandle()" class="search-btn">Search</button>
    </div>
  </div>
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
      logcontent: [],
      pageOffset: 0,
      pageLimit: 20,
      reload: null,
      searchKey: "",
    };
  },
  computed: {
    getSelectedLog() {
      return this.$store.state.selectedLog;
    },
  },
  methods: {
    // handleClick(tab, event) {
    //   console.log(event.target.outerText);
    //   if (event.target.outerText === "Content") {
    //     this.loadContent();
    //   }
    // },
    // loadContent() {
    //   api.table
    //     .getTableContent(this.getSelectedTable, this.pageOffset, this.pageLimit)
    //     .then((res) => {
    //       console.log(res);
    //       this.tablecontent = res.data;
    //     })
    //     .catch((err) => {
    //       console.log(err);
    //     });
    // },
    searchHandle() {
      clearInterval(this.reload);
      api.logs
        .getSearchContent(this.getSelectedLog, this.searchKey, this.pageLimit)
        .then((res) => {
          console.log(res);
          this.logcontent = res.data;
        })
        .catch((err) => {
          console.log(err);
        });
    },
    changeTable() {
      this.searchKey = "";
      api.logs
        .getLogContent(this.getSelectedLog, this.pageOffset, this.pageLimit)
        .then((res) => {
          console.log(res);
          this.logcontent = res.data.Content;
        })
        .catch((err) => {
          console.log(err);
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
    getSelectedLog(newValue, oldValue) {
      clearInterval(this.reload);
      console.log(oldValue, newValue);
      this.changeTable();
      this.reload = setInterval(() => {
        this.changeTable();
      }, 30000);
    },
  },
  mounted() {
    this.changeTable();
    this.reload = setInterval(() => {
      this.changeTable();
    }, 30000);
  },
};
</script>
<style scoped>
.content-page {
  height: 100%;
}
.content-body {
  padding: 0 10px;
  height: 80vh;
  overflow-y: scroll;
}
.content-item {
  padding-bottom: 10px;
  text-align: left;
}
.search input {
  width: 85%;
  height: 30px;
  margin-right: 10px;
}
.search button {
  width: 10%;
  height: 35px;
  background-color: #c6ddfc;
  border: none;
}
.base-time {
  color: #75715e;
}
.base-host {
  color: #c63;
  border-color: #c63;
}
.base-app {
  color: #66d9ef;
  border-color: #66d9ef;
}
.base-text {
  color: #272822 !important;
}
.base-lvl-warn {
  background-color: #fd971f;
}
.base-lvl-error {
  background-color: #f92672;
}
</style>
