import Vue from "vue";
import store from "../store";

export function AddHistory(CreatedTime) {
  const api = "/history/addHistory";
  const username = store.state.userInfo.username;
  const org_step1 = store.state.step1;
  const org_step2 = store.state.step2;
  const created_time = CreatedTime;
  const echart_list = store.state.echartList.length;
  const HistoryRec = {
    username: username,
    step1: JSON.stringify(org_step1),
    step2: JSON.stringify(org_step2),
    created_time: created_time,
    echart_list: echart_list,
    zip_code: org_step1.zip_code,
    project_name: org_step2.Crop,
    comp_name: org_step1.CompName,
    county: org_step1.county,
    muname: org_step1.Muname,
  };
  return Vue.axios.post(api, HistoryRec);
}

export function CheckHistory() {
  const username = store.state.userInfo.username;
  const api = "/history/checkHistory?username=" + username;
  return Vue.axios.get(api).then((res) => {
    const Data = res.data.data;
    const HistoryList = Data.history_list;
    console.log("historyList", HistoryList);
    return HistoryList;
  });
}

export async function ToVisualization(HistoryInfo) {
  console.log("historyInfo", HistoryInfo);
  const step1 = HistoryInfo.Step1;
  const step2 = HistoryInfo.Step2;
  const created_time = HistoryInfo.CreatedDate;
  const echart_list = HistoryInfo.EchartList;
  let echart_list_res = [];
  for (let i = 0; i < echart_list; i += 1) {
    echart_list_res.push({
      bmp_id: i + 1,
    });
  }

  store.commit("SetStep1Info", JSON.parse(step1));
  store.commit("SetStep2Info", JSON.parse(step2));
  store.commit("SetCreatedTime", created_time);
  store.commit("SetEchartList", echart_list_res);
}
