<template>
  <v-card class="mx-auto mt-3 pa-3" width="1100px">
    <v-card-title>
      <h3>Trade</h3>
    </v-card-title>
    <v-tabs
      centered
      fixed-tabs
      v-model="tab"
      background-color="indigo"
      dark
      slider-color="purple"
    >
      <v-tab
        v-for="item in items"
        :key="item.tab"
      >
        {{ item.tab }}
      </v-tab>
    </v-tabs>


    <v-tabs-items v-model="tab">
      <v-tab-item
        v-for="item in items"
        :key="item.tab"
      >
      
        <v-card flat color="">
          <v-card-text class="text-h5">{{ item.content }}</v-card-text>
          <div class="text-h4 mx-auto">
            {{ item.value }} JPY
          </div>
            <td>
              <v-btn
                depressed
                color="error"
                x-large
              >
                SELL
              </v-btn>
              <v-btn
                depressed
                color="success"
                x-large
              >
                BUY
              </v-btn>
            </td>
            <div class="mx-auto mt-3 mb-3">
              Japanese Yen {{ item.jpy }} JPY
            </div>
        </v-card>
      </v-tab-item>
    </v-tabs-items>
  </v-card>
</template>

<script>
import axios from 'axios'
export default {
    name: 'TableTrade',
    data () {
        return {
        tab: null,
        balance: [],
        tickers: [],
        items: [
          { tab: 'BTC/JPY', content: 'Bitcoin Marketplace', value: null, jpy: null },
          { tab: 'XRP/JPY', content: 'XRP Marketplace', value: null, jpy: null },
          { tab: 'ETH/JPY', content: 'ETH Marketplace', value: null, jpy: null },
          { tab: 'XLM/JPY', content: 'XLM Marketplace', value: null, jpy: null },
          { tab: 'MONA/JPY', content: 'MONA Marketplace', value: null, jpy: null },
        ],
        }
    },
    mounted () {
        this.apiCalc();
        // this.intervalFetchData();
    },
    methods: {
    apiCalc: async function () {
        let res_assets = await axios.get(process.env.API_BASE_URL + '/getmyassets')
        this.tickers = res_assets.data
        let res = await axios.get(process.env.API_BASE_URL + '/balance')
        this.balance = res.data
        console.log(this.tickers)
        let tickers_len = this.tickers.length;
        for (let i = 0; i < tickers_len; i++) {
            this.items[i]["value"] = this.tickers[i].price
            this.items[i]["jpy"] = this.balance[0]["available"]
        }
    },
    // intervalFetchData: function () {
    //     setInterval(() => {    
    //         this.apiCalc();
    //         }, 1000000);    
    // }
   }
  }
</script>

<style scoped>
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}
</style>

