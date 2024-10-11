<template>
  <div class="Manual">
    <b-card class="mt-5">
      <h1>CalBMP User Manual</h1>
      <p>(Version 1.0)</p>

      <p>Contact Information:</p>
      <p>
        Jingyuan Xue <a href="mailto: jiyxue@ucdavis.edu">jiyxue@ucdavis.edu</a>
      </p>
      <p>
        Developer
        <a href="mailto: Naqin0701@protonmail.com">
          Naqin0701@protonmail.com
        </a>
      </p>
      <p>
        Department of Land Air and Water Resources, University of California,
        Davis, 95616
      </p>

      <strong><i>Overview</i></strong>
      <p>
        CalBMP is a web-based interface for site-specific best management
        practices evaluations in California. It is designed as an advisory and
        regulatory modeling tool to mitigate pesticide loading leaving upland
        crop field for different stakeholders including growers, pesticide
        advisors, extension agents, modelers and policy makers, etc. The current
        version of the CalBMP interface enable users to do pesticide modeling
        and BMP evaluations based on site-specific field conditions in terms of
        location, topography, soil, weather, crop management, pesticide
        application, and BMP selections. The CalBMP interface has many great
        features including:
      </p>

      <ol>
        <li>No need to install offline softwares</li>
        <li>Skip the long-time preparation of input files</li>
        <li>No need to learn and understand the model output structures.</li>
        <li>
          Complete database: SSURGO soil database, GridMET weather database,
          pesticide usage (Pesticide Use Risk Evaluation (PURE) decision support
          system (DSS)) and pesticide properties database (PPDB).
        </li>
        <li>
          BMP scenarios: enabling simulation of the mitigation effectiveness for
          different BMP options including pesticide application reduction,
          pesticide application timing, cover crop, crop rotation and residue
          management, contour farming, parallel terracing, strip cropping.
        </li>
        <li>
          Visualization: enabling users to review different outputs including
          water depths, sediment, pesticide loading/concentration in runoff,
          erosion and volatilization in graphs and texts. Both time-series and
          event-based values could be checked in the front-end.
        </li>
        <li>
          Historical projects accessible: Projects could be saved after
          simulation, inputs and outputs files could be downloaded and saved for
          future use.
        </li>
      </ol>
      <p>Below are webpages demonstrating the steps for operating CalBMP</p>

      <strong><i>Welcome page</i></strong>
      <p>
        The first-time users will need to register and login to the interface.
        (Fig.1)
      </p>
      <!-- img of home page -->
      <img src="../assets/manualImg/welcome.png" alt="" />
      <p style="text-align: center">(a)</p>
      <!-- img of register page -->
      <img src="../assets/manualImg/register.png" alt="" />
      <p style="text-align: center">(b)</p>
      <p>
        <strong>Fig.1</strong> CalBMP page: welcome page, registration and login
      </p>
      <strong>
        <i>Make a run in Baseline scenario (Example with filled input info.)</i>
      </strong>
      <ol>
        <li>Configure the user field (Fig.2):</li>
        <ol>
          <li>
            Find your field location by selecting the county and zip code.
          </li>
          <li>Specifying the soil series and input the field area.</li>
        </ol>

        <li>Specify the crop management and pesticide application (Fig.3):</li>
        <ol>
          <li>
            Select crop, emergence date, maturity date, harvest date, irrigation
            type. (Put mouse on the date selection window, there will be pop-ups
            beside showing recommended time period guiding user’s selection.)
          </li>
          <li>
            Select application Numbers, pesticide name, application rate, date,
            method and location (Multiple applications could be applied in each
            run. Up to three chemicals could be evaluated in each run.)
          </li>
        </ol>

        <li>Visualization (Fig.4):</li>
        <ol>
          <li>
            Select water depths, sediment, pesticide loading in runoff, erosion
            and volatilization in the left side bar.
          </li>
          <li>
            Switch between “Water”, “Pesticide”, “Loading”, and “Concentration”
            taps on the top to review the outputs varying with time. (Drag the
            time slide bar below the graph to change the time period from year
            to week. Put the mouse on the curve to see the value and date of the
            index. Other legends could be grayed out to leave only one index in
            each graph for review.)
          </li>
          <li>
            Click on “summary” button on the right to review the summarized
            basic field and management information for this project. Values of
            event-based outputs could be reviewed in the table as well.
          </li>
        </ol>
      </ol>
      <!-- img of step1 -->
      <img src="../assets/manualImg/step1.png" alt="" />
      <p>
        <strong>Fig.2</strong> CalBMP page: finding the field. (Location, soil,
        area need to be specified by user to configure one field. The nearest
        weather station is automatically visualized in the map for user’s
        review.)
      </p>
      <!-- img of step2 -->
      <img src="../assets/manualImg/step2.png" alt="" />

      <p>
        <strong>Fig.3</strong> CalBMP page: specifying the crop management and
        pesticide application.
      </p>
      <!-- img of baseline -->
      <img src="../assets/manualImg/baseline/water/runoff.png" alt="" />
      <p style="text-align: center">(a)</p>
      <img src="../assets/manualImg/baseline/sediment/sediment.png" alt="" />
      <p style="text-align: center">(b)</p>
      <img src="../assets/manualImg/baseline/pesticide/total.png" alt="" />
      <p style="text-align: center">(c)</p>
      <img src="../assets/manualImg/baseline/pesticide/runoff.png" alt="" />
      <p style="text-align: center">(d)</p>
      <img src="../assets/manualImg/baseline/pesticide/erosion.png" alt="" />
      <p style="text-align: center">(e)</p>
      <img
        src="../assets/manualImg/baseline/concentration/concentration.png"
        alt=""
      />
      <p style="text-align: center">(f)</p>
      <img src="../assets/manualImg/baseline/summary.png" alt="" />
      <p style="text-align: center">(g)</p>

      <p style="text-align: center">
        <strong>Fig.4</strong> CalBMP page: outputs visualization in the
        baseline scenario.
      </p>

      <strong>
        <i>Make a run in BMP scenario (Example with filled input info.)</i>
      </strong>
      <ol>
        <li>
          Select BMP options (Fig.5a). (Single or multiple BMP options could be
          selected in each BMP scenario.)
        </li>
        <li>
          Input required parameters in pop-ups for selected BMP and click the
          “submit” button to run BMP scenario. Examples with filled information
          were provided for each BMP option (Fig.5b-h).
        </li>
        <li>Visualization</li>
        <ol>
          <li>
            Select “water depth”, “sediment”, “pesticide concentration in
            runoff”, “pesticide loading in runoff”, “pesticide loading in
            erosion” and “pesticide loading in volatilization” in the left side
            bar. Switch between “Water”, “Pesticide”, “Loading”, and
            “Concentration” in the top taps to review the outputs varying with
            time. (Same as visualization in baseline scenario)
          </li>
          <li>
            Click on “Try other BMP” button on the right to run another BMP
            scenario. (Optional)
          </li>
          <li>
            Click on “BMP comparison” button on the right to review the BMP
            effectiveness for single BMP or effectiveness comparison between BMP
            scenarios. (Fig. 6a-b)
          </li>
          <li>
            Click on “BMP summary” button on the right to review the table
            including summarized BMP scenario information and pesticide
            reduction percentage in main pathways. (Fig.6c)
          </li>
        </ol>
      </ol>
      <!-- img of bmp options -->
      <img src="../assets/manualImg/bmpOptions/totalOptions.png" alt="" />
      <p style="text-align: center">(a)</p>
      <img src="../assets/manualImg/bmpOptions/appReduction.png" alt="" />
      <p style="text-align: center">(b)</p>
      <img src="../assets/manualImg/bmpOptions/coverCrop.png" alt="" />
      <p style="text-align: center">(c)</p>
      <img src="../assets/manualImg/bmpOptions/rotation.png" alt="" />
      <p style="text-align: center">(d)</p>
      <img src="../assets/manualImg/bmpOptions/rotationPesticide.png" alt="" />
      <p style="text-align: center">(e)</p>
      <img src="../assets/manualImg/bmpOptions/contourFarming.png" alt="" />
      <p style="text-align: center">(f)</p>
      <img src="../assets/manualImg/bmpOptions/parallelTerracing.png" alt="" />
      <p style="text-align: center">(g)</p>
      <img src="../assets/manualImg/bmpOptions/stripCropping.png" alt="" />
      <p style="text-align: center">(h)</p>

      <p style="text-align: center">
        <strong>Fig.5</strong> CalBMP page: selecting BMP/BMPs in the BMP
        scenario.
      </p>
      <!-- img of comparision -->
      <img src="../assets/manualImg/comparison/single.png" alt="" />
      <p style="text-align: center">(a)</p>
      <img src="../assets/manualImg/comparison/five.png" alt="" />
      <p style="text-align: center">(b)</p>
      <p style="text-align: center">
        <strong>Fig.6</strong> CalBMP page: outputs visualization in the BMP
        scenario.
      </p>

      <strong><i>Help</i></strong>
      <p>
        Help contains this user manual and the developer’s contact information.
      </p>

      <b-row class="mt-5">
        <b-button variant="primary" class="mx-auto" @click="downloadManual">
          Download Manual
        </b-button>
      </b-row>
    </b-card>
  </div>
</template>

<script>
export default {
  methods: {
    downloadManual() {
      // window.location.href = "../assets/CalBMP_online_user_manual.docx";
      var link = document.createElement("a");
      // If you don't know the name or want to use
      // the webserver default set name = ''
      link.setAttribute("download", "manual.docx");
      link.href = "../assets/CalBMP_online_user_manual.docx";
      document.body.appendChild(link);
      link.click();
      link.remove();
    },
  },
};
</script>

<style>
img {
  border-radius: 8px;
  max-width: 70%;
  height: auto;
  margin-top: 2vh;
  display: block;
  margin-left: auto;
  margin-right: auto;
}
</style>
