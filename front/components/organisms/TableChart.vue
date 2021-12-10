<template>
 <v-card class="mt-3 mr-3 pa-4">
    <v-card-title class="pb-5">
      <h3>Home</h3>
    </v-card-title>
    <v-simple-table dense>
      <template v-slot:default>
        <thead>
          <v-btn
            class="ma-2"
            @click="apiBalance"
            depressed
            color="primary"
          >
            SHOW
          </v-btn>
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
            <td>{{ balance.currency_code }}</td>
            <td>{{ balance.amount }}</td>
          </tr>
        </tbody>
      </template>
    </v-simple-table>
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
    }
  }
}
</script>
