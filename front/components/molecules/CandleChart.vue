<template>
<v-card class="ml-3" width="500px" height="365px" v-if="dataTable.length !== 1">
   <v-card-title class="pb-5">
      <h3>Chart</h3>
    </v-card-title>
  <GChart
      type="CandlestickChart"
      :settings="chartsLib"
      :data="dataTable"
      :options="chartOptions"
  />
</v-card>
  <v-card class="ml-3" width="500px" height="365px" v-else>
    <v-card-actions class="justify-center pt-16">
       <v-progress-circular
         class="pt-16"
          indeterminate
          color="primary"
          ></v-progress-circular>
 </v-card-actions>
  </v-card>
</template>

<script>
import { GChart } from 'vue-google-charts'
import axios from 'axios'
export default {
  components: {
    GChart
  },
  data () {
    return {
      chartsLib: {packages: ['corechart']},
      dataTable: [["x", "Low", "Close", "Open", "High"]],
      chartOptions: {
        legend:'none',
        candlestick: {
          fallingColor: {strokeWidth: 0, fill: '#a52714' },
          risingColor: {strokeWidth: 0, fill: '#0f9d58' }, 
      }
      }
    }
  },
  mounted() {
    this.intervalFetchData();
  },
  methods: {
  drawChart() {
    this.dataTable = [["x", "Low", "Close", "Open", "High"]]; 
    for (var i = 0; i < this.dataflame.candles.length; i++) {
        var candleData = []
        candleData.push(this.dataflame.candles[i]["time"])
        candleData.push(this.dataflame.candles[i]["low"])
        candleData.push(this.dataflame.candles[i]["open"])
        candleData.push(this.dataflame.candles[i]["close"])
        candleData.push(this.dataflame.candles[i]["high"])
        this.dataTable.push(candleData)
        candleData = []
    } 
  },  
  getCandle: async function () {
      let res = await axios.get(process.env.API_BASE_URL + '/api/candle/')
      this.dataflame = res.data
      this.drawChart();
    },
  intervalFetchData: function () {
    setInterval(() => {    
      this.getCandle();
    }, 1000);    
  }
  }
}
</script>