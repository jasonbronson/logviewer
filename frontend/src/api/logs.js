import { instance as axios } from "../plugins/axios";

export const logs = {
  async getLogs(table, pageOffset, pageLimit) {
    try {
      let paging = "";
      if(table == undefined){
        table = "";
      }else{
        paging = `?pageOffset=${pageOffset}&pageLimit=${pageLimit}`
      }
      return await axios.get(
        `/logs${table}${paging}`
      );
    } catch (err) {
      if (err.response) {
        return err.response.data;
      }
      return err;
    }
  },
  async getLogContent(filename, pageOffset, pageLimit) {
    try {
      let paging = "";
      if(filename == undefined){
        filename = "";
      }else{
        paging = `?pageOffset=${pageOffset}&pageLimit=${pageLimit}`
      }
      return await axios.get(
        `/logs/${filename}${paging}`
      );
    } catch (err) {
      if (err.response) {
        return err.response.data;
      }
      return err;
    }
  },
  async getSearchContent(filename, keysearch, pageLimit, pageOffset) {
    try {
      return await axios.get(
        `/logs/search/${filename}?searchKey=${keysearch}&pageLimit=${pageLimit}&pageOffset=${pageOffset}`
      );
    } catch (err) {
      if (err.response) {
        return err.response.data;
      }
      return err;
    }
  },
};
