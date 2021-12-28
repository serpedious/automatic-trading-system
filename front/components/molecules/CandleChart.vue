<template>
<v-card lass=fluid style="margin: 2px; padding: 30px; width: 100%" v-if="dataTable.length !== 1">
   <v-card-title class="pb-5">
      <h3>Chart</h3>
    </v-card-title>
    <div class="pb-2">
    <v-btn small @click="changeDuration('1s');">1s</v-btn>
    <v-btn small @click="changeDuration('1m');">1m</v-btn>
    <v-btn small @click="changeDuration('1h');">1h</v-btn>
    </div>
  <GChart
      type="CandlestickChart"
      :settings="chartsLib"
      :data="dataTable"
      :options="chartOptions"
  />
</v-card>
  <v-card lass=fluid style="margin: 5px; padding: 30px; width: 100%" v-else>
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
      },
      config: {
        candlestick:{
            product_code: 'BTC_JPY',
            duration: '1m',
            limit: 365,
            numViews: 5,
        },
        dataTable: {
            index : 0,
            value: null
        },
        rsi: {
            enable: false,
            indexes: {'up': 0, 'value': 0, 'down': 0},
            period: 14,
            up: 70,
            values: [],
            down: 30
        },
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
    var params = {
        "product_code": this.config.candlestick.product_code,
        "duration": this.config.candlestick.duration,
        "limit": this.config.candlestick.limit,
      }
      let res = await axios.get(process.env.API_BASE_URL + '/api/candle/', {params: params})
      this.dataflame = res.data
      this.drawChart();
    },
  changeDuration(s){
    this.config.candlestick.duration = s;
    this.getCandle();
  },
  intervalFetchData: function () {
    setInterval(() => {    
      this.getCandle();
    }, 1000);    
  }
  }
}
</script>