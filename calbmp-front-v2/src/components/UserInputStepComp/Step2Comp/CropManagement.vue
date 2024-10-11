<template>
  <div class="CropManagement">
    <b-card title="Specify Your Crop Management" border-variant="light">
      <b-row>
        <!-- select Crop -->
        <b-col md="6" lg="3">
          <b-form-group label="Crop">
            <b-form-select
              v-model="crop_management_temp.Crop"
              :options="Crops"
              required
            >
            </b-form-select>
          </b-form-group>
        </b-col>

        <!-- select Emergence Data -->
        <b-col md="6" lg="2">
          <b-form-group label="Emergence Date">
            <b-form-datepicker
              id="EmergenceDatePicker"
              :date-format-options="{
                year: 'numeric',
                month: 'numeric',
                day: 'numeric',
              }"
              locale="en"
              v-model="crop_management_temp.EmergenceDate"
              placeholder="Emergence"
              required
            ></b-form-datepicker>
          </b-form-group>
        </b-col>

        <!-- select Maturity Data -->
        <b-col md="6" lg="2">
          <b-form-group label="Maturity Date">
            <b-form-datepicker
              id="MaturityDatePicker"
              :date-format-options="{
                year: 'numeric',
                month: 'numeric',
                day: 'numeric',
              }"
              locale="en"
              v-model="crop_management_temp.MaturityDate"
              placeholder="Maturity"
              required
            ></b-form-datepicker>
          </b-form-group>
        </b-col>

        <!-- select Harvest Data -->
        <b-col md="6" lg="2">
          <b-form-group label="Harvest Date">
            <b-form-datepicker
              id="HarvestDatePicker"
              :date-format-options="{
                year: 'numeric',
                month: 'numeric',
                day: 'numeric',
              }"
              locale="en"
              v-model="crop_management_temp.HarvestDate"
              placeholder="Harvest"
              v-b-tooltip.hover title="Tooltip directive content"
              required
            ></b-form-datepicker>
          </b-form-group>
        </b-col>

        <!-- select Irrigation Type -->
        <b-col md="6" lg="2" offset-md="0" offset-lg="1">
          <b-form-group label="Irrigation Type">
            <b-form-select
              v-model="crop_management_temp.IrrigationType"
              :options="IrrigationTypes"
              required
            >
            </b-form-select>
          </b-form-group>
        </b-col>
      </b-row>
    </b-card>
  </div>
</template>

<script>
import { GetCrops } from "@/server/basicDataServer";
import { IrrigationTypes } from "@/server/inputServer";
export default {
  name: "crop-management",
  props: {
    crop_management: Object,
  },
  model: {
    prop: "crop_management",
    event: "cropManagement",
  },
  data() {
    return {
      crop_management_temp: {
        Crop: null,
        EmergenceDate: null,
        MaturityDate: null,
        HarvestDate: null,
        IrrigationType: null,
      },
      IrrigationTypes: IrrigationTypes,
      Crops: [],
    };
  },
  // methods
  methods: {
    returnCropManagement() {
      this.$emit("cropManagement", this.crop_management_temp);
    },
    async getCrops() {
      this.Crops = await GetCrops();
    },
  },

  mounted() {
    // init crops
    this.getCrops();
  },
  computed: {
    Crop() {
      return this.crop_management_temp.Crop;
    },
    EmergenceDate() {
      return this.crop_management_temp.EmergenceDate;
    },
    MaturityDate() {
      return this.crop_management_temp.MaturityDate;
    },
    HarvestDate() {
      return this.crop_management_temp.HarvestDate;
    },
    IrrigationType() {
      return this.crop_management_temp.IrrigationType;
    },
  },

  watch: {},
};
</script>
