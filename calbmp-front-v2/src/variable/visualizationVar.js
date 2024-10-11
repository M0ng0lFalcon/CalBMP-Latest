// static variable for baseline scenario and bmp scenario
export var WaterBalanceComponents = {
  item0: { name: "Runoff", status: 0, value: "RUNF" },
  item1: { name: "Precipitation", status: 0, value: "PRCP" },
  item2: { name: "Irrigation", status: 0, value: "IRRG" },
};

export var SedimentComponent = {
  item0: { name: "Eroded Soilds", status: 0, value: "ESLS" },
};

export var PesticideComponentsOpt = {
  item0: { name: "Loading in runoff", status: 0, value: "RFLX" },
  item1: { name: "Loading in erosion", status: 0, value: "EFLX" },
  // value: "FPVL,VFLX",
  item2: { name: "Loading in volatilization", status: 0, value: "VFLX" },
};

export var ConcentrationOpt = {
  item0: { name: "Concentration in runoff", status: 0, value: "concentration" },
};
