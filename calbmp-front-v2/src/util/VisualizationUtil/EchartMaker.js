import { MyEcharts } from "@/server/echartsServer";

export function MakeChartObj(
  data_obj,
  chart_id,
  date,
  y_axis_name,
  benchmark = false
) {
  // make legend
  let legend;
  const legend_keys = Object.keys(data_obj);
  if (!benchmark) {
    legend = LegendMapper(legend_keys);
  } else {
    legend = legend_keys;
  }
  // make series
  let series = [];
  legend.forEach((key, index) => {
    let chart_type = "line";
    if (key.includes("Irr") || key.includes("Pre")) {
      chart_type = "bar";
    }
    series.push({
      name: key,
      type: chart_type,
      data: data_obj[legend_keys[index]],
    });
  });

  // Special cases
  if (benchmark !== false) {
    let extra_legend = [];
    legend.forEach((key) => {
      const tmp_name = "US EPA aquatic life benchmark:" + key;
      extra_legend.push(tmp_name);
      const benchmark_data = new Array(data_obj[key].length).fill(
        parseFloat(benchmark)
      );
      // add 2 series
      series.push({
        name: tmp_name,
        type: "line",
        data: benchmark_data,
        symbolSize: 0,
      });
    });
    legend.push(...extra_legend);
  }

  // make echart obj
  return MyEcharts(chart_id, legend, date, series, y_axis_name);
}

function LegendMapper(legend_keys) {
  // legend map
  const mp = {
    IRRG: "Irrigation",
    PRCP: "Precipitation",
    RUNF: "Runoff",
    ESLS: "Eroded solids",
    EFLX: "Loading in erosion",
    RFLX: "Loading in runoff",
    VFLX: "Loading in volatilization",
  };
  // result
  let res = [];
  // console.log("legend mp:", Object.keys(mp));
  legend_keys = legend_keys.sort(function (a, b) {
    let a_val = -1;
    let b_val = -1;
    if (a.includes("RFLX")) {
      a_val = 2;
    } else if (a.includes("EFLX")) {
      a_val = 1;
    } else {
      a_val = 0;
    }

    if (b.includes("RFLX")) {
      b_val = 2;
    } else if (b.includes("EFLX")) {
      b_val = 1;
    } else {
      b_val = 0;
    }

    if (a === b || a === -1 || b === -1) {
      return 0;
    }

    return a_val - b_val > 0 ? -1 : 1;
  });
  for (let i = 0; i < legend_keys.length; i++) {
    const key_now = legend_keys[i];
    Object.keys(mp).forEach((key) => {
      if (key_now.includes(key)) {
        if (key_now.includes("TCUM")) {
          res.push(mp[key] + " (cum)");
        } else if (key_now.includes("TSER")) {
          res.push(mp[key] + " (daily)");
        } else {
          res.push(mp[key]);
        }
      }
    });
  }
  // console.log("legend res:", res);
  return res;
}
