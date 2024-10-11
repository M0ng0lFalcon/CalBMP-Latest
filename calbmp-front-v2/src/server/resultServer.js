import Vue from "vue";
import store from "../store";

export function getBasicResult(
  waterBalanceComponents,
  pesticideComponentsOpt,
  sedimentComponentsOpt,
  concentrationOpt,
  postData
) {
  // add to post data
  postData.water = GetSelectedOptions(waterBalanceComponents);
  postData.pesticide = GetSelectedOptions(pesticideComponentsOpt);
  postData.sediment = GetSelectedOptions(sedimentComponentsOpt);
  postData.concentration = GetSelectedOptions(concentrationOpt);
  postData.created_time = store.state.created_time;

  // get data
  const api = "/result/getBasicByName";
  return Vue.axios
    .post(api, postData)
    .then((res) => {
      return res.data.data.result;
    })
    .catch((err) => {
      console.log(err);
    });
}

export function GetComparisonData(BmpCnt, Harvest) {
  const created_time = store.state.created_time;
  const url = `/result/getComparisonData`;
  const params = `?bmp_cnt=${BmpCnt}&harvest=${Harvest}&created_time=${created_time}`;
  const api = url + params;
  return Vue.axios.get(api);
}

// ! util fun for get selected data
export function GetSelectedOptions(OptionsObj) {
  let seletedData = [];
  Object.keys(OptionsObj).forEach((key) => {
    const now = OptionsObj[key];
    if (now.status === "1") {
      seletedData.push(now.value);
    }
  });
  return seletedData;
}

export function GetEventBasedOutputs(
  chemical_name,
  first_emergence,
  last_harvest,
  created_time
) {
  const api = `/result/GetEventBasedOutputs?chemical_name=${chemical_name}&created_time=${created_time}`;
  return Vue.axios.get(api).then((res) => {
    console.log(res);
    const Data = res.data.data;
    const Res = Data.event_based_outputs;

    const runoff_date = Res.runoff_date;
    const runoff_len = runoff_date.length;
    const scenario_value = Res.scenario_value;
    const benchmark = Res.benchmark;

    // location emergence date and harvest date
    const emergenceObj = new Date(first_emergence);
    const harvestObj = new Date(last_harvest);
    let idx1 = 0;
    let idx2 = runoff_len;
    for (let i = 0; i < runoff_len; i += 1) {
      const temp = new Date(runoff_date[i]);
      if (temp >= emergenceObj) {
        idx1 = i;
        break;
      }
    }

    for (let i = 0; i < runoff_len; i += 1) {
      const temp = new Date(runoff_date[i]);
      if (temp <= harvestObj) {
        idx2 = i;
      }
    }
    console.log("idx1", idx1);
    console.log("idx2", idx2);

    let event_based_outputs = [
      { EventBasedOutputs: "Runoff date" },
      { EventBasedOutputs: "Runoff (cm)" },
      { EventBasedOutputs: "Precipitation (cm)" },
      { EventBasedOutputs: "Irrigation (cm)" },
      { EventBasedOutputs: "Sediment (t)" },
      { EventBasedOutputs: "Pesticide Loading in Runoff (lbs/acre)" },
      { EventBasedOutputs: "Pesticide Loading in Erosion (lbs/acre)" },
      { EventBasedOutputs: "Pesticide Loading in Volatilization (lbs/acre)" },
      { EventBasedOutputs: "Pesticide Concentration in Runoff (ppm)" },
      { EventBasedOutputs: "US EPA aquatic life benchmark (ppm)" },
    ];

    const len = event_based_outputs.length;
    let column_num = 1;
    for (let i = idx1; i <= idx2; i += 1) {
      // set runoff date
      event_based_outputs[0]["_" + column_num.toString()] = runoff_date[i];
      // set benchmark data
      event_based_outputs[len - 1]["_" + column_num.toString()] = benchmark;

      column_num += 1;
    }

    for (let i = 1; i < len - 1; i += 1) {
      column_num = 1;
      for (let j = idx1; j <= idx2; j += 1) {
        event_based_outputs[i]["_" + column_num.toString()] = KeepDecimalPlaces(
          parseFloat(scenario_value[i - 1][j])
        );
        column_num += 1;
      }
    }

    console.log(JSON.stringify(event_based_outputs));

    return event_based_outputs;
  });
}

