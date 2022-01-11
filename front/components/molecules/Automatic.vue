<template>
<v-card style="margin: 2px; padding: 30px; width: 100%; height: 50%;">
   <v-card-title>
    <h3>Automatic Trade</h3>
   </v-card-title>
   <v-alert
      dense
      shaped
      outlined
      text
      type="success"
    >
      RSI Border Condition: <strong>  Upper: 70 - Lower: 30 </strong>
    </v-alert>
    <div class="pl-7">
      <h5>RSI Values</h5>
     <v-simple-table dense>
    <template v-slot:default>
      <thead>
        <tr>
          <th class="text-left">
            Term
          </th>
          <th class="text-left">
            Value
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="rsi of rsi_data" :key="rsi.msg"
        >
          <td>{{ rsi.msg }}</td>
          <td>{{ rsi.value }}</td>
        </tr>
      </tbody>
    </template>
  </v-simple-table><br>
      <h5>Last Trade Notification</h5>
      <v-simple-table dense>
      <template v-slot:default>
        <thead>
          <tr>
            <th class="text-left">
              Crypto Name
            </th>
            <th class="text-left">
              Price
            </th>
            <th class="text-left">
              Side
            </th>
            <th class="text-left">
              Time
            </th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>{{ signal_data.product_code }}</td>
            <td>{{ signal_data.price }}</td>
            <td>{{ signal_data.side }}</td>
            <td>{{ signal_data.time }}</td>
          </tr>
        </tbody>
      </template>
    </v-simple-table>
    </div>
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
        { msg: 'Current the second value to last', value: null},
        { msg: 'Current last value', value: null}
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
     this.getRsi();
     this.getSignal();
    }, 1000);    
   }
  }
}
</script>
