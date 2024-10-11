<template>
  <div class="CoverCropsComp">
    <b-modal
      id="coverCrops"
      centered
      title="Cover Crop"
      size="lg"
      @ok="onOk"
      @cancel="onCancel"
      @close="onCancel"
      no-close-on-backdrop
      no-close-on-esc
    >
      <b-container fluid>
        <b-row>
          <!-- show first crop -->
          <b-col>
            <b-form-group label="First Crop">
              <b-form-input
                :value="preCropInfo.preCrop"
                disabled
              ></b-form-input>
            </b-form-group>
          </b-col>

          <!-- show emergence date of first crop -->
          <b-col>
            <b-form-group label="Emergence">
              <b-form-input
                disabled
                :value="preCropInfo.preEmergence"
              ></b-form-input>
            </b-form-group>
          </b-col>

          <!-- show harvest date of first crop -->
          <b-col>
            <b-form-group label="Final Harvest">
              <b-form-input
                disabled
                :value="preCropInfo.preHarvest"
              ></b-form-input>
            </b-form-group>
          </b-col>
        </b-row>

        <b-row>
          <!-- select cover crop -->
          <b-col>
            <b-form-group label="Cover Crop">
              <b-form-select
                v-model="cover_crop_temp.cover_crop"
                :options="crops"
              >
              </b-form-select>
            </b-form-group>
          </b-col>

          <!-- select emergence date of cover crop -->
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
                v-model="cover_crop_temp.cover_crop_emergence"
                no-flip
              ></b-form-datepicker>
            </b-form-group>
          </b-col>

          <!-- select harvest date of cover crop -->
          <b-col>
            <b-form-group label="Mowed">
              <b-form-datepicker
                :date-format-options="{
                  year: 'numeric',
                  month: 'numeric',
                  day: 'numeric',
                }"
                locale="en"
                placeholder="date"
                v-model="cover_crop_temp.cover_crop_harvest"
                no-flip
              ></b-form-datepicker>
            </b-form-group>
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
import { CompareDate } from "@/util/TimeUtil";
export default {
  name: "cover-crops",
  data() {
    return {
      cover_crop_temp: {
        cover_crop: null,
        cover_crop_emergence: null,
        cover_crop_harvest: null,
      },
      preCropInfo: {
        preCrop: this.$store.state.step2.Crop,
        preEmergence: this.$store.state.step2.EmergenceDate,
        preHarvest: this.$store.state.step2.HarvestDate,
      },
    };
  },
  props: {
    crops: {},
    cover_crop: {},
  },
  model: {
    prop: "cover_crop",
    event: "cp",
  },

  // methods
  methods: {
    timeCompare() {
      // -------------------- time interval --------------------
      // check global time interval
      const vuex_first_emergence = this.$store.state.first_emergence;
      const vuex_last_harvest = this.$store.state.last_harvest;
      const cur_first_emergence = this.cover_crop_temp.cover_crop_emergence;
      const cur_last_harvest = this.cover_crop_temp.cover_crop_harvest;
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
    },
    returnCP() {
      this.timeCompare();
      this.$emit("cp", this.cover_crop_temp);
    },
    onOk() {
      this.returnCP();
    },
    onCancel() {
      this.$emit("modal-cancel");
    },
  },
};
</script>
