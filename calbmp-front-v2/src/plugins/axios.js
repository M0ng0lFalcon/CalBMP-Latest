"use strict";

import Vue from "vue";
import axios from "axios";
import store from "@/store";

const HOST = process.env.VUE_APP_HOST;
const PORT = process.env.VUE_APP_PORT;
const SUFFIX = process.env.VUE_APP_SUFFIX;
const BASE_URL = HOST + ":" + PORT + SUFFIX;
let config = {
  baseURL: BASE_URL, // Base URL
  timeout: 60 * 1000, // Timeout
};

const _axios = axios.create(config);

_axios.interceptors.request.use(
  function (config) {
    config.headers = {
      Authorization: store.state.token,
    };
    return config;
  },
  function (error) {
    // Do something with request error
    // eslint-disable-next-line no-debugger
    debugger;
    return Promise.reject(error);
  }
);

// Add a response interceptor
_axios.interceptors.response.use(
  function (response) {
    // Do something with response data
    return response;
  },
  function (error) {
    // Do something with response error
    return Promise.reject(error);
  }
);

// Org: Plugin.install = function (Vue, options) {
Plugin.install = function (Vue) {
  Vue.axios = _axios;
  window.axios = _axios;
  Object.defineProperties(Vue.prototype, {
    axios: {
      get() {
        return _axios;
      },
    },
    $axios: {
      get() {
        return _axios;
      },
    },
  });
};

Vue.use(Plugin);

export default Plugin;
