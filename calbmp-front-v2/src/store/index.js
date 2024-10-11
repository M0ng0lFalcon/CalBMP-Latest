import Vue from "vue";
import Vuex from "vuex";
import Storage from "../util/StorageHandle";
import SessionStorage from "../util/SessionHandle";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    // user info
    isLogin: SessionStorage.getItem("isLogin") || false,
    userInfo: SessionStorage.getItem("userInfo") || "",
    token: SessionStorage.getItem("token") || "",

    // user input step 1 and 2
    step1: Storage.getItem("step1") || null,
    step2: Storage.getItem("step2") || null,
    created_time: Storage.getItem("created_time") || null,

    // bmp list
    echartList: Storage.getItem("echartList") || [],
    bmp_options_list: Storage.getItem("bmp_options_list") || [],

    // vfs list
    vfsm_list: Storage.getItem("vfsm_list") || [],
    vfsm_options_list: Storage.getItem("vfsm_option_list") || [],
    vfsm_max_id: SessionStorage.getItem("vfsm_max_id") || [],
    vfsm_min_id: SessionStorage.getItem("vfsm_min_id") || [],

    // global time interval
    first_emergence: Storage.getItem("first_emergence") || null,
    last_harvest: Storage.getItem("last_harvest") || null,
  },
  mutations: {
    // * ===================================================== user input step 1 and 2
    SetStep1Info(state, Step1Info) {
      state.step1 = Step1Info;
      Storage.setItem("step1", Step1Info);
    },

    ClearStep1Info(state) {
      state.step1 = null;
      Storage.removeItem("step1");
    },

    SetStep2Info(state, Step2Info) {
      state.step2 = Step2Info;
      Storage.setItem("step2", Step2Info);
    },

    ClearStep2Info(state) {
      state.step2 = null;
      Storage.removeItem("step2");
    },

    SetCreatedTime(state, _created_time) {
      state.created_time = _created_time;
      Storage.setItem("created_time", _created_time);
    },

    ClearCreatedTime(state) {
      state.created_time = null;
      Storage.removeItem("created_time");
    },

    // * ===================================================== bmp
    SetEchartList(state, echartList) {
      state.echartList = echartList;
      Storage.setItem("echartList", echartList);
    },

    ClearBmp(state) {
      state.echartList = [];
      Storage.removeItem("echartList");
    },

    SetBmpOptionsList(state, bmp_options_list) {
      state.bmp_options_list = bmp_options_list;
      Storage.setItem("bmp_options_list", bmp_options_list);
    },

    ClearBmpOptionsList(state) {
      state.bmp_options_list = [];
      Storage.removeItem("bmp_options_list");
    },

    // * ===================================================== vfs
    SetVfsmList(state, vfsm_list) {
      state.vfsm_list = vfsm_list;
      Storage.setItem("vfsm_list", vfsm_list);
    },

    ClearVfsm(state) {
      state.vfsm_list = [];
      Storage.removeItem("vfsm_list");
    },

    SetVfsmOptionList(state, vfsm_options_list) {
      state.vfsm_options_list = vfsm_options_list;
      Storage.setItem("vfsm_options_list", vfsm_options_list);
    },

    ClearVfsmOptionList(state) {
      state.vfsm_options_list = [];
      Storage.removeItem("vfsm_options_list");
    },

    SetVfsmMinId(state, vfsm_min_id) {
      state.vfsm_min_id = vfsm_min_id;
      SessionStorage.setItem("vfsm_min_id", vfsm_min_id);
    },

    SetVfsmMaxId(state, vfsm_max_id) {
      state.vfsm_max_id = vfsm_max_id;
      SessionStorage.setItem("vfsm_max_id", vfsm_max_id);
    },

    // * ===================================================== user info
    SetIsLogin(state, isLogin) {
      state.isLogin = isLogin;
      SessionStorage.setItem("isLogin", isLogin);
    },

    SetToken(state, _token) {
      state.token = _token;
      SessionStorage.setItem("token", _token);
    },

    ClearIsLogin(state) {
      state.isLogin = false;
      SessionStorage.removeItem("isLogin");
    },

    SetUserInfo(state, _userInfo) {
      state.userInfo = _userInfo;
      SessionStorage.setItem("userInfo", _userInfo);
    },
    ClearUserInfo(state) {
      state.userInfo = "";
      SessionStorage.removeItem("userInfo");
    },

    // * ===================================================== global time interval
    SetTimeInterval(state, TimeInterval) {
      console.log(TimeInterval);
      state.first_emergence = TimeInterval.first_emergence;
      state.last_harvest = TimeInterval.last_harvest;
      Storage.setItem("first_emergence", TimeInterval.first_emergence);
      Storage.setItem("last_harvest", TimeInterval.last_harvest);
    },
  },
  actions: {},
  modules: {},
});
