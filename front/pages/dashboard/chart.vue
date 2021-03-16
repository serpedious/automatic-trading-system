<template>
  <div>
    <div class="pa-10">
      <GraphChart/>
    </div>
    <TableChart/>
  </div>
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

<style>
</style>

