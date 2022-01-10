<template>


<v-row dense>
  <v-col xs="1" sm="10" md="10" lg="6" xl="2">
   <TabCsv/>
  </v-col>
  <v-col xs="1" sm="10" md="10" lg="6" xl="2">
   <TabProfit/> 
  </v-col>
   <v-col xs="1" sm="10" md="10" lg="12" xl="2">
   <TableHistory/>
  </v-col>
</v-row>
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
