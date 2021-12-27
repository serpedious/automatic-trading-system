<template>
<v-card class="ml-3 pa-8" width="500px" height="365px">
   <v-card-title class="pb-5">
    <h3>Automatic Trade</h3>
    </v-card-title>
    <h4>Border</h4>
    <h5>upper: 70</h5>
    <h5>lower: 30</h5>
     <tbody>
        <tr
        v-for="rsi of rsi_data" :key="rsi.msg"
        >
        <td><h4>{{ rsi.msg }} -> {{ rsi.value }}</h4></td>
        </tr>
    </tbody>
    <h4>Last Trade Notification</h4>
    <h5>{{ signal_data.product_code }}: (Price ->{{ signal_data.price }}, Side ->{{ signal_data.side }}, Time ->{{ signal_data.time }})</h5>
</v-card>
</template>

<script>
import axios from 'axios'
export default {
  data () {
    return {
      signal_data: {
          product_code: null,
          time: null,
          price: null,
          side: null,
      }, 
      signal: [],
      rsis: [],
      rsi_data: [
        { msg: 'second value', value: null},
        { msg: 'last value', value: null}
      ],
    }
  },
  mounted() {
    this.getRsi();
    this.getSignal();
    this.intervalFetchRsi();
  },
  methods: {
    getRsi: async function () {
      let res = await axios.get(process.env.API_BASE_URL + '/getrsi')
      this.rsis = res.data
      for (var i = 0; i < this.rsi_data.length; i++) {
        this.rsi_data[i].value = this.rsis[i]
      }
   },
   getSignal: async function() {
      let res = await axios.get(process.env.API_BASE_URL + '/getsignal') 
      this.signal = res.data
      this.signal_data.product_code = this.signal.signals[0].product_code
      this.signal_data.time = this.signal.signals[0].time
      this.signal_data.price = this.signal.signals[0].price
      this.signal_data.side = this.signal.signals[0].side
   },
   intervalFetchRsi: function () {
    setInterval(() => {    
     console.log("(((((((((((9((((99")
     this.getRsi();
     this.getSignal();
    }, 1000 * 60);    
   }
  }
}
</script>