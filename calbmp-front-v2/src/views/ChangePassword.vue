<template>
  <div class="changePassword">
    <b-row>
      <b-col md="8" offset-md="2" lg="6" offset-lg="3">
        <b-card title="Change Password" class="mt-5">
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
            <b-form-group label="New Password">
              <b-form-input
                v-model="user.password"
                type="password"
                placeholder="Please input password"
                required
              ></b-form-input>
            </b-form-group>

            <b-form-group label="Confirm Password">
              <b-form-input
                v-model="user.confirm_password"
                type="password"
                placeholder="Please input password"
                required
              ></b-form-input>
            </b-form-group>

            <b-form-group>
              <b-button variant="danger" block type="submit">Submit</b-button>
            </b-form-group>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>

<script>
import { ChangePassword } from "@/server/userServer";

export default {
  name: "ChangePassword",
  data() {
    return {
      user: {
        username: "",
        password: "",
        confirm_password: "",
      },
    };
  },

  // method
  methods: {
    checkLogin() {
      const isLogin = this.$store.state.isLogin;
      if (isLogin === true) {
        this.$router.replace("/userInputStep1");
      } else {
        this.$router.replace("/");
      }
    },
    async onSubmit(event) {
      event.preventDefault();
      if (this.user.password !== this.user.confirm_password) {
        this.$bvToast.toast(
          "The password does not match the confirmation password!!!",
          {
            title: "Change Password Error",
            toaster: "b-toaster-top-center",
            variant: "danger",
            solid: true,
          }
        );
      } else {
        const resObj = await ChangePassword(this.user);
        if (!resObj.flag) {
          this.$bvToast.toast(resObj.msg, {
            title: "Change Password Error",
            toaster: "b-toaster-top-center",
            variant: "danger",
            solid: true,
          });
        } else {
          this.$store.commit("ClearIsLogin");
          this.checkLogin();
        }
      }
    },
  },
};
</script>

<style scoped></style>
