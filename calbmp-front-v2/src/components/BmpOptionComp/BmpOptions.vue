<template>
  <div class="bmpOptions">
    <!-- options part -->
    <b-sidebar
      id="bmpOptions"
      title="BMP Options"
      shadow="lg"
      sidebar-class="border-right border-primary"
      backdrop-variant="dark"
      backdrop
    >
      <b-overlay :show="overlayShow">
        <div class="px-3 py-2">
          <b-card class="mt-5">
            <!-- Non-structural BMPs -->
            <!-- 非结构化 BMPs -->
            <b-form-group label="Non-structural BMPs">
              <div
                class="checkBoxList"
                v-for="item in bmpOptions"
                :key="item.modalId"
              >
                <!-- @change:选项改变的时候添加到bmp_opts列表当中 -->
                <b-form-checkbox
                  v-model="item.status"
                  @change="onCheckBoxChange(item.modalId, item.status)"
                  @click="removeModalId(item.modalId)"
                >
                  {{ item.text }}
                </b-form-checkbox>
              </div>
            </b-form-group>

            <hr />

            <!-- Structural BMPs -->
            <!-- 结构化 BMPs -->
            <b-form-group label="Structural BMPs">
              <div
                class="checkBoxList"
                v-for="item in structuralBmpOptions"
                :key="item.modalId"
              >
                <b-form-checkbox
                  v-model="item.status"
                  @change="onCheckBoxChange(item.modalId, item.status)"
                  @click="removeModalId(item.modalId)"
                >
                  {{ item.text }}
                </b-form-checkbox>
              </div>
            </b-form-group>
          </b-card>

          <!-- BMP buttons -->
          <b-card class="mt-3" border-variant="light">
            <b-row>
              <b-col class="mx-auto">
                <b-button variant="success" block @click="tryBmp">
                  Submit
                </b-button>
              </b-col>
            </b-row>
          </b-card>
        </div>
      </b-overlay>

      <!-- Close Button -->
      <template #footer="{ hide }">
        <div class="d-flex bg-dark text-light align-items-center px-3 py-2">
          <!-- <strong class="mr-auto">Footer</strong> -->
          <b-button variant="danger" size="sm" class="ml-auto" @click="hide">
            Close
          </b-button>
        </div>
      </template>
    </b-sidebar>

    <!-- Modal part -->
    <!-- 每个BMP的弹窗选项 -->
    <div class="ModalPar">
      <!-- Modal: contour farming -->
      <!--    have no value, just add modal id -->
      <contour-farming
        @modal-cancel="removeModalId('contourFarming')"
      ></contour-farming>

      <!-- Modal: strip cropping -->
      <strip-cropping
        :crops="crops"
        v-model="non_structural_form.strip_cropping"
        @modal-cancel="removeModalId('stripCropping')"
      ></strip-cropping>

      <!-- Modal: parallel terracing -->
      <parallel-terracing
        v-model="non_structural_form.parallel_terracing"
        @modal-cancel="removeModalId('parallelTerracing')"
      ></parallel-terracing>

      <!-- Modal: cover crops -->
      <cover-crops
        :crops="crops"
        v-model="non_structural_form.cover_crops"
        @modal-cancel="removeModalId('coverCrops')"
      ></cover-crops>

      <!-- Modal: pesticide application reduction -->
      <pesticideApp-reduction
        v-model="non_structural_form.pesticide_app_reduction"
        @modal-cancel="removeModalId('pesticideAppReduction')"
      ></pesticideApp-reduction>

      <!-- Modal: pesticide application timing -->
      <pesticideApp-timing
        @modal-cancel="removeModalId('pesticideAppTiming')"
      ></pesticideApp-timing>

      <!-- Modal: crop rotation -->
      <crop-rotation
        :crops="crops"
        v-model="non_structural_form.crop_rotation_and_residue_management"
        @modal-cancel="removeModalId('cropRotation')"
      ></crop-rotation>

      <!-- Modal: VFS模型-->
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
import { BmpOptionsList, StructuralBmpOptionList } from "@/variable/bmpVar";
import { RunVfsmModel } from "@/server/vfsmServer";
// components
import ContourFarming from "./ContourFarmingComp.vue";
import StripCropping from "./StripCroppingComp.vue";
import ParallelTerracing from "./ParallelTerracingComp.vue";
import CoverCrops from "./CoverCropsComp.vue";
import CropRotation from "./CropRotationComp.vue";
import PesticideAppReduction from "./PesticideAppReductionComp.vue";
import PesticideAppTiming from "./PesticideAppTimingComp.vue";
import VfsmBasic from "@/components/BmpOptionComp/VfsmBasic";

