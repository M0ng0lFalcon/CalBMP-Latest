<template>
  <div class="UserinputStep2">
    <b-overlay :show="overlayShow">
      <b-card class="mt-1">
        <b-form @submit="submit">
          <b-row>
            <b-col md="12" lg="12">
              <!-- step1:
                Crop;
                Emergence Data;
                Maturity Data;
                Harvest Data;
                Irrigation Type;
            -->
              <b-card title="Crop Management" border-variant="light">
                <b-row>
                  <b-col md="6" lg="3">
                    <!-- select Crop -->
                    <b-form-group label="Crop">
                      <b-form-select
                        v-model="userinputStep2.Crop"
                        :options="Crops"
                        required
                      >
                      </b-form-select>
                    </b-form-group>
                  </b-col>
                  <b-col md="6" lg="2">
                    <!-- select Emergence Data -->
                    <b-form-group label="Emergence Date">
                      <b-form-datepicker
                        id="EmergenceDatePicker"
                        :date-format-options="{
                          year: 'numeric',
                          month: 'numeric',
                          day: 'numeric',
                        }"
                        locale="en"
                        v-model="userinputStep2.EmergenceDate"
                        placeholder="Emergence"
                        no-flip
                        :max="new Date('2023-12-31')"
                        :min="new Date('1998-01-01')"
                        hide-header
                        required
                      ></b-form-datepicker>
                    </b-form-group>
                  </b-col>
                  <b-col md="6" lg="2">
                    <!-- select Maturity Data -->
                    <b-form-group label="Maturity Date">
                      <b-form-datepicker
                        id="MaturityDatePicker"
                        :date-format-options="{
                          year: 'numeric',
                          month: 'numeric',
                          day: 'numeric',
                        }"
                        locale="en"
                        v-model="userinputStep2.MaturityDate"
                        placeholder="Maturity"
                        :max="new Date('2023-12-31')"
                        :min="new Date('1998-01-01')"
                        hide-header
                        required
                      ></b-form-datepicker>
                    </b-form-group>
                  </b-col>
                  <b-col md="6" lg="2" v-b-tooltip.hover title="This is the final harvest date to ensure multiple harvests be included.">
                    <!-- select Harvest Data -->
                    <b-form-group label="Harvest Date">
                      <b-form-datepicker
                        id="HarvestDatePicker"
                        :date-format-options="{
                          year: 'numeric',
                          month: 'numeric',
                          day: 'numeric',
                        }"
                        locale="en"
                        v-model="userinputStep2.HarvestDate"
                        placeholder="Harvest"
                        no-flip
                        :max="new Date('2023-12-31')"
                        :min="new Date('1998-01-01')"
                        hide-header
                        required
                      ></b-form-datepicker>
                    </b-form-group>
                  </b-col>
                  <b-col md="6" lg="2" offset-md="0" offset-lg="1">
                    <!-- select Irrigation Type -->
                    <b-form-group label="Irrigation Type">
                      <b-form-select
                        v-model="userinputStep2.IrrigationType"
                        :options="IrrigationTypes"
                        @click="clear_irrg_type"
                        @change="select_irrg_type"
                        required
                      >
                      </b-form-select>
                    </b-form-group>
                  </b-col>
                </b-row>
                <b-row v-if="userinputStep2.IrrigationType === 0">
                  <b-col align-self="center" md="2">
                    <b-form-group label="No. of Irrigation">
                      <b-form-spinbutton
                        v-model="NOofIrrigation"
                        min="1"
                        max="100"
                        inline
                      ></b-form-spinbutton>
                    </b-form-group>
                  </b-col>
                  <div v-for="n in NOofIrrigation" :key="n + 100">
                    <b-col>
                      <b-form-group label="Date">
                        <b-form-datepicker
                          :id="generateId('IrriDateSelect_', n)"
                          :date-format-options="{
                            year: 'numeric',
                            month: 'numeric',
                            day: 'numeric',
                          }"
                          locale="en"
                          v-model="userinputStep2.irrigation_date[n - 1]"
                          placeholder="date"
                          no-flip
                          :max="new Date('2023-12-31')"
                          :min="new Date('1998-01-01')"
                          hide-header
                          required
                        ></b-form-datepicker>
                      </b-form-group>
                      <b-form-group label="Amount (inch/day)">
                        <b-form-input
                          :id="generateId('IrriAmountInput_', n)"
                          v-model="userinputStep2.irrigation_amount[n - 1]"
                          placeholder="Enter Amount"
                          type="number"
                          step="0.00001"
                          required
                        >
                        </b-form-input>
                      </b-form-group>
                    </b-col>
                  </div>
                </b-row>
              </b-card>
            </b-col>

            <b-col md="12" lg="12">
              <b-card class="mt-1" border-variant="light">
                <b-card-title>
                  Pesticide Application
                  <b-icon
                    v-b-popover.hover.top="
                      'Up to three pesticides could be simulated for each running.'
                    "
                    title="Tips"
                    icon="info-circle-fill"
                    variant="primary"
                  ></b-icon>
                  <b-icon
                      class="ml-2"
                      v-b-popover.hover.top="
                      'Please input pesticide applications from previous crop cycle as well as current crop if they exist.'
                    "
                      title="Tips"
                      icon="info-circle-fill"
                      variant="success"
                  ></b-icon>
                </b-card-title>
                <b-row>
                  <!-- step3:
                Number of Applications
                Date
                -->
                  <b-col md="6" lg="2">
                    <b-form-group label="NO. of App">
                      <!-- <template v-slot:label>
                        <p>NO. of App</p>
                      </template> -->
                      <b-form-select
                        v-model="userinputStep2.NOofApp"
                        :options="NOofApps"
                        required
                      >
                      </b-form-select>
                    </b-form-group>
                  </b-col>
                </b-row>
                <!-- for select -->
                <b-row v-for="n in userinputStep2.NOofApp" :key="n">
                  <b-col md="6" lg="2">
                    <b-form-group label="Date">
                      <b-form-datepicker
                        :id="generateId('DateSelect_', n)"
                        :date-format-options="{
                          year: 'numeric',
                          month: 'numeric',
                          day: 'numeric',
                        }"
                        locale="en"
                        v-model="userinputStep2.date[n - 1]"
                        placeholder="date"
                        no-flip
                        :max="new Date('2023-12-31')"
                        :min="new Date('1998-01-01')"
                        hide-header
                        required
                      ></b-form-datepicker>
                    </b-form-group>
                  </b-col>

                  <b-col md="6" lg="2">
                    <b-form-group label="Pesticide">
                      <b-form-select
                        :id="generateId('Pesticide_', n)"
                        v-model="userinputStep2.Pesticide[n - 1]"
                        :options="Pesticides"
                        required
                      >
                      </b-form-select>
                    </b-form-group>
                  </b-col>

                  <b-col md="6" lg="2">
                    <b-form-group label="Amount (lb/acre)">
                      <b-form-input
                        :id="generateId('AmountInput_', n)"
                        v-model="userinputStep2.amount[n - 1]"
                        placeholder="Enter Amount"
                        type="number"
                        step="0.00001"
                        v-b-popover.focus="{
                          variant: 'info',
                          content:
                            'Please refer to the maximum value on the product label',
                        }"
                        required
                      >
                      </b-form-input>
                    </b-form-group>
                  </b-col>

                  <b-col md="6" lg="2">
                    <b-form-group label="App Method">
                      <b-form-select
                        :id="generateId('AppEquipment_', n)"
                        v-model="userinputStep2.ApplicationEquipment[n - 1]"
                        :options="ApplicationEquipments"
                        required
                      >
                      </b-form-select>
                    </b-form-group>
                  </b-col>

                  <b-col
                    md="6"
                    lg="2"
                    v-if="userinputStep2.ApplicationEquipment[n - 1] != null"
                  >
                    <b-form-group label="App Location">
                      <b-form-select
                        :id="generateId('AppMethod_', n)"
                        v-model="userinputStep2.ApplicationMethod[n - 1]"
                        :options="ApplicationMethods"
                        required
                      >
                      </b-form-select>
                    </b-form-group>
                  </b-col>

                  <b-col
                    md="6"
                    lg="2"
                    v-if="userinputStep2.ApplicationMethod[n - 1] > 2"
                  >
                    <b-form-group label="Add Depth">
                      <b-form-input
                        :id="generateId('Depth_', n)"
                        type="number"
                        v-model="userinputStep2.depth[n - 1]"
                        required
                      >
                      </b-form-input>
                    </b-form-group>
                  </b-col>
                </b-row>
                <!-- end for -->
              </b-card>
            </b-col>
          </b-row>
          <b-card class="mt-1" border-variant="light" body-class="text-center">
            <b-row>
              <b-col md="6" lg="4" offset-md="8" offset-lg="8">
                <b-button
                  variant="danger"
                  class="mt-2"
                  style="margin-right: 10px"
                  >Reset</b-button
                >

                <b-button variant="primary" class="mt-2" type="submit">
                  Submit
                </b-button>
              </b-col>
            </b-row>
          </b-card>
        </b-form>
      </b-card>
    </b-overlay>

    <div class="irrigation_modal_div">
      <b-modal id="irrigation_modal">
        <b-form-group>
          <h5>Do you have runoff generation in your field?</h5>
          <b-form-radio-group
            :options="over_canopy_list"
            v-model="userinputStep2.IrrigationType"
            @change="irrigation_type_view"
          ></b-form-radio-group>
        </b-form-group>

        <b-form-group class="mt-3" v-if="userinputStep2.IrrigationType !== 3">
          <h5>Irrigation rate</h5>
          <b-form-input placeholder="cm" v-model="userinputStep2.user_defined_irrg_rate"></b-form-input>
        </b-form-group>
      </b-modal>
    </div>
  </div>
