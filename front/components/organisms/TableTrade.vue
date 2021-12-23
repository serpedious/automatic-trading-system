<template>
  <v-card class="mx-auto mt-3 pa-5" width="1300">
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
          <div class="d-flex justify-start">
          <img width="100px" height="100px" :src="item.img"/>
          <div class="text-h4 mx-auto">
            Value: {{ item.value }} JPY
          </div>
            <OrderButton/>
          <div class="mx-auto mt-3 mb-3">
            Your Balance: Japanese Yen {{ item.jpy }} JPY
          </div>
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
          { tab: 'BTC/JPY', content: 'Bitcoin Marketplace', value: null, jpy: null, img: "/bitcoin.svg" },
          { tab: 'XRP/JPY', content: 'XRP Marketplace', value: null, jpy: null, img: "/ripple.svg"},
          { tab: 'ETH/JPY', content: 'ETH Marketplace', value: null, jpy: null, img: "/ethereum.svg" },
          { tab: 'XLM/JPY', content: 'XLM Marketplace', value: null, jpy: null, img: "/stellar.svg"},
          { tab: 'MONA/JPY', content: 'MONA Marketplace', value: null, jpy: null, img: "/mona.svg" },
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


