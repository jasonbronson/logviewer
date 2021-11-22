import { instance as axios } from "../plugins/axios";

export const logs = {
  async getLogs(logName, pageLimit, pageOffset) {
    try {
      let paging = "";
      if(logName == undefined){
        logName = "";
      }else{
        paging = `?pageOffset=${pageOffset}&pageLimit=${pageLimit}`
      }
      return await axios.get(
        `/logs/${logName}${paging}`
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
