<template>
  <div class="Register">
    <b-row>
      <b-col md="8" offset-md="2" lg="6" offset-lg="3">
        <b-card title="Register" class="mt-5">
          <b-form @submit="onSubmit">
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

            <!-- button -->
            <b-form-group>
              <b-button variant="danger" block type="submit">
                register
              </b-button>
            </b-form-group>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>

<script>
import { Register } from "@/server/userServer";

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
    onSubmit(event) {
      event.preventDefault();
      this.register();
    },
    async register() {
      console.log("register");
      // check password
      if (this.user.password.length < 8) {
        this.$bvToast.toast("Password at least 8 characters.", {
          title: "Register Error",
          toaster: "b-toaster-top-center",
          variant: "danger",
          solid: true,
        });
        return;
      }

      // make formdata
      let form_data = new FormData();
      form_data.append("username", this.user.username);
      form_data.append("password", this.user.password);
      const register_flag = await Register(form_data);
      const _this = this;

      if (register_flag) {
        this.$bvToast.toast(`Add new user: ${this.user.username}`, {
          title: "Register successfully",
          toaster: "b-toaster-top-center",
          variant: "success",
          solid: true,
        });
        // replace router when register successfully
        // timeout 1 second
        setTimeout(function () {
          console.log("replace router");
          _this.$router.replace("/login");
        }, 1000);
      } else {
        this.$bvToast.toast("User already exists.", {
          title: "Register Error",
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

.Register {
  font-family: "amplesoft";
}
</style>
