<template>
  <div class="BMPSummaryTable">
    <b-table
      striped
      hover
      :fields="fields"
      :items="BmpSummaryProvider"
      busy.sync
    >
      <template #cell(show_details)="row">
        <b-button
          v-if="row.item.BMP_Options === 'VFS Mode'"
          variant="success"
          size="sm"
          @click="row.toggleDetails"
          class="mr-2"
        >
          {{ row.detailsShowing ? "Hide" : "Show" }}
        </b-button>
      </template>

      <template #row-details="row">
        <b-table striped hover :fields="vfsm_fields" :items="vfsm_items">
        </b-table>
        <b-button variant="danger" size="sm" @click="row.toggleDetails">
          Hide Details
        </b-button>
      </template>
    </b-table>
  </div>
</template>

<script>
import { GetBmpSummaryData } from "@/server/resultServer";

export default {
  name: "bmp-table",
  data() {
    return {
      fields: [
        { key: "BMP_Scenario", label: "BMP Scenario" },
        { key: "BMP_Options", label: "BMP Options" },
        {
          key: "pesticide_reduction_in_Runoff",
          label: "Pesticide Reduction in Runoff (%)",
        },
        {
          key: "pesticide_reduction_in_Erosion",
          label: "Pesticide Reduction in Erosion (%)",
        },
        {
          key: "pesticide_reduction_in_Volatilization",
          label: "Pesticide Reduction in Volatilization (%)",
        },
        {
          key: "show_details",
          label: "Show Details",
        },
      ],

      vfsm_fields: [
        {
          key: "event",
          label: "Event",
        },
        {
          key: "date",
          label: "Date",
        },
        {
          key: "sediment_trapping_efficiency",
          label: "Sediment Trapping Efficiency (%)",
        },
        {
          key: "runoff_inflow_reduction",
          label: "Runoff Inflow Reduction (%)",
        },
        {
          key: "pesticide_reduction",
          label: "Pesticide Reduction (%)",
        },
        {
          key: "pesticide_trapped_in_vfs",
          label: "Pesticide Trapped In VFS (mg)",
        },
      ],

      vfsm_items: [
        {
          event: 1,
          date: "2012/4/13",
          sediment_trapping_efficiency: 78.87,
          runoff_inflow_reduction: 66.3,
          pesticide_reduction: 99.299,
          pesticide_trapped_in_vfs: 1.04135,
        },
      ],
    };
  },

  // methods
  methods: {
    async BmpSummaryProvider() {
      const bmp_cnt = this.$store.state.echartList.length;
      const harvest = this.$store.state.last_harvest;
      return GetBmpSummaryData(bmp_cnt, harvest);
    },
  },
};
</script>
