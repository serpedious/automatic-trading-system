<template>
<div>
<div class="d-flex justify-start">
<CandleChart/>
<Automatic/>
 </div>
 <div class="d-flex justify-start">
 <v-card lass=fluid style="margin: 2px; padding: 30px; width: 100%" v-if="balancelists.length">
    <v-card-title class="pb-5">
      <h3>My Assets</h3>
    </v-card-title>
    <v-simple-table dense>
      <template v-slot:default>
        <thead>
          <tr>
            <th class="text-left">
              Crypto Name
            </th>
            <th class="text-left">
              Balance
            </th>
            <th class="text-left">
              Price
            </th>
            <th class="text-left">
              Value(JPY)
            </th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="balance of balancelists" :key="balance.currency_code"
          >
            <td>{{ balance.crpto }}</td>
            <td>{{ balance.amount }}</td>
            <td>{{ balance.price }}</td>
            <td>{{ balance.value }}</td>
          </tr>
        </tbody>
      </template>
    </v-simple-table>
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
<TradeTable/>
</div>
</div>
</template>

<script>
import axios from 'axios'
import CandleChart from '../molecules/CandleChart'
import Automatic from '../molecules/Automatic.vue'
import TradeTable from './TableTrade'
export default {
  components: {
    CandleChart,
    Automatic,
    TradeTable,
  },
  name: 'TableChart',
  data() {
    return {
      balancelists: [],
    }
  },
  mounted() {
    this.apiBalance();
  },
  methods: {
    apiBalance: async function () {
      let res = await axios.get(process.env.API_BASE_URL + '/getmyassets')
      this.balancelists = res.data
    }
  }
}
</script>

