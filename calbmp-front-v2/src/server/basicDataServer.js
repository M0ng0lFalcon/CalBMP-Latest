import Vue from "vue";

// ---------------- user step 1 ----------------
// 1. get all soil
export function GetAllSoil() {
  const api = "/data/get_all_soil";
  return Vue.axios.get(api).then((res) => {
    return res.data.data.soil_list;
  });
}

export function GetZipcodeBySoil(mukey, cokey) {
  const api = `/data/get_zipcode_by_soil?mukey=${mukey}&cokey=${cokey}`;
  return Vue.axios.get(api).then((res) => {
    console.log(res);
  });
}

// 1. get county names
export function GetCouties() {
  const api = "/data/getCounties";
  return Vue.axios.get(api).then((response) => {
    console.log("[*] getting counties");

    let counties = [];

    const countiesData = response.data.data.counties;
    for (let i = 0; i < countiesData.length; i += 1) {
      const newData = {
        value: countiesData[i],
        text: countiesData[i],
      };
      counties.push(newData);
    }

    return counties;
  });
}

// 2. get zip code
export function getZipcode(county) {
  const api = `/data/getZipcode?CountyName=${county}`;

  return Vue.axios.get(api).then((response) => {
    console.log("[*] getting zipcode");

    let zipCodes = [];

    const zipCodeData = response.data.data.zipCode;
    for (let i = 0; i < zipCodeData.length; i += 1) {
      const newData = {
        value: zipCodeData[i],
        text: zipCodeData[i],
      };
      zipCodes.push(newData);
    }

    return zipCodes;
  });
}

// 3. get comp names
export function getCompnameMukeyCokey(zip_code) {
  const api = `/data/getCompnameMukeyCokey?zip_code=${zip_code}`;
  return Vue.axios.get(api).then((res) => {
    console.log("[*] get Compname Mukey Cokey");
    const data = res.data.data;
    const compname_mukey_cokey = data.compname_mukey_cokey;

    let resData = [];
    for (let it in compname_mukey_cokey) {
      const obj = {
        value: {
          compName: compname_mukey_cokey[it].compname,
          mukey: compname_mukey_cokey[it].mukey,
          cokey: compname_mukey_cokey[it].cokey,
          muname: compname_mukey_cokey[it].muname,
        },
        text: compname_mukey_cokey[it].muname,
      };
      resData.push(obj);
    }
    return resData;
  });
}

export function getStation(ZipCode) {
  const api = `/data/getStation?ZipCode=${ZipCode}`;
  return Vue.axios.get(api).then((response) => {
    console.log("[*] getting station");

    const data = response.data.data;
    const stationData = data.Station;

    const Log = stationData.Log;
    const Lat = stationData.Lat;

    const resData = {
      ClimateId: stationData.ClimateId,
      log: Log,
      lat: Lat,
      center: [Lat, Log],
      zoom: 8,
    };
    return resData;
  });
}

// ---------------- user step 2 ----------------

export function GetCrops() {
  const api = `/data/getCropName`;
  console.log("[*] getting crops");
  return Vue.axios.get(api).then((res) => {
    console.log(res);
    let cropsRes = new Array();
    const crops = res.data.data.CropNames;
    for (let i = 0; i < crops.length; i += 1) {
      const newData = {
        value: crops[i],
        text: crops[i],
      };
      cropsRes.push(newData);
    }
    return cropsRes;
  });
}

export function GetPesticide() {
  const api = "/data/getPesticide";
  return Vue.axios.get(api).then((res) => {
    console.log("[*] getting pesticides");
    const pesticides = res.data.data.Pesticides;
    let PesticideRes = new Array();
    for (let i = 0; i < pesticides.length; i += 1) {
      const newData = {
        value: pesticides[i],
        text: pesticides[i],
      };
      PesticideRes.push(newData);
    }
    return PesticideRes;
  });
}

export function submitStep1(UserStep1) {
  const api = "/input/step1";
  console.log(JSON.stringify(UserStep1));
  return Vue.axios.post(api, UserStep1).then((res) => {
    console.log(res);
  });
}

export function submitStep2(UserStep2) {
  const api = "/input/step2";
  console.log(JSON.stringify(UserStep2));
  return Vue.axios.post(api, UserStep2);
}
