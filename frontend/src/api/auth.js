import { instance as axios } from "../plugins/axios";

export const auth = {
  async login(payload) {
    try {
      return await axios.post(
        `/login`, payload
      );
    } catch (err) {
      if (err.response) {
        return err.response.data;
      }
      return err;
    }
  },
};