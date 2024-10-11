<template>
  <div class="CropRotationComp">
    <b-modal
      id="cropRotation"
      centered
      title="Crop Rotation"
      size="lg"
      @ok="onOk"
      @cancel="onCancel"
      @close="onCancel"
      no-close-on-backdrop
      no-close-on-esc
    >
      <b-container fluid>
        <div class="cropUnit" v-for="(item, index) in cropUnitLi" :key="index">
          <!-- container -->
          <b-row>
            <b-col>
              <b-row>
                <!-- select crop -->
                <b-col>
                  <b-form-group :label="index + 1 + '. Crop for Rotation'">
                    <b-form-select
                      v-model="
                        crop_rotation_and_residue_management.crop_infos[index]
                          .crop_name
                      "
                      :options="crops"
                    >
                    </b-form-select>
                  </b-form-group>
                </b-col>

                <!-- select emergence -->
                <b-col>
                  <b-form-group label="Emergence">
                    <b-form-datepicker
                      :date-format-options="{
                        year: 'numeric',
                        month: 'numeric',
                        day: 'numeric',
                      }"
                      locale="en"
                      placeholder="date"
                      v-model="
                        crop_rotation_and_residue_management.crop_infos[index]
                          .emergence_date
                      "
                    ></b-form-datepicker>
                  </b-form-group>
                </b-col>

                <!-- select harvest -->
                <b-col>
                  <b-form-group label="Final Harvest">
                    <b-form-datepicker
                      :date-format-options="{
                        year: 'numeric',
                        month: 'numeric',
                        day: 'numeric',
                      }"
                      locale="en"
                      placeholder="date"
                      v-model="
                        crop_rotation_and_residue_management.crop_infos[index]
                          .harvest_date
                      "
                    ></b-form-datepicker>
                  </b-form-group>
                </b-col>
              </b-row>

              <!-- residue place -->
              <residue-management
                :ResidueIndex="index"
                :ResidueShow="item"
                v-model="
                  crop_rotation_and_residue_management.crop_infos[index]
                    .residue_value
                "
              ></residue-management>

              <b-row class="mt-2">
                <b-col></b-col>
                <b-col>
                  <!-- add pesticide application -->
                  <b-button
                    @click="SetPesticideApplication(index)"
                    block
                    variant="outline-primary"
                  >
                    Pesticide Application
                  </b-button>

                  <!-- pesticide application modal -->
                  <pesticide-application
                    :ref="`pesticide_${index}`"
                    :pesticideId="`pesticideId_${index}`"
                  ></pesticide-application>
                </b-col>
                <b-col></b-col>
              </b-row>
            </b-col>

            <!-- delete crop unit button -->
            <b-col lg="1">
              <b-button
                variant="outline-danger"
                pill
                style="height: 100%"
                @click="DeleteCropUnit(index)"
              >
                <b-icon icon="x-circle-fill" aria-hidden="true"></b-icon>
              </b-button>
            </b-col>
          </b-row>
          <hr />
        </div>

        <!-- add another crop button -->
        <b-row class="mt-2">
          <b-col>
            <b-button block @click="AddCropUnit" variant="success">
              <b-icon icon="plus-square" aria-hidden="true"></b-icon>
              Add Another Crop
            </b-button>
          </b-col>
        </b-row>
      </b-container>

      <template #modal-footer="{ ok, cancel }">
        <!-- Emulate built in modal footer ok and cancel button actions -->
        <b-button variant="danger" @click="cancel()"> Cancel </b-button>
        <b-button variant="success" @click="ok()"> OK </b-button>
      </template>
    </b-modal>
  </div>
</template>

<script>
import PesticideApplication from "./CropRotationComps/PesticideApplication.vue";
import ResidueManagement from "./CropRotationComps/ResidueManagement.vue";
import { CompareDate } from "@/util/TimeUtil";

export default {
  name: "crop-rotation",
  components: {
    PesticideApplication,
    ResidueManagement,
  },
  data() {
    return {
      cropUnitLi: [false],
      // pesticide application stuff
      crop_rotation_and_residue_management: {
        crop_infos: [
          {
            crop_name: null,
            emergence_date: null,
            harvest_date: null,
            residue_value: null,
          },
        ],
      },
    };
  },
  props: {
    crops: {},
    crop_rotation_and_residue_management_res: {},
  },
  model: {
    prop: "crop_rotation_and_residue_management_res",
    event: "crrm",
  },

  // methods
  methods: {
    AddCropUnit() {
      this.crop_rotation_and_residue_management.crop_infos.push({
        crop_name: null,
        emergence_date: null,
        harvest_date: null,
        residue_value: null,
      });
      this.cropUnitLi.push(false);
    },

    DeleteCropUnit(index) {
      let arr = this.cropUnitLi.concat();
      arr.splice(index, 1);
      this.cropUnitLi = arr;
    },

    SetPesticideApplication(index) {
      this.$refs[`pesticide_${index}`][0].showPesticideModel();
    },

    timeCompare() {
      const len = this.crop_rotation_and_residue_management.crop_infos.length;
      const CropInfos = this.crop_rotation_and_residue_management.crop_infos;
      for (let i = 0; i < len; i += 1) {
        // -------------------- time interval --------------------
        // check global time interval
        const vuex_first_emergence = this.$store.state.first_emergence;
        const vuex_last_harvest = this.$store.state.last_harvest;
        const cur_first_emergence = CropInfos[i].emergence_date;
        const cur_last_harvest = CropInfos[i].harvest_date;
        let res_first_emergence = cur_first_emergence;
        let res_last_harvest = cur_last_harvest;

        // compare emergence date
        if (vuex_first_emergence === null) {
          res_first_emergence = cur_first_emergence;
        } else {
          const flag_emergence = CompareDate(
            vuex_first_emergence,
            cur_first_emergence
          );
          // false: vuex >  cur
          // true : vuex <= cur
          if (flag_emergence === false) {
            res_first_emergence = vuex_first_emergence;
          }
        }

        // compare harvest date
        if (vuex_last_harvest === null) {
          res_last_harvest = cur_last_harvest;
        } else {
          const flag_harvest = CompareDate(vuex_last_harvest, cur_last_harvest);
          // false: vuex > cur
          if (flag_harvest === true) {
            res_last_harvest = vuex_last_harvest;
          }
        }

        const TimeInterval = {
          first_emergence: res_first_emergence,
          last_harvest: res_last_harvest,
        };
        this.$store.commit("SetTimeInterval", TimeInterval);
        // -------------------- time interval --------------------
      }
    },

    onOk() {
      this.timeCompare();
      this.crop_rotation_and_residue_management.crop_infos.push({
        crop_name: this.$store.state.step2.Crop,
        emergence_date: this.$store.state.step2.EmergenceDate,
        harvest_date: this.$store.state.step2.HarvestDate,
        residue_value: -1,
      });
      this.$emit("crrm", this.crop_rotation_and_residue_management);
    },

    onCancel() {
      this.$emit("modal-cancel");
    },
  },
};
</script>
