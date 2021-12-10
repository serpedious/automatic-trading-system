<template>
  <v-card class="mx-auto mt-3 pa-3" width="1100px">
    <v-card-title>
      <h3>Trade</h3>
    </v-card-title>
    <v-tabs
      centered
      fixed-tabs
      v-model="tab"
      background-color="primary"
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
          <v-card-text >{{ item.content }}</v-card-text>
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
        </v-card>
      </v-tab-item>
    </v-tabs-items>
  </v-card>
</template>

<script>
  export default {
    name: 'Csv',
    data () {
      return {
        tab: null,
        items: [
          { tab: 'BTC/JPY', content: 'Bitcoin Marketplace' },
          { tab: 'XRP/JPY', content: 'Bitcoin Marketplace' },
          { tab: 'ETH/JPY', content: 'Bitcoin Marketplace' },
          { tab: 'DOT/JPY', content: 'Bitcoin Marketplace' },
          { tab: 'XTZ/JPY', content: 'Bitcoin Marketplace' },
          { tab: 'XLM/JPY', content: 'Bitcoin Marketplace' },
          { tab: 'XEM/JPY', content: 'Bitcoin Marketplace' },
          { tab: 'BAT/JPY', content: 'Bitcoin Marketplace' },
          { tab: 'ETC/JPY', content: 'Bitcoin Marketplace' },
          { tab: 'LTC/JPY', content: 'Bitcoin Marketplace' },
          { tab: 'BCH/JPY', content: 'Bitcoin Marketplace' },
          { tab: 'MONA/JPY', content: 'Bitcoin Marketplace' },
          { tab: 'LSK/JPY', content: 'Bitcoin Marketplace' },
        ],
      }
    },
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