export default {
  name: "bmp-options",
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
  data() {
    return {
      // 在外部定义的变量
      bmpOptions: BmpOptionsList, // 非结构化BMP选项
      structuralBmpOptions: StructuralBmpOptionList, // 结构化BMP选项
      // 非结构化BMP 表单数据
      non_structural_form: {
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
  methods: {
    /**
     * 监测复选框是否选中
     * @param modalId
     * @param status
     */
    onCheckBoxChange(modalId, status) {
      // 1. show modal
      this.showBmpModal(modalId, status);

      // 2. add or remove modal id
      //    status=true  : add modal id
      //    status=false : remove modal id
      if (status === true) {
        if (modalId === "vfsm_basic_modal") {
          // 表明选择的是结构化BMP
          this.non_structural_form.vfsm_opts.push(modalId);
        } else {
          this.non_structural_form.bmp_opts.push(modalId);
          console.log("bmp_opts_on change:", this.non_structural_form.bmp_opts);
        }
      }
    },

    /**
     * 显示选中的BMP弹窗
     * @param modalId
     * @param status
     */
    showBmpModal(modalId, status) {
      if (status === true) {
        this.$bvModal.show(modalId);
      }
    },

    /**
     * 移除未选的BMP选项
     * @param modalId
     */
    removeModalId(modalId) {
      // 先从form表单中删掉
      let temp = this.non_structural_form.bmp_opts;
      for (let i = 0; i < temp.length; i += 1) {
        if (temp[i] === modalId) {
          this.non_structural_form.bmp_opts.splice(i, 1);
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
      temp = this.non_structural_form.vfsm_opts;
      for (let i = 0; i < temp.length; i += 1) {
        if (temp[i].modalId === modalId) {
          this.non_structural_form.bmp_opts.splice(i, 1);
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

    /**
     * 运行BMP，这是核心
     */
    tryBmp() {
      // 非结构化BMP的数量
      let non_structural_bmp_len = this.non_structural_form.bmp_opts.length;
      // 结构化BMP的数量
      let structural_bmp_len = this.non_structural_form.vfsm_opts.length;

      // 判断非结构化BMP的数量
      if (non_structural_bmp_len !== 0) {
        this.overlayShow = true; // 加载动画开始

        // 非结构化BMP的运行函数
        BmpScenario(this.non_structural_form).then((res) => {
          console.log(res);

          // 将最新的BMP ID加入到浏览器内存
          const echartList = this.$store.state.echartList;
          echartList.push({
            bmp_id: `${echartList.length + 1}`,
          });
          this.$store.commit("SetEchartList", echartList);

          // 将BMP的名字添加进内存
          let bmp_options_list = this.$store.state.bmp_options_list;
          bmp_options_list.push(this.non_structural_form.bmp_opts);
          this.$store.commit("SetBmpOptionsList", bmp_options_list);

          this.overlayShow = false;

          location.reload();
        });
      }

      // 判断结构化BMP的数量
      if (structural_bmp_len !== 0) {
        this.overlayShow = true;
        // 将vfsm添加到浏览器内存
        let vfsm_list = this.$store.state.vfsm_list;
        vfsm_list.push({
          vfsm_id: `${vfsm_list.length + 1}`,
        });
        this.$store.commit("SetVfsmList", vfsm_list);

        // 将vfsm配置存到浏览器内存
        let vfsm_option_list = this.$store.state.vfsm_options_list;
        vfsm_option_list.push(this.non_structural_form.vfsm_opts);
        this.$store.commit("SetVfsmOptionList", vfsm_option_list);

        this.overlayShow = false;
        // run vfsm model
        this.$emit("get_vfsm_progress");
        RunVfsmModel(this.vfsm_param).then(() => {
          location.reload();
        });
      }

      // 重置表单
      this.non_structural_form = {
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
      };

      // end of 'try bmp'
    },
  },

  async mounted() {
    this.crops = await GetCrops();
  },
};
</script>
