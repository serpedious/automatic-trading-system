<template>
  <div>
    <div class="d-flex justify-start">
      <TabCsv/>
      <TabProfit/>
    </div>
    <TableHistory/>
  </div>
</template>

<script>
import TableHistory from '~/components/organisms/TableHistory.vue'
import TabProfit from '../../components/organisms/TabProfit.vue'
import TabCsv from '../../components/organisms/TabCsv.vue'
export default {
  components: {
  TableHistory,
  TabProfit,
  TabCsv
  }
}
</script>
