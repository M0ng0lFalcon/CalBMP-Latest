<template>
  <div class="PesticideApplication">
    <b-modal :id="pesticideId" size="xl" title="Pesticide Application">
      <!-- number of pesticide -->
      <b-row>
        <b-col>
          <b-form-group label="No. of Application">
            <b-form-spinbutton
              v-model="NOofApp"
              min="1"
              max="100"
              inline
            ></b-form-spinbutton>
          </b-form-group>
        </b-col>
      </b-row>

      <b-row v-for="n in NOofApp" :key="n">
        <!-- pesticide date -->
        <b-col md="6" lg="2">
          <b-form-group label="Date">
            <b-form-datepicker
              :id="'DateSelect_' + n"
              :date-format-options="{
                year: 'numeric',
                month: 'numeric',
                day: 'numeric',
              }"
              locale="en"
              v-model="date[n - 1]"
              placeholder="date"
            ></b-form-datepicker>
          </b-form-group>
        </b-col>

        <!-- pesticide name -->
        <b-col md="6" lg="2">
          <b-form-group label="Pesticide">
            <b-form-select
              :id="'Pesticide_' + n"
              v-model="Pesticide[n - 1]"
              :options="Pesticides"
            >
            </b-form-select>
          </b-form-group>
        </b-col>

        <!-- pesticide amount -->
        <b-col md="6" lg="2">
          <b-form-group label="Amount (lb/acre)">
            <b-form-input
              :id="'AmountInput_' + n"
              v-model="amount[n - 1]"
              placeholder="Enter Amount"
              type="number"
              v-b-popover.focus="{
                variant: 'info',
                content:
                  'Please refer to the maximum value on the product label',
              }"
            >
            </b-form-input>
          </b-form-group>
        </b-col>

        <!-- pesticide equipment -->
        <b-col md="6" lg="2">
          <b-form-group label="App Method">
            <b-form-select
              :id="'AppEquipment_' + n"
              v-model="ApplicationEquipment[n - 1]"
              :options="ApplicationEquipments"
            >
            </b-form-select>
          </b-form-group>
        </b-col>

        <!-- pesticide app method -->
        <b-col md="6" lg="2" v-if="ApplicationEquipment[n - 1] != null">
          <b-form-group label="App Location">
            <b-form-select
              :id="'AppMethod_' + n"
              v-model="ApplicationMethod[n - 1]"
              :options="ApplicationMethods"
            >
            </b-form-select>
          </b-form-group>
        </b-col>

        <!-- pesticide depth -->
        <b-col md="6" lg="2" v-if="ApplicationMethod[n - 1] > 2">
          <b-form-group label="Add Depth">
            <b-form-input
              :id="'Depth_' + name"
              type="number"
              v-model="depth[n - 1]"
            >
            </b-form-input>
          </b-form-group>
        </b-col>
      </b-row>
    </b-modal>
  </div>
</template>

<script>
import { GetPesticide } from "@/server/basicDataServer";

export default {
  name: "pesticide-application",
  props: {
    pesticideId: String,
  },
  data() {
    return {
      NOofApp: 1,
      date: [],
      Pesticide: [],
      amount: [],
      ApplicationEquipment: [],
      ApplicationMethod: [],
      depth: [],

      // options
      Pesticides: [],
      ApplicationEquipments: [
        { value: 0.99, text: "Ground-applied Sprayer" },
        { value: 0.95, text: "Ariel sprayer" },
        { value: 1, text: "Others" },
      ],
      ApplicationMethods: [
        { value: 1, text: "Below crop" },
        { value: 4, text: "Below crop (DEPI)" },
        { value: 7, text: "Below crop (T-band)" },
        { value: 2, text: "Above crop" },
        { value: 9, text: "Above crop (DEPI)" },
      ],
    };
  },
  async mounted() {
    this.Pesticides = await GetPesticide();
  },
  // methods
  methods: {
    showPesticideModel() {
      this.$bvModal.show(this.pesticideId);
    },
  },
};
</script>
