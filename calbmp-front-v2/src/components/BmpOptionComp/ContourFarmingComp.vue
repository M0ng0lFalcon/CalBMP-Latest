<template>
  <div class="ContourFarmingComp">
    <b-modal
      id="contourFarming"
      centered
      title="Contour Farming"
      @ok="onOk"
      @cancel="onCancel"
      @close="onCancel"
      no-close-on-backdrop
      no-close-on-esc
    >
      <p>Contour Farming</p>
      <p>(Default scenario, no input needed.)</p>

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
  name: "contour-farming",
  methods: {
    onOk() {},

    onCancel() {
      this.$emit("modal-cancel");
    },
  },
};
</script>
