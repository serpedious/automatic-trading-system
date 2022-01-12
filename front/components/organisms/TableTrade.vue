<template>
  <v-card style="margin: 2px; padding: 30px; width: 100%; height: 50%;">
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
    <v-alert
      class="pa-3 ma-2"
      dense
      shaped
      outlined
      icon="mdi-currency-usd"
      color="purple"
    >
      {{ item.content }} --> <strong> Value: {{ item.value }} JPY </strong> 
    </v-alert>
          <div class="d-flex justify-center">
            <img class="mt-6" width="100px" height="100px" :src="item.img"/>
            <order-button :message="item.tab"/><br>
          </div>
        </v-card>
      </v-tab-item>
    </v-tabs-items>
  </v-card>
</template>

<script>
import axios from 'axios'
import OrderButton from './OrderButton.vue'
export default {
    components: {
      OrderButton,
    },
    name: 'TableTrade',
    data () {
        return {
        tab: null,
        balance: [],
        tickers: [],
        items: [
          { tab: 'BTC_JPY', content: 'Bitcoin Marketplace', value: null, jpy: null, img: "/bitcoin.svg" },
          { tab: 'XRP_JPY', content: 'XRP Marketplace', value: null, jpy: null, img: "/ripple.svg"},
          { tab: 'ETH_JPY', content: 'ETH Marketplace', value: null, jpy: null, img: "/ethereum.svg" },
          { tab: 'XLM_JPY', content: 'XLM Marketplace', value: null, jpy: null, img: "/stellar.svg"},
          { tab: 'MONA_JPY', content: 'MONA Marketplace', value: null, jpy: null, img: "/mona.svg" },
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