export function GetTextResult(
  chemical_name,
  first_emergence,
  last_harvest,
  created_time,
  field_size
) {
  const api = "/result/get_text_result";
  return Vue.axios
    .post(api, {
      pesticide_list: chemical_name,
      created_time: created_time,
      field_size: field_size,
    })
    .then((res) => {
      const text_res = res.data.data.text_res;

      const date = text_res.date;
      const date_len = date.length;
      const water = text_res.water;
      const pesticide = text_res.pesticide;
      const sediment = text_res.sediment;
      const concentration = text_res.concentration;
      const benchmark = text_res.benchmark;
      const benchmark_len = Object.keys(benchmark).length;

      // location emergence date and harvest date
      const emergenceObj = new Date(first_emergence);
      const harvestObj = new Date(last_harvest);
      let idx1 = 0;
      let idx2 = date_len;
      for (let i = 0; i < date_len; i += 1) {
        const temp = new Date(date[i]);
        if (temp >= emergenceObj) {
          idx1 = i;
          break;
        }
      }

      for (let i = 0; i < date_len; i += 1) {
        const temp = new Date(date[i]);
        if (temp <= harvestObj) {
          idx2 = i;
        }
      }
      console.log("text_res:", text_res);
      console.log("idx1, idx2:", idx1, idx2);

      let event_based_outputs = [
        { EventBasedOutputs: "Runoff date" },
        { EventBasedOutputs: "Runoff (cm)" },
        { EventBasedOutputs: "Precipitation (cm)" },
        { EventBasedOutputs: "Irrigation (cm)" },
        { EventBasedOutputs: "Sediment (t)" },
        { EventBasedOutputs: "Pesticide Loading in Runoff (lbs/acre)" },
        { EventBasedOutputs: "Pesticide Loading in Erosion (lbs/acre)" },
        { EventBasedOutputs: "Pesticide Loading in Volatilization (lbs/acre)" },
        { EventBasedOutputs: "Pesticide Concentration in Runoff (ppm)" },
        { EventBasedOutputs: "US EPA aquatic life benchmark (ppm)" },
      ];
      const len = event_based_outputs.length;
      let column_num = 1;
      for (let i = idx1; i <= idx2; i += 1) {
        let col_name = "_" + column_num.toString();
        // set runoff date
        event_based_outputs[0][col_name] = date[i];
        // set benchmark data
        let benchmark_list = [];
        Object.keys(benchmark).forEach((key) => {
          benchmark_list.push(key + ":" + benchmark[key]);
        });
        event_based_outputs[len - 1][col_name] = benchmark_list.join("\n");

        // water
        event_based_outputs[1][col_name] = water["RUNF"][i]; // runoff
        event_based_outputs[2][col_name] = water["PRCP"][i]; // precipitation
        event_based_outputs[3][col_name] = water["IRRG"][i]; // irrigation

        // sediment
        event_based_outputs[4][col_name] = sediment["ESLS"][i]; // sediment

        // pesticide
        let rflx_data = [];
        let eflx_data = [];
        let vflx_data = [];
        for (let j = 1; j <= benchmark_len; j += 1) {
          rflx_data.push(
            pesticide[`RFLX_${chemical_name[j - 1]}_${j}_TSER`][i]
          );
          eflx_data.push(
            pesticide[`EFLX_${chemical_name[j - 1]}_${j}_TSER`][i]
          );
          vflx_data.push(
            pesticide[`VFLX_${chemical_name[j - 1]}_${j}_TSER`][i]
          );
        }
        event_based_outputs[5][col_name] = rflx_data.join(" ; ");
        event_based_outputs[6][col_name] = eflx_data.join(" ; ");
        event_based_outputs[7][col_name] = vflx_data.join(" ; ");

        // concentration
        let concentration_data = [];
        Object.keys(benchmark).forEach((key) => {
          concentration_data.push(concentration[key][i]);
        });
        event_based_outputs[8][col_name] = concentration_data.join(" ; ");

        column_num += 1;
      }

      return event_based_outputs;
    })
    .catch((err) => {
      console.log(err);
    });
}

export function GetBmpSummaryData(bmp_cnt, harvest) {
  return GetComparisonData(bmp_cnt, harvest).then((res) => {
    console.log("comparison", res);
    const Res = res.data.data;
    const runoff = Res.comparison_runoff;
    const erosion = Res.comparison_erosion;
    const volatilization = Res.comparison_volatilization;

    let res_array = [];
    const len = runoff.length;
    let i = 0;
    for (i; i < len; i += 1) {
      res_array.push({
        BMP_Scenario: i + 1,
        BMP_Options: store.state.bmp_options_list[i].toString(),
        pesticide_reduction_in_Runoff: KeepDecimalPlaces(runoff[i]),
        pesticide_reduction_in_Erosion: KeepDecimalPlaces(erosion[i]),
        pesticide_reduction_in_Volatilization: KeepDecimalPlaces(
          volatilization[i]
        ),
      });
    }

    // VFSM summary测试
    if (store.state.vfsm_list.length !== 0) {
      res_array.push({
        BMP_Scenario: i + 1,
        BMP_Options: "VFS Mode",
        pesticide_reduction_in_Runoff: 74.6,
        pesticide_reduction_in_Erosion: 76.3,
        pesticide_reduction_in_Volatilization: 0,
      });
    }

    return res_array;
  });
}

export function GetInputFiles(created_at) {
  const api = `/result/get_input_files?created_time=${created_at}`;
  return Vue.axios.get(api).then((res) => {
    console.log("get input files:", res.data.data);
    return res.data.data.input_files;
  });
}

export function DownloadInputFile(file_list, created_time) {
  const api = "/result/zip_input_file";
  const api_down = `/result/download_input_file?created_time=${created_time}`;
  const data = {
    file_list: file_list,
    created_time: created_time,
  };
  return Vue.axios.post(api, data).then((res) => {
    console.log(res);
    Vue.axios.get(api_down, { responseType: "blob" }).then((fres) => {
      const elink = document.createElement("a");
      elink.style.display = "none";
      elink.href = URL.createObjectURL(fres.data);
      console.log(elink.href);
      document.body.appendChild(elink);
      elink.click();
      URL.revokeObjectURL(elink.href);
      document.body.removeChild(elink);
    });
  });
}

// ! util of GetBmpSummaryData for Keep several decimal places
function KeepDecimalPlaces(DecimalNum) {
  const decimalStr = DecimalNum.toString();
  const len = decimalStr.length;
  // find out index of first not zero at float
  let idx = 0;
  let start = decimalStr.indexOf(".");
  for (let i = start + 1; i < len; i += 1) {
    if (decimalStr[i] !== "0") {
      idx = i;
      break;
    }
  }

  const cnt_of_zero = idx - start - 1;

  if (cnt_of_zero > 5) {
    // if cnt of zero > 5
    return parseFloat(decimalStr.slice(0, start)).toFixed(2);
  } else if (cnt_of_zero === 4) {
    return DecimalNum.toFixed(5);
  } else {
    return DecimalNum.toFixed(4);
  }
}
