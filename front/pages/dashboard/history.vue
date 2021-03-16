<template>
  <div>
    <div class="pa-10">
      <TabProfit/>
    </div>
    <TableHistory/>
  </div>
</template>

<script>
import TableHistory from '~/components/organisms/TableHistory.vue'
import TabProfit from '../../components/organisms/TabProfit.vue'
export default {
  components: {
  TableHistory,
  TabProfit
  }
}
</script>

<style>
</style>

