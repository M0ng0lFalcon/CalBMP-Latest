<template>
  <div class="CusNav">
    <b-navbar toggleable="lg" type="dark" variant="dark">
      <b-container>
        <b-navbar-brand class="cus_nav" @click="checkLogin">
          CalBMP
        </b-navbar-brand>

        <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

        <b-collapse id="nav-collapse" is-nav>
          <!-- Text -->
          <b-navbar-nav>
            <b-nav-text
              style="color: rgb(255 255 255); font-family: 'amplesoft'"
            >
              —— Apps for Pesticide Management on Farm
            </b-nav-text>
          </b-navbar-nav>
          <!-- Right aligned nav items -->
          <b-navbar-nav class="ml-auto" v-if="$route.name != 'Userinput'">
            <b-button-group>
              <b-button
                @click="$router.push('login')"
                variant="success"
                v-if="$route.name === 'Register' || $route.name === 'Home'"
              >
                Login
              </b-button>

              <b-button
                @click="$router.push('register')"
                variant="danger"
                v-if="$route.name === 'Login' || $route.name === 'Home'"
              >
                Register
              </b-button>
            </b-button-group>
          </b-navbar-nav>
          <b-navbar-nav v-if="$store.state.isLogin" class="ml-auto">
            <b-nav-item-dropdown :text="$store.state.userInfo.username">
              <b-dropdown-item @click="toHistory">History</b-dropdown-item>
              <b-dropdown-item @click="toManual">Manual</b-dropdown-item>
              <b-dropdown-item @click="toChangePassword">
                Change Pwd
              </b-dropdown-item>
              <b-dropdown-divider></b-dropdown-divider>
              <b-dropdown-item @click="signOut" variant="danger">
                sign out
              </b-dropdown-item>
            </b-nav-item-dropdown>
          </b-navbar-nav>
        </b-collapse>
      </b-container>
    </b-navbar>
  </div>
</template>

<script>
export default {
  name: "cus-nav",
  methods: {
    checkLogin() {
      const isLogin = this.$store.state.isLogin;
      if (isLogin === true) {
        this.$router.push("/userInputStep1");
      } else {
        this.$router.replace("/");
      }
    },

    signOut() {
      this.$store.commit("ClearIsLogin");
      this.checkLogin();
    },

    toHistory() {
      this.$router.push("/history");
    },

    toManual() {
      this.$router.push("/manual");
    },

    toChangePassword() {
      this.$router.push("/change_password");
    },
  },
};
</script>

<style>
@font-face {
  font-family: "AbrilFatface";
  src: url("../fonts/AbrilFatface-Regular.otf");
}

@font-face {
  font-family: "amplesoft";
  src: url("../fonts/FontsFree-Net-AmpleSoftProMedium.ttf");
}

.CusNav {
  font-family: "amplesoft";
}

.cus_nav {
  font-family: "AbrilFatface";
}

.cus_nav:hover {
  cursor: pointer;
}
</style>
