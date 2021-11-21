<template>
  <div class="content-page">
    <div class="content-body" v-on:scroll="handleScroll">
      <div
        v-if="!nothingFound"
        v-for="(content, idx) in logcontent"
        :key="idx"
        class="content-item"
        ref="logviewer"
      >
        <span v-html="highlight(content)"></span>
      </div>
      <div v-if="nothingFound">
        Nothing Found in Search
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
      pageOffset: 20,
      pageLimit: 20,
      reload: null,
      searchKey: "",
      nothingFound: false,
      loadingLogs: false,
    };
  },
  computed: {
    getSelectedLog() {
      return this.$store.state.selectedLog;
    },
  },
  methods: {
    handleScroll(event) {
      if (!this.loadingLogs && event.target.scrollTop < 80) {
        console.log("scrolling ", event.target.scrollTop);
        this.pageOffset += this.pageLimit;
        this.getLogs();
      }
    },
    highlight(data) {
      //console.log("********", this.searchKey, data);
      if (!this.searchKey) return data;
      if (data != typeof string) {
        data = data.toString();
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
        console.log("scroll to end");
        el.scrollIntoView({ block: "end" });
      }
    },
    getLogs() {
      this.loadingLogs = true;
      api.logs
        .getSearchContent(
          this.getSelectedLog,
          this.searchKey,
          this.pageLimit,
          this.pageOffset
        )
        .then((res) => {
          this.loadingLogs = false;
          if (!res.data) {
            this.nothingFound = true;
            return;
          }
          this.logcontent = res.data;
          this.scrollToElement();
        })
        .catch((err) => {
          console.log(err);
        });
    },
    searchHandle() {
      this.nothingFound = false;
      clearInterval(this.reload);
      this.getLogs();
    },
    changeLog() {
      this.nothingFound = false;
      this.searchKey = "";
      this.getLogs();
    },
  },
  watch: {
    getSelectedLog(newValue, oldValue) {
      clearInterval(this.reload);
      console.log(oldValue, newValue);
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
