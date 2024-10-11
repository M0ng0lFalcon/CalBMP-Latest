<template>
  <div class="ResidueManagement">
    <b-row v-if="show === false">
      <b-col></b-col>
      <b-col>
        <b-button block variant="outline-success" @click="ShowResidueInput">
          <b-icon icon="plus-circle"></b-icon>
          Add Residue
        </b-button>
      </b-col>
      <b-col></b-col>
    </b-row>

    <b-row v-if="show === true">
      <b-col lg="9">
        <b-row>
          <b-col lg="4">
            <label :for="'cropUnitResidue_' + ResidueIndex">
              Residue Cover(t/ha):
            </label>
          </b-col>
          <b-col>
            <b-form-input
              :id="'cropUnitResidue_' + ResidueIndex"
              v-model="ResidueValueTemp"
              type="range"
              min="0.5"
              max="10"
              step="0.5"
              @change="returnRM"
            ></b-form-input>
            <b-tooltip
              :target="'cropUnitResidue_' + ResidueIndex"
              variant="primary"
            >
              {{ ResidueValueTemp != null ? ResidueValueTemp : 0 }}
            </b-tooltip>
          </b-col>
        </b-row>
      </b-col>
      <b-col>
        <b-button variant="outline-danger" @click="HideResidueInput">
          Delete Residue
        </b-button>
      </b-col>
    </b-row>
  </div>
</template>

<script>
export default {
  name: "residue-management",
  props: {
    ResidueShow: {
      type: Boolean,
      default: false,
    },
    ResidueIndex: Number,
    ResidueValue: Number,
  },
  model: {
    prop: "ResidueValue",
    event: "rm",
  },
  data() {
    return {
      ResidueValueTemp: 0,
      show: this.ResidueShow,
    };
  },
  methods: {
    HideResidueInput() {
      this.show = false;
      this.ResidueValueTemp = 0;
    },

    ShowResidueInput() {
      this.show = true;
    },

    returnRM() {
      this.$emit("rm", parseInt(this.ResidueValueTemp));
    },
  },
};
</script>
