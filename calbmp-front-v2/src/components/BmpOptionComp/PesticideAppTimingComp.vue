<template>
  <div class="PesticideAppTimingComp">
    <b-modal
      id="pesticideAppTiming"
      centered
      title="Pesticide Application Timing"
      @ok="onOk"
      @cancel="onCancel"
      @close="onCancel"
      no-close-on-backdrop
      no-close-on-esc
    >
      <p>Adjust your pesticide application time to aviod heavy rain!</p>
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
  name: "pesticideApp-timing",
  methods: {
    onOk() {},

    onCancel() {
      this.$emit("modal-cancel");
    },
  },
};
</script>
