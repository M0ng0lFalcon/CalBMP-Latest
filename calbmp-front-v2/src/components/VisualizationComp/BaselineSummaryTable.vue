<template>
  <div class="BaselineSummaryTable">
    <b-card style="max-width: 57vw">
      <b-row>
        <!-- basic info -->
        <b-card title="Basic Info" style="width: 100%" border-variant="light">
          <b-row>
            <b-col>
              <div v-for="(value, key) in basic_info" :key="key">
                <strong>{{ refactorKey(key) }} :</strong> {{ value }}
                <br />
              </div>
            </b-col>
            <b-col>
              <div v-for="n in basic_info_pesticide.pesticide.length" :key="n">
                <b-form-group :label="'Pesticide Application Number: ' + n">
                  <div v-for="(value, key) in basic_info_pesticide" :key="key">
                    <strong>{{ refactorKey(key) }} :</strong>
                    {{ value[n - 1] }}
                  </div>
                </b-form-group>
              </div>
            </b-col>
          </b-row>
        </b-card>
      </b-row>

      <b-row>
        <!-- Event-based outputs -->
        <b-table responsive hover :items="eventBasedOutputsProvider" busy.sync>
        </b-table>
      </b-row>
    </b-card>
  </div>
</template>

<script>
import { GetTextResult } from "@/server/resultServer";

export default {
  name: "baseline-table",
  data() {
    return {
      basic_info: {
        project_id: this.$store.state.step2.Crop,
        created_time: this.$store.state.created_time,
        county_name: this.$store.state.step1.county,
        zip_code: this.$store.state.step1.ZipCode,
        area: this.$store.state.step1.FieldSize,
        soil: this.$store.state.step1.muname,
        emergence_date: this.changeDateFormat(
          this.$store.state.step2.EmergenceDate
        ),
        maturity_date: this.changeDateFormat(
          this.$store.state.step2.MaturityDate
        ),
        harvest_date: this.changeDateFormat(
          this.$store.state.step2.HarvestDate
        ),
        irrigation_types: this.$store.state.step2.IrrigationTypeName,
      },
      basic_info_pesticide: {
        // The following is a list
        pesticide: this.$store.state.step2.Pesticide,
        application_date: this.changeDateFormat(this.$store.state.step2.date),
        rate: this.addRateUnit(this.$store.state.step2.amount),
        method: this.$store.state.step2.ApplicationEquipmentName,
        location: this.$store.state.step2.ApplicationMethodName,
      },
    };
  },

  // Methods
  methods: {
    refactorKey(key) {
      const keyLi = key.split("_");

      return keyLi
        .join(" ")
        .replace(/(^\w{1})|(\s+\w{1})/g, (letter) => letter.toUpperCase());
    },

    async eventBasedOutputsProvider() {
      const first_emergence = this.$store.state.first_emergence;
      const last_harvest = this.$store.state.last_harvest;
      const chemical_name = this.$store.state.step2.Pesticide;
      const created_time = this.$store.state.created_time;
      const field_size = parseFloat(this.$store.state.step1.FieldSize);
      return await GetTextResult(
        chemical_name,
        first_emergence,
        last_harvest,
        created_time,
        field_size
      );
      // return await GetEventBasedOutputs(
      //   chemical_name,
      //   first_emergence,
      //   last_harvest,
      //   created_time
      // );
    },

    changeDateFormat(org_date) {
      if (org_date instanceof Array) {
        let res = [];
        for (let it in org_date) {
          const li = org_date[it].split("-");
          res[it] = li[1] + "/" + li[2] + "/" + li[0];
        }
        return res;
      } else {
        const li = org_date.split("-");
        return li[1] + "/" + li[2] + "/" + li[0];
      }
    },

    addRateUnit(org_rate) {
      let res = [];
      for (let it in org_rate) {
        res.push(org_rate[it] + " lb/acre");
      }
      return res;
    },
  },
};
</script>
