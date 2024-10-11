<template>
  <div class="Download">
    <b-table
      :items="inputFileProvider"
      :fields="fields"
      select-mode="multi"
      ref="selectableTable"
      selectable
      @row-selected="onRowSelected"
    >
      <!-- Example scoped slot for select state illustrative purposes -->
      <template #cell(selected)="{ rowSelected }">
        <template v-if="rowSelected">
          <span aria-hidden="true">&check;</span>
          <span class="sr-only">Selected</span>
        </template>
        <template v-else>
          <span aria-hidden="true">&nbsp;</span>
          <span class="sr-only">Not selected</span>
        </template>
      </template>
    </b-table>
    <p style="text-align: center">
      <b-button variant="primary" @click="selectAllRows" class="mr-3">
        Select all
      </b-button>
      <b-button variant="success" @click="clearSelected" class="mr-3">
        Clear selected
      </b-button>
      <b-button variant="outline-danger" @click="downloadFile" class="mr-3">
        Download
      </b-button>
      <b-button variant="danger" @click="get_vfsm_inputs" class="mr-3">
        VFS Input
      </b-button>
      <b-button variant="info" @click="get_vfsm_outputs"> VFS Output </b-button>
    </p>
  </div>
</template>

<script>
import { GetInputFiles, DownloadInputFile } from "@/server/resultServer";
import { GetVfsInputs, GetVfsOutputs } from "@/server/vfsmServer";

export default {
  name: "Download",
  data() {
    return {
      fields: ["selected", "filename"],
      items: [
        { filename: "test" },
        { filename: "test" },
        { filename: "test" },
        { filename: "test" },
        { filename: "test" },
      ],
      selected: [],
    };
  },

  computed: {
    file_list() {
      let res = [];
      this.selected.forEach((item) => {
        res.push(item.filename);
      });
      return res;
    },
  },

  methods: {
    onRowSelected(items) {
      this.selected = items;
    },
    selectAllRows() {
      this.$refs.selectableTable.selectAllRows();
    },
    clearSelected() {
      this.$refs.selectableTable.clearSelected();
    },
    downloadFile() {
      console.log("download file");
      DownloadInputFile(this.file_list, this.$store.state.created_time);
    },

    async inputFileProvider() {
      return await GetInputFiles(this.$store.state.created_time);
    },

    get_vfsm_inputs() {
      GetVfsInputs(this.$store.state.created_time);
    },

    get_vfsm_outputs() {
      GetVfsOutputs(this.$store.state.created_time);
    },
  },
};
</script>

<style scoped></style>
