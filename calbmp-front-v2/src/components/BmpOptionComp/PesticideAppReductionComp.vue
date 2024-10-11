<template>
  <div class="PesticideAppReductionComp">
    <b-modal
      id="pesticideAppReduction"
      centered
      title="Pesticide Application Reduction"
      @ok="onOk"
      @cancel="onCancel"
      @close="onCancel"
      no-close-on-backdrop
      no-close-on-esc
    >
      <b-row>
        <b-col>
          <b-form-group label="App. Reduction Rate(%)">
            <!-- input reduction rate -->
            <b-form-input
              id="pesticideAppReductionInput"
              v-model="pesticide_app_reduction_temp.rate"
              type="range"
              min="0"
              max="100"
              step="1"
            ></b-form-input>

            <!-- show reduction rate -->
            <b-tooltip target="pesticideAppReductionInput" variant="primary">
              {{ pesticide_app_reduction_temp.rate }}
            </b-tooltip>
          </b-form-group>
        </b-col>
      </b-row>

      <template #modal-footer="{ ok, cancel }">
        <!-- Emulate built in modal footer ok and cancel button actions -->
        <b-button variant="danger" @click="cancel()"> Cancel </b-button>
        <b-button variant="success" @click="ok()"> OK </b-button>
      </template>
    </b-modal>
  </div>
</template>

<script>
export default {
  name: "pesticideApp-reduction",
  data() {
    return {
      pesticide_app_reduction_temp: {
        rate: 0,
      },
    };
  },
  props: {
    pesticide_app_reduction: {},
  },
  model: {
    prop: "pesticide_app_reduction",
    event: "par",
  },

  // methods
  methods: {
    returnPAR() {
      // convert rate to integer
      const tmpRate = this.pesticide_app_reduction_temp.rate;
      this.pesticide_app_reduction_temp.rate = parseInt(tmpRate);
      this.$emit("par", this.pesticide_app_reduction_temp);
    },

    onOk() {
      this.returnPAR();
    },

    onCancel() {
      this.$emit("modal-cancel");
    },
  },
};
</script>
