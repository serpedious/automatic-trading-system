<template>
  <v-card class="mx-auto mt-2 pa-6" width="400px">
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
          <v-card-text >{{ item.content }}</v-card-text>
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
          { tab: 'Day', content: 'Tab 1 Content' },
          { tab: 'Week', content: 'Tab 2 Content' },
          { tab: 'Month', content: 'Tab 3 Content' },
          { tab: 'All', content: 'Tab 4 Content' },
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