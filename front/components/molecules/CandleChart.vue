<template>
<div>
    <script type="text/javascript" src="https://www.gstatic.com/charts/loader.js"></script>
    <script type="text/javascript">
      google.charts.load('current', {'packages':['corechart']});
      google.charts.setOnLoadCallback(drawChart);
    </script>
 <v-card width="700">
    <v-card-title class="pl-8">
      <h3>Bitcoin Chart</h3>
    </v-card-title>
    <div id="chart_div" style="width: 700px; height: 300px;"></div>
 </v-card>
</div>
</template>

<script>
import axios from 'axios'
export default {
  props: {
    crypto: '',
    side: '',
    amount: ''
  },
  data() {
    return {
      dataflame: null,
      returnData: {
        crypto: this.crypto,
        side: this.side,
        amount: this.amount
      }
    }
  },
  mounted () {
    this.getCandle();
    this.intervalFetchData();
  },
  methods: {
    getCandle: async function () {
      let res = await axios.get(process.env.API_BASE_URL + '/api/candle/')
      this.dataflame = res.data
      this.drawChart();
    },
    submit() {
      this.$emit('clickSubmit', this.returnData)
    },
    cancel() {
      this.$emit('clickSubmit', "")
    },
    drawChart() {
    var dataTable = [];
    for (var i = 0; i < this.dataflame.candles.length; i++) {
        var candleData = []
        candleData.push(this.dataflame.candles[i]["time"])
        candleData.push(this.dataflame.candles[i]["low"])
        candleData.push(this.dataflame.candles[i]["open"])
        candleData.push(this.dataflame.candles[i]["close"])
        candleData.push(this.dataflame.candles[i]["high"])
        dataTable.push(candleData)
        candleData = []
    } 
    
    var data = google.visualization.arrayToDataTable(dataTable, true);
    var options = {
      legend:'none',
      candlestick: {
          fallingColor: {strokeWidth: 0, fill: '#a52714' },
          risingColor: {strokeWidth: 0, fill: '#0f9d58' }, 
      }
    };

    var chart = new google.visualization.CandlestickChart(document.getElementById('chart_div'));

    chart.draw(data, options);
    this.intervalFetchData();
  },
   intervalFetchData: function () {
    setInterval(() => {    
      this.getCandle();
    }, 10000);    
  }
  }
}
</script>
