<template>
  <div class="StripCroppingComp">
    <b-modal
      id="stripCropping"
      centered
      title="Strip Cropping"
      @ok="onOk"
      @cancel="onCancel"
      @close="onCancel"
      no-close-on-backdrop
      no-close-on-esc
    >
      <b-row>
        <!-- first crop -->
        <b-col>
          <b-form-group label="First Crop">
            <b-form-input disabled :value="strip_cropping_temp.pre_crop">
            </b-form-input>
          </b-form-group>
        </b-col>

        <!-- first rate -->
        <b-col>
          <b-form-group label="Area Percentage">
            <b-form-spinbutton
              :value="100 - strip_cropping_temp.rate"
              min="1"
              max="100"
              disabled
            ></b-form-spinbutton>
          </b-form-group>
        </b-col>
      </b-row>

      <b-row>
        <!-- strip crop name -->
        <b-col>
          <b-form-group label="Second Crop">
            <b-form-select v-model="strip_cropping_temp.crop" :options="crops">
            </b-form-select>
          </b-form-group>
        </b-col>

        <!-- strip crop rate -->
        <b-col>
          <b-form-group label="Area Percentage">
            <b-form-spinbutton
              v-model="strip_cropping_temp.rate"
              min="1"
              max="100"
            ></b-form-spinbutton>
          </b-form-group>
        </b-col>
      </b-row>

      <template #modal-footer="{ ok, cancel }">
        <!-- Emulate built in modal footer ok and cancel button actions -->
        <b-button variant="success" @click="ok()"> OK </b-button>
        <b-button variant="danger" @click="cancel()"> Cancel </b-button>
      </template>
    </b-modal>
  </div>
</template>

<script>
export default {
  name: "strip-cropping",
  data() {
    return {
      strip_cropping_temp: {
        pre_crop: this.$store.state.step2.Crop,
        crop: null,
        rate: 50,
      },
    };
  },
  props: {
    crops: {},
    strip_cropping: {},
  },
  model: {
    prop: "strip_cropping",
    event: "sc",
  },
  // methods
  methods: {
    returnSC() {
      this.$emit("sc", this.strip_cropping_temp);
    },

    // button event
    onOk() {
      console.log(this.strip_cropping_temp);
      this.returnSC();
    },

    onCancel() {
      this.strip_cropping_temp = {
        pre_crop: this.$store.state.step2.Crop,
        crop: null,
        rate: 50,
      };
      this.$emit("modal-cancel");
    },
  },
};
</script>
