<template>
  <div class="ComponentsOpts">
    <!-- water balance components place : start -->
    <b-form-group label="Water Depth">
      <b-form-checkbox
        v-for="item in waterBalanceComponents"
        :key="item.name"
        v-model="item.status"
        value="1"
        unchecked-value="0"
        switch
      >
        {{ item.name }}
      </b-form-checkbox>
    </b-form-group>

    <b-form-group label="Sediment">
      <b-form-checkbox
        v-for="item in sedimentComponent"
        :key="item.name"
        v-model="item.status"
        value="1"
        unchecked-value="0"
        switch
      >
        {{ item.name }}
      </b-form-checkbox>
    </b-form-group>

    <!-- make a horizontal rule -->
    <hr />

    <!-- pesticide components place : start -->
    <b-form-group label="Pesticide">
      <!-- concentration options -->
      <b-form-checkbox
        v-for="item in concentrationOpt"
        :key="item.name"
        v-model="item.status"
        value="1"
        unchecked-value="0"
        switch
      >
        {{ item.name }}
      </b-form-checkbox>
      <!-- loading options -->
      <b-form-checkbox
        v-for="item in pesticideComponentsOpt"
        :key="item.name"
        v-model="item.status"
        value="1"
        unchecked-value="0"
        switch
      >
        {{ item.name }}
      </b-form-checkbox>
    </b-form-group>
  </div>
</template>

<script>
import {
  WaterBalanceComponents,
  PesticideComponentsOpt,
  SedimentComponent,
  ConcentrationOpt,
} from "@/variable/visualizationVar";
export default {
  name: "components-opts",
  data() {
    return {
      waterBalanceComponents: WaterBalanceComponents,
      pesticideComponentsOpt: PesticideComponentsOpt,
      sedimentComponent: SedimentComponent,
      concentrationOpt: ConcentrationOpt,
    };
  },
};
</script>
