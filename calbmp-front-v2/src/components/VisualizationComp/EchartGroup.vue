<template>
  <div class="EchartGroup">
    <b-overlay :show="overlayShow" rounded="sm">
      <b-tabs card active-nav-item-class="font-weight-bold">
        <!-- waterBalanceComponents -->
        <b-tab title="Water" no-body active>
          <custom-echart :echarId="waterEchartId"></custom-echart>
        </b-tab>

        <!-- sedimentComponentsOpt -->
        <b-tab title="Sediment" no-body @click="resizeEchart">
          <custom-echart :echarId="sedimentEchartId"></custom-echart>
        </b-tab>

        <!-- loadingComponentsOpt -->
        <div v-for="(item, index) in postData.pesticide_list" :key="index">
          <b-tab :title="'Loading:' + item" no-body @click="resizeEchart">
            <custom-echart :echarId="pesticideEchartId + item"></custom-echart>
          </b-tab>

          <b-tab
            :title="'Concentration:' + item"
            no-body
            @click="toConcentration"
          >
            <custom-echart
              :echarId="
                concentrationEchartId + item + '_' + (index + 1).toString()
              "
            ></custom-echart>
          </b-tab>
        </div>
      </b-tabs>
    </b-overlay>
  </div>
</template>

<script>
import CustomEchart from "./CustomEchart.vue";
import { GetCharts } from "@/server/visualization.js";

export default {
  name: "echart-group",
  components: {
    CustomEchart,
  },
  props: {
    // id of echarts
    waterEchartId: String,
    pesticideEchartId: String,
    sedimentEchartId: String,
    concentrationEchartId: String,
    // scenario type : baseline or bmp
    ScenarioType: {
      type: String,
      default: "baseline",
    },
    // bmp id
    BmpId: {
      type: String,
      default: "0",
    },
  },
  data() {
    return {
      // echart object
      waterEchart: null,
      pesticideEchartList: [],
      sedimentEchart: null,
      concentrationEchart: null,
      // echart width and height
      echartWidth: Number,
      echartHeight: Number,

      // if benchmark higher than concentration
      recommend_flag: false,

      // form data
      postData: {
        // selected option of per type
        water: [],
        pesticide: [],
        sediment: [],
        concentration: [],
        // scenario type and bmp_id control bmp scenario
        scenario_type: this.ScenarioType,
        bmp_id: this.BmpId,
        // for get benchmark
        pesticide_list: this.$store.state.step2.Pesticide,
        last_harvest: this.$store.state.last_harvest,
        first_emergence: this.$store.state.first_emergence,
        field_size: parseFloat(this.$store.state.step1.FieldSize),
      },

      // overlay
      overlayShow: false,
    };
  },
  methods: {
    async getEcharts() {
      this.overlayShow = true;

      // get charts by new func
      // get echart object and width and height
      /**
       * 获取图表对象以及默认图表的长宽
       * @type {{recommend_flag: boolean, width: *, SedimentEchart: *, PesticideEchartList: [], overlayShow: boolean, WaterEchart: *, ConcentrationEchart: [], height: *}}
       */
      const EchartInfo = await GetCharts(
        this.postData,
        this.waterEchartId,
        this.pesticideEchartId,
        this.sedimentEchartId,
        this.concentrationEchartId
      );

      // get echart object
      /**
       * 分别获取图表对象
       */
      this.waterEchart = EchartInfo.WaterEchart;
      this.pesticideEchartList = EchartInfo.PesticideEchartList; // list
      this.sedimentEchart = EchartInfo.SedimentEchart;
      this.concentrationEchart = EchartInfo.ConcentrationEchart; // list

      // get width and height
      this.echartWidth = EchartInfo.width;
      this.echartHeight = EchartInfo.height;

      // 取消loading
      this.overlayShow = EchartInfo.overlayShow;

      // 是否弹出推荐弹窗
      this.recommend_flag = EchartInfo.recommend_flag;
    },

    toConcentration() {
      this.resizeEchart();
      if (this.recommend_flag) {
        this.$bvToast.toast(
          "The pesticide concentration in runoff (blue line) is the on-site concentration, not the value in a receiving water body. If the concentration of a pesticide is lower than a benchmark value (green line), there is no need of BMP on farm. If the pesticide concentration is much higher than a benchmark value, BMP to help with pesticide mitigation is highly recommended.",
          {
            title: "Note",
            toaster: "b-toaster-top-left",
            variant: "danger",
            solid: true,
          }
        );

        this.$bvToast.toast("Please click on [Try BMP] button on the left.", {
          title: "Note",
          toaster: "b-toaster-top-left",
          variant: "info",
          solid: true,
        });
      }
    },

    resizeEchart() {
      if (this.echartWidth != null) {
        // console.log("resizing");
        // resize loading echart
        if (this.pesticideEchartList !== []) {
          this.pesticideEchartList.forEach((chart) => {
            chart.resize({
              width: this.echartWidth,
              height: this.echartHeight,
            });
          });
        }

        // resize sediment echart
        if (this.sedimentEchart !== null) {
          this.sedimentEchart.resize({
            width: this.echartWidth,
            height: this.echartHeight,
          });
        }

        if (this.concentrationEchart !== []) {
          this.concentrationEchart.forEach((chart) => {
            chart.resize({
              width: this.echartWidth,
              height: this.echartHeight,
            });
          });
        }
      }
    },
  },
};
</script>
