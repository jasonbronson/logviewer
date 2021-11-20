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
      <div
        v-for="(content, idx) in logcontent"
        :key="idx"
        class="content-item"
        ref="logviewer"
      >
        <span v-html="highlight(content)"></span>
      </div>
    </div>
    <div class="search">
      <input
        placeholder="Enter text for searching..."
        v-model="searchKey"
        class="search-input"
        @keyup.enter="searchHandle()"
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
    highlight(data) {
      //console.log("********", this.searchKey, data);
      if (!this.searchKey || !data) return data;
      if (data != typeof string) {
        //data = data.toString();
        return data;
      }
      var regex = new RegExp(this.searchKey, "gi");
      return data.replace(regex, function(str) {
        return "<span style='background-color: yellow;'>" + str + "</span>";
      });
    },
    scrollToElement() {
      const el = this.$refs.logviewer;
      if (el) {
        // Use el.scrollIntoView() to instantly scroll to the element
        el.scrollIntoView({ behavior: "smooth" });
      }
    },
    searchHandle() {
      clearInterval(this.reload);
      api.logs
        .getSearchContent(this.getSelectedLog, this.searchKey, this.pageLimit)
        .then((res) => {
          this.logcontent = res;
        })
        .catch((err) => {
          console.log(err);
        });
    },
    changeLog() {
      this.searchKey = "";
      api.logs
        .getLogContent(this.getSelectedLog, this.pageOffset, this.pageLimit)
        .then((res) => {
          this.logcontent = res.data.Content;
        })
        .catch((err) => {
          console.log(err);
        });
    },
  },
  watch: {
    getSelectedLog(newValue, oldValue) {
      clearInterval(this.reload);
      //console.log(oldValue, newValue);
      this.changeLog();
      // this.reload = setInterval(() => {
      //   this.changeLog();
      // }, 5000);
    },
  },
  mounted() {
    this.changeLog();
    // this.reload = setInterval(() => {
    //   this.changeLog();
    //   this.scrollToElement();
    // }, 5000);
    this.scrollToElement();
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
.search {
  padding: 10px;
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
