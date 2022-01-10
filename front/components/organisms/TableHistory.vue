<template>
    <v-card style="margin: 2px; padding: 30px; width: 100%">
 <v-simple-table dense>
    <template v-slot:default>
      <thead>
        <tr>
          <th class="text-left">
            Trade Date
          </th>
          <th class="text-left">
            Product Code
          </th>
          <th class="text-left">
            Price
          </th>
          <th class="text-left">
            Crypto Amount
          </th>
          <th class="text-left">
            Trade Type
          </th>
          <th class="text-left">
            Fee
          </th>
          <th class="text-left">
            Currency
          </th>
          <th class="text-left">
            Currency Amount
          </th>
          <th class="text-left">
            Order ID
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="trade in trades"
          :key="trade.date"
        >
          <td>{{ trade.trade_date }}</td>
          <td>{{ trade.product_code }}</td>
          <td>{{ trade.price }}</td>
          <td>{{ trade.quantity}}</td>
          <td>{{ trade.trade_type }}</td>
          <td>{{ trade.commission }}</td>
          <td>{{ trade.currency_code }}</td>
          <td>{{ trade.amount}}</td>
          <td>{{ trade.order_id }}</td>
        </tr>
      </tbody>
    </template>
  </v-simple-table>
  </v-card>
</template>

<script>
import axios from 'axios'

export default {
  name: 'History',
  data () {
      return {
        trades: [
          // {
          //   date: '2020-20-28',
          //   type: 'SELL',
          //   price: 100,
          //   crypto: 'XRP',
          //   crp_amount: 9000,
          //   fee: 0.009,
          //   currency: 'JPY',
          //   amount: 5000,
          //   counter: "bitfyler",
          //   order_id: 'HDLOEE9987',
          // },
          // {
          //   date: '2020-20-28',
          //   type: 'SELL',
          //   price: 100,
          //   crypto: 'XRP',
          //   crp_amount: 9000,
          //   fee: 0.009,
          //   currency: 'JPY',
          //   amount: 5000,
          //   counter: "bitfyler",
          //   order_id: 'HDLOEE9987',
          // },
        ],
      }
    },
    mounted() {
      this.getBalanceHistory();
    },
    methods: {
      getBalanceHistory: async function () {
        let res = await axios.get(process.env.API_BASE_URL + '/getbalancehistory')
        this.trades = res.data
      }
  }
}
</script>
