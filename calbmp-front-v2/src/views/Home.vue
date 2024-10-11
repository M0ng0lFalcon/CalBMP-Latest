<template>
  <div class="home">
    <div id="headerwrap">
      <header class="clearfix mt-5">
        <p>Welcome to CalBMP</p>
        <p>The platform to manage your farm with best</p>
        <p>management practice.</p>
      </header>
    </div>

    <!-- ==== FOOTERWRAP ==== -->
    <!-- <div id="footerwrap">
      <div class="container">
        <h5>
          Modeler <a href="#">Jingyang Xue</a> Created by <a href="#">Naqin</a>
          - Copyright 2021
        </h5>
      </div>
    </div> -->
  </div>
</template>

<script>
export default {
  methods: {
    clear_local_storage() {
      window.localStorage.removeItem("first_emergence");
      window.localStorage.removeItem("last_harvest");
    },
  },
  mounted() {
    this.clear_local_storage();
  },
};
</script>

<style>
body {
  height: 100%;
  display: flex;
  flex-direction: column;
}

#headerwrap {
  height: 100vh;
  margin-top: -50px;
  padding-top: 120px;
  text-align: center;
  background-size: cover;
  background: url("../assets/header_bg.jpg") no-repeat center center;
  width: 100%;
  background-size: cover;
}

#headerwrap h1 {
  color: #ffffff;
  margin-top: 90px;
  padding-bottom: 30px;
  letter-spacing: 8px;
  font-size: 100px;
  font-weight: bold;
  text-shadow: 2px 2px 2px #000;
}

#headerwrap p {
  color: #ffffff;
  font-size: 40px;
  font-weight: 400;
  text-align: left;
  line-height: 40px;
  letter-spacing: 3px;
  margin-left: 15px;
  text-shadow: 1px 1px 1px #000;
}

@font-face {
  font-family: "amplesoft";
  src: url("../fonts/FontsFree-Net-AmpleSoftProMedium.ttf");
}

.clearfix {
  font-family: "amplesoft";
}

#footerwrap {
  height: 10vh;
  background: #2b2b2b;
  margin-top: 0px;
  padding-top: 40px;
  padding-bottom: 20px;
  border-bottom: 1px solid #ddd;
  text-align: center;
}

#footerwrap h5 {
  color: white;
}
</style>
