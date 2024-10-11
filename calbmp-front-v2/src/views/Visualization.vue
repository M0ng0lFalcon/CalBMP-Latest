<template>
  <div class="Visualization">
    <b-row>
      <!-- left side -->
      <b-col md="8" lg="3">
        <!-- Baseline Options -->
        <b-card class="mt-5">
          <components-opts></components-opts>
        </b-card>

        <!-- Button sets -->
        <b-card class="mt-2">
          <!-- submit button -->
          <b-row>
            <b-col class="mx-auto">
              <b-button variant="primary" @click="submitFun" block>
                Submit
              </b-button>
            </b-col>
          </b-row>

          <!-- save project button -->
          <b-row class="mt-2">
            <b-col class="mx-auto">
              <b-button variant="success" @click="saveProject" block>
                Save Project
              </b-button>
            </b-col>
          </b-row>

          <!-- BMP settings -->
          <b-row class="mt-2">
            <b-col class="mx-auto">
              <b-button v-b-toggle.bmpOptions variant="outline-danger" block>
                {{
                  $store.state.echartList.length === 0 &&
                  $store.state.vfsm_list.length === 0
                    ? "Try BMP"
                    : "Try Another BMP"
                }}
              </b-button>
            </b-col>
            <!-- bmp options shown as sidebar -->
            <bmp-options @get_vfsm_progress="get_vfsm_progress"></bmp-options>
          </b-row>

          <b-row
            class="mt-2"
            v-if="
              $store.state.echartList.length !== 0 ||
              $store.state.vfsm_list.length !== 0
            "
          >
            <b-col class="mx-auto">
              <b-button variant="danger" @click="clearBmp" block>
                Clear BMP
              </b-button>
            </b-col>
          </b-row>
        </b-card>
      </b-col>

      <!-- right side: visualization part -->
      <b-col md="4" lg="9">
        <b-card no-body class="mt-5">
          <b-tabs v-model="tab_index" card end>
            <!-- baseline side -->
            <!-- baseline side -->
            <!-- baseline side -->
            <b-tab no-body title="Baseline" @click="changeRef('baseline')">
              <echart-group
                waterEchartId="baselineWater"
                pesticideEchartId="baselinePesticide"
                sedimentEchartId="baselineSediment"
                concentrationEchartId="baselineConcentration"
                ref="baseline"
              ></echart-group>
            </b-tab>

            <!-- summary side -->
            <!-- summary side -->
            <!-- summary side -->
            <b-tab no-body title="Summary">
              <baseline-table></baseline-table>
            </b-tab>

            <!-- bmp side -->
            <!-- bmp side -->
            <!-- bmp side -->
            <b-tab
              v-for="(item, index) in echartList"
              :key="item.bmp_id"
              :title-link-class="bmpTabClass(index)"
              @click="changeRef(`bmp_${item.bmp_id}`)"
              no-body
            >
              <template #title>
                BMP {{ item.bmp_id }}
                <b-icon
                  v-b-popover.hover.top="'Recommended BMP: ' + bmp_opts[index]"
                  title="Tips"
                  icon="info-circle-fill"
                  variant="primary"
                  v-if="index + 1 === best_bmp"
                ></b-icon>
              </template>
              <echart-group
                :waterEchartId="`bmp_water_${item.bmp_id}`"
                :pesticideEchartId="`bmp_pesticide_${item.bmp_id}`"
                :sedimentEchartId="`bmp_sediment_${item.bmp_id}`"
                :concentrationEchartId="`bmp_concentration_${item.bmp_id}`"
                :ref="`bmp_${item.bmp_id}`"
                ScenarioType="bmp"
                :BmpId="item.bmp_id"
              ></echart-group>
            </b-tab>

            <!-- vfs side -->
            <!-- vfs side -->
            <!-- vfs side -->
            <b-tab
              v-for="(item, index) in vfsm_list"
              :key="index"
              @click="changeRef(`vfsm_${item.vfsm_id}`)"
              no-body
            >
              <template #title>
                <!-- vfs 进度条 -->
                <b-progress
                  v-if="!vfsm_status"
                  :value="vfsm_progress"
                  max="1"
                  show-progress
                  animated
                ></b-progress>
                VFS Model
              </template>
              <echart-group
                :waterEchartId="`vfsm_water_${item.vfsm_id}`"
                :pesticideEchartId="`vfsm_pesticide_${item.vfsm_id}`"
                :sedimentEchartId="`vfsm_sediment_${item.vfsm_id}`"
                :concentrationEchartId="`vfsm_concentration_${item.vfsm_id}`"
                :ref="`vfsm_${item.vfsm_id}`"
                ScenarioType="vfsm"
                :BmpId="item.vfsm_id"
              ></echart-group>
            </b-tab>

            <!-- bmp summary -->
            <!-- bmp summary -->
            <!-- bmp summary -->
            <b-tab
              title-link-class="text-dark"
              no-body
              title="BMP Summary"
              v-if="echartList.length !== 0 || vfsm_list.length !== 0"
            >
              <bmp-table></bmp-table>
            </b-tab>

            <!-- bmp comparison -->
            <!-- bmp comparison -->
            <!-- bmp comparison -->
            <b-tab
              no-body
              title="BMP Comparison"
              title-link-class="text-dark"
              v-if="echartList.length !== 0 || vfsm_list.length !== 0"
              @click="getComparisonEchart"
            >
              <bmp-comparison ref="pc"></bmp-comparison>
            </b-tab>

            <!-- download -->
            <!-- download -->
            <!-- download -->
            <b-tab no-body title="Download" title-item-class="text-dark">
              <download></download>
            </b-tab>
          </b-tabs>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>

