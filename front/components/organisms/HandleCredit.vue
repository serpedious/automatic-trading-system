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
        <!-- <li v-for="item in deposit" :key="item.message">
          {{ item.amount }}
       </li> -->
        <tr
          v-for="item in deposit"
          :key="item.id"
          :class="[item.flag === 'MDP' ? 'red lighten-5' : 'blue lighten-5']"
        >
          <td>{{ item.message }} {{ item.amount }} {{item.price}} YEN {{item.product_code}} {{ item.side}}</td>
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
        withdraws: [],
      }
    },

   mounted() {
      this.getAlert();
    },
    methods: {
    getAlert: async function () {
      let res = await axios.get(process.env.API_BASE_URL + '/getalert')
      this.alerts = res.data
      this.deposit = this.alerts.deposit
      this.withdraws = this.alerts.withdrawals
      const len_with = this.withdraws.length;
      for (var i = 0; i < len_with; i++) {
        this.deposit.push(this.withdraws[i])
      }
      const len_all = this.deposit.length;
      for (var i = 0; i < len_all; i++) {
        this.deposit[i].flag = this.deposit[i].order_id.substring(0,3)
        if (this.deposit[i].flag === 'MDP') {
          this.deposit[i].message = "You deposited"
        } else {
          this.deposit[i].message = "You withdrawed"
        }
      }
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
