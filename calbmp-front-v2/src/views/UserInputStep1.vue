<template>
  <div class="UserInputStep1">
    <b-overlay :show="over_lay_show">
      <b-row>
        <!-- User input step 1 panel  -->
        <b-col md="6" lg="4">
          <b-card title="Find your field" class="mt-5">
            <b-form @submit="onSubmit" @reset="onReset">
              <b-row>
                <b-col>
                  <!-- step1: select county -->
                  <b-form-group label="County">
                    <b-form-select
                        v-model="userinputStep1.county"
                        :options="counties"
                        @change="getZipcode"
                        required
                    >
                    </b-form-select>
                  </b-form-group>
                </b-col>

                <b-col>
                  <!-- step2: select zipcode -->
                  <b-form-group label="ZIP CODE">
                    <b-form-select
                        v-model="userinputStep1.ZipCode"
                        :options="zipCodes"
                        @change="getNeededData"
                        required
                    >
                    </b-form-select>
                  </b-form-group>
                </b-col>
              </b-row>

              <!-- step3: select soil name : comp name -->
              <b-form-group>
                <template v-slot:label>
                  <p>
                    Soil&nbsp;&nbsp;&nbsp;
                    <a
                        href="https://websoilsurvey.sc.egov.usda.gov/App/WebSoilSurvey.aspx"
                        target="_blank"
                    >
                      Web soil survey
                      <b-icon
                          v-b-popover.hover.top="
                          'looking for the major soil series in your field.'
                        "
                          title="Tips"
                          icon="info-circle-fill"
                          variant="primary"
                      ></b-icon>
                    </a>
                  </p>
                </template>

                <b-form-select
                    v-model="compValue"
                    :options="compNames"
                    @change="setCokeyMukey"
                    required
                >
                </b-form-select>

                <!--                <b-form-input-->
                <!--                  list="soil_list_opts"-->
                <!--                  v-model="soil_select"-->
                <!--                  @change="get_zipcode_by_soil"-->
                <!--                ></b-form-input>-->

                <!--                <datalist id="soil_list_opts">-->
                <!--                  <option v-for="soil in soil_list" :key="soil.ID">-->
                <!--                    {{ soil.Muname }}-->
                <!--                  </option>-->
                <!--                </datalist>-->

                <div class="mt-3" v-if="compValue != null">
                  <p>Soil Series : {{ compValue.compName }}</p>
                </div>
              </b-form-group>

              <b-form-group>
                <b-form-checkbox v-model="userinputStep1.know_slope">
                  I know the specific field slope.
                </b-form-checkbox>

                <b-form-input
                    v-if="userinputStep1.know_slope"
                    type="number"
                    step="0.00001"
                    placeholder="Slope value"
                    :required="userinputStep1.know_slope"
                    v-model="userinputStep1.slope"
                ></b-form-input>
              </b-form-group>

              <b-form-group label="Area (acre)">
                <b-form-input
                    type="number"
                    step="0.00001"
                    v-model="userinputStep1.FieldSize"
                    required
                ></b-form-input>
              </b-form-group>

              <b-form-group label="Soil Crusting(%)">
                <b-form-input
                    type="number"
                    step="0.00001"
                    v-b-popover.focus="{
                          variant: 'info',
                          content:
                            'This is the proportion of soil crusting area in the field area',
                        }"
                    v-model="userinputStep1.soil_crusting"
                ></b-form-input>
              </b-form-group>

              <b-form-group>
                <b-button variant="danger" type="reset" class="mr-3">
                  Reset
                </b-button>
                <b-button variant="primary" type="submit" class="mr-3">
                  Next
                </b-button>
                <b-button variant="danger" @click="clearSteps"> Clear</b-button>
              </b-form-group>
            </b-form>
          </b-card>
        </b-col>

        <!-- show map -->
        <b-col md="6" lg="8">
          <b-card title="Your Nearest Weather Station" class="mt-5">
            <!-- 34.683815, 241.976076 -->
            <div style="height: 400px">
              <div class="info"></div>
              <l-map
                  style="height: 100%; width: 100%"
                  :zoom="zoom"
                  :center="center"
                  @update:zoom="zoomUpdated"
                  @update:center="centerUpdated"
                  @update:bounds="boundsUpdated"
              >
                <l-tile-layer :url="url"></l-tile-layer>
                <l-circle
                    v-if="userinputStep1.ZipCode != null"
                    :lat-lng="circle.center"
                    :radius="circle.radius"
                    :color="circle.color"
                />
              </l-map>
            </div>
          </b-card>
        </b-col>
      </b-row>
    </b-overlay>
  </div>
</template>

<script>
import {LMap, LTileLayer, LCircle} from "vue2-leaflet";

import {
  GetCouties,
  getZipcode,
  getStation,
  getCompnameMukeyCokey,
  submitStep1,
} from "@/server/basicDataServer";

