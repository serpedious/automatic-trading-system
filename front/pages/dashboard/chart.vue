<template>
  <!-- <v-container> -->
    <v-row>
      <v-col md="8">
        <GraphChart/>
      </v-col>
      <v-col md="4">
      <TableChart/>
      </v-col>
    </v-row>
  <!-- </v-container> -->
</template>

<script>
import TableChart from '~/components/organisms/TableChart.vue'
import GraphChart from '~/components/organisms/GraphChart.vue'
export default {
  components: {
    GraphChart,
    TableChart,
  }
}
</script>
