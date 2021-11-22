export const localstorage = {
  setToken(token) {
    localStorage.setItem("token_user", token);
  },
  getToken() {
    let token = localStorage.getItem("token_user");
    return token;
  },
  clearAuth() {
    localStorage.removeItem("token_user");
  },
};
