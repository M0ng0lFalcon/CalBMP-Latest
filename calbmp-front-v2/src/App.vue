<template>
  <div id="app">
    <cus-nav></cus-nav>

    <!-- if route is home, no container -->
    <div v-if="$route.name === 'Home'">
      <router-view />
    </div>

    <div v-if="$route.name !== 'Home' && $route.name !== 'Visualization'">
      <b-container>
        <router-view />
      </b-container>
    </div>

    <div v-if="$route.name !== 'Home' && $route.name === 'Visualization'">
      <b-container fluid>
        <router-view />
      </b-container>
    </div>
  </div>
</template>

<script>
import CusNav from "./components/CusNav.vue";

export default {
  components: {
    CusNav,
  },
};
</script>
