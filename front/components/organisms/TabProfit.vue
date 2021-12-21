<template>
  <v-card class="mx-auto mt-2 pa-5" width="400px">
    <v-card-title class="pb-5">
      <h3>Profit</h3>
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
          <v-card-text class="text-h5">{{ item.content }} YEN</v-card-text>
        </v-card>
      </v-tab-item>
    </v-tabs-items>
  </v-card>
</template>

<script>
import axios from 'axios'
  export default {
    name: 'Profit',
    data () {
      return {
        tab: null,
        items: [
          { tab: 'All', content: null },
          { tab: 'Day', content: 'no value' },
          { tab: 'Week', content: 'no value' },
          { tab: 'Month', content: 'no value' },
        ],
      }
    },
    mounted() {
      this.getProfit();
    },
    methods: {
    getProfit: async function () {
      let res = await axios.get(process.env.API_BASE_URL + '/calcprofit')
      this.items[0]["content"] = res.data
    },
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