<template>

<v-row dense>
  <v-col xs="1" sm="10" md="10" lg="6" xl="2">
    <CandleChart/>
    <MyAssets/>
  </v-col>
  <v-col xs="1" sm="10" md="10" lg="6" xl="2">
    <Automatic/>
    <TradeTable/>
  </v-col>
</v-row>


</template>

<script>
import CandleChart from '../molecules/CandleChart'
import Automatic from '../molecules/Automatic.vue'
import MyAssets from '../molecules/MyAssets.vue'
import TradeTable from './TableTrade'
export default {
  components: {
    CandleChart,
    Automatic,
    MyAssets,
    TradeTable,
  },
  name: 'TableChart',
  data() {
    return {
    }
  }
}
</script>

