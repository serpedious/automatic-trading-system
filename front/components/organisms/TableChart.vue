<template>
 <v-card class="mt-3 mr-3 pa-4">
    <v-card-title class="pb-5">
      <h3>Home</h3>
    </v-card-title>
    <v-simple-table dense>
      <template v-slot:default>
        <thead>
          <tr>
            <th class="text-left">
              Name
            </th>
            <!-- <th class="text-left">
              Bid
            </th>
            <th class="text-left">
              Ask
            </th> -->
            <th class="text-left">
              Balance
            </th>
            <th class="text-left">
              Transation
            </th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="balance of balancelists" :key="balance.currency_code"
          >
            <td>{{ balance.currency_code }}</td>
            <!-- <td>{{ item.buy }}</td>
            <td>{{ item.sell }}</td> -->
            <td>{{ balance.amount }}</td>
            <td>
              <v-btn
                depressed
                color="primary"
                x-small
              >
                BID
              </v-btn>
              <v-btn
                depressed
                color="error"
                x-small
              >
                ASK
              </v-btn>
            </td>
          </tr>
        </tbody>
      </template>
    </v-simple-table>
    <v-btn
      class="ma-2"
      @click="apiBalance"
      depressed
      color="primary"
    >
      SHOW
    </v-btn>
    <v-btn
      class="ma-2"
      depressed
      color="grey"
      @click="apiBalanceHide"
    >
      CLOSE
    </v-btn>
 </v-card>
</template>

<script>
import axios from 'axios'
export default {
  name: 'TableChart',
  data() {
    return {
      balancelists: [],
    }
  },
  methods: {
    apiBalance: async function () {
      let res = await axios.get(process.env.API_BASE_URL + '/balance')
      this.balancelists = res.data
    },
    apiBalanceHide: async function () {
      this.balancelists = ''
    }
  }
}
</script>
