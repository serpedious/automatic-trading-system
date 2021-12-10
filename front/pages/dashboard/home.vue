<template>
  <v-row>
    <v-col class="mx-auto pa-5" width="800px">
    <TableChart/>
    </v-col>
  </v-row>
</template>

<script>
import TableChart from '~/components/organisms/TableChart.vue'
// import GraphChart from '~/components/organisms/GraphChart.vue'
export default {
  components: {
    // GraphChart,
    TableChart,
  }
}
</script>
