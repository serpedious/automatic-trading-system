<template>
  <section>
    <v-card class=fluid style="margin: 5px; padding: 30px; width: 100%">
      <v-card-title class="pb-3">
        <h3>Alert</h3>
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
          { tab: 'Activity', content: 'Tab 1 Content' },
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
