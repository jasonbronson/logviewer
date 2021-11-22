<template>
  <div class="form_wrapper">
    <div class="form_container">
      <div class="title_container">
        <h2>Login</h2>
      </div>
      <div class="row clearfix">
        <div class="">
          <form>
            <div class="input_field">
              <span><i aria-hidden="true" class="fa fa-envelope"></i></span>
              <input
                v-model="email"
                type="email"
                name="email"
                placeholder="Email"
                required
              />
            </div>
            <div class="input_field">
              <span><i aria-hidden="true" class="fa fa-lock"></i></span>
              <input
                v-model="password"
                type="password"
                name="password"
                placeholder="Password"
                required
              />
            </div>
            <div class="btn__login">
              <button @click.prevent="login">Login</button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import api from "@/api";
export default {
  data() {
    return {
      email: "",
      password: "",
    };
  },
  components: {
  },

  computed: {
  },

  methods: {
    async login() {
      if ((this.email !== "", this.password !== "")) {
        try {
          let payload = {
            email: this.email,
            password: this.password,
          };
          api.auth
            .login(payload)
            .then((res) => {
              if (res.status === 200) {
                this.$store.commit("setAuthenticated", res.data);
                this.$router.push("/");
              }
            })
            .catch((err) => {
              console.log(err);
              this.$notify.error({
                title: "Error",
                message: "Email or Password incorrect",
              });
            });
        } catch (error) {
          console.log("err", error)
        }
      } else {
        this.$notify({
          title: "Network error logging in",
          message: "Email or Password not empty!",
          type: "error",
        });
      }
    },
  },
};
</script>
<style  scoped>
.input_field {
  padding: 5px;
}
.btn__login {
  margin-top: 10px;
  display: flex;
  width: 100%;
  justify-content: center;
}
.btn__login span a{
    color: #3498db;
    font-size: 13px;
    line-height: 1.6;
}
.btn__login button {
  border: 0;
  font-size: 14px;
  cursor: pointer;
  font-weight: 400;
  padding: 10px 15px;
  position: relative;
  background: #c6ddfc;
  white-space: nowrap;
  display: inline-block;
  text-decoration: none;
}

p {
  color: #555;
  margin: 0 0 10px;
  font-size: 13px;
  line-height: 1.6;
}
p a {
  color: #3498db;
  outline: 0 !important;
  text-decoration: none;
}
.form_wrapper input[type="text"]:focus,
.form_wrapper input[type="email"]:focus,
.form_wrapper input[type="password"]:focus {
  border: 1px solid;
  background: #fafafa;
  border-color: #bbb;
  box-shadow: 0 0 2px #c9c9c9;
}
</style>
