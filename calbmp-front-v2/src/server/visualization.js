import { getBasicResult, GetComparisonData } from "@/server/resultServer";
import { MakeChartObj } from "@/util/VisualizationUtil/EchartMaker";
import { Result2Visualization } from "./vfsmServer";
// comparison
import { ComparisonEchartMaker } from "@/server/echartsServer";
import store from "@/store";
// import variable
import {
  ConcentrationOpt,
  PesticideComponentsOpt,
  SedimentComponent,
  WaterBalanceComponents,
} from "@/variable/visualizationVar";

/**
 * 获取多个图表对象
 * @param postData 请求参数
 * @param WaterEchartId 水系目标图表id
 * @param PesticideEchartId 农药目标图表id
 * @param SedimentEchartId 残留目标图表id
 * @param ConcentrationEchartId 浓度目标图表id
 * @returns {Promise<{recommend_flag: boolean, width: *, SedimentEchart: *, PesticideEchartList: *[], overlayShow: boolean, WaterEchart: *, ConcentrationEchart: *[], height: *}>}
 * @constructor
 */
export async function GetCharts(
  postData,
  // id of echarts
  WaterEchartId,
  PesticideEchartId,
  SedimentEchartId,
  ConcentrationEchartId
) {
  // 1. get result data
  let result = Object();

  // 根据bmp类型，请求相应的API，返回dict[type]字典类型的数据
  // VFSM模式
  if (postData.scenario_type === "vfsm") {
    const min_id = store.state.vfsm_min_id;
    const max_id = store.state.vfsm_max_id;
    const created_time = store.state.created_time;
    result = await Result2Visualization(created_time, min_id, max_id); // 获取VFSM的结果
  } else {
    // 非结构化BMP模式
    result = await getBasicResult(
      // selected data of options
      WaterBalanceComponents,
      PesticideComponentsOpt,
      SedimentComponent,
      ConcentrationOpt,
      // form data
      postData
    );
  }

  // console.log("basic result:", result);

  // 2. val 2 var
  let benchmark = result.benchmark; // list
  let concentration = result.concentration; // obj: {chemical name 1:[...], chemical name 2:[...]}
  let date = result.date; // daily date list
  let pesticide = result.pesticide; // obj: {eflx1_tcum:[...], eflx1_tser:[...],...}
  let sediment = result.sediment; // obj: {esls:[...]}
  let water = result.water; // obj: {irrg:[...], ...}

  // 定义图表的y轴文字
  let water_y = "";
  let pesticide_y = "";
  let sediment_y = "";
  let concentration_y = "";

  if (postData.scenario_type === "vfsm") {
    water_y = "Water depth (Acre-feet)";
    pesticide_y = "Pesticide loading (lbs/acre)";
    sediment_y = "Eroded solids (lbs)";
    concentration_y = "Concentration in runoff (ug/L)";
  } else {
    water_y = "Water depth (Acre-feet)";
    pesticide_y = "Pesticide loading (lbs/acre)";
    sediment_y = "Eroded solids (lbs)";
    concentration_y = "Concentration in runoff (ug/L)";
  }

  // water charts
  if (
    Object.keys(water).includes("PRCP") &&
    store.state.step2.IrrigationType === 0
  ) {
    const irrigation_data = store.state.step2.irrigation_date;
    const irrigation_amount = store.state.step2.irrigation_amount;
    for (let idx = 0; idx < irrigation_data.length; idx += 1) {
      const irri_split = irrigation_data[idx].split("-");
      const irri_date = `${irri_split[0]}/${parseInt(irri_split[1])}/${parseInt(
        irri_split[2]
      )}`;
      for (let i = 0; i < date.length; i += 1) {
        if (date[i] === irri_date) {
          water["PRCP"][i] -=
            (irrigation_amount[idx] * 8.1071318210885) / 100.0;
        }
      }
    }
  }
  if (
    Object.keys(water).includes("IRRG") &&
    store.state.step2.IrrigationType === 0
  ) {
    const irrigation_data = store.state.step2.irrigation_date;
    const irrigation_amount = store.state.step2.irrigation_amount;
    for (let idx = 0; idx < irrigation_data.length; idx += 1) {
      const irri_split = irrigation_data[idx].split("-");
      const irri_date = `${irri_split[0]}/${parseInt(irri_split[1])}/${parseInt(
        irri_split[2]
      )}`;
      for (let i = 0; i < date.length; i += 1) {
        if (date[i] === irri_date && irrigation_amount[idx] !== -1) {
          water["IRRG"][i] +=
            (irrigation_amount[idx] * 8.1071318210885) / 100.0;
        }
      }
    }
  }
  const water_chart = MakeChartObj(water, WaterEchartId, date, water_y);

  // make multiple pesticide charts
  let pesticide_chart_list = [];
  let pesticide_data_list = {};
  Object.keys(pesticide).forEach((key) => {
    const msg = key.split("_");
    const dataKey = msg[1] + "_" + msg[2];
    if (dataKey in pesticide_data_list) {
      pesticide_data_list[dataKey][key] = pesticide[key];
    } else {
      pesticide_data_list[dataKey] = {};
      pesticide_data_list[dataKey][key] = pesticide[key];
    }
  });
  // console.log(pesticide_data_list);
  Object.keys(pesticide_data_list).forEach((key) => {
    const msg = key.split("_");
    pesticide_chart_list.push(
      MakeChartObj(
        pesticide_data_list[key],
        PesticideEchartId + msg[0],
        date,
        pesticide_y
      )
    );
  });

  const sediment_chart = MakeChartObj(
    sediment,
    SedimentEchartId,
    date,
    sediment_y
  );

  let concentration_chart_list = [];
  Object.keys(concentration).forEach((key) => {
    let conData = {};
    conData[key] = concentration[key];
    concentration_chart_list.push(
      MakeChartObj(
        conData,
        ConcentrationEchartId + key,
        date,
        concentration_y,
        benchmark[key]
      )
    );
  });

  // store width and height
  const echartsWidth = water_chart.getWidth();
  const echartsHeight = water_chart.getHeight();

  // compare concentration and benchmark
  let recommend_flag = compareConcentrationAndBenchmark(
    concentration,
    benchmark
  );
  console.log("recommend_flag:", recommend_flag);

  // return objs
  return {
    // obj of echart
    WaterEchart: water_chart,
    PesticideEchartList: pesticide_chart_list,
    SedimentEchart: sediment_chart,
    ConcentrationEchart: concentration_chart_list,
    // width and height
    width: echartsWidth,
    height: echartsHeight,
    recommend_flag: recommend_flag,
    overlayShow: false,
  };
}

