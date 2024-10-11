<template>
  <div class="Login">
    <b-row>
      <b-col md="8" offset-md="2" lg="6" offset-lg="3">
        <b-card title="Login" class="mt-5">
          <b-form>
            <!-- username -->
            <b-form-group label="Username">
              <b-form-input
                v-model="user.username"
                type="text"
                placeholder="Username/PhoneNumber/E-Mail"
                required
              ></b-form-input>
            </b-form-group>

            <!-- password -->
            <b-form-group label="Password">
              <b-form-input
                v-model="user.password"
                type="password"
                placeholder="Please input password"
                required
              ></b-form-input>
            </b-form-group>

            <b-form-group>
              <b-button variant="link" @click="toChangePassword">
                Forget password?
              </b-button>
              <b-button variant="success" block @click="login">Login</b-button>
            </b-form-group>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>

<script>
import { Login } from "@/server/userServer";

export default {
  data() {
    return {
      user: {
        username: "",
        password: "",
      },
    };
  },
  methods: {
    toChangePassword() {
      this.$router.push("/change_password");
    },
    async login() {
      // 以下仅供测试样例
      console.log("login");
      // make formdata
      let form_data = new FormData();
      form_data.append("username", this.user.username);
      form_data.append("password", this.user.password);
      const login_flag = await Login(form_data);
      if (login_flag) {
        await this.$router.push("userinputStep1");
      } else {
        this.$bvToast.toast("Username or Passwor error !!!", {
          title: "Login Error",
          toaster: "b-toaster-top-center",
          variant: "danger",
          solid: true,
        });
      }
    },
  },
};
</script>

<style>
@font-face {
  font-family: "amplesoft";
  src: url("../fonts/FontsFree-Net-AmpleSoftProMedium.ttf");
}

.Login {
  font-family: "amplesoft";
}
</style>
