import Vue from "vue";

// id         : water component || pesticide component
// legendData : title
// date       : date info
// seriesData : shown data
// yAxisName  : y axis names
export function MyEcharts(id, legendData, date, seriesData, yAxisName) {
  const myChart = Vue.prototype.$echarts.init(document.getElementById(id));
  const option = {
    tooltip: {},
    toolbox: {
      feature: {
        dataZoom: {
          yAxisIndex: "none",
        },
        restore: {},
        saveAsImage: {},
      },
    },
    legend: {
      top: -5,
      width: "60%",
      data: legendData,
    },
    xAxis: {
      name: "Time",
      type: "category",
      data: date,
      axisPointer: {
        type: "shadow",
      },
    },
    yAxis: {
      type: "value",
      name: yAxisName,
      minorTick: {
        show: true,
      },
      axisLabel: {
        formatter(value) {
          if (value === 0) {
            return "0";
          }
          if (`${value}`.indexOf("e") > 0) {
            return `${value}`.replace(/e/, "E");
          }
          const res = value.toString();
          let numN1 = 0;
          let numN2 = 1;
          let num1 = 0;
          let num2 = 0;
          let t1 = 1;
          for (let k = 0; k < res.length; k += 1) {
            if (res[k] === ".") {
              t1 = 0;
            }
            if (t1) {
              num1 += 1;
            } else {
              num2 += 1;
            }
          }
          if (Math.abs(value) < 1) {
            for (let i = 2; i < res.length; i += 1) {
              if (res[i] === "0" && res[i] !== ".") {
                numN2 += 1;
              } else {
                break;
              }
            }
            let v = parseFloat(value);
            v *= 10 ** numN2;
            v = v.toFixed(1);
            return `${v.toString()}E-${numN2}`;
          }
          if (num1 > 1) {
            numN1 = num1 - 1;
            let v = parseFloat(value);
            v /= 10 ** numN1;
            if (num2 > 1) {
              v = v.toFixed(1);
            }
            return `${v.toString()}E${numN1}`;
          }
          return value;
        },
      },
    },
    dataZoom: [
      {
        show: true,
        start: 0,
        end: 100,
      },
      {
        type: "inside",
        start: 0,
        end: 100,
      },
    ],
    series: seriesData,
  };

  myChart.setOption(option);

  return myChart;
}

export function ComparisonEchartMaker(
  id,
  Comparison_runoff,
  Comparison_erosion,
  Comparison_volatilization,
  xAxisLabel
) {
  var app = {};

  // init echart object
  var comparisonEchart = Vue.prototype.$echarts.init(
    document.getElementById(id)
  );
  var option;

  app.config = {
    rotate: 90,
    align: "left",
    verticalAlign: "middle",
    position: "insideBottom",
    distance: 15,
    onChange: function () {
      var labelOption = {
        normal: {
          rotate: app.config.rotate,
          align: app.config.align,
          verticalAlign: app.config.verticalAlign,
          position: app.config.position,
          distance: app.config.distance,
        },
      };
      comparisonEchart.setOption({
        series: [
          {
            label: labelOption,
          },
          {
            label: labelOption,
          },
          {
            label: labelOption,
          },
        ],
      });
    },
  };

  var labelOption = {
    show: true,
    position: app.config.position,
    distance: app.config.distance,
    align: app.config.align,
    verticalAlign: app.config.verticalAlign,
    rotate: app.config.rotate,
    formatter: "{a}",
    color: "#240006",
    fontSize: 16,
    fontWeight: "bolder",
    rich: {
      name: {},
    },
  };

  // init options
  option = {
    tooltip: {
      trigger: "axis",
      axisPointer: {
        type: "shadow",
      },
    },
    legend: {
      data: ["Runoff", "Erosion", "Volatilization"],
    },
    xAxis: [
      {
        type: "category",
        axisTick: { show: false },
        data: xAxisLabel,
      },
    ],
    yAxis: [
      {
        type: "value",
        name: "Pesticide Reduction(%)",
      },
    ],
    series: [
      {
        name: "Runoff",
        type: "bar",
        barGap: 0,
        label: labelOption,
        emphasis: {
          focus: "series",
        },
        data: Comparison_runoff,
      },
      {
        name: "Erosion",
        type: "bar",
        label: labelOption,
        emphasis: {
          focus: "series",
        },
        data: Comparison_erosion,
      },
      {
        name: "Volatilization",
        type: "bar",
        label: labelOption,
        emphasis: {
          focus: "series",
        },
        data: Comparison_volatilization,
      },
    ],
  };

  comparisonEchart.setOption(option);

  return comparisonEchart;
}
