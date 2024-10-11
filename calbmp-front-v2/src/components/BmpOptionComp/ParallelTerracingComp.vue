<template>
  <div class="ParallelTerracingComp">
    <b-modal
      id="parallelTerracing"
      centered
      title="Parallel Terracing"
      @ok="onOk"
      @cancel="onCancel"
      @close="onCancel"
      no-close-on-backdrop
      no-close-on-esc
    >
      <b-form-group>
        <b-form-radio-group
          v-model="parallel_terracing_temp.type"
          :options="parallelTerracingOpts"
        >
        </b-form-radio-group>
      </b-form-group>

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
  name: "parallel-terracing",
  data() {
    return {
      // parallelTerracing options
      parallelTerracingOpts: [
        { value: "type1", text: "Graded channels sod outlets" },
        { value: "type2", text: "Steep backslope underground outlets" },
      ],

      parallel_terracing_temp: {
        type: null,
      },
    };
  },
  props: {
    parallel_terracing: {},
  },
  model: {
    prop: "parallel_terracing",
    event: "pt",
  },

  // methods
  methods: {
    returnPT() {
      this.$emit("pt", this.parallel_terracing_temp);
    },
    onOk() {
      this.returnPT();
    },
    onCancel() {
      this.parallel_terracing_temp.type = null;
      this.$emit("modal-cancel");
    },
  },
};
</script>
