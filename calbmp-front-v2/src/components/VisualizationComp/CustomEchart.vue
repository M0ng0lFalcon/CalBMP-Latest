<template>
  <div class="CustomEchart">
    <b-card border-variant="light">
      <div style="width: 100%; height: 50vh">
        <!-- echart body -->
        <div :id="echarId" style="width: 95%; height: 100%"></div>
      </div>
    </b-card>
  </div>
</template>

<script>
export default {
  name: "custom-echart",
  props: {
    echarId: String,
  },
};
</script>
