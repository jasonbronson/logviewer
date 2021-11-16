import { instance as axios } from "../plugins/axios";

export const logs = {
  // async setLogs(table) {
  //   try {
  //     return await axios.post("/logs/create/" + table);
  //   } catch (err) {
  //     if (err.response) {
  //       return err.response.data;
  //     }
  //     return err;
  //   }
  // },
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
};