/**
 *
 * @param {Int} BmpCnt count of bmp
 * @param {String} Harvest string of harvest : yyyy-mm-dd
 * @param {String} id bmp id
 * @returns echart obj oj comparison chart
 */
export function GetComparisonEchart(BmpCnt, Harvest, id) {
  return GetComparisonData(BmpCnt, Harvest).then((res) => {
    const Data = res.data.data;
    // get data
    let Comparison_runoff = Data.comparison_runoff;
    let Comparison_erosion = Data.comparison_erosion;
    let Comparison_volatilization = Data.comparison_volatilization;

    // VFSM 对比测试
    if (store.state.vfsm_list.length !== 0) {
      Comparison_runoff.push(74.6);
      Comparison_erosion.push(76.3);
      Comparison_volatilization.push(0);
    }

    let xAxisLabel = [];
    for (let i = 1; i <= BmpCnt; i += 1) {
      xAxisLabel.push("BMP " + i);
    }
    xAxisLabel.push("VFS");

    return ComparisonEchartMaker(
      id,
      Comparison_runoff,
      Comparison_erosion,
      Comparison_volatilization,
      xAxisLabel
    );
  });
}

function compareConcentrationAndBenchmark(concentration, benchmark) {
  // max value of concentration
  let max_concentration = -1;
  Object.keys(concentration).forEach((key) => {
    let tmp = Math.max(concentration[key]);
    if (tmp > max_concentration) {
      max_concentration = tmp;
    }
  });
  // max value of benchmark
  let max_bench = -1;
  Object.keys(benchmark).forEach((key) => {
    if (benchmark[key] > max_bench) {
      max_bench = benchmark[key];
    }
  });

  return max_bench >= max_concentration;
}