<script>
import BmpOptions from "@/components/BmpOptionComp/BmpOptions.vue";
import EchartGroup from "@/components/VisualizationComp/EchartGroup.vue";
import ComponentsOpts from "@/components/VisualizationComp/ComponentsOpts.vue";
import BmpComparison from "@/components/VisualizationComp/BMPComparison.vue";
import BaselineTable from "../components/VisualizationComp/BaselineSummaryTable.vue";
import BmpTable from "../components/VisualizationComp/BMPSummaryTable.vue";
import Download from "@/components/VisualizationComp/Download";
import { AddHistory } from "@/server/historyServer";
import { GetBmpSummaryData } from "@/server/resultServer";
import { GetVfsmProgress } from "@/server/vfsmServer";

export default {
  components: {
    BmpOptions,
    ComponentsOpts,
    EchartGroup,
    BmpComparison,
    BaselineTable,
    BmpTable,
    Download,
  },
  data() {
    return {
      tab_index: 0,

      // bmp list
      echartList: this.$store.state.echartList,
      bmp_opts: this.$store.state.bmp_options_list,

      // vfs list
      vfsm_list: this.$store.state.vfsm_list,
      vfsm_options_list: this.$store.state.vfsm_options_list,
      vfsm_status: false,
      vfsm_progress: 0,

      // current ref
      curRef: "baseline",
      best_bmp: 0,
    };
  },

  methods: {
    get_vfsm_progress() {
      console.log("get progress");
      let get_progress_mission = setInterval(async () => {
        const res = await GetVfsmProgress();
        console.log(res);
        this.vfsm_progress = res.progress;
        if (res.progress === undefined) {
          this.vfsm_status = true;
          this.vfsm_progress = 1;
          this.$store.commit("SetVfsmMinId", res.min);
          this.$store.commit("SetVfsmMaxId", res.max);
          clearInterval(get_progress_mission);
        }
      }, 10000);
    },
    async submitFun() {
      this.$bvToast.toast(
          'Since a default input of initial soil mositure content assumed in CalBMP, modeling uncertainty of the first 1-2 irrigation events are expected.',
          {
            title: "Note",
            toaster: "b-toaster-top-left",
            variant: "success",
            solid: true,
          }
      );
      // get specific visualization
      if (this.curRef !== "baseline") {
        await this.$refs[this.curRef][0].getEcharts();
      } else {
        this.$refs[this.curRef].getEcharts();
      }
    },

    changeRef(curRef) {
      this.curRef = curRef;
      console.log("[!] Changing ref:", this.curRef);
    },

    clearBmp() {
      this.$store.commit("ClearBmp");
      this.$store.commit("ClearBmpOptionsList");
      this.$store.commit("ClearVfsm");
      this.$store.commit("ClearVfsmOptionList");
      location.reload();
    },

    getComparisonEchart() {
      this.$refs["pc"].getEcharts();
    },

    saveProject() {
      // add history
      const CreatedTime = this.$store.state.created_time;
      AddHistory(CreatedTime).then((res) => {
        console.log(res);
        this.$bvToast.toast("Add History", {
          title: "Add History",
          variant: "success",
          solid: true,
        });
      });
    },

    bmpTabClass(idx) {
      if (idx + 1 === this.best_bmp) {
        return "text-danger";
      } else if (idx + 2 === this.tab_index) {
        return "text-success bg-white";
      } else {
        return "text-success";
      }
    },
  },

  async mounted() {
    const bmp_cnt = this.$store.state.echartList.length;
    const harvest = this.$store.state.last_harvest;
    let comparison_array = await GetBmpSummaryData(bmp_cnt, harvest);
    comparison_array = comparison_array.sort(function (a, b) {
      const a_weight =
        parseFloat(a.pesticide_reduction_in_Runoff) +
        parseFloat(a.pesticide_reduction_in_Erosion);
      const b_weight =
        parseFloat(b.pesticide_reduction_in_Runoff) +
        parseFloat(b.pesticide_reduction_in_Erosion);
      return b_weight - a_weight;
    });
    // console.log("comparison array:", comparison_array);
    this.best_bmp = comparison_array[0].BMP_Scenario;

    // 获取vfs的进度
    this.get_vfsm_progress();
  },
};
</script>
