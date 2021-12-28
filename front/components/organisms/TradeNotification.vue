<template>
 <v-simple-table dense>
    <template v-slot:default>
      <thead>
        <tr>
          <th class="text-left">
            Transaction
          </th>
          <th class="text-left">
            Date
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="item in deposit"
          :key="item.id"
        >
          <td>{{item.product_code}} is recommended to trade with {{ item.side}} positon</td>
          <td>{{ item.event_date}}  {{item.time}} </td>
        </tr>
      </tbody>
    </template>
  </v-simple-table>
</template>

<script>
import axios from 'axios'
export default {
  name: 'History',
  data () {
      return {
        alerts: [],
        deposit: [],
      }
    },
    
   mounted() {
      this.getAlert();
    },
    methods: {
    getAlert: async function () {
      let resp = await axios.get(process.env.API_BASE_URL + '/getsignalall') 
      this.signal = resp.data
      var size = Object.values(this.signal)[0].length
      for (var i = 0; i < size; i++) {
        console.log(Object.values(this.signal)[0][i].message)
        this.deposit.push(Object.values(this.signal)[0][i])
      }
      console.log(this.deposit)
      this.deposit = this.deposit.sort(function(a,b){
        return new Date(b.event_date) - new Date(a.event_date);
      });
    },
  //   getSignal: async function() {
  //     let res = await axios.get(process.env.API_BASE_URL + '/getsignalall') 
  //     this.signal = res.data
  //     var size = Object.values(this.signal)[0].length
  //     for (var i = 0; i < size; i++) {
  //       console.log(Object.values(this.signal)[0][i])
  //       this.deposit.push(Object.values(this.signal)[0][i])
  //     }
  //     console.log(this.deposit)
  //  },
    }
}
</script>

<style scoped>
.container {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}
</style>
