import Vue from "vue";

export function Result2Visualization(created_time, min_id, max_id) {
  const api = `/vfsm/get_res_visualization?created_time=${created_time}&min_id=${min_id}&max_id=${max_id}`;
  return Vue.axios.get(api).then((res) => {
    return res.data.data.result;
  });
}

export function GetVegetation() {
  const api = "/vfsm/get_vegetation";
  return Vue.axios
    .get(api)
    .then((res) => {
      // console.log("vegetation", res.data.data.vegetation_list);
      return res.data.data.vegetation_list;
    })
    .catch((err) => {
      console.log(err);
      return false;
    });
}

export function GetVegetationModel(vegetation_name) {
  const api = `/vfsm/get_vegetation_model?vegetation=${vegetation_name}`;
  return Vue.axios
    .get(api)
    .then((res) => {
      return res.data.data.vegetation;
    })
    .catch((err) => {
      console.log(err);
      return false;
    });
}

export function RunVfsmModel(vfsm_param) {
  const api = "/vfsm/run";
  return Vue.axios.post(api, vfsm_param).then((res) => {
    console.log(res);
  });
}

export function GetVfsmProgress() {
  const api = "/vfsm/get_progress";
  return Vue.axios.get(api).then((res) => {
    return res.data.data;
  });
}

export function GetVfsInputs(created_time) {
  const api = `/vfsm/get_inputs?created_time=${created_time}`;
  Vue.axios
    .get(api, {
      responseType: "blob",
      header: {
        "Content-Disposition": `attachment; filename=VFS_input_files-${created_time}.zip`,
      },
    })
    .then((res) => {
      const elink = document.createElement("a");
      elink.style.display = "none";
      elink.href = URL.createObjectURL(res.data);
      elink.download = `VFS_input_files-${created_time}.zip`;
      document.body.appendChild(elink);
      elink.click();
      URL.revokeObjectURL(elink.href);
      document.body.removeChild(elink);
    });
}

export function GetVfsOutputs(created_time) {
  const api = `/vfsm/get_outputs?created_time=${created_time}`;
  Vue.axios
    .get(api, {
      responseType: "blob",
      header: {
        "Content-Disposition": `attachment; filename=VFS_output_files-${created_time}.zip`,
      },
    })
    .then((res) => {
      const elink = document.createElement("a");
      elink.style.display = "none";
      elink.href = URL.createObjectURL(res.data);
      elink.download = `VFS_output_files-${created_time}.zip`;
      document.body.appendChild(elink);
      elink.click();
      URL.revokeObjectURL(elink.href);
      document.body.removeChild(elink);
    });
}
