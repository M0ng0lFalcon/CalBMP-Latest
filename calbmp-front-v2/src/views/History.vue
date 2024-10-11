<template>
  <div class="History">
    <b-card title="History" class="mt-5" border-variant="light">
      <div class="HistoryListCard" v-for="item in historyList" :key="item.ID">
        <b-card class="mt-1">
          <b-row>
            <b-col lg="10">
              <h4>
                Project Name:
                <span class="infoText">{{ item.ProjectName }}</span>
              </h4>
              <p>
                <strong>County: </strong>
                <span class="infoText">{{ item.CountyName }}</span>
                <span>&nbsp;&nbsp;</span>

                <strong>Zip Code: </strong
                ><span class="infoText">{{ item.ZipCode }}</span>
                <span>&nbsp;&nbsp;</span>

                <strong>Soil: </strong
                ><span class="infoText">{{ item.Soil }}</span>
                <span>&nbsp;&nbsp;</span>

                <strong>Created Date: </strong>
                <span class="infoText">{{ item.CreatedDate }}</span>
              </p>
              <p>
                <strong>Muname: </strong>
                <span class="infoText">{{ item.Muname }}</span>
              </p>
            </b-col>
            <b-col class="my-auto">
              <b-button variant="primary" @click="toVisualization(item)">
                Visualization
              </b-button>
            </b-col>
          </b-row>
        </b-card>
      </div>
    </b-card>
  </div>
</template>

<script>
import { CheckHistory, ToVisualization } from "@/server/historyServer";

export default {
  data() {
    return {
      historyList: [],
    };
  },

  methods: {
    async toVisualization(HistoryInfo) {
      await ToVisualization(HistoryInfo);
      await this.$router.replace("/visualization");
    },
  },

  // mounted
  async mounted() {
    this.historyList = await CheckHistory();
  },
};
</script>

<style>
.infoText {
  color: #f48282;
}
</style>
