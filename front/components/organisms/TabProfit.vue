<template>
  <v-card style="margin: 2px; padding: 30px; width: 100%;">
    <v-card-title class="pb-5">
      <h3>Profit</h3>
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
        <v-card flat color="" v-if="item.content !== null">
          <v-card-text class="text-center pt-8 text-h4 indigo--text"><b>{{ item.content }} YEN</b></v-card-text>
        </v-card>
        <v-card style="padding-bottom: 24px; width: 100%; height: 100%;" class="text-center pt-8" v-else>
          <v-progress-circular
            indeterminate
            color="primary"
          ></v-progress-circular>
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