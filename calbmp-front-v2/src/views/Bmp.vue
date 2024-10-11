<template>
  <div class="bmp_page">
    <b-card class="mt-5 mx-auto" style="max-width: 500px">
      <div class="px-3 py-2">
        <!-- Non-structural BMPs -->
        <b-form-group label="Non-structural BMPs">
          <div
            class="checkBoxList"
            v-for="item in bmpOptions"
            :key="item.modalId"
          >
            <b-form-checkbox
              v-model="item.status"
              @change="onCheckBoxChange(item.modalId, item.status)"
            >
              {{ item.text }}
            </b-form-checkbox>
          </div>
        </b-form-group>

        <hr />

        <!-- Structural BMPs -->
        <b-form-group label="Structural BMPs">
          <div
            class="checkBoxList"
            v-for="item in structuralBmpOptions"
            :key="item.modalId"
          >
            <b-form-checkbox
              v-model="item.status"
              @change="onCheckBoxChange(item.modalId, item.status)"
            >
              {{ item.text }}
            </b-form-checkbox>
          </div>
        </b-form-group>
      </div>

      <!-- BMP buttons -->
      <template #footer>
        <b-button variant="success" @click="submit_test"> Submit </b-button>
      </template>
    </b-card>

    <!-- Modal part -->
    <div class="ModalPar">
      <!-- Modal: contour farming -->
      <!--    have no value, just add modal id -->
      <contour-farming
        @modal-cancel="removeModalId('contourFarming')"
      ></contour-farming>

      <!-- Modal: strip cropping -->
      <strip-cropping
        :crops="crops"
        v-model="form.strip_cropping"
        @modal-cancel="removeModalId('stripCropping')"
      ></strip-cropping>

      <!-- Modal: parallel terracing -->
      <parallel-terracing
        v-model="form.parallel_terracing"
        @modal-cancel="removeModalId('parallelTerracing')"
      ></parallel-terracing>

      <!-- Modal: cover crops -->
      <cover-crops
        :crops="crops"
        v-model="form.cover_crops"
        @modal-cancel="removeModalId('coverCrops')"
      ></cover-crops>

      <!-- Modal: pesticide application reduction -->
      <pesticideApp-reduction
        v-model="form.pesticide_app_reduction"
        @modal-cancel="removeModalId('pesticideAppReduction')"
      ></pesticideApp-reduction>

      <!-- Modal: pesticide application timing -->
      <pesticideApp-timing
        @modal-cancel="removeModalId('pesticideAppTiming')"
      ></pesticideApp-timing>

      <!-- Modal: crop rotation -->
      <crop-rotation
        :crops="crops"
        v-model="form.crop_rotation_and_residue_management"
        @modal-cancel="removeModalId('cropRotation')"
      ></crop-rotation>

      <vfsm-basic
        v-model="vfsm_param"
        @modal-cancel="removeModalId('vfsm_basic_modal')"
      ></vfsm-basic>
    </div>
  </div>
</template>

<script>
// functions and variables
import { GetCrops } from "@/server/basicDataServer";
import { BmpScenario } from "@/server/bmpServer";
import { RunVfsmModel } from "@/server/vfsmServer";
import { BmpOptionsList, StructuralBmpOptionList } from "@/variable/bmpVar"; // 变量
// components
import ContourFarming from "../components/BmpOptionComp/ContourFarmingComp.vue";
import StripCropping from "../components/BmpOptionComp/StripCroppingComp.vue";
import ParallelTerracing from "../components/BmpOptionComp/ParallelTerracingComp.vue";
import CoverCrops from "../components/BmpOptionComp/CoverCropsComp.vue";
import CropRotation from "../components/BmpOptionComp/CropRotationComp.vue";
import PesticideAppReduction from "../components/BmpOptionComp/PesticideAppReductionComp.vue";
import PesticideAppTiming from "../components/BmpOptionComp/PesticideAppTimingComp.vue";
import VfsmBasic from "@/components/BmpOptionComp/VfsmBasic";

