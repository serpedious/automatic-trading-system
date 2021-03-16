<template>
  <section class="container">
  <v-card class="mx-auto mt-5 pa-5" width="400px">
    <v-card-title class="pb-5">
      <h3>CSV</h3>
    </v-card-title>
    <v-tabs
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
        <v-card flat>
          <v-btn
          color="purple"
          dark
          class="ma-10"
        >
        DownLoad
        </v-btn>
        </v-card>
      </v-tab-item>
    </v-tabs-items>
  </v-card>
  </section>
</template>

<script>
  export default {
    name: 'Csv',
    data () {
      return {
        tab: null,
        items: [
          { tab: 'history', content: 'Tab 1 Content' },
          { tab: 'profit', content: 'Tab 2 Content' },
        ],
      }
    },
  }
</script>

<style scoped>
.container {
  min-height: 50vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}
</style>