<template>
  <div class="BMPComparison">
    <b-tabs card active-nav-item-class="font-weight-bold">
      <b-tab title="Effectiveness" no-body active>
        <custom-echart echarId="comparison"></custom-echart>
      </b-tab>

      <!-- <b-tab title="Price" no-body @click="resizeEchart">
        <custom-echart echarId="Price"></custom-echart>
      </b-tab> -->
    </b-tabs>
  </div>
</template>

<script>
import CustomEchart from "./CustomEchart.vue";
import { GetComparisonEchart } from "@/server/visualization";

export default {
  name: "bmp-comparison",
  components: {
    CustomEchart,
  },
  data() {
    return {
      comparisonEchart: null,
      costEchart: null,

      echartWidth: null,
      echartHeight: null,
    };
  },

  methods: {
    async getEcharts() {
      const BmpCnt = this.$store.state.echartList.length;
      const Harvest = this.$store.state.step2.HarvestDate;
      // console.log(this.$store.state);
      const comparisonEcharts = await GetComparisonEchart(
        BmpCnt,
        Harvest,
        "comparison"
      );

      this.comparisonEchart = comparisonEcharts;
      this.costEchart = comparisonEcharts;
    },

    resizeEchart() {
      if (this.echartWidth != null) {
        console.log("resizing");
        this.pesticideEchart.resize({
          width: this.echartWidth,
          height: this.echartHeight,
        });
      }
    },
  },
};
</script>
