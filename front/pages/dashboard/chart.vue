<template>
  <div class="home">
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
            v-for="balance of balancelists" v-bind:key="balance.amount"
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
    <button @click="apiBalance">show balance</button>
    <button @click="apiBalanceHide">hide balance</button>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'Memo',
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

<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
button {
  margin: 10px 0;
  padding: 10px;
}
</style>