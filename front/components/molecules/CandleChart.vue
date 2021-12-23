<template>
<div>
    <script type="text/javascript" src="https://www.gstatic.com/charts/loader.js"></script>
    <script type="text/javascript">
      google.charts.load('current', {'packages':['corechart']});
      google.charts.setOnLoadCallback(drawChart);
    </script>
 <v-card width="700">
    <div id="chart_div" style="width: 700px; height: 350px;"></div>
 </v-card>
</div>
</template>

<script>
export default {
  props: {
    crypto: '',
    side: '',
    amount: ''
  },
  data() {
    return {
      returnData: {
        crypto: this.crypto,
        side: this.side,
        amount: this.amount
      }
    }
  },
  mounted () {
    this.drawChart();
  },
  methods: {
    submit() {
      this.$emit('clickSubmit', this.returnData)
    },
    cancel() {
      this.$emit('clickSubmit', "")
    },
    drawChart() {
    var data = google.visualization.arrayToDataTable([
      ['Mon', 20, 28, 38, 45],
      ['Tue', 31, 38, 55, 66],
      ['Wed', 50, 55, 77, 80],
      ['Thu', 77, 77, 66, 50],
      ['Fri', 68, 66, 22, 15]
      // Treat first row as data as well.
    ], true);

    var options = {
      legend:'none'
    };

    var chart = new google.visualization.CandlestickChart(document.getElementById('chart_div'));

    chart.draw(data, options);
  }
  }
}
</script>