export default {
  components: {
    LMap,
    LTileLayer,
    LCircle,
  },
  data() {
    return {
      over_lay_show: false,
      // form place
      userinputStep1: {
        county: null,
        comp_name: null,
        muname: null,
        // new items
        ZipCode: null,
        ClimateId: null,
        FieldSize: null,
        ifBmp: false,
        if_vfsm: false,
        muKey: null,
        coKey: null,
        log: null,
        lat: null,
        know_slope: false,
        slope: null,
        soil_crusting: 0,
      },

      // select options
      counties: [],
      zipCodes: [],
      compNames: [],
      soil_list: [],
      soil_select: null,

      // temp variable
      mukeys: [],
      compValue: null,

      // map place
      // 34.683815, 241.976076
      url: "https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png",
      center: [34.683815, 241.976076],
      bounds: null,
      zoom: 6,
      circle: {
        center: [34.683815, 241.976076],
        radius: 4500,
        color: "red",
      },
    };
  },
  validations: {
    userinputStep1: {},
  },
  methods: {
    async getCounties() {
      this.counties = await GetCouties();
    },

    async getZipcode() {
      this.userinputStep1.ZipCode = null;
      this.zipCodes = await getZipcode(this.userinputStep1.county);
    },

    async getNeededData() {
      // get station data
      const resData = await getStation(this.userinputStep1.ZipCode);

      this.userinputStep1.ClimateId = resData.ClimateId;
      this.userinputStep1.log = resData.log;
      this.userinputStep1.lat = resData.lat;
      this.center = resData.center;
      this.circle.center = this.center;
      this.zoom = resData.zoom;

      // get CompnameMukeyCokey
      this.compNames = await getCompnameMukeyCokey(this.userinputStep1.ZipCode);
    },

    setCokeyMukey() {
      this.userinputStep1.coKey = this.compValue.cokey;
      this.userinputStep1.muKey = this.compValue.mukey;
      this.userinputStep1.comp_name = this.compValue.compName;
      this.userinputStep1.muname = this.compValue.muname;

      console.log("UserStep1", this.userinputStep1);
    },

    // ------------------------
    // map place
    zoomUpdated(zoom) {
      this.zoom = zoom;
    },
    centerUpdated(center) {
      this.center = center;
    },
    boundsUpdated(bounds) {
      this.bounds = bounds;
    },
    // ------------------------

    clearSteps() {
      this.$store.commit("ClearStep1Info");
      this.$store.commit("ClearStep2Info");
      this.$store.commit("ClearCreatedTime");
      window.location.reload();
    },

    // next page fun
    onSubmit(event) {
      event.preventDefault();
      // validate form
      // save to local storage
      this.userinputStep1.ifBmp = false;
      this.userinputStep1.if_vfsm = false;
      if (this.userinputStep1.soil_crusting === null) {
        console.log('soil crusting:', this.userinputStep1.soil_crusting);
        this.userinputStep1.soil_crusting = 0;
      }

      // change type
      this.userinputStep1.slope = parseFloat(this.userinputStep1.slope);
      this.userinputStep1.soil_crusting = parseFloat(this.userinputStep1.soil_crusting);

      // store into vuex store
      this.$store.commit("SetStep1Info", this.userinputStep1);


      submitStep1(this.userinputStep1);
      this.$router.push("/userinputStep2");
    },

    onReset() {
      this.$store.commit("ClearStep1Info");
      this.userinputStep1 = {
        county: null,
        comp_name: null,
        muname: null,
        // new items
        ZipCode: null,
        ClimateId: null,
        FieldSize: null,
        ifBmp: false,
        if_vfsm: false,
        muKey: null,
        coKey: null,
        log: null,
        lat: null,
      };
      this.zoom = 6;
      this.center = [34.683815, 241.976076];
      // this.counties = [];
      this.zipCodes = [];
      this.compNames = [];
      this.compValue = null;
      this.bounds = null;
    },
  },
  async mounted() {
    this.over_lay_show = true;
    // init counties list
    await this.getCounties();
    // await this.getAllSoil();

    // init map
    this.center = [34.683815, 241.976076];

    this.$store.commit("ClearBmp");
    this.$store.commit("ClearBmpOptionsList");
    this.$store.commit("ClearVfsm");
    this.$store.commit("ClearVfsmOptionList");
    this.$store.commit("ClearCreatedTime");

    const step1 = this.$store.state.step1;
    if (step1 !== null) {
      this.userinputStep1.county = step1.county;
      await this.getZipcode();
      this.userinputStep1.know_slope = step1.know_slope;
      this.userinputStep1.slope = step1.slope;
      this.userinputStep1.ZipCode = step1.ZipCode;
      this.userinputStep1.FieldSize = step1.FieldSize;
      this.userinputStep1.soil_crusting = step1.soil_crusting;
      await this.getNeededData();
      this.userinputStep1 = step1;
      // console.log(this.compNames);
      this.compValue = {
        cokey: step1.coKey,
        mukey: step1.muKey,
        compName: step1.comp_name,
        muname: step1.muname,
      };
      console.log("this.compvalue:", this.compValue);
      this.setCokeyMukey();
    }
    if (this.userinputStep1.soil_crusting === null) {
      this.userinputStep1.soil_crusting = 0;
    }
    this.over_lay_show = false;
    // this.$store.commit("ClearStep1Info");
    // this.$store.commit("ClearStep2Info");
    // this.$store.commit("ClearCreatedTime");
  },
};
</script>

<style>
@font-face {
  font-family: "amplesoft";
  src: url("../fonts/FontsFree-Net-AmpleSoftProMedium.ttf");
}

.UserInputStep1 {
  font-family: "amplesoft", serif;
}
</style>
