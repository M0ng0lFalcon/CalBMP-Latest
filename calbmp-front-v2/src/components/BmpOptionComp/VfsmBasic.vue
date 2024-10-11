<template>
  <div class="vfsm_basic">
    <b-modal
      id="vfsm_basic_modal"
      centered
      title="VFS Model"
      @ok="onOk"
      @cancel="onCancel"
      @close="onCancel"
      no-close-on-backdrop
      no-close-on-esc
    >
      <b-img src="@/assets/VFSM.png" fluid alt="vfsm"></b-img>
      <b-form-group id="vegetation">
        <b-row>
          <b-col>
            <b-form-group label="Vegetation">
              <!-- vegetation select -->
              <b-form-select
                v-model="vegetation"
                :options="vegetation_list"
                @change="get_max_height"
              ></b-form-select>
            </b-form-group>
          </b-col>

          <b-col>
            <b-form-group label="Max height">
              <!-- max height of vegetation -->
              <b-form-input
                id="max_height"
                placeholder="max height"
                type="number"
                v-model="maximum_height"
              ></b-form-input>
              <b-tooltip target="max_height">
                Maximum Height of vegetation.
              </b-tooltip>
            </b-form-group>
          </b-col>
        </b-row>
      </b-form-group>

      <!-- f_width and vl -->
      <b-form-group>
        <b-row>
          <b-col>
            <b-form-group label="FWidth">
              <b-form-input
                id="f-width"
                placeholder="f-width"
                type="number"
                v-model="f_width"
              ></b-form-input>
              <b-tooltip target="f-width"> Width of the strip(m).</b-tooltip>
              <b-tooltip target="f-width" placement="bottom">
                FWidth and SWidth can be different.
              </b-tooltip>
            </b-form-group>
          </b-col>

          <b-col>
            <b-form-group label="VL">
              <b-form-input
                id="vl"
                placeholder="vl"
                type="number"
                v-model="vl"
              ></b-form-input>
              <b-tooltip target="vl"> Length of the VFS(m).</b-tooltip>
            </b-form-group>
          </b-col>
        </b-row>
      </b-form-group>

      <b-form-group label="WTD">
        <b-row>
          <b-col md="3">
            <b-form-checkbox id="wtd_checkbox" v-model="wtd_status" v-b-tooltip.hover title="Water Table Depth">
              Use WTD
            </b-form-checkbox>
          </b-col>
          <b-col>
            <!-- wtd -->
            <b-form-input
              v-if="wtd_status"
              id="wtd"
              placeholder="WTD"
              type="number"
              v-model="wtd"
            ></b-form-input>
          </b-col>
        </b-row>

        <b-tooltip target="wtd"> Water table depth(m).</b-tooltip>
      </b-form-group>

      <b-form-group>
        <!-- swidth and slength -->
        <b-row>
          <b-col>
            <b-form-group label="SWidth">
              <b-form-input
                id="swidth"
                placeholder="swidth"
                type="number"
                v-model="swidth"
                @change="change_length"
              ></b-form-input>
              <b-tooltip target="swidth"> Source area width.</b-tooltip>
            </b-form-group>
          </b-col>
          <b-col>
            <b-form-group label="SLength">
              <b-form-input
                id="slength"
                placeholder="slength"
                type="number"
                v-model="slength"
                disabled
              ></b-form-input>
              <b-tooltip target="slength">
                Source area flow path length.
              </b-tooltip>
            </b-form-group>
          </b-col>
        </b-row>
      </b-form-group>

      <b-form-group>
        <b-form-checkbox v-model="design_mode"> Design Mode</b-form-checkbox>
      </b-form-group>

      <hr />

      <div v-if="design_mode">
        <b-form-group>
          <b-row>
            <b-col md="3" sm="3" lg="3">
              <b-form-group label="Storm Type">
                Precipitation Depth(inch)
              </b-form-group>
            </b-col>
            <b-col>
              <b-form-group label="Ⅰ">
                <b-form-input type="number"></b-form-input>
              </b-form-group>
            </b-col>
            <b-col>
              <b-form-group label="Ⅱ">
                <b-form-input type="number"></b-form-input>
              </b-form-group>
            </b-col>
            <b-col>
              <b-form-group label="Ⅲ">
                <b-form-input type="number"></b-form-input>
              </b-form-group>
            </b-col>
            <b-col>
              <b-form-group label="IA">
                <b-form-input type="number"></b-form-input>
              </b-form-group>
            </b-col>
          </b-row>
        </b-form-group>
        <b-form-group>
          <b-row>
            <b-col md="3" lg="3" sm="3">
              <b-form-group> Storm Duration(h) </b-form-group>
            </b-col>
            <b-col>
              <b-form-group>
                <b-form-input type="number"></b-form-input>
              </b-form-group>
            </b-col>
            <b-col>
              <b-form-group>
                <b-form-input type="number"></b-form-input>
              </b-form-group>
            </b-col>
            <b-col>
              <b-form-group>
                <b-form-input type="number"></b-form-input>
              </b-form-group>
            </b-col>
            <b-col>
              <b-form-group>
                <b-form-input type="number"></b-form-input>
              </b-form-group>
            </b-col>
          </b-row>
        </b-form-group>

        <hr />

        <b-form-group label="Specify VFS length(m)">
          <b-row>
            <b-col>
              <b-form-group label="Lower">
                <b-form-input type="number"></b-form-input>
              </b-form-group>
            </b-col>
            <b-col>
              <b-form-group label="Upper">
                <b-form-input type="number"></b-form-input>
              </b-form-group>
            </b-col>
            <b-col>
              <b-form-group label="Increment">
                <b-form-input type="number"></b-form-input>
              </b-form-group>
            </b-col>
          </b-row>
        </b-form-group>
      </div>

      <template #modal-footer="{ ok, cancel }">
        <!-- Emulate built in modal footer ok and cancel button actions -->
        <b-button variant="danger" @click="cancel()"> Cancel</b-button>
        <b-button variant="success" @click="ok()"> OK</b-button>
      </template>
    </b-modal>
  </div>
