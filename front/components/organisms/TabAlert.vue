<template>
  <section class="container">
    <v-card class="mx-auto mt-5 pa-5" width="700px">
      <v-card-title class="pb-5">
        <h3>Alert</h3>
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
       <TableAlert/>
      </v-tab-item>
    </v-tabs-items>
  </v-card>
</section>
</template>

<script>
import TableAlert from '~/components/organisms/TableAlert.vue'
  export default {
    name: 'Csv',
    components: {
      TableAlert,
    },
    data () {
      return {
        tab: null,
        items: [
          { tab: 'Twitter', content: 'Tab 1 Content' },
          { tab: 'Balance', content: 'Tab 2 Content' },
        ],
      }
    },
  }
</script>

<style scoped>
.container {
  min-height: 90vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}
</style>
