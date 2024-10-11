import store from "../store";
import Vue from "vue";

export function Login(user) {
  const api = "/user/login";
  console.log("login");
  console.log(user);
  return Vue.axios
    .post(api, user)
    .then((res) => {
      console.log(res);
      store.commit("SetToken", `Bearer ${res.data.data.token}`);
      store.commit("SetIsLogin", true);
      const UserInfo = {
        username: user.get("username"),
      };
      store.commit("SetUserInfo", UserInfo);
      return true;
    })
    .catch((err) => {
      console.log(err);
      return false;
    });
}

export function Register(user) {
  const api = "/user/register";
  return Vue.axios
    .post(api, user)
    .then((res) => {
      console.log(res);
      return true;
    })
    .catch((err) => {
      console.log(err);
      return false;
    });
}

export function ChangePassword(user) {
  const api = "/user/change_password";
  return Vue.axios
    .post(api, user)
    .then((res) => {
      console.log(res);
      return {
        flag: true,
        msg: "Change password successfully",
      };
    })
    .catch((err) => {
      console.log(err.response.data.msg);
      return {
        flag: false,
        msg: err.response.data.msg,
      };
    });
}