export default {
  name: "Bmp",
  // ------------------------------ component ------------------------------
  components: {
    // PesticideApplication,
    ContourFarming,
    StripCropping,
    ParallelTerracing,
    CoverCrops,
    CropRotation,
    PesticideAppReduction,
    PesticideAppTiming,
    VfsmBasic,
  },
  // ------------------------------ data ------------------------------
  data() {
    return {
      // imported variable
      bmpOptions: BmpOptionsList,
      structuralBmpOptions: StructuralBmpOptionList,
      // form data
      form: {
        bmp_id: `${this.$store.state.echartList.length + 1}`,
        // basic data
        bmp_opts: [],
        vfsm_opts: [],

        // optional data
        step_1_params: this.$store.state.step1,
        step_2_params: this.$store.state.step2,
        pesticide_app_reduction: {},
        strip_cropping: {},
        parallel_terracing: {},
        cover_crops: {},
        crop_rotation_and_residue_management: {},
      },
      vfsm_param: {},

      // from back end
      crops: null,

      // overlay
      overlayShow: false,
    };
  },
  // ------------------------------ methods ------------------------------
  methods: {
    onCheckBoxChange(modalId, status) {
      // 1. show modal
      this.showBmpModal(modalId, status);

      // 2. add or remove modal id
      //    status=true  : add modal id
      //    status=false : remove modal id
      if (status === true) {
        if (modalId === "vfsm_basic_modal") {
          this.form.vfsm_opts.push(modalId);
        } else {
          this.form.bmp_opts.push(modalId);
        }
      }
    },
    showBmpModal(modalId, status) {
      if (status === true) {
        this.$bvModal.show(modalId);
      }
    },
    removeModalId(modalId) {
      // 先从form表单中删掉
      let temp = this.form.bmp_opts;
      for (let i = 0; i < temp.length; i += 1) {
        if (temp[i] === modalId) {
          this.form.bmp_opts.splice(i, 1);
          break;
        }
      }

      // 常规bmp中删掉
      temp = this.bmpOptions;
      for (let i = 0; i < temp.length; i += 1) {
        if (temp[i].modalId === modalId) {
          temp[i].status = false;
          break;
        }
      }

      // vfsm表单中删掉
      temp = this.form.vfsm_opts;
      for (let i = 0; i < temp.length; i += 1) {
        if (temp[i].modalId === modalId) {
          this.form.bmp_opts.splice(i, 1);
          break;
        }
      }

      // 结构化bmp中删掉
      temp = this.structuralBmpOptions;
      for (let i = 0; i < temp.length; i += 1) {
        if (temp[i].modalId === modalId) {
          temp[i].status = false;
          break;
        }
      }
    },

    submit_test() {
      console.log(this.bmpOptions);
      console.log(this.structuralBmpOptions);
      console.log(this.form.bmp_opts);
      console.log(this.form.vfsm_opts);
    },

    // bmp submit fun
    tryBmp() {
      // set this variable to get bmp data
      if (this.form.bmp_opts.length !== 0) {
        this.overlayShow = true;
        BmpScenario(this.form).then((res) => {
          console.log(res);

          const echartList = this.$store.state.echartList;
          echartList.push({
            bmp_id: `${echartList.length + 1}`,
          });
          this.$store.commit("SetEchartList", echartList);

          let bmp_options_list = this.$store.state.bmp_options_list;
          bmp_options_list.push(this.form.bmp_opts);
          this.$store.commit("SetBmpOptionsList", bmp_options_list);

          this.overlayShow = false;

          location.reload();
        });
      }
      if (this.form.vfsm_opts.length !== 0) {
        this.overlayShow = true;
        RunVfsmModel(this.vfsm_param).then(() => {
          const echartList = this.$store.state.echartList;
          echartList.push({
            bmp_id: `${echartList.length + 1}`,
          });
          this.$store.commit("SetEchartList", echartList);

          let bmp_options_list = this.$store.state.bmp_options_list;
          bmp_options_list.push(this.form.vfsm_opts);
          this.$store.commit("SetBmpOptionsList", bmp_options_list);
          this.overlayShow = false;
          location.reload();
        });
      }
    },
  },
  // ------------------------------ mounted ------------------------------
  async mounted() {
    this.crops = await GetCrops();
  },
};
</script>

<style scoped></style>