</template>

<script>
import { GetCrops, GetPesticide, submitStep2 } from "@/server/basicDataServer";

export default {
  data() {
    return {
      userinputStep2: {
        Crop: null,
        EmergenceDate: null,
        MaturityDate: null,
        HarvestDate: null,
        IrrigationType: null,
        IrrigationTypeName: null,
        irrigation_amount: [],
        irrigation_date: [],
        Pesticide: [],
        cntPesticide: 0,
        PesticideSet: null,
        ApplicationEquipment: [],
        ApplicationEquipmentName: [],
        ApplicationMethod: [],
        ApplicationMethodName: [],
        depth: [],
        NOofApp: 1,
        date: [],
        amount: [],
        ifBmp: false,
        // public params
        ZipCode: null,
        ClimateId: null,
        coKey: null,
        muKey: null,
        log: null,
        lat: null,
        comp_name: null,
        user_defined_irrg_rate: null,
      },
      // options
      // row 1
      Crops: [],
      IrrigationTypes: [
        { value: 0, text: "Flood irrigation" },
        { value: 3, text: "Sprinkler irrigation" },
        // { value: 4, text: "Drip irrigation" },
        { value: 2, text: "No irrigation" },
      ],
      // row 2
      Pesticides: [],
      ApplicationEquipments: [
        { value: 0.99, text: "Ground Application (Above Crop)" },
        { value: 0.95, text: "Spray" },
        { value: 1, text: "Chemigate" },
      ],
      ApplicationMethods: [
        { value: 2, text: "Above crop" },
        {
          value: 1,
          text: "Below crop (soil incorporation depth of 4cm, linearly decreasing with depth)",
        },
        {
          value: 4,
          text: "Below crop (user defined soil depth, uniform with depth)",
        },
        {
          value: 5,
          text: "Below crop (user defined incorporation depth, linearly increasing with depth)",
        },
        {
          value: 6,
          text: "Below crop (user defined incorporation depth, linearly decreasing with depth)",
        },
        {
          value: 7,
          text: "Below crop (user defined incorporation depth, T-Band)",
        },
        {
          value: 8,
          text: "Below crop (user defined incorporation depth, chemical incorporated entirely into depth specified)",
        },
      ],
      NOofApps: [],
      NOofIrrigation: 1,
      over_canopy_title: "Do you have runoff generation in your field?",
      over_canopy_list: [
        { text: "No clue", value: 3 },
        {
          text: "Yes",
          value: 6,
        },
        {
          text: "No",
          value: 7,
        },
      ],

      // overlay show
      overlayShow: false,
    };
  },
  methods: {
    // todo: 当重新选的时候，把之前添加的临时的选项给去除
    clear_irrg_type() {
      let tmp_irrg_li = [];
      for (let i = 0; i < this.IrrigationTypes.length; i += 1) {
        let tmp = this.IrrigationTypes[i];
        if (tmp.value !== 6 || tmp.value !== 7) {
          tmp_irrg_li.push(tmp);
        }
      }
      this.IrrigationTypes = tmp_irrg_li;
      console.log("clear irrg:", this.IrrigationTypes);
    },
    select_irrg_type() {
      if (this.userinputStep2.IrrigationType === 3) {
        this.$bvModal.show("irrigation_modal");
      }
    },
    irrigation_type_view() {
      if (this.userinputStep2.IrrigationType === 6) {
        this.IrrigationTypes.push({
          value: 6,
          text: "over canopy sprinkler",
          disabled: true,
        });
      } else if (this.userinputStep2.IrrigationType === 7) {
        this.IrrigationTypes.push({
          value: 7,
          text: "over canopy sprinkler",
          disabled: true,
        });
      }
    },
    generateId(pre, n) {
      return pre + n.toString();
    },
    submit(event) {
      event.preventDefault();
      // show overlay
      this.overlayShow = true;

      // slice pesticide application and irrigation application
      let sliceEnd = this.userinputStep2.NOofApp;
      this.userinputStep2.date = this.userinputStep2.date.slice(0, sliceEnd);
      this.userinputStep2.Pesticide = this.userinputStep2.Pesticide.slice(
        0,
        sliceEnd
      );
      this.userinputStep2.amount = this.userinputStep2.amount.slice(
        0,
        sliceEnd
      );
      this.userinputStep2.ApplicationEquipment =
        this.userinputStep2.ApplicationEquipment.slice(0, sliceEnd);
      this.userinputStep2.ApplicationMethod =
        this.userinputStep2.ApplicationMethod.slice(0, sliceEnd);
      this.userinputStep2.depth = this.userinputStep2.depth.slice(0, sliceEnd);

      sliceEnd = this.NOofIrrigation;
      this.userinputStep2.irrigation_date =
        this.userinputStep2.irrigation_date.slice(0, sliceEnd);
      this.userinputStep2.irrigation_amount =
        this.userinputStep2.irrigation_amount.slice(0, sliceEnd);

      // make pesticide set to list
      this.userinputStep2.PesticideSet = new Set();
      this.userinputStep2.Pesticide.forEach((elem) => {
        this.userinputStep2.PesticideSet.add(elem);
      });
      // get cnt of pesticide
      this.userinputStep2.cntPesticide = this.userinputStep2.PesticideSet.size;
      this.userinputStep2.PesticideSet = Array.from(
        this.userinputStep2.PesticideSet
      );

      // set irrigation type name
      const UserIrrigation = this.userinputStep2.IrrigationType;
      const IrrgOpts = this.IrrigationTypes;
      for (let i = 0; i < IrrgOpts.length; i += 1) {
        if (UserIrrigation === IrrgOpts[i].value) {
          this.userinputStep2.IrrigationTypeName = IrrgOpts[i].text;
        }
      }
      // set equipment name
      const UserEquipment = this.userinputStep2.ApplicationEquipment;
      const EquipOpts = this.ApplicationEquipments;
      for (let i = 0; i < UserEquipment.length; i += 1) {
        for (let j = 0; j < EquipOpts.length; j += 1) {
          if (UserEquipment[i] === EquipOpts[j].value) {
            this.userinputStep2.ApplicationEquipmentName[i] = EquipOpts[j].text;
          }
        }
      }

      // convert irrigation amount to float
      for (
        let i = 0;
        i < this.userinputStep2.irrigation_amount.length;
        i += 1
      ) {
        this.userinputStep2.irrigation_amount[i] = parseFloat(
          this.userinputStep2.irrigation_amount[i]
        ) * 2.54;
      }

      // convert user defined irrigation rate to float
      this.userinputStep2.user_defined_irrg_rate = parseFloat(this.userinputStep2.user_defined_irrg_rate);

      // convert depth amount to float
      for (let i = 0; i < this.userinputStep2.depth.length; i += 1) {
        this.userinputStep2.depth[i] = parseFloat(this.userinputStep2.depth[i]);
      }

      // set method name
      const UserMethod = this.userinputStep2.ApplicationMethod;
      const MethodOpts = this.ApplicationMethods;
      for (let i = 0; i < UserMethod.length; i += 1) {
        for (let j = 0; j < MethodOpts.length; j += 1) {
          if (UserMethod[i] === MethodOpts[j].value) {
            this.userinputStep2.ApplicationMethodName[i] = MethodOpts[j].text;
          }
        }
      }

      // -------------------- time interval: start --------------------
      const res_first_emergence = this.userinputStep2.EmergenceDate; // visual first day
      let res_last_harvest = this.userinputStep2.HarvestDate; // visual last day

      if (this.userinputStep2.IrrigationType === 0) {
        this.userinputStep2.irrigation_date.forEach((irrg_date) => {
          const date1 = new Date(irrg_date);
          const date2 = new Date(res_last_harvest);
          if (date1.getTime() >= date2.getTime()) {
            res_last_harvest = irrg_date;
          }
        });
      }

      this.userinputStep2.if_vfsm = false;
      this.userinputStep2.ifBmp = false;

      const TimeInterval = {
        first_emergence: res_first_emergence,
        last_harvest: res_last_harvest,
      };
      this.$store.commit("SetTimeInterval", TimeInterval);
      // -------------------- time interval: end --------------------

      // submit
      submitStep2(this.userinputStep2).then((res) => {
        console.log(res);
        const Data = res.data.data;
        const CreatedTime = Data.created_time;
        // hide overlay
        this.overlayShow = false;
        // save step2 info to vuex
        this.$store.commit("SetStep2Info", this.userinputStep2);

        // clear bmp
        this.$store.commit("ClearBmp");
        // set created time vuex
        this.$store.commit("SetCreatedTime", CreatedTime);

        this.$router.replace("/visualization");
      });
    },
    async getCrops() {
      this.Crops = await GetCrops();
    },
    async getPesticide() {
      this.Pesticides = await GetPesticide();
    },
  },
  async mounted() {
    this.overlayShow = true;
    // init NOofApps
    this.NOofApps = [];
    for (let i = 1; i <= 20; i += 1) {
      this.NOofApps.push({
        value: i,
        text: i.toString(),
      });
    }
    // init crops
    await this.getCrops();
    // init pesticide
    await this.getPesticide();
    // init public form params
    const step1 = this.$store.state.step1;
    console.log("getting step1 info:", step1);
    this.userinputStep2.ZipCode = step1.ZipCode;
    this.userinputStep2.ClimateId = step1.ClimateId;
    this.userinputStep2.coKey = step1.coKey;
    this.userinputStep2.muKey = step1.muKey;
    this.userinputStep2.log = step1.log;
    this.userinputStep2.lat = step1.lat;
    this.userinputStep2.comp_name = step1.comp_name;
    this.userinputStep2.soil_crusting = step1.soil_crusting;

    const step2 = this.$store.state.step2;
    if (step2 !== null) {
      this.userinputStep2.Crop = step2.Crop;
      this.userinputStep2.EmergenceDate = step2.EmergenceDate;
      this.userinputStep2.MaturityDate = step2.MaturityDate;
      this.userinputStep2.HarvestDate = step2.HarvestDate;
      this.userinputStep2.IrrigationType = step2.IrrigationType;
      this.userinputStep2.IrrigationTypeName = step2.IrrigationTypeName;
      this.userinputStep2.irrigation_amount = step2.irrigation_amount;
      this.userinputStep2.irrigation_date = step2.irrigation_date;
      this.userinputStep2.Pesticide = step2.Pesticide;
      this.userinputStep2.cntPesticide = step2.cntPesticide;
      this.userinputStep2.PesticideSet = step2.PesticideSet;
      this.userinputStep2.ApplicationEquipment = step2.ApplicationEquipment;
      this.userinputStep2.ApplicationEquipmentName =
        step2.ApplicationEquipmentName;
      this.userinputStep2.ApplicationMethod = step2.ApplicationMethod;
      this.userinputStep2.ApplicationMethodName = step2.ApplicationMethodName;
      this.userinputStep2.depth = step2.depth;
      this.userinputStep2.NOofApp = step2.NOofApp;
      this.userinputStep2.date = step2.date;
      this.userinputStep2.amount = step2.amount;
      this.userinputStep2.ifBmp = step2.ifBmp;
      this.userinputStep2.ZipCode = step1.ZipCode;
      this.userinputStep2.ClimateId = step1.ClimateId;
      this.userinputStep2.coKey = step1.coKey;
      this.userinputStep2.muKey = step1.muKey;
      this.userinputStep2.log = step1.log;
      this.userinputStep2.lat = step1.lat;
      this.userinputStep2.comp_name = step1.comp_name;
    }
    this.overlayShow = false;
  },
};
</script>

<style>
@font-face {
  font-family: "amplesoft";
  src: url("../fonts/FontsFree-Net-AmpleSoftProMedium.ttf");
}

.UserinputStep2 {
  font-family: "amplesoft";
}
</style>
