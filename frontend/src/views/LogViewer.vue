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
        <div class="index"><span>{{indexFormat(idx)}}</span></div>
        <div><span v-html="highlight(content)"></span></div>
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
      scrollFetch: false,
    };
  },
  computed: {
    getSelectedLog() {
      return this.$store.state.selectedLog;
    },
  },
  methods: {
    handleScroll(event) {
      if (
        !this.scrollFetch &&
        !this.loadingLogs &&
        event.target.scrollTop < 80
      ) {
        console.log("scrolling ", event.target.scrollTop);
        this.pageOffset += this.pageLimit;
        this.getMoreLogs();
      }
    },
    highlight(data) {
      //console.log("********", this.searchKey, data);
      let pre = data.split('"-"')[0]
      let host = pre.split("- -")[0]
      let time = pre.substring(
                    data.indexOf("["),
                    data.indexOf("]") + 1
                );
      let text = data.replace(host, "").replace(time, "").replace("- -", "")
      let content  = `<span class='base-host'> ${host} </span> - -
                      <span class='base-time'> ${time} </span>
                      <span class='base-text'> ${text} </span>`
      data = data.replace(data, content)
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
    getMoreLogs() {
      this.loadingLogs = true;
      api.logs
        .getLogs(this.getSelectedLog, this.pageLimit, this.pageOffset)
        .then((res) => {
          this.loadingLogs = false;
          if (res.data.Total == 0) {
            this.scrollFetch = true;
            return;
          }
          this.logcontent = res.data.Content.concat(this.logcontent)
          this.scrollToElement();
        })
        .catch((err) => {
          console.log(err);
        });
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
          if (res.data.Total == 0) {
            this.nothingFound = true;
            return;
          }
          this.logcontent = res.data.Content;
          this.scrollToElement();
        })
        .catch((err) => {
          console.log(err);
        });
    },
    searchHandle() {
      this.nothingFound = false;
      this.scrollFetch = false;
      clearInterval(this.reload);
      this.getMoreLogs();
    },
    changeLog() {
      this.scrollFetch = false;
      this.nothingFound = false;
      this.pageOffset = 20;
      this.searchKey = "";
      this.getLogs();
    },
    indexFormat(i) {
      let idx = i +1 
      if (idx < 10) {
        idx = "0" + idx
      }
      return idx
    }
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
  height: 80vh;
  overflow-y: scroll;
}
.content-item {
  padding-bottom: 10px;
  text-align: left;
  display: flex;
}
.index{
  margin-right: 10px;
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
:deep() .base-time {
  color: #75715e !important;
}
:deep() .base-host {
  color: #c63 !important;
  border-color: #c63;
}
:deep() .base-app {
  color: #66d9ef;
  border-color: #66d9ef;
}
:deep() .base-text {
  color: #272822 !important;
}
:deep() .base-lvl-warn {
  background-color: #fd971f;
}
:deep() .base-lvl-error {
  background-color: #f92672;
}
</style>