</template>

<script>
import { GetVegetation, GetVegetationModel } from "@/server/vfsmServer";

export default {
  name: "vfsm-basic",
  computed: {
    vfsm_param() {
      return {
        created_at: this.$store.state.created_time,
        ikw_param: {
          vegetation: this.vegetation,
          f_width: parseFloat(this.f_width),
          vl: parseFloat(this.vl),
        },
        igr_param: {
          maximum_height: parseFloat(this.maximum_height),
        },
        iso_param: {
          wtd: parseFloat(this.wtd),
        },
        iro_param: {
          swidth: parseFloat(this.swidth),
          slength: parseFloat(this.slength),
        },
      };
    },

    area() {
      return this.$store.state.step1.FieldSize;
    },
  },
  data() {
    return {
      vegetation_list: [],
      vegetation: "",
      maximum_height: 0,
      vl: 0,
      wtd: -1,
      f_width: 0,
      swidth: 0,
      slength: 0,
      wtd_status: false,

      design_mode: false,
    };
  },

  props: {
    vfsm_model: {},
  },
  model: {
    prop: "vfsm_model",
    event: "vfsm_e",
  },

  methods: {
    async get_vegetation() {
      this.vegetation_list = await GetVegetation();
    },

    returnVfsmE() {
      // convert rate to integer
      this.vfsm_model = this.vfsm_param;
      this.$emit("vfsm_e", this.vfsm_model);
    },

    onOk() {
      this.returnVfsmE();
      console.log(this.vfsm_param);
    },

    async get_max_height() {
      let vegmodel = await GetVegetationModel(this.vegetation);
      this.maximum_height = vegmodel.MaximumHeight;
    },

    change_length() {
      this.slength = this.area / this.swidth;
    },

    onCancel() {
      this.$emit("modal-cancel");
    },
  },

  mounted() {
    this.get_vegetation();
  },
};
</script>
