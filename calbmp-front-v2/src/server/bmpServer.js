import Vue from "vue";

export function GetCropNames() {
  console.log("get corpus names");
  return Vue.axios.get("/api/data/getSiteName").then((res) => {
    return res.data.crops;
  });
}

export function BmpScenario(formData) {
  const api = "/bmp/bmpScenario";
  console.log("bmp form:", JSON.stringify(formData));
  formData.step_1_params.ifBmp = true;
  formData.step_2_params.ifBmp = true;
  formData.step_1_params.wea_file_path = `${formData.bmp_id}/weatherData.wea`;
  formData.step_1_params.zts_file_path = `${formData.bmp_id}/bmp.zts`;

  // convert pesticide_app_reduction rate to int
  formData.pesticide_app_reduction.rate = parseInt(
    formData.pesticide_app_reduction.rate
  );
  return Vue.axios.post(api, formData);
}
